package adapter

import (
	"fmt"
)

type Client struct{}

func NewClient() Client {
	return Client{}
}

func (c *Client) InsertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
