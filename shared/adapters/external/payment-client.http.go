package external

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/course-sphere/course-sphere-backend/shared/domain"
	"github.com/course-sphere/course-sphere-backend/shared/ports"
)

type HTTPPaymentClient struct {
	ProxyURL string
}

var _ ports.PaymentClient = &HTTPPaymentClient{}

func (c *HTTPPaymentClient) GetWalletByUser(ctx context.Context, token string) (*domain.Wallet, error) {
	endpoint := fmt.Sprintf("%s/wallet", c.ProxyURL)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var wallet domain.Wallet
	err = json.NewDecoder(resp.Body).Decode(&wallet)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (c *HTTPPaymentClient) Withdraw(ctx context.Context, token string, amount int64, detail string) error {
	endpoint := fmt.Sprintf("%s/payment/withdraw", c.ProxyURL)

	data := map[string]any{
		"amount": amount,
		"detail": detail,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", token)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Not enough money")
	}

	return nil
}
