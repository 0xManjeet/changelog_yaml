package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Change struct {
	Added   []string `yaml:"added"`
	Fixed   []string `yaml:"fixed"`
	Removed []string `yaml:"removed"`
}

type Version struct {
	Version string `yaml:"version"`
	Build   string `yaml:"build"`
	Date    string `yaml:"date"`
	Changes Change `yaml:"changes"`
}

type Changelog struct {
	VersioningFormat string    `yaml:"versioning_format"`
	Versions         []Version `yaml:"versions"`
}

func main() {
	yamlFile, err := os.ReadFile("changelog.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var changelog Changelog

	err = yaml.Unmarshal(yamlFile, &changelog)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	file, err := os.Create("changelog.md")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "# Changelog\n\nversioning format: `%s`\n\n", changelog.VersioningFormat)
	sort.Slice(changelog.Versions, func(i, j int) bool {
		iInt, _ := strconv.Atoi(changelog.Versions[i].Build)
		jInt, _ := strconv.Atoi(changelog.Versions[j].Build)
		return iInt > jInt
	})

	for _, version := range changelog.Versions {

		fmt.Fprintf(file, "## v%s+%s - %s\n\n", version.Version, version.Build, version.Date)
		if len(version.Changes.Added) > 0 {
			fmt.Fprintln(file, "### Added")
			fmt.Fprintln(file)
			for _, added := range version.Changes.Added {
				fmt.Fprintf(file, "- %s\n", added)
			}
			fmt.Fprintln(file)
		}

		if len(version.Changes.Fixed) > 0 {
			fmt.Fprintln(file, "### Fixed")
			fmt.Fprintln(file)
			for _, fixed := range version.Changes.Fixed {
				fmt.Fprintf(file, "- %s\n", fixed)
			}
			fmt.Fprintln(file)
		}
		// print new line
		fmt.Fprintln(file)
	}
	// now spit out the latest build on console
	fmt.Println(changelog.Versions[0].Build)
}
