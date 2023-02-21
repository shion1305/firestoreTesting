package question

import (
	"firestoreTesting/domain/model/util"
)

type (
	ID       util.ID
	Type     int
	Question interface {
		GetID() ID
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
