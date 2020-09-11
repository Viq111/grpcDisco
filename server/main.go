package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	proto "github.com/Viq111/protoDisco"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
)

var options struct {
	Port int
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

type TestServer struct {
	port        int
	serviceName string
}

var testServer *TestServer

func (s *TestServer) GetFeature(ctx context.Context, req *empty.Empty) (*proto.TestResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	serviceRequested := ""
	if len(md["service"]) > 0 {
		serviceRequested = md["service"][0]
	}
	if serviceRequested != s.serviceName {
		return nil, status.Error(codes.FailedPrecondition, "nope")
	}

	resp := &proto.TestResponse{
		Port: int64(s.port),
	}
	return resp, nil
}

func init() {
	flag.IntVar(&options.Port, "p", 4545, "port to listen to")
}

func readInput() {
	time.Sleep(time.Second)
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("Command (q/s): ")
		text, _ := reader.ReadString('\n')
		if text == "q\n" {
			fmt.Println("Bye.")
			return
		}
		if text == "s\n" {
			serviceName := "A"
			if testServer.serviceName == "A" {
				serviceName = "B"
			}
			fmt.Printf("Swapping service name to %s\n", serviceName)
			testServer.serviceName = serviceName
		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", options.Port))
	panicOnErr(err)
	grpcServer := grpc.NewServer()
	testServer = &TestServer{
		port:        options.Port,
		serviceName: "A",
	}

	proto.RegisterTestService(grpcServer, proto.NewTestService(testServer))
	go func() {
		fmt.Printf("Listing on :%d...\n", options.Port)
		grpcServer.Serve(lis)
	}()
	readInput()
}
