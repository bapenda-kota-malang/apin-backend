package logbankjatim

import (
	"encoding/json"
	"time"

	"gorm.io/datatypes"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/logbankjatim"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "logbankjatim"

func Create(input m.CreateDynamicDto) (any, error) {
	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	cb, err := json.Marshal(input.Content)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal marshal json content:"+err.Error(), input.Content)
	}
	rb, err := json.Marshal(input.Response)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal marshal json response:"+err.Error(), input.Response)
	}
	data := m.LogBankJatim{
		Path:         input.Path,
		Content:      datatypes.JSON(cb),
		Response:     datatypes.JSON(rb),
		ResponseCode: input.ResponseCode,
		User_Id:      input.User_Id,
		Timestamp:    time.Now(),
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := a.DB.Create(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data log"+err.Error(), data)
	}

	return data, nil
}
