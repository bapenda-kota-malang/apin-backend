package bankjatim

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type VirtualAccount struct {
	url    string
	client *http.Client
}

func NewVirtualAccount(url string, client *http.Client) *VirtualAccount {
	return &VirtualAccount{url, client}
}

func (va *VirtualAccount) Registration(ctx context.Context, payload RequestRegistration) (data *ResponseRegistration, err error) {
	if err = payload.Verify(); err != nil {
		err = fmt.Errorf("failed verify registration payload: %w", err)
		return
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		err = fmt.Errorf("failed create registration payload: %w", err)
		return
	}
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/registration", va.url), bytes.NewBuffer(jsonBody))
	if err != nil {
		err = fmt.Errorf("failed create request registration virtual account: %w", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := va.client.Do(req)
	if err != nil {
		err = fmt.Errorf("error api request registration virtual account: %w", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http response code error: got: %d, want: %d", resp.StatusCode, http.StatusOK)
		return
	}

	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		err = fmt.Errorf("failed decode response: %w", err)
		return
	}

	switch data.Status.ResponseCode {
	case CodeTidakDitemukan:
		data.Status.ErrorDesc = DescTidakDitemukan
	case CodeLunas:
		data.Status.ErrorDesc = DescLunas
	case CodeTidakSesuai:
		data.Status.ErrorDesc = DescTidakSesuai
	case CodeError:
		data.Status.ErrorDesc = DescError + data.Status.ErrorDesc
	}
	return
}
