package graph

import (
	"net/http"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	httpClient	*http.Client
	baseUrl		string
}

func NewResolver() (*Resolver, error) {

	r := &Resolver{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		baseUrl: "http://18.158.69.134:8080/v1",
	}

	return r, nil
}


