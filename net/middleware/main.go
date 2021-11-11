package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/",myMiddleware(MyHandler))
	http.ListenAndServe(":8080",nil)
}

func MyHandler(w http.ResponseWriter,_ *http.Request){
	w.Write([]byte("Hello:\n"))
	w.Write([]byte("This is a new page.\n"))
	w.WriteHeader(http.StatusOK)
}

func myMiddleware(next http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("next函数进入之前\r\n"))
		next(w,r)
		w.Write([]byte("next函数进入之后\r\n"))
	}
}