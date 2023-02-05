package worker

import (
	"fmt"

	"github.com/19chonm/461_1_23/fileio"
)

func runTask(url string, ratingch chan<- fileio.Rating) {
	fmt.Println("My job is", url)
	r := fileio.Rating{NetScore: 5.0, Url: "url_for_module_A"}
	ratingch <- r
}
