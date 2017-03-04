package pack

import (
	"html/template"
	"os"
)

func SayHelloReportallocs(user UserReportallocs) {
	t, _ := template.New("").Parse("Hello {{.Name}}\n")
	t.Execute(os.Stdout, user)
}

type UserReportallocs struct {
	Name string
}

func NewUserReportallocs(name string) UserReportallocs {
	result := UserReportallocs{name}
	return result
}
