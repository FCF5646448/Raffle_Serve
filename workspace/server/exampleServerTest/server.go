package exampleServerTest

import (
	"context"
	pb "sniper/rpc/example/v1"
)

// hello 接口实现
type Server struct{}


func (s *Server)SayHello(ctx context.Context, req *pb.HelloRequest) (rsp *pb.HelloResponse, err error) {
	rsp = &pb.HelloResponse{}
	if req.Name == "" {
		return
	}
	rsp.Data = &pb.HelloMessage{Message:"hello " + req.Name}
	return
}


//func (s *Server) Hello(ctx context.Context, req *pb.HelloRequest) (rsp *pb.HelloResponse, err error) {
//	rsp = &pb.HelloResponse{}
//	if req.Message == "" {
//		rsp.Msg = "参数异常"
//		return
//	}
//
//	rsp.Data = &pb.HelloMessage{Message: "hello " + req.Message}
//	return
//}