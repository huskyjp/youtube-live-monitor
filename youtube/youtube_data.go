package youtubestats

// Store response JSON from Youtube API
// Items are slice of other statistics iofo
type Response struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

type Items struct {
	Kind string `json:"kind"`
	Id   string `json:"id"`
	Data Data   `json:"statistics"`
}

// We stores actual youtube account info here
// views, subscribers, and how many video they have posted
type Data struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}
