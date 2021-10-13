# config-template
### example of how to write a highly usable cli  program
```go
// config/xxx_devconfig.json is used for dev mod
// config/xxx_proconfig.json is used for production mod
// you can overwrite attributes of config/xxx_xxxconfig.json with the cli input params
// according to your cli input , a unique app is selected and started
// put all the global singleton component in global.go file 

// !important , for any long-term-runing app like 'service','background job'...
// a http hearbeat (controller) check is strongly suggest for error detection

// start default app with command :
go build && ./cli-config-template --dev=true
go build && ./cli-config-template  // default for pro mod , need proconfig.json

// start the log app with command :
go build && ./cli-config-template logs --num=10   
go build && ./cli-config-template logs --onlyerr  --num=10   //only print error logs
```

### publish
```
after build your exe file ,don't forget to bring your 'assets' and 'config' folder together 
as these folders are usded when running
```