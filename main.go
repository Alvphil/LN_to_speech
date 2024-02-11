package main

import (
	"fmt"
	"log"

	"github.com/Alvphil/LN_to_speech/internal/websites"
)

func main() {
	content, err := websites.GetChapterWI("https://wanderinginn.com/2022/06/18/interlude-singing-ships/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
}
