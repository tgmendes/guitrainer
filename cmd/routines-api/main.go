package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/tgmendes/guitrainer/cmd/routines-api/handlers"
	"github.com/tgmendes/guitrainer/internal/platform/db"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	cfg := db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "somepassword",
		DBName:   "guitrainer",
	}
	conn, err := db.Open(ctx, cfg)
	if err != nil {
		panic(err.Error())
	}
	defer closeConn(ctx, conn)

	router := handlers.API(conn)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func closeConn(ctx context.Context, conn *pgx.Conn) {
	_ = conn.Close(ctx)
}
