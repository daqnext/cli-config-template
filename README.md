# config-template
### example of how to write a highly usable cli  program
```
// the parameters will be overwrite to gconfig.json file
go run ./main.go  firstcmd --optionbool=false --optionnum=3123
go run ./main.go   
go run ./main.go  firstcmd --optionbool=true --optionnum=333
go run ./main.go   

```