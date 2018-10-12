package math

import (
	"errors"

	"github.com/jcsw/go-math-api/pkg/driver/syslog"
	"github.com/jcsw/go-math-api/pkg/entity"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// ExecuteMathOperation execute a math operation
func (s *Service) ExecuteMathOperation(request entity.MathOperationRequest) (*entity.MathOperationResponse, error) {

	response := entity.MathOperationResponse{}

	switch request.Operation {
	case entity.Sum:
		response.Result = executeSumOperation(request.Parameters)
	case entity.Sub:
		return nil, errors.New("not implemented")
	default:
		return nil, errors.New("invalid operation")
	}

	if err := s.repo.PersistOperationLog(request, response); err != nil {
		return nil, errors.New("could not persist operation")
	}

	syslog.Info("p=math f=ExecuteMathOperation request=%+v response=%+v", request, response)
	return &response, nil
}

func executeSumOperation(parameters []float64) float64 {

	result := 0.0
	for _, p := range parameters {
		result += p
	}

	return result
}
