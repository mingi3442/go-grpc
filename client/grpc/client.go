package grpc

import (
  "fmt"

  "google.golang.org/grpc"
)

func Connect() *grpc.ClientConn {
  conn, err := grpc.Dial("cosmos-grpc.polkachu.com:14990", grpc.WithInsecure())
  if err != nil {
    fmt.Printf("did not connect: %v", err)

  }
  return conn
}

func DisConnect(conn *grpc.ClientConn) {
  conn.Close()
}
