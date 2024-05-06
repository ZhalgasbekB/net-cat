package internal

import (
	"fmt"
	"time"
)

// Функция получения методанных для сообщения ввиде текущей даты, времени и имени (никнейма) пользователя
func metadata(client client) string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("[%s][%s]:", currentTime, client.name)
	return msg
}
