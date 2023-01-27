package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main() {
	content, err := ioutil.ReadFile("./resume.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	log.Print(payload["basics"])

	var tmpl = `
    {{with .basics}}
    hello {{.email}}
    {{end}}`
	t2 := template.Must(template.New(tmpl).Parse(tmpl))

	err2 := t2.Execute(os.Stdout, payload)
	if err != nil {
		log.Println("executing template:", err2)
	}

	// Define a template.
	const letter = `
----
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

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

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// create a new file
	file, _ := os.Create("greeting2.txt")
	defer file.Close()

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(file, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
