package logging

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"user-service-client/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Service interface {
	GetLogger() *zap.SugaredLogger

	DebugWithContext(ctx context.Context, args ...interface{})
	DebugWithContextf(ctx context.Context, format string, args ...interface{})

	ErrorWithContext(ctx context.Context, args ...interface{})
	ErrorWithContextf(ctx context.Context, format string, args ...interface{})

	InfoWithContext(ctx context.Context, args ...interface{})
	InfoWithContextf(ctx context.Context, format string, args ...interface{})
}

// StandardLogger initializes the standard logger
type standardLogger struct {
	logger *zap.SugaredLogger
}

type bufwriter chan []byte

func (bw bufwriter) Write(p []byte) (int, error) {
	pCopy := make([]byte, len(p))
	copy(pCopy, p)
	bw <- pCopy
	return len(p), nil
}

func NewBufwriter(n int, logFile string) bufwriter {
	w := make(bufwriter, n)
	logwriter := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10, // megabytes
		MaxBackups: 50,
		// MaxAge:     1, //days
	}

	go func(l *lumberjack.Logger, c bufwriter) {
		for p := range c {
			os.Stdout.Write(p)
			l.Write(p)
		}
	}(logwriter, w)

	return w
}

// NewService initializes the standard logger
func NewService(filename string) *standardLogger {
	atom := zap.NewAtomicLevel()
	atom.SetLevel(zap.DebugLevel) // level has been set

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
		if err != nil { // unable to open file, logging to stdout
			fmt.Println("errrrr #1", err)
		}
	}

	out, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_EXCL, 0666)
	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "file exists") {
			_, err = os.Stat(filename)
			if err != nil {
				panic(err)
			}
			err = os.Rename(filename, fmt.Sprintf(".logs/logger-%v.log", time.Now().Format("2006-01-02T15-04-05")))
			if err != nil {
				log.Println(err)
				panic(err)
			}
			out, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	out.Close()     // was there just to check
	if err != nil { // unable to open file, logging to stdout
		cfg := zap.Config{
			Encoding:         "json",
			Level:            atom,
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey: "message",

				LevelKey:    "level",
				EncodeLevel: zapcore.CapitalLevelEncoder,

				TimeKey:    "time",
				EncodeTime: zapcore.ISO8601TimeEncoder,

				// Commented as we manually add caller
				// CallerKey:    "caller",
				// EncodeCaller: zapcore.FullCallerEncoder,

				LineEnding: "\n",
			},
		}
		logger, nerr := cfg.Build() // error shouldn't happen but still I am handling it
		var sugar *zap.SugaredLogger
		if nerr != nil {
			sugar = zap.NewExample().Sugar()
			sugar.Error("Was unable to create logger file!")
			sugar.Error("Was unable to create desired logger, running on simple logger!")
		} else {
			sugar = logger.Sugar()
			sugar.Error("Was unable to create logger file!")
		}
		sugar = sugar.WithOptions(zap.AddCallerSkip(1))
		defer sugar.Sync()
		return &standardLogger{sugar}
	} else {
		cfg := zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			// Commented as we manually add caller
			CallerKey:    "file",
			EncodeCaller: zapcore.FullCallerEncoder,

			// FunctionKey: "func",
			LineEnding: "\n",
		}
		w := zapcore.AddSync(NewBufwriter(10000, filename))
		// mw := io.MultiWriter(os.Stdout, l)
		core := zapcore.NewCore(zapcore.NewJSONEncoder(cfg),
			w,
			atom,
		)
		logger := zap.New(core, zap.AddCaller())
		sugar := logger.Sugar()
		sugar = sugar.WithOptions(zap.AddCallerSkip(1))
		defer sugar.Sync()
		return &standardLogger{sugar}
	}
}

func (s *standardLogger) GetLogger() *zap.SugaredLogger {
	return s.logger
}

func (s *standardLogger) DebugWithContext(ctx context.Context, args ...interface{}) {
	var ctxKey *config.CtxKey
	if c, ok := ctx.Value(config.CtxKey{}).(*config.CtxKey); !ok {
		s.logger.Debug(args...)
		return
	} else {
		ctxKey = c
	}
	s.logger.WithOptions(zap.Fields(zap.Field{
		Key:    "request_id",
		Type:   15,
		String: ctxKey.RequestID,
	},
		zap.Field{
			Key:    "session_key",
			Type:   15,
			String: ctxKey.Session,
		})).Debug(args...)
}

func (s *standardLogger) DebugWithContextf(ctx context.Context, format string, args ...interface{}) {
	var ctxKey *config.CtxKey
	if c, ok := ctx.Value(config.CtxKey{}).(*config.CtxKey); !ok {
		s.logger.Debugf(format, args...)
		return
	} else {
		ctxKey = c
	}
	s.logger.WithOptions(zap.Fields(zap.Field{
		Key:    "request_id",
		Type:   15,
		String: ctxKey.RequestID,
	},
		zap.Field{
			Key:    "session_key",
			Type:   15,
			String: ctxKey.Session,
		},
	)).Debugf(format, args...)
}

func (s *standardLogger) ErrorWithContextf(ctx context.Context, format string, args ...interface{}) {
	var ctxKey *config.CtxKey
	if c, ok := ctx.Value(config.CtxKey{}).(*config.CtxKey); !ok {
		s.logger.Errorf(format, args...)
		return
	} else {
		ctxKey = c
	}

	reponseMessage := "unknown"
	if len(args) > 0 {
		if err, ok := args[len(args)-1].(error); ok {
			errString := err.Error()
			args = append(args[:len(args)-1], " "+errString)
			reponseMessage = err.Error()
		}
	}
	s.logger.WithOptions(zap.Fields(
		zap.Field{
			Key:    "response_message",
			Type:   15,
			String: reponseMessage,
		},
		zap.Field{
			Key:    "request_id",
			Type:   15,
			String: ctxKey.RequestID,
		},
		zap.Field{
			Key:    "session_key",
			Type:   15,
			String: ctxKey.Session,
		})).Errorf(format, args...)
}

func (s *standardLogger) ErrorWithContext(ctx context.Context, args ...interface{}) {
	var ctxKey *config.CtxKey
	if c, ok := ctx.Value(config.CtxKey{}).(*config.CtxKey); !ok {
		s.logger.Error(args...)
		return
	} else {
		ctxKey = c
	}
	reponseMessage := "unknown"
	if len(args) > 0 {
		if err, ok := args[len(args)-1].(error); ok {
			errString := err.Error()
			// if res, ok := err.(errors.Errors); ok {
			//  errString = res.ActualError()
			//  if errString == "" {
			//      errString = err.Error()
			//  }
			// }
			// fmt.Println("errString", errString)
			args = append(args[:len(args)-1], " "+errString)
			reponseMessage = err.Error()
		}
	}
	s.logger.WithOptions(zap.Fields(
		zap.Field{
			Key:    "response_message",
			Type:   15,
			String: reponseMessage,
		},
		zap.Field{
			Key:    "request_id",
			Type:   15,
			String: ctxKey.RequestID,
		},
		zap.Field{
			Key:    "session_key",
			Type:   15,
			String: ctxKey.Session,
		})).Error(args...)
}

func (s *standardLogger) InfoWithContext(ctx context.Context, args ...interface{}) {
	var ctxKey *config.CtxKey
	if c, ok := ctx.Value(config.CtxKey{}).(*config.CtxKey); !ok {
		s.logger.Info(args...)
		return
	} else {
		ctxKey = c
	}
	s.logger.WithOptions(zap.Fields(zap.Field{
		Key:    "request_id",
		Type:   15,
		String: ctxKey.RequestID,
	},
		zap.Field{
			Key:    "session_key",
			Type:   15,
			String: ctxKey.Session,
		})).Info(args...)
}

func (s *standardLogger) InfoWithContextf(ctx context.Context, format string, args ...interface{}) {
	var ctxKey *config.CtxKey
	if c, ok := ctx.Value(config.CtxKey{}).(*config.CtxKey); !ok {
		s.logger.Infof(format, args...)
		return
	} else {
		ctxKey = c
	}
	s.logger.WithOptions(zap.Fields(zap.Field{
		Key:    "request_id",
		Type:   15,
		String: ctxKey.RequestID,
	},
		zap.Field{
			Key:    "session_key",
			Type:   15,
			String: ctxKey.Session,
		})).Infof(format, args...)
}
