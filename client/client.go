package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	proto "github.com/Viq111/protoDisco"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func call(client proto.TestClient) int {
	md := metadata.MD{
		"service": []string{"A"},
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := client.GetFeature(ctx, &empty.Empty{})
	if err != nil {
		fmt.Printf("failed to call, err=%s\n", err)
	} else {
		//fmt.Printf("called on port=%d\n", resp.GetPort())
	}
	return int(resp.GetPort())
}

func readInput(client proto.TestClient) {
	time.Sleep(time.Second)
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("Command (q/c/u): ")
		text, _ := reader.ReadString('\n')
		if text == "q\n" {
			fmt.Println("Bye.")
			return
		}
		if text == "c\n" {
			ports := make(map[int]int)
			for i := 0; i < 2; i++ {
				p := call(client)
				ports[p]++
			}
			fmt.Printf("called %v\n", ports)
		}
		if text == "u\n" {
			if len(ports) == 1 {
				ports = []int{4545, 4546}
			} else {
				ports = []int{4546}
			}
			fmt.Printf("Now targetting ports=%v\n", ports)
		}
	}
}

func init() {
	resolver.Register(&LocalResolverBuilder{})
}

func main() {
	ports = []int{4545, 4546}
	fmt.Println("Hello World!")
	conn, err := grpc.Dial("test:///hello", grpc.WithInsecure())
	panicOnErr(err)
	defer conn.Close()
	client := proto.NewTestClient(conn)
	readInput(client)

}
