package json

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
