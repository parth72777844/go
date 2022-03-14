package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

var certlocs []string = []string{
	"/Users/k1n1/Desktop/k1n1-server/src/app.mrk1n1.tk",
	"/Users/k1n1/Desktop/k1n1-server/src/app.mrk1n1.tk"}

type IHandler struct{}

type SHandler struct{}

func (ih IHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostdom := strings.Split(r.Host, ":")[0]
	http.Redirect(w, r, "https://"+hostdom+":8181"+r.URL.Path, 302)
}

func (sh SHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I hope you feel secure now you are here")
}

func (sh SHandler) router(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w,r,"https://app.mrk1n1.tk", 302)
}


func main() {
	tconf := &tls.Config{}

	for _, v := range certlocs {
		cert, err := tls.LoadX509KeyPair(path.Join(v, "certificate.crt"), path.Join(v, "private.key"))
		if err != nil {
			log.Fatal(err)
		}
		tconf.Certificates = append(tconf.Certificates, cert)
	}

	tconf.BuildNameToCertificate()

	go func() {
		log.Fatal(http.ListenAndServe(":80", IHandler{}))
	}()

	sserv := http.Server{
		Addr:      ":8181",
		Handler:   SHandler{},
		TLSConfig: tconf,
	}

	log.Fatal(sserv.ListenAndServeTLS("", ""))

}