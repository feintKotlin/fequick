package logger

import "log"

func LogI(ctx ...interface{}) {
	log.Println("[INFO]" , ctx)
}
