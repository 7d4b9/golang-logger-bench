package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	loggerlog "github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
	zerolog "github.com/rs/zerolog/log"
	logruslog "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	zap "go.uber.org/zap"
)

type dummy struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

func (d *dummy) String() string {
	b := bytes.NewBuffer([]byte{})
	json.NewEncoder(b).Encode(&d)
	return b.String()
}

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

	textRandom := make([]string, msgNum)
	for i := range textRandom {
		textRandom[i] = strconv.Itoa(rand.Int())
	}

	dummyLog := make([]dummy, msgNum)
	for i := range dummyLog {
		dummyLog[i].Foo = strconv.Itoa(rand.Int())
		dummyLog[i].Bar = strconv.Itoa(rand.Int())
	}

	// ZAP logger extra lines (POO, whatever ?)
	zlogger, _ := zap.NewProduction()
	defer zlogger.Sync()

	zSugarLogger, _ := zap.NewProduction()
	defer zSugarLogger.Sync()

	logrus := logruslog.New()

	logrus.SetFormatter(&logruslog.JSONFormatter{})

	logger := loggerlog.NewLogger(handler.Stream(os.Stderr, formatter.NewJSONEncoder()))

	type loggerType func(msg string, dummy *dummy)

	func() /*chronometer*/ {
		tic := time.Now()
		defer func() { fmt.Println("\n>Total duration:", time.Since(tic)) }()

		var wg sync.WaitGroup
		defer wg.Wait()
		for _, f := range []struct {
			name   string
			logger loggerType
		}{
			{"Logger", func(msg string, dummy *dummy) {
				logger.Info(msg, loggerlog.Ctx("Logger", dummy))
			}},
			{"Logrus", func(msg string, dummy *dummy) {
				logrus.WithField("Logrus", dummy).Info(msg)
			}},
			{"Zap", func(msg string, dummy *dummy) {
				zlogger.Info(msg, zap.Stringer("Zap", dummy))
			}},
			{"ZapSugar", func(msg string, dummy *dummy) {
				zSugarLogger.Sugar().Infow(msg, "ZapSugar", dummy)
			}},
			{"Zerolog", func(msg string, dummy *dummy) {
				zerolog.Info().Interface("Zerolog", dummy).Msg(msg)
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
				go func(s string, d *dummy) {
					defer semDone()
					f.logger(s, d)
				}(textRandom[i], &dummyLog[i])
			}
			semWait()
			fmt.Println(f.name+":", int(time.Since(t).Nanoseconds())/len(textRandom), "ns per request")
		}
		wg.Wait()
	} /*chronometer*/ ()
}
