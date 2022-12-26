package sql

import (
	"strings"
	"testing"
)

func TestNewGenericCon(t *testing.T) {
	type args struct {
		driverName     string
		dataSourceName string
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		errConclusion string
	}{
		{name: "failed unknown driverName", args: args{driverName: "", dataSourceName: ""}, wantErr: true, errConclusion: "unknown driver"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewGenericCon(tt.args.driverName, tt.args.dataSourceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGenericCon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) && !strings.Contains(err.Error(), tt.errConclusion) {
				t.Errorf("NewGenericCon() error = %v, errConclusion %v", err, tt.errConclusion)
				return
			}
		})
	}
}
