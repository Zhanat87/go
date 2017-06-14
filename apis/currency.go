package apis

import (
	"github.com/go-ozzo/ozzo-routing"
	"github.com/Zhanat87/go/grpc/currency"
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/Zhanat87/go/grpc/currency"
	"os"
)

// getExchangeRates calls the RPC method GetExchangeRates of CurrencyServer
func getExchangeRates(client pb.CurrencyClient, emptyRequest *pb.EmptyRequest) *currency.ExchangeRatesResponse {
	// calling the streaming API
	stream, err := client.GetExchangeRates(context.Background(), emptyRequest)
	if err != nil {
		log.Fatalf("Error on get exchange rates: %v", err)
	}
	var data *currency.ExchangeRatesResponse
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetExchangeRates(_) = _, %v", client, err)
		}
		log.Printf("data: %v", data)
		data = resp
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
		// Creates a new CurrencyClient
		client := pb.NewCurrencyClient(conn)

		return c.Write(getExchangeRates(client, &pb.EmptyRequest{}))
	}
}
