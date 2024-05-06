package internal

import "fmt"

// Функция широковешательного распределения данных
func broadcast() {
	for {
		select {
		case msg := <-join:
			for _, client := range clients {
				welcome := fmt.Sprintf("[%s][%s]:", msg.date, client.name)
				if client.name != msg.name {
					client.conn.Write([]byte(msg.text + "\n"))
					client.conn.Write([]byte(welcome))
				}
			}
		case msg := <-messages:
			text := fmt.Sprintf("[%s][%s]:%s", msg.date, msg.name, msg.text)
			for _, client := range clients {
				metaDataClient := fmt.Sprintf("[%s][%s]:", msg.date, client.name)
				if client.name != msg.name {
					client.conn.Write([]byte("\r\n" + text + "\n"))
					client.conn.Write([]byte(metaDataClient))
				}
			}
			cahe = append(cahe, text+"\n")
		case msg := <-leave:
			for _, client := range clients {
				metaDataClient := fmt.Sprintf("[%s][%s]:", msg.date, client.name)
				if client.name != msg.name {
					client.conn.Write([]byte("\r\n" + msg.text + "\n"))
					client.conn.Write([]byte(metaDataClient))
				}
			}
		}
	}
}
