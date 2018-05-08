package main

import (
	"flag"
	"html/template"
	"os"
	"regexp"
	"strings"
)

var (
	repo        = flag.String("repo", "", "repository name")
	branchesStr = flag.String("branches", "", "branches")
)

var regSp = regexp.MustCompile(`\s`)

func main() {
	flag.Parse()
	if len(*repo) <= 0 {
		panic("repository name is required")
	}
	if len(*branchesStr) <= 0 {
		panic("branches are required")
	}

	htmlBase := `
<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title></title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css" integrity="sha384-WskhaSGFgHYWDcbwN70/dfYBj47jz9qbsMId/iRN3ewGhXQFZCSftd1LZCfmhktB" crossorigin="anonymous">
  </head>
  <body>
    <div class="container">
		<h1>API List for {{- .Repo -}}</h1>
		{{ $r := .Repo -}}
		<ul>
			{{ range .Branches -}}
				<li><a href="swagger-ui/index.html?repo={{ $r -}}&branch={{ . -}}">{{ . -}}</a></li>
			{{ end }}
        </ul>
    </div>
  </body>
</html>
`
	tmpBranches := regSp.Split(*branchesStr, -1)
	branches := make([]string, 0, 0)
	for _, b := range tmpBranches {
		if len(b) > 0 {
			branches = append(branches, strings.Replace(b, " ", "", -1))
		}
	}

	t := template.New("swagger ui template")
	t, err := t.Parse(htmlBase)
	if err != nil {
		panic(err)
	}

	t.Execute(os.Stdout, struct {
		Repo     string
		Branches []string
	}{
		Repo:     *repo,
		Branches: branches,
	})
}
