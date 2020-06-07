package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/tgmendes/guitrainer/internal/platform/db"
	"github.com/tgmendes/guitrainer/internal/routines"
	"log"
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

	rout, err := routines.Add(ctx, conn, routines.Routine{
		Name:        "1hr workout",
		Description: "workout for an houasdfasdfr",
		Level:       "this",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", rout)

	routs, err := routines.List(ctx, conn)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Routines: %+v\n", routs)

	rGet, _ := routines.Get(ctx, conn, 1)
	fmt.Printf("%+v\n", rGet)

	err = routines.Delete(ctx, conn, 1)
	if err != nil {
		println("ERROR?")
		fmt.Println(err.Error())
		return
	}

	err = routines.Delete(ctx, conn, 2)
	if err != nil {
		println("ERROR?")

		println(err.Error())
		return
	}
	println("NO ERROR")

}

func closeConn(ctx context.Context, conn *pgx.Conn) {
	_ = conn.Close(ctx)
}
