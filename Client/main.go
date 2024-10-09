package main

import (
	"GRPC_Weather_API/Weather"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)


func main() {


    creds, err := credentials.NewClientTLSFromFile("Certs/ca.cert", "localhost.backloop.dev")

    if err != nil {
        log.Fatal(err)
    }

    conn, err := grpc.NewClient("0.0.0.0:8080", grpc.WithTransportCredentials(creds))

    if err != nil {
        log.Fatal(err)
        log.Fatal("could not connect to server on port 8080")
    }

    defer conn.Close()

    client := Weather.NewWeatherReporterClient(conn)

    resp, err := client.GetTodaysWeather(context.Background(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp)

}
