package main

import (
	"context"
	"log"
	pb "proto"
	"time"

	"google.golang.org/grpc"
)


func main(){
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewTranslateServiceClient(conn)

	text := []string{"Salom", "dunyo", "kitob", "darsturlash", "olma"}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.TranslateText(ctx, &pb.TranslateRequest{Words: text})
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("Translated text: %v", r.GetTranslatedText())
}