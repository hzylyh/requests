/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/7/25 11:43
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/7/25 11:43
 */

package requests

import (
	"io"
	"net/http"
)

var (
	// client is a shared http Client.
	client HttpClient = &http.Client{}
)

// HttpClient interface has the method required to use a type as custom http client.
// The net/*http.Client type satisfies this interface.
type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

func Get(url string, param map[string]string) (respBytes []byte, err error) {
	//url := "http://192.168.1.201:1273/api/helm/repositories/%5Blocal%5D"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	respBytes, err = io.ReadAll(resp.Body)
	return
}

func Post(url string) {

}
