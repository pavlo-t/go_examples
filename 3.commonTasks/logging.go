package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// create custom logger
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// create a logger that writes to a buffer, it can be any io.Writer
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello")
	fmt.Print("from buflog:", buf.String())

	// JSON logger
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	myslog.Info("hello again", "key", "val", "age", 25)
	userLog := myslog.WithGroup("user")
	userLog.Info("grouped info", "name", "John Doe")

	jamesLog := myslog.With("user", struct {
		Name  string `json:"name"`
		phone string
	}{"James", "1234567890"})
	jamesLog.Debug("debug message", "someKey", "someVal")
	jamesLog.Info("info message", "someKey", "someVal")
	jamesLog.Warn("warn message", "someKey", "someVal")
	jamesLog.Error("error message", "someKey", "someVal")
}
