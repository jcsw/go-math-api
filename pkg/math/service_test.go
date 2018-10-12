package math

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jcsw/go-math-api/pkg/entity"
)

func TestShouldReturnResultWhenOperationIsSum(t *testing.T) {
	request := entity.MathOperationRequest{Operation: "SUM", Parameters: []float64{1, 1}}

	repositoryMock := &RepositoryMock{}
	repositoryMock.On("PersistOperationLog", request, mock.Anything).Return(nil)

	service := NewService(repositoryMock)
	response, err := service.ExecuteMathOperation(request)

	assert.Nil(t, err)

	if assert.NotNil(t, response) {
		assert.Equal(t, response.Result, 2.0)
	}

	repositoryMock.AssertCalled(t, "PersistOperationLog", request, mock.Anything)
}

func TestShouldReturnNotImplementedWhenOperationIsSub(t *testing.T) {
	request := entity.MathOperationRequest{Operation: "SUB", Parameters: []float64{1, 1}}

	repositoryMock := &RepositoryMock{}
	repositoryMock.On("PersistOperationLog", request, mock.Anything).Return(nil)

	service := NewService(repositoryMock)
	response, err := service.ExecuteMathOperation(request)

	assert.Nil(t, response)

	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), "not implemented")
	}

	repositoryMock.AssertNotCalled(t, "PersistOperationLog", request, mock.Anything)
}

func TestShouldReturnErrorWhenOperationIsInvalid(t *testing.T) {
	request := entity.MathOperationRequest{Operation: "ABC", Parameters: []float64{1, 1}}

	repositoryMock := &RepositoryMock{}
	repositoryMock.On("PersistOperationLog", request, mock.Anything).Return(nil)

	service := NewService(repositoryMock)
	response, err := service.ExecuteMathOperation(request)

	assert.Nil(t, response)

	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), "invalid operation")
	}

	repositoryMock.AssertNotCalled(t, "PersistOperationLog", request, mock.Anything)
}

func TestShouldReturnErrorWhenRepositoryReturnError(t *testing.T) {
	request := entity.MathOperationRequest{Operation: "SUM", Parameters: []float64{1, 1}}

	repositoryMock := &RepositoryMock{}
	repositoryMock.On("PersistOperationLog", request, mock.Anything).Return(errors.New("mock error"))

	service := NewService(repositoryMock)
	response, err := service.ExecuteMathOperation(request)

	assert.Nil(t, response)

	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), "could not persist operation")
	}

	repositoryMock.AssertCalled(t, "PersistOperationLog", request, mock.Anything)
}
