package apis

import (
	"testing"
	grpc_local "github.com/Zhanat87/go/grpc"
	"google.golang.org/grpc"
	"github.com/stretchr/testify/assert"
)

// go test apis/currency_test.go
func TestWeatherInfo(t *testing.T) {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	assert.Nil(t, err)

	request := &grpc_local.WeatherRequest{
		Latitude:  43.241688,
		Longitude: 76.877550,
	}
	res := getWeatherInfo(grpc_local.NewGrpcServiceClient(conn), request)

	assert.NotNil(t, res)

	assert.NotNil(t, res.Temp)
	assert.NotNil(t, res.Humidity)
	assert.NotNil(t, res.Pressure)
}
