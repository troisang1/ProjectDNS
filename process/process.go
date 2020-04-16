package process

import (
	"Tiamat/config"
	Api "Tiamat/grpc"
	"Tiamat/log"
	"Tiamat/prisma"
	"context"
	"google.golang.org/grpc"
	"net"
)

type handleRequestServer struct {
	Api.UnimplementedHandleRequestServer
}

func HandleRequest() {
	listener, err := net.Listen(config.CONN_TYPE, config.LOCAL_HOST+
		":"+config.LOCAL_PORT)
	if log.CheckErr("Error listening", err) {
		return
	}

	log.Log.Info("Serving on " + config.LOCAL_HOST +
		":" + config.LOCAL_PORT + " for Manager")

	MakingGrpcServer(listener)
}
func NewServer() *handleRequestServer {
	s := &handleRequestServer{}
	return s
}
func MakingGrpcServer(listener net.Listener) {
	grpcServer := grpc.NewServer()
	Api.RegisterHandleRequestServer(grpcServer, NewServer())
	grpcServer.Serve(listener)
}

func (entity *handleRequestServer) Login(ctx context.Context, data *Api.User) (*Api.Mess, error) {
	tmp, err := prisma.Login(data.UserName, data.Password)
	if err != 0 {
		return &Api.Mess{Content: tmp}, nil
	}
	return &Api.Mess{Content: "Denied"}, nil
}
func (entity *handleRequestServer) Service(ctx context.Context, data *Api.Command) (*Api.Mess, error) {

	if data.Command == "GET" {
		tmp := prisma.GetIp(data.Domain)
		return &Api.Mess{Content: tmp}, nil
	}

	if data.Command == "ADD" {
		tmp := prisma.Add(data.Domain, data.Ip, data.User)
		return &Api.Mess{Content: tmp}, nil
	}

	if data.Command == "REMOVE" {
		tmp := prisma.Remove(data.Domain, data.User)
		return &Api.Mess{Content: tmp}, nil
	}

	if data.Command == "UPDATE" {
		tmp := prisma.Update(data.Domain, data.Ip, data.User)
		return &Api.Mess{Content: tmp}, nil
	}
	return &Api.Mess{Content: "Undefine"}, nil
}
