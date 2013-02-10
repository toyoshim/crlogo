package main

import (
  "fmt"
  "net/http"
  "os"
)

func main() {
  http.Handle("/tmalib/", http.StripPrefix("/tmalib/", http.FileServer(http.Dir("tmalib"))))
  http.Handle("/", http.RedirectHandler("/tmalib/mozapp/chrome_logo.html", 301))
  fmt.Println("listening...")
  err := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
  if err != nil {
    panic(err)
  }
}
