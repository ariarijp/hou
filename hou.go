package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

func createParams(filename string) slack.PostMessageParameters {
	var params slack.PostMessageParameters
	params.Username = "hou"
	params.IconEmoji = ":bell:"
	params.Parse = "full"

	if filename == "" {
		return params
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return params
	}

	err = json.Unmarshal(bytes, &params)
	if err != nil {
		return params
	}

	return params
}

func handleError(err error, quiet bool) {
	if !quiet {
		log.Fatal(err)
	}

	os.Exit(1)
}

func main() {
	var channel, filename string
	var asCode, quiet, silent bool
	flag.StringVar(&channel, "channel", "", "Required: Which channel to send messages to")
	flag.StringVar(&filename, "params-file", "", "Optional: Override slack.PostMessageParameters with JSON file")
	flag.BoolVar(&asCode, "as-code", true, "Optional: Send message fenced by three backticks")
	flag.BoolVar(&quiet, "quiet", true, "Optional: Suppress error message")
	flag.BoolVar(&silent, "silent", false, "Optional: Suppress any output")
	flag.Parse()

	if silent {
		quiet = true
	}

	if channel == "" {
		handleError(errors.New("please set an argument: -channel"), quiet)
	}

	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		handleError(errors.New("please set an environment variable: SLACK_API_TOKEN"), quiet)
	}

	api := slack.New(token)

	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		handleError(err, quiet)
	}

	text := string(body)
	if !silent {
		fmt.Print(text)
	}

	text = strings.TrimSpace(text)
	if text == "" {
		os.Exit(0)
	}

	if asCode {
		text = fmt.Sprintf("```\n%s\n```", text)
	}

	params := createParams(filename)

	_, _, err = api.PostMessage(channel, text, params)
	if err != nil {
		handleError(err, quiet)
	}
}
