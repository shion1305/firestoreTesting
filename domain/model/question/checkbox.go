package question

import (
	"errors"
	"firestoreTesting/domain/model/util"
	"firestoreTesting/pkg/identity"
)

type (
	CheckBoxOptionID util.ID

	CheckBoxQuestion struct {
		ID           ID
		Options      []CheckBoxOption
		OptionsOrder []CheckBoxOptionID
	}

	CheckBoxOption struct {
		ID   CheckBoxOptionID
		Text string
	}
)

func NewCheckBoxQuestion(id ID, options []CheckBoxOption) *CheckBoxQuestion {
	return &CheckBoxQuestion{
		ID:      id,
		Options: options,
	}
}

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
	return NewCheckBoxQuestion(q.ID, options), nil
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
