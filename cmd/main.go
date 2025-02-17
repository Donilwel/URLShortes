package cmd

import (
	"context"
	"fmt"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewURLShortenerClient(conn)

	req := &proto.CreateShortURLRequest{
		OriginalUrl: "https://www.example.com",
	}
	res, err := client.CreateShortURL(context.Background(), req)
	if err != nil {
		log.Fatalf("could not create short URL: %v", err)
	}
	fmt.Printf("Shortened URL: %s\n", res.GetShortUrl())

	getReq := &proto.GetOriginalURLRequest{
		ShortCode: "shortCode", // используйте реальный сокращённый код
	}
	getRes, err := client.GetOriginalURL(context.Background(), getReq)
	if err != nil {
		log.Fatalf("could not get original URL: %v", err)
	}
	fmt.Printf("Original URL: %s\n", getRes.GetOriginalUrl())
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
