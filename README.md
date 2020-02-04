# Golang logger benchmark

Generated README.md
> Makefile : bench burst_bench json_bench burst_json_bench parallele_bench burst_parallele_bench json_parallele_bench burst_json_parallele_bench bench_plot burst_bench_plot json_bench_plot burst_json_bench_plot parallele_bench_plot burst_parallele_bench_plot json_parallele_bench_plot burst_json_parallele_bench_plot

This benchmark introduces a conscious bias. A considered fair bias considering UX.
This should be interpreted as modulo the "most user-friendly call".

Each library also have the benefit/penality of logging its own name. It's considered fair in this benchmark.

This benchmark does not intend to cover the advanced usages of the different loggers, just basics:

* `logger.Info(msg, loggerlog.Ctx("Logger", dummy))`
* `logrus.WithField("Logrus", dummy).Info(msg)`
* `zap.Info(msg, zap.Stringer("Zap", dummy))`
* `zap.Sugar().Infow(msg, "ZapSugar", dummy)`
* `zerolog.Info().Interface("Zerolog", dummy).Msg (msg)`

This is considered fair as some libraries doesn't support disabling the
addition of a call stack inside there logs while using the standard jsonEncoder by example.
This is a fair UX criteria not to forget in this benchmark.

```sh
make all-plot
```

```sh
TEST: bench 10
sample:
{"Message":"Logger: 830578489689995947","Level":6,"Timestamp":1580803191.475,"Context":null}
{"level":"info","msg":"Logrus: 830578489689995947","time":"2020-02-04T08:59:51+01:00"}
{"level":"info","ts":1580803191.475771,"caller":"logrus/bench.go:66","msg":"Zap: 830578489689995947"}
{"level":"info","ts":1580803191.475858,"caller":"logrus/bench.go:69","msg":"ZapSugar: 830578489689995947"}
{"level":"debug","time":"2020-02-04T08:59:51+01:00","message":"Zerolog: 830578489689995947"}
--------------------
results:
Logger: 5406 ns per request
Logrus: 36240 ns per request
Zap: 18383 ns per request
ZapSugar: 14489 ns per request
Zerolog: 7863 ns per request

>Total duration: 857.071µs
_____________________
---------------------
TEST: bench 100
sample:
{"Message":"Logger: 7899417380333218364","Level":6,"Timestamp":1580803194.784,"Context":null}
{"level":"info","msg":"Logrus: 7899417380333218364","time":"2020-02-04T08:59:54+01:00"}
{"level":"info","ts":1580803194.783882,"caller":"logrus/bench.go:66","msg":"Zap: 7899417380333218364"}
{"level":"info","ts":1580803194.783977,"caller":"logrus/bench.go:69","msg":"ZapSugar: 7899417380333218364"}
{"level":"debug","time":"2020-02-04T08:59:54+01:00","message":"Zerolog: 7899417380333218364"}
--------------------
results:
Logger: 3387 ns per request
Logrus: 10276 ns per request
Zap: 7425 ns per request
ZapSugar: 7239 ns per request
Zerolog: 3316 ns per request

>Total duration: 3.198151ms
_____________________
---------------------
TEST: bench 1000
sample:
{"Message":"Logger: 4635908850881362380","Level":6,"Timestamp":1580803197.999,"Context":null}
{"level":"info","msg":"Logrus: 4635908850881362380","time":"2020-02-04T08:59:57+01:00"}
{"level":"info","ts":1580803197.9988842,"caller":"logrus/bench.go:66","msg":"Zap: 4635908850881362380"}
{"level":"info","ts":1580803197.9989781,"caller":"logrus/bench.go:69","msg":"ZapSugar: 4635908850881362380"}
{"level":"debug","time":"2020-02-04T08:59:57+01:00","message":"Zerolog: 4635908850881362380"}
--------------------
results:
Logger: 2495 ns per request
Logrus: 8235 ns per request
Zap: 7607 ns per request
ZapSugar: 14156 ns per request
Zerolog: 4404 ns per request

>Total duration: 37.007595ms
_____________________
---------------------
TEST: bench 10000
sample:
{"Message":"Logger: 6401009507432232634","Level":6,"Timestamp":1580803201.359,"Context":null}
{"level":"info","msg":"Logrus: 6401009507432232634","time":"2020-02-04T09:00:01+01:00"}
{"level":"info","ts":1580803201.359564,"caller":"logrus/bench.go:66","msg":"Zap: 6401009507432232634"}
{"level":"info","ts":1580803201.359652,"caller":"logrus/bench.go:69","msg":"ZapSugar: 6401009507432232634"}
{"level":"debug","time":"2020-02-04T09:00:01+01:00","message":"Zerolog: 6401009507432232634"}
--------------------
results:
Logger: 2591 ns per request
Logrus: 8093 ns per request
Zap: 6169 ns per request
ZapSugar: 6493 ns per request
Zerolog: 3486 ns per request

>Total duration: 268.41729ms
_____________________
---------------------
TEST: bench 100000
sample:
{"Message":"Logger: 8316538580727726880","Level":6,"Timestamp":1580803205.026,"Context":null}
{"level":"info","msg":"Logrus: 8316538580727726880","time":"2020-02-04T09:00:05+01:00"}
{"level":"info","ts":1580803205.026262,"caller":"logrus/bench.go:66","msg":"Zap: 8316538580727726880"}
{"level":"info","ts":1580803205.026349,"caller":"logrus/bench.go:69","msg":"ZapSugar: 8316538580727726880"}
{"level":"debug","time":"2020-02-04T09:00:05+01:00","message":"Zerolog: 8316538580727726880"}
--------------------
results:
Logger: 2413 ns per request
Logrus: 7306 ns per request
Zap: 4237 ns per request
ZapSugar: 4730 ns per request
Zerolog: 4041 ns per request

>Total duration: 2.27298506s
_____________________
---------------------
TEST: bench 1000000
sample:
{"Message":"Logger: 2858248458070622657","Level":6,"Timestamp":1580803211.455,"Context":null}
{"level":"info","msg":"Logrus: 2858248458070622657","time":"2020-02-04T09:00:11+01:00"}
{"level":"info","ts":1580803211.455721,"caller":"logrus/bench.go:66","msg":"Zap: 2858248458070622657"}
{"level":"info","ts":1580803211.4558342,"caller":"logrus/bench.go:69","msg":"ZapSugar: 2858248458070622657"}
{"level":"debug","time":"2020-02-04T09:00:11+01:00","message":"Zerolog: 2858248458070622657"}
--------------------
results:
Logger: 4148 ns per request
Logrus: 8055 ns per request
Zap: 5546 ns per request
ZapSugar: 7455 ns per request
Zerolog: 4347 ns per request

>Total duration: 29.552864s
_____________________
---------------------
TEST: burst_bench 10
sample:
{"Message":"Logger: 7012122857076182561","Level":6,"Timestamp":1580803246.773,"Context":null}
{"level":"info","msg":"Logrus: 7012122857076182561","time":"2020-02-04T09:00:46+01:00"}
{"level":"info","ts":1580803246.774511,"caller":"logrus/burst_bench.go:74","msg":"Zap: 7012122857076182561"}
{"level":"info","ts":1580803246.775213,"caller":"logrus/burst_bench.go:77","msg":"ZapSugar: 7012122857076182561"}
{"level":"debug","time":"2020-02-04T09:00:46+01:00","message":"Zerolog: 7012122857076182561"}
--------------------
results:
Logger: 61415 ns per request
Logrus: 101361 ns per request
Zap: 71250 ns per request
ZapSugar: 69634 ns per request
Zerolog: 69151 ns per request

>Total duration: 3.790261ms
_____________________
---------------------
TEST: burst_bench 100
sample:
{"Message":"Logger: 5948115836595637092","Level":6,"Timestamp":1580803251.043,"Context":null}
{"level":"info","msg":"Logrus: 5948115836595637092","time":"2020-02-04T09:00:51+01:00"}
{"level":"info","ts":1580803251.043745,"caller":"logrus/burst_bench.go:74","msg":"Zap: 5948115836595637092"}
{"level":"info","ts":1580803251.0447202,"caller":"logrus/burst_bench.go:77","msg":"ZapSugar: 5948115836595637092"}
{"level":"debug","time":"2020-02-04T09:00:51+01:00","message":"Zerolog: 5948115836595637092"}
--------------------
results:
Logger: 10939 ns per request
Logrus: 37421 ns per request
Zap: 14766 ns per request
ZapSugar: 14081 ns per request
Zerolog: 12076 ns per request

>Total duration: 9.063292ms
_____________________
---------------------
TEST: burst_bench 1000
sample:
{"Message":"Logger: 7790391007109879882","Level":6,"Timestamp":1580803255.023,"Context":null}
{"level":"info","msg":"Logrus: 7790391007109879882","time":"2020-02-04T09:00:55+01:00"}
{"level":"info","ts":1580803255.0247831,"caller":"logrus/burst_bench.go:74","msg":"Zap: 7790391007109879882"}
{"level":"info","ts":1580803255.0255182,"caller":"logrus/burst_bench.go:77","msg":"ZapSugar: 7790391007109879882"}
{"level":"debug","time":"2020-02-04T09:00:55+01:00","message":"Zerolog: 7790391007109879882"}
--------------------
results:
Logger: 5141 ns per request
Logrus: 36678 ns per request
Zap: 6820 ns per request
ZapSugar: 13138 ns per request
Zerolog: 8393 ns per request

>Total duration: 70.455803ms
_____________________
---------------------
TEST: burst_bench 10000
sample:
{"Message":"Logger: 4526820632292169691","Level":6,"Timestamp":1580803258.767,"Context":null}
{"level":"info","msg":"Logrus: 4526820632292169691","time":"2020-02-04T09:00:58+01:00"}
{"level":"info","ts":1580803258.7685812,"caller":"logrus/burst_bench.go:74","msg":"Zap: 4526820632292169691"}
{"level":"info","ts":1580803258.769071,"caller":"logrus/burst_bench.go:77","msg":"ZapSugar: 4526820632292169691"}
{"level":"debug","time":"2020-02-04T09:00:58+01:00","message":"Zerolog: 4526820632292169691"}
--------------------
results:
Logger: 7282 ns per request
Logrus: 25445 ns per request
Zap: 9153 ns per request
ZapSugar: 7797 ns per request
Zerolog: 6012 ns per request

>Total duration: 557.867671ms
_____________________
---------------------
TEST: burst_bench 100000
sample:
{"Message":"Logger: 6296607132726060105","Level":6,"Timestamp":1580803262.917,"Context":null}
{"level":"info","msg":"Logrus: 6296607132726060105","time":"2020-02-04T09:01:02+01:00"}
{"level":"info","ts":1580803262.918098,"caller":"logrus/burst_bench.go:74","msg":"Zap: 6296607132726060105"}
{"level":"info","ts":1580803262.918742,"caller":"logrus/burst_bench.go:77","msg":"ZapSugar: 6296607132726060105"}
{"level":"debug","time":"2020-02-04T09:01:02+01:00","message":"Zerolog: 6296607132726060105"}
--------------------
results:
Logger: 5261 ns per request
Logrus: 26306 ns per request
Zap: 6023 ns per request
ZapSugar: 6359 ns per request
Zerolog: 4385 ns per request

>Total duration: 4.833834961s
_____________________
---------------------
TEST: burst_bench 1000000
sample:
{"Message":"Logger: 6805207991721480818","Level":6,"Timestamp":1580803271.559,"Context":null}
{"level":"info","msg":"Logrus: 6805207991721480818","time":"2020-02-04T09:01:11+01:00"}
{"level":"info","ts":1580803271.560467,"caller":"logrus/burst_bench.go:74","msg":"Zap: 6805207991721480818"}
{"level":"info","ts":1580803271.56124,"caller":"logrus/burst_bench.go:77","msg":"ZapSugar: 6805207991721480818"}
{"level":"debug","time":"2020-02-04T09:01:11+01:00","message":"Zerolog: 6805207991721480818"}
--------------------
results:
Logger: 4305 ns per request
Logrus: 25714 ns per request
Zap: 4881 ns per request
ZapSugar: 5640 ns per request
Zerolog: 3844 ns per request

>Total duration: 44.387688199s
_____________________
---------------------
TEST: json_bench 10
sample:
{"Message":"5283214376117843306","Level":6,"Timestamp":1580803319.978,"Context":{"Logger":{"foo":"1070365298610024604","bar":"5128499292799961676"}}}
{"Logrus":{"foo":"1070365298610024604","bar":"5128499292799961676"},"level":"info","msg":"5283214376117843306","time":"2020-02-04T09:01:59+01:00"}
{"level":"info","ts":1580803319.978488,"caller":"logrus/json_bench.go:87","msg":"5283214376117843306","Zap":"{\"foo\":\"1070365298610024604\",\"bar\":\"5128499292799961676\"}\n"}
{"level":"info","ts":1580803319.978595,"caller":"logrus/json_bench.go:90","msg":"5283214376117843306","ZapSugar":"{\"foo\":\"1070365298610024604\",\"bar\":\"5128499292799961676\"}\n"}
{"level":"info","Zerolog":{"foo":"1070365298610024604","bar":"5128499292799961676"},"time":"2020-02-04T09:01:59+01:00","message":"5283214376117843306"}
--------------------
results:
Logger: 20014 ns per request
Logrus: 36462 ns per request
Zap: 24591 ns per request
ZapSugar: 23191 ns per request
Zerolog: 7935 ns per request

>Total duration: 1.161575ms
_____________________
---------------------
TEST: json_bench 100
sample:
{"Message":"7052382950402538072","Level":6,"Timestamp":1580803323.570,"Context":{"Logger":{"foo":"7058768294071258557","bar":"1115121960138211449"}}}
{"Logrus":{"foo":"7058768294071258557","bar":"1115121960138211449"},"level":"info","msg":"7052382950402538072","time":"2020-02-04T09:02:03+01:00"}
{"level":"info","ts":1580803323.5700579,"caller":"logrus/json_bench.go:87","msg":"7052382950402538072","Zap":"{\"foo\":\"7058768294071258557\",\"bar\":\"1115121960138211449\"}\n"}
{"level":"info","ts":1580803323.5701828,"caller":"logrus/json_bench.go:90","msg":"7052382950402538072","ZapSugar":"{\"foo\":\"7058768294071258557\",\"bar\":\"1115121960138211449\"}\n"}
{"level":"info","Zerolog":{"foo":"7058768294071258557","bar":"1115121960138211449"},"time":"2020-02-04T09:02:03+01:00","message":"7052382950402538072"}
--------------------
results:
Logger: 7514 ns per request
Logrus: 13661 ns per request
Zap: 12933 ns per request
ZapSugar: 11228 ns per request
Zerolog: 4005 ns per request

>Total duration: 4.973859ms
_____________________
---------------------
TEST: json_bench 1000
sample:
{"Message":"8961797361046763198","Level":6,"Timestamp":1580803327.109,"Context":{"Logger":{"foo":"3161328539749888324","bar":"3772091577924314950"}}}
{"Logrus":{"foo":"3161328539749888324","bar":"3772091577924314950"},"level":"info","msg":"8961797361046763198","time":"2020-02-04T09:02:07+01:00"}
{"level":"info","ts":1580803327.1093671,"caller":"logrus/json_bench.go:87","msg":"8961797361046763198","Zap":"{\"foo\":\"3161328539749888324\",\"bar\":\"3772091577924314950\"}\n"}
{"level":"info","ts":1580803327.109486,"caller":"logrus/json_bench.go:90","msg":"8961797361046763198","ZapSugar":"{\"foo\":\"3161328539749888324\",\"bar\":\"3772091577924314950\"}\n"}
{"level":"info","Zerolog":{"foo":"3161328539749888324","bar":"3772091577924314950"},"time":"2020-02-04T09:02:07+01:00","message":"8961797361046763198"}
--------------------
results:
Logger: 6247 ns per request
Logrus: 17418 ns per request
Zap: 11198 ns per request
ZapSugar: 16082 ns per request
Zerolog: 5167 ns per request

>Total duration: 56.278035ms
_____________________
---------------------
TEST: json_bench 10000
sample:
{"Message":"5626419769269438031","Level":6,"Timestamp":1580803330.714,"Context":{"Logger":{"foo":"4354847356322698109","bar":"5488891864822756458"}}}
{"Logrus":{"foo":"4354847356322698109","bar":"5488891864822756458"},"level":"info","msg":"5626419769269438031","time":"2020-02-04T09:02:10+01:00"}
{"level":"info","ts":1580803330.714426,"caller":"logrus/json_bench.go:87","msg":"5626419769269438031","Zap":"{\"foo\":\"4354847356322698109\",\"bar\":\"5488891864822756458\"}\n"}
{"level":"info","ts":1580803330.7145119,"caller":"logrus/json_bench.go:90","msg":"5626419769269438031","ZapSugar":"{\"foo\":\"4354847356322698109\",\"bar\":\"5488891864822756458\"}\n"}
{"level":"info","Zerolog":{"foo":"4354847356322698109","bar":"5488891864822756458"},"time":"2020-02-04T09:02:10+01:00","message":"5626419769269438031"}
--------------------
results:
Logger: 5530 ns per request
Logrus: 11021 ns per request
Zap: 9010 ns per request
ZapSugar: 10237 ns per request
Zerolog: 4003 ns per request

>Total duration: 398.120795ms
_____________________
---------------------
TEST: json_bench 100000
sample:
{"Message":"8621911908547641021","Level":6,"Timestamp":1580803334.921,"Context":{"Logger":{"foo":"421291888734744325","bar":"8262984496092421075"}}}
{"Logrus":{"foo":"421291888734744325","bar":"8262984496092421075"},"level":"info","msg":"8621911908547641021","time":"2020-02-04T09:02:14+01:00"}
{"level":"info","ts":1580803334.921416,"caller":"logrus/json_bench.go:87","msg":"8621911908547641021","Zap":"{\"foo\":\"421291888734744325\",\"bar\":\"8262984496092421075\"}\n"}
{"level":"info","ts":1580803334.921538,"caller":"logrus/json_bench.go:90","msg":"8621911908547641021","ZapSugar":"{\"foo\":\"421291888734744325\",\"bar\":\"8262984496092421075\"}\n"}
{"level":"info","Zerolog":{"foo":"421291888734744325","bar":"8262984496092421075"},"time":"2020-02-04T09:02:14+01:00","message":"8621911908547641021"}
--------------------
results:
Logger: 5280 ns per request
Logrus: 10021 ns per request
Zap: 8009 ns per request
ZapSugar: 8962 ns per request
Zerolog: 3543 ns per request

>Total duration: 3.581886616s
_____________________
---------------------
TEST: json_bench 1000000
sample:
{"Message":"3114066517732179985","Level":6,"Timestamp":1580803342.137,"Context":{"Logger":{"foo":"1960936633738194982","bar":"6946391176883919181"}}}
{"Logrus":{"foo":"1960936633738194982","bar":"6946391176883919181"},"level":"info","msg":"3114066517732179985","time":"2020-02-04T09:02:22+01:00"}
{"level":"info","ts":1580803342.1369128,"caller":"logrus/json_bench.go:87","msg":"3114066517732179985","Zap":"{\"foo\":\"1960936633738194982\",\"bar\":\"6946391176883919181\"}\n"}
{"level":"info","ts":1580803342.137036,"caller":"logrus/json_bench.go:90","msg":"3114066517732179985","ZapSugar":"{\"foo\":\"1960936633738194982\",\"bar\":\"6946391176883919181\"}\n"}
{"level":"info","Zerolog":{"foo":"1960936633738194982","bar":"6946391176883919181"},"time":"2020-02-04T09:02:22+01:00","message":"3114066517732179985"}
--------------------
results:
Logger: 5297 ns per request
Logrus: 9951 ns per request
Zap: 7893 ns per request
ZapSugar: 10286 ns per request
Zerolog: 4269 ns per request

>Total duration: 37.699473409s
_____________________
---------------------
TEST: burst_json_bench 10
sample:
{"Message":"6945027614566471655","Level":6,"Timestamp":1580803384.327,"Context":{"Logger":{"foo":"7314391177768420210","bar":"5256069732696711628"}}}
{"Logrus":{"foo":"7314391177768420210","bar":"5256069732696711628"},"level":"info","msg":"6945027614566471655","time":"2020-02-04T09:03:04+01:00"}
{"level":"info","ts":1580803384.327805,"caller":"logrus/burst_json_bench.go:95","msg":"6945027614566471655","Zap":"{\"foo\":\"7314391177768420210\",\"bar\":\"5256069732696711628\"}\n"}
{"level":"info","ts":1580803384.328408,"caller":"logrus/burst_json_bench.go:98","msg":"6945027614566471655","ZapSugar":"{\"foo\":\"7314391177768420210\",\"bar\":\"5256069732696711628\"}\n"}
{"level":"info","Zerolog":{"foo":"7314391177768420210","bar":"5256069732696711628"},"time":"2020-02-04T09:03:04+01:00","message":"6945027614566471655"}
--------------------
results:
Logger: 74853 ns per request
Logrus: 96540 ns per request
Zap: 72302 ns per request
ZapSugar: 80320 ns per request
Zerolog: 59869 ns per request

>Total duration: 3.931207ms
_____________________
---------------------
TEST: burst_json_bench 100
sample:
{"Message":"8851791325569881173","Level":6,"Timestamp":1580803388.280,"Context":{"Logger":{"foo":"3451203371616141010","bar":"474148338560469017"}}}
{"Logrus":{"foo":"3451203371616141010","bar":"474148338560469017"},"level":"info","msg":"8851791325569881173","time":"2020-02-04T09:03:08+01:00"}
{"level":"info","ts":1580803388.2815301,"caller":"logrus/burst_json_bench.go:95","msg":"8851791325569881173","Zap":"{\"foo\":\"3451203371616141010\",\"bar\":\"474148338560469017\"}\n"}
{"level":"info","ts":1580803388.282485,"caller":"logrus/burst_json_bench.go:98","msg":"8851791325569881173","ZapSugar":"{\"foo\":\"3451203371616141010\",\"bar\":\"474148338560469017\"}\n"}
{"level":"info","Zerolog":{"foo":"3451203371616141010","bar":"474148338560469017"},"time":"2020-02-04T09:03:08+01:00","message":"8851791325569881173"}
--------------------
results:
Logger: 13292 ns per request
Logrus: 35871 ns per request
Zap: 16533 ns per request
ZapSugar: 45565 ns per request
Zerolog: 14755 ns per request

>Total duration: 12.751369ms
_____________________
---------------------
TEST: burst_json_bench 1000
sample:
{"Message":"2639750199502742843","Level":6,"Timestamp":1580803392.067,"Context":{"Logger":{"foo":"8725054823586061395","bar":"3234593949396262646"}}}
{"Logrus":{"foo":"8725054823586061395","bar":"3234593949396262646"},"level":"info","msg":"2639750199502742843","time":"2020-02-04T09:03:12+01:00"}
{"level":"info","ts":1580803392.068388,"caller":"logrus/burst_json_bench.go:95","msg":"2639750199502742843","Zap":"{\"foo\":\"8725054823586061395\",\"bar\":\"3234593949396262646\"}\n"}
{"level":"info","ts":1580803392.068983,"caller":"logrus/burst_json_bench.go:98","msg":"2639750199502742843","ZapSugar":"{\"foo\":\"8725054823586061395\",\"bar\":\"3234593949396262646\"}\n"}
{"level":"info","Zerolog":{"foo":"8725054823586061395","bar":"3234593949396262646"},"time":"2020-02-04T09:03:12+01:00","message":"2639750199502742843"}
--------------------
results:
Logger: 6988 ns per request
Logrus: 37987 ns per request
Zap: 11381 ns per request
ZapSugar: 27245 ns per request
Zerolog: 9898 ns per request

>Total duration: 93.628901ms
_____________________
---------------------
TEST: burst_json_bench 10000
sample:
{"Message":"8511436930961940684","Level":6,"Timestamp":1580803395.894,"Context":{"Logger":{"foo":"1568200645076255955","bar":"280473137991347242"}}}
{"Logrus":{"foo":"1568200645076255955","bar":"280473137991347242"},"level":"info","msg":"8511436930961940684","time":"2020-02-04T09:03:15+01:00"}
{"level":"info","ts":1580803395.89468,"caller":"logrus/burst_json_bench.go:95","msg":"8511436930961940684","Zap":"{\"foo\":\"1568200645076255955\",\"bar\":\"280473137991347242\"}\n"}
{"level":"info","ts":1580803395.895234,"caller":"logrus/burst_json_bench.go:98","msg":"8511436930961940684","ZapSugar":"{\"foo\":\"1568200645076255955\",\"bar\":\"280473137991347242\"}\n"}
{"level":"info","Zerolog":{"foo":"1568200645076255955","bar":"280473137991347242"},"time":"2020-02-04T09:03:15+01:00","message":"8511436930961940684"}
--------------------
results:
Logger: 9189 ns per request
Logrus: 27750 ns per request
Zap: 8328 ns per request
ZapSugar: 24373 ns per request
Zerolog: 6246 ns per request

>Total duration: 758.997224ms
_____________________
---------------------
TEST: burst_json_bench 100000
sample:
{"Message":"6030821340924911119","Level":6,"Timestamp":1580803400.308,"Context":{"Logger":{"foo":"1022994903273051460","bar":"1927072687754001579"}}}
{"Logrus":{"foo":"1022994903273051460","bar":"1927072687754001579"},"level":"info","msg":"6030821340924911119","time":"2020-02-04T09:03:20+01:00"}
{"level":"info","ts":1580803400.30944,"caller":"logrus/burst_json_bench.go:95","msg":"6030821340924911119","Zap":"{\"foo\":\"1022994903273051460\",\"bar\":\"1927072687754001579\"}\n"}
{"level":"info","ts":1580803400.3100019,"caller":"logrus/burst_json_bench.go:98","msg":"6030821340924911119","ZapSugar":"{\"foo\":\"1022994903273051460\",\"bar\":\"1927072687754001579\"}\n"}
{"level":"info","Zerolog":{"foo":"1022994903273051460","bar":"1927072687754001579"},"time":"2020-02-04T09:03:20+01:00","message":"6030821340924911119"}
--------------------
results:
Logger: 7668 ns per request
Logrus: 32432 ns per request
Zap: 8678 ns per request
ZapSugar: 23611 ns per request
Zerolog: 4862 ns per request

>Total duration: 7.725488647s
_____________________
---------------------
TEST: burst_json_bench 1000000
sample:
{"Message":"7527495019834895092","Level":6,"Timestamp":1580803411.952,"Context":{"Logger":{"foo":"3912792941688477261","bar":"9060708812903372617"}}}
{"Logrus":{"foo":"3912792941688477261","bar":"9060708812903372617"},"level":"info","msg":"7527495019834895092","time":"2020-02-04T09:03:31+01:00"}
{"level":"info","ts":1580803411.952996,"caller":"logrus/burst_json_bench.go:95","msg":"7527495019834895092","Zap":"{\"foo\":\"3912792941688477261\",\"bar\":\"9060708812903372617\"}\n"}
{"level":"info","ts":1580803411.953629,"caller":"logrus/burst_json_bench.go:98","msg":"7527495019834895092","ZapSugar":"{\"foo\":\"3912792941688477261\",\"bar\":\"9060708812903372617\"}\n"}
{"level":"info","Zerolog":{"foo":"3912792941688477261","bar":"9060708812903372617"},"time":"2020-02-04T09:03:31+01:00","message":"7527495019834895092"}
--------------------
results:
Logger: 6075 ns per request
Logrus: 32950 ns per request
Zap: 7890 ns per request
ZapSugar: 21918 ns per request
Zerolog: 4837 ns per request

>Total duration: 1m13.670992759s
_____________________
---------------------
TEST: parallele_bench 10
sample:
{"Message":"Logger: 8592567625004535262","Level":6,"Timestamp":1580803490.076,"Context":null}
{"level":"info","ts":1580803490.075626,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 8592567625004535262"}
{"level":"debug","time":"2020-02-04T09:04:50+01:00","message":"Zerolog: 8592567625004535262"}
{"level":"info","ts":1580803490.075798,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 8592567625004535262"}
{"level":"info","msg":"Logrus: 8592567625004535262","time":"2020-02-04T09:04:50+01:00"}
--------------------
results:
Logger: 20049 ns per request
Zap: 41993 ns per request
Zerolog: 53799 ns per request
ZapSugar: 45681 ns per request
Logrus: 74616 ns per request

>Total duration: 813.73µs
_____________________
---------------------
TEST: parallele_bench 100
sample:
{"Message":"Logger: 5331221227307209071","Level":6,"Timestamp":1580803493.676,"Context":null}
{"level":"info","ts":1580803493.676228,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 5331221227307209071"}
{"level":"info","ts":1580803493.6762009,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 5331221227307209071"}
{"level":"debug","time":"2020-02-04T09:04:53+01:00","message":"Zerolog: 5331221227307209071"}
{"level":"info","msg":"Logrus: 5331221227307209071","time":"2020-02-04T09:04:53+01:00"}
--------------------
results:
Zerolog: 14119 ns per request
Logrus: 26808 ns per request
ZapSugar: 28497 ns per request
Logger: 28763 ns per request
Zap: 35719 ns per request

>Total duration: 3.634515ms
_____________________
---------------------
TEST: parallele_bench 1000
sample:
{"Message":"Logger: 7173577491414600157","Level":6,"Timestamp":1580803497.344,"Context":null}
{"level":"info","ts":1580803497.3443341,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 7173577491414600157"}
{"level":"info","ts":1580803497.3443189,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 7173577491414600157"}
{"level":"debug","time":"2020-02-04T09:04:57+01:00","message":"Zerolog: 7173577491414600157"}
{"level":"info","msg":"Logrus: 7173577491414600157","time":"2020-02-04T09:04:57+01:00"}
--------------------
results:
Logger: 29607 ns per request
Zerolog: 29808 ns per request
ZapSugar: 33720 ns per request
Zap: 35028 ns per request
Logrus: 37351 ns per request

>Total duration: 37.424478ms
_____________________
---------------------
TEST: parallele_bench 10000
sample:
{"Message":"Logger: 9010250686390591043","Level":6,"Timestamp":1580803501.033,"Context":null}
{"level":"info","ts":1580803501.032895,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 9010250686390591043"}
{"level":"debug","time":"2020-02-04T09:05:01+01:00","message":"Zerolog: 9010250686390591043"}
{"level":"info","ts":1580803501.0329049,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 9010250686390591043"}
{"level":"info","msg":"Logrus: 9010250686390591043","time":"2020-02-04T09:05:01+01:00"}
--------------------
results:
Logger: 25541 ns per request
Zap: 29423 ns per request
Zerolog: 31171 ns per request
ZapSugar: 33199 ns per request
Logrus: 36466 ns per request

>Total duration: 364.73604ms
_____________________
---------------------
TEST: parallele_bench 100000
sample:
{"Message":"Logger: 2779792397377100593","Level":6,"Timestamp":1580803505.019,"Context":null}
{"level":"info","ts":1580803505.018708,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 2779792397377100593"}
{"level":"debug","time":"2020-02-04T09:05:05+01:00","message":"Zerolog: 2779792397377100593"}
{"level":"info","ts":1580803505.0190709,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 2779792397377100593"}
{"level":"info","msg":"Logrus: 2779792397377100593","time":"2020-02-04T09:05:05+01:00"}
--------------------
results:
Zerolog: 25601 ns per request
Logger: 26079 ns per request
Zap: 28804 ns per request
ZapSugar: 30032 ns per request
Logrus: 31108 ns per request

>Total duration: 3.110900856s
_____________________
---------------------
TEST: parallele_bench 1000000
sample:
{"Message":"Logger: 5397658448497168842","Level":6,"Timestamp":1580803511.899,"Context":null}
{"level":"info","ts":1580803511.899143,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 5397658448497168842"}
{"level":"info","ts":1580803511.899321,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 5397658448497168842"}
{"level":"debug","time":"2020-02-04T09:05:11+01:00","message":"Zerolog: 5397658448497168842"}
{"level":"info","msg":"Logrus: 5397658448497168842","time":"2020-02-04T09:05:11+01:00"}
--------------------
results:
Zerolog: 27049 ns per request
Logger: 27319 ns per request
Zap: 31875 ns per request
ZapSugar: 32543 ns per request
Logrus: 33313 ns per request

>Total duration: 33.313985756s
_____________________
---------------------
TEST: burst_parallele_bench 10
sample:
{"level":"debug","time":"2020-02-04T09:05:49+01:00","message":"Zerolog: 7460017127044346173"}
{"Message":"Logger: 7460017127044346173","Level":6,"Timestamp":1580803549.196,"Context":null}
{"level":"info","ts":1580803549.195528,"caller":"logrus/burst_parallele_bench.go:78","msg":"Zap: 7460017127044346173"}
{"level":"info","ts":1580803549.195702,"caller":"logrus/burst_parallele_bench.go:81","msg":"ZapSugar: 7460017127044346173"}
{"level":"info","msg":"Logrus: 7460017127044346173","time":"2020-02-04T09:05:49+01:00"}
--------------------
results:
Logger: 110399 ns per request
Zap: 137520 ns per request
Zerolog: 167389 ns per request
ZapSugar: 73399 ns per request
Logrus: 173904 ns per request

>Total duration: 2.166984ms
_____________________
---------------------
TEST: burst_parallele_bench 100
sample:
{"Message":"Logger: 4196264493281438798","Level":6,"Timestamp":1580803552.889,"Context":null}
{"level":"info","ts":1580803552.889023,"caller":"logrus/burst_parallele_bench.go:78","msg":"Zap: 4196264493281438798"}
{"level":"debug","time":"2020-02-04T09:05:52+01:00","message":"Zerolog: 4196264493281438798"}
{"level":"info","ts":1580803552.8892841,"caller":"logrus/burst_parallele_bench.go:81","msg":"ZapSugar: 4196264493281438798"}
{"level":"info","msg":"Logrus: 4196264493281438798","time":"2020-02-04T09:05:52+01:00"}
--------------------
results:
Logger: 42205 ns per request
Zerolog: 46138 ns per request
ZapSugar: 39533 ns per request
Zap: 54112 ns per request
Logrus: 43858 ns per request

>Total duration: 6.83203ms
_____________________
---------------------
TEST: burst_parallele_bench 1000
sample:
{"Message":"Logger: 5968371273625419068","Level":6,"Timestamp":1580803556.569,"Context":null}
{"level":"info","ts":1580803556.568525,"caller":"logrus/burst_parallele_bench.go:78","msg":"Zap: 5968371273625419068"}
{"level":"info","ts":1580803556.568773,"caller":"logrus/burst_parallele_bench.go:81","msg":"ZapSugar: 5968371273625419068"}
{"level":"info","msg":"Logrus: 5968371273625419068","time":"2020-02-04T09:05:56+01:00"}
{"level":"debug","time":"2020-02-04T09:05:56+01:00","message":"Zerolog: 5968371273625419068"}
--------------------
results:
Zerolog: 53057 ns per request
Logger: 53413 ns per request
ZapSugar: 49268 ns per request
Zap: 60317 ns per request
Logrus: 78495 ns per request

>Total duration: 78.671795ms
_____________________
---------------------
TEST: burst_parallele_bench 10000
sample:
{"Message":"Logger: 9023497999184866722","Level":6,"Timestamp":1580803560.409,"Context":null}
{"level":"debug","time":"2020-02-04T09:06:00+01:00","message":"Zerolog: 9023497999184866722"}
{"level":"info","ts":1580803560.409238,"caller":"logrus/burst_parallele_bench.go:81","msg":"ZapSugar: 9023497999184866722"}
{"level":"info","ts":1580803560.409055,"caller":"logrus/burst_parallele_bench.go:78","msg":"Zap: 9023497999184866722"}
{"level":"info","msg":"Logrus: 9023497999184866722","time":"2020-02-04T09:06:00+01:00"}
--------------------
results:
Zerolog: 34832 ns per request
ZapSugar: 38740 ns per request
Logger: 40182 ns per request
Zap: 40189 ns per request
Logrus: 57551 ns per request

>Total duration: 576.038908ms
_____________________
---------------------
TEST: burst_parallele_bench 100000
sample:
{"Message":"Logger: 1610496247855233552","Level":6,"Timestamp":1580803564.587,"Context":null}
{"level":"debug","time":"2020-02-04T09:06:04+01:00","message":"Zerolog: 1610496247855233552"}
{"level":"info","ts":1580803564.586539,"caller":"logrus/burst_parallele_bench.go:78","msg":"Zap: 1610496247855233552"}
{"level":"info","ts":1580803564.586762,"caller":"logrus/burst_parallele_bench.go:81","msg":"ZapSugar: 1610496247855233552"}
{"level":"info","msg":"Logrus: 1610496247855233552","time":"2020-02-04T09:06:04+01:00"}
--------------------
results:
Zerolog: 16000 ns per request
Logger: 15996 ns per request
ZapSugar: 35046 ns per request
Zap: 36520 ns per request
Logrus: 48249 ns per request

>Total duration: 4.826633144s
_____________________
---------------------
TEST: burst_parallele_bench 1000000
sample:
{"level":"debug","time":"2020-02-04T09:06:13+01:00","message":"Zerolog: 1178701856559542585"}
{"Message":"Logger: 1178701856559542585","Level":6,"Timestamp":1580803573.106,"Context":null}
{"level":"info","ts":1580803573.106133,"caller":"logrus/burst_parallele_bench.go:81","msg":"ZapSugar: 1178701856559542585"}
{"level":"info","ts":1580803573.105953,"caller":"logrus/burst_parallele_bench.go:78","msg":"Zap: 1178701856559542585"}
{"level":"info","msg":"Logrus: 1178701856559542585","time":"2020-02-04T09:06:13+01:00"}
--------------------
results:
Logger: 9274 ns per request
Zerolog: 9459 ns per request
ZapSugar: 29686 ns per request
Zap: 29999 ns per request
Logrus: 42221 ns per request

>Total duration: 42.221673139s
_____________________
---------------------
TEST: json_parallele_bench 10
sample:
{"Message":"8105909129440206206","Level":6,"Timestamp":1580803619.485,"Context":{"Logger":{"foo":"5204513136237229151","bar":"1022444170871319651"}}}
{"level":"info","ts":1580803619.485198,"caller":"logrus/json_parallele_bench.go:90","msg":"8105909129440206206","Zap":"{\"foo\":\"5204513136237229151\",\"bar\":\"1022444170871319651\"}\n"}
{"level":"info","Zerolog":{"foo":"5204513136237229151","bar":"1022444170871319651"},"time":"2020-02-04T09:06:59+01:00","message":"8105909129440206206"}
{"Logrus":{"foo":"5204513136237229151","bar":"1022444170871319651"},"level":"info","msg":"8105909129440206206","time":"2020-02-04T09:06:59+01:00"}
{"level":"info","ts":1580803619.485448,"caller":"logrus/json_parallele_bench.go:93","msg":"8105909129440206206","ZapSugar":"{\"foo\":\"5204513136237229151\",\"bar\":\"1022444170871319651\"}\n"}
--------------------
results:
Logger: 42497 ns per request
Zap: 75316 ns per request
Zerolog: 90442 ns per request
Logrus: 91887 ns per request
ZapSugar: 72115 ns per request

>Total duration: 1.038464ms
_____________________
---------------------
TEST: json_parallele_bench 100
sample:
{"level":"info","ts":1580803652.54211,"caller":"logrus/json_parallele_bench.go:90","msg":"2066576595888828204","Zap":"{\"foo\":\"6113110706468453339\",\"bar\":\"4077683177464495369\"}\n"}
{"Message":"2066576595888828204","Level":6,"Timestamp":1580803652.542,"Context":{"Logger":{"foo":"6113110706468453339","bar":"4077683177464495369"}}}
{"level":"info","ts":1580803652.542294,"caller":"logrus/json_parallele_bench.go:93","msg":"2066576595888828204","ZapSugar":"{\"foo\":\"6113110706468453339\",\"bar\":\"4077683177464495369\"}\n"}
{"Logrus":{"foo":"6113110706468453339","bar":"4077683177464495369"},"level":"info","msg":"2066576595888828204","time":"2020-02-04T09:07:32+01:00"}
{"level":"info","Zerolog":{"foo":"6113110706468453339","bar":"4077683177464495369"},"time":"2020-02-04T09:07:32+01:00","message":"2066576595888828204"}
--------------------
results:
Zerolog: 32251 ns per request
Logger: 35441 ns per request
Zap: 40630 ns per request
Logrus: 43803 ns per request
ZapSugar: 42796 ns per request

>Total duration: 4.564567ms
_____________________
---------------------
TEST: json_parallele_bench 1000
sample:
{"level":"info","ts":1580803656.636359,"caller":"logrus/json_parallele_bench.go:90","msg":"3846083950304569234","Zap":"{\"foo\":\"2246166969926055771\",\"bar\":\"9078291574514180965\"}\n"}
{"level":"info","ts":1580803656.6371572,"caller":"logrus/json_parallele_bench.go:93","msg":"3846083950304569234","ZapSugar":"{\"foo\":\"2246166969926055771\",\"bar\":\"9078291574514180965\"}\n"}
{"Message":"3846083950304569234","Level":6,"Timestamp":1580803656.636,"Context":{"Logger":{"foo":"2246166969926055771","bar":"9078291574514180965"}}}
{"level":"info","Zerolog":{"foo":"2246166969926055771","bar":"9078291574514180965"},"time":"2020-02-04T09:07:36+01:00","message":"3846083950304569234"}
{"Logrus":{"foo":"2246166969926055771","bar":"9078291574514180965"},"level":"info","msg":"3846083950304569234","time":"2020-02-04T09:07:36+01:00"}
--------------------
results:
Zerolog: 27371 ns per request
Logger: 35999 ns per request
Zap: 37527 ns per request
ZapSugar: 40710 ns per request
Logrus: 42323 ns per request

>Total duration: 42.441288ms
_____________________
---------------------
TEST: json_parallele_bench 10000
sample:
{"Message":"6915004838842578048","Level":6,"Timestamp":1580803660.482,"Context":{"Logger":{"foo":"7516966390404920572","bar":"2636270078778034862"}}}
{"level":"info","ts":1580803660.4820101,"caller":"logrus/json_parallele_bench.go:90","msg":"6915004838842578048","Zap":"{\"foo\":\"7516966390404920572\",\"bar\":\"2636270078778034862\"}\n"}
{"level":"info","ts":1580803660.482195,"caller":"logrus/json_parallele_bench.go:93","msg":"6915004838842578048","ZapSugar":"{\"foo\":\"7516966390404920572\",\"bar\":\"2636270078778034862\"}\n"}
{"level":"info","Zerolog":{"foo":"7516966390404920572","bar":"2636270078778034862"},"time":"2020-02-04T09:07:40+01:00","message":"6915004838842578048"}
{"Logrus":{"foo":"7516966390404920572","bar":"2636270078778034862"},"level":"info","msg":"6915004838842578048","time":"2020-02-04T09:07:40+01:00"}
--------------------
results:
Zerolog: 29331 ns per request
Logger: 31280 ns per request
ZapSugar: 39708 ns per request
Zap: 39920 ns per request
Logrus: 41388 ns per request

>Total duration: 413.992202ms
_____________________
---------------------
TEST: json_parallele_bench 100000
sample:
{"level":"info","ts":1580803664.580865,"caller":"logrus/json_parallele_bench.go:93","msg":"8677322341749369062","ZapSugar":"{\"foo\":\"3690756395777593483\",\"bar\":\"7337956383581651355\"}\n"}
{"level":"info","ts":1580803664.580896,"caller":"logrus/json_parallele_bench.go:90","msg":"8677322341749369062","Zap":"{\"foo\":\"3690756395777593483\",\"bar\":\"7337956383581651355\"}\n"}
{"level":"info","Zerolog":{"foo":"3690756395777593483","bar":"7337956383581651355"},"time":"2020-02-04T09:07:44+01:00","message":"8677322341749369062"}
{"Message":"8677322341749369062","Level":6,"Timestamp":1580803664.581,"Context":{"Logger":{"foo":"3690756395777593483","bar":"7337956383581651355"}}}
{"Logrus":{"foo":"3690756395777593483","bar":"7337956383581651355"},"level":"info","msg":"8677322341749369062","time":"2020-02-04T09:07:44+01:00"}
--------------------
results:
Zerolog: 29846 ns per request
Logger: 32305 ns per request
Zap: 40839 ns per request
ZapSugar: 40931 ns per request
Logrus: 46019 ns per request

>Total duration: 4.602062875s
_____________________
---------------------
TEST: json_parallele_bench 1000000
sample:
{"Message":"9131318169390706583","Level":6,"Timestamp":1580803673.640,"Context":{"Logger":{"foo":"8236909805278618653","bar":"256868382541535432"}}}
{"level":"info","ts":1580803673.639942,"caller":"logrus/json_parallele_bench.go:90","msg":"9131318169390706583","Zap":"{\"foo\":\"8236909805278618653\",\"bar\":\"256868382541535432\"}\n"}
{"level":"info","ts":1580803673.639946,"caller":"logrus/json_parallele_bench.go:93","msg":"9131318169390706583","ZapSugar":"{\"foo\":\"8236909805278618653\",\"bar\":\"256868382541535432\"}\n"}
{"level":"info","Zerolog":{"foo":"8236909805278618653","bar":"256868382541535432"},"time":"2020-02-04T09:07:53+01:00","message":"9131318169390706583"}
{"Logrus":{"foo":"8236909805278618653","bar":"256868382541535432"},"level":"info","msg":"9131318169390706583","time":"2020-02-04T09:07:53+01:00"}
--------------------
results:
Zerolog: 29154 ns per request
Logger: 31633 ns per request
ZapSugar: 37107 ns per request
Zap: 38126 ns per request
Logrus: 38445 ns per request

>Total duration: 38.446060811s
_____________________
---------------------
TEST: burst_json_parallele_bench 10
sample:
{"level":"info","ts":1580803717.808532,"caller":"logrus/burst_json_parallele_bench.go:98","msg":"4890149223483016737","ZapSugar":"{\"foo\":\"3006818645199260969\",\"bar\":\"1114848918198568120\"}\n"}
{"Logrus":{"foo":"3006818645199260969","bar":"1114848918198568120"},"level":"info","msg":"4890149223483016737","time":"2020-02-04T09:08:37+01:00"}
{"level":"info","Zerolog":{"foo":"3006818645199260969","bar":"1114848918198568120"},"time":"2020-02-04T09:08:37+01:00","message":"4890149223483016737"}
{"level":"info","ts":1580803717.808516,"caller":"logrus/burst_json_parallele_bench.go:95","msg":"4890149223483016737","Zap":"{\"foo\":\"3006818645199260969\",\"bar\":\"1114848918198568120\"}\n"}
{"Message":"4890149223483016737","Level":6,"Timestamp":1580803717.809,"Context":{"Logger":{"foo":"3006818645199260969","bar":"1114848918198568120"}}}
--------------------
results:
Zap: 104053 ns per request
Logger: 121118 ns per request
Zerolog: 157281 ns per request
ZapSugar: 115848 ns per request
Logrus: 102424 ns per request

>Total duration: 2.20511ms
_____________________
---------------------
TEST: burst_json_parallele_bench 100
sample:
{"Message":"6803733806585055887","Level":6,"Timestamp":1580803721.633,"Context":{"Logger":{"foo":"8356580605078383281","bar":"3894362141193927429"}}}
{"level":"info","Zerolog":{"foo":"8356580605078383281","bar":"3894362141193927429"},"time":"2020-02-04T09:08:41+01:00","message":"6803733806585055887"}
{"level":"info","ts":1580803721.6327622,"caller":"logrus/burst_json_parallele_bench.go:95","msg":"6803733806585055887","Zap":"{\"foo\":\"8356580605078383281\",\"bar\":\"3894362141193927429\"}\n"}
{"Logrus":{"foo":"8356580605078383281","bar":"3894362141193927429"},"level":"info","msg":"6803733806585055887","time":"2020-02-04T09:08:41+01:00"}
{"level":"info","ts":1580803721.633172,"caller":"logrus/burst_json_parallele_bench.go:98","msg":"6803733806585055887","ZapSugar":"{\"foo\":\"8356580605078383281\",\"bar\":\"3894362141193927429\"}\n"}
--------------------
results:
Zerolog: 37695 ns per request
Logger: 44469 ns per request
ZapSugar: 64326 ns per request
Zap: 71712 ns per request
Logrus: 77380 ns per request

>Total duration: 9.57114ms
_____________________
---------------------
TEST: burst_json_parallele_bench 1000
sample:
{"Message":"8568529625651102581","Level":6,"Timestamp":1580803725.253,"Context":{"Logger":{"foo":"4547780269294871953","bar":"9174417994459996783"}}}
{"level":"info","ts":1580803725.253386,"caller":"logrus/burst_json_parallele_bench.go:95","msg":"8568529625651102581","Zap":"{\"foo\":\"4547780269294871953\",\"bar\":\"9174417994459996783\"}\n"}
{"level":"info","ts":1580803725.253637,"caller":"logrus/burst_json_parallele_bench.go:98","msg":"8568529625651102581","ZapSugar":"{\"foo\":\"4547780269294871953\",\"bar\":\"9174417994459996783\"}\n"}
{"level":"info","Zerolog":{"foo":"4547780269294871953","bar":"9174417994459996783"},"time":"2020-02-04T09:08:45+01:00","message":"8568529625651102581"}
{"Logrus":{"foo":"4547780269294871953","bar":"9174417994459996783"},"level":"info","msg":"8568529625651102581","time":"2020-02-04T09:08:45+01:00"}
--------------------
results:
Zerolog: 72033 ns per request
Logger: 72670 ns per request
ZapSugar: 73402 ns per request
Zap: 74130 ns per request
Logrus: 64079 ns per request

>Total duration: 94.456914ms
_____________________
---------------------
TEST: burst_json_parallele_bench 10000
sample:
{"Message":"2088163991776841699","Level":6,"Timestamp":1580803729.016,"Context":{"Logger":{"foo":"685346454193709866","bar":"2583271878712146251"}}}
{"level":"info","Zerolog":{"foo":"685346454193709866","bar":"2583271878712146251"},"time":"2020-02-04T09:08:49+01:00","message":"2088163991776841699"}
{"level":"info","ts":1580803729.016018,"caller":"logrus/burst_json_parallele_bench.go:95","msg":"2088163991776841699","Zap":"{\"foo\":\"685346454193709866\",\"bar\":\"2583271878712146251\"}\n"}
{"Logrus":{"foo":"685346454193709866","bar":"2583271878712146251"},"level":"info","msg":"2088163991776841699","time":"2020-02-04T09:08:49+01:00"}
{"level":"info","ts":1580803729.016426,"caller":"logrus/burst_json_parallele_bench.go:98","msg":"2088163991776841699","ZapSugar":"{\"foo\":\"685346454193709866\",\"bar\":\"2583271878712146251\"}\n"}
--------------------
results:
Zerolog: 48467 ns per request
Logger: 52827 ns per request
ZapSugar: 61152 ns per request
Zap: 62548 ns per request
Logrus: 81915 ns per request

>Total duration: 819.313654ms
_____________________
---------------------
TEST: burst_json_parallele_bench 100000
sample:
{"level":"info","ts":1580803733.790849,"caller":"logrus/burst_json_parallele_bench.go:95","msg":"3924186224648295497","Zap":"{\"foo\":\"6028116635171577506\",\"bar\":\"8306332625045299096\"}\n"}
{"level":"info","Zerolog":{"foo":"6028116635171577506","bar":"8306332625045299096"},"time":"2020-02-04T09:08:53+01:00","message":"3924186224648295497"}
{"level":"info","ts":1580803733.791001,"caller":"logrus/burst_json_parallele_bench.go:98","msg":"3924186224648295497","ZapSugar":"{\"foo\":\"6028116635171577506\",\"bar\":\"8306332625045299096\"}\n"}
{"Message":"3924186224648295497","Level":6,"Timestamp":1580803733.791,"Context":{"Logger":{"foo":"6028116635171577506","bar":"8306332625045299096"}}}
{"Logrus":{"foo":"6028116635171577506","bar":"8306332625045299096"},"level":"info","msg":"3924186224648295497","time":"2020-02-04T09:08:53+01:00"}
--------------------
results:
Zerolog: 18809 ns per request
Logger: 20521 ns per request
Zap: 56511 ns per request
Logrus: 67614 ns per request
ZapSugar: 67448 ns per request

>Total duration: 6.78472409s
_____________________
---------------------
TEST: burst_json_parallele_bench 1000000
sample:
{"Message":"5497736081717594285","Level":6,"Timestamp":1580803744.477,"Context":{"Logger":{"foo":"9043163335128089091","bar":"4012040850115119794"}}}
{"level":"info","Zerolog":{"foo":"9043163335128089091","bar":"4012040850115119794"},"time":"2020-02-04T09:09:04+01:00","message":"5497736081717594285"}
{"level":"info","ts":1580803744.4771688,"caller":"logrus/burst_json_parallele_bench.go:95","msg":"5497736081717594285","Zap":"{\"foo\":\"9043163335128089091\",\"bar\":\"4012040850115119794\"}\n"}
{"Logrus":{"foo":"9043163335128089091","bar":"4012040850115119794"},"level":"info","msg":"5497736081717594285","time":"2020-02-04T09:09:04+01:00"}
{"level":"info","ts":1580803744.477549,"caller":"logrus/burst_json_parallele_bench.go:98","msg":"5497736081717594285","ZapSugar":"{\"foo\":\"9043163335128089091\",\"bar\":\"4012040850115119794\"}\n"}
--------------------
results:
Zerolog: 7915 ns per request
Logger: 12352 ns per request
Zap: 42624 ns per request
ZapSugar: 53928 ns per request
Logrus: 56130 ns per request

>Total duration: 56.131343361s
_____________________
---------------------
