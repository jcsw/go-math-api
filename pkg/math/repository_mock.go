package math

import (
	"github.com/jcsw/go-math-api/pkg/entity"
	"github.com/stretchr/testify/mock"
)

// RepositoryMock mock to Repository
type RepositoryMock struct {
	mock.Mock
}

// PersistOperationLog mock to PersistOperationLog
func (m *RepositoryMock) PersistOperationLog(request entity.MathOperationRequest, response entity.MathOperationResponse) error {
	args := m.Called(request, response)
	return args.Error(0)
}
