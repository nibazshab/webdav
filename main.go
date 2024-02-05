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
	web := "/web/"

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PROPPATCH", "MKCOL", "DELETE", "PUT", "COPY", "MOVE", "LOCK", "UNLOCK":
			http.Error(w, "Permission denied", http.StatusForbidden)
			return
		}

		if len(r.URL.Path) >= len(web) && r.URL.Path[:len(web)] == web {
			http.StripPrefix(web, http.FileServer(http.Dir(rootfs))).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}

		logRecord(r)
	})

	fmt.Println("port:", port, "\nrootfs:", rootfs, "\nweb:", web)
	http.ListenAndServe(":"+port, nil)
}

func logRecord(r *http.Request) {
	log.Print(r.RemoteAddr + " - " + r.Method + " - " + r.URL.String() + " - " + r.Header.Get("user-agent"))
}
