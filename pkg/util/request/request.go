package request

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// type RequestOptionsFunc func(*RequestOptions)

// type RequestOptions struct {
// 	retryTimes int
// 	userAgent  string
// 	refer      string
// }

// func WithRetryTimes(v int) RequestOptionsFunc {
// 	return func(o *RequestOptions) {
// 		o.retryTimes = v
// 	}
// }

// func WithUserAgent(v string) RequestOptionsFunc {
// 	return func(o *RequestOptions) {
// 		o.userAgent = v
// 	}
// }

// func WithRefer(v string) RequestOptionsFunc {
// 	return func(o *RequestOptions) {
// 		o.refer = v
// 	}
// }

func Request(method, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		DisableCompression:  true,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Minute,
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if _, ok := headers["Referer"]; !ok {
		req.Header.Set("Referer", url)
	}

	var (
		res          *http.Response
		requestError error
	)

	for i := 0; ; i++ {
		res, requestError = client.Do(req)
		if requestError == nil && res.StatusCode < 400 {
			break
		} else if i+1 >= 3 {
			var err error
			if requestError != nil {
				err = errors.Errorf("request error: %v", requestError)
			} else {
				err = errors.Errorf("%s request error: HTTP %d", url, res.StatusCode)
			}
			return nil, errors.WithStack(err)
		}
		time.Sleep(1 * time.Second)
	}

	return res, nil
}

func Get(rawURL string, params map[string]string, headers map[string]string) (string, error) {
	urlValues := url.Values{}
	Url, _ := url.Parse(rawURL)
	for key, val := range params {
		urlValues.Set(key, val)
	}
	Url.RawQuery = urlValues.Encode()
	urlPath := Url.String()

	body, err := GetByte(http.MethodGet, urlPath, nil, headers)
	return string(body), err
}

func PostForm(url string, data map[string]string, headers map[string]string) (string, error) {
	return Post(url, "application/x-www-form-urlencoded", data, headers)
}

func PostJson(url string, data map[string]string, headers map[string]string) (string, error) {
	return Post(url, "application/json", data, headers)
}

// func PostFile(url string, data map[string]string) (string, error) {
// 	return Post(url, "multipart/form-data", data, nil)
// }

func Post(url, contentType string, data map[string]string, headers map[string]string) (string, error) {
	body := GetReader(data, contentType)
	headers["Content-Type"] = contentType
	resp, err := GetByte(http.MethodPost, url, body, headers)
	return string(resp), err
}

func GetReader(data map[string]string, contentType string) io.Reader {
	if strings.Contains(contentType, "json") {
		bytesData, _ := json.Marshal(data)
		return bytes.NewReader(bytesData)
	}

	urlValues := url.Values{}
	for key, val := range data {
		urlValues.Set(key, val)
	}
	body := urlValues.Encode()
	return strings.NewReader(body)
}

func GetByte(method, url string, body io.Reader, headers map[string]string) ([]byte, error) {
	res, err := Request(method, url, body, headers)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer res.Body.Close()

	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(res.Body)
	case "deflate":
		reader = flate.NewReader(res.Body)
	default:
		reader = res.Body
	}
	defer reader.Close()

	response, err := ioutil.ReadAll(reader)
	if err != nil && err != io.EOF {
		return nil, errors.WithStack(err)
	}

	return response, nil
}

// func Headers(url string, headers map[string]string) (http.Header, error) {
// 	res, err := Request(http.MethodGet, url, nil, headers)
// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}
// 	defer res.Body.Close()

// 	return res.Header, nil
// }

// func Size(url string, headers map[string]string) (int64, error) {
// 	h, err := Headers(url, headers)
// 	if err != nil {
// 		return 0, err
// 	}
// 	s := h.Get("Content-Length")
// 	if s == "" {
// 		return 0, errors.New("Content-Length is not present")
// 	}
// 	size, err := strconv.ParseInt(s, 10, 64)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return size, nil
// }

// func ContentType(url string, headers map[string]string) (string, error) {
// 	h, err := Headers(url, headers)
// 	if err != nil {
// 		return "", err
// 	}
// 	s := h.Get("Content-Type")
// 	return strings.Split(s, ";")[0], nil
// }
