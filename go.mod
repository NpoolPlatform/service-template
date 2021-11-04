module github.com/NpoolPlatform/go-service-app-template

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211031094530-5ab1b9965285
	github.com/go-resty/resty/v2 v2.7.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.0.0-20211015200801-69063c4bb744 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/genproto v0.0.0-20211027162914-98a5263abeca
	google.golang.org/grpc v1.41.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
