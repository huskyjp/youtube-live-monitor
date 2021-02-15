# YouTube Live Data Monitor
Golang based live Youtube Data (Subscribers, Views, Videos) Monitor with simple HTML client.
Use own YouTube Channel_ID & YouTube_APIKey to access the actual data.

Topic includes:
- Websocket
- REST API
- Gorilla Web Toolkit
- Basic Go Routine

# Run with Golang
```
export CHANNEL_ID='CHANNEL_ID'
export YOUTUBE_KEY='API_KEY'
```
which enables you to get a personal YouTube data.
- 'CHANNEL_ID' => You can get it from [here](https://www.youtube.com/account_advanced)
- 'API_KEY`    => You can get it from [here](https://developers.google.com/youtube/v3/getting-started)

```
go get ["github.com/gorilla/websocket"](https://github.com/gorilla/websocket)
go run main.go
```
which will kick of a server on http://localhost:8080 and display your personal YouTube Live Data every 3 seconds!
You need to import Gorilla Web Toolkit package.