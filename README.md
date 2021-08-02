# 3n+1 with Golang

After watching Veritasium's video (https://odysee.com/@veritasium:f/the-simplest-math-problem-no-one-can:7) I decided to make a simple version of it on Go.

Requirements: GO. grab it from here: https://go.dev

To run the code directly:

`go run 3x+1.go (StartValue) (EndValue) (m|h)`

You can just run it with no arguments and that will return the length of 1 to 5.

Optional Arguments:

StartValue: Give a positive integer. It has to be an integer! Example:

`go run 3x+1.go 300`

EndValue : Give another positive integer that is bigger than the start Value. If it is smaller than the StartValue it will be StartValue + 1.Example:

`go run 3x+1.go 250 70000`

Optional Flags:

You can put either no, n or h to make the function less verbose.Ex:

`go run 3x+1.go 70000 h`

Or you can use m for matrix mode where every operation will be printed on your terminal and it will be a mess!

`go run 3x+1.go 5000 300000 m`


You can build an executable using command:

`go build 3x+1.go`

then just use that binary: 

`./3x+1`
