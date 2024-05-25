# Go-Head

## Description

This is a Go implementation of the Unix command line tool head.

## Installation

To install `go-head`, you need to have Go installed on your machine. Once you have Go installed, you can clone this repository and build the project:

```bash
git clone https://github.com/hash167/go-head.git
cd go-head
go build
```
This will create an executable file named `go-head` in the current directory.

## Usage

To use head-go, you can run the head-go command followed by the names of the files you want to read:

```
./head-go file1.txt file2.txt
```

By default, head-go prints the first 10 lines of each file. You can change the number of lines with the -n flag:

```
./head-go -n 20 file1.txt file2.txt
```

You can also read the first n bytes of each file with the -b flag:

```
./head-go -n 20 -b file1.txt file2.txt
```
