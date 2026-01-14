package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type ConverterService struct {
	APIKey string
	Cache  sync.Map
	TTL    time.Duration
}

type rateResult struct {
	Rate     float64
	ExpireAt time.Time
}

type exchangeResponse struct {
	Result         string  `json:"result"`
	ConversionRate float64 `json:"conversion_rate"`
}

func NewConverterService(apiKey string) *ConverterService {
	return &ConverterService{
		APIKey: apiKey,
		TTL:    10 * time.Minute,
	}
}

func (s *ConverterService) Convert(from, to string, amount float64) (float64, float64, error) {
	key := fmt.Sprintf("%s->%s", from, to)

	if cached, ok := s.Cache.Load(key); ok {
		rate := cached.(rateResult)
		if time.Now().Before(rate.ExpireAt) {
			converted := amount * rate.Rate
			return converted, rate.Rate, nil
		}
	}

	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s", s.APIKey, from, to)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, errors.New("erro ao consultar API externa")
	}

	var data exchangeResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, 0, err
	}

	if data.Result != "success" {
		return 0, 0, errors.New("API externa retornou erro")
	}

	s.Cache.Store(key, rateResult{
		Rate:     data.ConversionRate,
		ExpireAt: time.Now().Add(s.TTL),
	})

	converted := amount * data.ConversionRate
	return converted, data.ConversionRate, nil
}
