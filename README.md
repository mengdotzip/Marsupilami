# Marsupilami 🦘

A minimal Go REPL.

## What is this?

Marsupilami (mspl) lets you write Go code interactively. Type your code, hit `return` to execute it.

## Install

```
git clone https://github.com/mengdotzip/Marsupilami
cd marsupilami
go build
```


## Usage

Run it:
```
./marsupilami
```

### Example
```
MSPL ambrose@RedWitch:~/Code/Marsupilami$ x := 10 + 5
MSPL ambrose@RedWitch:~/Code/Marsupilami$ y := x * 2
MSPL ambrose@RedWitch:~/Code/Marsupilami$ fmt.Println(y)
MSPL ambrose@RedWitch:~/Code/Marsupilami$ return
30
```
You can always check the current code you are building
```
MSPL ambrose@RedWitch:~/Code/Marsupilami$ x := 10 + 5
MSPL ambrose@RedWitch:~/Code/Marsupilami$ y := x * 2
MSPL ambrose@RedWitch:~/Code/Marsupilami$ current
x := 10 + 5
y := x * 2

```
Or the history of already ran code
```
MSPL ambrose@RedWitch:~/Code/Marsupilami$ history
Index:0
Code:
x := 10 + 5
y := x * 2
fmt.Println(y)

Timestamp:2025-10-01 17:38:16.30717469 +0200 CEST m=+16.838390291
MSPL ambrose@RedWitch:~/Code/Marsupilami$ 
```