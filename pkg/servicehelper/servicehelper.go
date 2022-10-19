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
	"time"

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

func GetUuidv4() (id string, err error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return
	}
	id = uid.String()
	return
}

// get path for resource root folder
func GetResourcesPath() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	basePath := filepath.Join(wd, "../../", "resources")
	os.MkdirAll(basePath, os.ModePerm)
	return basePath
}

// get path for resource/img folder
func getImgPath() (string, error) {
	// path
	resourcesPath := GetResourcesPath()
	basePath := filepath.Join(resourcesPath, "img")
	os.MkdirAll(basePath, os.ModePerm)
	return basePath, nil
}

// get path for resource/pdf folder
func GetPdfPath() string {
	// path
	resourcesPath := GetResourcesPath()
	basePath := filepath.Join(resourcesPath, "pdf")
	os.MkdirAll(basePath, os.ModePerm)
	return basePath
}

// check base64 is image with supported format
func CheckImage(b64Raw string) (err error) {
	coI := strings.Index(string(b64Raw), ",")
	switch strings.TrimSuffix(b64Raw[5:coI], ";base64") {
	case "image/png":
		return
	case "image/jpeg":
		return
	default:
		err = errors.New("unsupported mime type")
	}
	return
}

// check base64 is pdf format
func CheckPdf(b64Raw string) (err error) {
	coI := strings.Index(string(b64Raw), ",")
	if strings.TrimSuffix(b64Raw[5:coI], ";base64") != "application/pdf" {
		err = errors.New("unsupported mime type")
	}
	return
}

// save new image from base64 string to img eiter jpeg/png to resource/img folder with filename uuidv4 generated
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
	uid, err := GetUuidv4()
	if err != nil {
		errCh <- err
		return
	}
	imgName := uid

	// write img according file type
	switch strings.TrimSuffix(b64Raw[5:coI], ";base64") {
	case "image/png":
		imgName += ".png"
		f, err := os.OpenFile(fmt.Sprintf("%s/%s", basePath, imgName), os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			errCh <- err
			return
		}
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
		if err != nil {
			errCh <- err
			return
		}
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

// save new image and delete old image from resource/img folder
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
	errCh <- nil
}

// save base64 string to file
func SaveFile(b64Raw, fileName, path string, errCh chan error) {
	defer close(errCh)
	coI := strings.Index(string(b64Raw), ",")
	rawFiles := string(b64Raw)[coI+1:]
	dec, err := base64.StdEncoding.DecodeString(rawFiles)
	if err != nil {
		errCh <- err
		return
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", path, fileName))
	if err != nil {
		errCh <- err
		return
	}
	defer f.Close()
	if _, err := f.Write(dec); err != nil {
		errCh <- err
		return
	}
	if err := f.Sync(); err != nil {
		errCh <- err
		return
	}
	errCh <- nil
}

func ReplaceFile(oldFileName, b64Raw, newFileName, path string, errCh chan error) {
	defer close(errCh)
	saveErrCh := make(chan error)

	go SaveFile(b64Raw, newFileName, path, saveErrCh)
	if err := <-saveErrCh; err != nil {
		errCh <- err
		return
	}
	if err := os.Remove(fmt.Sprintf("%s/%s", path, oldFileName)); err != nil {
		errCh <- err
		return
	}
	errCh <- nil
}

// get beginning of month
//
// e.g now = 2022-10-11 22:33:44 +0700 WIB
//
// BeginningOfMonth(now) = 2022-10-01 00:00:00 +0700 WIB
func BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// get beginning of previous month
//
// e.g now = 2022-10-11 22:33:44 +0700 WIB
//
// BeginningOfPreviosMonth() = 2022-09-01 00:00:00 +0700 WIB
func BeginningOfPreviosMonth() time.Time {
	return BeginningOfMonth(time.Now()).AddDate(0, -1, 0)
}

// get end of month
//
// e.g now = 2022-10-11 22:33:44 +0700 WIB
//
// EndOfMonth(now) = 2022-10-31 23:59:59.999999999 +0700 WIB
func EndOfMonth(date time.Time) time.Time {
	return BeginningOfMonth(date).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// array photo
func GetArrayPhoto(input []string) string {
	var result []string
	for _, v := range input {
		var imgNameChan = make(chan string)
		var errChan = make(chan error)

		go SaveImage(v, imgNameChan, errChan)
		if err := <-errChan; err != nil {
			return ""
		}
		var tmp string = <-imgNameChan
		result = append(result, tmp)
	}

	bytes, _ := json.Marshal(result)
	return string(bytes)
}

// add multiply photo
func AddMorePhotos(input []string, dataBefore string) string {
	var result []string
	cnvUnmarshal := json.Unmarshal([]byte(dataBefore), &result)
	if cnvUnmarshal != nil {
		return ""
	}
	for _, v := range input {
		var imgNameChan = make(chan string)
		var errChan = make(chan error)

		go SaveImage(v, imgNameChan, errChan)
		if err := <-errChan; err != nil {
			return ""
		}
		var tmp string = <-imgNameChan
		result = append(result, tmp)
	}

	bytes, _ := json.Marshal(result)
	return string(bytes)
}

// remove photo from array
func DeletePhoto(input string, data string) string {
	var result []string
	cnvUnmarshal := json.Unmarshal([]byte(data), &result)
	if cnvUnmarshal != nil {
		return ""
	}
	for k, v := range result {
		if v == input {
			result = append(result[:k], result[k+1:]...)

		}
	}
	// basePath, err := getImgPath()
	// if err != nil {
	// 	fmt.Println("HERE")
	// 	return ""
	// }

	// if err := os.Remove(fmt.Sprintf("%s/%s", basePath, input)); err != nil {
	// 	fmt.Println("here")
	// 	return ""
	// }

	bytes, _ := json.Marshal(result)
	return string(bytes)
}
