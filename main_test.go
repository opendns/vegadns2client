package vegadns2client

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ipPort = "127.0.0.1:2112"

func TestGetDomainID(t *testing.T) {
	ts := startTestServer(
		[]byte(`{"domains":[{"domain_id" :1,"domain":"example.com","status":"active","owner_id":0}]}`),
		"application/json",
	)
	defer ts.Close()

	v := initMockClient()
	v.User = "user@example.com"
	v.Pass = "secret"
	domainID, err := v.GetDomainID("example.com")

	assert.Equal(t, 1, domainID)
	assert.Nil(t, err)
}

// Starts and returns a test server using a custom ip/port. Defer close() afterwards.
func startTestServer(response []byte, contentType string) *httptest.Server {
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Write(response)
	}))
	l, _ := net.Listen("tcp", ipPort)
	ts.Listener = l
	ts.Start()
	return ts
}

func initMockClient() VegaDNSClient {
	return NewVegaDNSClient("http://" + ipPort)
}
