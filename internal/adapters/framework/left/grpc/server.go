package rpc

import (
	"log"
	"net"

	"hexArch/internal/adapters/framework/left/grpc/pb"
	"hexArch/internal/ports"

	"google.golang.org/grpc"
)

type Adapter struct {
  api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter  {
  return &Adapter{api: api}
}

func (grpca Adapter) Run()  {
  var err error

  listen, err := net.Listen("tcp", ":9000")
  if err != nil {
    log.Fatalf("failed to listen on port 9000: %v", err)
  }

  arithmeticServiceServer := grpca
  grpcServer := grpc.NewServer()
  pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

  if err := grpcServer.Serve(listen); err != nil {
    log.Fatalf("failed to serve gRPC server on port 9000: %v", err)
  }
}
