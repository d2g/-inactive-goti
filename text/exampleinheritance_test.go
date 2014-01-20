package goti_test

import (
	template "github.com/d2g/goti/text"
	"log"
	"os"
)

func Example_inheritance() {
	// Define a templates
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

	//Output: 	<html>
	//				<body>
	//					<ol>
	//						<li>Address Line 1</li>
	//						<li>Address Line 2</li>
	//						<li>Address Line 3</li>
	//						<li>Address Line 4</li>
	//						<li>Address Line 5</li>
	//						<li>Post Code</li>
	//					</ol>
	//					Dear Aunt Mildred,
	//						It was a great to see you at my birthday.
	//						Thank you for the lovely bone china tea set.
	//					Best wishes,
	//					Dan
	//				</body>
	//			</html>

	//			<html>
	//				<body>
	//					<ol>
	//						<li>Address Line 1</li>
	//						<li>Address Line 2</li>
	//						<li>Address Line 3</li>
	//						<li>Address Line 4</li>
	//						<li>Address Line 5</li>
	//						<li>Post Code</li>
	//					</ol>
	//					Dear Uncle John,
	//						It is a shame you couldn't make it to my birthday.
	//						Thank you for the lovely moleskin pants.
	//					Best wishes,
	//					Dan
	//				</body>
	//			</html>

	//			<html>
	//				<body>
	//					<ol>
	//						<li>Address Line 1</li>
	//						<li>Address Line 2</li>
	//						<li>Address Line 3</li>
	//						<li>Address Line 4</li>
	//						<li>Address Line 5</li>
	//						<li>Post Code</li>
	//					</ol>
	//					Dear Cousin Rodney,
	//						It is a shame you couldn't make it to my birthday.
	//
	//					Best wishes,
	//					Dan
	//				</body>
	//			</html>

	//			<html>
	//				<body>
	//					<ol>
	//						<li>Address Line 1</li>
	//						<li>Address Line 2</li>
	//						<li>Address Line 3</li>
	//						<li>Address Line 4</li>
	//						<li>Address Line 5</li>
	//						<li>Post Code</li>
	//					</ol>
	//					Dear Aunt Mildred,
	//						It was a great to see you at Christmas.
	//						Thank you for the bone china tea set.
	//					Best wishes,
	//					Dan
	//				</body>
	//			</html>

	//			<html>
	//				<body>
	//					<ol>
	//						<li>Address Line 1</li>
	//						<li>Address Line 2</li>
	//						<li>Address Line 3</li>
	//						<li>Address Line 4</li>
	//						<li>Address Line 5</li>
	//						<li>Post Code</li>
	//					</ol>
	//					Dear Uncle John,
	//						It is a shame I didn't see you at Christmas.
	//						Thank you for the moleskin pants.
	//					Best wishes,
	//					Dan
	//				</body>
	//			</html>

	//			<html>
	//				<body>
	//					<ol>
	//						<li>Address Line 1</li>
	//						<li>Address Line 2</li>
	//						<li>Address Line 3</li>
	//						<li>Address Line 4</li>
	//						<li>Address Line 5</li>
	//						<li>Post Code</li>
	//					</ol>
	//					Dear Cousin Rodney,
	//						It is a shame I didn't see you at Christmas.
	//
	//					Best wishes,
	//					Dan
	//				</body>
	//			</html>
}
