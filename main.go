package main

import (
  context "context"
  "fmt"

  // sdk "github.com/cosmos/cosmos-sdk/types"

  slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
  stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
  client "github.com/mingi3442/go-grpc/client"
  log "github.com/mingi3442/go-grpc/log"
)

func main() {
  log.Log(log.INFO, "Hello, World!")
  log.Log(log.WARN, "Hello, World!")
  // conn, err := grpc.Dial("cosmos-grpc.polkachu.com:14990", grpc.WithInsecure())
  // if err != nil {
  //   fmt.Printf("did not connect: %v", err)

  // }
  // defer conn.Close()
  conn := client.Connect()
  defer client.DisConnect(conn)

  slashingClient := slashingtypes.NewQueryClient(conn)
  stakingClient := stakingtypes.NewQueryClient(conn)

  // a41ConsAddr, err := sdk.ConsAddressFromBech32("cosmosvalcons1v78emy9d2xe3tj974l7tmn2whca2nh9z4drnsy")
  // if err != nil {
  //   fmt.Printf("invalid address: %v", err)
  // }

  slashingParamsRes, err := slashingClient.Params(context.Background(), &slashingtypes.QueryParamsRequest{})
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Printf("Slashing Param: %v\n", slashingParamsRes.Params.SignedBlocksWindow)

  // slashingInfosRes, err := slashingClient.SigningInfos(context.Background(), &slashingtypes.QuerySigningInfosRequest{})
  // if err != nil {
  //   fmt.Printf("Error: %v\n", err)
  // }
  // fmt.Printf("Slashing Infos: %v\n", slashingInfosRes.Info)

  // validatorDelegationsRes, err := stakingClient.ValidatorDelegations(context.Background(), &stakingtypes.QueryValidatorDelegationsRequest{
  //   ValidatorAddr: "cosmosvaloper1v78emy9d2xe3tj974l7tmn2whca2nh9zp7s0u9",
  // })
  // if err != nil {
  //   fmt.Printf("Error: %v\n", err)
  // }
  // fmt.Printf("Validator Delegations : %v\n", validatorDelegationsRes)
  validatorRes, err := stakingClient.Validator(context.Background(), &stakingtypes.QueryValidatorRequest{
    ValidatorAddr: "cosmosvaloper1v78emy9d2xe3tj974l7tmn2whca2nh9zp7s0u9",
  })
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Printf("Validator : %v\n", validatorRes.Validator)
}
