package main

import (
	"GRPC_Weather_API/Weather"
	"context"
	"fmt"
	"io"
	"log"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)


func main() {


    creds, err := credentials.NewClientTLSFromFile("Certs/ca.cert", "localhost.backloop.dev")

    if err != nil {
        log.Fatal(err)
    }

    perRPC := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(getOauthAccessToken()) }

    opts := []grpc.DialOption{
        grpc.WithPerRPCCredentials(perRPC),
        grpc.WithTransportCredentials(creds),
    }

    conn, err := grpc.NewClient("0.0.0.0:8080", opts...) 

    if err != nil {
        log.Fatal(err)
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

func getOauthAccessToken() *oauth2.Token {
    return &oauth2.Token{
        AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJKYWNrIEdpdHRlciIsIm5hbWUiOiJUZXN0aW5nIEF1dGggdG9rZW5zISIsImlhdCI6OTUxNjIzOTAyMn0.tOf9G7H3QtMlbulEyL99HjqkEycd0jIHNRHODbCNxZg",
    }
}
