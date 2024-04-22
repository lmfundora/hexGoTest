package main

import (
  "os"

  "hexArch/internal/adapters/app/api"
  "hexArch/internal/adapters/core/arithmetic"
  "hexArch/internal/adapters/framework/right/db"
  "hexArch/internal/ports"

  gRPC "hexArch/internal/adapters/framework/left/grpc"

)

func main()  {

  //ports
  var dbaseAdapter ports.DbPort
  var core ports.ArithmeticPort
  var appAdapter ports.APIPort
  var gRPCAdapter ports.GRPCPort

  dbaseDriver := os.Getenv("DB_DRIVER")
  dsourceName := os.Getenv("DS_NAME")

  dbaseAdapter = db.NewAdapter(dbaseDriver, dsourceName)

  defer dbaseAdapter.CloseDbConnection()

  core = arithmetic.NewAdapter()

  appAdapter = api.NewAdapter(dbaseAdapter, core)

  gRPCAdapter = gRPC.NewAdapter(appAdapter)

  gRPCAdapter.Run()
}
