package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/cavaliergopher/grab/v3"
	flag "github.com/ogier/pflag"
)

var (
	directoryPath string
)
var (
	urlLink string
)

func main() {
	flag.Parse()

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options] [arguments]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("Arguments:")
		fmt.Println("[URL]: starting URL")
		os.Exit(1)
	}

	urlLink = flag.Args()[0]

	downloadPage(directoryPath, urlLink)
}

func init() {
	flag.StringVarP(&directoryPath, "dir", "d", "", "Destination directory for downloaded pages")
}

func downloadPage(destPath string, url string) {
	resp, err := grab.Get(destPath, url)
	if err != nil {
		fmt.Println("Error occurred!")
		log.Fatal(err)
	}
	fmt.Println("Download saved to", resp.Filename)

	getValidLinks(url, resp.Filename)
}

func getValidLinks(startURL string, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	anchorMatcher := regexp.MustCompile(`<a[^>]* href="([^"]*)">`)
	for scanner.Scan() {
		text := scanner.Text()
		href := anchorMatcher.FindAllString(text, -1)
		if len(href) > 0 {
			getChildrenLink(startURL, href)
		}

	}

	file.Close()
}

func getChildrenLink(startURL string, href []string) {
	linkMatcher := regexp.MustCompile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
	for i := 0; i < len(href); i++ {
		link := linkMatcher.FindString(href[i])
		if strings.Contains(link, startURL) {
			fmt.Println(link)
		}

	}
}
