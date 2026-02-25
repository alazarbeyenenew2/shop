package initiator

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/alazarbeyenenew2/shop/internal/app/transport/grpc/product"
	"github.com/alazarbeyenenew2/shop/internal/pkg/logger"
	"github.com/alazarbeyenenew2/shop/pb"
	"google.golang.org/grpc"
)

func Initiate() {
	ctx := context.Background()
	log := logger.New()
	client := InitDB(log, ctx)
	repo := InitRepo(ctx, client, log)
	services := InitServices(repo, log)
	port := os.Getenv("APP_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &product.ProductHandler{
		ProductService: services.ProductService,
	})
	fmt.Printf("server running on port %s \n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
