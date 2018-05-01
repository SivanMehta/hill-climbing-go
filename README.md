Trying to implement [hill climbing](https://en.wikipedia.org/wiki/Hill_climbing) to estimate the maximum value between x = [0, 10] for the following function in `go`:

```
y = sin(x)
```

The correct answers(s) is are (π / 2, 1) and/or (5π / 2, 1)


Run `make` to run the optimization:

```sh
$ make
go build -o hill.out hill.go
./hill.out
2018/04/30 18:32:41 (probable) global max for math.Sin at x = 7.854477651926955
2018/04/30 18:32:41 (probable) global max for math.Cos at x = 12.566748149871607
2018/04/30 18:32:41 (probable) global max for math.Tan at x = 10.995501257174165
```
