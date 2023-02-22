package question

import (
	"errors"
	"firestoreTesting/domain/model/util"
)

type (
	ID       util.ID
	Type     int
	Question interface {
		Export() StandardQuestion
		GetType() Type
		GetID() ID
		GetText() string
	}

	StandardQuestion struct {
		ID      ID
		Text    string
		Type    Type
		Customs map[string]interface{}
	}
)

const (
	TypeCheckBox Type = 1
	TypeRadio    Type = 2
	TypeFile     Type = 3
)

func NewStandardQuestion(t Type, id ID, text string, customs map[string]interface{}) StandardQuestion {
	return StandardQuestion{
		ID:      id,
		Text:    text,
		Type:    t,
		Customs: customs,
	}
}

func (q StandardQuestion) ToQuestion() (Question, error) {
	switch q.Type {
	case TypeCheckBox:
		return ImportCheckBoxQuestion(q)
	case TypeRadio:
		return ImportRadioButtonsQuestion(q)
	case TypeFile:
		return ImportFileQuestion(q)
	default:
		return nil, errors.New("invalid question type")
	}
}
