package main

import (
	"encoding/json"
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

func main() {
	var channel, filename string
	var asCode bool
	flag.StringVar(&channel, "channel", "", "Required: Which channel to send messages to")
	flag.StringVar(&filename, "params-file", "", "Optional: Override slack.PostMessageParameters with JSON file")
	flag.BoolVar(&asCode, "as-code", true, "Optional: Send message fenced by three backticks")
	flag.Parse()

	if channel == "" {
		log.Fatal("Please set an argument: -channel")
	}

	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		log.Fatal("Please set an environment variable: SLACK_API_TOKEN")
	}

	api := slack.New(token)

	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	text := strings.TrimSpace(string(body))
	if asCode {
		text = fmt.Sprintf("```\n%s\n```", text)
	}

	params := createParams(filename)

	api.PostMessage(channel, text, params)
}
