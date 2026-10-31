package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
	"yatter-backend-go/app/handler/request"
)

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := request.IDOf(r)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	ctx := r.Context()
	status, err := h.app.Dao.Status().FindByID(ctx, id)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	if status == nil {
		httperror.Error(w, http.StatusNotFound)
		return
	}
	account := auth.AccountOf(r)
	if account.ID != status.AccountID {
		httperror.Error(w, http.StatusForbidden)
		return
	}
	if err = h.app.Dao.Status().Delete(ctx, status.ID); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(struct{}{}); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
