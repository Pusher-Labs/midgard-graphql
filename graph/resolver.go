package graph

import (
	"net/http"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var DefaultHttpClient = &http.Client{
	Timeout: time.Second * 10,
}

const BaseUrl = "http://18.158.69.134:8080/v1"

