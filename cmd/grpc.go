package cmd

import (
    "context"
    "net"
    "os"
    "os/signal"
    "syscall"

    _ "github.com/go-sql-driver/mysql"
    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "google.golang.org/grpc"

    "github.com/dinhtp/lets-run-customer/customer"
    pb "github.com/dinhtp/lets-run-pbtype/gateway"
)

var grpcCmd = &cobra.Command{
    Use:   "grpc",
    Short: "Run Customer service grpc command",
    Run:   runGrpcCommand,
}

func init() {
    serveCmd.AddCommand(grpcCmd)

    grpcCmd.Flags().StringP("backend", "", "grpc-address", "gRPC address")
    grpcCmd.Flags().StringP("platform", "", "platform-grpc-address", "platform grpc address")
    grpcCmd.Flags().StringP("shopifyCustomer", "", "shopify-grpc-address", "shopify customer grpc address")
    grpcCmd.Flags().StringP("wooCommerceCustomer", "", "woo-grpc-address", "woocommerce customer grpc address")

    _ = viper.BindPFlag("backend", grpcCmd.Flags().Lookup("backend"))
    _ = viper.BindPFlag("platform", grpcCmd.Flags().Lookup("platform"))
    _ = viper.BindPFlag("shopifyCustomer", grpcCmd.Flags().Lookup("shopifyCustomer"))
    _ = viper.BindPFlag("wooCommerceCustomer", grpcCmd.Flags().Lookup("wooCommerceCustomer"))
}

func runGrpcCommand(cmd *cobra.Command, args []string) {
    ctx := context.Background()
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    // services
    grpcServer := initializeServices(grpc.NewServer())

    // init GRPC backend
    grpcAddr := viper.GetString("backend")
    lis, err := net.Listen("tcp", grpcAddr)
    if err != nil {
        panic(err)
    }

    // Serve GRPC
    go func() {
        err = grpcServer.Serve(lis)
        if err != nil {
            panic(err)
        }
    }()

    logrus.WithFields(logrus.Fields{
        "service": "run-customer-service",
        "type":    "grpc",
        "address": grpcAddr,
    }).Info("run customer service server started")

    <-c
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    logrus.WithFields(logrus.Fields{
        "service": "run-customer-service",
        "type":    "grpc",
        "address": grpcAddr,
    }).Info("run customer service gracefully shutdowns")

}

func initializeServices(grpcServer *grpc.Server) *grpc.Server {
    customerService := customer.NewService()
    pb.RegisterCustomerServiceServer(grpcServer, customerService)

    return grpcServer
}
