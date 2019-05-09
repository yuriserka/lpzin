package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	temp, err := template.ParseFiles("./static/views/xotomate.html")
	if err != nil {
		panic(err)
	}

	t := template.Must(temp, err)
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
	var err error
	temp, err := template.ParseFiles("./static/views/index.html")
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(temp, err)
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
	err = tmpl.Execute(w, formatedHour(time.Now()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
