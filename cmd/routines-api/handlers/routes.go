package handlers

import (
	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
)

func API(db *pgx.Conn) *httprouter.Router {
	router := httprouter.New()

	rh := Routine{db: db}
	router.GET("/api/routines", rh.List)
	router.POST("/api/routines", rh.Add)

	return router
}
