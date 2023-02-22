package entity

import (
	"firestoreTesting/domain/model/question"
	"firestoreTesting/pkg/identity"
)

type Question struct {
	ID      int64                  `firestore:"id"`
	Text    string                 `firestore:"text"`
	Type    int                    `firestore:"type"`
	Customs map[string]interface{} `firestore:"customs"`
}

func (q Question) ToModel() (question.Question, error) {
	sq := question.NewStandardQuestion(
		question.Type(q.Type),
		identity.NewID(q.ID),
		q.Text,
		q.Customs,
	)
	return sq.ToQuestion()
}
