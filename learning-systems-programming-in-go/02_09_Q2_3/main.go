package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "Wirld",
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "     ")
	encoder.Encode(source)

	file, err := os.Create("text.json.gz")
	if err != nil {
		panic(err)
	}

	gzipw := gzip.NewWriter(file)
	gzipw.Header.Name = "text.json"
	gzipw.Flush()
	io.MultiWriter(w, gzipw)
	gzipw.Close()

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
