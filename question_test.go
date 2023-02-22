package main

import (
	fs "cloud.google.com/go/firestore"
	"context"
	"firestoreTesting/domain/model/question"
	"firestoreTesting/pkg/identity"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"log"
	"os"
	"testing"
)

func TestQuestions(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
		return
	}
	ctx := context.Background()
	client := getFirestore(ctx)
	defer client.Close()

	fileQuestion := question.NewFileQuestion(identity.IssueID(), "file question", question.Image,
		question.NewImageFileConstraint(
			1, 1, 1, 1, 200, 200, 200, []string{"jpg", "png"}))
	checkBoxOptions := []question.CheckBoxOption{
		{ID: identity.IssueID(), Text: "option1"},
		{ID: identity.IssueID(), Text: "option2"},
	}
	checkBoxOptionsOrder := []question.CheckBoxOptionID{checkBoxOptions[1].ID, checkBoxOptions[0].ID}
	checkboxQuestion := question.NewCheckBoxQuestion(
		identity.IssueID(), "checkbox question", checkBoxOptions, checkBoxOptionsOrder)
	t1 := genTest(client, fileQuestion)
	t2 := genTest(client, checkboxQuestion)
	t.Run("file question", t1)
	t.Run("checkbox question", t2)
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
