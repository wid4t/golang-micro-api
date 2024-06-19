package main

import (
	"context"
	"fmt"
	pb "golang-micro-api/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	serviceName = "golang-micro-api"
	version     = "latest"
)

func main() {

	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(version),
	)

	service.Init()

	client := pb.NewGolangMicroProductService("golang-micro-product", service.Client())

	req := &pb.CallRequest{
		Name: "widat",
	}

	resp, err := client.Call(context.Background(), req)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(resp.Msg)

}
