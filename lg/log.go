package lg

import (
	"log"
	"runtime"
	"strconv"
	"time"
)

var lev = 0

// I debug log
var I = printer{level: "info", color: 36, pri: 0}

// E warning log
var E = printer{level: "error", color: 33, pri: 1}

// F fatal log. does not panic
var F = printer{level: "fatal", color: 31, pri: 2}

// SetLevel ログレベルを設定
func SetLevel(l string) {
	switch l {
	case "info":
		lev = 0
	case "error":
		lev = 1
	case "fatal":
		lev = 2
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
	return decorateColor("["+p.level+"]", p.color) + " " + decorateColor(timestamp(), 34) + " " + caller(3) + "\n\t"
}

// Print Print
func (p *printer) Print(v ...interface{}) {
	if p.pri >= lev {
		v = append([]interface{}{p.prefix()}, v...)
		log.Print(v...)
	}
}

// Printf Printf
func (p *printer) Printf(format string, v ...interface{}) {
	if p.pri >= lev {
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
	return "\x1b[" + strconv.Itoa(color) + "m" + str + "\x1b[0m"
}

func caller(depth int) string {
	pc, file, line, _ := runtime.Caller(depth)
	f := runtime.FuncForPC(pc)
	return f.Name() + "() " + file + ":" + strconv.Itoa(line) + ""
}

func timestamp() string {
	return time.Now().Format("[2006-01-02 15:04:05]")
}

func init() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}
