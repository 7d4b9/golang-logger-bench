package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	zerolog "github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	zap "go.uber.org/zap"
)

const messagesNumConf = "messages_num"

func init() {
	viper.AutomaticEnv()
	viper.SetDefault(messagesNumConf, 10000)
}

func main() {

	rand.Seed(time.Now().Local().Unix())

	textRandom := make([]string, viper.GetInt(messagesNumConf))

	for i := range textRandom {
		textRandom[i] = strconv.Itoa(rand.Int())
	}

	zlogger, _ := zap.NewProduction()
	defer zlogger.Sync()

	zSugarLogger, _ := zap.NewProduction()
	defer zSugarLogger.Sync()

	logrus.SetFormatter(&logrus.JSONFormatter{})

	func() /*chronometer*/ {
		tic := time.Now()
		defer func() { fmt.Println("\n>Total duration:", time.Since(tic)) }()

		var wg sync.WaitGroup
		defer wg.Wait()

		for _, current := range []struct {
			name   string
			logger func(msg string)
		}{
			{"Zerolog", func(msg string) { zerolog.Print("Zerolog: " + msg) }},
			{"Zap", func(msg string) { zlogger.Info("Zap: " + msg) }},
			{"ZapSugar", func(msg string) { zSugarLogger.Sugar().Infow("ZapSugar: " + msg) }},
			{"Logrus", func(msg string) { logrus.Info("Logrus: " + msg) }},
			{"StdLog", func(msg string) { log.Print("StdLog: " + msg) }},
		} {
			wg.Add(1)
			go func(current struct {
				name   string
				logger func(msg string)
			}) {
				defer wg.Done()
				t := time.Now()
				for i := range textRandom {
					current.logger(textRandom[i])
				}
				fmt.Println(current.name+":", int(time.Since(t).Nanoseconds())/len(textRandom), "ns per request")
			}(current)
		}
	} /*chronometer*/ ()
}
