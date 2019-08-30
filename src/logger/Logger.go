package logger

import (
	"io"
	"log"
)

var (
	/*Trace is getting used globally*/
	Trace *log.Logger
	/*Info is getting used globally*/
	Info *log.Logger
	/*Warning is getting used globally*/
	Warning *log.Logger
	/*Error is getting used globally*/
	Error *log.Logger
)

/*Init is for intiate the logging*/
func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"LotterySystem:TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"LotterySystem:INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"LotterySystem:WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"LotterySystem:ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
