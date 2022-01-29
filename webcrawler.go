package main

import (
	"fmt"
	"log"
	"os"

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
	fmt.Printf("Destination: %s\n", directoryPath)
	fmt.Println("URL:", urlLink)

	//downloadPage(directoryPath, urlLink)
	getValidLinks("https://start.url/abc", "/Users/tonym/Desktop/Hackathon/Project_Plato/WgetExample/sites/start.html")
}

func init() {
	flag.StringVarP(&directoryPath, "dir", "d", "", "Destination directory for downloaded pages")
}

func downloadPage(destPath string, url string) {
	resp, err := grab.Get(destPath, url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Download saved to", resp.Filename)
}

func getValidLinks(startURL string, filePath string) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(dat))
}
