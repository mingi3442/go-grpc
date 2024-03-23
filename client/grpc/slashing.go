package grpc

import (
  "context"
  "encoding/json"
  "fmt"

  slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

  log "github.com/mingi3442/go-grpc/log"
)

func GetSlashingParamsResponse() []byte {
  conn := Connect()
  log.Log(log.INFO, "Slashing Client Connected")
  defer DisConnect(conn)
  slashingClient := slashingTypes.NewQueryClient(conn)
  slashingParamsRes, err := slashingClient.Params(context.Background(), &slashingTypes.QueryParamsRequest{})
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }

  jsonData, err := json.MarshalIndent(slashingParamsRes, "", "  ")
  if err != nil {
    fmt.Printf("Failed to marshal to JSON: %v", err)
  }
  return jsonData
}
