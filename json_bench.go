package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	zerolog "github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
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

	logrus.SetFormatter(&logrus.JSONFormatter{})

	type loggerType func(msg string, dummy *dummy)

	func() /*chronometer*/ {
		tic := time.Now()
		defer func() { fmt.Println("\n>Total duration:", time.Since(tic)) }()

		for _, f := range []struct {
			name   string
			logger loggerType
		}{
			{"Zerolog", func(msg string, dummy *dummy) { zerolog.Print(msg, "Zerolog", dummy) }},
			{"Logrus", func(msg string, dummy *dummy) { logrus.WithField("Logrus", dummy).Info(msg) }},
			{"Zap", func(msg string, dummy *dummy) { zlogger.Info(msg, zap.Stringer("Zap", dummy)) }},
			{"ZapSugar", func(msg string, dummy *dummy) { zSugarLogger.Sugar().Infow(msg, "ZapSugar", dummy) }},
			{"StdLog", func(msg string, dummy *dummy) { log.Print(msg, "StdLog", dummy) }},
		} {
			t := time.Now()
			for i := range textRandom {
				f.logger(textRandom[i], &dummyLog[i])
			}
			fmt.Println(f.name+":", int(time.Since(t).Nanoseconds())/len(textRandom), "ns per request")
		}
	} /*chronometer*/ ()
}
