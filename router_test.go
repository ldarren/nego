package nego

import (
	"net/http"
	"testing"
)

func TestHTTPHandlerConforms (t *testing.T) {
	// Make sure the Router conforms with the http.Handler interface
	var _ http.Handler = New()
}
