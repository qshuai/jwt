package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	port     = flag.String("port", "443", "the port of http server listened on")
	https    = flag.Bool("tls", false, "whether to enable https server")
	certfile = flag.String("cert-file", "server.cert", "the certification file path")
	keyfile  = flag.String("key-file", "server.key", "the private key file path")
)

func main() {
	flag.Parse()

	http.HandleFunc("/setcookie", setCookie)
	http.HandleFunc("/override_cookie", overrideCookie)

	if *https {
		err := http.ListenAndServeTLS(":"+*port, *certfile, *keyfile, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "setup https server err: %v", err)
			os.Exit(1)
		}
	} else {
		err := http.ListenAndServe(":"+*port, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "setup http server err: %v", err)
			os.Exit(1)
		}
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
