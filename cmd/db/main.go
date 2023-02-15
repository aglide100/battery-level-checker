package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	grpcAddr = flag.String("apid grpc addr", "0.0.0.0:50011", "grpc address")
	usingTls = flag.Bool("grpc.tls", false, "using http2")
)

func main() {
	if err := realMain(); err != nil {
		log.Printf("err :%s", err)
		os.Exit(1)
	}
}

func realMain() error {
	gRPCL, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		return err
	}
	defer gRPCL.Close()
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	wg, _ := errgroup.WithContext(context.Background())

	wg.Go(func() error {
		logger.Info("Starting grpc server...")
		err := grpcServer.Serve(gRPCL)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}

		return nil
	})

	return wg.Wait()
}
