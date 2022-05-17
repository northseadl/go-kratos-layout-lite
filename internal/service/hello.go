package service

import (
	"context"
	"go-kratos-layout-lite/internal/biz"

	pb "go-kratos-layout-lite/api/hello/v1"
)

type HelloService struct {
	pb.UnimplementedHelloServer
	euc *biz.HelloUC
}

func NewHelloService(euc *biz.HelloUC) *HelloService {
	return &HelloService{
		euc: euc,
	}
}

func (s *HelloService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *HelloService) Hi(ctx context.Context, req *pb.HiRequest) (*pb.HiReply, error) {
	return &pb.HiReply{}, nil
}
