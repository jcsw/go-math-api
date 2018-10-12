package math

import "github.com/jcsw/go-math-api/pkg/entity"

// Repository define the data math repository
type Repository interface {
	PersistOperationLog(request entity.MathOperationRequest, response entity.MathOperationResponse) error
}

//UseCase use case interface
type UseCase interface {
	ExecuteMathOperation(request entity.MathOperationRequest) (entity.MathOperationResponse, error)
}
