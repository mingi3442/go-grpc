package main

import (
  "fmt"
  "sync"
  // sdk "github.com/cosmos/cosmos-sdk/types"
  grpcClient "github.com/mingi3442/go-grpc/client/grpc"
  log "github.com/mingi3442/go-grpc/log"
)

func main() {
  var wg sync.WaitGroup
  wg.Add(2)

  log.Log(log.DEBUG, "Starting the goroutines")

  go func() {
    defer wg.Done()
    validatorRes := grpcClient.GetValidatorResponse()
    fmt.Printf("Validator : %v\n", validatorRes.Validator)
    log.Log(log.DEBUG, "End of goroutine 1")
  }()
  go func() {
    defer wg.Done()
    slashingParamsRes := grpcClient.GetSlashingParamsResponse()
    fmt.Printf("Slashing Params : %v\n", string(slashingParamsRes))
    log.Log(log.DEBUG, "End of goroutine 2")
  }()

  wg.Wait()
  log.Log(log.DEBUG, "Waiting for goroutines to finish")

}
