package redirector

import (
	"fmt"
	"net/http"
	"time"
)

//Redirector returns all links with redirect
func Redirector(myURL string) map[string]int {
	redirs := make(map[string]int)
	timeout := time.Duration(5 * time.Second)
	nextURL := myURL
	var i int
	for i < 100 {
		redirs[nextURL] = i
		client := &http.Client{
			Timeout: timeout,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse

			}}

		resp, err := client.Get(nextURL)

		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		//fmt.Println("StatusCode:", resp.StatusCode)
		//fmt.Println(resp.Request.URL)

		if resp.StatusCode == 200 {
			fmt.Println("Done!")
			break
		} else {
			nextURL = resp.Header.Get("Location")
			i++
		}
	}
	return (redirs)
}
