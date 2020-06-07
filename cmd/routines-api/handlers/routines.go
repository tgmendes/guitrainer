package handlers

import (
	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/tgmendes/guitrainer/internal/platform/web"
	"github.com/tgmendes/guitrainer/internal/routines"
	"net/http"
)

type Routine struct {
	db *pgx.Conn
}

type ListResponse struct {
	Routines []routines.Routine `json:"routines"`
}

func (rt *Routine) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	routs, _ := routines.List(r.Context(), rt.db)

	resp := ListResponse{Routines: routs}
	_ = web.Respond(w, http.StatusOK, resp)

}

func (rt *Routine) Add(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var nr routines.Routine

	if err := web.Decode(r, &nr); err != nil {
		_ = web.RespondError(w, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	_, err := routines.Add(r.Context(), rt.db, nr)
	if err != nil {
		_ = web.RespondError(w, http.StatusInternalServerError, "couldn't create routine")
		return
	}

	_ = web.Respond(w, http.StatusCreated, nil)
}
