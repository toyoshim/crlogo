package main

import (
	"fmt"
	"net/http"
	"os"

//	"github.com/toyoshim/gomolog"
)

func main() {
//	log := gomolog.Open(os.Getenv("MONGOLAB_URI"), "log")
//	defer log.Close()

	http.Handle("/tmalib/", http.StripPrefix("/tmalib/", http.FileServer(http.Dir("tmalib"))))
	http.Handle("/", http.RedirectHandler("/tmalib/mozapp/chrome_logo.html", 301))
	fmt.Println("listening on port " + os.Getenv("PORT") + "...")
//	err := http.ListenAndServe(":"+os.Getenv("PORT"), log.Logger())
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
