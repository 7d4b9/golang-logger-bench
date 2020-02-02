# Golang logger benchmark

Generated README.md
> Makefile : bench json_bench parallele_bench json_parallele_bench bench_plot json_bench_plot parallele_bench_plot json_parallele_bench_plot
```sh
make all-plot
```
```sh
TEST: bench 10
sample:
{"Message":"Logger: 6279746941961538829","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 6279746941961538829","time":"2020-02-02T11:56:12+01:00"}
{"level":"info","ts":1580640972.3553731,"caller":"logrus/bench.go:66","msg":"Zap: 6279746941961538829"}
{"level":"info","ts":1580640972.355463,"caller":"logrus/bench.go:69","msg":"ZapSugar: 6279746941961538829"}
{"level":"debug","time":"2020-02-02T11:56:12+01:00","message":"Zerolog: 6279746941961538829"}
--------------------
results:
Logger: 4650 ns per request
Logrus: 36036 ns per request
Zap: 15554 ns per request
ZapSugar: 10565 ns per request
Zerolog: 5126 ns per request

>Total duration: 839.307µs
_____________________
---------------------
TEST: bench 100
sample:
{"Message":"Logger: 8043102384864134523","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 8043102384864134523","time":"2020-02-02T11:56:16+01:00"}
{"level":"info","ts":1580640976.3777862,"caller":"logrus/bench.go:66","msg":"Zap: 8043102384864134523"}
{"level":"info","ts":1580640976.377897,"caller":"logrus/bench.go:69","msg":"ZapSugar: 8043102384864134523"}
{"level":"debug","time":"2020-02-02T11:56:16+01:00","message":"Zerolog: 8043102384864134523"}
--------------------
results:
Logger: 2518 ns per request
Logrus: 11564 ns per request
Zap: 6459 ns per request
ZapSugar: 6708 ns per request
Zerolog: 2946 ns per request

>Total duration: 3.149507ms
_____________________
---------------------
TEST: bench 1000
sample:
{"Message":"Logger: 726665127020169825","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 726665127020169825","time":"2020-02-02T11:56:20+01:00"}
{"level":"info","ts":1580640980.164404,"caller":"logrus/bench.go:66","msg":"Zap: 726665127020169825"}
{"level":"info","ts":1580640980.1645,"caller":"logrus/bench.go:69","msg":"ZapSugar: 726665127020169825"}
{"level":"debug","time":"2020-02-02T11:56:20+01:00","message":"Zerolog: 726665127020169825"}
--------------------
results:
Logger: 2368 ns per request
Logrus: 8582 ns per request
Zap: 5356 ns per request
ZapSugar: 7435 ns per request
Zerolog: 3569 ns per request

>Total duration: 27.520539ms
_____________________
---------------------
TEST: bench 10000
sample:
{"Message":"Logger: 3652555429967922895","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 3652555429967922895","time":"2020-02-02T11:56:24+01:00"}
{"level":"info","ts":1580640984.1060228,"caller":"logrus/bench.go:66","msg":"Zap: 3652555429967922895"}
{"level":"info","ts":1580640984.1061141,"caller":"logrus/bench.go:69","msg":"ZapSugar: 3652555429967922895"}
{"level":"debug","time":"2020-02-02T11:56:24+01:00","message":"Zerolog: 3652555429967922895"}
--------------------
results:
Logger: 2264 ns per request
Logrus: 7484 ns per request
Zap: 5253 ns per request
ZapSugar: 5984 ns per request
Zerolog: 2725 ns per request

>Total duration: 237.45366ms
_____________________
---------------------
TEST: bench 100000
sample:
{"Message":"Logger: 5558590184485769013","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 5558590184485769013","time":"2020-02-02T11:56:28+01:00"}
{"level":"info","ts":1580640988.190465,"caller":"logrus/bench.go:66","msg":"Zap: 5558590184485769013"}
{"level":"info","ts":1580640988.190561,"caller":"logrus/bench.go:69","msg":"ZapSugar: 5558590184485769013"}
{"level":"debug","time":"2020-02-02T11:56:28+01:00","message":"Zerolog: 5558590184485769013"}
--------------------
results:
Logger: 2230 ns per request
Logrus: 8417 ns per request
Zap: 5396 ns per request
ZapSugar: 6084 ns per request
Zerolog: 2838 ns per request

>Total duration: 2.49702142s
_____________________
---------------------
TEST: bench 1000000
sample:
{"Message":"Logger: 8138231370733294038","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 8138231370733294038","time":"2020-02-02T11:56:34+01:00"}
{"level":"info","ts":1580640994.4666848,"caller":"logrus/bench.go:66","msg":"Zap: 8138231370733294038"}
{"level":"info","ts":1580640994.466779,"caller":"logrus/bench.go:69","msg":"ZapSugar: 8138231370733294038"}
{"level":"debug","time":"2020-02-02T11:56:34+01:00","message":"Zerolog: 8138231370733294038"}
--------------------
results:
Logger: 2145 ns per request
Logrus: 7624 ns per request
Zap: 5151 ns per request
ZapSugar: 5935 ns per request
Zerolog: 2812 ns per request

>Total duration: 23.670762235s
_____________________
---------------------
TEST: json_bench 10
sample:
{"Message":"4559346391730184767","Level":6,"Context":{"LoggerStringer":{"foo":"4174264162473799115","bar":"4409289114205181193"}}}
{"Message":"4559346391730184767","Level":6,"Context":{"LoggerReflect":{"foo":"4174264162473799115","bar":"4409289114205181193"}}}
{"Logrus":{"foo":"4174264162473799115","bar":"4409289114205181193"},"level":"info","msg":"4559346391730184767","time":"2020-02-02T11:57:02+01:00"}
{"level":"info","ts":1580641022.244892,"caller":"logrus/json_bench.go:90","msg":"4559346391730184767","Zap":"{\"foo\":\"4174264162473799115\",\"bar\":\"4409289114205181193\"}\n"}
{"level":"info","ts":1580641022.244996,"caller":"logrus/json_bench.go:93","msg":"4559346391730184767","ZapSugar":"{\"foo\":\"4174264162473799115\",\"bar\":\"4409289114205181193\"}\n"}
{"level":"info","Zerolog":{"foo":"4174264162473799115","bar":"4409289114205181193"},"time":"2020-02-02T11:57:02+01:00","message":"4559346391730184767"}
--------------------
results:
Logger.Stringer: 16021 ns per request
Logger.Reflect: 5224 ns per request
Logrus: 34013 ns per request
Zap: 20592 ns per request
ZapSugar: 15056 ns per request
Zerolog: 6123 ns per request

>Total duration: 1.099173ms
_____________________
---------------------
TEST: json_bench 100
sample:
{"Message":"2455379202271764944","Level":6,"Context":{"LoggerStringer":{"foo":"5655618399791757339","bar":"1530494768974504477"}}}
{"Message":"2455379202271764944","Level":6,"Context":{"LoggerReflect":{"foo":"5655618399791757339","bar":"1530494768974504477"}}}
{"Logrus":{"foo":"5655618399791757339","bar":"1530494768974504477"},"level":"info","msg":"2455379202271764944","time":"2020-02-02T11:57:05+01:00"}
{"level":"info","ts":1580641025.947343,"caller":"logrus/json_bench.go:90","msg":"2455379202271764944","Zap":"{\"foo\":\"5655618399791757339\",\"bar\":\"1530494768974504477\"}\n"}
{"level":"info","ts":1580641025.9474518,"caller":"logrus/json_bench.go:93","msg":"2455379202271764944","ZapSugar":"{\"foo\":\"5655618399791757339\",\"bar\":\"1530494768974504477\"}\n"}
{"level":"info","Zerolog":{"foo":"5655618399791757339","bar":"1530494768974504477"},"time":"2020-02-02T11:57:05+01:00","message":"2455379202271764944"}
--------------------
results:
Logger.Stringer: 6025 ns per request
Logger.Reflect: 4969 ns per request
Logrus: 14499 ns per request
Zap: 10024 ns per request
ZapSugar: 10232 ns per request
Zerolog: 3869 ns per request

>Total duration: 5.108693ms
_____________________
---------------------
TEST: json_bench 1000
sample:
{"Message":"4292919050662571574","Level":6,"Context":{"LoggerStringer":{"foo":"1867501916500423627","bar":"6448274562673477899"}}}
{"Message":"4292919050662571574","Level":6,"Context":{"LoggerReflect":{"foo":"1867501916500423627","bar":"6448274562673477899"}}}
{"Logrus":{"foo":"1867501916500423627","bar":"6448274562673477899"},"level":"info","msg":"4292919050662571574","time":"2020-02-02T11:57:09+01:00"}
{"level":"info","ts":1580641029.836094,"caller":"logrus/json_bench.go:90","msg":"4292919050662571574","Zap":"{\"foo\":\"1867501916500423627\",\"bar\":\"6448274562673477899\"}\n"}
{"level":"info","ts":1580641029.836202,"caller":"logrus/json_bench.go:93","msg":"4292919050662571574","ZapSugar":"{\"foo\":\"1867501916500423627\",\"bar\":\"6448274562673477899\"}\n"}
{"level":"info","Zerolog":{"foo":"1867501916500423627","bar":"6448274562673477899"},"time":"2020-02-02T11:57:09+01:00","message":"4292919050662571574"}
--------------------
results:
Logger.Stringer: 5239 ns per request
Logger.Reflect: 5465 ns per request
Logrus: 11512 ns per request
Zap: 8464 ns per request
ZapSugar: 9464 ns per request
Zerolog: 3625 ns per request

>Total duration: 44.054865ms
_____________________
---------------------
TEST: json_bench 10000
sample:
{"Message":"6168663394785585828","Level":6,"Context":{"LoggerStringer":{"foo":"7210810555893977420","bar":"220235859916644455"}}}
{"Message":"6168663394785585828","Level":6,"Context":{"LoggerReflect":{"foo":"7210810555893977420","bar":"220235859916644455"}}}
{"Logrus":{"foo":"7210810555893977420","bar":"220235859916644455"},"level":"info","msg":"6168663394785585828","time":"2020-02-02T11:57:13+01:00"}
{"level":"info","ts":1580641033.971068,"caller":"logrus/json_bench.go:90","msg":"6168663394785585828","Zap":"{\"foo\":\"7210810555893977420\",\"bar\":\"220235859916644455\"}\n"}
{"level":"info","ts":1580641033.971207,"caller":"logrus/json_bench.go:93","msg":"6168663394785585828","ZapSugar":"{\"foo\":\"7210810555893977420\",\"bar\":\"220235859916644455\"}\n"}
{"level":"info","Zerolog":{"foo":"7210810555893977420","bar":"220235859916644455"},"time":"2020-02-02T11:57:13+01:00","message":"6168663394785585828"}
--------------------
results:
Logger.Stringer: 8997 ns per request
Logger.Reflect: 5801 ns per request
Logrus: 13033 ns per request
Zap: 10265 ns per request
ZapSugar: 12492 ns per request
Zerolog: 4074 ns per request

>Total duration: 547.087467ms
_____________________
---------------------
TEST: json_bench 100000
sample:
{"Message":"8996321348197837245","Level":6,"Context":{"LoggerStringer":{"foo":"943107202896935780","bar":"7838501383458671136"}}}
{"Message":"8996321348197837245","Level":6,"Context":{"LoggerReflect":{"foo":"943107202896935780","bar":"7838501383458671136"}}}
{"Logrus":{"foo":"943107202896935780","bar":"7838501383458671136"},"level":"info","msg":"8996321348197837245","time":"2020-02-02T11:57:19+01:00"}
{"level":"info","ts":1580641039.967689,"caller":"logrus/json_bench.go:90","msg":"8996321348197837245","Zap":"{\"foo\":\"943107202896935780\",\"bar\":\"7838501383458671136\"}\n"}
{"level":"info","ts":1580641039.967807,"caller":"logrus/json_bench.go:93","msg":"8996321348197837245","ZapSugar":"{\"foo\":\"943107202896935780\",\"bar\":\"7838501383458671136\"}\n"}
{"level":"info","Zerolog":{"foo":"943107202896935780","bar":"7838501383458671136"},"time":"2020-02-02T11:57:19+01:00","message":"8996321348197837245"}
--------------------
results:
Logger.Stringer: 4825 ns per request
Logger.Reflect: 4572 ns per request
Logrus: 10671 ns per request
Zap: 9778 ns per request
ZapSugar: 9621 ns per request
Zerolog: 3605 ns per request

>Total duration: 4.308002335s
_____________________
---------------------
TEST: json_bench 1000000
sample:
{"Message":"5450230664996450636","Level":6,"Context":{"LoggerStringer":{"foo":"987833720859454117","bar":"9079100341858192133"}}}
{"Message":"5450230664996450636","Level":6,"Context":{"LoggerReflect":{"foo":"987833720859454117","bar":"9079100341858192133"}}}
{"Logrus":{"foo":"987833720859454117","bar":"9079100341858192133"},"level":"info","msg":"5450230664996450636","time":"2020-02-02T11:57:29+01:00"}
{"level":"info","ts":1580641049.324028,"caller":"logrus/json_bench.go:90","msg":"5450230664996450636","Zap":"{\"foo\":\"987833720859454117\",\"bar\":\"9079100341858192133\"}\n"}
{"level":"info","ts":1580641049.324147,"caller":"logrus/json_bench.go:93","msg":"5450230664996450636","ZapSugar":"{\"foo\":\"987833720859454117\",\"bar\":\"9079100341858192133\"}\n"}
{"level":"info","Zerolog":{"foo":"987833720859454117","bar":"9079100341858192133"},"time":"2020-02-02T11:57:29+01:00","message":"5450230664996450636"}
--------------------
results:
Logger.Stringer: 4755 ns per request
Logger.Reflect: 4390 ns per request
Logrus: 10784 ns per request
Zap: 8498 ns per request
ZapSugar: 9318 ns per request
Zerolog: 3594 ns per request

>Total duration: 41.342668723s
_____________________
---------------------
TEST: parallele_bench 10
sample:
{"Message":"Logger: 7989478113641876326","Level":6,"Context":null}
{"level":"info","ts":1580641096.4300191,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 7989478113641876326"}
{"level":"debug","time":"2020-02-02T11:58:16+01:00","message":"Zerolog: 7989478113641876326"}
{"level":"info","ts":1580641096.4300869,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 7989478113641876326"}
{"level":"info","msg":"Logrus: 7989478113641876326","time":"2020-02-02T11:58:16+01:00"}
--------------------
results:
Logger: 19876 ns per request
ZapSugar: 31976 ns per request
Zap: 48695 ns per request
Zerolog: 53911 ns per request
Logrus: 66017 ns per request

>Total duration: 744.42µs
_____________________
---------------------
TEST: parallele_bench 100
sample:
{"Message":"Logger: 526460676597058516","Level":6,"Context":null}
{"level":"info","ts":1580641100.330647,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 526460676597058516"}
{"level":"info","ts":1580641100.3308291,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 526460676597058516"}
{"level":"debug","time":"2020-02-02T11:58:20+01:00","message":"Zerolog: 526460676597058516"}
{"level":"info","msg":"Logrus: 526460676597058516","time":"2020-02-02T11:58:20+01:00"}
--------------------
results:
Logger: 24880 ns per request
Zerolog: 30402 ns per request
Zap: 32676 ns per request
ZapSugar: 33771 ns per request
Logrus: 36010 ns per request

>Total duration: 3.6799ms
_____________________
---------------------
TEST: parallele_bench 1000
sample:
{"Message":"Logger: 2442678073159455802","Level":6,"Context":null}
{"level":"debug","time":"2020-02-02T11:58:24+01:00","message":"Zerolog: 2442678073159455802"}
{"level":"info","ts":1580641104.0794141,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 2442678073159455802"}
{"level":"info","ts":1580641104.079447,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 2442678073159455802"}
{"level":"info","msg":"Logrus: 2442678073159455802","time":"2020-02-02T11:58:24+01:00"}
--------------------
results:
Logger: 25796 ns per request
Zerolog: 27572 ns per request
Zap: 34154 ns per request
ZapSugar: 35312 ns per request
Logrus: 35855 ns per request

>Total duration: 35.94148ms
_____________________
---------------------
TEST: parallele_bench 10000
sample:
{"Message":"Logger: 8333187893650326475","Level":6,"Context":null}
{"level":"info","ts":1580641107.868845,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 8333187893650326475"}
{"level":"info","ts":1580641107.868892,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 8333187893650326475"}
{"level":"debug","time":"2020-02-02T11:58:27+01:00","message":"Zerolog: 8333187893650326475"}
{"level":"info","msg":"Logrus: 8333187893650326475","time":"2020-02-02T11:58:27+01:00"}
--------------------
results:
Logger: 21482 ns per request
Zerolog: 27030 ns per request
Zap: 30283 ns per request
ZapSugar: 31330 ns per request
Logrus: 32245 ns per request

>Total duration: 322.53731ms
_____________________
---------------------
TEST: parallele_bench 100000
sample:
{"Message":"Logger: 7237910770202756494","Level":6,"Context":null}
{"level":"info","ts":1580641112.034805,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 7237910770202756494"}
{"level":"debug","time":"2020-02-02T11:58:32+01:00","message":"Zerolog: 7237910770202756494"}
{"level":"info","ts":1580641112.034772,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 7237910770202756494"}
{"level":"info","msg":"Logrus: 7237910770202756494","time":"2020-02-02T11:58:32+01:00"}
--------------------
results:
Logger: 24705 ns per request
Zerolog: 25096 ns per request
ZapSugar: 29656 ns per request
Zap: 30241 ns per request
Logrus: 31091 ns per request

>Total duration: 3.109221789s
_____________________
---------------------
TEST: parallele_bench 1000000
sample:
{"Message":"Logger: 658954141239483439","Level":6,"Context":null}
{"level":"info","ts":1580641118.9222558,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 658954141239483439"}
{"level":"info","ts":1580641118.922411,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 658954141239483439"}
{"level":"debug","time":"2020-02-02T11:58:38+01:00","message":"Zerolog: 658954141239483439"}
{"level":"info","msg":"Logrus: 658954141239483439","time":"2020-02-02T11:58:38+01:00"}
--------------------
results:
Logger: 22366 ns per request
Zerolog: 24172 ns per request
ZapSugar: 29840 ns per request
Zap: 30070 ns per request
Logrus: 30862 ns per request

>Total duration: 30.862767641s
_____________________
---------------------
TEST: json_parallele_bench 10
sample:
{"Message":"5971111271726804112","Level":6,"Context":{"LoggerStringer":{"foo":"837791428891411887","bar":"5946401255663496721"}}}
{"level":"info","ts":1580641153.806365,"caller":"logrus/json_parallele_bench.go:93","msg":"5971111271726804112","Zap":"{\"foo\":\"837791428891411887\",\"bar\":\"5946401255663496721\"}\n"}
{"Message":"5971111271726804112","Level":6,"Context":{"LoggerReflect":{"foo":"837791428891411887","bar":"5946401255663496721"}}}
{"level":"info","ts":1580641153.80654,"caller":"logrus/json_parallele_bench.go:96","msg":"5971111271726804112","ZapSugar":"{\"foo\":\"837791428891411887\",\"bar\":\"5946401255663496721\"}\n"}
{"level":"info","Zerolog":{"foo":"837791428891411887","bar":"5946401255663496721"},"time":"2020-02-02T11:59:13+01:00","message":"5971111271726804112"}
{"Logrus":{"foo":"837791428891411887","bar":"5946401255663496721"},"level":"info","msg":"5971111271726804112","time":"2020-02-02T11:59:13+01:00"}
--------------------
results:
Logger.Stringer: 40537 ns per request
Logger.Reflect: 60789 ns per request
Zerolog: 85187 ns per request
ZapSugar: 69183 ns per request
Zap: 73790 ns per request
Logrus: 102471 ns per request

>Total duration: 1.100504ms
_____________________
---------------------
TEST: json_parallele_bench 100
sample:
{"Message":"7877435496848272630","Level":6,"Context":{"LoggerStringer":{"foo":"6182341610775953743","bar":"1687699754048148733"}}}
{"level":"info","ts":1580641157.930756,"caller":"logrus/json_parallele_bench.go:93","msg":"7877435496848272630","Zap":"{\"foo\":\"6182341610775953743\",\"bar\":\"1687699754048148733\"}\n"}
{"level":"info","ts":1580641157.9308522,"caller":"logrus/json_parallele_bench.go:96","msg":"7877435496848272630","ZapSugar":"{\"foo\":\"6182341610775953743\",\"bar\":\"1687699754048148733\"}\n"}
{"Message":"7877435496848272630","Level":6,"Context":{"LoggerReflect":{"foo":"6182341610775953743","bar":"1687699754048148733"}}}
{"level":"info","Zerolog":{"foo":"6182341610775953743","bar":"1687699754048148733"},"time":"2020-02-02T11:59:17+01:00","message":"7877435496848272630"}
{"Logrus":{"foo":"6182341610775953743","bar":"1687699754048148733"},"level":"info","msg":"7877435496848272630","time":"2020-02-02T11:59:17+01:00"}
--------------------
results:
Zerolog: 34755 ns per request
Logger.Stringer: 37613 ns per request
Logger.Reflect: 39059 ns per request
ZapSugar: 43821 ns per request
Zap: 45684 ns per request
Logrus: 48646 ns per request

>Total duration: 4.943286ms
_____________________
---------------------
TEST: json_parallele_bench 1000
sample:
{"level":"info","ts":1580641161.885012,"caller":"logrus/json_parallele_bench.go:96","msg":"420012667130226020","ZapSugar":"{\"foo\":\"2958029448348868272\",\"bar\":\"4589322828674054986\"}\n"}
{"Message":"420012667130226020","Level":6,"Context":{"LoggerReflect":{"foo":"2958029448348868272","bar":"4589322828674054986"}}}
{"Message":"420012667130226020","Level":6,"Context":{"LoggerStringer":{"foo":"2958029448348868272","bar":"4589322828674054986"}}}
{"level":"info","ts":1580641161.88502,"caller":"logrus/json_parallele_bench.go:93","msg":"420012667130226020","Zap":"{\"foo\":\"2958029448348868272\",\"bar\":\"4589322828674054986\"}\n"}
{"level":"info","Zerolog":{"foo":"2958029448348868272","bar":"4589322828674054986"},"time":"2020-02-02T11:59:21+01:00","message":"420012667130226020"}
{"Logrus":{"foo":"2958029448348868272","bar":"4589322828674054986"},"level":"info","msg":"420012667130226020","time":"2020-02-02T11:59:21+01:00"}
--------------------
results:
Zerolog: 34003 ns per request
Logger.Reflect: 38299 ns per request
Logger.Stringer: 38540 ns per request
Zap: 45433 ns per request
ZapSugar: 45555 ns per request
Logrus: 46375 ns per request

>Total duration: 46.457737ms
_____________________
---------------------
TEST: json_parallele_bench 10000
sample:
{"Message":"2335127221385864778","Level":6,"Context":{"LoggerStringer":{"foo":"8280105674315579951","bar":"520671903339819635"}}}
{"Message":"2335127221385864778","Level":6,"Context":{"LoggerReflect":{"foo":"8280105674315579951","bar":"520671903339819635"}}}
{"level":"info","ts":1580641165.880037,"caller":"logrus/json_parallele_bench.go:93","msg":"2335127221385864778","Zap":"{\"foo\":\"8280105674315579951\",\"bar\":\"520671903339819635\"}\n"}
{"level":"info","ts":1580641165.880044,"caller":"logrus/json_parallele_bench.go:96","msg":"2335127221385864778","ZapSugar":"{\"foo\":\"8280105674315579951\",\"bar\":\"520671903339819635\"}\n"}
{"level":"info","Zerolog":{"foo":"8280105674315579951","bar":"520671903339819635"},"time":"2020-02-02T11:59:25+01:00","message":"2335127221385864778"}
{"Logrus":{"foo":"8280105674315579951","bar":"520671903339819635"},"level":"info","msg":"2335127221385864778","time":"2020-02-02T11:59:25+01:00"}
--------------------
results:
Logger.Stringer: 35645 ns per request
Zerolog: 37677 ns per request
Logger.Reflect: 38917 ns per request
Zap: 44304 ns per request
ZapSugar: 44599 ns per request
Logrus: 44806 ns per request

>Total duration: 448.141361ms
_____________________
---------------------
TEST: json_parallele_bench 100000
sample:
{"Message":"932575069855018893","Level":6,"Context":{"LoggerStringer":{"foo":"8188608732229792240","bar":"9036206741081842424"}}}
{"Message":"932575069855018893","Level":6,"Context":{"LoggerReflect":{"foo":"8188608732229792240","bar":"9036206741081842424"}}}
{"level":"info","ts":1580641170.3735838,"caller":"logrus/json_parallele_bench.go:93","msg":"932575069855018893","Zap":"{\"foo\":\"8188608732229792240\",\"bar\":\"9036206741081842424\"}\n"}
{"level":"info","Zerolog":{"foo":"8188608732229792240","bar":"9036206741081842424"},"time":"2020-02-02T11:59:30+01:00","message":"932575069855018893"}
{"level":"info","ts":1580641170.373605,"caller":"logrus/json_parallele_bench.go:96","msg":"932575069855018893","ZapSugar":"{\"foo\":\"8188608732229792240\",\"bar\":\"9036206741081842424\"}\n"}
{"Logrus":{"foo":"8188608732229792240","bar":"9036206741081842424"},"level":"info","msg":"932575069855018893","time":"2020-02-02T11:59:30+01:00"}
--------------------
results:
Zerolog: 36812 ns per request
Logger.Reflect: 37922 ns per request
Logger.Stringer: 38837 ns per request
Zap: 45259 ns per request
ZapSugar: 45447 ns per request
Logrus: 45858 ns per request

>Total duration: 4.585967878s
_____________________
---------------------
TEST: json_parallele_bench 1000000
sample:
{"Message":"575269020497523263","Level":6,"Context":{"LoggerStringer":{"foo":"3510565857353876841","bar":"4578135596460776054"}}}
{"Message":"575269020497523263","Level":6,"Context":{"LoggerReflect":{"foo":"3510565857353876841","bar":"4578135596460776054"}}}
{"level":"info","ts":1580641179.1443121,"caller":"logrus/json_parallele_bench.go:93","msg":"575269020497523263","Zap":"{\"foo\":\"3510565857353876841\",\"bar\":\"4578135596460776054\"}\n"}
{"level":"info","ts":1580641179.1443222,"caller":"logrus/json_parallele_bench.go:96","msg":"575269020497523263","ZapSugar":"{\"foo\":\"3510565857353876841\",\"bar\":\"4578135596460776054\"}\n"}
{"level":"info","Zerolog":{"foo":"3510565857353876841","bar":"4578135596460776054"},"time":"2020-02-02T11:59:39+01:00","message":"575269020497523263"}
{"Logrus":{"foo":"3510565857353876841","bar":"4578135596460776054"},"level":"info","msg":"575269020497523263","time":"2020-02-02T11:59:39+01:00"}
--------------------
results:
Zerolog: 34514 ns per request
Logger.Reflect: 36495 ns per request
Logger.Stringer: 36834 ns per request
ZapSugar: 44282 ns per request
Zap: 44636 ns per request
Logrus: 45052 ns per request

>Total duration: 45.053029345s
_____________________
---------------------
```
