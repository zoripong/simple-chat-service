package push

import (
	"context"
	"log"
	"time"

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
		log.Fatalf("did not connect: %v", err)
	}
	serviceClient := rpc.NewPushServiceClient(conn)
	client := PushClient{
		Connection:        conn,
		PushServiceClient: serviceClient,
	}
	return &client
}

func (client *PushClient) ToPushInstantlyReuest(message *PushMessage) *rpc.PushInstantlyRequest {
	return &rpc.PushInstantlyRequest{
		Content:    message.Content,
		SenderId:   message.SenderId,
		ReceiverId: message.ReceiverId,
		SendAt:     message.SendAt.Unix(),
	}
}

func (client *PushClient) SendPushInstantly(message *PushMessage) bool {
	defer client.Connection.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.PushInstantly(
		ctx,
		client.ToPushInstantlyReuest(message),
	)

	if err != nil {
		log.Fatalf("could not send push: %v", err)
		return false
	}

	return r.Success
}
