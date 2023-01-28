package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type Resume struct {
	Schema string `json:"$schema"`
	Meta   struct {
		Theme     string `json:"theme"`
		WorkLabel string `json:"workLabel"`
	} `json:"meta"`
	Basics struct {
		Name       string   `json:"name"`
		Email      string   `json:"email"`
		Phone      string   `json:"phone"`
		URL        string   `json:"url"`
		Summary    string   `json:"summary"`
		Highlights []string `json:"highlights,omitempty"`
		Location   struct {
			Address    string `json:"address"`
			City       string `json:"city"`
			Region     string `json:"region"`
			PostalCode string `json:"postalCode"`
		} `json:"location"`
		Profiles []struct {
			Network  string `json:"network"`
			Username string `json:"username"`
			URL      string `json:"url"`
		} `json:"profiles"`
	} `json:"basics"`
	Work []struct {
		Name        string   `json:"name"`
		Position    string   `json:"position"`
		Location    string   `json:"location"`
		StartDate   string   `json:"startDate"`
		Description string   `json:"description,omitempty"`
		EndDate     string   `json:"endDate,omitempty"`
		Summary     string   `json:"summary"`
		Highlights  []string `json:"highlights,omitempty"`
	} `json:"work"`
	Education []struct {
		Institution string `json:"institution"`
		URL         string `json:"url"`
		Area        string `json:"area"`
		StudyType   string `json:"studyType"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
	} `json:"education"`
	Certificates []struct {
		Name   string `json:"name"`
		Date   string `json:"date"`
		Issuer string `json:"issuer"`
		URL    string `json:"url"`
	} `json:"certificates"`
	Skills []struct {
		Name     string   `json:"name"`
		Keywords []string `json:"keywords"`
	} `json:"skills"`
	Projects []struct {
		Name        string   `json:"name"`
		Summary     string   `json:"summary"`
		Description string   `json:"description,omitempty"`
		Highlights  []string `json:"highlights,omitempty"`
		Keywords    []string `json:"keywords,omitempty"`
		EndDate     string   `json:"endDate"`
	} `json:"projects"`
	AdditionalWork []struct {
		Name        string   `json:"name"`
		Position    string   `json:"position"`
		Location    string   `json:"location"`
		StartDate   string   `json:"startDate"`
		Description string   `json:"description,omitempty"`
		EndDate     string   `json:"endDate,omitempty"`
		Summary     string   `json:"summary"`
		Highlights  []string `json:"highlights,omitempty"`
	} `json:"additinalWork"`
	Publications []struct {
		Name        string `json:"name"`
		Summary     string `json:"summary"`
		Publisher   string `json:"publisher"`
		ReleaseDate string `json:"releaseDate"`
		URL         string `json:"url"`
	} `json:"publications"`
	AdditionalPublications []struct {
		Name        string `json:"name"`
		Summary     string `json:"summary"`
		Publisher   string `json:"publisher"`
		ReleaseDate string `json:"releaseDate"`
		URL         string `json:"url"`
	} `json:"additionalPublications"`
}

func parse() {
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatal("invalid JSON Resume")
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}

var funcMap = template.FuncMap{
	"MMYYYY": func(value string) string {
		date, err := time.Parse("2006-01-02", value)
		if err != nil {
			log.Fatal(err)
		}
		return date.Format("01/2006")
	},
	"MMMMYYYY": func(value string) string {
		date, err := time.Parse("2006-01-02", value)
		if err != nil {
			log.Fatal(err)
		}
		return date.Format("January 2006")
	},
	"plus1": func(i int) int {
		return i + 1
	},
	"toUpper": strings.ToUpper,
}

func render() {
	base := filepath.Base(templateFile)
	t := template.Must(template.New(base).Funcs(funcMap).ParseFiles(templateFile))

	// // create a new file
	// err := os.Mkdir("artifact", os.ModePerm)
	// outpath := filepath.Join("artifact", strings.TrimSuffix(filepath.Base(templateFile), filepath.Ext((templateFile))))
	file, _ := os.Create(outputFile)
	defer file.Close()

	err := t.Execute(file, payload)
	// err = t.Execute(file, payload)
	if err != nil {
		log.Fatal("Error executing template: ", err)
	}
}

// func main() {
// 	content, err := ioutil.ReadFile("./resume.json")
// 	if err != nil {
// 		log.Fatal("Error when opening file: ", err)
// 	}

// 	// Now let's unmarshall the data into `payload`
// 	var payload map[string]interface{}
// 	err = json.Unmarshal(content, &payload)
// 	if err != nil {
// 		log.Fatal("Error during Unmarshal(): ", err)
// 	}

// 	log.Print(payload["basics"])

// 	var tmpl = `
//     {{with .basics}}
//     hello {{.email}}
//     {{end}}`
// 	t2 := template.Must(template.New(tmpl).Parse(tmpl))

// 	err2 := t2.Execute(os.Stdout, payload)
// 	if err != nil {
// 		log.Println("executing template:", err2)
// 	}

// 	// Define a template.
// 	const letter = `
// ----
// Dear {{.Name}},
// {{if .Attended}}
// It was a pleasure to see you at the wedding.
// {{- else}}
// It is a shame you couldn't make it to the wedding.
// {{- end}}
// {{with .Gift -}}
// Thank you for the lovely {{.}}.
// {{end}}
// Best wishes,
// Josie
// `

// 	// Prepare some data to insert into the template.
// 	type Recipient struct {
// 		Name, Gift string
// 		Attended   bool
// 	}
// 	var recipients = []Recipient{
// 		{"Aunt Mildred", "bone china tea set", true},
// 		{"Uncle John", "moleskin pants", false},
// 		{"Cousin Rodney", "", false},
// 	}

// 	// Create a new template and parse the letter into it.
// 	t := template.Must(template.New("letter").Parse(letter))

// 	// create a new file
// 	file, _ := os.Create("greeting2.txt")
// 	defer file.Close()

// 	// Execute the template for each recipient.
// 	for _, r := range recipients {
// 		err := t.Execute(file, r)
// 		if err != nil {
// 			log.Println("executing template:", err)
// 		}
// 	}

// }
