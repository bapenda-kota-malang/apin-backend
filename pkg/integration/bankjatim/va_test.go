package bankjatim

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/bapenda-kota-malang/apin-backend/pkg/integration"
)

func Test_virtualAccount_Registration(t *testing.T) {
	client := integration.NewClient()
	url := "https://fca830b2-dd1f-4b3b-99ac-f8d63ae38dbf.mock.pstmn.io"
	va := "0123456789012345"
	nama := "XXXXXX XXXXX XXXXXX XXXXXX XXXXXXX XXXXXXX XXXXX"

	type fields struct {
		url    string
		client *http.Client
	}
	type args struct {
		ctx     context.Context
		payload RequestRegistration
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData ResponseRegistration
		wantErr  bool
	}{
		{
			name: "check data",
			fields: fields{
				url:    url,
				client: client,
			},
			args: args{
				ctx: context.Background(),
				payload: RequestRegistration{
					VirtualAccount: va,
					Nama:           nama,
					TotalTagihan:   1234567890123,
					TanggalExp:     "20181231",
					Berita1:        "INFO 1",
					Berita2:        "INFO 2",
					Berita3:        "INFO 3",
					Berita4:        "INFO 4",
					Berita5:        "INFO 5",
					FlagProses:     ProsesCreate,
				},
			},
			wantData: ResponseRegistration{
				VirtualAccount: va,
				Nama:           nama,
				Status: ResponseStatus{
					IsError:      "False",
					ResponseCode: "00",
					ErrorDesc:    "Success",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			va := &VirtualAccount{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			gotData, err := va.Registration(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("virtualAccount.Registration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("virtualAccount.Posting() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
