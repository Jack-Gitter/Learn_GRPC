package Weather

import "context"


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
