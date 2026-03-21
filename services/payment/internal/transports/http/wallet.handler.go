package http

import (
	"fmt"
	"hash/fnv"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/payOSHQ/payos-lib-golang/v2"

	"github.com/course-sphere/course-sphere-backend/shared/transports/http/middleware"
)

var (
	cache map[int64]uuid.UUID = make(map[int64]uuid.UUID)
)

func (s *Server) GetWalletByUser(c fuego.ContextNoBody) (*Wallet, error) {
	ctx := c.Context()

	token := c.Value(middleware.TokenKey).(jwt.Token)
	sub, _ := token.Subject()
	userID, err := uuid.Parse(sub)
	if err != nil {
		return nil, fuego.UnauthorizedError{
			Err:    err,
			Detail: "Invalid token",
		}
	}

	raw, err := s.Wallet.GetByUser(ctx, userID)
	if err != nil {
		return nil, fuego.InternalServerError{
			Err: err,
		}
	}
	var wallet Wallet
	copier.Copy(&wallet, raw)

	return &wallet, nil
}

func (s *Server) GetWalletHistories(c fuego.ContextNoBody) ([]History, error) {
	ctx := c.Context()

	token := c.Value(middleware.TokenKey).(jwt.Token)
	sub, _ := token.Subject()
	userID, err := uuid.Parse(sub)
	if err != nil {
		return nil, fuego.UnauthorizedError{
			Err:    err,
			Detail: "Invalid token",
		}
	}

	wallet, err := s.Wallet.GetByUser(ctx, userID)
	if err != nil {
		return nil, fuego.InternalServerError{
			Err: err,
		}
	}

	raw, err := s.Wallet.GetHistories(ctx, wallet.ID)
	if err != nil {
		return nil, err
	}

	var histories []History
	copier.Copy(&histories, raw)

	return histories, nil
}

func (s *Server) CreatePaymentLink(c fuego.ContextWithBody[CreatePaymentLinkData]) (string, error) {
	ctx := c.Context()

	token := c.Value(middleware.TokenKey).(jwt.Token)
	sub, _ := token.Subject()
	userID, err := uuid.Parse(sub)
	if err != nil {
		return "", fuego.UnauthorizedError{
			Err:    err,
			Detail: "Invalid token",
		}
	}

	body, err := c.Body()
	if err != nil {
		return "", err
	}

	wallet, err := s.Wallet.GetByUser(ctx, userID)
	if err != nil {
		return "", fuego.InternalServerError{
			Err: err,
		}
	}

	// TODO: find better way to gen key
	h := fnv.New64a()
	h.Write([]byte(userID.String()))
	fmt.Fprintf(h, "%d", time.Now().UnixNano())
	orderCode := int64(h.Sum64())
	if orderCode < 0 {
		orderCode = -orderCode
	}
	orderCode %= (int64(1) << 32)

	cache[orderCode] = wallet.ID

	paymentRequest := payos.CreatePaymentLinkRequest{
		OrderCode:   orderCode,
		Amount:      int(body.Amount),
		Description: "Payment",
		ReturnUrl:   fmt.Sprintf("%s/payment/callback", s.Config.ApiURL),
	}
	payment, err := s.PayOSClient.PaymentRequests.Create(ctx, paymentRequest)
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "Failed to create payment link",
		}
	}

	return payment.CheckoutUrl, nil
}

func (s *Server) PaymentCallback(c fuego.Context[PaymentCallbackData, PaymentCallbackStatus]) (any, error) {
	ctx := c.Context()

	status, err := c.Params()
	if err != nil {
		return nil, err
	}
	if !status.Success {
		return nil, fuego.BadRequestError{
			Detail: "Failed to process payment",
		}
	}

	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	_, err = s.PayOSClient.Webhooks.VerifyData(ctx, body)
	if err != nil {
		return nil, err
	}

	walletID := cache[body.OrderCode]

	err = s.Wallet.Deposit(ctx, walletID, body.Amount, body.Description)

	return nil, err
}

func (s *Server) Withdraw(c fuego.ContextWithBody[WithdrawData]) (any, error) {

	ctx := c.Context()

	token := c.Value(middleware.TokenKey).(jwt.Token)
	sub, _ := token.Subject()
	userID, err := uuid.Parse(sub)
	if err != nil {
		return nil, fuego.UnauthorizedError{
			Err:    err,
			Detail: "Invalid token",
		}
	}

	wallet, err := s.Wallet.GetByUser(ctx, userID)
	if err != nil {
		return nil, fuego.InternalServerError{
			Err: err,
		}
	}

	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	err = s.Wallet.Withdraw(ctx, wallet.ID, body.Amount, body.Description)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Not enough balance",
		}
	}

	return nil, nil
}
