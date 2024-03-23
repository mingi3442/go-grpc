package grpc

import (
  "fmt"

  log "github.com/mingi3442/go-grpc/log"
  "google.golang.org/grpc"
)

func Connect() *grpc.ClientConn {
  conn, err := grpc.Dial("cosmos-grpc.polkachu.com:14990", grpc.WithInsecure())
  if err != nil {
    errorMessage := fmt.Sprintf("did not connect: %v", err)
    log.Log(log.ERROR, errorMessage)

  }
  log.Log(log.INFO, "Connected to cosmos-grpc")
  return conn
}

func DisConnect(conn *grpc.ClientConn) {
  log.Log(log.INFO, "Disconnected from cosmos-grpc")
  conn.Close()
}
