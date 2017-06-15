package apis

import (
	"io"
	"log"
	"strconv"
	"github.com/go-ozzo/ozzo-routing"
	"golang.org/x/net/context"
	grpc_local "github.com/Zhanat87/go/grpc"
	"os"
	"google.golang.org/grpc"
)

// getWeatherInfo calls the RPC method GetWeatherInfo of CurrencyServer
func getWeatherInfo(client grpc_local.GrpcServiceClient, weatherRequest *grpc_local.WeatherRequest) *grpc_local.WeatherResponse {
	// calling the streaming API
	stream, err := client.GetWeatherInfo(context.Background(), weatherRequest)
	if err != nil {
		log.Fatalf("Error on get exchange rates: %v", err)
	}
	var data *grpc_local.WeatherResponse
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetWeatherInfo(_) = _, %v", client, err)
		}
		log.Printf("data: %v", data)
		data = resp
	}

	return data
}

func WeatherInfo() routing.Handler {
	return func(c *routing.Context) error {
		lat, _ := strconv.ParseFloat(c.Param("lat"), 64)
		lon, _ := strconv.ParseFloat(c.Param("lon"), 64)
		request := &grpc_local.WeatherRequest{
			Latitude:  lat,
			Longitude: lon,
		}

		// Set up a connection to the gRPC server.
		conn, err := grpc.Dial(os.Getenv("GRPC_SERVER"), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		return c.Write(getWeatherInfo(grpc_local.NewGrpcServiceClient(conn), request))
	}
}
