package log

import (
  "fmt"
  "time"
)

type LogLevel int

const (
  DEBUG LogLevel = iota
  INFO
  WARN
  ERROR
)

type LogMessage struct {
  Level   LogLevel
  Message string
}

var logChan chan LogMessage

func init() {
  logChan = make(chan LogMessage, 100)
  go logWorker(logChan)
}

func logWorker(logChan <-chan LogMessage) {
  for logMsg := range logChan {
    // 로그 메시지 처리 로직
    fmt.Printf("[%s] %s: %s\n", logMsg.Level, time.Now().Format(time.RFC3339), logMsg.Message)
  }
}

func Log(level LogLevel, msg string) {
  logChan <- LogMessage{Level: level, Message: msg}
}

// var eventQueue chan func()

// // package가 읽혀질 때 최초 시작
// func init() {
//   // Logger setup
//   output := zerolog.ConsoleWriter{
//     Out:        os.Stderr,
//     TimeFormat: time.RFC1123,
//   }

//   // log Output 설정
//   log.Logger = log.Output(output)

//   zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
//   zerolog.ErrorStackFieldName = "trace"

//   // channel 생성
//   eventQueue = make(chan func())

//   // For thread safe
//   go func() {
//     // channel에 event가 들어올 때 마다 반복적으로 go function으로 실행
//     for event := range eventQueue {
//       event()
//     }
//   }()
// }

// func enqueue(event func()) {
//   eventQueue <- event
// }

// // log 종류

// func Info(msg string) {
//   event := func() {
//     log.Info().Msg(msg)
//   }
//   enqueue(event)
// }

// func Error(err error) {
//   stack := string(debug.Stack())
//   event := func() {
//     log.Error().Err(err).Msg("\n" + stack)
//   }
//   enqueue(event)
// }

// func Debug(msg any) {
//   message := fmt.Sprint(msg)
//   event := func() {
//     log.Debug().Msg(message)
//   }
//   enqueue(event)
// }
