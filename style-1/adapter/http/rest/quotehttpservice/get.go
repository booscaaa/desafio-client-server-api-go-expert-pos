package quotehttpservice

import (
	"encoding/json"
	"net/http"
)

func (service service) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	quotes, err := service.usecase.Get(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(quotes)
}
