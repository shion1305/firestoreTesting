package question

import (
	"firestoreTesting/domain/model/util"
)

type (
	ID       util.ID
	Type     int
	Question interface {
		Export() StandardQuestion
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
)

func NewStandardQuestion(t Type, id ID, text string, customs map[string]interface{}) StandardQuestion {
	return StandardQuestion{
		ID:      id,
		Text:    text,
		Type:    t,
		Customs: customs,
	}
}
