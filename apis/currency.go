package apis

import (
	"github.com/go-ozzo/ozzo-routing"
	"io"
	"log"
	"golang.org/x/net/context"
	grpc_local "github.com/Zhanat87/go/grpc"
	"os"
	"google.golang.org/grpc"
)

// getExchangeRates calls the RPC method GetExchangeRates of CurrencyServer
func getExchangeRates(client grpc_local.GrpcServiceClient, emptyRequest *grpc_local.EmptyRequest) *grpc_local.ExchangeRatesResponse {
	// calling the streaming API
	stream, err := client.GetExchangeRates(context.Background(), emptyRequest)
	if err != nil {
		log.Fatalf("Error on get exchange rates: %v", err)
	}
	var data *grpc_local.ExchangeRatesResponse
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetExchangeRates(_) = _, %v, test123", client, err)
		}
		data = resp
		log.Printf("data: %v", data)
	}

	return data
}

func ExchangeRates() routing.Handler {
	return func(c *routing.Context) error {
		// Set up a connection to the gRPC server.
		conn, err := grpc.Dial(os.Getenv("GRPC_SERVER"), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		return c.Write(getExchangeRates(grpc_local.NewGrpcServiceClient(conn), &grpc_local.EmptyRequest{}))
	}
}
