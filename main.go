package main

import (
  context "context"
  "fmt"

  sdk "github.com/cosmos/cosmos-sdk/types"
  banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

  "google.golang.org/grpc"
)

func main() {

  conn, err := grpc.Dial("cosmos-grpc.polkachu.com:14990", grpc.WithInsecure())
  if err != nil {
    fmt.Printf("did not connect: %v", err)
  }
  defer conn.Close()

  bankClient := banktypes.NewQueryClient(conn)

  addr, err := sdk.AccAddressFromBech32("cosmos1v78emy9d2xe3tj974l7tmn2whca2nh9zy2y6sk")
  if err != nil {
    fmt.Printf("invalid address: %v", err)
  }

  res2, err := bankClient.Balance(context.Background(), &banktypes.QueryBalanceRequest{
    Address: addr.String(),
    Denom:   "stake",
  })
  if err != nil {
    fmt.Printf("failed to query balance: %v", err)
  }

  fmt.Printf("Balance: %v", res2.Balance)
}
