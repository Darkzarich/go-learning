package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

func main() {
	tmpl, err := template.New("page").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>Hello</title>
</head>
<body>
	<p> You connected from {{.IP}} and request this page {{ .Count }} times </p>
</body>
</html>
`)

	if err != nil {
		panic(err)
	}

	var (
		countsByIP = make(map[string]uint, 10)
		mu         sync.Mutex
	)

	getTimes := func(IP string) uint {
		mu.Lock()
		defer mu.Unlock()
		countsByIP[IP]++
		return countsByIP[IP]
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		IP := r.RemoteAddr
		count := getTimes(IP)

		err := tmpl.Execute(w, struct {
			IP    string
			Count uint
		}{IP, count})
		if err != nil {
			log.Println(err)
		}
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
