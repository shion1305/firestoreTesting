package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
)

type (
	Form struct {
		Questions map[string]Question `firestore:"questions"`
	}

	Question struct {
		ID           string             `firestore:"-"`
		QuestionText string             `firestore:"question"`
		QuestionType int                `firestore:"question_type"`
		Order        int                `firestore:"order"`
		Properties   QuestionProperties `firestore:"properties"`
	}
	QuestionProperties map[string]interface{}
)

func testCase1(client *firestore.Client) {
	fmt.Println("testCase1")
	e := Form{
		Questions: map[string]Question{
			"q1": {
				ID:           "q1",
				QuestionText: "q1",
				QuestionType: 1,
				Order:        1,
				Properties: QuestionProperties{
					"optionID1": map[string]interface{}{
						"text":  "text1",
						"order": 1,
					},
					"optionID2": map[string]interface{}{
						"text":  "text2",
						"order": 2,
					},
				},
			},
			"q2": {
				ID:           "q2",
				QuestionText: "q2",
				QuestionType: 1,
				Order:        2,
				Properties: QuestionProperties{
					"optionID1": map[string]interface{}{
						"text":  "text1",
						"order": 1,
					},
					"optionID2": map[string]interface{}{
						"text":  "text2",
						"order": 2,
					},
				},
			},
		},
	}
	create, err := client.Collection("Forms").Doc("form1").Create(context.Background(), e)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		fmt.Println(errors.Is(err, errors.New("document already exists")))
		return
	}
	fmt.Println(create)
}

func testCase2(client *firestore.Client) {
	fmt.Println("testCase2")
	d, err := client.Collection("Forms").Doc("form1").Get(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	var e Form
	err = d.DataTo(&e)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", e)
}

func testCase3(client *firestore.Client) {
	fmt.Println("testCase3")
	_, err := client.Collection("Forms").Doc("non_exists").Get(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}

func testCase4(client *firestore.Client) {

}
