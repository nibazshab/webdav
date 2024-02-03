package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "golang.org/x/net/webdav"
)

func main() {
    port := "8088"
    var rootfs string

    o := len(os.Args)
    if o == 1 {
        fmt.Println("WARNING: Using the default directory (./)")
        rootfs = "."
    } else if o == 2 {
        rootfs = os.Args[1]
    } else {
        fmt.Println("ERROR: Just need 1 parameter")
        return
    }

    rootfs, _ = filepath.Abs(rootfs)

    fs := &webdav.Handler{
        FileSystem: webdav.Dir(rootfs),
        LockSystem: webdav.NewMemLS(),
    }

    fmt.Println("port:", port, "\nrootfs:", rootfs)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "PROPPATCH", "MKCOL", "DELETE", "PUT", "COPY", "MOVE", "LOCK", "UNLOCK":
            http.Error(w, "Permission denied", http.StatusForbidden)
            return
        }

        fs.ServeHTTP(w, r)

        xff := r.Header.Get("X-Forwarded-For")
        if xff == "" {
            xff = r.RemoteAddr
        }
        log.Print(xff + " - " + r.Method + " - " + r.URL.String() + " - " + r.Header.Get("user-agent"))
    })

    http.ListenAndServe(":"+port, nil)
}
