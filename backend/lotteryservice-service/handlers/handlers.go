package handlers

import (
	"context"

	pb "github.com/thebinarytrio/weighted-lottery/backend"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.LotteryServiceServer {
	return lotteryserviceService{}
}

type lotteryserviceService struct{}

func (s lotteryserviceService) LotteryService(ctx context.Context, in *pb.AllEventsRequest) (*pb.AllEventsResponse, error) {
	var resp pb.AllEventsResponse
	return &resp, nil
}
