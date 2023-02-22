package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firestoreTesting/domain/model/question"
	"firestoreTesting/infra/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func genTest(client *firestore.Client, target question.Question) func(t *testing.T) {
	e := entity.Question{
		ID:      target.GetID().GetValue(),
		Text:    target.GetText(),
		Type:    int(target.GetType()),
		Customs: target.Export().Customs,
	}
	return func(t *testing.T) {
		_, err := client.Collection("Questions").Doc(target.GetID().ExportID()).
			Set(context.Background(), e)
		if err != nil {
			t.Fatal(err)
			return
		}

		snap, err := client.Collection("Questions").Doc(target.GetID().ExportID()).
			Get(context.Background())
		if err != nil {
			t.Fatal(err)
			return
		}
		var e entity.Question
		err = snap.DataTo(&e)
		if err != nil {
			t.Fatal(err)
			return
		}
		model, err := e.ToModel()
		if err != nil {
			t.Fatal(err)
			return
		}
		var sm interface{}
		var ok bool
		switch model.GetType() {
		case question.TypeCheckBox:
			sm, ok = model.(*question.CheckBoxQuestion)
		case question.TypeFile:
			sm, ok = model.(*question.FileQuestion)
		default:
			t.Fatalf("test not implemented for type: %d", model.GetType())
		}
		if !ok {
			t.Fatalf("fail to make assertion, %#v", model)
			return
		}
		assert.Equal(t, target, sm)
	}
}
