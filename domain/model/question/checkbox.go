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
	// CheckBoxOptionsField should be map[string]string

	// Although you cannot cast map[string]interface{} to map[string]string,
	// you have to iterate over the map and cast each value to string.
	// First, check if customs has CheckBoxOptionsField as map[string]interface{}, return error if not.
	optionsDataI, has := q.Customs[CheckBoxOptionsField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for CheckBoxQuestion", CheckBoxOptionsField))
	}
	optionsData, ok := optionsDataI.(map[string]interface{})
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[string]string for CheckBoxQuestion", CheckBoxOptionsField))
	}

	// check if customs has "order" as []int64, return error if not
	optionsOrderDataI, has := q.Customs[CheckBoxOptionsOrderField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for CheckBoxQuestion", CheckBoxOptionsOrderField))
	}
	optionsOrderData, ok := optionsOrderDataI.([]interface{})
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be []int64 for CheckBoxQuestion", CheckBoxOptionsOrderField))
	}

	options := make([]CheckBoxOption, 0, len(optionsData))
	optionsOrder := make([]CheckBoxOptionID, 0, len(optionsOrderData))
	for _, id := range optionsOrderData {
		i, ok := id.(int64)
		if !ok {
			return nil, errors.New(
				fmt.Sprintf("Option order must be int64 for CheckBoxQuestion"))
		}
		optionsOrder = append(optionsOrder, identity.NewID(i))
	}

	for id, textI := range optionsData {
		// here we cast textI to string
		text, ok := textI.(string)
		if !ok {
			return nil, errors.New(
				fmt.Sprintf("Option text must be string for CheckBoxQuestion"))
		}
		i, err := identity.ImportID(id)
		if err != nil {
			return nil, err
		}
		options = append(options, CheckBoxOption{
			ID:   i,
			Text: text,
		})
	}
	return NewCheckBoxQuestion(q.ID, q.Text, options, optionsOrder), nil
}

func (q CheckBoxQuestion) Export() StandardQuestion {
	customs := make(map[string]interface{})
	options := make(map[string]string, len(q.Options))
	for _, option := range q.Options {
		options[option.ID.ExportID()] = option.Text
	}
	optionsOrder := make([]int64, 0, len(q.OptionsOrder))
	for _, id := range q.OptionsOrder {
		optionsOrder = append(optionsOrder, id.GetValue())
	}
	customs[CheckBoxOptionsField] = options
	customs[CheckBoxOptionsOrderField] = optionsOrder
	return NewStandardQuestion(TypeCheckBox, q.ID, q.Text, customs)
}

func (q CheckBoxQuestion) GetType() Type {
	return TypeCheckBox
}

func (q CheckBoxQuestion) GetID() ID {
	return q.ID
}

func (q CheckBoxQuestion) GetText() string {
	return q.Text
}
