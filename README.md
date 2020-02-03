# Golang logger benchmark

Generated README.md
> Makefile : bench json_bench parallele_bench json_parallele_bench bench_plot json_bench_plot parallele_bench_plot json_parallele_bench_plot
```sh
make all-plot
```
```sh
TEST: bench 10
sample:
{"Message":"Logger: 4233504706575583268","Level":6,"timestamp":1580702711.611,"Context":null}
{"level":"info","msg":"Logrus: 4233504706575583268","time":"2020-02-03T05:05:11+01:00"}
{"level":"info","ts":1580702711.611872,"caller":"logrus/bench.go:66","msg":"Zap: 4233504706575583268"}
{"level":"info","ts":1580702711.612224,"caller":"logrus/bench.go:69","msg":"ZapSugar: 4233504706575583268"}
{"level":"debug","time":"2020-02-03T05:05:11+01:00","message":"Zerolog: 4233504706575583268"}
--------------------
results:
Logger: 6694 ns per request
Logrus: 46379 ns per request
Zap: 21512 ns per request
ZapSugar: 16435 ns per request
Zerolog: 9241 ns per request

>Total duration: 1.068138ms
_____________________
---------------------
TEST: bench 100
sample:
{"Message":"Logger: 2880423983026293095","Level":6,"timestamp":1580702716.337,"Context":null}
{"level":"info","msg":"Logrus: 2880423983026293095","time":"2020-02-03T05:05:16+01:00"}
{"level":"info","ts":1580702716.337759,"caller":"logrus/bench.go:66","msg":"Zap: 2880423983026293095"}
{"level":"info","ts":1580702716.337868,"caller":"logrus/bench.go:69","msg":"ZapSugar: 2880423983026293095"}
{"level":"debug","time":"2020-02-03T05:05:16+01:00","message":"Zerolog: 2880423983026293095"}
--------------------
results:
Logger: 10277 ns per request
Logrus: 30069 ns per request
Zap: 13857 ns per request
ZapSugar: 14399 ns per request
Zerolog: 3935 ns per request

>Total duration: 7.518435ms
_____________________
---------------------
TEST: bench 1000
sample:
{"Message":"Logger: 5802080078091096200","Level":6,"timestamp":1580702722.416,"Context":null}
{"level":"info","msg":"Logrus: 5802080078091096200","time":"2020-02-03T05:05:22+01:00"}
{"level":"info","ts":1580702722.41598,"caller":"logrus/bench.go:66","msg":"Zap: 5802080078091096200"}
{"level":"info","ts":1580702722.4160872,"caller":"logrus/bench.go:69","msg":"ZapSugar: 5802080078091096200"}
{"level":"debug","time":"2020-02-03T05:05:22+01:00","message":"Zerolog: 5802080078091096200"}
--------------------
results:
Logger: 4031 ns per request
Logrus: 20961 ns per request
Zap: 7625 ns per request
ZapSugar: 8162 ns per request
Zerolog: 2997 ns per request

>Total duration: 43.840168ms
_____________________
---------------------
TEST: bench 10000
sample:
{"Message":"Logger: 3266273424374812236","Level":6,"timestamp":1580702727.685,"Context":null}
{"level":"info","msg":"Logrus: 3266273424374812236","time":"2020-02-03T05:05:27+01:00"}
{"level":"info","ts":1580702727.685049,"caller":"logrus/bench.go:66","msg":"Zap: 3266273424374812236"}
{"level":"info","ts":1580702727.685205,"caller":"logrus/bench.go:69","msg":"ZapSugar: 3266273424374812236"}
{"level":"debug","time":"2020-02-03T05:05:27+01:00","message":"Zerolog: 3266273424374812236"}
--------------------
results:
Logger: 4528 ns per request
Logrus: 11735 ns per request
Zap: 8029 ns per request
ZapSugar: 8556 ns per request
Zerolog: 4537 ns per request

>Total duration: 373.96032ms
_____________________
---------------------
TEST: bench 100000
sample:
{"Message":"Logger: 2217481443095153679","Level":6,"timestamp":1580702732.874,"Context":null}
{"level":"info","msg":"Logrus: 2217481443095153679","time":"2020-02-03T05:05:32+01:00"}
{"level":"info","ts":1580702732.874846,"caller":"logrus/bench.go:66","msg":"Zap: 2217481443095153679"}
{"level":"info","ts":1580702732.874978,"caller":"logrus/bench.go:69","msg":"ZapSugar: 2217481443095153679"}
{"level":"debug","time":"2020-02-03T05:05:32+01:00","message":"Zerolog: 2217481443095153679"}
--------------------
results:
Logger: 3517 ns per request
Logrus: 10281 ns per request
Zap: 8334 ns per request
ZapSugar: 7691 ns per request
Zerolog: 3369 ns per request

>Total duration: 3.319482422s
_____________________
---------------------
TEST: bench 1000000
sample:
{"Message":"Logger: 5894760961218256099","Level":6,"timestamp":1580702740.867,"Context":null}
{"level":"info","msg":"Logrus: 5894760961218256099","time":"2020-02-03T05:05:40+01:00"}
{"level":"info","ts":1580702740.867404,"caller":"logrus/bench.go:66","msg":"Zap: 5894760961218256099"}
{"level":"info","ts":1580702740.867521,"caller":"logrus/bench.go:69","msg":"ZapSugar: 5894760961218256099"}
{"level":"debug","time":"2020-02-03T05:05:40+01:00","message":"Zerolog: 5894760961218256099"}
--------------------
results:
Logger: 3651 ns per request
Logrus: 9336 ns per request
Zap: 6072 ns per request
ZapSugar: 6990 ns per request
Zerolog: 3398 ns per request

>Total duration: 29.450909841s
_____________________
---------------------
TEST: json_bench 10
sample:
{"Message":"1960278250673464644","Level":6,"timestamp":1580702775.810,"Context":{"LoggerStringer":{"foo":"3543276245033995829","bar":"4097785382807121587"}}}
{"Message":"1960278250673464644","Level":6,"timestamp":1580702775.810,"Context":{"LoggerReflect":{"foo":"3543276245033995829","bar":"4097785382807121587"}}}
{"Logrus":{"foo":"3543276245033995829","bar":"4097785382807121587"},"level":"info","msg":"1960278250673464644","time":"2020-02-03T05:06:15+01:00"}
{"level":"info","ts":1580702775.810712,"caller":"logrus/json_bench.go:90","msg":"1960278250673464644","Zap":"{\"foo\":\"3543276245033995829\",\"bar\":\"4097785382807121587\"}\n"}
{"level":"info","ts":1580702775.810931,"caller":"logrus/json_bench.go:93","msg":"1960278250673464644","ZapSugar":"{\"foo\":\"3543276245033995829\",\"bar\":\"4097785382807121587\"}\n"}
{"level":"info","Zerolog":{"foo":"3543276245033995829","bar":"4097785382807121587"},"time":"2020-02-03T05:06:15+01:00","message":"1960278250673464644"}
--------------------
results:
Logger.Stringer: 18628 ns per request
Logger.Reflect: 8883 ns per request
Logrus: 45828 ns per request
Zap: 25491 ns per request
ZapSugar: 19597 ns per request
Zerolog: 9486 ns per request

>Total duration: 1.329047ms
_____________________
---------------------
TEST: json_bench 100
sample:
{"Message":"8666075696248661767","Level":6,"timestamp":1580702780.584,"Context":{"LoggerStringer":{"foo":"3231059904283149733","bar":"5830176249980993364"}}}
{"Message":"8666075696248661767","Level":6,"timestamp":1580702780.584,"Context":{"LoggerReflect":{"foo":"3231059904283149733","bar":"5830176249980993364"}}}
{"Logrus":{"foo":"3231059904283149733","bar":"5830176249980993364"},"level":"info","msg":"8666075696248661767","time":"2020-02-03T05:06:20+01:00"}
{"level":"info","ts":1580702780.584202,"caller":"logrus/json_bench.go:90","msg":"8666075696248661767","Zap":"{\"foo\":\"3231059904283149733\",\"bar\":\"5830176249980993364\"}\n"}
{"level":"info","ts":1580702780.584337,"caller":"logrus/json_bench.go:93","msg":"8666075696248661767","ZapSugar":"{\"foo\":\"3231059904283149733\",\"bar\":\"5830176249980993364\"}\n"}
{"level":"info","Zerolog":{"foo":"3231059904283149733","bar":"5830176249980993364"},"time":"2020-02-03T05:06:20+01:00","message":"8666075696248661767"}
--------------------
results:
Logger.Stringer: 11055 ns per request
Logger.Reflect: 11836 ns per request
Logrus: 29010 ns per request
Zap: 12517 ns per request
ZapSugar: 17105 ns per request
Zerolog: 9139 ns per request

>Total duration: 9.161325ms
_____________________
---------------------
TEST: json_bench 1000
sample:
{"Message":"6503278325606638666","Level":6,"timestamp":1580702785.071,"Context":{"LoggerStringer":{"foo":"2433047025236777934","bar":"5185860088043466584"}}}
{"Message":"6503278325606638666","Level":6,"timestamp":1580702785.071,"Context":{"LoggerReflect":{"foo":"2433047025236777934","bar":"5185860088043466584"}}}
{"Logrus":{"foo":"2433047025236777934","bar":"5185860088043466584"},"level":"info","msg":"6503278325606638666","time":"2020-02-03T05:06:25+01:00"}
{"level":"info","ts":1580702785.0711,"caller":"logrus/json_bench.go:90","msg":"6503278325606638666","Zap":"{\"foo\":\"2433047025236777934\",\"bar\":\"5185860088043466584\"}\n"}
{"level":"info","ts":1580702785.071222,"caller":"logrus/json_bench.go:93","msg":"6503278325606638666","ZapSugar":"{\"foo\":\"2433047025236777934\",\"bar\":\"5185860088043466584\"}\n"}
{"level":"info","Zerolog":{"foo":"2433047025236777934","bar":"5185860088043466584"},"time":"2020-02-03T05:06:25+01:00","message":"6503278325606638666"}
--------------------
results:
Logger.Stringer: 13291 ns per request
Logger.Reflect: 11070 ns per request
Logrus: 23840 ns per request
Zap: 34016 ns per request
ZapSugar: 28515 ns per request
Zerolog: 12391 ns per request

>Total duration: 123.332668ms
_____________________
---------------------
TEST: json_bench 10000
sample:
{"Message":"203068497809304888","Level":6,"timestamp":1580702789.967,"Context":{"LoggerStringer":{"foo":"7777338899532831087","bar":"1035252674627393281"}}}
{"Message":"203068497809304888","Level":6,"timestamp":1580702789.967,"Context":{"LoggerReflect":{"foo":"7777338899532831087","bar":"1035252674627393281"}}}
{"Logrus":{"foo":"7777338899532831087","bar":"1035252674627393281"},"level":"info","msg":"203068497809304888","time":"2020-02-03T05:06:29+01:00"}
{"level":"info","ts":1580702789.967785,"caller":"logrus/json_bench.go:90","msg":"203068497809304888","Zap":"{\"foo\":\"7777338899532831087\",\"bar\":\"1035252674627393281\"}\n"}
{"level":"info","ts":1580702789.967948,"caller":"logrus/json_bench.go:93","msg":"203068497809304888","ZapSugar":"{\"foo\":\"7777338899532831087\",\"bar\":\"1035252674627393281\"}\n"}
{"level":"info","Zerolog":{"foo":"7777338899532831087","bar":"1035252674627393281"},"time":"2020-02-03T05:06:29+01:00","message":"203068497809304888"}
--------------------
results:
Logger.Stringer: 6323 ns per request
Logger.Reflect: 8626 ns per request
Logrus: 14532 ns per request
Zap: 15987 ns per request
ZapSugar: 14164 ns per request
Zerolog: 6566 ns per request

>Total duration: 662.169611ms
_____________________
---------------------
TEST: json_bench 100000
sample:
{"Message":"2847835954571374673","Level":6,"timestamp":1580702795.256,"Context":{"LoggerStringer":{"foo":"2664198489903170575","bar":"7004783195734700222"}}}
{"Message":"2847835954571374673","Level":6,"timestamp":1580702795.256,"Context":{"LoggerReflect":{"foo":"2664198489903170575","bar":"7004783195734700222"}}}
{"Logrus":{"foo":"2664198489903170575","bar":"7004783195734700222"},"level":"info","msg":"2847835954571374673","time":"2020-02-03T05:06:35+01:00"}
{"level":"info","ts":1580702795.2563908,"caller":"logrus/json_bench.go:90","msg":"2847835954571374673","Zap":"{\"foo\":\"2664198489903170575\",\"bar\":\"7004783195734700222\"}\n"}
{"level":"info","ts":1580702795.256511,"caller":"logrus/json_bench.go:93","msg":"2847835954571374673","ZapSugar":"{\"foo\":\"2664198489903170575\",\"bar\":\"7004783195734700222\"}\n"}
{"level":"info","Zerolog":{"foo":"2664198489903170575","bar":"7004783195734700222"},"time":"2020-02-03T05:06:35+01:00","message":"2847835954571374673"}
--------------------
results:
Logger.Stringer: 9177 ns per request
Logger.Reflect: 23575 ns per request
Logrus: 22967 ns per request
Zap: 17446 ns per request
ZapSugar: 14310 ns per request
Zerolog: 6498 ns per request

>Total duration: 9.397933345s
_____________________
---------------------
TEST: json_bench 1000000
sample:
{"Message":"6221968629135671716","Level":6,"timestamp":1580702810.933,"Context":{"LoggerStringer":{"foo":"1824446399818416752","bar":"7872015086999044677"}}}
{"Message":"6221968629135671716","Level":6,"timestamp":1580702810.933,"Context":{"LoggerReflect":{"foo":"1824446399818416752","bar":"7872015086999044677"}}}
{"Logrus":{"foo":"1824446399818416752","bar":"7872015086999044677"},"level":"info","msg":"6221968629135671716","time":"2020-02-03T05:06:50+01:00"}
{"level":"info","ts":1580702810.933088,"caller":"logrus/json_bench.go:90","msg":"6221968629135671716","Zap":"{\"foo\":\"1824446399818416752\",\"bar\":\"7872015086999044677\"}\n"}
{"level":"info","ts":1580702810.933231,"caller":"logrus/json_bench.go:93","msg":"6221968629135671716","ZapSugar":"{\"foo\":\"1824446399818416752\",\"bar\":\"7872015086999044677\"}\n"}
{"level":"info","Zerolog":{"foo":"1824446399818416752","bar":"7872015086999044677"},"time":"2020-02-03T05:06:50+01:00","message":"6221968629135671716"}
--------------------
results:
Logger.Stringer: 7384 ns per request
Logger.Reflect: 6802 ns per request
Logrus: 12593 ns per request
Zap: 9740 ns per request
ZapSugar: 10959 ns per request
Zerolog: 4200 ns per request

>Total duration: 51.679845786s
_____________________
---------------------
TEST: parallele_bench 10
sample:
{"Message":"Logger: 997238178200720982","Level":6,"timestamp":1580702870.003,"Context":null}
{"level":"info","ts":1580702870.002937,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 997238178200720982"}
{"level":"info","ts":1580702870.0028732,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 997238178200720982"}
{"level":"debug","time":"2020-02-03T05:07:50+01:00","message":"Zerolog: 997238178200720982"}
{"level":"info","msg":"Logrus: 997238178200720982","time":"2020-02-03T05:07:50+01:00"}
--------------------
results:
Logger: 21509 ns per request
ZapSugar: 27925 ns per request
Zap: 47035 ns per request
Zerolog: 58572 ns per request
Logrus: 69707 ns per request

>Total duration: 798.598µs
_____________________
---------------------
TEST: parallele_bench 100
sample:
{"Message":"Logger: 4009477010814587588","Level":6,"timestamp":1580702874.783,"Context":null}
{"level":"info","ts":1580702874.78302,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 4009477010814587588"}
{"level":"info","ts":1580702874.782846,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 4009477010814587588"}
{"level":"debug","time":"2020-02-03T05:07:54+01:00","message":"Zerolog: 4009477010814587588"}
{"level":"info","msg":"Logrus: 4009477010814587588","time":"2020-02-03T05:07:54+01:00"}
--------------------
results:
Zerolog: 33169 ns per request
Logger: 33765 ns per request
Zap: 38445 ns per request
ZapSugar: 40462 ns per request
Logrus: 46204 ns per request

>Total duration: 4.676638ms
_____________________
---------------------
TEST: parallele_bench 1000
sample:
{"level":"info","ts":1580702879.81532,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 1430894420084777095"}
{"level":"debug","time":"2020-02-03T05:07:59+01:00","message":"Zerolog: 1430894420084777095"}
{"Message":"Logger: 1430894420084777095","Level":6,"timestamp":1580702879.816,"Context":null}
{"level":"info","ts":1580702879.815559,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 1430894420084777095"}
{"level":"info","msg":"Logrus: 1430894420084777095","time":"2020-02-03T05:07:59+01:00"}
--------------------
results:
Zerolog: 31615 ns per request
Logger: 35316 ns per request
Zap: 45461 ns per request
ZapSugar: 45769 ns per request
Logrus: 46711 ns per request

>Total duration: 46.769055ms
_____________________
---------------------
TEST: parallele_bench 10000
sample:
{"level":"debug","time":"2020-02-03T05:08:04+01:00","message":"Zerolog: 8462226459573377611"}
{"Message":"Logger: 8462226459573377611","Level":6,"timestamp":1580702884.507,"Context":null}
{"level":"info","ts":1580702884.507041,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 8462226459573377611"}
{"level":"info","ts":1580702884.507335,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 8462226459573377611"}
{"level":"info","msg":"Logrus: 8462226459573377611","time":"2020-02-03T05:08:04+01:00"}
--------------------
results:
Zerolog: 30442 ns per request
Logger: 32262 ns per request
Zap: 40268 ns per request
ZapSugar: 49587 ns per request
Logrus: 50001 ns per request

>Total duration: 500.098899ms
_____________________
---------------------
TEST: parallele_bench 100000
sample:
{"Message":"Logger: 7289345803675786126","Level":6,"timestamp":1580702889.788,"Context":null}
{"level":"info","ts":1580702889.7881632,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 7289345803675786126"}
{"level":"info","ts":1580702889.788356,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 7289345803675786126"}
{"level":"debug","time":"2020-02-03T05:08:09+01:00","message":"Zerolog: 7289345803675786126"}
{"level":"info","msg":"Logrus: 7289345803675786126","time":"2020-02-03T05:08:09+01:00"}
--------------------
results:
Zerolog: 26537 ns per request
Logger: 28128 ns per request
ZapSugar: 35424 ns per request
Zap: 36761 ns per request
Logrus: 37877 ns per request

>Total duration: 3.787865877s
_____________________
---------------------
TEST: parallele_bench 1000000
sample:
{"Message":"Logger: 6677967560860132928","Level":6,"timestamp":1580702898.595,"Context":null}
{"level":"info","ts":1580702898.5949812,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 6677967560860132928"}
{"level":"info","ts":1580702898.5949922,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 6677967560860132928"}
{"level":"debug","time":"2020-02-03T05:08:18+01:00","message":"Zerolog: 6677967560860132928"}
{"level":"info","msg":"Logrus: 6677967560860132928","time":"2020-02-03T05:08:18+01:00"}
--------------------
results:
Zerolog: 28462 ns per request
Logger: 29011 ns per request
Zap: 34909 ns per request
ZapSugar: 34954 ns per request
Logrus: 36312 ns per request

>Total duration: 36.31250245s
_____________________
---------------------
TEST: json_parallele_bench 10
sample:
{"Message":"1322225383364148887","Level":6,"timestamp":1580702940.897,"Context":{"LoggerReflect":{"foo":"4970159107562610372","bar":"3326524632031252884"}}}
{"Message":"1322225383364148887","Level":6,"timestamp":1580702940.897,"Context":{"LoggerStringer":{"foo":"4970159107562610372","bar":"3326524632031252884"}}}
{"level":"info","Zerolog":{"foo":"4970159107562610372","bar":"3326524632031252884"},"time":"2020-02-03T05:09:00+01:00","message":"1322225383364148887"}
{"level":"info","ts":1580702940.8972669,"caller":"logrus/json_parallele_bench.go:93","msg":"1322225383364148887","Zap":"{\"foo\":\"4970159107562610372\",\"bar\":\"3326524632031252884\"}\n"}
{"Logrus":{"foo":"4970159107562610372","bar":"3326524632031252884"},"level":"info","msg":"1322225383364148887","time":"2020-02-03T05:09:00+01:00"}
{"level":"info","ts":1580702940.8972821,"caller":"logrus/json_parallele_bench.go:96","msg":"1322225383364148887","ZapSugar":"{\"foo\":\"4970159107562610372\",\"bar\":\"3326524632031252884\"}\n"}
--------------------
results:
Logger.Stringer: 48083 ns per request
Zerolog: 66889 ns per request
Logrus: 100564 ns per request
Logger.Reflect: 57072 ns per request
ZapSugar: 66774 ns per request
Zap: 123495 ns per request

>Total duration: 1.322118ms
_____________________
---------------------
TEST: json_parallele_bench 100
sample:
{"Message":"284747392667368410","Level":6,"timestamp":1580702945.683,"Context":{"LoggerStringer":{"foo":"4099096875926443509","bar":"2754620123431764537"}}}
{"level":"info","ts":1580702945.682831,"caller":"logrus/json_parallele_bench.go:93","msg":"284747392667368410","Zap":"{\"foo\":\"4099096875926443509\",\"bar\":\"2754620123431764537\"}\n"}
{"Message":"284747392667368410","Level":6,"timestamp":1580702945.683,"Context":{"LoggerReflect":{"foo":"4099096875926443509","bar":"2754620123431764537"}}}
{"level":"info","ts":1580702945.682925,"caller":"logrus/json_parallele_bench.go:96","msg":"284747392667368410","ZapSugar":"{\"foo\":\"4099096875926443509\",\"bar\":\"2754620123431764537\"}\n"}
{"level":"info","Zerolog":{"foo":"4099096875926443509","bar":"2754620123431764537"},"time":"2020-02-03T05:09:05+01:00","message":"284747392667368410"}
{"Logrus":{"foo":"4099096875926443509","bar":"2754620123431764537"},"level":"info","msg":"284747392667368410","time":"2020-02-03T05:09:05+01:00"}
--------------------
results:
Zerolog: 32873 ns per request
Logger.Reflect: 38602 ns per request
Zap: 43893 ns per request
Logger.Stringer: 50628 ns per request
ZapSugar: 50429 ns per request
Logrus: 54637 ns per request

>Total duration: 5.51247ms
_____________________
---------------------
TEST: json_parallele_bench 1000
sample:
{"Message":"6944894201763183134","Level":6,"timestamp":1580702950.668,"Context":{"LoggerStringer":{"foo":"3716014398094142597","bar":"4542937908549456570"}}}
{"Message":"6944894201763183134","Level":6,"timestamp":1580702950.668,"Context":{"LoggerReflect":{"foo":"3716014398094142597","bar":"4542937908549456570"}}}
{"level":"info","ts":1580702950.668185,"caller":"logrus/json_parallele_bench.go:93","msg":"6944894201763183134","Zap":"{\"foo\":\"3716014398094142597\",\"bar\":\"4542937908549456570\"}\n"}
{"level":"info","ts":1580702950.668343,"caller":"logrus/json_parallele_bench.go:96","msg":"6944894201763183134","ZapSugar":"{\"foo\":\"3716014398094142597\",\"bar\":\"4542937908549456570\"}\n"}
{"level":"info","Zerolog":{"foo":"3716014398094142597","bar":"4542937908549456570"},"time":"2020-02-03T05:09:10+01:00","message":"6944894201763183134"}
{"Logrus":{"foo":"3716014398094142597","bar":"4542937908549456570"},"level":"info","msg":"6944894201763183134","time":"2020-02-03T05:09:10+01:00"}
--------------------
results:
Zerolog: 43253 ns per request
Logger.Reflect: 44316 ns per request
Logger.Stringer: 44553 ns per request
Zap: 50065 ns per request
ZapSugar: 55028 ns per request
Logrus: 55994 ns per request

>Total duration: 56.059601ms
_____________________
---------------------
TEST: json_parallele_bench 10000
sample:
{"level":"info","ts":1580702955.385059,"caller":"logrus/json_parallele_bench.go:93","msg":"4661541960306149344","Zap":"{\"foo\":\"4253887256860228894\",\"bar\":\"3907353535080743836\"}\n"}
{"Message":"4661541960306149344","Level":6,"timestamp":1580702955.385,"Context":{"LoggerStringer":{"foo":"4253887256860228894","bar":"3907353535080743836"}}}
{"Message":"4661541960306149344","Level":6,"timestamp":1580702955.385,"Context":{"LoggerReflect":{"foo":"4253887256860228894","bar":"3907353535080743836"}}}
{"Logrus":{"foo":"4253887256860228894","bar":"3907353535080743836"},"level":"info","msg":"4661541960306149344","time":"2020-02-03T05:09:15+01:00"}
{"level":"info","Zerolog":{"foo":"4253887256860228894","bar":"3907353535080743836"},"time":"2020-02-03T05:09:15+01:00","message":"4661541960306149344"}
{"level":"info","ts":1580702955.385379,"caller":"logrus/json_parallele_bench.go:96","msg":"4661541960306149344","ZapSugar":"{\"foo\":\"4253887256860228894\",\"bar\":\"3907353535080743836\"}\n"}
--------------------
results:
Logger.Stringer: 42511 ns per request
Zerolog: 45083 ns per request
Logger.Reflect: 46671 ns per request
ZapSugar: 54238 ns per request
Zap: 65542 ns per request
Logrus: 69000 ns per request

>Total duration: 690.096344ms
_____________________
---------------------
TEST: json_parallele_bench 100000
sample:
{"Message":"3370306547525355812","Level":6,"timestamp":1580702960.965,"Context":{"LoggerReflect":{"foo":"3442037437415560415","bar":"5640324311325774880"}}}
{"Message":"3370306547525355812","Level":6,"timestamp":1580702960.965,"Context":{"LoggerStringer":{"foo":"3442037437415560415","bar":"5640324311325774880"}}}
{"level":"info","ts":1580702960.964727,"caller":"logrus/json_parallele_bench.go:93","msg":"3370306547525355812","Zap":"{\"foo\":\"3442037437415560415\",\"bar\":\"5640324311325774880\"}\n"}
{"level":"info","ts":1580702960.9648058,"caller":"logrus/json_parallele_bench.go:96","msg":"3370306547525355812","ZapSugar":"{\"foo\":\"3442037437415560415\",\"bar\":\"5640324311325774880\"}\n"}
{"Logrus":{"foo":"3442037437415560415","bar":"5640324311325774880"},"level":"info","msg":"3370306547525355812","time":"2020-02-03T05:09:20+01:00"}
{"level":"info","Zerolog":{"foo":"3442037437415560415","bar":"5640324311325774880"},"time":"2020-02-03T05:09:20+01:00","message":"3370306547525355812"}
--------------------
results:
Zerolog: 43901 ns per request
Logger.Stringer: 45113 ns per request
Logger.Reflect: 48649 ns per request
ZapSugar: 55128 ns per request
Logrus: 55879 ns per request
Zap: 57206 ns per request

>Total duration: 5.720913221s
_____________________
---------------------
TEST: json_parallele_bench 1000000
sample:
{"level":"info","Zerolog":{"foo":"4597087363093185535","bar":"6130153494376288746"},"time":"2020-02-03T05:09:33+01:00","message":"5964216177302604867"}
{"Message":"5964216177302604867","Level":6,"timestamp":1580702973.211,"Context":{"LoggerStringer":{"foo":"4597087363093185535","bar":"6130153494376288746"}}}
{"Message":"5964216177302604867","Level":6,"timestamp":1580702973.211,"Context":{"LoggerReflect":{"foo":"4597087363093185535","bar":"6130153494376288746"}}}
{"Logrus":{"foo":"4597087363093185535","bar":"6130153494376288746"},"level":"info","msg":"5964216177302604867","time":"2020-02-03T05:09:33+01:00"}
{"level":"info","ts":1580702973.211062,"caller":"logrus/json_parallele_bench.go:96","msg":"5964216177302604867","ZapSugar":"{\"foo\":\"4597087363093185535\",\"bar\":\"6130153494376288746\"}\n"}
{"level":"info","ts":1580702973.211062,"caller":"logrus/json_parallele_bench.go:93","msg":"5964216177302604867","Zap":"{\"foo\":\"4597087363093185535\",\"bar\":\"6130153494376288746\"}\n"}
--------------------
results:
Zerolog: 43430 ns per request
Logger.Stringer: 45112 ns per request
Logger.Reflect: 45116 ns per request
ZapSugar: 53824 ns per request
Zap: 54865 ns per request
Logrus: 55067 ns per request

>Total duration: 55.067893921s
_____________________
---------------------
```
