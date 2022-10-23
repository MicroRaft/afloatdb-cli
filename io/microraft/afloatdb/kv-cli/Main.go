package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	kv "io.microraft/afloatdb/kv"
)

var (
	serverAddr = flag.String("server", "localhost:6701", "The server address in the format of host:port")
	key        = flag.String("key", "", "key to insert. must be non-empty for put|set|get")
	value      = flag.String("value", "", "value to insert. must be non-empty for put|set")
	op         = flag.String("op", "get", "operation to execute. one of put|set|get|size")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	log.Println("Connecting to", *serverAddr)

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("Could not connect to server! Error: %v", err)
	}
	defer conn.Close()
	client := kv.NewKVRequestHandlerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch *op {
	case "size":
		resp, err := client.Size(ctx, &kv.SizeRequest{})
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Size:", resp.GetSizeResult().Size)
	case "put":
		validateNonEmpty(*key, "key")
		validateNonEmpty(*value, "value")
		resp, err := client.Put(ctx, &kv.PutRequest{
			Key: *key,
			Val: &kv.Val{Val: &kv.Val_Str{*value}},
		})
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Old value:", resp.GetPutResult().OldVal)
	case "set":
		validateNonEmpty(*key, "key")
		validateNonEmpty(*value, "value")
		resp, err := client.Set(ctx, &kv.SetRequest{
			Key: *key,
			Val: &kv.Val{Val: &kv.Val_Str{*value}},
		})

		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Has old value:", resp.GetSetResult().OldValExisted)
	case "get":
		validateNonEmpty(*key, "key")
		resp, err := client.Get(ctx, &kv.GetRequest{Key: *key})
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Value:", resp.GetGetResult().Val)
	default:
		log.Fatalln("Invalid op:", *op)
	}

}

func validateNonEmpty(value, component string) {
	if value == "" {
		log.Fatalf("Missing %v parameter", component)
	}
}
