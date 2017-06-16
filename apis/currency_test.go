package apis

import (
	"testing"
	grpc_local "github.com/Zhanat87/go/grpc"
	"google.golang.org/grpc"
	"github.com/stretchr/testify/assert"
)

// go test apis/currency_test.go
func TestExchangeRates(t *testing.T) {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	assert.Nil(t, err)

	res := getExchangeRates(grpc_local.NewGrpcServiceClient(conn), &grpc_local.EmptyRequest{})

	assert.Equal(t, 4, len(res.Data))
}
