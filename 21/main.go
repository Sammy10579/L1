package main

//Реализовать паттерн «адаптер» на любом примере.

func main() {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}

/*type kilograms int

func toGrams(g int) kilograms {
	return kilograms(g * 1000)
}

func kilo(gr kilograms) {
	fmt.Printf("grams - %d", gr)
}

func main() {
	kg := 3
	kilo(toGrams(kg))
}*/
