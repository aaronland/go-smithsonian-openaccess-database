package main

import (
	_ "github.com/mattn/go-sqlite3"
)

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-smithsonian-openaccess-database/oembed"
	"log"
	"net/http"
)

func RandomHandler(db oembed.OEmbedDatabase) (http.HandlerFunc, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		record, err := db.GetRandomOEmbed(ctx)

		if err != nil {
			return
		}

		rsp.Header().Set("Content-type", "application/json")

		enc := json.NewEncoder(rsp)
		enc.Encode(record)

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil

}

func main() {

	server_uri := flag.String("server-uri", "http://localhost:8080", "...")
	dsn := flag.String("database-dsn", "sql://sqlite3/oembed.db", "...")
	flag.Parse()

	ctx := context.Background()

	db, err := oembed.NewSQLOEmbedDatabase(ctx, *dsn)

	if err != nil {
		log.Fatalf("Failed to create database, %v", err)
	}

	defer db.Close()

	mux := http.NewServeMux()

	random_handler, err := RandomHandler(db)

	if err != nil {
		log.Fatalf("Failed to create random handler, %v", err)
	}

	mux.Handle("/random/", random_handler)

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Unable to create server (%s), %v", *server_uri, err)
	}

	log.Printf("Listening on %s", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
