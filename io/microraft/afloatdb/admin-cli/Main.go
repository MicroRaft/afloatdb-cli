package main

import (
	"flag"
	"log"

	admin "io.microraft/afloatdb/admin"
	// raftendpoint "io.microraft/afloatdb/raftendpoint"
)

const getReportOpName = "report"
const takeSnapshotOpName = "snapshot"

var (
	serverAddr  = flag.String("server", "localhost:6701", "The server address in the format of host:port")
	op          = flag.String("op", getReportOpName, "operation to execute. one of "+getReportOpName+"|"+takeSnapshotOpName)
	timeoutSecs = flag.Int64("timeoutSecs", 10, "timeout in seconds")
)

func main() {
	flag.Parse()
	client, err := admin.NewClient(*serverAddr, *timeoutSecs)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	switch *op {
	case getReportOpName:
		report, endpoints, err := client.GetRaftNodeReport()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Report:", *report)
		log.Println("Endpoints:", *endpoints)
	case takeSnapshotOpName:
		report, err := client.TakeSnapshot()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Report:", *report)
	default:
		log.Fatalln("Invalid op:", *op)
	}

}
