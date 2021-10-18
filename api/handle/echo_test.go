package handle

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-app-template/message/npool"
	"google.golang.org/grpc"
)

func TestEcho(t *testing.T) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		l, err := net.Listen("tcp", "0.0.0.0:8080")
		if err != nil {
			log.Fatal(err)
		}

		server := grpc.NewServer()
		go func() {
			time.Sleep(time.Second * 10)
			server.GracefulStop()
		}()

		Register(server)
		log.Println(server.Serve(l))
	}()

	go func() {
		done := make(chan struct{})
		defer wg.Done()
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
			conn, err := grpc.DialContext(ctx, "127.0.0.1:8080", grpc.WithInsecure(),
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
			fmt.Println(out.GetValue())
		}
		log.Println("grpc client close")
	}()

	t.Log("success")
	wg.Wait()
}
