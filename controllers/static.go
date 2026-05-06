package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {

	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "Is it free?",
			Answer:   "Yes it is fully free and Open-Source.",
		},

		{
			Question: "What is the meaning of life?",
			Answer:   "42.",
		},

		{
			Question: "Lorem?",
			Answer:   "Ipsum.",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
