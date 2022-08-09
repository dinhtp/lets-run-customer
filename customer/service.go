package customer

import (
    "context"

    "github.com/gogo/protobuf/types"
    "google.golang.org/grpc"

    pb "github.com/dinhtp/lets-run-pbtype/gateway"
    ppb "github.com/dinhtp/lets-run-pbtype/platform"
)

type Service struct {
    platformSelector Platform
    platformGrpc     string
}

func NewService(platformSelector Platform, platformGrpc string) *Service {
    return &Service{
        platformSelector: platformSelector,
        platformGrpc:     platformGrpc,
    }
}

func (s Service) Get(ctx context.Context, r *pb.OneCustomerRequest) (*pb.Customer, error) {
    if err := validateOne(r); err != nil {
        return nil, err
    }

    // dial grpc connection to platform service
    platformConn, err := grpc.Dial(s.platformGrpc, grpc.WithInsecure())
    if nil != err {
        return nil, err
    }

    defer platformConn.Close()

    // get platform data by id
    platformRequest := &pb.OnePlatformRequest{Id: r.GetPlatformId()}
    platformData, err := pb.NewPlatformServiceClient(platformConn).Get(ctx, platformRequest)
    if nil != err {
        return nil, err
    }

    // select ecommerce service address based on platform type
    ecomAddress := s.platformSelector.GetEndpoint(platformData.GetType())

    // dial grpc connection to ecommerce service
    ecomConn, err := grpc.Dial(ecomAddress, grpc.WithInsecure())
    if nil != err {
        return nil, err
    }

    defer ecomConn.Close()

    // forward request to ecommerce service to handle logic
    ecomRequest := &ppb.OneCustomerRequest{Platform: platformData, Id: r.GetId()}
    return ppb.NewCustomerServiceClient(ecomConn).Get(ctx, ecomRequest)
}

func (s Service) Create(ctx context.Context, r *pb.Customer) (*pb.Customer, error) {
    if err := validateCreate(r); err != nil {
        return nil, err
    }

    // dial grpc connection to platform service
    platformConn, err := grpc.Dial(s.platformGrpc, grpc.WithInsecure())
    if nil != err {
        return nil, err
    }

    defer platformConn.Close()

    // get platform data by id
    platformRequest := &pb.OnePlatformRequest{Id: r.GetPlatformId()}
    platformData, err := pb.NewPlatformServiceClient(platformConn).Get(ctx, platformRequest)
    if nil != err {
        return nil, err
    }

    // select ecommerce service address based on platform type
    ecomAddress := s.platformSelector.GetEndpoint(platformData.GetType())

    // dial grpc connection to ecommerce service
    ecomConn, err := grpc.Dial(ecomAddress, grpc.WithInsecure())
    if nil != err {
        return nil, err
    }

    defer ecomConn.Close()

    // forward request to ecommerce service to handle logic
    ecomRequest := &ppb.CreateUpdateCustomerRequest{Platform: platformData, Customer: r}
    return ppb.NewCustomerServiceClient(ecomConn).Create(ctx, ecomRequest)
}

func (s Service) Update(ctx context.Context, r *pb.Customer) (*pb.Customer, error) {
    // TODO: implement logic
    return &pb.Customer{}, nil
}

func (s Service) Delete(ctx context.Context, r *pb.OneCustomerRequest) (*types.Empty, error) {
    // TODO: implement logic
    return &types.Empty{}, nil
}

func (s Service) List(ctx context.Context, r *pb.ListCustomerRequest) (*pb.ListCustomerResponse, error) {
    // TODO: implement logic
    return &pb.ListCustomerResponse{}, nil
}
