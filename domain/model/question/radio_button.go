package question

import (
	"errors"
	"firestoreTesting/domain/model/util"
	"firestoreTesting/pkg/identity"
	"fmt"
)

type (
	RadioButtonsQuestion struct {
		ID      ID
		Options []RadioButtonOption
	}
	RadioButtonOption struct {
		ID   RadioButtonOptionID
		Text string
	}
	RadioButtonOptionID util.ID
)

const RadioButtonOptionsField = "options"

func NewRadioButtonsQuestion(id ID, options []RadioButtonOption) *RadioButtonsQuestion {
	return &RadioButtonsQuestion{
		ID:      id,
		Options: options,
	}
}

func ImportRadioButtonsQuestion(q StandardQuestion) (*RadioButtonsQuestion, error) {
	// check if customs has "options" as map[int64]string, return error if not
	optionsDataI, has := q.Customs[RadioButtonOptionsField]
	if !has {
		return nil, errors.New(fmt.Sprintf("\"%s\" is required for RadioButtonsQuestion", RadioButtonOptionsField))
	}
	optionsData, ok := optionsDataI.(map[int64]string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("\"%s\" must be map[int64]string for RadioButtonsQuestion", RadioButtonOptionsField))
	}

	options := make([]RadioButtonOption, 0, len(optionsData))
	for id, text := range optionsData {
		options = append(options, RadioButtonOption{
			ID:   identity.NewID(id),
			Text: text,
		})
	}
	return NewRadioButtonsQuestion(q.ID, options), nil
}
