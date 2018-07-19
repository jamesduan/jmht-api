package http

import (
	"html/template"
	"net/http"
)

func views() {

	// http.HandleFunc("/aboutus", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("templates/aboutus.html"))
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w, nil)
	// 		return
	// 	}
	// 	tmpl.Execute(w, struct{ Success bool }{true})
	// })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		tmpl.Execute(w, struct{ Success bool }{true})
	})

	// http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("templates/contact.html"))
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w, nil)
	// 		return
	// 	}
	// 	tmpl.Execute(w, struct{ Success bool }{true})
	// })

	// http.HandleFunc("/joinus", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("templates/joinus.html"))
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w, nil)
	// 		return
	// 	}
	// 	tmpl.Execute(w, struct{ Success bool }{true})
	// })

	// http.HandleFunc("/newsDetail", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("templates/news_detail.html"))
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w, nil)
	// 		return
	// 	}
	// 	tmpl.Execute(w, struct{ Success bool }{true})
	// })

	// http.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("templates/news.html"))
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w, nil)
	// 		return
	// 	}
	// 	tmpl.Execute(w, struct{ Success bool }{true})
	// })

	// http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("templates/product.html"))
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w, nil)
	// 		return
	// 	}
	// 	tmpl.Execute(w, struct{ Success bool }{true})
	// })

	// http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("templates/product.html"))
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w, nil)
	// 		return
	// 	}
	// 	tmpl.Execute(w, struct{ Success bool }{true})
	// })

	// http.HandleFunc("/aboutus", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("templates/aboutus.html"))
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w, nil)
	// 		return
	// 	}
	// 	tmpl.Execute(w, struct{ Success bool }{true})
	// })

	// http.HandleFunc("/apple", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("templates/apple.html"))
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w, nil)
	// 		return
	// 	}
	// 	tmpl.Execute(w, struct{ Success bool }{true})
	// })
}
