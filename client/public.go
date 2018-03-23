// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// unnamed API: public Resource Client
//
// Command:
// $ goagen
// --design=github.com/Microkubes/authorization-server/design
// --out=$(GOPATH)/src/github.com/Microkubes/authorization-server
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
)

// DownloadCSS downloads /files with the given filename and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadCSS(ctx context.Context, filename, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	p := path.Join("/auth/css/", filename)
	u := url.URL{Host: c.Host, Scheme: scheme, Path: p}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}

// DownloadJs downloads /files with the given filename and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadJs(ctx context.Context, filename, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	p := path.Join("/auth/js/", filename)
	u := url.URL{Host: c.Host, Scheme: scheme, Path: p}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}
