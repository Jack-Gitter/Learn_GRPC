package Weather

import (
	"context"

	"google.golang.org/grpc"
)


type WeatherServer struct { 
    UnimplementedWeatherReporterServer
}

func (s WeatherServer) GetTodaysWeather(context context.Context, request *WeatherRequest) (*WeatherReply, error) {
    return &WeatherReply{
        Temperature: 90,
        Humidity: 50,
        Cloudy: true,
    }, nil
}

func (s WeatherServer) GetPastTwoDaysWeatherServerStream(request *TwoDaysWeatherRequest, stream grpc.ServerStreamingServer[WeatherServerStreamReply]) error {

    stream.Send(&WeatherServerStreamReply{
        UV: 5,
        Percipitation: 25,
    })

    stream.Send(&WeatherServerStreamReply{
        UV: 10,
        Percipitation: 20,
    })

    return nil

}
