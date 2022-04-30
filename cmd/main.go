package cmd

import (
	"fmt"
	"github.com/fades-io/reg/internal/apperror"
	"github.com/fades-io/reg/internal/logging"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fades-io/reg/internal"
	"github.com/joho/godotenv"
)

// We make signallHandler receive a channel on which we will report the value of var quit
// signallHandler - обработчик сигналов
func signallHandler(q chan bool) {
	var quit bool

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// Для каждого полученного сигнала
	for signal := range c {
		log.Printf(logging.SignalReceived, signal.String())

		switch signal {
		case syscall.SIGINT, syscall.SIGTERM:
			quit = true
		case syscall.SIGHUP:
			quit = false
		}

		if quit {
			quit = false
			//              closeDb()
			//              closeLog()
			log.Print(logging.Terminating)
			os.Exit(0)
		}
		// Оповещаяем о прекращении работы
		q <- quit
	}
}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf(apperror.FileAccessFailed, err)
	} else {
		fmt.Println(logging.FileAccessSuccess)
	}

	// Канал для сигналов
	sig := make(chan bool)
	// Основной канал
	loop := make(chan error)

	// Начанаем мониторинг сигналов
	go signallHandler(sig)

	// Пока не пришло оповещение о прекращении работы
	for quit := false; !quit; {
		go func() {
			internal.Run()
			loop <- nil
		}()

		// Блокируем программу при получении сигнала
		select {
		// прерываем выполнение программы
		case quit = <-sig:
		// Продолжаем выполнение программы
		case <-loop:
		}
	}
}
