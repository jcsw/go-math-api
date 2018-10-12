package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jcsw/go-math-api/pkg/entity"
	"github.com/jcsw/go-math-api/pkg/math"
)

// MathHandler handler to "/math/operation"
type MathHandler struct {
	MathService *math.Service
}

// Register function to handle "/customer"
func (m *MathHandler) Register(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		m.executeMathOperation(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "invalid method")
	}

}

func (m *MathHandler) executeMathOperation(w http.ResponseWriter, r *http.Request) {

	reader := r.Body
	defer reader.Close()

	var request entity.MathOperationRequest
	if err := json.NewDecoder(reader).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if err := request.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := m.MathService.ExecuteMathOperation(request)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not complete operation")
		return
	}

	respondWithJSON(w, http.StatusOK, response)
}
