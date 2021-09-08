package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	hello "greeter/srv/proto/hello"

	_ "github.com/asim/go-micro/plugins/registry/etcd/v3"
	_ "github.com/asim/go-micro/plugins/transport/nats/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/util/log"

	"github.com/asim/go-micro/plugins/server/grpc/v3"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Log("Received Say.Hello request")
	var b strings.Builder
	b.WriteString("Hello " + req.Name + "\n")
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
		fmt.Fprintf(&b, "%s...\n", res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
		fmt.Fprintf(&b, "timeout 1...\n")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
		fmt.Fprintf(&b, "%s...\n", res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
		fmt.Fprintf(&b, "timeout 2...\n")
	}

	log.Log(ctx)
	log.Log(req)
	rsp.Msg = b.String()
	log.Log("Message you want: ", rsp.Msg)
	return nil
}

func (s *Say) Goodbye(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Log("Received Say.Goodbye request")
	var b strings.Builder
	b.WriteString("Goodbye " + req.Name + "\n")
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(15 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
		fmt.Fprintf(&b, "%s...\n", res)
	case <-time.After(10 * time.Second):
		fmt.Println("timeout 1")
		fmt.Fprintf(&b, "timeout 1...\n")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
		fmt.Fprintf(&b, "%s...\n", res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
		fmt.Fprintf(&b, "timeout 2...\n")
	}

	log.Log(ctx)
	log.Log(req)
	rsp.Msg = b.String()
	log.Log("Message you want: ", rsp.Msg)
	return nil
}

func main() {

	service := micro.NewService(
		micro.Server(grpc.NewServer()), micro.Name("go.micro.srv.greeter"))

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	err := hello.RegisterSayHandler(service.Server(), new(Say))
	if err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
