# Go Website Screenshotter

This tool takes screenshots of a list of websites provided in a JSON file using `gowitness`.

## Prerequisites

You must have `gowitness` installed on your system. If you haven't installed it, follow the instructions on the [official gowitness repository](https://github.com/sensepost/gowitness).

## Usage

1.  **Prepare your links**: Create a JSON file containing an array of URLs. For example, `mylinks.json`:

    ```json
    [
      "https://www.google.com",
      "https://github.com",
      "https://sensepost.com"
    ]
    ```

2.  **Build the program**:

    ```bash
    go build main.go
    ```

3.  **Run the program**:

    To take screenshots of URLs in `mylinks.json` and save them to the `website-screenshots` directory (the default behavior):

    ```bash
    ./main
    ```

    To specify a different input file:

    ```bash
    ./main --file links.json
    ```

    To specify a different output directory:

    ```bash
    ./main --output my-screenshots
    ```

    You can also combine flags:

    ```bash
    ./main --file mylinks.json --output my-screenshots
    ```

    Alternatively, you can run the program without building it using `go run`:
    ```bash
    go run main.go --file mylinks.json --output my-screenshots
    ```

    Screenshots will be saved in the specified output directory with filenames generated from the URLs.