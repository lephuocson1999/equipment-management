package httpfx

import (
	"net/http"
	"time"

	"go.uber.org/fx"
)

var Module = fx.Provide(provideHttpClient)

func provideHttpClient() http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := http.Client{Transport: tr}
	return client
}
