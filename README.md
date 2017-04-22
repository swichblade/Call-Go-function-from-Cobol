# Call-Go-function-from-Cobol
Show how Cobol arguments are passed to a Go function, and then back to Cobol

Using Gnucobol

From Ubuntu:

$ go build -buildmode=c-shared -o functions.so functions.go
$ export COB_PRE_LOAD=functions
$ cobc -x cobmain.cob
$ ./cobmain
