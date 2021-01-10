package accountMgr

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	pb "github.com/hzhyvinskyi/BcH/common/pb/generated"
)

type gRPCServer struct {
	createUser grpctransport.Handler
	getUser    grpctransport.Handler
}

func NewGRPCServer(endpoints Endpoints, logger log.Logger) pb.AccountMgrServiceServer {
	return &gRPCServer{
		createUser: grpctransport.NewServer(
			endpoints.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
		),
		getUser:    grpctransport.NewServer(
			endpoints.GetUser,
			decodeGetUserRequest,
			encodeGetUserResponse,
		),
	}
}

func (s *gRPCServer) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateUserResponse), nil
}

func (s *gRPCServer) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetUserResponse), nil
}

func decodeCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(pb.CreateUserRequest)
	return CreateUserReq{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

func encodeCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(pb.CreateUserResponse)
	return &resp, nil
}

func decodeGetUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(pb.GetUserRequest)
	return GetUserReq{
		ID: req.Id,
	}, nil
}

func encodeGetUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(pb.GetUserResponse)
	return &resp, nil
}
