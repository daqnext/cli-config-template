# config-template
### example of how to write a highly usable cli  program
```go
// configs/dev/appName.json is used for dev mod
// configs/pro/appName.json is used for production mod
// you can overwrite attributes of configs/xxx/xxxx.json with the cli input params
// according to your cli input , a unique app is selected and started
// all the global component is initialized in global.go file 

// !important , for any long-term-runing serverside app like 'service','background job'...
// a http hearbeat (controller) check is strongly suggested for error detection like [aws-ec2-loadbalancer]

// start default app with command :
go build && ./cli-config-template --dev=true
go build && ./cli-config-template  // default for pro mod , need proconfig.json

// start the log app with command :
go build && ./cli-config-template logs --num=10   
go build && ./cli-config-template logs --onlyerr  --num=10   //only print error logs
```

### publish
```
bring your 'assets' and 'config' folder together with your exe
as these folders are usded in runtime
```