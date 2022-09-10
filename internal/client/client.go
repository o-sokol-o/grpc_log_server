package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	grpc_log "github.com/o-sokol-o/grpc_log_server/pkg/domain"
)

type GrpcClient struct {
	client grpc_log.GrpcLogServiceClient
	conn   *grpc.ClientConn
	msg_id int64
}

func New(target string) (*GrpcClient, error) {

	// Set up a connection to the server.
	// Target as "localhost:9000"
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	gc := GrpcClient{
		client: grpc_log.NewGrpcLogServiceClient(conn),
		conn:   conn,
	}

	err = gc.Send(0, 0, 0)
	if err != nil {
		return nil, errors.New("grpc ping error: " + err.Error())
	}

	return &gc, nil
}

// Action    LogRequest_Actions     	`protobuf:"varint,1,opt,name=action,proto3,enum=grpc_log.LogRequest_Actions" json:"action,omitempty"`
// Entity    LogRequest_Entities    `protobuf:"varint,2,opt,name=entity,proto3,enum=grpc_log.LogRequest_Entities" json:"entity,omitempty"`
// EntityId  int64                  `protobuf:"varint,3,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
// Timestamp *timestamppb.Timestamp
func (c *GrpcClient) Send(Action grpc_log.LogRequest_Actions,
	Entity grpc_log.LogRequest_Entities,
	EntityId int64) error {

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	msg := grpc_log.LogRequest{
		Action:    Action,
		Entity:    Entity,
		EntityId:  EntityId,
		LogMsgId:  c.msg_id,
		Timestamp: timestamppb.Now(),
	}

	r, err := c.client.Log(ctx, &msg)
	if err != nil {
		return err // log.Fatalf("could not send: %v", err)
	}
	fmt.Printf("Msg %d sended. Respons %d\n", c.msg_id, r.LogMsgId)
	c.msg_id++
	return nil
}

func (c *GrpcClient) Close() {
	c.conn.Close()
}
