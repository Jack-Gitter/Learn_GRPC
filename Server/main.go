package main

import (
	"GRPC_Weather_API/Weather"
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {

    listener, err := net.Listen("tcp", "localhost.backloop.dev:8080")

    if err != nil {
        log.Fatal("unable to listen on port 8080")
    }

    creds, err := credentials.NewServerTLSFromFile("Certs/server.cert", "Certs/server.pem")

    if err != nil {
        log.Fatal(err)
    }

    opts := []grpc.ServerOption{
        grpc.UnaryInterceptor(PrintJWTInterceptor),
        grpc.Creds(creds),
    }

    grpcServer := grpc.NewServer(
        opts...
    )
    
    Weather.RegisterWeatherReporterServer(grpcServer, Weather.WeatherServer{})

    grpcServer.Serve(listener)

}

func PrintJWTInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return nil, errors.New("hi")
    }

    fmt.Println(md["authorization"])

    return handler(ctx, req)


}
