package question

import (
	"errors"
	"fmt"
)

type (
	FileQuestion struct {
		ID         ID
		Text       string
		Constraint FileConstraint
	}
	FileType int
)

const (
	Image                       FileType = 1
	PDF                         FileType = 2
	Any                         FileType = 3
	FileQuestionFileTypeField            = "fileType"
	FileConstraintsOptionsField          = "fileConstraints"
)

func NewFileQuestion(id ID, text string, constraint FileConstraint) *FileQuestion {
	return &FileQuestion{
		ID:         id,
		Text:       text,
		Constraint: constraint,
	}
}

func NewFileType(v int) (FileType, error) {
	switch FileType(v) {
	case Image, PDF, Any:
		return FileType(v), nil
	}
	return 0, errors.New("invalid file type")
}

func ImportFileQuestion(q StandardQuestion) (*FileQuestion, error) {
	// check if customs has "fileType" as int, return error if not
	fileTypeDataI, has := q.Customs[FileQuestionFileTypeField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for FileQuestion", FileQuestionFileTypeField))
	}
	fileTypeData, ok := fileTypeDataI.(int)
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be int for FileQuestion", FileQuestionFileTypeField))
	}
	fileType, err := NewFileType(fileTypeData)
	if err != nil {
		return nil, err
	}

	if fileType == Any {
		return NewFileQuestion(q.ID, q.Text, nil), nil
	}

	constraintsOptionsData, has := q.Customs[FileConstraintsOptionsField]
	if !has {
		return NewFileQuestion(q.ID, q.Text, nil), nil
	}

	constraintsOptions, ok := constraintsOptionsData.(map[string]interface{})
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[string]interface{} for FileQuestion", FileConstraintsOptionsField))
	}
	constraint := NewStandardFileConstraint(fileType, constraintsOptions)
	question := NewFileQuestion(q.ID, q.Text, ImportFileConstraint(constraint))
	if err != nil {
		return nil, err
	}
	return question, nil
}
