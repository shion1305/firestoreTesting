package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firestoreTesting/domain/model/question"
	"firestoreTesting/infra/entity"
	"firestoreTesting/pkg/identity"
	"fmt"
)

func testCase1(client *firestore.Client) {
	fileQuestion := question.NewFileQuestion(identity.IssueID(), "file question", question.Image,
		question.NewImageFileConstraint(
			1, 1, 1, 1, 200, 200, 200, []string{"jpg", "png"}))

	checkBoxOptions := []question.CheckBoxOption{
		{ID: identity.IssueID(), Text: "option1"},
		{ID: identity.IssueID(), Text: "option2"},
	}
	checkBoxOptionsOrder := []question.CheckBoxOptionID{checkBoxOptions[0].ID, checkBoxOptions[1].ID}
	checkboxQuestion := question.NewCheckBoxQuestion(
		identity.IssueID(), "checkbox question", checkBoxOptions, checkBoxOptionsOrder)
	outFileQuestion := entity.Question{
		ID:      fileQuestion.ID.GetValue(),
		Text:    fileQuestion.Text,
		Type:    int(question.TypeFile),
		Customs: fileQuestion.Export().Customs,
	}
	outCheckBoxQuestion := entity.Question{
		ID:      checkboxQuestion.ID.GetValue(),
		Text:    checkboxQuestion.Text,
		Type:    int(question.TypeCheckBox),
		Customs: checkboxQuestion.Export().Customs,
	}

	_, err := client.Collection("Questions").Doc("question1").
		Set(context.Background(), outFileQuestion)
	if err != nil {
		panic(err)
		return
	}

	_, err = client.Collection("Questions").Doc("question2").
		Set(context.Background(), outCheckBoxQuestion)
	if err != nil {
		panic(err)
		return
	}
	if err != nil {
		panic(err)
		return
	}

	snap, err := client.Collection("Questions").Doc("question1").
		Get(context.Background())
	if err != nil {
		panic(err)
		return
	}
	var e entity.Question
	err = snap.DataTo(&e)
	if err != nil {
		panic(err)
		return
	}
	model, err := e.ToModel()
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("%+v\n", model)

	snap, err = client.Collection("Questions").Doc("question2").
		Get(context.Background())
	if err != nil {
		panic(err)
		return
	}
	err = snap.DataTo(&e)
	if err != nil {
		panic(err)
		return
	}
	model, err = e.ToModel()
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("%+v\n", model)
}

func testCase2(client *firestore.Client) {
	//fmt.Println("testCase2")
	//d, err := client.Collection("Forms").Doc("form1").Get(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//var e Form
	//err = d.DataTo(&e)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("%+v\n", e)
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
