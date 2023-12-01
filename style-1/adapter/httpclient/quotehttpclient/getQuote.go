package quotehttpclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/desafio-client-server-api-go-expert-pos/style-1/core/dto"
)

const (
	BASEURL = "https://economia.awesomeapi.com.br/json/last"
)

func (httpclient httpclient) GetQuote(ctx context.Context, code string) (*dto.QuoteClientResponse, error) {
	ctxRequest, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	response := map[string]any{}
	req, err := http.NewRequestWithContext(ctxRequest, http.MethodGet, fmt.Sprintf("%s/%s", BASEURL, code), nil)

	if err != nil {
		return nil, err
	}

	res, err := httpclient.client.Do(req)

	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("Context Error:", err)
	}

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	quoteMapped, err := dto.Mapper(response[strings.Replace(code, "-", "", -1)], dto.QuoteClientResponse{})
	if err != nil {
		return nil, err
	}

	return quoteMapped, nil
}
