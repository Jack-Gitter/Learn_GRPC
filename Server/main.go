package main

import (
	"GRPC_Weather_API/Weather"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

    grpcServer := grpc.NewServer(
        grpc.Creds(creds),
    )
    
    Weather.RegisterWeatherReporterServer(grpcServer, Weather.WeatherServer{})

    grpcServer.Serve(listener)

}
