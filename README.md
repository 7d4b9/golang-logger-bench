# Golang logger benchmark

Generated README.md
> Makefile : bench json_bench parallele_bench json_parallele_bench bench_plot json_bench_plot parallele_bench_plot json_parallele_bench_plot
```sh
make all-plot
```
```sh
TEST: bench 10
sample:
{"Message":"Logger: 5890378885747324507","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 5890378885747324507","time":"2020-02-02T21:10:12+01:00"}
{"level":"info","ts":1580674212.4639502,"caller":"logrus/bench.go:66","msg":"Zap: 5890378885747324507"}
{"level":"info","ts":1580674212.4640589,"caller":"logrus/bench.go:69","msg":"ZapSugar: 5890378885747324507"}
{"level":"debug","time":"2020-02-02T21:10:12+01:00","message":"Zerolog: 5890378885747324507"}
--------------------
results:
Logger: 3780 ns per request
Logrus: 36137 ns per request
Zap: 17366 ns per request
ZapSugar: 13268 ns per request
Zerolog: 6199 ns per request

>Total duration: 803.641µs
_____________________
---------------------
TEST: bench 100
sample:
{"Message":"Logger: 7798283046524105409","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 7798283046524105409","time":"2020-02-02T21:10:16+01:00"}
{"level":"info","ts":1580674216.201472,"caller":"logrus/bench.go:66","msg":"Zap: 7798283046524105409"}
{"level":"info","ts":1580674216.201559,"caller":"logrus/bench.go:69","msg":"ZapSugar: 7798283046524105409"}
{"level":"debug","time":"2020-02-02T21:10:16+01:00","message":"Zerolog: 7798283046524105409"}
--------------------
results:
Logger: 2276 ns per request
Logrus: 12346 ns per request
Zap: 7154 ns per request
ZapSugar: 6970 ns per request
Zerolog: 3035 ns per request

>Total duration: 3.214946ms
_____________________
---------------------
TEST: bench 1000
sample:
{"Message":"Logger: 4462660558308884946","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 4462660558308884946","time":"2020-02-02T21:10:19+01:00"}
{"level":"info","ts":1580674219.7868538,"caller":"logrus/bench.go:66","msg":"Zap: 4462660558308884946"}
{"level":"info","ts":1580674219.786955,"caller":"logrus/bench.go:69","msg":"ZapSugar: 4462660558308884946"}
{"level":"debug","time":"2020-02-02T21:10:19+01:00","message":"Zerolog: 4462660558308884946"}
--------------------
results:
Logger: 2527 ns per request
Logrus: 8842 ns per request
Zap: 7909 ns per request
ZapSugar: 9667 ns per request
Zerolog: 3908 ns per request

>Total duration: 32.945172ms
_____________________
---------------------
TEST: bench 10000
sample:
{"Message":"Logger: 6308861019563998912","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 6308861019563998912","time":"2020-02-02T21:10:23+01:00"}
{"level":"info","ts":1580674223.3007538,"caller":"logrus/bench.go:66","msg":"Zap: 6308861019563998912"}
{"level":"info","ts":1580674223.300915,"caller":"logrus/bench.go:69","msg":"ZapSugar: 6308861019563998912"}
{"level":"debug","time":"2020-02-02T21:10:23+01:00","message":"Zerolog: 6308861019563998912"}
--------------------
results:
Logger: 2052 ns per request
Logrus: 9317 ns per request
Zap: 6159 ns per request
ZapSugar: 6802 ns per request
Zerolog: 3377 ns per request

>Total duration: 277.22489ms
_____________________
---------------------
TEST: bench 100000
sample:
{"Message":"Logger: 72151183495642918","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 72151183495642918","time":"2020-02-02T21:10:27+01:00"}
{"level":"info","ts":1580674227.006654,"caller":"logrus/bench.go:66","msg":"Zap: 72151183495642918"}
{"level":"info","ts":1580674227.006764,"caller":"logrus/bench.go:69","msg":"ZapSugar: 72151183495642918"}
{"level":"debug","time":"2020-02-02T21:10:27+01:00","message":"Zerolog: 72151183495642918"}
--------------------
results:
Logger: 1791 ns per request
Logrus: 6676 ns per request
Zap: 4368 ns per request
ZapSugar: 5046 ns per request
Zerolog: 2432 ns per request

>Total duration: 2.031605798s
_____________________
---------------------
TEST: bench 1000000
sample:
{"Message":"Logger: 6778029491635371369","Level":6,"Context":null}
{"level":"info","msg":"Logrus: 6778029491635371369","time":"2020-02-02T21:10:32+01:00"}
{"level":"info","ts":1580674232.4122062,"caller":"logrus/bench.go:66","msg":"Zap: 6778029491635371369"}
{"level":"info","ts":1580674232.412321,"caller":"logrus/bench.go:69","msg":"ZapSugar: 6778029491635371369"}
{"level":"debug","time":"2020-02-02T21:10:32+01:00","message":"Zerolog: 6778029491635371369"}
--------------------
results:
Logger: 1791 ns per request
Logrus: 6250 ns per request
Zap: 4427 ns per request
ZapSugar: 5221 ns per request
Zerolog: 2263 ns per request

>Total duration: 19.954929725s
_____________________
---------------------
TEST: json_bench 10
sample:
{"Message":"1707905717885836517","Level":6,"Context":{"LoggerStringer":{"foo":"1436280639265108592","bar":"5751164836506978321"}}}
{"Message":"1707905717885836517","Level":6,"Context":{"LoggerReflect":{"foo":"1436280639265108592","bar":"5751164836506978321"}}}
{"Logrus":{"foo":"1436280639265108592","bar":"5751164836506978321"},"level":"info","msg":"1707905717885836517","time":"2020-02-02T21:10:56+01:00"}
{"level":"info","ts":1580674256.286928,"caller":"logrus/json_bench.go:90","msg":"1707905717885836517","Zap":"{\"foo\":\"1436280639265108592\",\"bar\":\"5751164836506978321\"}\n"}
{"level":"info","ts":1580674256.287046,"caller":"logrus/json_bench.go:93","msg":"1707905717885836517","ZapSugar":"{\"foo\":\"1436280639265108592\",\"bar\":\"5751164836506978321\"}\n"}
{"level":"info","Zerolog":{"foo":"1436280639265108592","bar":"5751164836506978321"},"time":"2020-02-02T21:10:56+01:00","message":"1707905717885836517"}
--------------------
results:
Logger.Stringer: 14609 ns per request
Logger.Reflect: 6224 ns per request
Logrus: 43838 ns per request
Zap: 24094 ns per request
ZapSugar: 17797 ns per request
Zerolog: 8131 ns per request

>Total duration: 1.196903ms
_____________________
---------------------
TEST: json_bench 100
sample:
{"Message":"7557114853766237174","Level":6,"Context":{"LoggerStringer":{"foo":"3299533695051225816","bar":"2853641807636878533"}}}
{"Message":"7557114853766237174","Level":6,"Context":{"LoggerReflect":{"foo":"3299533695051225816","bar":"2853641807636878533"}}}
{"Logrus":{"foo":"3299533695051225816","bar":"2853641807636878533"},"level":"info","msg":"7557114853766237174","time":"2020-02-02T21:10:59+01:00"}
{"level":"info","ts":1580674259.774403,"caller":"logrus/json_bench.go:90","msg":"7557114853766237174","Zap":"{\"foo\":\"3299533695051225816\",\"bar\":\"2853641807636878533\"}\n"}
{"level":"info","ts":1580674259.7745779,"caller":"logrus/json_bench.go:93","msg":"7557114853766237174","ZapSugar":"{\"foo\":\"3299533695051225816\",\"bar\":\"2853641807636878533\"}\n"}
{"level":"info","Zerolog":{"foo":"3299533695051225816","bar":"2853641807636878533"},"time":"2020-02-02T21:10:59+01:00","message":"7557114853766237174"}
--------------------
results:
Logger.Stringer: 5846 ns per request
Logger.Reflect: 9511 ns per request
Logrus: 23802 ns per request
Zap: 15397 ns per request
ZapSugar: 15641 ns per request
Zerolog: 3762 ns per request

>Total duration: 7.44528ms
_____________________
---------------------
TEST: json_bench 1000
sample:
{"Message":"212484738915965028","Level":6,"Context":{"LoggerStringer":{"foo":"8651047162177509489","bar":"7793421855493032978"}}}
{"Message":"212484738915965028","Level":6,"Context":{"LoggerReflect":{"foo":"8651047162177509489","bar":"7793421855493032978"}}}
{"Logrus":{"foo":"8651047162177509489","bar":"7793421855493032978"},"level":"info","msg":"212484738915965028","time":"2020-02-02T21:11:03+01:00"}
{"level":"info","ts":1580674263.404565,"caller":"logrus/json_bench.go:90","msg":"212484738915965028","Zap":"{\"foo\":\"8651047162177509489\",\"bar\":\"7793421855493032978\"}\n"}
{"level":"info","ts":1580674263.404686,"caller":"logrus/json_bench.go:93","msg":"212484738915965028","ZapSugar":"{\"foo\":\"8651047162177509489\",\"bar\":\"7793421855493032978\"}\n"}
{"level":"info","Zerolog":{"foo":"8651047162177509489","bar":"7793421855493032978"},"time":"2020-02-02T21:11:03+01:00","message":"212484738915965028"}
--------------------
results:
Logger.Stringer: 4366 ns per request
Logger.Reflect: 7132 ns per request
Logrus: 19966 ns per request
Zap: 17392 ns per request
ZapSugar: 22153 ns per request
Zerolog: 9190 ns per request

>Total duration: 80.486341ms
_____________________
---------------------
TEST: json_bench 10000
sample:
{"Message":"6138302571367307252","Level":6,"Context":{"LoggerStringer":{"foo":"1198291932243507417","bar":"5460298374201116855"}}}
{"Message":"6138302571367307252","Level":6,"Context":{"LoggerReflect":{"foo":"1198291932243507417","bar":"5460298374201116855"}}}
{"Logrus":{"foo":"1198291932243507417","bar":"5460298374201116855"},"level":"info","msg":"6138302571367307252","time":"2020-02-02T21:11:06+01:00"}
{"level":"info","ts":1580674266.9912162,"caller":"logrus/json_bench.go:90","msg":"6138302571367307252","Zap":"{\"foo\":\"1198291932243507417\",\"bar\":\"5460298374201116855\"}\n"}
{"level":"info","ts":1580674266.991327,"caller":"logrus/json_bench.go:93","msg":"6138302571367307252","ZapSugar":"{\"foo\":\"1198291932243507417\",\"bar\":\"5460298374201116855\"}\n"}
{"level":"info","Zerolog":{"foo":"1198291932243507417","bar":"5460298374201116855"},"time":"2020-02-02T21:11:06+01:00","message":"6138302571367307252"}
--------------------
results:
Logger.Stringer: 4136 ns per request
Logger.Reflect: 4575 ns per request
Logrus: 10273 ns per request
Zap: 7787 ns per request
ZapSugar: 8676 ns per request
Zerolog: 3712 ns per request

>Total duration: 391.686417ms
_____________________
---------------------
TEST: json_bench 100000
sample:
{"Message":"9125921484521186394","Level":6,"Context":{"LoggerStringer":{"foo":"6233080925150074417","bar":"1178222766094942179"}}}
{"Message":"9125921484521186394","Level":6,"Context":{"LoggerReflect":{"foo":"6233080925150074417","bar":"1178222766094942179"}}}
{"Logrus":{"foo":"6233080925150074417","bar":"1178222766094942179"},"level":"info","msg":"9125921484521186394","time":"2020-02-02T21:11:10+01:00"}
{"level":"info","ts":1580674270.7633998,"caller":"logrus/json_bench.go:90","msg":"9125921484521186394","Zap":"{\"foo\":\"6233080925150074417\",\"bar\":\"1178222766094942179\"}\n"}
{"level":"info","ts":1580674270.76353,"caller":"logrus/json_bench.go:93","msg":"9125921484521186394","ZapSugar":"{\"foo\":\"6233080925150074417\",\"bar\":\"1178222766094942179\"}\n"}
{"level":"info","Zerolog":{"foo":"6233080925150074417","bar":"1178222766094942179"},"time":"2020-02-02T21:11:10+01:00","message":"9125921484521186394"}
--------------------
results:
Logger.Stringer: 3800 ns per request
Logger.Reflect: 3781 ns per request
Logrus: 8600 ns per request
Zap: 6720 ns per request
ZapSugar: 7762 ns per request
Zerolog: 2992 ns per request

>Total duration: 3.365913729s
_____________________
---------------------
TEST: json_bench 1000000
sample:
{"Message":"7705830738835612881","Level":6,"Context":{"LoggerStringer":{"foo":"4787959592275172402","bar":"3364533772448084945"}}}
{"Message":"7705830738835612881","Level":6,"Context":{"LoggerReflect":{"foo":"4787959592275172402","bar":"3364533772448084945"}}}
{"Logrus":{"foo":"4787959592275172402","bar":"3364533772448084945"},"level":"info","msg":"7705830738835612881","time":"2020-02-02T21:11:17+01:00"}
{"level":"info","ts":1580674277.5939672,"caller":"logrus/json_bench.go:90","msg":"7705830738835612881","Zap":"{\"foo\":\"4787959592275172402\",\"bar\":\"3364533772448084945\"}\n"}
{"level":"info","ts":1580674277.5940878,"caller":"logrus/json_bench.go:93","msg":"7705830738835612881","ZapSugar":"{\"foo\":\"4787959592275172402\",\"bar\":\"3364533772448084945\"}\n"}
{"level":"info","Zerolog":{"foo":"4787959592275172402","bar":"3364533772448084945"},"time":"2020-02-02T21:11:17+01:00","message":"7705830738835612881"}
--------------------
results:
Logger.Stringer: 3901 ns per request
Logger.Reflect: 3742 ns per request
Logrus: 8463 ns per request
Zap: 6798 ns per request
ZapSugar: 7686 ns per request
Zerolog: 2943 ns per request

>Total duration: 33.536558834s
_____________________
---------------------
TEST: parallele_bench 10
sample:
{"Message":"Logger: 516977753968667713","Level":6,"Context":null}
{"level":"info","ts":1580674315.161668,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 516977753968667713"}
{"level":"info","ts":1580674315.161632,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 516977753968667713"}
{"level":"debug","time":"2020-02-02T21:11:55+01:00","message":"Zerolog: 516977753968667713"}
{"level":"info","msg":"Logrus: 516977753968667713","time":"2020-02-02T21:11:55+01:00"}
--------------------
results:
Logger: 23636 ns per request
Zerolog: 42933 ns per request
Logrus: 76171 ns per request
ZapSugar: 57963 ns per request
Zap: 75945 ns per request

>Total duration: 876.48µs
_____________________
---------------------
TEST: parallele_bench 100
sample:
{"Message":"Logger: 6388132633363446738","Level":6,"Context":null}
{"level":"info","ts":1580674318.581384,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 6388132633363446738"}
{"level":"info","ts":1580674318.5812259,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 6388132633363446738"}
{"level":"debug","time":"2020-02-02T21:11:58+01:00","message":"Zerolog: 6388132633363446738"}
{"level":"info","msg":"Logrus: 6388132633363446738","time":"2020-02-02T21:11:58+01:00"}
--------------------
results:
Zerolog: 26273 ns per request
Logger: 27952 ns per request
ZapSugar: 35069 ns per request
Zap: 37031 ns per request
Logrus: 41006 ns per request

>Total duration: 4.165957ms
_____________________
---------------------
TEST: parallele_bench 1000
sample:
{"Message":"Logger: 3125374805351984994","Level":6,"Context":null}
{"level":"info","ts":1580674321.96557,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 3125374805351984994"}
{"level":"debug","time":"2020-02-02T21:12:01+01:00","message":"Zerolog: 3125374805351984994"}
{"level":"info","ts":1580674321.965476,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 3125374805351984994"}
{"level":"info","msg":"Logrus: 3125374805351984994","time":"2020-02-02T21:12:01+01:00"}
--------------------
results:
Logger: 22211 ns per request
Zerolog: 27934 ns per request
Zap: 29527 ns per request
ZapSugar: 36910 ns per request
Logrus: 37139 ns per request

>Total duration: 37.217366ms
_____________________
---------------------
TEST: parallele_bench 10000
sample:
{"Message":"Logger: 6120936179082459088","Level":6,"Context":null}
{"level":"info","ts":1580674325.7278578,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 6120936179082459088"}
{"level":"info","ts":1580674325.727894,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 6120936179082459088"}
{"level":"debug","time":"2020-02-02T21:12:05+01:00","message":"Zerolog: 6120936179082459088"}
{"level":"info","msg":"Logrus: 6120936179082459088","time":"2020-02-02T21:12:05+01:00"}
--------------------
results:
Logger: 23407 ns per request
Zerolog: 25706 ns per request
Zap: 31174 ns per request
ZapSugar: 31492 ns per request
Logrus: 34124 ns per request

>Total duration: 341.346393ms
_____________________
---------------------
TEST: parallele_bench 100000
sample:
{"Message":"Logger: 7954127219749747766","Level":6,"Context":null}
{"level":"info","ts":1580674329.497781,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 7954127219749747766"}
{"level":"info","ts":1580674329.4977841,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 7954127219749747766"}
{"level":"debug","time":"2020-02-02T21:12:09+01:00","message":"Zerolog: 7954127219749747766"}
{"level":"info","msg":"Logrus: 7954127219749747766","time":"2020-02-02T21:12:09+01:00"}
--------------------
results:
Logger: 23894 ns per request
Zerolog: 24438 ns per request
ZapSugar: 28631 ns per request
Zap: 28830 ns per request
Logrus: 29532 ns per request

>Total duration: 2.953375889s
_____________________
---------------------
TEST: parallele_bench 1000000
sample:
{"Message":"Logger: 1314305131273253591","Level":6,"Context":null}
{"level":"info","ts":1580674335.8707762,"caller":"logrus/parallele_bench.go:67","msg":"Zap: 1314305131273253591"}
{"level":"info","ts":1580674335.87091,"caller":"logrus/parallele_bench.go:70","msg":"ZapSugar: 1314305131273253591"}
{"level":"debug","time":"2020-02-02T21:12:15+01:00","message":"Zerolog: 1314305131273253591"}
{"level":"info","msg":"Logrus: 1314305131273253591","time":"2020-02-02T21:12:15+01:00"}
--------------------
results:
Logger: 21073 ns per request
Zerolog: 22806 ns per request
Zap: 27498 ns per request
ZapSugar: 27773 ns per request
Logrus: 28564 ns per request

>Total duration: 28.564433278s
_____________________
---------------------
TEST: json_parallele_bench 10
sample:
{"Message":"5936303275916165373","Level":6,"Context":{"LoggerReflect":{"foo":"3421102674176290467","bar":"8376938394956482517"}}}
{"Message":"5936303275916165373","Level":6,"Context":{"LoggerStringer":{"foo":"3421102674176290467","bar":"8376938394956482517"}}}
{"level":"info","ts":1580674368.2601888,"caller":"logrus/json_parallele_bench.go:93","msg":"5936303275916165373","Zap":"{\"foo\":\"3421102674176290467\",\"bar\":\"8376938394956482517\"}\n"}
{"level":"info","ts":1580674368.260382,"caller":"logrus/json_parallele_bench.go:96","msg":"5936303275916165373","ZapSugar":"{\"foo\":\"3421102674176290467\",\"bar\":\"8376938394956482517\"}\n"}
{"Logrus":{"foo":"3421102674176290467","bar":"8376938394956482517"},"level":"info","msg":"5936303275916165373","time":"2020-02-02T21:12:48+01:00"}
{"level":"info","Zerolog":{"foo":"3421102674176290467","bar":"8376938394956482517"},"time":"2020-02-02T21:12:48+01:00","message":"5936303275916165373"}
--------------------
results:
Zerolog: 63389 ns per request
Logrus: 46718 ns per request
ZapSugar: 46464 ns per request
Logger.Reflect: 58088 ns per request
Logger.Stringer: 72235 ns per request
Zap: 100820 ns per request

>Total duration: 1.79128ms
_____________________
---------------------
TEST: json_parallele_bench 100
sample:
{"Message":"2598807219331453965","Level":6,"Context":{"LoggerReflect":{"foo":"5769509002805628643","bar":"534148232602821962"}}}
{"Message":"2598807219331453965","Level":6,"Context":{"LoggerStringer":{"foo":"5769509002805628643","bar":"534148232602821962"}}}
{"level":"info","ts":1580674371.842724,"caller":"logrus/json_parallele_bench.go:93","msg":"2598807219331453965","Zap":"{\"foo\":\"5769509002805628643\",\"bar\":\"534148232602821962\"}\n"}
{"level":"info","ts":1580674371.842783,"caller":"logrus/json_parallele_bench.go:96","msg":"2598807219331453965","ZapSugar":"{\"foo\":\"5769509002805628643\",\"bar\":\"534148232602821962\"}\n"}
{"level":"info","Zerolog":{"foo":"5769509002805628643","bar":"534148232602821962"},"time":"2020-02-02T21:12:51+01:00","message":"2598807219331453965"}
{"Logrus":{"foo":"5769509002805628643","bar":"534148232602821962"},"level":"info","msg":"2598807219331453965","time":"2020-02-02T21:12:51+01:00"}
--------------------
results:
Logger.Reflect: 27942 ns per request
Logger.Stringer: 28765 ns per request
Zerolog: 39570 ns per request
Logrus: 38606 ns per request
ZapSugar: 46957 ns per request
Zap: 52527 ns per request

>Total duration: 5.305772ms
_____________________
---------------------
TEST: json_parallele_bench 1000
sample:
{"Message":"4514442873692687611","Level":6,"Context":{"LoggerStringer":{"foo":"1438489581110799452","bar":"3455407192681151399"}}}
{"Message":"4514442873692687611","Level":6,"Context":{"LoggerReflect":{"foo":"1438489581110799452","bar":"3455407192681151399"}}}
{"level":"info","ts":1580674375.26467,"caller":"logrus/json_parallele_bench.go:93","msg":"4514442873692687611","Zap":"{\"foo\":\"1438489581110799452\",\"bar\":\"3455407192681151399\"}\n"}
{"level":"info","ts":1580674375.264679,"caller":"logrus/json_parallele_bench.go:96","msg":"4514442873692687611","ZapSugar":"{\"foo\":\"1438489581110799452\",\"bar\":\"3455407192681151399\"}\n"}
{"Logrus":{"foo":"1438489581110799452","bar":"3455407192681151399"},"level":"info","msg":"4514442873692687611","time":"2020-02-02T21:12:55+01:00"}
{"level":"info","Zerolog":{"foo":"1438489581110799452","bar":"3455407192681151399"},"time":"2020-02-02T21:12:55+01:00","message":"4514442873692687611"}
--------------------
results:
Logger.Stringer: 33816 ns per request
Zerolog: 34447 ns per request
Logger.Reflect: 36348 ns per request
Zap: 45243 ns per request
Logrus: 49358 ns per request
ZapSugar: 50077 ns per request

>Total duration: 50.166227ms
_____________________
---------------------
TEST: json_parallele_bench 10000
sample:
{"Message":"1183410828625268876","Level":6,"Context":{"LoggerStringer":{"foo":"3794546198651393252","bar":"250849475571353915"}}}
{"Message":"1183410828625268876","Level":6,"Context":{"LoggerReflect":{"foo":"3794546198651393252","bar":"250849475571353915"}}}
{"level":"info","ts":1580674378.862939,"caller":"logrus/json_parallele_bench.go:93","msg":"1183410828625268876","Zap":"{\"foo\":\"3794546198651393252\",\"bar\":\"250849475571353915\"}\n"}
{"level":"info","ts":1580674378.862967,"caller":"logrus/json_parallele_bench.go:96","msg":"1183410828625268876","ZapSugar":"{\"foo\":\"3794546198651393252\",\"bar\":\"250849475571353915\"}\n"}
{"level":"info","Zerolog":{"foo":"3794546198651393252","bar":"250849475571353915"},"time":"2020-02-02T21:12:58+01:00","message":"1183410828625268876"}
{"Logrus":{"foo":"3794546198651393252","bar":"250849475571353915"},"level":"info","msg":"1183410828625268876","time":"2020-02-02T21:12:58+01:00"}
--------------------
results:
Logger.Reflect: 34206 ns per request
Zerolog: 34240 ns per request
Logger.Stringer: 38075 ns per request
ZapSugar: 42952 ns per request
Zap: 43249 ns per request
Logrus: 43993 ns per request

>Total duration: 440.024081ms
_____________________
---------------------
TEST: json_parallele_bench 100000
sample:
{"level":"info","ts":1580674382.732348,"caller":"logrus/json_parallele_bench.go:93","msg":"4166398562236008690","Zap":"{\"foo\":\"9065739298427179612\",\"bar\":\"5463493963801243560\"}\n"}
{"Message":"4166398562236008690","Level":6,"Context":{"LoggerReflect":{"foo":"9065739298427179612","bar":"5463493963801243560"}}}
{"Message":"4166398562236008690","Level":6,"Context":{"LoggerStringer":{"foo":"9065739298427179612","bar":"5463493963801243560"}}}
{"level":"info","ts":1580674382.7325282,"caller":"logrus/json_parallele_bench.go:96","msg":"4166398562236008690","ZapSugar":"{\"foo\":\"9065739298427179612\",\"bar\":\"5463493963801243560\"}\n"}
{"level":"info","Zerolog":{"foo":"9065739298427179612","bar":"5463493963801243560"},"time":"2020-02-02T21:13:02+01:00","message":"4166398562236008690"}
{"Logrus":{"foo":"9065739298427179612","bar":"5463493963801243560"},"level":"info","msg":"4166398562236008690","time":"2020-02-02T21:13:02+01:00"}
--------------------
results:
Zerolog: 32702 ns per request
Logger.Stringer: 34092 ns per request
Logger.Reflect: 35872 ns per request
Zap: 43094 ns per request
ZapSugar: 43377 ns per request
Logrus: 44017 ns per request

>Total duration: 4.401782954s
_____________________
---------------------
TEST: json_parallele_bench 1000000
sample:
{"Message":"7844581879076510278","Level":6,"Context":{"LoggerReflect":{"foo":"1319504056527448916","bar":"4019846894052576717"}}}
{"level":"info","ts":1580674390.933672,"caller":"logrus/json_parallele_bench.go:96","msg":"7844581879076510278","ZapSugar":"{\"foo\":\"1319504056527448916\",\"bar\":\"4019846894052576717\"}\n"}
{"Message":"7844581879076510278","Level":6,"Context":{"LoggerStringer":{"foo":"1319504056527448916","bar":"4019846894052576717"}}}
{"level":"info","ts":1580674390.933375,"caller":"logrus/json_parallele_bench.go:93","msg":"7844581879076510278","Zap":"{\"foo\":\"1319504056527448916\",\"bar\":\"4019846894052576717\"}\n"}
{"level":"info","Zerolog":{"foo":"1319504056527448916","bar":"4019846894052576717"},"time":"2020-02-02T21:13:10+01:00","message":"7844581879076510278"}
{"Logrus":{"foo":"1319504056527448916","bar":"4019846894052576717"},"level":"info","msg":"7844581879076510278","time":"2020-02-02T21:13:10+01:00"}
--------------------
results:
Zerolog: 36466 ns per request
Logger.Reflect: 37812 ns per request
Logger.Stringer: 37891 ns per request
ZapSugar: 44792 ns per request
Zap: 45100 ns per request
Logrus: 45622 ns per request

>Total duration: 45.622511986s
_____________________
---------------------
```
