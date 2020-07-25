package repositories

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"go-industry-rest-microservices/src/api/clients/restclient"
	"go-industry-rest-microservices/src/api/domain/repositories"
	"go-industry-rest-microservices/src/api/utils/errors"
	"go-industry-rest-microservices/src/api/utils/test_utils"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJSONBody(t *testing.T) {
	// response := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(response)
	// request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	// c.Request = request

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiError, err := errors.NewAPIErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, http.StatusBadRequest, apiError.GetStatusCode())
	assert.EqualValues(t, "Invalid JSON body", apiError.GetMessage())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	// response := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(response)
	// request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	// c.Request = request

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.github.com/v3"}`)),
		},
	})

	CreateRepo(c)

	assert.EqualValues(t, http.StatusUnauthorized, response.Code)

	apiError, err := errors.NewAPIErrorFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, http.StatusUnauthorized, apiError.GetStatusCode())
	assert.EqualValues(t, "Requires authentication", apiError.GetMessage())
}

func TestCreateRepoNoError(t *testing.T) {
	// response := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(response)
	// request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	// c.Request = request

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 1296269, "name": "testing", "owner": {"login": "gvu0110"}}`)),
		},
	})

	CreateRepo(c)
	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 1296269, result.ID)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "gvu0110", result.Owner)
}
