package socketSender

import (
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
	"strconv"
	"time"
)

var (
	_socketUrl, _socketPort, _socketToken string
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
	url := _socketUrl
	port, _ := strconv.Atoi(_socketPort)
	token := _socketToken

	builtUrl := gosocketio.GetUrl(url, port, false)
	builtUrl += "&token=" + token + "&domain=" + domain

	client, err := gosocketio.Dial(
		builtUrl,
		transport.GetDefaultWebsocketTransport())

	failOnError(err, "Cannot connect to socket server")

	return client, err
}

func Init(socketUrl, socketPort, socketToken string) {
	_socketPort = socketPort
	_socketToken = socketToken
	_socketUrl = socketUrl
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
