package main

import (
	"flag"
	"log"

	kv "io.microraft/afloatdb/kv"
)

var (
	serverAddr  = flag.String("server", "localhost:6701", "The server address in the format of host:port")
	key         = flag.String("key", "", "key to insert. must be non-empty for put|set|get")
	value       = flag.String("value", "", "value to insert. must be non-empty for put|set")
	op          = flag.String("op", "get", "operation to execute. one of put|set|get|size")
	timeoutSecs = flag.Int64("timeoutSecs", 10, "timeout in seconds")
)

const sizeOpName = "size"
const putOpName = "put"
const setOpName = "set"
const getOpName = "get"

func main() {
	flag.Parse()
	client, err := kv.NewClient(*serverAddr, *timeoutSecs)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	switch *op {
	case sizeOpName:
		size, err := client.Size()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Size:", size)
	case putOpName:
		validateNonEmpty(*key, "key")
		validateNonEmpty(*value, "value")
		oldVal, err := client.PutString(*key, *value)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Old value:", oldVal)
	case setOpName:
		validateNonEmpty(*key, "key")
		validateNonEmpty(*value, "value")
		oldValueExists, err := client.SetString(*key, *value)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Has old value:", oldValueExists)
	case getOpName:
		validateNonEmpty(*key, "key")
		value, err := client.Get(*key)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Value:", value)
	default:
		log.Fatalln("Invalid op:", *op)
	}

}

func validateNonEmpty(value, component string) {
	if value == "" {
		log.Fatalf("Missing %v parameter", component)
	}
}
