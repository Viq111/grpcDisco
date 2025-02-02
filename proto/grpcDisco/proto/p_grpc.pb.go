// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestClient interface {
	GetFeature(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TestResponse, error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	return &testClient{cc}
}

var testGetFeatureStreamDesc = &grpc.StreamDesc{
	StreamName: "GetFeature",
}

func (c *testClient) GetFeature(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/metricspb.Test/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestService is the service API for Test service.
// Fields should be assigned to their respective handler implementations only before
// RegisterTestService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type TestService struct {
	GetFeature func(context.Context, *empty.Empty) (*TestResponse, error)
}

func (s *TestService) getFeature(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/metricspb.Test/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetFeature(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterTestService registers a service implementation with a gRPC server.
func RegisterTestService(s grpc.ServiceRegistrar, srv *TestService) {
	srvCopy := *srv
	if srvCopy.GetFeature == nil {
		srvCopy.GetFeature = func(context.Context, *empty.Empty) (*TestResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "metricspb.Test",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "GetFeature",
				Handler:    srvCopy.getFeature,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "p.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewTestService creates a new TestService containing the
// implemented methods of the Test service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewTestService(s interface{}) *TestService {
	ns := &TestService{}
	if h, ok := s.(interface {
		GetFeature(context.Context, *empty.Empty) (*TestResponse, error)
	}); ok {
		ns.GetFeature = h.GetFeature
	}
	return ns
}

// UnstableTestService is the service API for Test service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableTestService interface {
	GetFeature(context.Context, *empty.Empty) (*TestResponse, error)
}
