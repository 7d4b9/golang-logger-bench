# agnostic and simple benchmark cases:

```sh
make all-plot
```
```sh
TEST: bench 10
--------------------
Zerolog :  23141  ns per request
Logrus :  10041  ns per request
Zap :  14458  ns per request
ZapSugar :  7407  ns per request
StdLog :  2236  ns per request
_____________________
---------------------
TEST: bench 100
--------------------
Zerolog :  4604  ns per request
Logrus :  5861  ns per request
Zap :  6273  ns per request
ZapSugar :  5737  ns per request
StdLog :  2096  ns per request
_____________________
---------------------
TEST: bench 1000
--------------------
Zerolog :  2758  ns per request
Logrus :  5150  ns per request
Zap :  5067  ns per request
ZapSugar :  5536  ns per request
StdLog :  1968  ns per request
_____________________
---------------------
TEST: bench 10000
--------------------
Zerolog :  2534  ns per request
Logrus :  5079  ns per request
Zap :  5060  ns per request
ZapSugar :  5548  ns per request
StdLog :  1984  ns per request
_____________________
---------------------
TEST: bench 100000
--------------------
Zerolog :  2532  ns per request
Logrus :  4990  ns per request
Zap :  4826  ns per request
ZapSugar :  5444  ns per request
StdLog :  1959  ns per request
_____________________
---------------------
TEST: bench 1000000
--------------------
Zerolog :  2540  ns per request
Logrus :  6302  ns per request
Zap :  4898  ns per request
ZapSugar :  5566  ns per request
StdLog :  1978  ns per request
_____________________
---------------------
TEST: json_bench 10
--------------------
Zerolog :  44884  ns per request
Logrus :  17658  ns per request
Zap :  18943  ns per request
ZapSugar :  11183  ns per request
StdLog :  4480  ns per request
_____________________
---------------------
TEST: json_bench 100
--------------------
Zerolog :  7856  ns per request
Logrus :  11837  ns per request
Zap :  9640  ns per request
ZapSugar :  9435  ns per request
StdLog :  3938  ns per request
_____________________
---------------------
TEST: json_bench 1000
--------------------
Zerolog :  5306  ns per request
Logrus :  11966  ns per request
Zap :  8517  ns per request
ZapSugar :  10073  ns per request
StdLog :  3863  ns per request
_____________________
---------------------
TEST: json_bench 10000
--------------------
Zerolog :  4967  ns per request
Logrus :  12503  ns per request
Zap :  8329  ns per request
ZapSugar :  9250  ns per request
StdLog :  3931  ns per request
_____________________
---------------------
TEST: json_bench 100000
--------------------
Zerolog :  5074  ns per request
Logrus :  10747  ns per request
Zap :  8099  ns per request
ZapSugar :  8958  ns per request
StdLog :  3958  ns per request
_____________________
---------------------
TEST: json_bench 1000000
--------------------
Zerolog :  5382  ns per request
Logrus :  11291  ns per request
Zap :  8414  ns per request
ZapSugar :  9489  ns per request
StdLog :  3902  ns per request
_____________________
---------------------
TEST: parallele_bench 10
--------------------
Zerolog 35986  ns per request
StdLog 48515  ns per request
Logrus 54040  ns per request
Zap 68678  ns per request
ZapSugar 70065  ns per request
_____________________
---------------------
TEST: parallele_bench 100
--------------------
StdLog 15767  ns per request
Zerolog 19278  ns per request
Logrus 28050  ns per request
Zap 31824  ns per request
ZapSugar 32251  ns per request
_____________________
---------------------
TEST: parallele_bench 1000
--------------------
StdLog 18331  ns per request
Zerolog 20763  ns per request
Logrus 26464  ns per request
Zap 27404  ns per request
ZapSugar 30618  ns per request
_____________________
---------------------
TEST: parallele_bench 10000
--------------------
Zerolog 16967  ns per request
StdLog 17200  ns per request
Logrus 24978  ns per request
Zap 28118  ns per request
ZapSugar 28226  ns per request
_____________________
---------------------
TEST: parallele_bench 100000
--------------------
StdLog 17108  ns per request
Zerolog 18559  ns per request
Logrus 23993  ns per request
Zap 27350  ns per request
ZapSugar 27723  ns per request
_____________________
---------------------
TEST: parallele_bench 1000000
--------------------
StdLog 16085  ns per request
Zerolog 17496  ns per request
Logrus 22826  ns per request
Zap 26134  ns per request
ZapSugar 26425  ns per request
_____________________
---------------------
TEST: json_parallele_bench 10
--------------------
Zap :  63708  ns per request
ZapSugar :  65388  ns per request
Zerolog :  85532  ns per request
StdLog :  97379  ns per request
Logrus :  102413  ns per request
_____________________
---------------------
TEST: json_parallele_bench 100
--------------------
StdLog :  30081  ns per request
Zerolog :  33836  ns per request
Logrus :  42774  ns per request
ZapSugar :  44255  ns per request
Zap :  44872  ns per request
_____________________
---------------------
TEST: json_parallele_bench 1000
--------------------
StdLog :  24600  ns per request
Zerolog :  28026  ns per request
Logrus :  38086  ns per request
Zap :  39980  ns per request
ZapSugar :  40323  ns per request
_____________________
---------------------
TEST: json_parallele_bench 10000
--------------------
StdLog :  26170  ns per request
Zerolog :  28222  ns per request
Logrus :  37299  ns per request
Zap :  40292  ns per request
ZapSugar :  40987  ns per request
_____________________
---------------------
TEST: json_parallele_bench 100000
--------------------
StdLog :  21958  ns per request
Zerolog :  22793  ns per request
Logrus :  34000  ns per request
Zap :  37154  ns per request
ZapSugar :  37332  ns per request
_____________________
---------------------
TEST: json_parallele_bench 1000000
--------------------
StdLog :  24425  ns per request
Zerolog :  24768  ns per request
Logrus :  35068  ns per request
Zap :  37870  ns per request
ZapSugar :  38086  ns per request
_____________________
---------------------
