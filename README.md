# Favicon Fetcher

This is a simple script that fetches the favicon of a website and saves it to a file.

## Usage

```bash
# Fetch github.com favicon and save it to icon.png
fetch-favicon github.com > icon.png

# Alternatively, you can use the -o flag to specify the output file
fetch-favicon -o icon.png github.com

# You can set the size of the favicon with the -s flag
fetch-favicon -s 32 github.com > icon.png
```
