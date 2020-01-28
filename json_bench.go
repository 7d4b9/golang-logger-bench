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

func main() {
	dummyLog := dummy{
		Foo: "foo",
		Bar: "bar",
	}

	randomText := [10000]string{}
	for i := range randomText {
		randomText[i] = strconv.Itoa(rand.Int())
	}

	zlogger, _ := zap.NewProduction()

	var wg sync.WaitGroup
	defer wg.Wait()

	for _, f := range []struct {
		name   string
		logger func(msg string)
	}{
		{"Zerolog", func(msg string) { zerolog.Print(msg, "Dummy", dummyLog) }},
		{"Logrus", func(msg string) { logrus.WithField("Dummy", dummyLog).Info(msg) }},
		{"Zap", func(msg string) { zlogger.Info(msg, zap.Stringer("Dummy", &dummyLog)) }},
		{"ZapSugar", func(msg string) { zlogger.Sugar().Infow(msg, "Dummy", dummyLog) }},
		{"StdLog", func(msg string) { log.Print(msg, "Dummy", dummyLog) }},
	} {
		t := time.Now()
		for i := range randomText {
			f.logger(randomText[i])
		}
		fmt.Println(f.name, ": ", int(time.Since(t).Nanoseconds())/len(randomText), " ns per request")
	}
}
