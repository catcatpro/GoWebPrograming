package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func page1(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("./templates/tmp1.html")
	t, _ := template.ParseGlob("./templates/*.html")
	t.Execute(w, "Hello, World")
}

func page2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/tmpl_base1.html", "./templates/tmpl_base2.html")
	t.Execute(w, "Hello, World")
}

func page3(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/tmpl_base1.html", "./templates/tmpl_base2.html")
	// t.Execute(w, "Hello, World")

	t.ExecuteTemplate(w, "tmpl_base2.html", "Hello, World")
}

func action_process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/tmpl_condition.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func action_loop(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/tmpl_loop.html")
	// daysOfWeek := []string{}
	daysOfWeek := []string{"Mon", "Tue"}
	t.Execute(w, daysOfWeek)
}

func set_action(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/templ_set_action.html")
	t.Execute(w, "hello")
}

func include_action(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/t1.html", "templates/t2.html")
	t.Execute(w, "Hello, World!")
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func do_format_date(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("formt_date.html").Funcs(funcMap)
	t, _ = t.ParseFiles("templates/formt_date.html")
	t.Execute(w, time.Now())
}

func context_aware(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/context_aware.html")
	content := `I asked <i>What's up?</i>`
	t.Execute(w, content)
}

func action_xss(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/xss.html")
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/form.html")
	t.Execute(w, nil)
}

func define_action(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("templates/layout.html")
	// t.ExecuteTemplate(w, "layout", "")

	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("templates/layout.html", "templates/red_hello.html")
	} else {
		// t, _ = template.ParseFiles("templates/layout.html", "templates/blue_hello.html")
		t, _ = template.ParseFiles("templates/layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}
