package http

import (
	"net/http"
)

func configTestRoutes() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, map[string]string{"Test": "ok", "Version": "1.0.1"})
	})
}

// func configAd() {
// 	http.HandleFunc("/ad", func(w http.ResponseWriter, r *http.Request) {
// 		tmpl := template.Must(template.ParseFiles("templates/ad.html"))
// 		if r.Method != http.MethodPost {
// 			tmpl.Execute(w, nil)
// 			return
// 		}
// 		tmpl.Execute(w, struct{ Success bool }{true})
// 	})
// }
