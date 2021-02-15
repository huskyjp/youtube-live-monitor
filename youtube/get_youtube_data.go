package youtubestats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GetYoutubeData requests YouTube API to get basic information depends users' YoutubeKey & ChannelID
func GetYoutubeData() (Items, error) {
	var response Response
	// Construct a new request with GET method to the youtube api
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		log.Fatalln(err)
		return Items{}, err
	}

	// construct query so we can access to each individual youtube account data
	q := req.URL.Query()
	// os.Gentenv allows us to manually insert the data from console
	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	q.Add("id", os.Getenv("CHANNEL_ID"))
	q.Add("part", "statistics") // want statistics part @json
	// encrypt query data and insert back to the header
	req.URL.RawQuery = q.Encode()

	// make a request and store response data to the client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return Items{}, err
	}
	defer resp.Body.Close()

	// get respone data (data is stored in the body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Items{}, err
	}
	// unmarshal current response whch is stored in body, and store it to address of response
	err = json.Unmarshal(body, &response) // json => struct
	if err != nil {
		return Items{}, err
	}
	fmt.Println("Current Items", response.Items[0])
	return response.Items[0], nil
}
