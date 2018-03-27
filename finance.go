package finance

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/piquette/finance-go/form"
)

// Printfer is an interface to be implemented by Logger.
type Printfer interface {
	Printf(format string, v ...interface{})
}

// init sets inital logger defaults.
func init() {
	Logger = log.New(os.Stderr, "", log.LstdFlags)
}

const (
	// YahooBackend is a constant representing the yahoo service backend.
	YahooBackend SupportedBackend = "yahoo"
	// YahooURL is the URL of the yahoo service backend.
	YahooURL string = "https://query2.finance.yahoo.com"
	// BATSBackend is a constant representing the uploads service backend.
	BATSBackend SupportedBackend = "bats"
	// BATSURL is the URL of the uploads service backend.
	BATSURL string = ""

	// Private constants.
	// ------------------

	defaultHTTPTimeout = 80 * time.Second
	yahooURL           = "https://query2.finance.yahoo.com"
	batsURL            = ""
)

var (
	// LogLevel is the logging level for this library.
	// 0: no logging
	// 1: errors only
	// 2: errors + informational (default)
	// 3: errors + informational + debug
	LogLevel = 2

	// Logger controls how this library performs logging at a package level. It is useful
	// to customise if you need it prefixed for your application to meet other
	// requirements
	Logger Printfer

	// Private vars.
	// -------------

	httpClient = &http.Client{Timeout: defaultHTTPTimeout}
	backends   Backends
)

// SupportedBackend is an enumeration of supported api endpoints.
type SupportedBackend string

// Backends are the currently supported endpoints.
type Backends struct {
	Yahoo, Bats Backend
	mu          sync.RWMutex
}

// BackendConfiguration is the internal implementation for making HTTP calls.
type BackendConfiguration struct {
	Type       SupportedBackend
	URL        string
	HTTPClient *http.Client
}

// Backend is an interface for making calls against an api service.
// This interface exists to enable mocking for during testing if needed.
type Backend interface {
	Call(method, path string, body *form.Values, params *Params, v interface{}) error
}

// SetHTTPClient overrides the default HTTP client.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func SetHTTPClient(client *http.Client) {
	httpClient = client
}

// NewBackends creates a new set of backends with the given HTTP client. You
// should only need to use this for testing purposes or on App Engine.
func NewBackends(httpClient *http.Client) *Backends {
	return &Backends{
		Yahoo: &BackendConfiguration{
			YahooBackend, YahooURL, httpClient,
		},
		Bats: &BackendConfiguration{
			BATSBackend, BATSURL, httpClient,
		},
	}
}

// GetBackend returns the currently used backend in the binding.
func GetBackend(backend SupportedBackend) Backend {
	switch backend {
	case YahooBackend:
		backends.mu.RLock()
		ret := backends.Yahoo
		backends.mu.RUnlock()
		if ret != nil {
			return ret
		}
		backends.mu.Lock()
		defer backends.mu.Unlock()
		backends.Yahoo = &BackendConfiguration{backend, yahooURL, httpClient}
		return backends.Yahoo
	case BATSBackend:
		backends.mu.RLock()
		ret := backends.Bats
		backends.mu.RUnlock()
		if ret != nil {
			return ret
		}
		backends.mu.Lock()
		defer backends.mu.Unlock()
		backends.Bats = &BackendConfiguration{backend, batsURL, httpClient}
		return backends.Bats
	}

	return nil
}

// SetBackend sets the backend used in the binding.
func SetBackend(backend SupportedBackend, b Backend) {
	switch backend {
	case YahooBackend:
		backends.Yahoo = b
	case BATSBackend:
		backends.Bats = b
	}
}

// Call is the Backend.Call implementation for invoking market data APIs.
func (s *BackendConfiguration) Call(method, path string, form *form.Values, params *Params, v interface{}) error {
	var body io.Reader
	if form != nil && !form.Empty() {
		data := form.Encode()
		if strings.ToUpper(method) == "GET" {
			path += "?" + data
		} else {
			body = bytes.NewBufferString(data)
		}
	}

	req, err := s.NewRequest(method, path, body, params)
	if err != nil {
		return err
	}

	if err := s.Do(req, v); err != nil {
		return err
	}

	return nil
}

// NewRequest is used by Call to generate an http.Request.
func (s *BackendConfiguration) NewRequest(method, path string, body io.Reader, params *Params) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.URL + path

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		if LogLevel > 0 {
			Logger.Printf("Cannot create api request: %v\n", err)
		}
		return nil, err
	}

	if params != nil {
		if params.Context != nil {
			req = req.WithContext(params.Context)
		}
	}

	return req, nil
}

// Do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (s *BackendConfiguration) Do(req *http.Request, v interface{}) error {
	if LogLevel > 1 {
		Logger.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
	}

	start := time.Now()

	res, err := s.HTTPClient.Do(req)

	if LogLevel > 2 {
		Logger.Printf("Completed in %v\n", time.Since(start))
	}

	if err != nil {
		if LogLevel > 0 {
			Logger.Printf("Request to api failed: %v\n", err)
		}
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if LogLevel > 0 {
			Logger.Printf("Cannot parse response: %v\n", err)
		}
		return err
	}

	if res.StatusCode >= 400 {
		return s.ResponseToError(res, resBody)
	}

	if LogLevel > 2 {
		Logger.Printf("API response: %q\n", resBody)
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}

// ResponseToError tbd.
func (s *BackendConfiguration) ResponseToError(res *http.Response, resBody []byte) error {

	var errMap map[string]interface{}

	err := json.Unmarshal(resBody, &errMap)
	if err != nil {
		return err
	}

	financeErr := &Error{}

	if LogLevel > 0 {
		Logger.Printf("Error encountered from api: %v\n", financeErr)
	}

	return financeErr
}
