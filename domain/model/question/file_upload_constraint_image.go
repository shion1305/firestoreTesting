package question

import (
	"errors"
	"firestoreTesting/domain/service/image"
	imagePkg "firestoreTesting/pkg/image"
	"fmt"
	"path/filepath"
)

type (
	ImageFileConstraint struct {
		Ratio               float64
		MinNumber           int
		MaxNumber           int
		MinResolutionWidth  int
		MaxResolutionWidth  int
		MinResolutionHeight int
		MaxResolutionHeight int
		Extensions          []string
		PNGInfo             image.InfoExtractor
		JPGInfo             image.InfoExtractor
		WEBPInfo            image.InfoExtractor
	}
	ImageType int
)

const (
	PNG  ImageType = 1
	JPG  ImageType = 2
	WEBP ImageType = 3
)

func NewImageFileConstraint(
	ratio float64, minNumber, maxNumber, minWidth, maxWidth, minHeight, maxHeight int, extensions []string,
) ImageFileConstraint {
	return ImageFileConstraint{
		Ratio:               ratio,
		MinNumber:           minNumber,
		MaxNumber:           maxNumber,
		MinResolutionWidth:  minWidth,
		MaxResolutionWidth:  maxWidth,
		MinResolutionHeight: minHeight,
		MaxResolutionHeight: maxHeight,
		Extensions:          extensions,
		PNGInfo:             imagePkg.NewPNGInfo(),
		JPGInfo:             imagePkg.NewJPEGInfo(),
		WEBPInfo:            imagePkg.NewWEBPInfo(),
	}
}

func ImportImageFileConstraint(standard StandardFileConstraint) ImageFileConstraint {
	ratio, _ := standard.Options["ratio"].(float64)
	minNumber, _ := standard.Options["minNumber"].(int)
	maxNumber, _ := standard.Options["maxNumber"].(int)
	minWidth, _ := standard.Options["minWidth"].(int)
	maxWidth, _ := standard.Options["maxWidth"].(int)
	minHeight, _ := standard.Options["minHeight"].(int)
	maxHeight, _ := standard.Options["maxHeight"].(int)
	exts, _ := standard.Options["extensions"].([]string)

	return NewImageFileConstraint(
		ratio, minNumber, maxNumber, minWidth, maxWidth, minHeight, maxHeight, exts)
}

func (c ImageFileConstraint) Export() StandardFileConstraint {
	return NewStandardFileConstraint(Image,
		map[string]interface{}{
			"ratio":      c.Ratio,
			"minNumber":  c.MinNumber,
			"maxNumber":  c.MaxNumber,
			"minWidth":   c.MinResolutionWidth,
			"maxWidth":   c.MaxResolutionWidth,
			"minHeight":  c.MinResolutionHeight,
			"maxHeight":  c.MaxResolutionHeight,
			"extensions": c.Extensions,
		})
}

func (c ImageFileConstraint) GetType() FileType {
	return Image
}

func (c ImageFileConstraint) ValidateFile(filename string, file []byte) error {
	// get file extension
	ext := filepath.Ext(filename)
	imgType, err := c.checkExtension(ext)
	if err != nil {
		return err
	}
	err = c.validateProperties(imgType, file)
	if err != nil {
		return err
	}
	return nil
}

func (c ImageFileConstraint) validateProperties(imgType ImageType, file []byte) error {
	var width, height int
	var err error
	switch imgType {
	case PNG:
		width, height, err = c.PNGInfo.ExtractInfo(file)
	case JPG:
		width, height, err = c.JPGInfo.ExtractInfo(file)
	case WEBP:
		width, height, err = c.WEBPInfo.ExtractInfo(file)
	default:
		// skip check for unimplemented image type
		return nil
	}
	if err != nil {
		return err
	}
	if c.MinResolutionWidth > 0 && width < c.MinResolutionWidth {
		return errors.New(
			fmt.Sprintf("width not satisfied. min width: %d, actual width: %d", c.MinResolutionWidth, width))
	}
	if c.MaxResolutionWidth > 0 && width > c.MaxResolutionWidth {
		return errors.New(
			fmt.Sprintf("width not satisfied. max width: %d, actual width: %d", c.MaxResolutionWidth, width))
	}
	if c.MinResolutionHeight > 0 && height < c.MinResolutionHeight {
		return errors.New(
			fmt.Sprintf("height not satisfied. min height: %d, actual height: %d", c.MinResolutionHeight, height))
	}
	if c.MaxResolutionHeight > 0 && height > c.MaxResolutionHeight {
		return errors.New(
			fmt.Sprintf("height not satisfied. max height: %d, actual height: %d", c.MaxResolutionHeight, height))
	}
	if c.Ratio > 0 {
		if ratio := float64(width) / float64(height); ratio != c.Ratio {
			return errors.New(
				fmt.Sprintf("ratio not satisfied. expected ratio: %f, actual ratio: %f", c.Ratio, ratio))
		}
	}
	return nil
}

func (c ImageFileConstraint) checkExtension(ext string) (ImageType, error) {
	if len(c.Extensions) == 0 {
		// if extension is not specified, check with default extensions
		return convertToImageType(ext)
	}
	for _, e := range c.Extensions {
		if e == ext {
			return convertToImageType(ext)
		}
	}
	return 0, errors.New(
		fmt.Sprintf("invalid file type. specified extensions: %v", c.Extensions))
}

func convertToImageType(ext string) (ImageType, error) {
	switch ext {
	case ".jpg", ".jpeg":
		return JPG, nil
	case ".png":
		return PNG, nil
	case ".webp":
		return WEBP, nil
	default:
		return 0, errors.New(
			fmt.Sprintf("invalid file type. available extensions: %v",
				[]string{".jpg", ".jpeg", ".png", ".webp"}))
	}
}
