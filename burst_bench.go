package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
	zerolog "github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	zap "go.uber.org/zap"
)

const (
	messagesNumConf = "messages_num"
	burstLimitConf  = "burst_limit"
)

func init() {

	rand.Seed(time.Now().Local().Unix())

	viper.AutomaticEnv()
	viper.SetDefault(messagesNumConf, 10000)
	viper.SetDefault(burstLimitConf, 10000)
}

func main() {

	msgNum := viper.GetInt(messagesNumConf)
	burstLimit := viper.GetInt(burstLimitConf)

	rand.Seed(time.Now().Local().Unix())

	textRandom := make([]string, msgNum)
	for i := range textRandom {
		textRandom[i] = strconv.Itoa(rand.Int())
	}

	zlogger, _ := zap.NewProduction()
	defer zlogger.Sync()

	zSugarLogger, _ := zap.NewProduction()
	defer zSugarLogger.Sync()

	logrusLog := logrus.New()

	logrusLog.SetFormatter(&logrus.JSONFormatter{})

	loggerLog := logger.NewLogger(handler.Stream(os.Stderr, formatter.NewJSONEncoder()))

	func() /*chronometer*/ {
		tic := time.Now()
		defer func() { fmt.Println("\n>Total duration:", time.Since(tic)) }()

		type loggerType func(msg string)

		for _, f := range []struct {
			name   string
			logger loggerType
		}{
			{"Logger", func(msg string) {
				loggerLog.Info("Logger: "+msg, nil)
			}},
			{"Logrus", func(msg string) {
				logrusLog.Info("Logrus: " + msg)
			}},
			{"Zap", func(msg string) {
				zlogger.Info("Zap: " + msg)
			}},
			{"ZapSugar", func(msg string) {
				zSugarLogger.Sugar().Infow("ZapSugar: " + msg)
			}},
			{"Zerolog", func(msg string) {
				zerolog.Print("Zerolog: " + msg)
			}},
		} {
			sem := make(chan struct{}, burstLimit)
			defer close(sem)
			semAdd := func() { sem <- struct{}{} }
			semDone := func() { <-sem }
			semWait := func() {
				for i := 0; i < burstLimit; i++ {
					semAdd()
				}
			}
			t := time.Now()
			for i := range textRandom {
				semAdd()
				go func(s string) {
					defer semDone()
					f.logger(s)
				}(textRandom[i])
			}
			semWait()
			fmt.Println(f.name+":", int(time.Since(t).Nanoseconds())/len(textRandom), "ns per request")
		}
	} /*chronometer*/ ()
}
