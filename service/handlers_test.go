package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unrolled/render"
)

var formatter = render.New(render.Options{IndentJSON: true})

func TestGetAssets(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(http.HandlerFunc(createGetAssetsHandler(formatter)))
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL, nil)

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in GET to createGetAssetsHandler? %v", err)
	}

	defer res.Body.Close()

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %s", res.Status)
	}

	var response = []Asset{}
	if err := json.Unmarshal(payload, &response); err != nil {
		t.Errorf("Can't parse JSON: %s", payload)
	}

}
