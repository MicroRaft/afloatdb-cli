package kv

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"context"
)

type AfloatDBClient struct {
	connection *grpc.ClientConn
	client     KVServiceClient
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewClient(serverAddr string, timeoutSecs int64) (*AfloatDBClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	log.Println("Connecting to", serverAddr)

	connection, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, err
	}

	client := new(AfloatDBClient)
	client.connection = connection
	client.client = NewKVServiceClient(connection)
	client.ctx, client.cancel = context.WithTimeout(context.Background(), time.Duration(timeoutSecs)*time.Second)

	return client, nil
}

func (client *AfloatDBClient) Close() {
	client.cancel()
	client.connection.Close()

	log.Println("Stopped client.")
}

func (client *AfloatDBClient) Size() (int32, error) {
	resp, err := client.client.Size(client.ctx, &SizeRequest{})
	if err != nil {
		return 0, err
	}

	return resp.GetSizeResult().Size, nil
}

func (client *AfloatDBClient) PutString(key, value string) (*Val, error) {
	resp, err := client.client.Put(client.ctx, &PutRequest{
		Key: key,
		Val: &Val{Val: &Val_Str{value}},
	})
	if err != nil {
		return nil, err
	}

	return resp.GetPutResult().OldVal, nil
}

func (client *AfloatDBClient) SetString(key, value string) (bool, error) {
	resp, err := client.client.Set(client.ctx, &SetRequest{
		Key: key,
		Val: &Val{Val: &Val_Str{value}},
	})

	if err != nil {
		return false, err
	}

	return resp.GetSetResult().OldValExisted, nil
}

func (client *AfloatDBClient) Get(key string) (*Val, error) {
	resp, err := client.client.Get(client.ctx, &GetRequest{Key: key})
	if err != nil {
		return nil, err
	}

	return resp.GetGetResult().Val, nil
}
