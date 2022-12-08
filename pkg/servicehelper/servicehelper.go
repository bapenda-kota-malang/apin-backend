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
	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	"github.com/google/uuid"
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

func GetUuidv4() (id uuid.UUID, err error) {
	return uuid.NewRandom()
}

// create filename with specific pattern: docsName_uuid_YYYYMMDD_userId.ext
func GenerateFilename(docsName string, uuid uuid.UUID, userId uint, extension string) (fileName string) {
	curTime := time.Now()
	if userId == 0 {
		fileName = fmt.Sprintf("%s_%s_%s.%s", docsName, uuid.String(), curTime.Format("20060102"), extension)
	} else {
		fileName = fmt.Sprintf("%s_%s_%s_%d.%s", docsName, uuid.String(), curTime.Format("20060102"), userId, extension)
	}
	return
}

// Get path based on file name
func GetPathByFilename(filename string) (filePath string) {
	filePath = GetResourcesPath()
	ext := filepath.Ext(filename)
	switch ext {
	case ".pdf":
		filePath = GetPdfPath()
	case ".jpeg", ".png":
		filePath = GetImgPath()
	case ".xlsx", ".xls":
		filePath = GetExcelPath()
	}
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

func getSpecificPath(typeFile string) string {
	// path
	resourcesPath := GetResourcesPath()
	basePath := filepath.Join(resourcesPath, typeFile)
	os.MkdirAll(basePath, os.ModePerm)
	return basePath
}

// get path for resource/img folder
func GetImgPath() string {
	return getSpecificPath("img")
}

// get path for resource/pdf folder
func GetPdfPath() string {
	return getSpecificPath("pdf")
}

// get path for resource/pdf folder
func GetExcelPath() string {
	return getSpecificPath("excel")
}

func RemoveFile(path, filename string) error {
	return os.Remove(fmt.Sprintf("%s/%s", path, filename))
}

// save base64 string to specific file
func SaveFile(b64Raw, fileName, path, extFile string, errCh chan error) {
	defer close(errCh)
	coI := strings.Index(string(b64Raw), ",")
	rawFiles := string(b64Raw)[coI+1:]
	dec, err := base64.StdEncoding.DecodeString(rawFiles)
	if err != nil {
		errCh <- err
		return
	}
	filePath := fmt.Sprintf("%s/%s", path, fileName)

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		errCh <- err
		return
	}
	defer f.Close()

	switch extFile {
	case "png":
		bReader := bytes.NewReader(dec)
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
	case "jpeg":
		bReader := bytes.NewReader(dec)
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
	case "pdf", "xlsx", "xls":
		if _, err := f.Write(dec); err != nil {
			errCh <- err
			return
		}
		if err := f.Sync(); err != nil {
			errCh <- err
			return
		}
	default:
		errCh <- errors.New("unsupported extension type")
		return
	}

	errCh <- nil
}

func ReplaceFile(oldFileName, b64Raw, newFileName, path, extFile string, errCh chan error) {
	defer close(errCh)
	saveErrCh := make(chan error)

	oldFileNameSave := oldFileName
	if oldFileName == newFileName {
		oldPath := filepath.Join(path, oldFileName)
		tmpOldFileName := "tmp" + oldFileName
		newPathTmp := filepath.Join(path, tmpOldFileName)
		if err := os.Rename(oldPath, newPathTmp); err != nil {
			errCh <- err
			return
		}
		oldFileName = tmpOldFileName
	}

	go SaveFile(b64Raw, newFileName, path, extFile, saveErrCh)
	if err := <-saveErrCh; err != nil {
		_ = os.Rename(filepath.Join(path, oldFileName), filepath.Join(path, oldFileNameSave))
		errCh <- err
		return
	}
	if err := RemoveFile(path, oldFileName); err != nil {
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

// get next midnight
//
// e.g now = 2022-10-11 22:33:44 +0700 WIB
//
// Midnight(now) = 2022-10-11 00:00:00.000000000 +0700 WIB
func Midnight(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// loop process to save image file from slice base64 and return slice filename
func loopArrayPhoto(input []string, docsName string, userId uint) (arrString []string, err error) {
	for v := range input {
		var errChan = make(chan error)
		path := GetImgPath()
		id, err2 := GetUuidv4()
		if err2 != nil {
			err = err2
			return
		}
		extFile, err2 := base64helper.GetExtensionBase64(input[v])
		if err2 != nil {
			err = err2
			return
		}
		switch extFile {
		case "png", "jpeg":
			fileName := GenerateFilename(docsName, id, userId, extFile)
			go SaveFile(input[v], fileName, path, extFile, errChan)
			if err2 := <-errChan; err2 != nil {
				err = err2
				return
			}
			arrString = append(arrString, fileName)
		default:
			err = fmt.Errorf("unsupported type for this process")
			return
		}
	}
	return
}

// process to convert string slice to json bytes and back it to raw string
func arrToJsonString(slc []string) string {
	bytes, _ := json.Marshal(slc)
	return string(bytes)
}

// array photo
func GetArrayPhoto(input []string, docsName string, userId uint) (arrString string, err error) {
	result, err := loopArrayPhoto(input, docsName, userId)
	if err == nil {
		arrString = arrToJsonString(result)
	}
	return
}

// array file
func GetArrayFile(input []string, docsName string, userId uint) (arrString string, err error) {
	result, err := loopArrayFile(input, docsName, userId)
	if err == nil {
		arrString = arrToJsonString(result)
	}
	return
}

// array pdf
func GetArrayPdfAndImage(input []string, docsName string, userId uint) (arrString string, err error) {
	result, err := loopArrayPdfAndImage(input, docsName, userId)
	if err == nil {
		arrString = arrToJsonString(result)
	}
	return
}

// loop process to save image file from slice base64 and return slice filename
func loopArrayFile(input []string, docsName string, userId uint) (arrString []string, err error) {
	for v := range input {
		var errChan = make(chan error)
		imagePath := GetImgPath()
		pdfPath := GetPdfPath()
		excelPath := GetExcelPath()
		id, err2 := GetUuidv4()
		if err2 != nil {
			err = err2
			return
		}
		extFile, err2 := base64helper.GetExtensionBase64(input[v])
		if err2 != nil {
			err = err2
			return
		}
		switch extFile {
		case "png", "jpeg":
			fileName := GenerateFilename(docsName, id, userId, extFile)
			go SaveFile(input[v], fileName, imagePath, extFile, errChan)
			if err2 := <-errChan; err2 != nil {
				err = err2
				return
			}
			arrString = append(arrString, fileName)
		case "pdf":
			fileName := GenerateFilename(docsName, id, userId, extFile)
			go SaveFile(input[v], fileName, pdfPath, extFile, errChan)
			if err2 := <-errChan; err2 != nil {
				err = err2
				return
			}
			arrString = append(arrString, fileName)
		case "xlsx", "xls":
			fileName := GenerateFilename(docsName, id, userId, extFile)
			go SaveFile(input[v], fileName, excelPath, extFile, errChan)
			if err2 := <-errChan; err2 != nil {
				err = err2
				return
			}
			arrString = append(arrString, fileName)
		default:
			err = fmt.Errorf("unsupported type for this process")
			return
		}
	}
	return
}

// loop process to save image file from slice base64 and return slice filename
func loopArrayPdfAndImage(input []string, docsName string, userId uint) (arrString []string, err error) {
	for v := range input {
		var errChan = make(chan error)
		pdfPath := GetPdfPath()
		imagePath := GetImgPath()
		id, err2 := GetUuidv4()
		if err2 != nil {
			err = err2
			return
		}
		extFile, err2 := base64helper.GetExtensionBase64(input[v])
		if err2 != nil {
			err = err2
			return
		}
		switch extFile {
		case "pdf":
			fileName := GenerateFilename(docsName, id, userId, extFile)
			go SaveFile(input[v], fileName, pdfPath, extFile, errChan)
			if err2 := <-errChan; err2 != nil {
				err = err2
				return
			}
			arrString = append(arrString, fileName)
		case "png", "jpeg":
			fileName := GenerateFilename(docsName, id, userId, extFile)
			go SaveFile(input[v], fileName, imagePath, extFile, errChan)
			if err2 := <-errChan; err2 != nil {
				err = err2
				return
			}
			arrString = append(arrString, fileName)
		default:
			err = fmt.Errorf("unsupported type for this process")
			return
		}
	}
	return
}

// add multiply photo
func AddMorePhotos(input []string, dataBefore, docsName string, userId uint) (arrString string, err error) {
	var result []string
	err = json.Unmarshal([]byte(dataBefore), &result)
	if err != nil {
		return
	}
	newData, err := loopArrayPhoto(input, docsName, userId)
	if err == nil {
		result = append(result, newData...)
		arrString = arrToJsonString(result)
	}
	return
}

// add multiply multiextension
func AddMoreFile(input []string, dataBefore, docsName string, userId uint) (arrString string, err error) {
	var result []string
	err = json.Unmarshal([]byte(dataBefore), &result)
	if err != nil {
		return
	}
	newData, err := loopArrayFile(input, docsName, userId)
	if err == nil {
		result = append(result, newData...)
		arrString = arrToJsonString(result)
	}
	return
}

// add multiply pdf
func AddMorePdfAndImage(input []string, dataBefore, docsName string, userId uint) (arrString string, err error) {
	var result []string
	err = json.Unmarshal([]byte(dataBefore), &result)
	if err != nil {
		return
	}
	newData, err := loopArrayPdfAndImage(input, docsName, userId)
	if err == nil {
		result = append(result, newData...)
		arrString = arrToJsonString(result)
	}
	return
}

// remove photo from array
func DeleteFile(input string, data string) (string, error) {
	var result []string
	var err error
	cnvUnmarshal := json.Unmarshal([]byte(data), &result)
	if cnvUnmarshal != nil {
		return "", errors.New("data tidak bisa diunmarshal")
	}

	for k, v := range result {
		if v == input {
			result = append(result[:k], result[k+1:]...)
		}
	}

	tmpSplit := strings.Split(input, ".")
	switch tmpSplit[1] {
	case "png", "jpeg":

		basePath := GetImgPath()

		if err := RemoveFile(basePath, input); err != nil {
			return "", errors.New("tidak dapat menghapus data dari resourse, filename tidak ditemukan")
		}
	case "pdf":
		basePath := GetPdfPath()

		if err := RemoveFile(basePath, input); err != nil {
			return "", errors.New("tidak dapat menghapus data dari resourse, filename tidak ditemukan")
		}
	case "xlxs", "xls":
		basePath := GetExcelPath()

		if err := RemoveFile(basePath, input); err != nil {
			return "", errors.New("tidak dapat menghapus data dari resourse, filename tidak ditemukan")
		}
	default:
		err = fmt.Errorf("unsupported type for this process")
		return func() string {
			bytes, _ := json.Marshal(result)
			return string(bytes)
		}(), err
	}

	bytes, _ := json.Marshal(result)
	return string(bytes), nil
}

// loop process to save image, excel, pdf file from base64 and return filename
func allTypeFile(input string, docsName string, userId uint) (resultString string, err error) {
	var errChan = make(chan error)
	imagePath := GetImgPath()
	pdfPath := GetPdfPath()
	excelPath := GetExcelPath()
	id, err2 := GetUuidv4()
	if err2 != nil {
		err = err2
		return
	}
	extFile, err2 := base64helper.GetExtensionBase64(input)
	if err2 != nil {
		err = err2
		return
	}
	switch extFile {
	case "png", "jpeg":
		fileName := GenerateFilename(docsName, id, userId, extFile)
		go SaveFile(input, fileName, imagePath, extFile, errChan)
		if err2 := <-errChan; err2 != nil {
			err = err2
			return
		}
		resultString = fileName
	case "pdf":
		fileName := GenerateFilename(docsName, id, userId, extFile)
		go SaveFile(input, fileName, pdfPath, extFile, errChan)
		if err2 := <-errChan; err2 != nil {
			err = err2
			return
		}
		resultString = fileName
	case "xlsx", "xls":
		fileName := GenerateFilename(docsName, id, userId, extFile)
		go SaveFile(input, fileName, excelPath, extFile, errChan)
		if err2 := <-errChan; err2 != nil {
			err = err2
			return
		}
		resultString = fileName
	default:
		err = fmt.Errorf("unsupported type for this process")
		return
	}

	return
}

// all type file
func GetAllTypeFile(input string, docsName string, userId uint) (resultString string, err error) {
	result, err := allTypeFile(input, docsName, userId)
	if err == nil {
		resultString = result
	}
	return
}

// loop process to save image, pdf file from base64 and return filename
func pdfAndImageFile(input string, docsName string, userId uint) (resultString string, err error) {
	var errChan = make(chan error)
	imagePath := GetImgPath()
	pdfPath := GetPdfPath()
	id, err2 := GetUuidv4()
	if err2 != nil {
		err = err2
		return
	}
	extFile, err2 := base64helper.GetExtensionBase64(input)
	if err2 != nil {
		err = err2
		return
	}
	switch extFile {
	case "png", "jpeg":
		fileName := GenerateFilename(docsName, id, userId, extFile)
		go SaveFile(input, fileName, imagePath, extFile, errChan)
		if err2 := <-errChan; err2 != nil {
			err = err2
			return
		}
		resultString = fileName
	case "pdf":
		fileName := GenerateFilename(docsName, id, userId, extFile)
		go SaveFile(input, fileName, pdfPath, extFile, errChan)
		if err2 := <-errChan; err2 != nil {
			err = err2
			return
		}
		resultString = fileName
	default:
		err = fmt.Errorf("unsupported type for this process")
		return
	}

	return
}

// pdf and image type file
func GetPdfOrImageFile(input string, docsName string, userId uint) (resultString string, err error) {
	result, err := pdfAndImageFile(input, docsName, userId)
	if err == nil {
		resultString = result
	}
	return
}

func FixedLengthString(length int, str string) string {
	var verb string = str
	if len(str) < length {
		tempL := length - len(str)
		for i := 0; i < tempL; i++ {
			verb = "0" + verb
		}
	}
	return verb
}

func NopParser(nop string) (result []string, area_kode string) {
	result = strings.Split(nop, ".")
	area_kode = result[0] + result[1] + result[2] + result[3]
	return result, area_kode
}
