package static

import (
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	mspt "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	sspt "github.com/bapenda-kota-malang/apin-backend/internal/services/spt"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"github.com/go-chi/chi/v5"
)

func AuthFile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		fileName := chi.URLParam(r, "filename")
		re := regexp.MustCompile(`^(\w+)_(.*)_\d+_(\d+)\.\w+$`)
		fileNameGroup := re.FindStringSubmatch(fileName)
		authIdFile, err := strconv.Atoi(fileNameGroup[3])
		if err != nil {
			http.Error(w, "gagal mendapatkan auth file id", http.StatusBadRequest)
			return
		}
		authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
		userId := authInfo.User_Id
		if authIdFile != userId {
			validId := false
			// oh no auth id in file name not same with auth id by login
			// check deeper based on prefix, and id in file name
			// should i move this to seperate handler(?) idk
			id, pass := handlerhelper.ValidateIdUuid(w, fileNameGroup[2])
			if !pass {
				http.Error(w, "file tidak memiliki uuid valid", http.StatusBadRequest)
				return
			}
			prefixFilename := fileNameGroup[1]
			if strings.Contains(prefixFilename, "sptpd") ||
				strings.Contains(prefixFilename, "skpd") ||
				strings.Contains(prefixFilename, "skpdkb") {
				respSpt, err := sspt.GetDetail(id, "", uint(userId))
				if err != nil {
					if respSpt == nil {
						http.Error(w, "file not found", http.StatusNotFound)
					} else {
						http.Error(w, err.Error(), http.StatusUnprocessableEntity)
					}
					return
				}
				dataSpt := respSpt.(rp.OKSimple).Data.(*mspt.Spt)
				if dataSpt != nil && *dataSpt.Npwpd.User_Id == uint64(userId) {
					validId = true
				}
			}
			if !validId {
				http.Error(w, "user tidak memiliki akses terhadap file", http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func JoinPrefix(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fileName := chi.URLParam(r, "filename")
		ext := filepath.Ext(fileName)
		prefixFolder := ""
		switch ext {
		case ".pdf":
			prefixFolder = "pdf"
		case ".xls", ".xlsx":
			prefixFolder = "excel"
		case ".jpeg", ".png":
			prefixFolder = "img"
		default:
			http.NotFound(w, r)
			return
		}
		newReq := new(http.Request)
		*newReq = *r
		newReq.URL = new(url.URL)
		*newReq.URL = *r.URL
		newReq.URL.Path = prefixFolder + "/" + r.URL.Path
		newReq.URL.RawPath = r.URL.RawPath

		next.ServeHTTP(w, newReq)
	})
}
