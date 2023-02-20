package main

import (
	fs "cloud.google.com/go/firestore"
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"log"
	"os"
)

func init() {

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
		return
	}
	ctx := context.Background()
	client := getFirestore(ctx)
	defer client.Close()
	testCase1(client)
	testCase2(client)
	testCase3(client)
	testCase01(client)
	testCase02(client)
}

func getFirestore(ctx context.Context) *fs.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		return nil
	}
	data, err := os.ReadFile(os.Getenv("GC_FIRESTORE_CREDENTIAL"))
	options := option.WithCredentialsJSON(data)
	client, err := fs.NewClient(ctx, os.Getenv("GC_PROJECT_ID"), options)
	if err != nil {
		log.Fatalf("firebase.NewClient err: %v", err)
		return nil
	}
	return client
}
