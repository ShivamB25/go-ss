# Go Website Screenshotter

A simple CLI tool to take screenshots of a list of websites from a JSON file using `gowitness`.

## Prerequisites

You must have `gowitness` installed on your system. If you haven't installed it, follow the instructions on the [official gowitness repository](https://github.com/sensepost/gowitness).

## Installation

You can install the screenshotter CLI using `go install`:

```bash
go install github.com/shivamb25/go-ss/cmd/screenshotter
```

Replace `your-username` with your actual GitHub username. This will install the `screenshotter` binary in your `$GOPATH/bin` directory.

## Usage

1.  **Prepare your links**: Create a JSON file containing an array of URLs. For example, `mylinks.json`:

    ```json
    [
      "https://www.google.com",
      "https://github.com",
      "https://sensepost.com"
    ]
    ```

2.  **Run the tool**:

    Once installed, you can run the `screenshotter` command from anywhere in your terminal.

    To take screenshots of URLs in `mylinks.json` and save them to the `website-screenshots` directory (the default behavior):

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

    You can also combine flags:

    ```bash
    screenshotter --file mylinks.json --output my-screenshots
    ```

    Screenshots will be saved in the specified output directory with filenames generated from the URLs.

## Development

To run the tool from the source code without installing:

```bash
go run ./cmd/screenshotter --file mylinks.json