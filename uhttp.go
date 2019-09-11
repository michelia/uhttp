package uhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/michelia/ulog"
)

// SetTimeout
// timeout 单位是秒 默认是 30秒 如改为 10
// maxIdleConns  连接池的最大连接数 默认是100 如改为 20
func SetClient(timeout, maxIdleConns int) {
	http.DefaultClient.Timeout = time.Second * time.Duration(timeout)
	t := http.DefaultTransport.(*http.Transport)
	t.MaxIdleConns = maxIdleConns
	http.DefaultTransport = t
}

// Get v must is pointer
// and Response must is json body
func Get(slog ulog.Logger, url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Get: resp err")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("resp.Status not is 200")
		slog.Error().Caller().Err(err).Msg("Get: ioutil.ReadAll err")
		return err
	}
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Get: ioutil.ReadAll err")
		return err
	}
	slog.Debug().RawJSON("body", rb).Msg("Get: resp_body")
	err = json.Unmarshal(rb, v)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Get: json.Unmarshal")
		return err
	}
	return nil
}

// GetAuth v must is pointer
// and Response must is json body
func GetAuth(slog ulog.Logger, url, username, password string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Get: req err")
		return err
	}
	req.SetBasicAuth(username, password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Get: resp err")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("resp.Status not is 200")
		slog.Error().Caller().Err(err).Msg("Get: ioutil.ReadAll err")
		return err
	}
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Get: ioutil.ReadAll err")
		return err
	}
	slog.Debug().RawJSON("body", rb).Msg("Get: resp_body")
	err = json.Unmarshal(rb, v)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Get: json.Unmarshal")
		return err
	}
	return nil
}

// Post v must is pointer
// and Request must is json body
// and Response must is json body
func Post(slog ulog.Logger, url string, body []byte, v interface{}) error {
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Post: resp err")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("resp.Status not is 200")
		slog.Error().Caller().Err(err).Msg("Post: ioutil.ReadAll err")
		return err
	}
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Post: ioutil.ReadAll err")
		return err
	}
	slog.Debug().RawJSON("body", rb).Msg("Post: resp_body")
	err = json.Unmarshal(rb, v)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Post: json.Unmarshal")
		return err
	}
	return nil
}

// PostAuth v must is pointer
// and Request must is json body
// and Response must is json body
func PostAuth(slog ulog.Logger, url, username, password string, body []byte, v interface{}) error {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Post: req err")
		return err
	}
	req.SetBasicAuth(username, password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Post: resp err")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("resp.Status not is 200")
		slog.Error().Caller().Err(err).Msg("Post: ioutil.ReadAll err")
		return err
	}
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Post: ioutil.ReadAll err")
		return err
	}
	slog.Debug().RawJSON("body", rb).Msg("Post: resp_body")
	err = json.Unmarshal(rb, v)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("Post: json.Unmarshal")
		return err
	}
	return nil
}

// PostForm v is must pointer
// data is url.Values
// and Response is must json body
func PostForm(slog ulog.Logger, url string, data url.Values, v interface{}) error {
	resp, err := http.PostForm(url, data)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("PostForm: resp err")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("resp.Status not is 200")
		slog.Error().Caller().Err(err).Msg("PostForm: ioutil.ReadAll err")
		return err
	}
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("PostForm: ioutil.ReadAll err")
		return err
	}
	slog.Debug().RawJSON("body", rb).Msg("PostForm: resp_body")
	err = json.Unmarshal(rb, v)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("PostForm: json.Unmarshal")
		return err
	}
	return nil
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// GinBodyLogJsonMiddleware print body log
func GinBodyLogJsonMiddleware(slog ulog.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		action := c.Request.URL.Path
		if c.Request.Method == "POST" {
			b, err := c.GetRawData()
			if err != nil {
				slog.Error().Caller().Err(err).Msg("c.GetRawData()")
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))
			slog.Debug().Str("action", action).RawJSON("body", b).Msg("gin_request_body")
		}
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		if c.Writer.Status() == 200 {
			slog.Debug().Str("action", action).RawJSON("body", blw.body.Bytes()).Msg("gin_response_body")
		}
	}
}
