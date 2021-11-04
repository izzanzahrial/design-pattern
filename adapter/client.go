package adapter

import "fmt"

type client struct {
}

func (c *client) insertLightningConnectorIntoComputer(pc computer) {
	fmt.Println("Client insert lighting connector into computer")
	pc.insertIntoLightningPort()
}
