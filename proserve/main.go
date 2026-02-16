package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jboursiquot/go-proverbs"
)

type Proverb struct {
	Saying string `json:"saying"`
	Link   string `json:"link"`
}

func HandleProverbJson(w http.ResponseWriter, r *http.Request) {
	randomProverb := proverbs.Random()

	p := Proverb{
		Saying: randomProverb.Saying,
		Link:   randomProverb.Link,
	}

	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		http.Error(w, "oops, that went wrong", http.StatusInternalServerError)
	}
}

func HandleProverbPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("page").Parse(`
	</html>
		<head></head>
		<body>
			<h1> Your Proverb: <a href="{{ .Link }}">{{ .Saying }}</a></h1>
		</body>
	</html>`)
	err := t.Execute(w, proverbs.Random())
	if err != nil {
		http.Error(w, "oops, that went wrong", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", HandleProverbPage)

	// PLain Text Endpoint
	http.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, proverbs.Random().Saying)
		if err != nil {
			http.Error(w, "oops, that went wrong", http.StatusInternalServerError)
		}
	})

	// REST API JSON Endpoint
	http.HandleFunc("/api", HandleProverbJson)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
