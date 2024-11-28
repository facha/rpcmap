package main

import (
	"fmt"
	"context"
	"rpcmap/server/mapservice"
)

type RPCServer struct {
	mapservice.UnimplementedMapServiceServer
	mapService *MapService 
}

func NewRPCServer(mapService *MapService) *RPCServer {
	return &RPCServer{mapService: mapService}
}

func (r *RPCServer) Put(ctx context.Context, req *mapservice.PutRequest) (*mapservice.PutResponse, error) {
	r.mapService.Put(req.Key, req.Value)
	return &mapservice.PutResponse{Success: true}, nil
}

func (r *RPCServer) Get(ctx context.Context, req *mapservice.GetRequest) (*mapservice.GetResponse, error) {
	value, ok := r.mapService.Get(req.Key)
	if !ok {
		return nil, fmt.Errorf("key %s not found", req.Key)
	}
	return &mapservice.GetResponse{Value: value}, nil
}

func (r *RPCServer) Delete(ctx context.Context, req *mapservice.DeleteRequest) (*mapservice.DeleteResponse, error) {
	success := r.mapService.Delete(req.Key)
	return &mapservice.DeleteResponse{Success: success}, nil
}

