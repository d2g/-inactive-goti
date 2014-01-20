GOTI
====

Go Template Inheritance

This is a tweak to the text/template package to removing the redefinition error (from template.go) and test (from multi_test.go). This allows you to redefine templates/define blocks in templates.

This with a tweaked version of html/template package to use the tweaked text/template allows you to use Inheritance in both packages.

To use replace:

`import ("html/template" )`

With:

`import (template "github.com/d2g/goti/html")`

Or:

`import ("text/template")`

With:

`import (template "github.com/d2g/goti/text")`


Although this is not optimal as both blocks get parsed it does work. I'm sure in the long term [This Issue/Change Request/Enhancement/Feature 3812](https://code.google.com/p/go/issues/detail?id=3812) will bring a better solution.

Example (Included in GoDoc):
	
	const grandparent = `
	{{define "body"}}
		<h1>Error:</h1>
		The Default "Body" Template has not Been Overridden.
	{{end}}
	<html>
		<body>{{template "body" .}}</body>
	</html>
	`

	const parent = `
	{{define "message"}}{{end}}
	{{define "body"}}
		<ol>
			<li>Address Line 1</li>
			<li>Address Line 2</li>
			<li>Address Line 3</li>
			<li>Address Line 4</li>
			<li>Address Line 5</li>
			<li>Post Code</li>
		</ol>{{template "message" .}}{{end}}
	`

	const child1 = `
	{{define "message"}}
		Dear {{.Name}},
			{{if .Attended}}It was a great to see you at my birthday.{{else}}It is a shame you couldn't make it to my birthday.{{end}}
			{{with .Gift}}Thank you for the lovely {{.}}.{{end}}
		Best wishes,
		Dan
		{{end}}
	`

	const child2 = `
	{{define "message"}}
		Dear {{.Name}},
			{{if .Attended}}It was a great to see you at Christmas.{{else}}It is a shame I didn't see you at Christmas.{{end}}
			{{with .Gift}}Thank you for the {{.}}.{{end}}
		Best wishes,
		Dan
		{{end}}
	`
	// Create a new template and parse the template hierarchy
	t := template.Must(template.New("Thanks").Parse(grandparent))
	t = template.Must(t.Parse(parent))
	t = template.Must(t.Parse(child1))

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Execute the Birthday template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Fatalf("Error Executing:%v Error:%v\n", r, err)
		}
	}

	//Change the Birthday Template For the Christmas Template
	t = template.Must(t.Parse(child2))

	// Execute the Christmas template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Fatalf("Error Executing:%v Error:%v\n", r, err)
		}
	}
`
