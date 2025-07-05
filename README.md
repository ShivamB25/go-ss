# Go Website Screenshotter

A simple, self-contained CLI tool to take screenshots of a list of websites concurrently using a headless browser. This tool uses `playwright-go` to control a headless Chromium instance.

## Prerequisites

- Go 1.16 or later.

That's it! The tool will automatically download the necessary browser binaries on the first run.

## Installation

You can install the screenshotter CLI using `go install`:

```bash
go install github.com/your-username/go-ss/cmd/screenshotter
```

Replace `your-username` with your actual GitHub username. This will install the `screenshotter` binary in your `$GOPATH/bin` directory.

## Usage

1.  **Prepare your links**: Create a JSON file containing an array of URLs. For example, `mylinks.json`:

    ```json
    [
      "https://www.google.com",
      "https://github.com",
      "https://www.bing.com"
    ]
    ```

2.  **Run the tool**:

    Once installed, you can run the `screenshotter` command from anywhere in your terminal.

    To take screenshots of URLs in `mylinks.json` and save them to the `website-screenshots` directory:

    ```bash
    screenshotter
    ```

    To specify a different input file:

    ```bash
    screenshotter --file links.json
    ```

    To specify a different output directory:

    ```bash
    screenshotter --output my-screenshots
    ```

    To control the number of concurrent browser instances:

    ```bash
    screenshotter --concurrency 5
    ```

    You can also combine flags:

    ```bash
    screenshotter --file mylinks.json --output my-screenshots --concurrency 20
    ```

    Screenshots will be saved in the specified output directory with filenames generated from the URLs.

## Development

To run the tool from the source code without installing:

```bash
go run ./cmd/screenshotter --file mylinks.json