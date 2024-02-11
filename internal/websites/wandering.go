package websites

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetChapterWI(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := extractEntryContent(resp)
	//bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)

	return bodyText, nil
}

// Extract the chapter from the response from the site.
func extractEntryContent(res *http.Response) (string, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var content string
	// Find the entry-content div and extract its content
	doc.Find(".entry-content").Each(func(i int, s *goquery.Selection) {
		// You can use .Html() to get the HTML content, or .Text() to get the plain text content
		divHtml, err := s.Html()
		if err != nil {
			fmt.Println("Error getting HTML content:", err)
		}
		content = divHtml
	})

	return content, nil

}
