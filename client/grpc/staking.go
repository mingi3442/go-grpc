package grpc

import (
  "context"
  "fmt"

  stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
  "github.com/mingi3442/go-grpc/log"
)

func GetValidatorResponse() *stakingTypes.QueryValidatorResponse {
  conn := Connect()
  log.Log(log.INFO, "Slashing Client Connected")
  defer DisConnect(conn)
  stakingClient := stakingTypes.NewQueryClient(conn)

  validatorRes, err := stakingClient.Validator(context.Background(), &stakingTypes.QueryValidatorRequest{
    ValidatorAddr: "cosmosvaloper1v78emy9d2xe3tj974l7tmn2whca2nh9zp7s0u9",
  })
  if err != nil {
    fmt.Printf("Failed to marshal to JSON: %v", err)
  }

  return validatorRes
}
