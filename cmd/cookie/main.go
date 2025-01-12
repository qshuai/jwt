package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/setcookie", setCookie)
	http.HandleFunc("/override_cookie", overrideCookie)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "setup http server err: %v", err)
		os.Exit(1)
	}
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "username=qshuai")
	w.Header().Add("Set-Cookie", "session_id=1a2997b8-555b-49cd-8837-95e6832a5a59")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Set-Cookie 测试"))
}

func overrideCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "username=github")
	w.Header().Add("Set-Cookie", "session_id=a97f9b39-3355-43ee-bf9b-36575e64e0ba")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("cookie覆盖测试"))
}
