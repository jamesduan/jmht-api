package http

import (
	"jmht-api/g"
	"net/http"
)

func configStatic() {
	fs := http.FileServer(http.Dir(g.Config().Image.FilePath))
	http.Handle("/", fs)
}
