# agnostic and simple benchmark cases:

```sh
make
```
```sh make
TEST: bench
Zerolog :  2199  ns per request
Logrus :  5170  ns per request
Zap :  4915  ns per request
ZapSugar :  6465  ns per request
StdLog :  2800  ns per request
--------------------
TEST: json_bench
Zerolog :  9098  ns per request
Logrus :  9274  ns per request
Zap :  6533  ns per request
ZapSugar :  6707  ns per request
StdLog :  2788  ns per request
--------------------
TEST: parallele_bench
StdLog 17588  ns per request
Zerolog 18470  ns per request
Logrus 23887  ns per request
ZapSugar 27032  ns per request
Zap 29815  ns per request
--------------------
TEST: json_parallele_bench
StdLog :  20921  ns per request
Zerolog :  32676  ns per request
Logrus :  44663  ns per request
ZapSugar :  49052  ns per request
Zap :  51944  ns per request
```