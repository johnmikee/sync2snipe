package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	snipeBaseURL         = "https://snipe-it.readme.io"
	jamfClassicSchemaURL = "https://raw.githubusercontent.com/jamf/Classic-API-Swagger-Specification/master/swagger.yaml"
)

func main() {
	schemaFlag := flag.String("schema", "", "Specify a schema version to download (e.g., 6.0.0)")
	jamfFlag := flag.Bool("jamf", false, "Download the Jamf schema")
	flag.Parse()

	if *schemaFlag != "" {
		downloadRequestedSchema(*schemaFlag)
	} else if *jamfFlag {
		downloadJamfSchema()
	} else {
		interactiveSelection()
	}
}

func downloadJamfSchema() {
	err := downloadSchema(jamfClassicSchemaURL, "jamf_swagger.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Jamf schema downloaded successfully.")
}

func interactiveSelection() {
	resp, err := http.Get("https://snipe-it.readme.io/openapi/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	htmlContent, err := extractHTMLContent(resp)
	if err != nil {
		log.Fatal(err)
	}

	options := extractOptions(htmlContent)

	fmt.Println("Available Versions:")
	for i, option := range options {
		fmt.Printf("%d: Version %s\n", i+1, option.version)
	}

	fmt.Print("Please select an option 1-", len(options), ": ")
	var choice int
	_, err = fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(options) {
		log.Fatal("Invalid choice")
	}

	selectedOption := options[choice-1]
	fmt.Printf("Selected Version: %s\n", selectedOption.version)
	schemaURL := snipeBaseURL + selectedOption.href

	err = downloadSchema(schemaURL, "snipe-"+selectedOption.version+".json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Schema downloaded successfully.")
}

type schemaOption struct {
	version string
	href    string
}

func extractHTMLContent(resp *http.Response) (string, error) {
	var contentBuilder strings.Builder
	buffer := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		contentBuilder.Write(buffer[:n])
	}
	return contentBuilder.String(), nil
}

func extractOptions(htmlContent string) []schemaOption {
	options := make([]schemaOption, 0)

	splitDivs := strings.Split(htmlContent, "<div>")
	for _, div := range splitDivs {
		if strings.Contains(div, "<a href") && strings.Contains(div, "</a>") {
			option := extractOption(div)
			options = append(options, option)
		}
	}

	return options
}

func extractOption(divContent string) schemaOption {
	startIdx := strings.Index(divContent, "<a href=\"") + len("<a href=\"")
	endIdx := strings.Index(divContent, "\">")
	href := divContent[startIdx:endIdx]

	linkTextStartIdx := endIdx + len("\">")
	linkTextEndIdx := strings.Index(divContent, "</a>")
	linkText := divContent[linkTextStartIdx:linkTextEndIdx]

	versionStartIdx := strings.LastIndex(linkText, " - v") + len(" - v")
	version := linkText[versionStartIdx:]

	return schemaOption{
		version: version,
		href:    href,
	}
}

func downloadRequestedSchema(version string) {
	resp, err := http.Get("https://snipe-it.readme.io/openapi/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	htmlContent, err := extractHTMLContent(resp)
	if err != nil {
		log.Fatal(err)
	}

	options := extractOptions(htmlContent)

	var schemaURL string
	for _, option := range options {
		if option.version == version {
			baseURL := "https://snipe-it.readme.io"
			schemaURL = baseURL + option.href
			break
		}
	}

	if schemaURL == "" {
		log.Fatalf("Schema version %s not found", version)
	}

	err = downloadSchema(schemaURL, "snipe-"+version+".json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Schema version %s downloaded successfully.\n", version)
}

func downloadSchema(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
