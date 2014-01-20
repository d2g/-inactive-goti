package goti_test

import (
	template "github.com/d2g/goti"
	"os"
	"testing"
)

func Test_Inheritance(test *testing.T) {

	/*
	 * This is tha Basic Template with our Standard CSS etc etc.
	 * If as a company we make a branding decision i.e. Colour Change.
	 * We only need to updated it in base to chacnge all templates.
	 */

	const base = `
	{{define "body"}}
		<h1>Error:</h1>
		The Default "Body" Template has not Been Overridden.
	{{end}}
	<html>
		<body>
			{{template "body"}}
		</body>
	</html>
	`

	/*
	 * This is our standard Department Template. Each Part of the company would then have their own
	 * Children templates Making renaming a departments (At least in the page design easier).
	 */
	const department_template = `
	{{define "department_title"}}Default Department{{end}}
	{{define "department_email"}}webmaster@company.com?Subject=Error%20Default%20Department{{end}}
	{{define "department_body"}}The Default Department Template("department_body") has not Been Overridden.{{end}}
	{{define "body"}}
		<h1><a href="mailto:{{template "department_email"}}" target="_top">{{template "department_title"}}</h1>
		<p>
			{{template "department_body"}}
		</p>
	{{end}}
	`

	/*
	 *  All departments within the company
	 * forfill the parts of this template. I.E Department Name.
	 */
	const integration_department = `
	{{define "department_title"}}Integration{{end}}
	{{define "department_email"}}integrationteam@company.com{{end}}
	{{define "department_body"}}The Integration Teams Default Page.. Please contact <Insert Name Here> in the integration team about this page..{{end}}
	`

	/*
	 * This is a specific page instance.
	 */
	const about = `
	{{define "department_body"}}This is the about page for the integration team{{end}}
	`

	/*
	 * Sometime with the best will in the world the integration department want access to the full page without the company borders, or want to change the css etc.
	 * So being able to overwrite the top template from the bottom is always worth doing.
	 */
	const bottom = `
	{{define "body"}}
		Overwrite the top from the bottom.. (We wasted all the processing in the middle but that not my problem...)
	{{end}}`

	t, err := template.New("Example").Parse(base)
	if err != nil {
		test.Fatalf("Unexpected Parse Error:%v\n", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		test.Fatalf("Base Error:%v\n", err)
	}

	t, err = t.Parse(department_template)
	if err != nil {
		test.Fatalf("Department Template Parse Error:%v\n", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		test.Fatalf("Department Template Error:%v\n", err)
	}

	t, err = t.Parse(integration_department)
	if err != nil {
		test.Fatalf("Department Parse Error:%v\n", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		test.Fatalf("Department Error:%v\n", err)
	}

	t, err = t.Parse(about)
	if err != nil {
		test.Fatalf("About Parse Error:%v\n", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		test.Fatalf("About Error:%v\n", err)
	}

	t, err = t.Parse(bottom)
	if err != nil {
		test.Fatalf("Bottom Parse Error:%v\n", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		test.Fatalf("Bottom Error:%v\n", err)
	}
}

func Test_InheritanceExampleTemplate(test *testing.T) {
	const grandparent = `
	{{define "body"}}
		<h1>Error:</h1>
		The Default "Body" Template has not Been Overridden.
	{{end}}
	<html>
		<body>
			{{template "body" .}}
		</body>
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
		</ol>
		{{template "message" .}}
	{{end}}
	`

	const child1 = `
	{{define "message"}}
		Dear {{.Name}},
			{{if .Attended}}
				It was a great to see you at my birthday.
			{{else}}
				It is a shame you couldn't make it to my birthday.
			{{end}}
				{{with .Gift}}Thank you for the lovely {{.}}.
			{{end}}
		Best wishes,
		Dan
	{{end}}
	`

	const child2 = `
	{{define "message"}}
		Dear {{.Name}},
			{{if .Attended}}
				It was a great to see you at Christmas.
			{{else}}
				It is a shame I didn't see you at Christmas.
			{{end}}
				{{with .Gift}}Thank you for the {{.}}.
			{{end}}
		Best wishes,
		Dan
	{{end}}
	`

	t, err := template.New("Thanks").Parse(grandparent)
	if err != nil {
		test.Fatalf("Unexpected Parse Error:%v\n", err)
	}

	t, err = t.Parse(parent)
	if err != nil {
		test.Fatalf("Parent Parse Error:%v\n", err)
	}

	t, err = t.Parse(child1)
	if err != nil {
		test.Fatalf("child1 Parse Error:%v\n", err)
	}

	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			test.Fatalf("Error Executing:%v Error:%v\n", r, err)
		}
	}

	t, err = t.Parse(child2)
	if err != nil {
		test.Fatalf("child2 Parse Error:%v\n", err)
	}

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			test.Fatalf("Error Executing:%v Error:%v\n", r, err)
		}
	}
}
