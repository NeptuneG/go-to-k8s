package handlers

import (
    "github.com/gorilla/mux"
)

/*
    だくさんのルートルールを対応するため、
    ルータの初期化とルールをpackageに含まれる
*/

// 必要なルートを登録し、ルータのインスタンスを返す
func Router() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/home", home).Methods("GET")
    return r
}