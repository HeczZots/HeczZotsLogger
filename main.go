package main

import "mylogger/logger"

func main() {
	log.Logger = log.Output(logger.ConsoleWriter{Out: os.Stdout})
	l.Debug().Msg()
}
