package servicehelper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"github.com/gofrs/uuid"
)

func SetError(scope, xtype, source, status, message string, data any) (any, error) {
	dataString, _ := json.Marshal(data)
	a.Logger.Error(scope,
		zap.String("type", xtype),
		zap.String("source", source),
		zap.String("status", status),
		zap.String("message", message),
		zap.String("data", string(dataString)))
	return nil, errors.New(message)
}

func getImgPath() (string, error) {
	// path
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	basePath := filepath.Join(wd, "../../", "resources", "img")
	os.MkdirAll(basePath, os.ModePerm)
	return basePath, nil
}

func SaveImage(b64Raw string, imgNameCh chan string, errCh chan error) {
	defer close(imgNameCh)
	defer close(errCh)
	coI := strings.Index(string(b64Raw), ",")
	rawImage := string(b64Raw)[coI+1:]
	unbased, err := base64.StdEncoding.DecodeString(rawImage)
	if err != nil {
		errCh <- err
		return
	}
	bReader := bytes.NewReader(unbased)

	basePath, err := getImgPath()
	if err != nil {
		errCh <- err
		return
	}

	// create img name
	uid, err := uuid.NewV4()
	if err != nil {
		errCh <- err
		return
	}
	imgName := uid.String()

	// write img according file type
	switch strings.TrimSuffix(b64Raw[5:coI], ";base64") {
	case "image/png":
		imgName += ".png"
		f, err := os.OpenFile(fmt.Sprintf("%s/%s", basePath, imgName), os.O_WRONLY|os.O_CREATE, 0777)
		defer f.Close()
		pngI, err := png.Decode(bReader)
		if err != nil {
			errCh <- err
			return
		}
		err = png.Encode(f, pngI)
		if err != nil {
			errCh <- err
			return
		}
	case "image/jpeg":
		imgName += ".jpeg"
		f, err := os.OpenFile(fmt.Sprintf("%s/%s", basePath, imgName), os.O_WRONLY|os.O_CREATE, 0777)
		defer f.Close()
		jpgI, err := jpeg.Decode(bReader)
		if err != nil {
			errCh <- err
			return
		}
		err = jpeg.Encode(f, jpgI, &jpeg.Options{Quality: 75})
		if err != nil {
			errCh <- err
			return
		}
	default:
		errCh <- errors.New("unsupported mime type")
		return
	}
	errCh <- nil
	imgNameCh <- imgName
}

func ReplaceImage(oImgName, b64Raw string, imgNameCh chan string, errCh chan error) {
	defer close(errCh)
	saveErrCh := make(chan error)
	go SaveImage(b64Raw, imgNameCh, saveErrCh)

	basePath, err := getImgPath()
	if err != nil {
		errCh <- err
		return
	}

	if err := <-saveErrCh; err != nil {
		errCh <- err
		return
	}

	if err := os.Remove(fmt.Sprintf("%s/%s", basePath, oImgName)); err != nil {
		errCh <- err
		return
	}

}
