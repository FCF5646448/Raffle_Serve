package exampleServerTest

import (
	"context"
	pb "workspace/rpc/example/v1"
)

// hello 接口实现
type Server struct{}

func (s *Server) Hello(ctx context.Context, req *pb.HelloRequest) (rsp *pb.HelloResponse, err error) {
	rsp = &pb.HelloResponse{}
	if req.Message == "" {
		rsp.Msg = "参数异常"
		return
	}

	rsp.Data = &pb.HelloMessage{Message: "hello " + req.Message}
	return
}