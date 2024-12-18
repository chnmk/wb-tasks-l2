package CurrentTime

import (
	"log"
	"os"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===
Создать программу печатающую точное время с использованием NTP-библиотеки.
Инициализировать как go module.
Использовать библиотеку github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Требования
    Программа должна быть оформлена как go module.
    Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS.
*/

func GetCurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		// По условию задания.
		// Это значение по умолчанию для пакета log, но если оно было перезаписано, вернём STDERR.
		log.SetOutput(os.Stderr)
		log.Fatal(err)
	}

	log.Print(time)
}
