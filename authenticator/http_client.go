package authenticator

import (
	"net/http"
	"time"
)

var httpClient = http.Client{
	Timeout: time.Second * 5,
}
