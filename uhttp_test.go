package uhttp

import (
	"net/url"
	"testing"

	"github.com/michelia/ulog"
)

func TestPostForm(t *testing.T) {
	slog := ulog.NewConsole()
	i := struct {
		Code  int    `json:"code"`
		Err   string `json:"err"`
		Logid string `json:"logid"`
	}{}
	data := url.Values{}
	data.Add("park_code", "1234")
	// data.Add("protocol", "1234")
	err := PostForm(slog, "http://116.62.132.145:9630/mul_vpl_clean", data, &i)
	if err != nil {
		slog.Error().Caller().Err(err).Msg("")
	}
}
