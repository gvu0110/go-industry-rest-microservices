package polo

import (
	"go-industry-rest-microservices/src/api/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestPoloNoError(t *testing.T) {
	// response := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(response)
	// request, _ := http.NewRequest(http.MethodPost, "/macro", strings.NewReader(``))
	// c.Request = request

	request, _ := http.NewRequest(http.MethodPost, "/macro", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	Macro(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}
