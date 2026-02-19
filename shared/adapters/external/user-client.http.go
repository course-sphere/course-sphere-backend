package external

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/shared/domain"
	"github.com/course-sphere/course-sphere-backend/shared/ports"
)

type HTTPUserClient struct {
	ProxyURL string
}

var _ ports.AuthClient = &HTTPAuthClient{}

func (c *HTTPUserClient) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	endpoint := fmt.Sprintf("%s/user/%s", c.ProxyURL, id)

	response, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var user domain.User
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
