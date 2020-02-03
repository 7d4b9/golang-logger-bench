# Golang logger benchmark

Generated README.md
> Makefile : bench json_bench parallele_bench json_parallele_bench bench_plot json_bench_plot parallele_bench_plot json_parallele_bench_plot

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
{"Message":"Logger: 470855719618197021","Level":6,"Timestamp":1580764055.128,"Context":null}
{"level":"info","msg":"Logrus: 470855719618197021","time":"2020-02-03T22:07:35+01:00"}
{"level":"info","ts":1580764055.12814,"caller":"logrus/bench.go:66","msg":"Zap: 470855719618197021"}
{"level":"info","ts":1580764055.128246,"caller":"logrus/bench.go:69","msg":"ZapSugar: 470855719618197021"}
{"level":"debug","time":"2020-02-03T22:07:35+01:00","message":"Zerolog: 470855719618197021"}
--------------------
results:
Logger: 10255 ns per request
Logrus: 60273 ns per request
Zap: 26250 ns per request
ZapSugar: 15064 ns per request
Zerolog: 8482 ns per request

>Total duration: 1.247724ms
_____________________
---------------------
TEST: bench 100
sample:
{"Message":"Logger: 3051736053718105278","Level":6,"Timestamp":1580764061.076,"Context":null}
{"level":"info","msg":"Logrus: 3051736053718105278","time":"2020-02-03T22:07:41+01:00"}
{"level":"info","ts":1580764061.076351,"caller":"logrus/bench.go:66","msg":"Zap: 3051736053718105278"}
{"level":"info","ts":1580764061.076457,"caller":"logrus/bench.go:69","msg":"ZapSugar: 3051736053718105278"}
{"level":"debug","time":"2020-02-03T22:07:41+01:00","message":"Zerolog: 3051736053718105278"}
--------------------
results:
Logger: 3741 ns per request
Logrus: 14792 ns per request
Zap: 12090 ns per request
ZapSugar: 7899 ns per request
Zerolog: 3280 ns per request

>Total duration: 4.256569ms
_____________________
---------------------
TEST: bench 1000
sample:
{"Message":"Logger: 812498502843828865","Level":6,"Timestamp":1580764066.136,"Context":null}
{"level":"info","msg":"Logrus: 812498502843828865","time":"2020-02-03T22:07:46+01:00"}
{"level":"info","ts":1580764066.1360588,"caller":"logrus/bench.go:66","msg":"Zap: 812498502843828865"}
{"level":"info","ts":1580764066.136238,"caller":"logrus/bench.go:69","msg":"ZapSugar: 812498502843828865"}
{"level":"debug","time":"2020-02-03T22:07:46+01:00","message":"Zerolog: 812498502843828865"}
--------------------
results:
Logger: 8984 ns per request
Logrus: 22619 ns per request
Zap: 11091 ns per request
ZapSugar: 15790 ns per request
Zerolog: 8862 ns per request

>Total duration: 67.501924ms
_____________________
---------------------
TEST: bench 10000
sample:
{"Message":"Logger: 8728139176410008644","Level":6,"Timestamp":1580764071.705,"Context":null}
{"level":"info","msg":"Logrus: 8728139176410008644","time":"2020-02-03T22:07:51+01:00"}
{"level":"info","ts":1580764071.705748,"caller":"logrus/bench.go:66","msg":"Zap: 8728139176410008644"}
{"level":"info","ts":1580764071.705861,"caller":"logrus/bench.go:69","msg":"ZapSugar: 8728139176410008644"}
{"level":"debug","time":"2020-02-03T22:07:51+01:00","message":"Zerolog: 8728139176410008644"}
--------------------
results:
Logger: 4796 ns per request
Logrus: 13456 ns per request
Zap: 9120 ns per request
ZapSugar: 13660 ns per request
Zerolog: 6410 ns per request

>Total duration: 474.525962ms
_____________________
---------------------
TEST: bench 100000
sample:
{"Message":"Logger: 2424893247991498469","Level":6,"Timestamp":1580764077.500,"Context":null}
{"level":"info","msg":"Logrus: 2424893247991498469","time":"2020-02-03T22:07:57+01:00"}
{"level":"info","ts":1580764077.500213,"caller":"logrus/bench.go:66","msg":"Zap: 2424893247991498469"}
{"level":"info","ts":1580764077.500332,"caller":"logrus/bench.go:69","msg":"ZapSugar: 2424893247991498469"}
{"level":"debug","time":"2020-02-03T22:07:57+01:00","message":"Zerolog: 2424893247991498469"}
--------------------
results:
Logger: 3988 ns per request
Logrus: 10212 ns per request
Zap: 6805 ns per request
ZapSugar: 7552 ns per request
Zerolog: 4033 ns per request

>Total duration: 3.259249807s
_____________________
---------------------
TEST: bench 1000000
sample:
{"Message":"Logger: 2841058939136120207","Level":6,"Timestamp":1580764086.283,"Context":null}
{"level":"info","msg":"Logrus: 2841058939136120207","time":"2020-02-03T22:08:06+01:00"}
{"level":"info","ts":1580764086.283672,"caller":"logrus/bench.go:66","msg":"Zap: 2841058939136120207"}
{"level":"info","ts":1580764086.2837849,"caller":"logrus/bench.go:69","msg":"ZapSugar: 2841058939136120207"}
{"level":"debug","time":"2020-02-03T22:08:06+01:00","message":"Zerolog: 2841058939136120207"}
--------------------
results:
Logger: 3561 ns per request
Logrus: 9380 ns per request
Zap: 6176 ns per request
ZapSugar: 7606 ns per request
Zerolog: 3518 ns per request

>Total duration: 30.242853185s
_____________________
---------------------
TEST: json_bench 10
sample:
{"Message":"7067303541522145647","Level":6,"Timestamp":1580764121.883,"Context":{"Logger":{"foo":"5787682404069524562","bar":"5311237587566724722"}}}
{"Logrus":{"foo":"5787682404069524562","bar":"5311237587566724722"},"level":"info","msg":"7067303541522145647","time":"2020-02-03T22:08:41+01:00"}
{"level":"info","ts":1580764121.884046,"caller":"logrus/json_bench.go:87","msg":"7067303541522145647","Zap":"{\"foo\":\"5787682404069524562\",\"bar\":\"5311237587566724722\"}\n"}
{"level":"info","ts":1580764121.884203,"caller":"logrus/json_bench.go:90","msg":"7067303541522145647","ZapSugar":"{\"foo\":\"5787682404069524562\",\"bar\":\"5311237587566724722\"}\n"}
{"level":"info","Zerolog":{"foo":"5787682404069524562","bar":"5311237587566724722"},"time":"2020-02-03T22:08:41+01:00","message":"7067303541522145647"}
--------------------
results:
Logger: 20726 ns per request
Logrus: 83508 ns per request
Zap: 32716 ns per request
ZapSugar: 17902 ns per request
Zerolog: 8585 ns per request

>Total duration: 1.674976ms
_____________________
---------------------
TEST: json_bench 100
sample:
{"Message":"1859378811898604688","Level":6,"Timestamp":1580764127.134,"Context":{"Logger":{"foo":"8761392782797303595","bar":"3704110018911162411"}}}
{"Logrus":{"foo":"8761392782797303595","bar":"3704110018911162411"},"level":"info","msg":"1859378811898604688","time":"2020-02-03T22:08:47+01:00"}
{"level":"info","ts":1580764127.135167,"caller":"logrus/json_bench.go:87","msg":"1859378811898604688","Zap":"{\"foo\":\"8761392782797303595\",\"bar\":\"3704110018911162411\"}\n"}
{"level":"info","ts":1580764127.1353111,"caller":"logrus/json_bench.go:90","msg":"1859378811898604688","ZapSugar":"{\"foo\":\"8761392782797303595\",\"bar\":\"3704110018911162411\"}\n"}
{"level":"info","Zerolog":{"foo":"8761392782797303595","bar":"3704110018911162411"},"time":"2020-02-03T22:08:47+01:00","message":"1859378811898604688"}
--------------------
results:
Logger: 8530 ns per request
Logrus: 20563 ns per request
Zap: 27489 ns per request
ZapSugar: 14878 ns per request
Zerolog: 4949 ns per request

>Total duration: 7.694092ms
_____________________
---------------------
TEST: json_bench 1000
sample:
{"Message":"4509817481686854449","Level":6,"Timestamp":1580764133.020,"Context":{"Logger":{"foo":"3322435699634591995","bar":"9094544119449398021"}}}
{"Logrus":{"foo":"3322435699634591995","bar":"9094544119449398021"},"level":"info","msg":"4509817481686854449","time":"2020-02-03T22:08:53+01:00"}
{"level":"info","ts":1580764133.020205,"caller":"logrus/json_bench.go:87","msg":"4509817481686854449","Zap":"{\"foo\":\"3322435699634591995\",\"bar\":\"9094544119449398021\"}\n"}
{"level":"info","ts":1580764133.020339,"caller":"logrus/json_bench.go:90","msg":"4509817481686854449","ZapSugar":"{\"foo\":\"3322435699634591995\",\"bar\":\"9094544119449398021\"}\n"}
{"level":"info","Zerolog":{"foo":"3322435699634591995","bar":"9094544119449398021"},"time":"2020-02-03T22:08:53+01:00","message":"4509817481686854449"}
--------------------
results:
Logger: 7347 ns per request
Logrus: 16037 ns per request
Zap: 13797 ns per request
ZapSugar: 29394 ns per request
Zerolog: 8180 ns per request

>Total duration: 74.841149ms
_____________________
---------------------
TEST: json_bench 10000
sample:
{"Message":"3182707266869963125","Level":6,"Timestamp":1580764138.681,"Context":{"Logger":{"foo":"2439590277576200300","bar":"1515538595311647621"}}}
{"Logrus":{"foo":"2439590277576200300","bar":"1515538595311647621"},"level":"info","msg":"3182707266869963125","time":"2020-02-03T22:08:58+01:00"}
{"level":"info","ts":1580764138.6816692,"caller":"logrus/json_bench.go:87","msg":"3182707266869963125","Zap":"{\"foo\":\"2439590277576200300\",\"bar\":\"1515538595311647621\"}\n"}
{"level":"info","ts":1580764138.681835,"caller":"logrus/json_bench.go:90","msg":"3182707266869963125","ZapSugar":"{\"foo\":\"2439590277576200300\",\"bar\":\"1515538595311647621\"}\n"}
{"level":"info","Zerolog":{"foo":"2439590277576200300","bar":"1515538595311647621"},"time":"2020-02-03T22:08:58+01:00","message":"3182707266869963125"}
--------------------
results:
Logger: 6745 ns per request
Logrus: 18049 ns per request
Zap: 11745 ns per request
ZapSugar: 14369 ns per request
Zerolog: 5737 ns per request

>Total duration: 566.549687ms
_____________________
---------------------
TEST: json_bench 100000
sample:
{"Message":"6004347388371117070","Level":6,"Timestamp":1580764144.746,"Context":{"Logger":{"foo":"6611497243221392460","bar":"6826586993795679647"}}}
{"Logrus":{"foo":"6611497243221392460","bar":"6826586993795679647"},"level":"info","msg":"6004347388371117070","time":"2020-02-03T22:09:04+01:00"}
{"level":"info","ts":1580764144.746148,"caller":"logrus/json_bench.go:87","msg":"6004347388371117070","Zap":"{\"foo\":\"6611497243221392460\",\"bar\":\"6826586993795679647\"}\n"}
{"level":"info","ts":1580764144.7462778,"caller":"logrus/json_bench.go:90","msg":"6004347388371117070","ZapSugar":"{\"foo\":\"6611497243221392460\",\"bar\":\"6826586993795679647\"}\n"}
{"level":"info","Zerolog":{"foo":"6611497243221392460","bar":"6826586993795679647"},"time":"2020-02-03T22:09:04+01:00","message":"6004347388371117070"}
--------------------
results:
Logger: 6693 ns per request
Logrus: 17383 ns per request
Zap: 10845 ns per request
ZapSugar: 13747 ns per request
Zerolog: 9459 ns per request

>Total duration: 5.813730786s
_____________________
---------------------
TEST: json_bench 1000000
sample:
{"Message":"8356612384658304685","Level":6,"Timestamp":1580764157.889,"Context":{"Logger":{"foo":"7188088639782562933","bar":"7475555167924573273"}}}
{"Logrus":{"foo":"7188088639782562933","bar":"7475555167924573273"},"level":"info","msg":"8356612384658304685","time":"2020-02-03T22:09:17+01:00"}
{"level":"info","ts":1580764157.88973,"caller":"logrus/json_bench.go:87","msg":"8356612384658304685","Zap":"{\"foo\":\"7188088639782562933\",\"bar\":\"7475555167924573273\"}\n"}
{"level":"info","ts":1580764157.889995,"caller":"logrus/json_bench.go:90","msg":"8356612384658304685","ZapSugar":"{\"foo\":\"7188088639782562933\",\"bar\":\"7475555167924573273\"}\n"}
{"level":"info","Zerolog":{"foo":"7188088639782562933","bar":"7475555167924573273"},"time":"2020-02-03T22:09:17+01:00","message":"8356612384658304685"}
--------------------
results:
Logger: 14976 ns per request
Logrus: 16109 ns per request
Zap: 11429 ns per request
ZapSugar: 12778 ns per request
Zerolog: 5126 ns per request

>Total duration: 1m0.419584359s
_____________________
---------------------
TEST: parallele_bench 10
sample:
{"Message":"Logger: 7927227360091622579","Level":6,"Timestamp":1580764225.247,"Context":null}
{"level":"info","ts":1580764225.24716,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 7927227360091622579"}
{"level":"info","ts":1580764225.247345,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 7927227360091622579"}
{"level":"debug","time":"2020-02-03T22:10:25+01:00","message":"Zerolog: 7927227360091622579"}
{"level":"info","msg":"Logrus: 7927227360091622579","time":"2020-02-03T22:10:25+01:00"}
--------------------
results:
Logger: 23322 ns per request
Zerolog: 59883 ns per request
ZapSugar: 38497 ns per request
Zap: 74941 ns per request
Logrus: 119984 ns per request

>Total duration: 1.283766ms
_____________________
---------------------
TEST: parallele_bench 100
sample:
{"Message":"Logger: 5400991969687165558","Level":6,"Timestamp":1580764230.616,"Context":null}
{"level":"info","ts":1580764230.6161149,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 5400991969687165558"}
{"level":"debug","time":"2020-02-03T22:10:30+01:00","message":"Zerolog: 5400991969687165558"}
{"level":"info","ts":1580764230.6163242,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 5400991969687165558"}
{"level":"info","msg":"Logrus: 5400991969687165558","time":"2020-02-03T22:10:30+01:00"}
--------------------
results:
Logger: 20286 ns per request
Zerolog: 30988 ns per request
ZapSugar: 46584 ns per request
Zap: 49995 ns per request
Logrus: 51945 ns per request

>Total duration: 5.320451ms
_____________________
---------------------
TEST: parallele_bench 1000
sample:
{"level":"info","ts":1580764236.268882,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 8266688679188292887"}
{"level":"debug","time":"2020-02-03T22:10:36+01:00","message":"Zerolog: 8266688679188292887"}
{"level":"info","ts":1580764236.269034,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 8266688679188292887"}
{"Message":"Logger: 8266688679188292887","Level":6,"Timestamp":1580764236.270,"Context":null}
{"level":"info","msg":"Logrus: 8266688679188292887","time":"2020-02-03T22:10:36+01:00"}
--------------------
results:
Logger: 34014 ns per request
Zerolog: 35653 ns per request
ZapSugar: 43048 ns per request
Zap: 45205 ns per request
Logrus: 48202 ns per request

>Total duration: 48.777153ms
_____________________
---------------------
TEST: parallele_bench 10000
sample:
{"Message":"Logger: 2815457077089264696","Level":6,"Timestamp":1580764242.265,"Context":null}
{"level":"info","ts":1580764242.2649212,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 2815457077089264696"}
{"level":"info","ts":1580764242.264899,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 2815457077089264696"}
{"level":"debug","time":"2020-02-03T22:10:42+01:00","message":"Zerolog: 2815457077089264696"}
{"level":"info","msg":"Logrus: 2815457077089264696","time":"2020-02-03T22:10:42+01:00"}
--------------------
results:
Logger: 35897 ns per request
Zerolog: 38876 ns per request
Zap: 44406 ns per request
ZapSugar: 46085 ns per request
Logrus: 48746 ns per request

>Total duration: 487.5645ms
_____________________
---------------------
TEST: parallele_bench 100000
sample:
{"Message":"Logger: 5475115292792446681","Level":6,"Timestamp":1580764248.125,"Context":null}
{"level":"info","ts":1580764248.125154,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 5475115292792446681"}
{"level":"info","ts":1580764248.125175,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 5475115292792446681"}
{"level":"debug","time":"2020-02-03T22:10:48+01:00","message":"Zerolog: 5475115292792446681"}
{"level":"info","msg":"Logrus: 5475115292792446681","time":"2020-02-03T22:10:48+01:00"}
--------------------
results:
Zerolog: 37418 ns per request
Logger: 39352 ns per request
ZapSugar: 45724 ns per request
Zap: 46149 ns per request
Logrus: 47924 ns per request

>Total duration: 4.792469099s
_____________________
---------------------
TEST: parallele_bench 1000000
sample:
{"Message":"Logger: 6999596966697032630","Level":6,"Timestamp":1580764259.260,"Context":null}
{"level":"info","ts":1580764259.259726,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 6999596966697032630"}
{"level":"info","ts":1580764259.2595682,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 6999596966697032630"}
{"level":"debug","time":"2020-02-03T22:10:59+01:00","message":"Zerolog: 6999596966697032630"}
{"level":"info","msg":"Logrus: 6999596966697032630","time":"2020-02-03T22:10:59+01:00"}
--------------------
results:
Zerolog: 37442 ns per request
Logger: 38884 ns per request
ZapSugar: 44577 ns per request
Zap: 45154 ns per request
Logrus: 46632 ns per request

>Total duration: 46.632758304s
_____________________
---------------------
TEST: json_parallele_bench 10
sample:
{"Message":"3213675607215033465","Level":6,"Timestamp":1580764312.526,"Context":{"Logger":{"foo":"539249446392763611","bar":"4330998641837815193"}}}
{"level":"info","ts":1580764312.5260391,"caller":"logrus/json_parallele_bench.go:90","msg":"3213675607215033465","Zap":"{\"foo\":\"539249446392763611\",\"bar\":\"4330998641837815193\"}\n"}
{"level":"info","ts":1580764312.5263019,"caller":"logrus/json_parallele_bench.go:93","msg":"3213675607215033465","ZapSugar":"{\"foo\":\"539249446392763611\",\"bar\":\"4330998641837815193\"}\n"}
{"level":"info","Zerolog":{"foo":"539249446392763611","bar":"4330998641837815193"},"time":"2020-02-03T22:11:52+01:00","message":"3213675607215033465"}
{"Logrus":{"foo":"539249446392763611","bar":"4330998641837815193"},"level":"info","msg":"3213675607215033465","time":"2020-02-03T22:11:52+01:00"}
--------------------
results:
Zap: 99476 ns per request
Logger: 110066 ns per request
Zerolog: 119537 ns per request
Logrus: 139276 ns per request
ZapSugar: 129299 ns per request

>Total duration: 1.497661ms
_____________________
---------------------
TEST: json_parallele_bench 100
sample:
{"level":"info","ts":1580764318.744614,"caller":"logrus/json_parallele_bench.go:90","msg":"6116450065998793618","Zap":"{\"foo\":\"3201064686323441548\",\"bar\":\"462468983286599794\"}\n"}
{"Logrus":{"foo":"3201064686323441548","bar":"462468983286599794"},"level":"info","msg":"6116450065998793618","time":"2020-02-03T22:11:58+01:00"}
{"level":"info","Zerolog":{"foo":"3201064686323441548","bar":"462468983286599794"},"time":"2020-02-03T22:11:58+01:00","message":"6116450065998793618"}
{"level":"info","ts":1580764318.7450082,"caller":"logrus/json_parallele_bench.go:93","msg":"6116450065998793618","ZapSugar":"{\"foo\":\"3201064686323441548\",\"bar\":\"462468983286599794\"}\n"}
{"Message":"6116450065998793618","Level":6,"Timestamp":1580764318.746,"Context":{"Logger":{"foo":"3201064686323441548","bar":"462468983286599794"}}}
--------------------
results:
Zerolog: 57116 ns per request
Zap: 67423 ns per request
Logrus: 73181 ns per request
Logger: 73351 ns per request
ZapSugar: 70719 ns per request

>Total duration: 7.848362ms
_____________________
---------------------
TEST: json_parallele_bench 1000
sample:
{"level":"info","ts":1580764324.988273,"caller":"logrus/json_parallele_bench.go:93","msg":"635742205378987571","ZapSugar":"{\"foo\":\"7255036059103365733\",\"bar\":\"8077510712641395243\"}\n"}
{"level":"info","ts":1580764324.988318,"caller":"logrus/json_parallele_bench.go:90","msg":"635742205378987571","Zap":"{\"foo\":\"7255036059103365733\",\"bar\":\"8077510712641395243\"}\n"}
{"level":"info","Zerolog":{"foo":"7255036059103365733","bar":"8077510712641395243"},"time":"2020-02-03T22:12:04+01:00","message":"635742205378987571"}
{"Message":"635742205378987571","Level":6,"Timestamp":1580764324.989,"Context":{"Logger":{"foo":"7255036059103365733","bar":"8077510712641395243"}}}
{"Logrus":{"foo":"7255036059103365733","bar":"8077510712641395243"},"level":"info","msg":"635742205378987571","time":"2020-02-03T22:12:04+01:00"}
--------------------
results:
Zerolog: 47968 ns per request
Logger: 51035 ns per request
Logrus: 68408 ns per request
Zap: 69845 ns per request
ZapSugar: 71071 ns per request

>Total duration: 71.575528ms
_____________________
---------------------
TEST: json_parallele_bench 10000
sample:
{"level":"info","ts":1580764331.3837922,"caller":"logrus/json_parallele_bench.go:90","msg":"8457729495265972778","Zap":"{\"foo\":\"4947976740066988069\",\"bar\":\"751996309793918460\"}\n"}
{"Message":"8457729495265972778","Level":6,"Timestamp":1580764331.384,"Context":{"Logger":{"foo":"4947976740066988069","bar":"751996309793918460"}}}
{"level":"info","ts":1580764331.383658,"caller":"logrus/json_parallele_bench.go:93","msg":"8457729495265972778","ZapSugar":"{\"foo\":\"4947976740066988069\",\"bar\":\"751996309793918460\"}\n"}
{"level":"info","Zerolog":{"foo":"4947976740066988069","bar":"751996309793918460"},"time":"2020-02-03T22:12:11+01:00","message":"8457729495265972778"}
{"Logrus":{"foo":"4947976740066988069","bar":"751996309793918460"},"level":"info","msg":"8457729495265972778","time":"2020-02-03T22:12:11+01:00"}
--------------------
results:
Zerolog: 46326 ns per request
Logger: 51242 ns per request
Zap: 63368 ns per request
ZapSugar: 64909 ns per request
Logrus: 72648 ns per request

>Total duration: 727.276382ms
_____________________
---------------------
TEST: json_parallele_bench 100000
sample:
{"Message":"8120408174395297448","Level":6,"Timestamp":1580764338.091,"Context":{"Logger":{"foo":"2898406726615832230","bar":"2801804268354163194"}}}
{"level":"info","ts":1580764338.0911932,"caller":"logrus/json_parallele_bench.go:93","msg":"8120408174395297448","ZapSugar":"{\"foo\":\"2898406726615832230\",\"bar\":\"2801804268354163194\"}\n"}
{"level":"info","ts":1580764338.0910149,"caller":"logrus/json_parallele_bench.go:90","msg":"8120408174395297448","Zap":"{\"foo\":\"2898406726615832230\",\"bar\":\"2801804268354163194\"}\n"}
{"level":"info","Zerolog":{"foo":"2898406726615832230","bar":"2801804268354163194"},"time":"2020-02-03T22:12:18+01:00","message":"8120408174395297448"}
{"Logrus":{"foo":"2898406726615832230","bar":"2801804268354163194"},"level":"info","msg":"8120408174395297448","time":"2020-02-03T22:12:18+01:00"}
--------------------
results:
Zerolog: 46489 ns per request
Logger: 48245 ns per request
ZapSugar: 57384 ns per request
Logrus: 58583 ns per request
Zap: 59415 ns per request

>Total duration: 5.941678899s
_____________________
---------------------
TEST: json_parallele_bench 1000000
sample:
{"level":"info","ts":1580764350.447042,"caller":"logrus/json_parallele_bench.go:93","msg":"4373790944994367586","ZapSugar":"{\"foo\":\"422616754546682887\",\"bar\":\"6565927336806332111\"}\n"}
{"level":"info","ts":1580764350.44703,"caller":"logrus/json_parallele_bench.go:90","msg":"4373790944994367586","Zap":"{\"foo\":\"422616754546682887\",\"bar\":\"6565927336806332111\"}\n"}
{"Message":"4373790944994367586","Level":6,"Timestamp":1580764350.448,"Context":{"Logger":{"foo":"422616754546682887","bar":"6565927336806332111"}}}
{"Logrus":{"foo":"422616754546682887","bar":"6565927336806332111"},"level":"info","msg":"4373790944994367586","time":"2020-02-03T22:12:30+01:00"}
{"level":"info","Zerolog":{"foo":"422616754546682887","bar":"6565927336806332111"},"time":"2020-02-03T22:12:30+01:00","message":"4373790944994367586"}
--------------------
results:
Zerolog: 53411 ns per request
Logger: 54356 ns per request
ZapSugar: 63026 ns per request
Logrus: 64472 ns per request
Zap: 64641 ns per request

>Total duration: 1m4.64178239s
_____________________
---------------------
