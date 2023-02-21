package question

type (
	StandardFileConstraint struct {
		FileType FileType
		Options  map[string]interface{}
	}
	FileConstraint interface {
		GetFileType() FileType
		Export() StandardFileConstraint
		ValidateFiles(file []File) error
	}
	File struct {
		FileName string
		Data     []byte
	}
)

func NewStandardFileConstraint(fileType FileType, options map[string]interface{}) StandardFileConstraint {
	return StandardFileConstraint{
		FileType: fileType,
		Options:  options,
	}
}

func ImportFileConstraint(standard StandardFileConstraint) FileConstraint {
	switch standard.FileType {
	case Image:
		return ImportImageFileConstraint(standard)
	default:
		return nil
	}
}
