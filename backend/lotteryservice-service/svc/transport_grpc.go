// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: c48ffb9f04
// Version Date: Wed Aug 12 00:08:17 UTC 2020

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/thebinarytrio/weighted-lottery/backend"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC LotteryServiceServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.LotteryServiceServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// lotteryservice

		lotteryservice: grpctransport.NewServer(
			endpoints.LotteryServiceEndpoint,
			DecodeGRPCLotteryServiceRequest,
			EncodeGRPCLotteryServiceResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the LotteryServiceServer interface
type grpcServer struct {
	lotteryservice grpctransport.Handler
}

// Methods for grpcServer to implement LotteryServiceServer interface

func (s *grpcServer) LotteryService(ctx context.Context, req *pb.AllEventsRequest) (*pb.AllEventsResponse, error) {
	_, rep, err := s.lotteryservice.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AllEventsResponse), nil
}

// Server Decode

// DecodeGRPCLotteryServiceRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC lotteryservice request to a user-domain lotteryservice request. Primarily useful in a server.
func DecodeGRPCLotteryServiceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AllEventsRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCLotteryServiceResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain lotteryservice response to a gRPC lotteryservice reply. Primarily useful in a server.
func EncodeGRPCLotteryServiceResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.AllEventsResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
