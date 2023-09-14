package models

import (
	"fmt"
	"sync"
)

type client struct {
	Api_Key      string
	Organization string
}

var lock = &sync.Mutex{}

var GptClient *client

func GetGPTClient(Api_key string, organization string) {
	if GptClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if GptClient == nil {
			fmt.Println("Creating single instance now")
			GptClient = &client{
				Api_Key:      Api_key,
				Organization: organization,
			}
		} else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}
}
