package internal

import (
	"fmt"
	"log"
	"net"
	"os"
)

func StartServer(port string) {
	// Читаем файл с приветствием
	welcome, err := os.ReadFile("static/welcome.txt")
	if err != nil {
		log.Println("File reading error!")
		return
	}

	// Запуск сервера для прослушивания и принятия входящих запросов
	// по протоколу TCP на локальных хост с указанным портом
	l, err := net.Listen("tcp", ":"+port)

	// Отложенный вызов для закрытия подключения
	defer l.Close()

	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	fmt.Printf("Listening on the port: %s\n", port)

	// Вызов горутины для широковещательной рассылки данных
	// между участниками чата с помощью каналов в фоновом
	// режиме
	go broadcast()
	
	// Получаем входящие подключения
	for {
		// Получаем объект ввиде подключенного клиента
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Unable to connect")
			return
		}
		// Счетчик для подсчета соединений на указанный порт
		if conn != nil {
			countConn++
		}
		// Проверяем что в данный момент в чате не более 10 клиентов
		switch {
		case countConn < 11:
			// Запускаем в отдельной горутине логику обработки
			// каждого подключенного клиента
			go newClient(conn, string(welcome))
		default:
			conn.Write([]byte(sorryNoAddTxtMsg))
			countConn-- 
			conn.Close()
		}
	}
}
