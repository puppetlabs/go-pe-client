package puppetdb

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func init() {
	pdbClient = NewClient(hostURL, "xxxx", nil, time.Second*1000)
	httpmock.Activate()
	httpmock.ActivateNonDefault(pdbClient.resty.GetClient())
}

func setupGetResponder(t *testing.T, url, query, responseFilename string) {
	httpmock.Reset()
	responseBody, err := os.ReadFile("testdata/" + responseFilename)
	require.Nil(t, err)
	response := httpmock.NewBytesResponse(200, responseBody)
	response.Header.Set("Content-Type", "application/json")
	if query != "" {
		httpmock.RegisterResponder(http.MethodGet, hostURL+url, httpmock.ResponderFromResponse(response))
	} else {
		httpmock.RegisterResponderWithQuery(http.MethodGet, hostURL+url, query, httpmock.ResponderFromResponse(response))
	}
	response.Body.Close()
}

type mockPaginatedGetOptions struct {
	limit             int
	total             int
	pageFilenames     []string
	returnErrorOnPage *int
}

func setupPaginatedGetResponder(t *testing.T, url, query string, opts mockPaginatedGetOptions) {
	var pages [][]byte

	for _, pfn := range opts.pageFilenames {
		responseBody, err := os.ReadFile(filepath.Join("testdata", pfn))
		require.NoError(t, err)

		pages = append(pages, responseBody)
	}

	responder := func(r *http.Request) (*http.Response, error) {
		var (
			offset  int
			pageNum int
			err     error
		)

		offsetS := r.URL.Query().Get("offset")
		if offsetS != "" {
			offset, err = strconv.Atoi(offsetS)
			if err != nil {
				return nil, err
			}
		}

		if offset > 0 {
			pageNum = offset / opts.limit
		}

		if opts.returnErrorOnPage != nil && *opts.returnErrorOnPage == pageNum {
			response := httpmock.NewBytesResponse(http.StatusInternalServerError, []byte("{\"error\": \"oops\"}"))
			response.Header.Set("Content-Type", "application/json")

			defer response.Body.Close()

			return response, nil
		}

		responseBody := pages[pageNum]

		response := httpmock.NewBytesResponse(http.StatusOK, responseBody)
		response.Header.Set("Content-Type", "application/json")
		response.Header.Set("X-Records", fmt.Sprintf("%d", opts.total))

		defer response.Body.Close()

		return response, nil
	}

	httpmock.Reset()
	httpmock.RegisterResponderWithQuery(http.MethodGet, hostURL+url, query, responder)
}

func setupURLErrorResponder(t *testing.T, url string) {
	setupURLResponderWithStatusCode(t, url, http.StatusNotFound)
}

func setupURLResponderWithStatusCode(t *testing.T, url string, statusCode int) {
	setupResponderWithStatusCodeAndBody(t, url, statusCode, expectedURLError)
}

func setupResponderWithStatusCodeAndBody(t *testing.T, url string, statusCode int, response interface{}) {
	httpmock.Reset()
	responder, err := httpmock.NewJsonResponder(statusCode, response)
	require.Nil(t, err)
	httpmock.RegisterResponder(http.MethodGet, hostURL+url, responder)
}

var pdbClient *Client

var hostURL = "https://test-host:8081"

var expectedURLError = url.Error{Op: "nil", URL: hostURL, Err: nil}
