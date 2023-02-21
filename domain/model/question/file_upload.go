package question

type (
	FileUploadQuestion struct {
		ID         ID
		Text       string
		Constraint FileConstraint
	}
	FileType int
)

const (
	Image FileType = 1
	PDF   FileType = 2
)
