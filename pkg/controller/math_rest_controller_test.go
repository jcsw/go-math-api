package controller_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jcsw/go-math-api/pkg/controller"
	"github.com/jcsw/go-math-api/pkg/math"
)

func TestPostCustomerHandler(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		description        string
		repositoryMock     *math.RepositoryMock
		method             string
		url                string
		payload            []byte
		expectedStatusCode int
		expectedBody       string
	}{
		{
			description:        "should return error 400 when body is not valid",
			repositoryMock:     mockPersistOperationLogSuccesfull(),
			method:             "POST",
			url:                "/math/operation",
			payload:            []byte(`"a=b"`),
			expectedStatusCode: 400,
			expectedBody:       `{"error":"invalid request payload"}`,
		},
		{
			description:        "should return error 400 when is missing an argument",
			repositoryMock:     mockPersistOperationLogSuccesfull(),
			method:             "POST",
			url:                "/math/operation",
			payload:            []byte(`{"operation":"SUM"}`),
			expectedStatusCode: 400,
			expectedBody:       `{"error":"invalid value 'parameters'"}`,
		},
		{
			description:        "should return 200 when successful",
			repositoryMock:     mockPersistOperationLogSuccesfull(),
			method:             "POST",
			url:                "/math/operation",
			payload:            []byte(`{"operation":"SUM", "parameters":[1, 1]}`),
			expectedStatusCode: 200,
			expectedBody:       `{"result":[0-9\.]*}`,
		},
		{
			description:        "should return 500 when occurs internal error",
			repositoryMock:     mockPersistOperationLogError(),
			method:             "POST",
			url:                "/math/operation",
			payload:            []byte(`{"operation":"SUM", "parameters":[1, 1]}`),
			expectedStatusCode: 500,
			expectedBody:       `{"error":"could not complete operation"}`,
		},
	}

	for _, tc := range tests {

		req, err := http.NewRequest(tc.method, tc.url, bytes.NewBuffer(tc.payload))
		assert.NoError(err)

		resp := httptest.NewRecorder()

		service := math.NewService(tc.repositoryMock)
		mathHandler := controller.MathHandler{MathService: service}
		mathHandler.Register(resp, req)

		assert.Equal(tc.expectedStatusCode, resp.Code, tc.description)
		assert.Regexp(tc.expectedBody, string(resp.Body.Bytes()), tc.description)
	}
}

func mockPersistOperationLogSuccesfull() *math.RepositoryMock {
	repositoryMock := &math.RepositoryMock{}
	repositoryMock.On("PersistOperationLog", mock.Anything, mock.Anything).Return(nil)
	return repositoryMock
}

func mockPersistOperationLogError() *math.RepositoryMock {
	repositoryMock := &math.RepositoryMock{}
	repositoryMock.On("PersistOperationLog", mock.Anything, mock.Anything).Return(errors.New("mock error"))
	return repositoryMock
}
