package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	//restApi config
	BaseUrl       = "https://api.bitget.com"
	ApiKey        = ""
	SecretKey     = ""
	PASSPHRASE    = ""
	TimeoutSecond = 30

	//websocket config
	WsUrl = "wss://ws.bitget.com/spot/v1/stream"
)

func initConfig() {
	err := godotenv.Load("bitget.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	ApiKey = os.Getenv("BITGET_APIKEY")
	SecretKey = os.Getenv("BITGET_SECRETKEY")
	PASSPHRASE = os.Getenv("BITGET_PASSPHRASE")
}
