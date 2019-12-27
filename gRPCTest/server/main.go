package main

import ( 
	"context"
	"github.com/uzdigital12345/golang/gRPCTest/proto"
)

const (
	port = 4040
)
type server struct{}

func (s *server) Add(ctx context.Context,request *proto.Request) (*proto.Responce,error) {

	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Responce{Result:result},nil	
}

func (s *server) Multiply(ctx context.Context,request *proto.Request) (*proto.Responce,error) {

	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Responce{Result:result},nil	
}

func main() {

	p := fmt.Sprintf(":%d",port)
	listener, err := net.Listen("tcp",a)
	if err!=nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv,&server{})
	reflection.Register(srv)

	if e:=srv.Serve(listener); e!=nil {
		panic(e)
	}
}
