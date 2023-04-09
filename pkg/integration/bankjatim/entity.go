package bankjatim

import (
	"fmt"
	"time"
)

type RequestRegistration struct {
	VirtualAccount string     `json:"VirtualAccount"`
	Nama           string     `json:"Nama"`
	TotalTagihan   uint64     `json:"TotalTagihan"`
	TanggalExp     string     `json:"TanggalExp"`
	Berita1        string     `json:"Berita1"`
	Berita2        string     `json:"Berita2"`
	Berita3        string     `json:"Berita3"`
	Berita4        string     `json:"Berita4"`
	Berita5        string     `json:"Berita5"`
	FlagProses     FlagProses `json:"FlagProses"`
}

func (data RequestRegistration) Verify() error {
	if len(data.VirtualAccount) != 16 {
		return fmt.Errorf("length virtual account number not 16")
	}
	if data.Nama == "" {
		return fmt.Errorf("name cannot empty")
	}
	if data.TanggalExp == "" {
		return fmt.Errorf("tanggal exp cannot empty")
	}
	if _, err := time.Parse("20060102", data.TanggalExp); err != nil {
		return fmt.Errorf("tanggal exp wrong format")
	}
	switch data.FlagProses {
	case ProsesCreate, ProsesUpdate, ProsesDelete:
		// do nothing
	default:
		return fmt.Errorf("flag process undefined")
	}

	return nil
}

type RequestPosting struct {
	VirtualAccount string `json:"VirtualAccount"`
	Amount         uint64 `json:"Amount"`
	Reference      string `json:"Reference"`
	Tanggal        string `json:"Tanggal"`
}

func (data RequestPosting) Verify() error {
	if len(data.VirtualAccount) != 16 {
		return fmt.Errorf("length virtual account number not 16")
	}
	if data.Reference == "" {
		return fmt.Errorf("reference cannot empty")
	}
	if data.Tanggal == "" {
		return fmt.Errorf("tanggal cannot empty")
	}
	if _, err := time.Parse("2006-01-02 15:04:05", data.Tanggal); err != nil {
		return fmt.Errorf("tanggal wrong format")
	}
	return nil
}

type ResponseStatus struct {
	IsError      string     `json:"IsError"`
	ResponseCode StatusCode `json:"ResponseCode"`
	ErrorDesc    string     `json:"ErrorDesc"`
}

type ResponseRegistration struct {
	VirtualAccount string         `json:"VirtualAccount"`
	Nama           string         `json:"Nama"`
	Status         ResponseStatus `json:"Status"`
}

type ResponsePosting struct {
	VirtualAccount string         `json:"VirtualAccount"`
	Amount         uint64         `json:"Amount"`
	Tanggal        string         `json:"Tanggal"`
	Status         ResponseStatus `json:"Status"`
}
