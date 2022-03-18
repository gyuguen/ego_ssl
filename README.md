## Build

```shell
ego-go build app.go
ego sign app
ego run app
```

Result
```shell
$ ego run app
EGo v0.4.2 (dafde90ae7a95bb0591f21225c8806b345a53b0d)
[erthost] loading enclave ...
[erthost] entering enclave ...
[ego] starting application ...
failed get https url: Get "https://json-schema.org/draft/2020-12/schema": x509: certificate signed by un
known authority
panic: Get "https://json-schema.org/draft/2020-12/schema": x509: certificate signed by unknown authority

goroutine 1 [running]:
main.main()
        /home/gyuguen.jang/workspace/gyuguen/ego_ssl/app.go:38 +0x57a
```