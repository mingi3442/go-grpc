package main

import (
  context "context"
  "fmt"

  sdk "github.com/cosmos/cosmos-sdk/types"

  slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

  "google.golang.org/grpc"
)

func main() {

  conn, err := grpc.Dial("cosmos-grpc.polkachu.com:14990", grpc.WithInsecure())
  if err != nil {
    fmt.Printf("did not connect: %v", err)
  }
  defer conn.Close()

  slashingClient := slashingtypes.NewQueryClient(conn)

  a41ConsAddr, err := sdk.ConsAddressFromBech32("cosmosvalcons1v78emy9d2xe3tj974l7tmn2whca2nh9z4drnsy")
  if err != nil {
    fmt.Printf("invalid address: %v", err)
  }

  slashingParamsRes, err := slashingClient.Params(context.Background(), &slashingtypes.QueryParamsRequest{})
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Printf("Slashing Param: %v\n", slashingParamsRes.Params.SignedBlocksWindow)

  slashingInfosRes, err := slashingClient.SigningInfos(context.Background(), &slashingtypes.QuerySigningInfosRequest{})
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Printf("Slashing Infos: %v\n", slashingInfosRes.Info)
  slashingInfoRes, err := slashingClient.SigningInfo(context.Background(), &slashingtypes.QuerySigningInfoRequest{
    ConsAddress: a41ConsAddr.String(),
  })
  fmt.Printf("Slashing Info: %v\n", slashingInfoRes.GetValSigningInfo())
  if err != nil {
    fmt.Printf("Error: %v", err)
  }

}
