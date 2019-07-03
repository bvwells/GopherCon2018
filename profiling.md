# Profiling

This document details how to profile go applications.

# Generating CPU profiles

It is possible to generate CPU profiles when running benchmarks using the following command:

```
go test -bench=. -cpuprofile=cpu.pprof
```

This will generate a pprof file called ```cpu.pprof``` which can be examined using pprof.

# Examining CPUs profiles

To examine CPU profiles run the following command:

```
go tool pprof cpu.pprof
```

# Generating memory profiles

It is possible to generate memory profiles when running benchmarks using the following command:

```
go test -bench=. -memprofile=mem.pprof
```

This will generate a pprof file called ```mem.pprof``` which can be examined using pprof.

# Examining memory profiles

To examine memory profiles run the following command:
```
go tool pprof --alloc_objects mem.pprof
```

# Running Google pprof

Google pprof can be used to examine profiles and generate flamecharts. The google pprof tool can be obtained by running the following command:

```
go get github.com/google/pprof
```

To profile an existing profile run:
```
pprof cpu.pprof
```

The profile can be examined on localhost by running:
```
pprof -http=localhost:8181 cpu.pprof
```