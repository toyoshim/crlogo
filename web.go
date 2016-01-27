package main

import (
  "fmt"
  "net/http"
  "os"
  "time"
)

type wrappedWriter struct {
  http.ResponseWriter
  status int
  length int
}

func (w *wrappedWriter) WriteHeader(status int) {
  w.status = status
  w.ResponseWriter.WriteHeader(status)
}

func (w *wrappedWriter) Write(b []byte) (int, error) {
  if w.status == 0 {
    w.status = 200
  }
  w.length = len(b)
  return w.ResponseWriter.Write(b)
}

func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    writer := wrappedWriter{ w, 0, 0 }
    start := time.Now()
    handler.ServeHTTP(&writer, r)
    duration := time.Now().Sub(start)
    fmt.Printf("%s %s %s %s %s %s %s %d %d %f %s - %s\n",
        time.Now().Format(time.RFC3339),
        r.Header["Referer"][0],
        r.Method,
        r.Host,
        r.URL.Path,
        r.Proto,
        r.Header["Accept-Language"][0],
        writer.status,
        writer.length,
        duration.Seconds() * 1000,
        r.RemoteAddr,
        r.UserAgent())
  })
}

func main() {
  http.Handle("/tmalib/", http.StripPrefix("/tmalib/", http.FileServer(http.Dir("tmalib"))))
  http.Handle("/", http.RedirectHandler("/tmalib/mozapp/chrome_logo.html", 301))
  fmt.Println("listening on port " + os.Getenv("PORT") + "...")
  err := http.ListenAndServe(":" + os.Getenv("PORT"), Log(http.DefaultServeMux))
  if err != nil {
    panic(err)
  }
}
