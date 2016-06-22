Experiment with gRPC and twitter API


update protobuf
`protoc -I twitter/ twitter/twitter.proto --go_out=plugins=grpc:twitter`