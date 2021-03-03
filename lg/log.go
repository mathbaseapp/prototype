package lg

import (
	"log"
	"runtime"
	"strconv"
)

var pri = 0

// D debug log
var D = printer{level: "debug", color: 36, pri: 0}

// W warning log
var W = printer{level: "warinig", color: 33, pri: 1}

// F fatal log. does not panic
var F = printer{level: "fatal", color: 31, pri: 2}

// SetPri ログレベルを設定
func SetPri(p string) {
	switch p {
	case "info":
		pri = 0
	case "error":
		pri = 1
	case "fatal":
		pri = 2
	default:
		F.Println("log level not defined.")
		panic("")
	}
}

type printer struct {
	level string
	color int // 34:blue, 35:magenta, 33:yellow, 32:green, 36:cyan, 31:red
	pri   int
}

func (p *printer) prefix() string {
	return decorateColor(p.level, p.color) + " " + caller(3)
}

// Print Print
func (p *printer) Print(v ...interface{}) {
	if p.pri >= pri {
		v = append([]interface{}{p.prefix()}, v...)
		log.Print(v...)
	}
}

// Printf Printf
func (p *printer) Printf(format string, v ...interface{}) {
	if p.pri >= pri {
		v = append([]interface{}{p.prefix()}, v...)
		log.Printf(format, v...)
	}
}

// Println Println
func (p *printer) Println(v ...interface{}) {
	v = append([]interface{}{p.prefix()}, v...)
	log.Println(v...)
}

func decorateColor(str string, color int) string {
	return "\x1b[" + strconv.Itoa(color) + "m[" + str + "]\x1b[0m"
}

func caller(depth int) string {
	pc, file, line, _ := runtime.Caller(depth)
	f := runtime.FuncForPC(pc)
	return "@" + f.Name() + " (" + file + ":" + strconv.Itoa(line) + ")"
}
