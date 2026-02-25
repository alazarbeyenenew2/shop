package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alazarbeyenenew2/shop/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := pb.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.CreateProduct(ctx, &pb.CreateProductRequest{
		Name:        "Test Product",
		Description: "test product description",
		Category:    "unknown",
		BasePrice:   300,
	})

	fmt.Println(resp)

}
