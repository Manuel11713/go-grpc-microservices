package main

import (
	"context"
	"log"

	pb "github.com/Manuel11713/mongo-api/proto"
)

func CreateBlog(client pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Manuel",
		Title:    "My first Blog",
		Content:  "Content of the fist blog",
	}

	res, err := client.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected err %v\n", err)
	}

	log.Printf("Blog has been created %s\n", res.Id)

	return res.Id
}
