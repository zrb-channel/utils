package utils

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
)

func NetworkImageToBase64(ctx context.Context, addr string) (string, error) {
	resp, err := Request(ctx).Get(addr)
	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		return "", errors.New(resp.Status())
	}

	return base64.StdEncoding.EncodeToString(resp.Body()), nil
}
