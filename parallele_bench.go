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
	zap "go.uber.org/zap"
)

func main() {

	rand.Seed(time.Now().Local().Unix())

	textRandom := [10000]string{}

	for i := range textRandom {
		textRandom[i] = strconv.Itoa(rand.Int())
	}

	zlogger, _ := zap.NewProduction()

	var wg sync.WaitGroup
	defer wg.Wait()

	for _, current := range []struct {
		name   string
		logger func(msg string)
	}{
		{"Zerolog", func(msg string) { zerolog.Print(msg, "Dummy") }},
		{"Zap", func(msg string) { zlogger.Info(msg) }},
		{"ZapSugar", func(msg string) { zlogger.Sugar().Infow(msg) }},
		{"Logrus", func(msg string) { logrus.Info(msg) }},
		{"StdLog", func(msg string) { log.Print(msg) }},
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
			totalDuration := time.Since(t)
			fmt.Println(current.name, int(totalDuration.Nanoseconds())/len(textRandom), " ns per request")
		}(current)
	}
}
