package ports

import (
  "context"
  "hexArch/internal/adapters/framework/left/grpc/pb"
) 

type GRPCPort interface {
  Run()
  GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
  GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
  GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
  GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
}
