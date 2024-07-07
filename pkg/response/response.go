package response

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, r *http.Request, apiResponse *APIResponse) {
	response, err := json.Marshal(apiResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Headers", "Set-Cookie")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiResponse.StatusCode)
	w.Write(response)
}

// APIResponse defines general API response
type APIResponse struct {
	Message    string      `json:"message,omitempty"`
	Status     string      `json:"status,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	StatusCode int         `json:"-"`
}

// NewAPIResponse returns *APIResponse
func NewAPIResponse(status string, statusCode int) *APIResponse {
	return &APIResponse{Status: status, StatusCode: statusCode}
}
