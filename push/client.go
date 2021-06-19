package push

import (
	"context"
	"time"

	"github.com/zoripong/simple-chat-service/public"
	"github.com/zoripong/simple-chat-service/rpc"
	"google.golang.org/grpc"
)

type PushClient struct {
	Connection *grpc.ClientConn
	rpc.PushServiceClient
}

type PushSender interface {
	SendPushInstantly()
}

func NewPushClient() *PushClient {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		public.GetErrorLogger().Printf("did not connect: %v\n", err)
	}
	serviceClient := rpc.NewPushServiceClient(conn)
	client := PushClient{
		Connection:        conn,
		PushServiceClient: serviceClient,
	}
	return &client
}

func (client *PushClient) SendPushInstantly(message *PushMessage) bool {
	defer client.Connection.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.PushInstantly(
		ctx,
		message.ToPushInstantlyReuest(),
	)

	if err != nil {
		public.GetErrorLogger().Printf("could not send push: %v\n", err)
		return false
	}

	return r.Success
}
