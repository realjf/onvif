package networking

import (
	"bytes"
	"net/http"

	"github.com/juju/errors"
)

// SendSoap send soap message
func SendSoap(httpClient *http.Client, endpoint, message string) (*http.Response, error) {
	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return resp, errors.Annotate(err, "Post")
	}
	if resp.StatusCode == 400 {
		err = errors.New("Bad Request")
		return resp, errors.Annotate(err, "Bad Request")
	} else if resp.StatusCode == 401 {
		err = errors.New("Unauthorized")
		return resp, errors.Annotate(err, "Unauthorized")
	} else if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return resp, errors.Annotate(err, resp.Status)
	}

	return resp, nil
}
