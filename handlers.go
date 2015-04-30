package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/garyburd/redigo/redis"
)

func StorageGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	vars := mux.Vars(r)
	key := vars["key"]

	conn := pool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", storageKey(key)))
	if err == redis.ErrNil {
		w.WriteHeader(http.StatusNotFound)
		return;
	}
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, value)
}

func StoragePut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	key := vars["key"]

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	value := string(body) // TODO: encoding?

	conn := pool.Get()
	defer conn.Close()

	set, err := redis.Bool(conn.Do("SETNX", storageKey(key), value))
	if err != nil {
		panic(err)
	}
	if !set {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func StoragePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	key := vars["key"]

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		writeError(w, err)
		return
	}
	if err := r.Body.Close(); err != nil {
		writeError(w, err)
		return
	}

	value := string(body) // TODO: encoding?

	conn := pool.Get()
	defer conn.Close()

	_, err = redis.String(conn.Do("SET", storageKey(key), value, "XX"))
	if err == redis.ErrNil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func StorageDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	key := vars["key"]

	conn := pool.Get()
	defer conn.Close()
	count, err := redis.Int(conn.Do("Del", key))
	if err != nil {
		panic(err)
	}
	if count != 1 {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}

func StorageOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "60")
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, err)
}

func storageKey(key string) string {
	return "storage:" + key
}