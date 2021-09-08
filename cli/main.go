package main

import (
	"context"
	"fmt"

	hello "greeter/srv/proto/hello"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/selector"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := hello.NewSayService("go.micro.srv.greeter", service.Client())

	// TEST XKDGO
	var opts []selector.Option

	// new selector
	s := selector.NewSelector(opts...)
	selectorOpts := s.Options()
	listOfServices, _ := selectorOpts.Registry.GetService("go.micro.srv.greeter")
	fmt.Println("listOfServices", listOfServices[0])
	fmt.Println("listOfServices", listOfServices[0].Nodes[0])
	fmt.Println("Service_Name", listOfServices[0].Name)
	fmt.Println("Node_Id", listOfServices[0].Nodes[0].Id)
	fmt.Println("Node_Address", listOfServices[0].Nodes[0].Address)
	fmt.Println("listOfServices", listOfServices[0].Endpoints[0])

	// Make request
	rsp, err := cl.Hello(context.Background(), &hello.Request{
		Name: "Someone",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
