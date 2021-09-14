# config-template
### example of how to write a highly usable cli  program
```go
// config/devconfig.json is used for dev mod
// config/proconfig.json is used for production mod
// you can overwrite attributes of config/xxxconfig.json with the cli input params
// according to your cli input , a unique app is selected and started
// put all the global singleton component in global.go file 

// start with command :

go run ./main.go --dev=true // for dev mod
go run ./main.go            // default for pro mod

```