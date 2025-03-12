# File Downloader

File Downloader is a Go application that reads URLs from an input file, downloads the content from those URLs, and stores the results on disk.

## Features

- Supports reading URLs from a CSV file.
- Concurrently downloads content from multiple URLs.
- Stores downloaded content on disk.
- Logs the progress and errors using Logrus.

## Prerequisites

- Go 1.16 or later

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Meghalam/file_downloader.git
    cd file_downloader
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

## Usage

1. Prepare an input file with URLs (e.g., `input.csv`).

2. Run the application:

    ```sh
    go run cmd/file_downloader/main.go /path/to/your/inputfile.csv
    ```

## Project Structure

- [main.go]: The main entry point of the application.
- [data_store]: Package responsible for storing downloaded content.
- [file_reader]: Package responsible for reading URLs from input files.
- [definitions]: Package containing application-wide definitions and constants.
- [utility]: Package containing utility functions.


## Configuration

- [def.MaxWorkers]: Maximum number of concurrent workers for downloading URLs.
- [def.ResultPath]: Path where the downloaded content will be stored.

## Logging

The application uses Logrus for logging. Logs include information about the progress of the file reading, URL processing, and any errors encountered.
