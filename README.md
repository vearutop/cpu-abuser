# cpu-abuser

CPU abuser burns cycles to keep you warm. 🔥

## Why?

Some cloud providers (*cough* Oracle *cough*) don't like when machines are idle, they may even think that if running
programs are not CPU-intensive enough machine should be killed with no mercy.

This ridiculous program watches CPU usage and starts spamming CPU with meaningless work once the load in lower than a
threshold.

## Install

```
go install github.com/vearutop/cpu-abuser@latest
$(go env GOPATH)/bin/cpu-abuser --help
```

Or download binary from [releases](https://github.com/vearutop/cpu-abuser/releases).

### Linux AMD64

```
wget https://github.com/vearutop/cpu-abuser/releases/latest/download/linux_amd64.tar.gz && tar xf linux_amd64.tar.gz && rm linux_amd64.tar.gz
./cpu-abuser -version
```

## Usage

```
cpu-abuser -target 20
```

This would start a foreground `cpu-abuser` process that tries to keep CPU usage above 20%.

In practice, you'd want to have it background, for example with `screen`.

```
screen -dmS cpu-abuser ./cpu-abuser -target 18
```

(You can attach later to that screen with `screen -r cpu-abuser` and detach again by )# cpu-abuser

```
2023/06/15 00:40:23 😡 CPU Usage: 0.00%, punch(2388, 32.297µs)
2023/06/15 00:40:24 😡 CPU Usage: 15.81%, punch(2149, 32.297µs)
2023/06/15 00:40:25 😡 CPU Usage: 13.67%, punch(1934, 32.297µs)
2023/06/15 00:40:26 😡 CPU Usage: 12.56%, punch(2127, 32.297µs)
2023/06/15 00:40:30 😴 CPU Usage: 58.88%
2023/06/15 00:40:30 😴 CPU Usage: 61.54%
2023/06/15 00:40:35 😴 CPU Usage: 84.58%
2023/06/15 00:40:40 😴 CPU Usage: 100.00%
2023/06/15 00:40:45 😴 CPU Usage: 91.36%
2023/06/15 00:40:50 😴 CPU Usage: 96.32%
2023/06/15 00:40:55 😴 CPU Usage: 97.89%
2023/06/15 00:41:00 😴 CPU Usage: 79.19%
2023/06/15 00:41:05 😴 CPU Usage: 93.24%
2023/06/15 00:41:10 😴 CPU Usage: 94.14%
2023/06/15 00:41:15 😴 CPU Usage: 100.00%
2023/06/15 00:41:20 😴 CPU Usage: 78.34%
2023/06/15 00:41:25 😴 CPU Usage: 89.90%
2023/06/15 00:41:30 😴 CPU Usage: 94.10%
2023/06/15 00:41:35 😴 CPU Usage: 98.21%
2023/06/15 00:41:40 😴 CPU Usage: 99.01%
2023/06/15 00:41:45 😴 CPU Usage: 97.80%
2023/06/15 00:41:50 😴 CPU Usage: 25.82%
2023/06/15 00:41:55 😡 CPU Usage: 0.30%, punch(1915, 32.297µs)
2023/06/15 00:41:56 😡 CPU Usage: 12.76%, punch(1723, 32.297µs)
2023/06/15 00:41:57 😡 CPU Usage: 10.28%, punch(1551, 32.297µs)
2023/06/15 00:41:58 😴 CPU Usage: 19.70%
2023/06/15 00:41:58 😡 CPU Usage: 9.09%, punch(1706, 32.297µs)
2023/06/15 00:41:59 😴 CPU Usage: 26.43%
2023/06/15 00:41:59 😡 CPU Usage: 0.00%, punch(1877, 35.527µs)
2023/06/15 00:42:00 😡 CPU Usage: 12.09%, punch(1689, 35.527µs)
2023/06/15 00:42:00 😡 CPU Usage: 17.31%, punch(1858, 35.527µs)
2023/06/15 00:42:01 😡 CPU Usage: 16.85%, punch(2044, 35.527µs)
2023/06/15 00:42:02 😡 CPU Usage: 13.25%, punch(1839, 35.527µs)
2023/06/15 00:42:03 😡 CPU Usage: 9.69%, punch(2023, 35.527µs)
2023/06/15 00:42:04 😡 CPU Usage: 12.98%, punch(1821, 31.974µs)
2023/06/15 00:42:05 😡 CPU Usage: 17.09%, punch(2003, 28.777µs)
2023/06/15 00:42:06 😡 CPU Usage: 16.49%, punch(2203, 25.899µs)
2023/06/15 00:42:07 😡 CPU Usage: 16.67%, punch(2424, 23.309µs)
2023/06/15 00:42:08 😡 CPU Usage: 11.54%, punch(2181, 20.978µs)
2023/06/15 00:42:09 😴 CPU Usage: 25.57%
2023/06/15 00:42:09 😡 CPU Usage: 0.00%, punch(2399, 18.88µs)
```
