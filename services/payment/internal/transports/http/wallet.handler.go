package http

import (
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/lestrrat-go/jwx/v3/jwt"

	"github.com/course-sphere/course-sphere-backend/shared/transports/http/middleware"
)

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
