package main

import (
  "fmt"
  // sdk "github.com/cosmos/cosmos-sdk/types"

  grpcClient "github.com/mingi3442/go-grpc/client/grpc"
)

func main() {

  go func() {
    slashingParamsRes := grpcClient.GetSlashingParamsResponse()
    fmt.Printf("Slashing Params : %v\n", string(slashingParamsRes))
  }()

  validatorRes := grpcClient.GetValidatorResponse()

  fmt.Printf("Validator : %v\n", validatorRes.Validator)
}
