package main

import (
  "fmt"
  "net/http"
  "os"
)

func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
  })
}

func main() {
  http.Handle("/tmalib/", http.StripPrefix("/tmalib/", http.FileServer(http.Dir("tmalib"))))
  http.Handle("/", http.RedirectHandler("/tmalib/mozapp/chrome_logo.html", 301))
  fmt.Println("listening...")
  err := http.ListenAndServe(":" + os.Getenv("PORT"), Log(http.DefaultServeMux))
  if err != nil {
    panic(err)
  }
}
