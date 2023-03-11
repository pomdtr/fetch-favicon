package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fetch-favicon <domain_url>",
	Short: "Fetches the favicon for a given domain.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		var output io.Writer

		if cmd.Flags().Changed("output") {
			outputPath, _ := cmd.Flags().GetString("output")
			if outputPath == "-" {
				output = os.Stdout
			} else {
				file, err := os.Create(outputPath)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error:", err)
					os.Exit(1)
				}
				defer file.Close()

				output = file
			}
		} else {
			if isatty.IsTerminal(os.Stdout.Fd()) {
				fmt.Fprintln(os.Stderr, "Error: output is a terminal. Use -o to specify an output file.")
				os.Exit(1)
			}

			output = os.Stdout
		}

		err = fetchFavicon(url, size, output)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().Int("size", 16, "size of the favicon")
	rootCmd.Flags().StringP("output", "o", "", "output file")
}

func fetchFavicon(domain_url string, size int, output io.Writer) error {
	faviconUrl := url.URL{
		Scheme: "https",
		Host:   "www.google.com",
		Path:   "/s2/favicons",
	}

	query := faviconUrl.Query()
	query.Set("sz", strconv.Itoa(size))
	query.Set("domain_url", url.QueryEscape(domain_url))
	faviconUrl.RawQuery = query.Encode()

	log.Printf("Fetching favicon from %s", faviconUrl.String())
	request, err := http.NewRequest("GET", faviconUrl.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to fetch favicon: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch favicon: %s", response.Status)
	}

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
