package main

import (
	"bytes"
	"encoding/json"
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

	// ZAP logger extra lines (POO, whatever ?)
	zlogger, _ := zap.NewProduction()
	defer zlogger.Sync()

	var wg sync.WaitGroup
	defer wg.Wait()

	type loggerType func(msg string, dummy *dummy)

	for _, f := range []struct {
		name   string
		logger loggerType
	}{
		{"Zerolog", func(msg string, dummy *dummy) { zerolog.Print(msg, "Dummy", dummy) }},
		{"Logrus", func(msg string, dummy *dummy) { logrus.WithField("Dummy", dummy).Info(msg) }},
		{"Zap", func(msg string, dummy *dummy) { zlogger.Info(msg, zap.Stringer("Dummy", dummy)) }},
		{"ZapSugar", func(msg string, dummy *dummy) { zlogger.Sugar().Infow(msg, "Dummy", dummy) }},
		{"StdLog", func(msg string, dummy *dummy) { log.Print(msg, "Dummy", dummy) }},
	} {
		wg.Add(1)
		go func(f struct {
			name   string
			logger loggerType
		}) {
			defer wg.Done()
			t := time.Now()
			for i := range textRandom {
				f.logger(textRandom[i], &dummyLog[i])
			}
			fmt.Println(f.name, ": ", int(time.Since(t).Nanoseconds())/len(textRandom), " ns per request")
		}(f)
	}
}
