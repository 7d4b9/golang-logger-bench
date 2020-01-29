# agnostic and simple benchmark cases:
```sh
make all-plot
```
```sh
TEST: bench 10
--------------------
Zerolog :  21689  ns per request
Logrus :  9221  ns per request
Zap :  12595  ns per request
ZapSugar :  7821  ns per request
StdLog :  1798  ns per request

>Total duration: 626.419µs
_____________________
---------------------
TEST: bench 100
--------------------
Zerolog :  3928  ns per request
Logrus :  4526  ns per request
Zap :  5225  ns per request
ZapSugar :  5007  ns per request
StdLog :  1626  ns per request

>Total duration: 2.121109ms
_____________________
---------------------
TEST: bench 1000
--------------------
Zerolog :  2264  ns per request
Logrus :  4223  ns per request
Zap :  4747  ns per request
ZapSugar :  5138  ns per request
StdLog :  1668  ns per request

>Total duration: 18.158651ms
_____________________
---------------------
TEST: bench 10000
--------------------
Zerolog :  2215  ns per request
Logrus :  4317  ns per request
Zap :  5396  ns per request
ZapSugar :  6165  ns per request
StdLog :  1572  ns per request

>Total duration: 196.875283ms
_____________________
---------------------
TEST: bench 100000
--------------------
Zerolog :  2185  ns per request
Logrus :  4303  ns per request
Zap :  4112  ns per request
ZapSugar :  4469  ns per request
StdLog :  1559  ns per request

>Total duration: 1.663393687s
_____________________
---------------------
TEST: bench 1000000
--------------------
Zerolog :  2165  ns per request
Logrus :  4087  ns per request
Zap :  4075  ns per request
ZapSugar :  4504  ns per request
StdLog :  1644  ns per request

>Total duration: 16.479261346s
_____________________
---------------------
TEST: json_bench 10
--------------------
Zerolog :  41481  ns per request
Logrus :  19300  ns per request
Zap :  19484  ns per request
ZapSugar :  13802  ns per request
StdLog :  4417  ns per request

>Total duration: 1.082547ms
_____________________
---------------------
TEST: json_bench 100
--------------------
Zerolog :  7668  ns per request
Logrus :  10016  ns per request
Zap :  7861  ns per request
ZapSugar :  8365  ns per request
StdLog :  3183  ns per request

>Total duration: 3.796567ms
_____________________
---------------------
TEST: json_bench 1000
--------------------
Zerolog :  4217  ns per request
Logrus :  9673  ns per request
Zap :  7439  ns per request
ZapSugar :  8954  ns per request
StdLog :  3136  ns per request

>Total duration: 33.582966ms
_____________________
---------------------
TEST: json_bench 10000
--------------------
Zerolog :  4464  ns per request
Logrus :  9920  ns per request
Zap :  8525  ns per request
ZapSugar :  8813  ns per request
StdLog :  3496  ns per request

>Total duration: 352.438076ms
_____________________
---------------------
TEST: json_bench 100000
--------------------
Zerolog :  4775  ns per request
Logrus :  10985  ns per request
Zap :  8074  ns per request
ZapSugar :  9126  ns per request
StdLog :  3953  ns per request

>Total duration: 3.691762954s
_____________________
---------------------
TEST: json_bench 1000000
--------------------
Zerolog :  5207  ns per request
Logrus :  10849  ns per request
Zap :  8034  ns per request
ZapSugar :  9117  ns per request
StdLog :  3830  ns per request

>Total duration: 37.039654189s
_____________________
---------------------
TEST: parallele_bench 10
--------------------
StdLog 40139  ns per request
Logrus 33676  ns per request
Zerolog 49149  ns per request
Zap 60682  ns per request
ZapSugar 65615  ns per request

>Total duration: 729.022µs
_____________________
---------------------
TEST: parallele_bench 100
--------------------
StdLog 22615  ns per request
Logrus 25308  ns per request
Zerolog 30714  ns per request
Zap 36771  ns per request
ZapSugar 37244  ns per request

>Total duration: 3.797542ms
_____________________
---------------------
TEST: parallele_bench 1000
--------------------
StdLog 24670  ns per request
Zerolog 27247  ns per request
Logrus 30155  ns per request
Zap 30751  ns per request
ZapSugar 31189  ns per request

>Total duration: 31.282967ms
_____________________
---------------------
TEST: parallele_bench 10000
--------------------
StdLog 24681  ns per request
Zerolog 26648  ns per request
Zap 30620  ns per request
Logrus 30772  ns per request
ZapSugar 30847  ns per request

>Total duration: 308.575531ms
_____________________
---------------------
TEST: parallele_bench 100000
--------------------
StdLog 24420  ns per request
Zerolog 25709  ns per request
Logrus 30592  ns per request
Zap 30737  ns per request
ZapSugar 30912  ns per request

>Total duration: 3.091298458s
_____________________
---------------------
TEST: parallele_bench 1000000
--------------------
StdLog 23625  ns per request
Zerolog 24746  ns per request
Zap 29146  ns per request
Logrus 29526  ns per request
ZapSugar 29584  ns per request

>Total duration: 29.584799602s
_____________________
---------------------
TEST: json_parallele_bench 10
--------------------
Zap :  67841  ns per request
ZapSugar :  71884  ns per request
Zerolog :  77985  ns per request
StdLog :  92419  ns per request
Logrus :  104008  ns per request

>Total duration: 1.117981ms
_____________________
---------------------
TEST: json_parallele_bench 100
--------------------
StdLog :  40368  ns per request
Zerolog :  40084  ns per request
Zap :  44443  ns per request
ZapSugar :  46603  ns per request
Logrus :  47896  ns per request

>Total duration: 4.871386ms
_____________________
---------------------
TEST: json_parallele_bench 1000
--------------------
StdLog :  35501  ns per request
Zerolog :  36623  ns per request
Zap :  39144  ns per request
Logrus :  40314  ns per request
ZapSugar :  41153  ns per request

>Total duration: 41.262887ms
_____________________
---------------------
TEST: json_parallele_bench 10000
--------------------
StdLog :  30671  ns per request
Zerolog :  34732  ns per request
Zap :  39925  ns per request
ZapSugar :  40657  ns per request
Logrus :  41779  ns per request

>Total duration: 417.90932ms
_____________________
---------------------
TEST: json_parallele_bench 100000
--------------------
StdLog :  30265  ns per request
Zerolog :  32603  ns per request
Zap :  36910  ns per request
ZapSugar :  38520  ns per request
Logrus :  39209  ns per request

>Total duration: 3.921070083s
_____________________
---------------------
TEST: json_parallele_bench 1000000
--------------------
Zerolog :  32218  ns per request
StdLog :  32790  ns per request
ZapSugar :  39015  ns per request
Zap :  39117  ns per request
Logrus :  39979  ns per request

>Total duration: 39.980037107s
_____________________
---------------------
```
