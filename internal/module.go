package internal

import (
	"net"
	"sync"
)

const (
	enterNameTxtMsg  = "\n[ENTER YOUR NAME]:"
	joinTxtMsg       = " has joined our chat..."
	leftTxtMsg       = " has left our chat..."
	emptyTxtMsg      = "[Empty name is unavailable.Write your name]:"
	longTxtMsg       = "[TOO LONG NAME. ENTER YOU NAME]:"
	alredyUsedTxtMsg = "[That name is already used. Choose another one]:"
	sorryNoAddTxtMsg = "Sorry, there is no more free space in the chat"
)

type client struct {
	name string   // Хранит имя (никнейм) клиента
	addr string   // Хранит адрес локальной сети
	conn net.Conn // Хранит соединение
}

type message struct {
	name string // Хранит имя (никнейм) отправителя сообщения
	date string // Хранит методанные даты и времени сообщения
	text string // Хранит текст сообщения
}

var (
	clients   = make(map[net.Conn]client, 10) // Хеш-таблица для хранения данных об подключенных клиентах 
	mu        sync.Mutex           // Мьютекс для синхронизации горутин при использовании разделемых ресурсов
	join      = make(chan message) // Канал для передачи сообщения об подключении клиента
	messages  = make(chan message) // Канал для передачи сообщений
	leave     = make(chan message) // Канал для передачи сообщений об отключении клиента
	cahe      []string             // Срез для хранения истории сообщений из чата
	countConn int                  // Счетчик соединений
)
