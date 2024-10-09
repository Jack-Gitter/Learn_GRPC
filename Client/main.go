package main

import (
	"GRPC_Weather_API/Weather"
	"context"
	"fmt"
	"io"
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

    resp, err := client.GetTodaysWeather(context.Background(), &Weather.WeatherRequest{})

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp)

    readStream, err := client.GetPastTwoDaysWeatherServerStream(context.Background(), &Weather.TwoDaysWeatherRequest{})

    if err != nil {
        log.Fatal(err)
    }

    for {
        weatherReport, err := readStream.Recv()
        if err == io.EOF {
            break;
        }
        if err != nil {
            log.Fatal(err)
        }
        log.Println(weatherReport)
    }

}
