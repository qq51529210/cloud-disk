package api

import context "context"

type grpcService struct {
	UnimplementedApiServer
}

func (s *grpcService) AddUser(context.Context, *AddUserReq) (*AddUserRes, error) {
	return &AddUserRes{}, nil
}
