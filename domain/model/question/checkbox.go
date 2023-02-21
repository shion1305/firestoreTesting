package question

import (
	"errors"
	"firestoreTesting/domain/model/util"
	"firestoreTesting/pkg/identity"
	"fmt"
)

type (
	CheckBoxOptionID util.ID

	CheckBoxQuestion struct {
		ID           ID
		Text         string
		Options      []CheckBoxOption
		OptionsOrder []CheckBoxOptionID
	}

	CheckBoxOption struct {
		ID   CheckBoxOptionID
		Text string
	}
)

const (
	CheckBoxOptionsField      = "options"
	CheckBoxOptionsOrderField = "order"
)

func NewCheckBoxQuestion(
	id ID, text string, options []CheckBoxOption, optionsOrder []CheckBoxOptionID,
) *CheckBoxQuestion {
	return &CheckBoxQuestion{
		ID:           id,
		Text:         text,
		Options:      options,
		OptionsOrder: optionsOrder,
	}
}

func ImportCheckBoxQuestion(q StandardQuestion) (*CheckBoxQuestion, error) {
	// check if customs has CheckBoxOptionsField as map[int64]string, return error if not
	optionsDataI, has := q.Customs[CheckBoxOptionsField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for CheckBoxQuestion", CheckBoxOptionsField))
	}
	optionsData, ok := optionsDataI.(map[int64]string)
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[int64]string for CheckBoxQuestion", CheckBoxOptionsField))
	}

	// check if customs has "order" as []int64, return error if not
	optionsOrderDataI, has := q.Customs[CheckBoxOptionsOrderField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for CheckBoxQuestion", CheckBoxOptionsOrderField))
	}
	optionsOrderData, ok := optionsOrderDataI.([]int64)
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be []int64 for CheckBoxQuestion", CheckBoxOptionsOrderField))
	}

	options := make([]CheckBoxOption, 0, len(optionsData))
	optionsOrder := make([]CheckBoxOptionID, 0, len(optionsOrderData))
	for _, id := range optionsOrderData {
		optionsOrder = append(optionsOrder, identity.NewID(id))
	}

	for id, text := range optionsData {
		options = append(options, CheckBoxOption{
			ID:   identity.NewID(id),
			Text: text,
		})
	}
	return NewCheckBoxQuestion(q.ID, q.Text, options, optionsOrder), nil
}

func (q CheckBoxQuestion) Export() StandardQuestion {
	customs := make(map[string]interface{})
	options := make(map[int64]string, len(q.Options))
	for _, option := range q.Options {
		options[option.ID.GetValue()] = option.Text
	}
	customs[CheckBoxOptionsField] = options
	return NewStandardQuestion(TypeCheckBox, q.ID, q.Text, customs)
}
