package main

import "net/http"

func main(){
	// routing 
	http.HandleFunc("/hello", func(w http.ResponseWriter , r *http.Request){
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":3000",nil)
}