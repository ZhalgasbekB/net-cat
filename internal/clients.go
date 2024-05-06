package internal

import (
	"bufio"
	"net"
	"strings"
	"time"
)

// Функция обработки логики подключения клиента
func newClient(conn net.Conn, welcome string) {
	defer conn.Close()
	// Передаем новому пользователю (клиенту) приветственное сообщение и просим ввести никнейм
	conn.Write([]byte(welcome + enterNameTxtMsg))

	// Проверяем на валидность введенного пользователем имени (никнейма)
	nameUser := checkName(clients, conn)

	// Вносим параметры нового клиента как объекта с определенными
	// параметрами
	client := client{nameUser, conn.LocalAddr().String(), conn}

	mu.Lock()
	// Вносим данные о новом клиенте как об объекте в хеш-таблицу клиентов
	clients[conn] = client
	mu.Unlock()

	// Передаем в канал сообщение об подключении нового клиента для широковещания
	join <- message{client.name, time.Now().Format("2006-01-02 15:04:05"), "\r\n" + client.name + joinTxtMsg}

	mu.Lock()
	// Проверяем кэш сообщение из чата
	// при наличии истории сообщений подгружаем теущему
	// новому клиенту
	if len(cache) != 0 {
		for _, v := range cache {
			conn.Write([]byte(v))
		}
	}
	mu.Unlock()

	// Регистрируем для соединения сканнер получающий ввод текста сообщений
	inputUser := bufio.NewScanner(conn)
	// Передаем текущему клиенту метаданные
	conn.Write([]byte(metadata(client)))
	// Читаем в цикле из стандартного ввода терминала клиента
	// ввод сообщения
	for inputUser.Scan() {
		// Возвращаем последний прочитанный токен (в нашем случае построчно текст сообщения)
		// из стандратного ввода терминала пользователя
		textMsg := inputUser.Text()
		if len(strings.Trim(textMsg, " \r\n")) == 0 {

			conn.Write([]byte(metadata(client)))
			continue
		}
		// Передаем в канал сообщение от клиента для широковещания
		messages <- message{client.name, time.Now().Format("2006-01-02 15:04:05"), textMsg}
		conn.Write([]byte(metadata(client)))
	}

	// Формируем сообщение об отключении клиента
	leaveText := client.name + leftTxtMsg
	// Вычитаем из счетчика соединение
	countConn--

	mu.Lock()
	// Удаляем данные клиента из хеш-таблицы
	delete(clients, conn)
	mu.Unlock()
	// Передаем в канал сообщение об отключении клиента для широковещания
	leave <- message{client.name, time.Now().Format("2006-01-02 15:04:05"), leaveText}
}

// Функция проверки введенного пользователем имени на валидность
func checkName(clients map[net.Conn]client, conn net.Conn) string {
	// Регистрируем для соединения сканнер получающий ввод текста имени (никнейма) клиента
	compaund := bufio.NewScanner(conn)

	// Читаем в цикле из стандартного ввода терминала клиента
	// ввод
	for compaund.Scan() {
		flag := true
		// Возвращае последний прочитанный токен (в нашем случае имя)
		// из стандратного ввода терминала пользователя
		name := compaund.Text()
		// Проверка на пустое имя
		if len(strings.Trim(name, " \r\n")) == 0 {
			conn.Write([]byte(emptyTxtMsg))
			continue
		}
		// Проверка размера строки с именем (никнеймом) больше 30 байт
		if len(name) > 30 {
			conn.Write([]byte(longTxtMsg))
			continue
		}
		// Проверка на уникальность имени (никнейма) в чате
		mu.Lock()
		for _, client := range clients {
			if client.name == name {
				conn.Write([]byte(alredyUsedTxtMsg))
				flag = false
				break
			}
		}
		mu.Unlock()
		if flag == false {
			continue
		} else {
			return name
		}
	}
	return ""
}
