module io.microraft/afloatdb/admin-cli

go 1.19

require (
	io.microraft/afloatdb/admin v1.0.0
	io.microraft/afloatdb/raftendpoint v1.0.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect
	google.golang.org/grpc v1.50.1 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace io.microraft/afloatdb/admin => ../admin

replace io.microraft/afloatdb/raftendpoint => ../raftendpoint
