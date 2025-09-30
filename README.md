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