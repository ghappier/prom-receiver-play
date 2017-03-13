package main

import (
	"net/http"

	"github.com/lizq/prom-receiver-play/pool"
	"github.com/lizq/prom-receiver-play/storage"
)

func main() {
	kafka := new(storage.KafkaStorage)
	mongodb := new(storage.MongodbStorage)
	holder := storage.NewStorageHolder()
	holder.Attach(kafka)
	holder.Attach(mongodb)

	dispatcher := pool.NewDispatcher(pool.MaxWorker, holder)
	dispatcher.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("/receive", pool.PayloadHandler)
	//mux.HandleFunc("/", echoHandler)
	http.ListenAndServe(":9990", mux)
}
