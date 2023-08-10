# window 文件及文件夹深度快速删除程序

> 利用了 robocopy 程序 删除文件名太长或者路径太深的文件

- exe 添加图标

```shell
go get github.com/akavel/rsrc

rsrc -manifest main.manifest -ico icon.ico -o main.syso
```

- 编译

```shell
go build -ldflags "-X 'main.version=$(git rev-parse --short HEAD)' -H windowsgui" .
```
