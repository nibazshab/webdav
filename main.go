package main

import (
  "net/http"
  "golang.org/x/net/webdav"
  "fmt"
  "os"
  "path/filepath"
  "golang.org/x/net/context"
  "strconv"
)

func handleDirList(fs webdav.FileSystem, w http.ResponseWriter, req *http.Request) bool {
  ctx := context.Background()
  f, _ := fs.OpenFile(ctx, req.URL.Path, os.O_RDONLY, 0)
  defer f.Close()
  if fi, _ := f.Stat(); fi != nil && !fi.IsDir() {
    return false
  }
  dirs, _ := f.Readdir(-1)

  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  fmt.Fprintf(w, "<pre>\n")
  for _, d := range dirs {
    name := d.Name()
    if d.IsDir() {
      name += "/"
    }
    size := ""
    if !d.IsDir() {
      size = strconv.FormatInt(d.Size(), 10)
    }
    fmt.Fprintf(w, "<a href=\"%s\">%s</a> %s\n", name, name, size)
  }
  fmt.Fprintf(w, "</pre>\n")
  return true
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Please provide a directory path.")
    return
  }

  rootDir := os.Args[1]
  absRootDir, _ := filepath.Abs(rootDir)

  fs := &webdav.Handler{
    FileSystem: webdav.Dir(absRootDir),
    LockSystem: webdav.NewMemLS(),
  }

  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    if req.Method == "GET" && handleDirList(fs.FileSystem, w, req) {
      return
    }
    fs.ServeHTTP(w, req)
  })

  addr := "8088"
  fmt.Printf("WebDAV server listening on %s\n", addr)
  _ = http.ListenAndServe(":" + addr, nil)
}
