Generated _README.md
```sh
make all-plot
```
```sh
TEST: bench 10
sample:
{"level":"debug","time":"2020-01-30T10:01:38+01:00","message":"Zerolog: 2551171196965782110"}
{"level":"info","msg":"Logrus: 2551171196965782110","time":"2020-01-30T10:01:38+01:00"}
{"level":"info","ts":1580374898.658868,"caller":"logrus/bench.go:53","msg":"Zap: 2551171196965782110"}
{"level":"info","ts":1580374898.658977,"caller":"logrus/bench.go:54","msg":"ZapSugar: 2551171196965782110"}
2020/01/30 10:01:38 StdLog: 2551171196965782110
--------------------
results:
Zerolog: 28934 ns per request
Logrus: 19323 ns per request
Zap: 16155 ns per request
ZapSugar: 10734 ns per request
StdLog: 2423 ns per request

>Total duration: 877.904µs
_____________________
---------------------
TEST: bench 100
sample:
{"level":"debug","time":"2020-01-30T10:01:42+01:00","message":"Zerolog: 4311770505602433868"}
{"level":"info","msg":"Logrus: 4311770505602433868","time":"2020-01-30T10:01:42+01:00"}
{"level":"info","ts":1580374902.603581,"caller":"logrus/bench.go:53","msg":"Zap: 4311770505602433868"}
{"level":"info","ts":1580374902.603678,"caller":"logrus/bench.go:54","msg":"ZapSugar: 4311770505602433868"}
2020/01/30 10:01:42 StdLog: 4311770505602433868
--------------------
results:
Zerolog: 4650 ns per request
Logrus: 9258 ns per request
Zap: 6362 ns per request
ZapSugar: 6305 ns per request
StdLog: 2108 ns per request

>Total duration: 2.971398ms
_____________________
---------------------
TEST: bench 1000
sample:
{"level":"debug","time":"2020-01-30T10:01:46+01:00","message":"Zerolog: 7344154867412652978"}
{"level":"info","msg":"Logrus: 7344154867412652978","time":"2020-01-30T10:01:46+01:00"}
{"level":"info","ts":1580374906.255639,"caller":"logrus/bench.go:53","msg":"Zap: 7344154867412652978"}
{"level":"info","ts":1580374906.255796,"caller":"logrus/bench.go:54","msg":"ZapSugar: 7344154867412652978"}
2020/01/30 10:01:46 StdLog: 7344154867412652978
--------------------
results:
Zerolog: 3273 ns per request
Logrus: 9400 ns per request
Zap: 5065 ns per request
ZapSugar: 6245 ns per request
StdLog: 2230 ns per request

>Total duration: 26.430308ms
_____________________
---------------------
TEST: bench 10000
sample:
{"level":"debug","time":"2020-01-30T10:01:51+01:00","message":"Zerolog: 5109066305323995517"}
{"level":"info","msg":"Logrus: 5109066305323995517","time":"2020-01-30T10:01:51+01:00"}
{"level":"info","ts":1580374911.724465,"caller":"logrus/bench.go:53","msg":"Zap: 5109066305323995517"}
{"level":"info","ts":1580374911.724596,"caller":"logrus/bench.go:54","msg":"ZapSugar: 5109066305323995517"}
2020/01/30 10:01:51 StdLog: 5109066305323995517
--------------------
results:
Zerolog: 2702 ns per request
Logrus: 7462 ns per request
Zap: 4778 ns per request
ZapSugar: 5274 ns per request
StdLog: 2129 ns per request

>Total duration: 223.69951ms
_____________________
---------------------
TEST: bench 100000
sample:
{"level":"debug","time":"2020-01-30T10:01:55+01:00","message":"Zerolog: 6951289732495515107"}
{"level":"info","msg":"Logrus: 6951289732495515107","time":"2020-01-30T10:01:55+01:00"}
{"level":"info","ts":1580374915.729366,"caller":"logrus/bench.go:53","msg":"Zap: 6951289732495515107"}
{"level":"info","ts":1580374915.7294662,"caller":"logrus/bench.go:54","msg":"ZapSugar: 6951289732495515107"}
2020/01/30 10:01:55 StdLog: 6951289732495515107
--------------------
results:
Zerolog: 2678 ns per request
Logrus: 8460 ns per request
Zap: 6652 ns per request
ZapSugar: 6181 ns per request
StdLog: 2103 ns per request

>Total duration: 2.607918806s
_____________________
---------------------
TEST: bench 1000000
sample:
{"level":"debug","time":"2020-01-30T10:02:02+01:00","message":"Zerolog: 6610212479150430682"}
{"level":"info","msg":"Logrus: 6610212479150430682","time":"2020-01-30T10:02:02+01:00"}
{"level":"info","ts":1580374922.0724962,"caller":"logrus/bench.go:53","msg":"Zap: 6610212479150430682"}
{"level":"info","ts":1580374922.072601,"caller":"logrus/bench.go:54","msg":"ZapSugar: 6610212479150430682"}
2020/01/30 10:02:02 StdLog: 6610212479150430682
--------------------
results:
Zerolog: 2950 ns per request
Logrus: 7604 ns per request
Zap: 4823 ns per request
ZapSugar: 5296 ns per request
StdLog: 2161 ns per request

>Total duration: 22.837146136s
_____________________
---------------------
TEST: json_bench 10
sample:
{"level":"debug","time":"2020-01-30T10:02:29+01:00","message":"7216188272204197094Zerolog{\"foo\":\"7102752046737485505\",\"bar\":\"6713045383846035508\"}\n"}
{"Logrus":{"foo":"7102752046737485505","bar":"6713045383846035508"},"level":"info","msg":"7216188272204197094","time":"2020-01-30T10:02:29+01:00"}
{"level":"info","ts":1580374949.941916,"caller":"logrus/json_bench.go:74","msg":"7216188272204197094","Zap":"{\"foo\":\"7102752046737485505\",\"bar\":\"6713045383846035508\"}\n"}
{"level":"info","ts":1580374949.942035,"caller":"logrus/json_bench.go:75","msg":"7216188272204197094","ZapSugar":"{\"foo\":\"7102752046737485505\",\"bar\":\"6713045383846035508\"}\n"}
2020/01/30 10:02:29 7216188272204197094StdLog{"foo":"7102752046737485505","bar":"6713045383846035508"}
--------------------
results:
Zerolog: 34367 ns per request
Logrus: 16973 ns per request
Zap: 19552 ns per request
ZapSugar: 24674 ns per request
StdLog: 5856 ns per request

>Total duration: 1.124936ms
_____________________
---------------------
TEST: json_bench 100
sample:
{"level":"debug","time":"2020-01-30T10:02:34+01:00","message":"4939596942468343466Zerolog{\"foo\":\"6797877694076855873\",\"bar\":\"8601247144494993588\"}\n"}
{"Logrus":{"foo":"6797877694076855873","bar":"8601247144494993588"},"level":"info","msg":"4939596942468343466","time":"2020-01-30T10:02:34+01:00"}
{"level":"info","ts":1580374954.3188748,"caller":"logrus/json_bench.go:74","msg":"4939596942468343466","Zap":"{\"foo\":\"6797877694076855873\",\"bar\":\"8601247144494993588\"}\n"}
{"level":"info","ts":1580374954.318976,"caller":"logrus/json_bench.go:75","msg":"4939596942468343466","ZapSugar":"{\"foo\":\"6797877694076855873\",\"bar\":\"8601247144494993588\"}\n"}
2020/01/30 10:02:34 4939596942468343466StdLog{"foo":"6797877694076855873","bar":"8601247144494993588"}
--------------------
results:
Zerolog: 7927 ns per request
Logrus: 11357 ns per request
Zap: 10706 ns per request
ZapSugar: 9605 ns per request
StdLog: 3990 ns per request

>Total duration: 4.480428ms
_____________________
---------------------
TEST: json_bench 1000
sample:
{"level":"debug","time":"2020-01-30T10:02:37+01:00","message":"1668060262027978298Zerolog{\"foo\":\"8622487033753602090\",\"bar\":\"5412059188166712793\"}\n"}
{"Logrus":{"foo":"8622487033753602090","bar":"5412059188166712793"},"level":"info","msg":"1668060262027978298","time":"2020-01-30T10:02:37+01:00"}
{"level":"info","ts":1580374957.871028,"caller":"logrus/json_bench.go:74","msg":"1668060262027978298","Zap":"{\"foo\":\"8622487033753602090\",\"bar\":\"5412059188166712793\"}\n"}
{"level":"info","ts":1580374957.871138,"caller":"logrus/json_bench.go:75","msg":"1668060262027978298","ZapSugar":"{\"foo\":\"8622487033753602090\",\"bar\":\"5412059188166712793\"}\n"}
2020/01/30 10:02:37 1668060262027978298StdLog{"foo":"8622487033753602090","bar":"5412059188166712793"}
--------------------
results:
Zerolog: 10761 ns per request
Logrus: 30060 ns per request
Zap: 15277 ns per request
ZapSugar: 11270 ns per request
StdLog: 4485 ns per request

>Total duration: 72.153556ms
_____________________
---------------------
TEST: json_bench 10000
sample:
{"level":"debug","time":"2020-01-30T10:02:42+01:00","message":"258887970314695678Zerolog{\"foo\":\"8889058973204651738\",\"bar\":\"7146247278408535550\"}\n"}
{"Logrus":{"foo":"8889058973204651738","bar":"7146247278408535550"},"level":"info","msg":"258887970314695678","time":"2020-01-30T10:02:42+01:00"}
{"level":"info","ts":1580374962.505049,"caller":"logrus/json_bench.go:74","msg":"258887970314695678","Zap":"{\"foo\":\"8889058973204651738\",\"bar\":\"7146247278408535550\"}\n"}
{"level":"info","ts":1580374962.505193,"caller":"logrus/json_bench.go:75","msg":"258887970314695678","ZapSugar":"{\"foo\":\"8889058973204651738\",\"bar\":\"7146247278408535550\"}\n"}
2020/01/30 10:02:42 258887970314695678StdLog{"foo":"8889058973204651738","bar":"7146247278408535550"}
--------------------
results:
Zerolog: 5190 ns per request
Logrus: 11023 ns per request
Zap: 7921 ns per request
ZapSugar: 8636 ns per request
StdLog: 3957 ns per request

>Total duration: 367.54452ms
_____________________
---------------------
TEST: json_bench 100000
sample:
{"level":"debug","time":"2020-01-30T10:02:47+01:00","message":"7238173094249297225Zerolog{\"foo\":\"8521804788262623083\",\"bar\":\"6656528774611607166\"}\n"}
{"Logrus":{"foo":"8521804788262623083","bar":"6656528774611607166"},"level":"info","msg":"7238173094249297225","time":"2020-01-30T10:02:47+01:00"}
{"level":"info","ts":1580374967.4649138,"caller":"logrus/json_bench.go:74","msg":"7238173094249297225","Zap":"{\"foo\":\"8521804788262623083\",\"bar\":\"6656528774611607166\"}\n"}
{"level":"info","ts":1580374967.4650612,"caller":"logrus/json_bench.go:75","msg":"7238173094249297225","ZapSugar":"{\"foo\":\"8521804788262623083\",\"bar\":\"6656528774611607166\"}\n"}
2020/01/30 10:02:47 7238173094249297225StdLog{"foo":"8521804788262623083","bar":"6656528774611607166"}
--------------------
results:
Zerolog: 5057 ns per request
Logrus: 10238 ns per request
Zap: 8113 ns per request
ZapSugar: 8285 ns per request
StdLog: 3888 ns per request

>Total duration: 3.558578021s
_____________________
---------------------
TEST: json_bench 1000000
sample:
{"level":"debug","time":"2020-01-30T10:02:55+01:00","message":"1672769870155017885Zerolog{\"foo\":\"818598830912180347\",\"bar\":\"5347952139347106808\"}\n"}
{"Logrus":{"foo":"818598830912180347","bar":"5347952139347106808"},"level":"info","msg":"1672769870155017885","time":"2020-01-30T10:02:55+01:00"}
{"level":"info","ts":1580374975.254807,"caller":"logrus/json_bench.go:74","msg":"1672769870155017885","Zap":"{\"foo\":\"818598830912180347\",\"bar\":\"5347952139347106808\"}\n"}
{"level":"info","ts":1580374975.2549179,"caller":"logrus/json_bench.go:75","msg":"1672769870155017885","ZapSugar":"{\"foo\":\"818598830912180347\",\"bar\":\"5347952139347106808\"}\n"}
2020/01/30 10:02:55 1672769870155017885StdLog{"foo":"818598830912180347","bar":"5347952139347106808"}
--------------------
results:
Zerolog: 4952 ns per request
Logrus: 10647 ns per request
Zap: 7936 ns per request
ZapSugar: 10695 ns per request
StdLog: 4865 ns per request

>Total duration: 39.098729543s
_____________________
---------------------
TEST: parallele_bench 10
sample:
2020/01/30 10:03:39 StdLog: 7537215850295871783
{"level":"info","ts":1580375019.8331819,"caller":"logrus/parallele_bench.go:55","msg":"ZapSugar: 7537215850295871783"}
{"level":"debug","time":"2020-01-30T10:03:39+01:00","message":"Zerolog: 7537215850295871783"}
{"level":"info","ts":1580375019.833243,"caller":"logrus/parallele_bench.go:54","msg":"Zap: 7537215850295871783"}
{"level":"info","msg":"Logrus: 7537215850295871783","time":"2020-01-30T10:03:39+01:00"}
--------------------
results:
Zerolog: 31422 ns per request
Zap: 60258 ns per request
StdLog: 73718 ns per request
ZapSugar: 72353 ns per request
Logrus: 74950 ns per request

>Total duration: 866.087µs
_____________________
---------------------
TEST: parallele_bench 100
sample:
{"level":"info","ts":1580375024.206125,"caller":"logrus/parallele_bench.go:54","msg":"Zap: 5247004873706100587"}
2020/01/30 10:03:44 StdLog: 5247004873706100587
{"level":"debug","time":"2020-01-30T10:03:44+01:00","message":"Zerolog: 5247004873706100587"}
{"level":"info","ts":1580375024.206229,"caller":"logrus/parallele_bench.go:55","msg":"ZapSugar: 5247004873706100587"}
{"level":"info","msg":"Logrus: 5247004873706100587","time":"2020-01-30T10:03:44+01:00"}
--------------------
results:
StdLog: 39131 ns per request
Zerolog: 45639 ns per request
ZapSugar: 55961 ns per request
Logrus: 53423 ns per request
Zap: 57895 ns per request

>Total duration: 5.930056ms
_____________________
---------------------
TEST: parallele_bench 1000
sample:
{"level":"info","ts":1580375029.253558,"caller":"logrus/parallele_bench.go:54","msg":"Zap: 2765698756405531822"}
2020/01/30 10:03:49 StdLog: 2765698756405531822
{"level":"info","msg":"Logrus: 2765698756405531822","time":"2020-01-30T10:03:49+01:00"}
{"level":"debug","time":"2020-01-30T10:03:49+01:00","message":"Zerolog: 2765698756405531822"}
{"level":"info","ts":1580375029.254185,"caller":"logrus/parallele_bench.go:55","msg":"ZapSugar: 2765698756405531822"}
--------------------
results:
Zap: 24899 ns per request
StdLog: 25960 ns per request
Zerolog: 27039 ns per request
Logrus: 31711 ns per request
ZapSugar: 32272 ns per request

>Total duration: 32.369489ms
_____________________
---------------------
TEST: parallele_bench 10000
sample:
{"level":"info","ts":1580375033.2812932,"caller":"logrus/parallele_bench.go:54","msg":"Zap: 5761257768688307612"}
2020/01/30 10:03:53 StdLog: 5761257768688307612
{"level":"debug","time":"2020-01-30T10:03:53+01:00","message":"Zerolog: 5761257768688307612"}
{"level":"info","ts":1580375033.281302,"caller":"logrus/parallele_bench.go:55","msg":"ZapSugar: 5761257768688307612"}
{"level":"info","msg":"Logrus: 5761257768688307612","time":"2020-01-30T10:03:53+01:00"}
--------------------
results:
StdLog: 24568 ns per request
Zerolog: 27077 ns per request
ZapSugar: 32384 ns per request
Zap: 35344 ns per request
Logrus: 35676 ns per request

>Total duration: 357.086999ms
_____________________
---------------------
TEST: parallele_bench 100000
sample:
{"level":"info","ts":1580375038.768476,"caller":"logrus/parallele_bench.go:54","msg":"Zap: 3487508695484243680"}
2020/01/30 10:03:58 StdLog: 3487508695484243680
{"level":"debug","time":"2020-01-30T10:03:58+01:00","message":"Zerolog: 3487508695484243680"}
{"level":"info","ts":1580375038.768509,"caller":"logrus/parallele_bench.go:55","msg":"ZapSugar: 3487508695484243680"}
{"level":"info","msg":"Logrus: 3487508695484243680","time":"2020-01-30T10:03:58+01:00"}
--------------------
results:
StdLog: 19094 ns per request
Zerolog: 20926 ns per request
ZapSugar: 25577 ns per request
Zap: 27248 ns per request
Logrus: 27892 ns per request

>Total duration: 2.789514059s
_____________________
---------------------
TEST: parallele_bench 1000000
sample:
{"level":"info","ts":1580375046.211794,"caller":"logrus/parallele_bench.go:54","msg":"Zap: 7202437140349226036"}
{"level":"debug","time":"2020-01-30T10:04:06+01:00","message":"Zerolog: 7202437140349226036"}
2020/01/30 10:04:06 StdLog: 7202437140349226036
{"level":"info","ts":1580375046.2121012,"caller":"logrus/parallele_bench.go:55","msg":"ZapSugar: 7202437140349226036"}
{"level":"info","msg":"Logrus: 7202437140349226036","time":"2020-01-30T10:04:06+01:00"}
--------------------
results:
StdLog: 25427 ns per request
Zerolog: 25947 ns per request
Zap: 31364 ns per request
ZapSugar: 31713 ns per request
Logrus: 32788 ns per request

>Total duration: 32.789106405s
_____________________
---------------------
TEST: json_parallele_bench 10
sample:
{"level":"info","ts":1580375083.3615649,"caller":"logrus/json_parallele_bench.go:79","msg":"4092665920117515847","ZapSugar":"{\"foo\":\"6663618552046423790\",\"bar\":\"7286187599510025567\"}\n"}
{"level":"info","ts":1580375083.361568,"caller":"logrus/json_parallele_bench.go:78","msg":"4092665920117515847","Zap":"{\"foo\":\"6663618552046423790\",\"bar\":\"7286187599510025567\"}\n"}
{"level":"debug","time":"2020-01-30T10:04:43+01:00","message":"4092665920117515847Zerolog{\"foo\":\"6663618552046423790\",\"bar\":\"7286187599510025567\"}\n"}
2020/01/30 10:04:43 4092665920117515847StdLog{"foo":"6663618552046423790","bar":"7286187599510025567"}
{"Logrus":{"foo":"6663618552046423790","bar":"7286187599510025567"},"level":"info","msg":"4092665920117515847","time":"2020-01-30T10:04:43+01:00"}
--------------------
results:
Zap: 65571 ns per request
StdLog: 84503 ns per request
Zerolog: 88716 ns per request
ZapSugar: 88863 ns per request
Logrus: 101861 ns per request

>Total duration: 1.095242ms
_____________________
---------------------
TEST: json_parallele_bench 100
sample:
{"level":"info","ts":1580375087.201748,"caller":"logrus/json_parallele_bench.go:79","msg":"6009594360842759989","ZapSugar":"{\"foo\":\"2783633245467298950\",\"bar\":\"1011674774904777804\"}\n"}
{"level":"info","ts":1580375087.2017338,"caller":"logrus/json_parallele_bench.go:78","msg":"6009594360842759989","Zap":"{\"foo\":\"2783633245467298950\",\"bar\":\"1011674774904777804\"}\n"}
{"level":"debug","time":"2020-01-30T10:04:47+01:00","message":"6009594360842759989Zerolog{\"foo\":\"2783633245467298950\",\"bar\":\"1011674774904777804\"}\n"}
2020/01/30 10:04:47 6009594360842759989StdLog{"foo":"2783633245467298950","bar":"1011674774904777804"}
{"Logrus":{"foo":"2783633245467298950","bar":"1011674774904777804"},"level":"info","msg":"6009594360842759989","time":"2020-01-30T10:04:47+01:00"}
--------------------
results:
StdLog: 39601 ns per request
Zerolog: 41191 ns per request
Zap: 42617 ns per request
Logrus: 45376 ns per request
ZapSugar: 46399 ns per request

>Total duration: 4.745862ms
_____________________
---------------------
TEST: json_parallele_bench 1000
sample:
{"level":"info","ts":1580375090.8870678,"caller":"logrus/json_parallele_bench.go:78","msg":"3809845119943323333","Zap":"{\"foo\":\"3994538292824829166\",\"bar\":\"6977118017069742433\"}\n"}
{"level":"info","ts":1580375090.887087,"caller":"logrus/json_parallele_bench.go:79","msg":"3809845119943323333","ZapSugar":"{\"foo\":\"3994538292824829166\",\"bar\":\"6977118017069742433\"}\n"}
{"level":"debug","time":"2020-01-30T10:04:50+01:00","message":"3809845119943323333Zerolog{\"foo\":\"3994538292824829166\",\"bar\":\"6977118017069742433\"}\n"}
2020/01/30 10:04:50 3809845119943323333StdLog{"foo":"3994538292824829166","bar":"6977118017069742433"}
{"Logrus":{"foo":"3994538292824829166","bar":"6977118017069742433"},"level":"info","msg":"3809845119943323333","time":"2020-01-30T10:04:50+01:00"}
--------------------
results:
StdLog: 34357 ns per request
Zerolog: 36886 ns per request
Zap: 38787 ns per request
ZapSugar: 40727 ns per request
Logrus: 41415 ns per request

>Total duration: 41.49768ms
_____________________
---------------------
TEST: json_parallele_bench 10000
sample:
{"level":"info","ts":1580375094.919905,"caller":"logrus/json_parallele_bench.go:79","msg":"5651318541934194475","ZapSugar":"{\"foo\":\"187328805340793295\",\"bar\":\"3035672743497359373\"}\n"}
{"level":"info","ts":1580375094.919887,"caller":"logrus/json_parallele_bench.go:78","msg":"5651318541934194475","Zap":"{\"foo\":\"187328805340793295\",\"bar\":\"3035672743497359373\"}\n"}
{"level":"debug","time":"2020-01-30T10:04:54+01:00","message":"5651318541934194475Zerolog{\"foo\":\"187328805340793295\",\"bar\":\"3035672743497359373\"}\n"}
2020/01/30 10:04:54 5651318541934194475StdLog{"foo":"187328805340793295","bar":"3035672743497359373"}
{"Logrus":{"foo":"187328805340793295","bar":"3035672743497359373"},"level":"info","msg":"5651318541934194475","time":"2020-01-30T10:04:54+01:00"}
--------------------
results:
StdLog: 30541 ns per request
Zerolog: 33325 ns per request
Zap: 39230 ns per request
ZapSugar: 39663 ns per request
Logrus: 40407 ns per request

>Total duration: 404.174322ms
_____________________
---------------------
TEST: json_parallele_bench 100000
sample:
{"level":"info","ts":1580375099.209841,"caller":"logrus/json_parallele_bench.go:78","msg":"3380046655242664175","Zap":"{\"foo\":\"435822572708970895\",\"bar\":\"4768704971931816079\"}\n"}
{"level":"info","ts":1580375099.20985,"caller":"logrus/json_parallele_bench.go:79","msg":"3380046655242664175","ZapSugar":"{\"foo\":\"435822572708970895\",\"bar\":\"4768704971931816079\"}\n"}
{"level":"debug","time":"2020-01-30T10:04:59+01:00","message":"3380046655242664175Zerolog{\"foo\":\"435822572708970895\",\"bar\":\"4768704971931816079\"}\n"}
2020/01/30 10:04:59 3380046655242664175StdLog{"foo":"435822572708970895","bar":"4768704971931816079"}
{"Logrus":{"foo":"435822572708970895","bar":"4768704971931816079"},"level":"info","msg":"3380046655242664175","time":"2020-01-30T10:04:59+01:00"}
--------------------
results:
Zerolog: 29581 ns per request
StdLog: 31129 ns per request
Zap: 35550 ns per request
ZapSugar: 35957 ns per request
Logrus: 36127 ns per request

>Total duration: 3.612837951s
_____________________
---------------------
TEST: json_parallele_bench 1000000
sample:
{"level":"info","ts":1580375106.758386,"caller":"logrus/json_parallele_bench.go:79","msg":"3114230902425843949","ZapSugar":"{\"foo\":\"6919128875777902408\",\"bar\":\"6803128349924459680\"}\n"}
{"level":"info","ts":1580375106.75839,"caller":"logrus/json_parallele_bench.go:78","msg":"3114230902425843949","Zap":"{\"foo\":\"6919128875777902408\",\"bar\":\"6803128349924459680\"}\n"}
2020/01/30 10:05:06 3114230902425843949StdLog{"foo":"6919128875777902408","bar":"6803128349924459680"}
{"level":"debug","time":"2020-01-30T10:05:06+01:00","message":"3114230902425843949Zerolog{\"foo\":\"6919128875777902408\",\"bar\":\"6803128349924459680\"}\n"}
{"Logrus":{"foo":"6919128875777902408","bar":"6803128349924459680"},"level":"info","msg":"3114230902425843949","time":"2020-01-30T10:05:06+01:00"}
--------------------
results:
StdLog: 35590 ns per request
Zerolog: 35619 ns per request
ZapSugar: 43483 ns per request
Zap: 43487 ns per request
Logrus: 43964 ns per request

>Total duration: 43.964588385s
_____________________
---------------------
```
