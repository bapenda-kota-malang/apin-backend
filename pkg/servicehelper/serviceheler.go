package servicehelper

import (
	"encoding/json"
	"errors"

	"go.uber.org/zap"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
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
