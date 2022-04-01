package logger

//import (
//	"github.com/gin-gonic/gin"
//	"github.com/okikechinonso/internals/ports"
//	"io"
//	"log"
//	"os"
//	"strings"
//)
//
//type Logger struct {
//	logger *log.Logger
//}
//
//// NewLogger create an instance of the logger
//func NewLogger(l *log.Logger) ports.ILogger {
//	return &Logger{logger: l}
//}
//
//func (l *Logger) SetFormater() {
//	formatter := &log.TextFormatter{
//		TimestampFormat: "02-01-2006 15:04:05", // the "time" field configuration
//		DisableColors: false,
//	}
//
//	l.logger.SetReportCaller(true)
//	l.logger.SetFormatter(formatter)
//}
//
//// MakeLogger creates an instance of a logger
//func (l *Logger) MakeLogger(filename string, display bool) *log.Logger {
//	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	l.SetFormater()
//
//	if display {
//		l.logger.SetOutput(io.MultiWriter(os.Stdout, f))
//	} else {
//		l.logger.SetOutput(io.MultiWriter(f))
//	}
//	return l.logger
//
//}
//
////func (l *logger) Hook() *log.Logger {
////	client, err := elastic.NewClient(elastic.SetURL(config.Instance.ElasticURL))
////	if err != nil {
////		log.Panic(err)
////	}
////	hook, err := elogrus.NewAsyncElasticHook(client, "localhost", log.DebugLevel, "mylog")
////	if err != nil {
////		log.Panic(err)
////	}
////	l.SetFormater()
////	l.logger.Hooks.Add(hook)
////	return l.logger
////}
//
//func formatFilePath(path string) string {
//	arr := strings.Split(path, "/")
//	return arr[len(arr)-1]
//}
//
//type RequestMessage struct {
//	StatusCode  interface{} `json:"status_code"`
//	LatencyTime interface{} `json:"latency_time"`
//	ClientIp    interface{} `json:"client_ip"`
//	ReqMethod   interface{} `json:"req_method"`
//	ReqUri      interface{} `json:"req_uri"`
//}
//
//
//
//func(l *Logger) RequestLog (ctx *gin.Context) {
//	l.SetFormater()
//	log.WithFields(log.Fields{
//		"statusCode" : ctx.Writer.Status(),
//		"client_ip": ctx.ClientIP(),
//		"request_method": ctx.Request.Method,
//		"request_uri": ctx.Request.RequestURI,
//	}).Infof("")
//
//}
