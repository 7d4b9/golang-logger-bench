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

const messagesNumConf = "messages_num"

func init() {

	rand.Seed(time.Now().Local().Unix())

	viper.AutomaticEnv()
	viper.SetDefault(messagesNumConf, 10000)
}

func main() {

	msgNum := viper.GetInt(messagesNumConf)

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

		for _, f := range []struct {
			name   string
			logger func(msg string)
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
			t := time.Now()
			for i := range textRandom {
				f.logger(textRandom[i])
			}
			fmt.Println(f.name+":", int(time.Since(t).Nanoseconds())/len(textRandom), "ns per request")
		}
	} /*chronometer*/ ()
}
