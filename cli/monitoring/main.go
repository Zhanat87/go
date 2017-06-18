/*
go run cli/monitoring/main.go
cd cli/monitoring && go build && cd ../..
cli/monitoring/monitoring
https://stackoverflow.com/questions/12486691/how-do-i-get-my-golang-web-server-to-run-in-the-background
https://askubuntu.com/questions/38126/how-to-redirect-output-to-screen-as-well-as-a-file
/root/zhanat.site/cli/monitoring/monitoring &
echo $! | tee ../pid.txt
 */
package main

import (
	"net/http"
	"github.com/joho/godotenv"
	"github.com/Zhanat87/go/helpers"
	"os"
	"github.com/pkg/errors"
	"io/ioutil"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"time"
	"golang.org/x/net/context"
	grpc_local "github.com/Zhanat87/go/grpc"
	"google.golang.org/grpc"
	"io"
	"fmt"
)

// go run ~/go/src/github.com/Zhanat87/go/server.go
func checkApiServerLiveness() {
	resp, err := http.Get(os.Getenv("API_BASE_URL") + "liveness")
	helpers.SendErrorIfNeed(err)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if (string(body) != "api works!") {
		helpers.LogError(errors.New("golang service not responsed!"))
	}
}

// go run ~/go/src/github.com/Zhanat87/golang-socketio-server/main.go
func checkSocketIoServerLiveness() {
	// connect to server, you can use your own transport settings
	conn, err := gosocketio.Dial(
		gosocketio.GetUrl(os.Getenv("DOMAIN_NAME"), 5000, false),
		transport.GetDefaultWebsocketTransport(),
	)
	defer conn.Close()
	helpers.SendErrorIfNeed(err)

	conn.On("livenessMessage", func(c *gosocketio.Channel, msg int) string {
		if msg != 2 {
			helpers.LogError(errors.New("socket io service bad response!"))
		}
		return "OK"
	})
	conn.Emit("liveness", 1)
	time.Sleep(1 * time.Second)
}

// go run ~/go/src/github.com/Zhanat87/golang-grpc-protobuf-server/main.go
func checkGrpcServerLiveness() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(os.Getenv("GRPC_SERVER"), grpc.WithInsecure())
	helpers.SendErrorIfNeed(err)
	defer conn.Close()

	client := grpc_local.NewGrpcServiceClient(conn)
	// calling the streaming API
	stream, err := client.CheckGrpcServerLiveness(context.Background(), &grpc_local.EmptyRequest{})
	helpers.SendErrorIfNeed(err)

	var data *grpc_local.LivenessResponse
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			helpers.SendErrorIfNeed(err)
		}
		data = resp
	}

	if data.Msg != "gRPC server works!" {
		helpers.LogError(errors.New("gRPC service bad response!"))
	}
}

func main() {
	for {
		err := godotenv.Load()
		helpers.SendErrorIfNeed(err)

		checkApiServerLiveness()

		checkSocketIoServerLiveness()

		checkGrpcServerLiveness()

		// report about app liveness
		msg := "All services and microservices works fine!"
		_, err = helpers.SendEmail(os.Getenv("MAIL_TO_ADDRESS"), "Zhanat site liveness", msg)
		helpers.SendErrorIfNeed(err)

		fmt.Println(msg)

		time.Sleep(1 * time.Hour)
	}
}
