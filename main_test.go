package main

import "testing"
import "net"
import "bufio"
import "fmt"
import "io/ioutil"

func TestConnect(t *testing.T) {
	go StartServer()

	// se connecter à localhost:8000
	conn, err := net.Dial("tcp", "localhost:8000")

	// vérifier qu'on peut
	if err != nil {
		t.Errorf("Could not connect: %s", err)
	}

	// envoyer une requête HTTP
	_, err = fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")

	if err != nil {
		t.Errorf("Could not write: %s", err)
	}

	reader := bufio.NewReader(conn)

	status, err := reader.ReadString('\n')

	if err != nil {
		t.Errorf("Could not read status: %s", err)
	} else if status != "HTTP/1.1 200 OK\r\n" {
		t.Errorf("Invalid response: %s", status)
	}

	blankline, err := reader.ReadString('\n')

	if err != nil {
		t.Errorf("Could not read blank: %s", err)
	} else if blankline != "\r\n" {
		t.Errorf("Invalid response: %s", blankline)
	}

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Errorf("Could not read content: %s", err)
	} else if string(content) != "Hello!" {
		t.Errorf("Invalid response: %s", content)
	}
}
