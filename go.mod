module github.com/NpoolPlatform/go-service-app-template

go 1.16

require (
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211019090045-7235f7a363f4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20211015200801-69063c4bb744 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211019152133-63b7e35f4404
	google.golang.org/grpc v1.41.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
