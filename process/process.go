package process

import (
	"Tiamat/config"
	Api "Tiamat/grpc"
	"Tiamat/log"
	"Tiamat/prisma"
	"context"
	"google.golang.org/grpc"
	"net"
	"math/rand"
)

var RegisUser  = make(map[string] string)

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
		id := Random()
		RegisUser[id] = tmp
		return &Api.Mess{Content: id}, nil
	}
	return &Api.Mess{Content: "Denied"}, nil
}
func (entity *handleRequestServer) Service(ctx context.Context, data *Api.Command) (*Api.Mess, error) {

	if data.Command == "GET" {
		tmp := prisma.GetIp(data.Domain)
		return &Api.Mess{Content: tmp}, nil
	}
	UserId := RegisUser[data.User]
	if data.Command == "ADD" {
		tmp := prisma.Add(data.Domain, data.Ip, UserId)
		return &Api.Mess{Content: tmp}, nil
	}

	if data.Command == "REMOVE" {
		tmp := prisma.Remove(data.Domain, UserId)
		return &Api.Mess{Content: tmp}, nil
	}

	if data.Command == "UPDATE" {
		tmp := prisma.Update(data.Domain, data.Ip, UserId)
		return &Api.Mess{Content: tmp}, nil
	}
	if data.Command == "LOGOUT" {
		delete(RegisUser, data.User)
		return &Api.Mess{Content: "Accept"}, nil
	}
	return &Api.Mess{Content: "Undefine"}, nil
}

func Random() string {
	char := "qwertyuioplkjhgfdsazxcvbnm1234567890."
	id := ""
	for i:=0; i<10; i++ {
		id = id + string(char[rand.Intn(len(char))])
	}
	return id
}
