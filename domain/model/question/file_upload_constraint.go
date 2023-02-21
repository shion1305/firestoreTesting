package question

type (
	StandardFileConstraint struct {
		FileType FileType
		Options  map[string]interface{}
	}
	FileConstraint interface {
		GetFileType() FileType
		Export() StandardFileConstraint
		ValidateFiles(filename string, file [][]byte) error
	}
)

func NewStandardFileConstraint(fileType FileType, options map[string]interface{}) StandardFileConstraint {
	return StandardFileConstraint{
		FileType: fileType,
		Options:  options,
	}
}
