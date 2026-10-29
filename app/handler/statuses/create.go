package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

type AddRequest struct {
	Status string `json:"status"`
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	account := auth.AccountOf(r)

	status := &object.Status{AccountID: account.ID, Content: req.Status}

	ctx := r.Context()
	statusRepo := h.app.Dao.Status()
	createdStatus, err := statusRepo.Create(ctx, status)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	createdStatus.Account = account

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(createdStatus); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
