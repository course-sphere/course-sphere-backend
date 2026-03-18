package external

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/course-sphere/course-sphere-backend/shared/ports"
	"github.com/google/uuid"
)

type HTTPGeneralClient struct {
	ProxyURL string
}

var _ ports.GeneralClient = &HTTPGeneralClient{}

func (c *HTTPGeneralClient) GetEnrolledCourses(ctx context.Context, token string) ([]uuid.UUID, error) {
	endpoint := fmt.Sprintf("%s/course/my", c.ProxyURL)

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

	var ids []uuid.UUID
	err = json.NewDecoder(resp.Body).Decode(&ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
