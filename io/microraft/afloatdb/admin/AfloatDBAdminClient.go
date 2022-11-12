package admin

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"context"
)

type AfloatDBAdminClient struct {
	connection *grpc.ClientConn
	client     AdminServiceClient
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewClient(serverAddr string, timeoutSecs int64) (*AfloatDBAdminClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	log.Println("Connecting to", serverAddr)

	connection, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, err
	}

	client := new(AfloatDBAdminClient)
	client.connection = connection
	client.client = NewAdminServiceClient(connection)
	client.ctx, client.cancel = context.WithTimeout(context.Background(), time.Duration(timeoutSecs)*time.Second)

	return client, nil
}

func (client *AfloatDBAdminClient) Close() {
	client.cancel()
	client.connection.Close()

	log.Println("Stopped client.")
}

func (client *AfloatDBAdminClient) GetRaftNodeReport() (*RaftNodeReportProto, *map[string]string, error) {
	resp, err := client.client.GetRaftNodeReport(client.ctx, &GetRaftNodeReportRequest{})

	if err != nil {
		return nil, nil, err
	}

	return resp.Report, &resp.EndpointAddress, nil
}

func (client *AfloatDBAdminClient) TakeSnapshot() (*RaftNodeReportProto, error) {
	resp, err := client.client.TakeSnapshot(client.ctx, &TakeSnapshotRequest{})

	if err != nil {
		return nil, err
	}

	return resp.Report, nil
}

// func (client *AfloatDBClient) Size() (int32, error) {
// 	resp, err := client.client.Size(client.ctx, &SizeRequest{})
// 	if err != nil {
// 		return 0, err
// 	}

// 	return resp.GetSizeResult().Size, nil
// }

// func (client *AfloatDBClient) PutString(key, value string) (*Val, error) {
// 	resp, err := client.client.Put(client.ctx, &PutRequest{
// 		Key: key,
// 		Val: &Val{Val: &Val_Str{value}},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp.GetPutResult().OldVal, nil
// }

// func (client *AfloatDBClient) SetString(key, value string) (bool, error) {
// 	resp, err := client.client.Set(client.ctx, &SetRequest{
// 		Key: key,
// 		Val: &Val{Val: &Val_Str{value}},
// 	})

// 	if err != nil {
// 		return false, err
// 	}

// 	return resp.GetSetResult().OldValExisted, nil
// }

// func (client *AfloatDBClient) Get(key string) (*Val, error) {
// 	resp, err := client.client.Get(client.ctx, &GetRequest{Key: key})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp.GetGetResult().Val, nil
// }
