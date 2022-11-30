package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type Telnet struct { // структура с полученными данными
	host    string
	port    string
	timeout time.Duration
}

func connect() *Telnet {
	timeoutString := flag.String("timeout", "10s", "max timeout")

	flag.Parse() // парсим флаги
	seconds, err := strconv.Atoi(strings.TrimRight(*timeoutString, "s"))
	if err != nil {
		panic("invalid params")
	}
	timeout := time.Duration(seconds) * time.Second

	args := flag.Args()
	if len(args) < 2 {
		panic("invalid params")
	}
	a := &Telnet{
		host:    args[0],
		port:    args[1],
		timeout: timeout,
	}
	return a
}

func send(conn net.Conn, sigChan chan os.Signal, newQuery chan struct{}) { // читаем из консоли
	for {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			sigChan <- syscall.SIGQUIT
		}
		newQuery <- struct{}{}
	}
}

func read(conn net.Conn, sigChan chan os.Signal, newQuery chan struct{}) {
	for _ = range newQuery {
		fmt.Println("new message")
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			sigChan <- syscall.SIGQUIT
			return
		}
	}
}

func main() {
	a := connect()
	newQuery := make(chan struct{}, 1)
	sigChan := make(chan os.Signal, 1)                      // канал для отправки сигналов о завершении работы клиента
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // отлавливаем сигналы о завершении работы
	address := net.JoinHostPort(a.host, a.port)             // объединяем хост и порт
	conn, err := net.DialTimeout("tcp", address, a.timeout) // подключаемся к сети
	if err != nil {
		log.Fatal(err)
	}

	go send(conn, sigChan, newQuery)
	go read(conn, sigChan, newQuery)

	<-sigChan
	fmt.Println("Telnet client is closed")
	conn.Close()
}
