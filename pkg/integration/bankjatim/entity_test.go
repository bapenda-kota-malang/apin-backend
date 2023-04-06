package bankjatim

import (
	"testing"
)

func TestRequestRegistration_Verify(t *testing.T) {
	validVa := "1234567890123456"
	validTglExp := "20230309"
	type fields struct {
		VirtualAccount string
		Nama           string
		TotalTagihan   uint64
		TanggalExp     string
		Berita1        string
		Berita2        string
		Berita3        string
		Berita4        string
		Berita5        string
		FlagProses     FlagProses
	}
	tests := []struct {
		name           string
		fields         fields
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "failed va must 16 character",
			fields: fields{
				VirtualAccount: "22",
			},
			wantErr:        true,
			wantErrMessage: "length virtual account number not 16",
		},
		{
			name: "failed name can't empty",
			fields: fields{
				VirtualAccount: validVa,
			},
			wantErr:        true,
			wantErrMessage: "name cannot empty",
		},
		{
			name: "failed tanggal exp can't empty",
			fields: fields{
				VirtualAccount: validVa,
				Nama:           "test nama",
			},
			wantErr:        true,
			wantErrMessage: "tanggal exp cannot empty",
		},
		{
			name: "failed tanggal exp wrong format",
			fields: fields{
				VirtualAccount: validVa,
				Nama:           "test nama",
				TanggalExp:     "2023-01-31T19:53:03.82888+07:00",
			},
			wantErr:        true,
			wantErrMessage: "tanggal exp wrong format",
		},
		{
			name: "failed unknown flag proses",
			fields: fields{
				VirtualAccount: validVa,
				Nama:           "test nama",
				TanggalExp:     validTglExp,
				FlagProses:     5,
			},
			wantErr:        true,
			wantErrMessage: "flag process undefined",
		},
		{
			name: "success",
			fields: fields{
				VirtualAccount: validVa,
				Nama:           "test nama",
				TanggalExp:     validTglExp,
				FlagProses:     ProsesCreate,
			},
			wantErr:        false,
			wantErrMessage: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := RequestRegistration{
				VirtualAccount: tt.fields.VirtualAccount,
				Nama:           tt.fields.Nama,
				TotalTagihan:   tt.fields.TotalTagihan,
				TanggalExp:     tt.fields.TanggalExp,
				Berita1:        tt.fields.Berita1,
				Berita2:        tt.fields.Berita2,
				Berita3:        tt.fields.Berita3,
				Berita4:        tt.fields.Berita4,
				Berita5:        tt.fields.Berita5,
				FlagProses:     tt.fields.FlagProses,
			}
			err := data.Verify()
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyRegistration() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.wantErrMessage {
				t.Errorf("VerifyRegistration() error = %v, wantErrMessage %v", err, tt.wantErrMessage)
			}
		})
	}
}

func TestRequestPosting_Verify(t *testing.T) {
	type fields struct {
		VirtualAccount string
		Amount         uint64
		Reference      string
		Tanggal        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := RequestPosting{
				VirtualAccount: tt.fields.VirtualAccount,
				Amount:         tt.fields.Amount,
				Reference:      tt.fields.Reference,
				Tanggal:        tt.fields.Tanggal,
			}
			if err := data.Verify(); (err != nil) != tt.wantErr {
				t.Errorf("RequestPosting.VerifyPosting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
