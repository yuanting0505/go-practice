##Exercise 8.1 

Modify clock2 to accept a port number, and write a program, clockwall, that acts
as a client of several clock servers at once, reading the times from each one
and displaying the results in a table, akin to the wall of clocks seen in some
business offices. If you have access to geographically distributed computers,
run instances remotely; otherwise run local instances on different ports with
fake time zones.

My implementation is to serve as multiple servers, not a client of multiple servers.

clock2 link: https://github.com/adonovan/gopl.io/blob/master/ch8/clock2/clock.go

### How to run
```
go build -o build/clockwall section8/clockwall.go
./build/clockwall localhost:8000 localhost:8001
```
Use command to verify:
```
nc localhost 8000
nc localhost 8001 // open a new tab
```