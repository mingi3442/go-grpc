package log

import (
  "fmt"
  "time"
)

// LogLevel type
type LogLevel int

// Log levels
const (
  DEBUG LogLevel = iota
  INFO
  WARN
  ERROR
)

// // String returns the string representation of a LogLevel(0 => "DEBUG" etc.)
func (l LogLevel) String() string {
  switch l {
  case DEBUG:
    return "\033[34mDEBUG\033[0m" // 파란색
  case INFO:
    return "\033[32mINFO\033[0m" // 녹색
  case WARN:
    return "\033[33mWARN\033[0m" // 노란색
  case ERROR:
    return "\033[31mERROR\033[0m" // 빨간색
  default:
    return "\033[37mUNKNOWN\033[0m" // 흰색
  }
}
func (l LogLevel) Color() string {
  switch l {
  case DEBUG:
    return "\033[34m" // 파란색
  case INFO:
    return "\033[32m" // 녹색
  case WARN:
    return "\033[33m" // 노란색
  case ERROR:
    return "\033[31m" // 빨간색
  default:
    return "\033[37m" // 흰색
  }
}

// LogMessage struct
type LogMessage struct {
  Level   LogLevel
  Message string
}

// logChan is a channel for log messages
var logChan chan LogMessage

// package가 읽혀질 때 최초 시작
func init() {
  logChan = make(chan LogMessage)
  go logPrint(logChan)
}

func logPrint(logChan <-chan LogMessage) {
  for logMsg := range logChan {
    colorCode := logMsg.Level.Color()
    // 로그 레벨과 메시지 전체에 색상을 적용하고, 메시지 출력 후 색상을 리셋합니다.
    fmt.Printf("%s[%s] %s: %s%s\n", colorCode, logMsg.Level.String(), time.Now().Format(time.RFC3339), logMsg.Message, "\033[0m")
  }
}

func Log(level LogLevel, msg string) {
  logChan <- LogMessage{Level: level, Message: msg}
}

//-------------------------------------------------------------

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
