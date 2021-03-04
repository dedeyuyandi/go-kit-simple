package user

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser     endpoint.Endpoint
	GetUser        endpoint.Endpoint
	DeleteUserByID endpoint.Endpoint
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		CreateUser:     makeCreateUserEndpoint(svc),
		GetUser:        makeGetUserEndpoint(svc),
		DeleteUserByID: makeDeleteUserByIDEndpoint(svc),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req)
		return CreateUserResponse{Message: ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		fmt.Println(req, "endpoint -> 2")
		resp, err := s.GetUser(ctx, req.Id)
		fmt.Println(resp, "response from service -> 6")
		return GetUserResponse{
			ID:       resp.ID,
			Email:    resp.Email,
			Password: resp.Password,
		}, err
	}
}

func makeDeleteUserByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserByIDRequest)
		ok, err := s.DeleteUserByID(ctx, req.Id)
		return DeleteUserByIDResponse{Message: ok}, err
	}
}
