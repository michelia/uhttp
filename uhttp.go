package uhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/michelia/ulog"
)

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
