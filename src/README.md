# Introduction


# How to
Compile soil.go like this
```
go get github.com/stianeikeland/go-rpio
env GOOS=linux GOARCH=arm GOARM=6 go build soil.go
```
Copy the generated file to your raspberry pi device and execute it with this command

```
./soil
```

Happy blinking 