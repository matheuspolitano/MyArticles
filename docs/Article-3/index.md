## Elevate Your Go Logging: Why Zerolog is Your Best Choice

 I bet that YOUR first code was print "Hello World" in a log. Log is essencial to every application doesn't metter if is a huge application or your college project. Even if your app have a lot service to hendle with metrics, log will be there. So why not choice the Zerolog, that's a high-performance logging library for Go, and it is considered most fast and simple to implement.


 Thit blog post we gonna dive into the log world. Starting by the most simple use up to how use in a real world application.

 ## Topics

- Why Zerolog is so fast?
- My first zerolog




# Why Zerolog is so fast?

Zerolog is designed to avoid memory allocations as much as possible. This design priciple is crucial for achieving high performance in logging system. By minimizing allocation adn reduce pressure on the garbage collectior. leading to less GS pause times and improced overal application performance.

Implementation take advantage of modern CPU architetctures, using eddicient algorithms and data structure that reduce computational overhead and improve cache utilization


## Benchmark results

## Benchmarks

See [logbench](http://bench.zerolog.io/) for more comprehensive and up-to-date benchmarks.

All operations are allocation free (those numbers *include* JSON encoding):

```text
BenchmarkLogEmpty-8        100000000    19.1 ns/op     0 B/op       0 allocs/op
BenchmarkDisabled-8        500000000    4.07 ns/op     0 B/op       0 allocs/op
BenchmarkInfo-8            30000000     42.5 ns/op     0 B/op       0 allocs/op
BenchmarkContextFields-8   30000000     44.9 ns/op     0 B/op       0 allocs/op
BenchmarkLogFields-8       10000000     184 ns/op      0 B/op       0 allocs/op
```

There are a few Go logging benchmarks and comparisons that include zerolog.

* [imkira/go-loggers-bench](https://github.com/imkira/go-loggers-bench)
* [uber-common/zap](https://github.com/uber-go/zap#performance)

Using Uber's zap comparison benchmark:

Log a message and 10 fields:

| Library | Time | Bytes Allocated | Objects Allocated |
| :--- | :---: | :---: | :---: |
| zerolog | 767 ns/op | 552 B/op | 6 allocs/op |
| :zap: zap | 848 ns/op | 704 B/op | 2 allocs/op |
| :zap: zap (sugared) | 1363 ns/op | 1610 B/op | 20 allocs/op |
| go-kit | 3614 ns/op | 2895 B/op | 66 allocs/op |
| lion | 5392 ns/op | 5807 B/op | 63 allocs/op |
| logrus | 5661 ns/op | 6092 B/op | 78 allocs/op |
| apex/log | 15332 ns/op | 3832 B/op | 65 allocs/op |
| log15 | 20657 ns/op | 5632 B/op | 93 allocs/op |

Log a message with a logger that already has 10 fields of context:

| Library | Time | Bytes Allocated | Objects Allocated |
| :--- | :---: | :---: | :---: |
| zerolog | 52 ns/op | 0 B/op | 0 allocs/op |
| :zap: zap | 283 ns/op | 0 B/op | 0 allocs/op |
| :zap: zap (sugared) | 337 ns/op | 80 B/op | 2 allocs/op |
| lion | 2702 ns/op | 4074 B/op | 38 allocs/op |
| go-kit | 3378 ns/op | 3046 B/op | 52 allocs/op |
| logrus | 4309 ns/op | 4564 B/op | 63 allocs/op |
| apex/log | 13456 ns/op | 2898 B/op | 51 allocs/op |
| log15 | 14179 ns/op | 2642 B/op | 44 allocs/op |

Log a static string, without any context or `printf`-style templating:

| Library | Time | Bytes Allocated | Objects Allocated |
| :--- | :---: | :---: | :---: |
| zerolog | 50 ns/op | 0 B/op | 0 allocs/op |
| :zap: zap | 236 ns/op | 0 B/op | 0 allocs/op |
| standard library | 453 ns/op | 80 B/op | 2 allocs/op |
| :zap: zap (sugared) | 337 ns/op | 80 B/op | 2 allocs/op |
| go-kit | 508 ns/op | 656 B/op | 13 allocs/op |
| lion | 771 ns/op | 1224 B/op | 10 allocs/op |
| logrus | 1244 ns/op | 1505 B/op | 27 allocs/op |
| apex/log | 2751 ns/op | 584 B/op | 11 allocs/op |
| log15 | 5181 ns/op | 1592 B/op | 26 allocs/op |














