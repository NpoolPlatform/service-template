package echo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-app-template/message/npool"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func run(wg *sync.WaitGroup) {
	defer wg.Done()

	l, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Println(err)
	}

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		)),
	)

	go func() {
		npool.RegisterServiceExampleServer(server, &Server{})
		grpc_prometheus.EnableHandlingTimeHistogram()
		grpc_prometheus.Register(server)
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			http.ListenAndServe(":9999", nil) // nolint
		}()
		log.Println(server.Serve(l))
	}()

	mux := runtime.NewServeMux()
	svc := http.Server{
		Addr:    ":9091",
		Handler: mux,
	}
	go func() {
		time.Sleep(time.Second * 10)
		if err := svc.Shutdown(context.Background()); err != nil {
			log.Panic(err)
		}
		server.GracefulStop()
	}()
	if err := npool.RegisterServiceExampleHandlerFromEndpoint(context.Background(), mux, "127.0.0.1:9090", []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		log.Panic(err)
	}
	if err := mux.HandlePath(http.MethodGet, "/healthz", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("PONG")) // nolint
	}); err != nil {
		log.Panic(err)
	}
	if err := svc.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panic(err)
	}
	log.Println("server done")
}

func grpct(wg *sync.WaitGroup) {
	defer wg.Done()
	done := make(chan struct{})
	tick := time.NewTicker(time.Second)
	go func() {
		time.Sleep(time.Second * 5)
		tick.Stop()
		done <- struct{}{}
	}()
loop:
	for {
		select {
		case <-tick.C:
		case <-done:
			close(done)
			break loop
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		conn, err := grpc.DialContext(ctx, "127.0.0.1:9090", grpc.WithInsecure(),
			grpc.WithBlock(),
		)
		if err != nil {
			cancel()
			continue
		}
		client := npool.NewServiceExampleClient(conn)
		out, err := client.Echo(ctx, &npool.StringMessage{Value: "hello world"})
		if err != nil {
			cancel()
			continue
		}
		cancel()
		fmt.Println("from grpc: ", out.GetValue())
	}
	log.Println("grpc done")
}

func httpt(wg *sync.WaitGroup) {
	defer wg.Done()
	done := make(chan struct{})
	tick := time.NewTicker(time.Second)
	go func() {
		time.Sleep(time.Second * 5)
		tick.Stop()
		done <- struct{}{}
	}()
loop:
	for {
		select {
		case <-tick.C:
		case <-done:
			close(done)
			break loop
		}

		buf := bytes.Buffer{}
		info, err := json.Marshal(struct {
			Value string `json:"value"`
		}{Value: "hello grpc"})
		if err != nil {
			break loop
		}
		_, err = buf.Write(info)
		if err != nil {
			break loop
		}

		resp, err := http.Post("http://127.0.0.1:9091/v1/echo", "application/json", &buf)
		if err != nil {
			break loop
		}

		rest, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			break loop
		}
		resp.Body.Close()
		fmt.Println("from http: ", string(rest))

		err = healthz()
		if err != nil {
			break loop
		}
	}
	log.Println("http done")
}

func healthz() error {
	resp, err := http.Get("http://127.0.0.1:9091/healthz")
	if err != nil {
		return err
	}

	rest, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return err
	}
	resp.Body.Close()
	fmt.Println("from http healthz: ", string(rest))
	return nil
}

func TestEcho(t *testing.T) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go run(&wg)
	go grpct(&wg)
	go httpt(&wg)
	t.Log("success")
	wg.Wait()
}
