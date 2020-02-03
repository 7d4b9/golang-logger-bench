package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

	dummyLog := make([]dummy, msgNum)
	for i := range dummyLog {
		dummyLog[i].Foo = strconv.Itoa(rand.Int())
		dummyLog[i].Bar = strconv.Itoa(rand.Int())
	}

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

		for _, f := range []struct {
			name string
			log  loggerType
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
			t := time.Now()
			for i := range textRandom {
				f.log(textRandom[i], &dummyLog[i])
			}
			fmt.Println(f.name+":", int(time.Since(t).Nanoseconds())/len(textRandom), "ns per request")
		}
	} /*chronometer*/ ()
}
