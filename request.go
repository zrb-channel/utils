package utils

import (
	"context"

	resty "github.com/go-resty/resty/v2"
)

var (
	httpClient = resty.New()
)

func Request(ctx context.Context) *resty.Request {
	return httpClient.R().SetContext(ctx)
}
