package customer

import (
    "context"

    "github.com/gogo/protobuf/types"

    pb "github.com/dinhtp/lets-run-pbtype/gateway"
)

type Service struct {
}

func NewService() *Service {
    return &Service{}
}

func (s Service) Get(ctx context.Context, r *pb.OneCustomerRequest) (*pb.Customer, error) {
    panic("implement me")
}

func (s Service) Create(ctx context.Context, r *pb.Customer) (*pb.Customer, error) {
    panic("implement me")
}

func (s Service) Update(ctx context.Context, r *pb.Customer) (*pb.Customer, error) {
    panic("implement me")
}

func (s Service) Delete(ctx context.Context, r *pb.OneCustomerRequest) (*types.Empty, error) {
    panic("implement me")
}

func (s Service) List(ctx context.Context, r *pb.ListCustomerRequest) (*pb.ListCustomerResponse, error) {
    panic("implement me")
}
