package logger

import "log"

func LogI(ctx ...interface{}) {
	log.Println("[INFO]" , ctx)
}

func LogE(ctx ...interface{}) {
	log.Fatalln("[ERROR]", ctx)
}

func LogW(ctx ...interface{})  {
	log.Println("[WRONG]", ctx)
}