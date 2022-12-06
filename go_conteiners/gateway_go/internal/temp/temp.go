package temp

import (
	"net/http"
	"text/template"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("some template") // Create a template.
	data := 1
	t, _ = t.ParseFiles("site/index.html") // Parse template file.
	// Get current user infomration.
	t.Execute(w, data) // merge.
}
