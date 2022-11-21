package pengurangan

import (
	"errors"

	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

func FilePreProcess(b64String string, userId uint, docsName string) (fileName, path, extFile string, err error) {
	extFile, err = base64helper.GetExtensionBase64(b64String)
	if err != nil {
		return
	}
	path = sh.GetResourcesPath()
	switch extFile {
	case "png", "jpeg":
		path = sh.GetImgPath()
	default:
		err = errors.New("file bukan gambar")
		return
	}
	id, err := sh.GetUuidv4()
	if err != nil {
		err = errors.New("gagal generate uuid")
		return
	}
	fileName = sh.GenerateFilename(docsName, id, userId, extFile)
	return
}
