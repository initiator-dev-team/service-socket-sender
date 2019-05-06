package socketSender

import (
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
	"os"
	"strconv"
	"time"
)

type directMessageData struct {
	socketId string `json:"id"`
	data     string `json:"data"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Println("FAIL:", msg, err)
	}
}

func initClient(domain string) (*gosocketio.Client, error) {
	url := os.Getenv("SOCKET_URL")
	port, _ := strconv.Atoi(os.Getenv("SOCKET_PORT"))
	token := os.Getenv("SOCKET_TOKEN")

	builtUrl := gosocketio.GetUrl(url, port, false)
	builtUrl += "&token=" + token + "&domain=" + domain

	client, err := gosocketio.Dial(
		builtUrl,
		transport.GetDefaultWebsocketTransport())

	failOnError(err, "Cannot connect to socket server")

	return client, err
}

func SendDirectMessage(domain, socketId, data string) {

	client, error := initClient(domain)

	if error != nil {
		return
	}
	_, err := client.Ack("DIRECT_MESSAGE", directMessageData{socketId, data}, time.Second*1)

	failOnError(err, "Dont believe, it was sent")
	log.Println("emitted", "method:", "DIRECT_MESSAGE", "socketId:", socketId, "data:", data)

	client.Close()
}

func SendToRoom(domain, eventName, data string) {
	client, error := initClient(domain)

	if error != nil {
		return
	}

	_, err := client.Ack(eventName, data, time.Second*1)

	failOnError(err, "Dont believe, it was sent")
	log.Println("emitted", "method:", eventName, "data:", data)

	client.Close()
}
