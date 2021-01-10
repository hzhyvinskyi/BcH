package accountMgr

import (
	"context"
	"log"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

type CreateUserReq struct {
	Email    string
	Password string
}

type CreateUserResp struct {
	user *User
}

type GetUserReq struct {
	ID string
}

type GetUserResp struct {
	user *User
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserReq)
		user, err := s.CreateUser(ctx, req.Email, req.Password)
		if err != nil {
			log.Fatalln(err.Error())
		}
		return CreateUserResp{user: user}, nil
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserReq)
		user, err := s.GetUser(ctx, req.ID)
		if err != nil {
			log.Fatalln(err.Error())
		}
		return GetUserResp{user: user}, nil
	}
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}
