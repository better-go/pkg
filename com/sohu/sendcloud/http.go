package sendcloud

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/better-go/pkg/log"
)

// post form:
func QuickPost(cli *http.Client, url string, data url.Values) error {
	// post:
	resp, err := cli.PostForm(url, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// resp:
	body, err := ioutil.ReadAll(resp.Body)
	log.Debugf("sendcould post resp: %+v, err: %v", string(body), err)
	if err != nil {
		return err
	}

	// ok:
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return errors.New(string(body))

}
