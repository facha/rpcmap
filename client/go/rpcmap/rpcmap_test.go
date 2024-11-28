package rpcmap

import (
	"context"
	"testing"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"rpcmap/client/go/mapservice"
)

type MockMapServiceClient struct {
	mock.Mock
}

func (m *MockMapServiceClient) Get(ctx context.Context, req *mapservice.GetRequest, opts ...grpc.CallOption) (*mapservice.GetResponse, error) {
	args := m.Called(ctx, req)
	if resp, ok := args.Get(0).(*mapservice.GetResponse); ok {
		return resp, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMapServiceClient) Put(ctx context.Context, req *mapservice.PutRequest, opts ...grpc.CallOption) (*mapservice.PutResponse, error) {
	args := m.Called(ctx, req)
	if resp, ok := args.Get(0).(*mapservice.PutResponse); ok {
		return resp, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMapServiceClient) Delete(ctx context.Context, req *mapservice.DeleteRequest, opts ...grpc.CallOption) (*mapservice.DeleteResponse, error) {
	args := m.Called(ctx, req)
	if resp, ok := args.Get(0).(*mapservice.DeleteResponse); ok {
		return resp, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestMapClient_Get(t *testing.T) {
	mockClient := new(MockMapServiceClient)
	mapClient := &MapClient{client: mockClient}

	key := "testKey"
	expectedValue := "testValue"

	mockClient.On("Get", mock.Anything, &mapservice.GetRequest{Key: key}).
		Return(&mapservice.GetResponse{Value: expectedValue}, nil)

	value, err := mapClient.Get(key)
	if err != nil {
		t.Fatalf("Get() returned an error: %v", err)
	}
	if value != expectedValue {
		t.Errorf("Get() = %v; want %v", value, expectedValue)
	}

	mockClient.AssertExpectations(t)
}

