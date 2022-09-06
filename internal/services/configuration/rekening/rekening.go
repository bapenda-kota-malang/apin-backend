package rekening

import (
	"net/http"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/configurationmodel"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"gorm.io/gorm"
)

func GetAll(r *http.Request) (interface{}, error) {
	var rekenings []*configurationmodel.Rekening
	err := apicore.DB.Model(&configurationmodel.Rekening{}).Find(&rekenings).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return responses.OK{
		Data: rekenings,
	}, nil
}
