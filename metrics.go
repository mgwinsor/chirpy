package main

import (
	"bytes"
	"html/template"
	"net/http"
)

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("admin/metrics.html")
	if err != nil {
		http.Error(w, "Couldn't parse template", http.StatusInternalServerError)
		return
	}

	data := struct{ Hits int32 }{Hits: cfg.fileserverHits.Load()}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		http.Error(w, "Couldn't execute template", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
