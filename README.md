# e-wallet-ums

Generate proto:

1. Create file `.proto`

2. Enter the dir where proto located

   example: `cd cmd/proto/tokenvalidation`

3. Type command

   `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative {your_proto_filename}.proto`
