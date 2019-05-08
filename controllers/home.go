package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/views/xotomate.html"))
	t.Execute(w, struct {
		Nome     string
		Idade    int
		Solteiro bool
	}{
		Nome:     "Yuri",
		Idade:    20,
		Solteiro: true,
	})
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/views/index.html"))
	formatedHour := func(t time.Time) string {
		var h, m string
		if t.Hour() < 10 {
			h = "0" + strconv.Itoa(t.Hour())
		} else {
			h = strconv.Itoa(t.Hour())
		}
		if t.Minute() < 10 {
			m = "0" + strconv.Itoa(t.Minute())
		} else {
			m = strconv.Itoa(t.Minute())
		}
		return h + ":" + m
	}
	tmpl.Execute(w, formatedHour(time.Now()))
}
