package adapter

func Execute() {
	client := &client{}
	mac := &mac{}

	client.insertLightningConnectorIntoComputer(mac)

	windowsMachine := &windows{}
	windwosMachineAdapter := &windwosAdapter{
		windowsMachine: windowsMachine,
	}

	client.insertLightningConnectorIntoComputer(windwosMachineAdapter)
}
