package cache

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CacheTestSuite struct {
	suite.Suite
}

func (suite *CacheTestSuite) SetupTest() {
	suite.T().Log("This is when you can configure your test setup")
}

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("content-type", "application/json")
	}

	return httptest.NewServer(http.HandlerFunc(f))
}
func TestGetRequest_200(t *testing.T) {

	server := mockServer()
	defer server.Close()

	t.Log("Here is server info: ", server.URL)

	response, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("error in get request to server")
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		t.Fatal("should receive 200 status code")
	}

	t.Log("This is when your feature test done!")
}
