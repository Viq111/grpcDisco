package main

import (
	"fmt"
	"google.golang.org/grpc/resolver"
	"time"
)

var ports []int

type LocalResolver struct {
	cc resolver.ClientConn
}

func (r *LocalResolver) Close()                                {}
func (r *LocalResolver) ResolveNow(resolver.ResolveNowOptions) {}
func (r *LocalResolver) Up() {
	state := resolver.State{
		Addresses: make([]resolver.Address, len(ports)),
	}
	for i, p := range ports {
		state.Addresses[i] = resolver.Address{
			Addr: fmt.Sprintf("127.0.0.1:%d", p),
		}
	}
	//fmt.Printf("Update resolved addresses to: %v\n", state)
	r.cc.UpdateState(state)
}

type LocalResolverBuilder struct{}

func (b *LocalResolverBuilder) Scheme() string {
	return "test"
}

func (b *LocalResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	resolver := &LocalResolver{
		cc: cc,
	}
	resolver.Up()
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for range ticker.C {
			resolver.Up()
		}
	}()
	return resolver, nil
}
