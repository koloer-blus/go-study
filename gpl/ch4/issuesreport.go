package main

import (
	"go-study/gpl/ch4/github"
	"html/template"
	"log"
	"os"
	"time"
)

const temp1 = `{{.TotalCount}} issues:
{{range.Items}}----------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	report, _ := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(temp1)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
