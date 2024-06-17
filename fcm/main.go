package main

func main() {
	cli, err := newClient()
	if err != nil {
		return
	}

	sendMessage(cli)
}
