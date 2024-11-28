package rpcmap

import (
	"fmt"
	"context"
	"google.golang.org/grpc"
    "rpcmap/client/go/mapservice" 
)

type MapClient struct {
	client mapservice.MapServiceClient
	conn   *grpc.ClientConn
}

func NewMapClient() *MapClient {
	return &MapClient{}
}

func (c *MapClient) Connect(hostPort string) error {
	conn, err := grpc.Dial(hostPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("failed to connect to server: %v", err)
	}
	c.client = mapservice.NewMapServiceClient(conn)
	c.conn = conn
	return nil
}

func (c *MapClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *MapClient) Get(key string) (string, error) {
	req := &mapservice.GetRequest{Key: key}
	resp, err := c.client.Get(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("failed to get value: %v", err)
	}
	return resp.GetValue(), nil
}

func (c *MapClient) Put(key, value string) error {
	req := &mapservice.PutRequest{Key: key, Value: value}
	_, err := c.client.Put(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to put value: %v", err)
	}
	return nil
}

func (c *MapClient) Del(key string) error {
	req := &mapservice.DeleteRequest{Key: key}
	_, err := c.client.Delete(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to delete value: %v", err)
	}
	return nil
}
