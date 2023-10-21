package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// On version update, we always start counting revisions from 1
var revision = "1"

var (
	semverTagRe      = regexp.MustCompile(`(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)?$`)
	revisionSuffixRe = regexp.MustCompile(`-\d+\s*$`)
)

// this script accepts release name as an argument, extracts semantically versioned tag from the release name
// and writes it to VERSION file in the root of the repository if the version is different
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		exit("argument <release name> is required!")
	}
	releaseName := args[0]
	if releaseName == "" {
		exit("release name is empty")
	}

	fmt.Printf("checking relese: '%s'\n", releaseName)

	latest := semverTagRe.FindString(releaseName)
	if latest == "" {
		exit("semantic version tag is not recognized")
	}

	content, err := os.ReadFile("VERSION")
	if err != nil {
		exit("file not found")
	}

	current := revisionSuffixRe.ReplaceAllString(string(content), "")
	if current != latest {
		tag := strings.Join([]string{latest, revision}, "-")
		err = os.WriteFile("VERSION", []byte(tag+"\n"), 0744)
		if err != nil {
			exit(err.Error())
		}
		fmt.Printf("version updated to %s\n", latest)
	}
	fmt.Println("already at the latest version")
}

// exit with a non-zero exit code and an error message
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
