# Favicon Fetcher

This is a simple script that fetches the favicon of a website and saves it to a file.

> **Note** This cli uses an [hidden Google API](https://dev.to/derlin/get-favicons-from-any-website-using-a-hidden-google-api-3p1e) to fetch the favicon.

## Installation

```bash
go install github.com/pomdtr/fetch-favicon@latest
```

## Usage

```bash
# Fetch github.com favicon and save it to icon.png
fetch-favicon github.com > icon.png

# Alternatively, you can use the -o flag to specify the output file
fetch-favicon -o icon.png github.com

# You can set the size of the favicon with the -s flag
fetch-favicon -s 32 github.com > icon.png
```
