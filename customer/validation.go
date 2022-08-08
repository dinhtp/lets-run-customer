package customer

import (
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    pb "github.com/dinhtp/lets-run-pbtype/gateway"
)

func validateOne(r *pb.OneCustomerRequest) error {
    if r.GetPlatformId() == "" {
        return status.Error(codes.InvalidArgument, "platform id is required")
    }

    if r.GetId() == "" {
        return status.Error(codes.InvalidArgument, "customer id is required")
    }

    return nil
}

func validateCreate(r *pb.Customer) error {
    if r.GetPlatformId() == "" {
        return status.Error(codes.InvalidArgument, "platform id is required")
    }

    if r.GetFirstName() == "" {
        return status.Error(codes.InvalidArgument, "first name is required")
    }

    if r.GetLastName() == "" {
        return status.Error(codes.InvalidArgument, "last name is required")
    }

    if r.GetEmail() == "" {
        return status.Error(codes.InvalidArgument, "email is required")
    }

    return nil
}
