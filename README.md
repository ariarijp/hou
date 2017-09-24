# hou

Send STDIN as a Slack message.

Inspired by [Songumu/horenso](https://github.com/Songmu/horenso).

## Installation

```bash
$ go get -u github.com/ariarijp/hou
```

## Usage

```bash
$ export SLACK_API_TOKEN=YOUR_SLACK_API_TOKEN
$ w | hou -channel hello
```

Then, you will receive a message in Slack like below.

![screenshot.png](screenshot.png)

## License

MIT

## Author

[ariarijp](https://github.com/ariarijp)
