package base64helper

import (
	"errors"
	"strings"
)

func GetExtensionBase64(b64Raw string) (extension string, err error) {
	coI := strings.Index(string(b64Raw), ",")
	switch strings.TrimSuffix(b64Raw[5:coI], ";base64") {
	case "image/png":
		extension = "png"
		return
	case "image/jpeg":
		extension = "jpeg"
		return
	case "application/pdf":
		extension = "pdf"
		return
	case "vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		extension = "xlsx"
		return
	case "msexcel":
		extension = "xls"
		return
	default:
		err = errors.New("unsupported mime type")
	}
	return
}
