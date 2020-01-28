package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	zerolog "github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	zap "go.uber.org/zap"
)

func main() {

	textRandom := [10000]string{}
	for i := range textRandom {
		textRandom[i] = strconv.Itoa(rand.Int())
	}

	zlogger, _ := zap.NewProduction()

	for _, f := range []struct {
		name   string
		logger func(msg string)
	}{
		{"Zerolog", func(msg string) { zerolog.Print(msg) }},
		{"Logrus", func(msg string) { logrus.Info(msg) }},
		{"Zap", func(msg string) { zlogger.Info(msg) }},
		{"ZapSugar", func(msg string) { zlogger.Sugar().Infow(msg) }},
		{"StdLog", func(msg string) { log.Print(msg) }},
	} {
		t := time.Now()
		for i := range textRandom {
			f.logger(textRandom[i])
		}
		fmt.Println(f.name, ": ", int(time.Since(t).Nanoseconds())/len(textRandom), " ns per request")
	}
}
