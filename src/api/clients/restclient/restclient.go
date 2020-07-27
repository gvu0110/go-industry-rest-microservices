package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	enableMocks = false
	mocks       = make(map[string]*Mock)
)

type Mock struct {
	URL        string
	HTTPMethod string
	BodyText   string
	Response   *http.Response
	Err        error
}

func (m *Mock) GetResponse() *http.Response {
	m.Response.Body = ioutil.NopCloser(strings.NewReader(m.BodyText))
	return m.Response
}

func getMockID(HTTPMethod, URL string) string {
	return fmt.Sprintf("%s_%s", HTTPMethod, URL)
}

func StartMockups() {
	enableMocks = true
}

func StopMockups() {
	enableMocks = false
}

func FlushMockups() {
	mocks = make(map[string]*Mock)
}

func AddMockups(mock Mock) {
	mocks[getMockID(mock.HTTPMethod, mock.URL)] = &mock
}

func Post(URL string, body interface{}, headers http.Header) (*http.Response, error) {
	if enableMocks {
		mock := mocks[getMockID(http.MethodPost, URL)]
		if mock == nil {
			return nil, errors.New("No mockup found for given request")
		}
		if mock.BodyText != "" {
			return mock.GetResponse(), mock.Err
		}
		return mock.Response, mock.Err
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, URL, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
