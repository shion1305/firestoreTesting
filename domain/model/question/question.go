package question

import (
	"errors"
	"firestoreTesting/domain/model/util"
	"firestoreTesting/pkg/identity"
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

type (
	CheckBoxQuestion struct {
		ID           ID
		Options      []CheckBoxOption
		OptionsOrder []OptionID
	}

	OptionID util.ID

	CheckBoxOption struct {
		ID   OptionID
		Text string
	}

	RadioButtonsQuestion struct {
		ID      ID
		Options []RadioButtonOption
	}
	RadioButtonOption struct {
		ID   OptionID
		Text string
	}
)

func ImportCheckBoxQuestion(q StandardQuestion) (*CheckBoxQuestion, error) {
	// check if customs has "options" as map[int64]string, return error if not
	optionsDataI, has := q.Customs["options"]
	if !has {
		return nil, errors.New("\"options\" is required for CheckBoxQuestion")
	}
	optionsData, ok := optionsDataI.(map[int64]string)
	if !ok {
		return nil, errors.New("\"options\" must be map[int64]string for CheckBoxQuestion")
	}

	options := make([]CheckBoxOption, 0, len(optionsData))
	for id, text := range optionsData {
		options = append(options, CheckBoxOption{
			ID:   identity.NewID(id),
			Text: text,
		})
	}

	return &CheckBoxQuestion{
		ID:      q.ID,
		Options: options,
	}, nil
}

func (q CheckBoxQuestion) GetID() ID {
	return q.ID
}

func (q CheckBoxQuestion) Export() StandardQuestion {
	customs := make(map[string]interface{})
	options := make(map[int64]string, len(q.Options))
	for _, option := range q.Options {
		options[option.ID.GetValue()] = option.Text
	}
	customs["options"] = options
	return StandardQuestion{
		ID:      q.ID,
		Text:    q.Options[0].Text,
		Type:    TypeCheckBox,
		Customs: customs,
	}
}
