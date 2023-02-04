"# sample_go"

# [install]

https://go.dev/

echo %GOPATH%

# [go doc]

go install golang.org/x/tools/cmd/godoc@latest

※非推奨 go get コマンド

```
go: 'go install' requires a version when current directory is not in a module

Try 'go install golang.org/x/tools/cmd/godoc@latest' to install the latest version
```

# [GOROOT path 問題]

go mod init {project 名} // root で実行

# [import 問題]

- プライベートとパブリックを理解する
  - 大文字小文字

# [VsCode Format]

参考: https://zenn.dev/t_morooka/articles/24ce28f4844895

setting.json > 下記追加

"[go]": {
"editor.defaultFormatter": "golang.go"
},

```
windows: ctrl + ,(comma)
mac: cmd + ,(comma)
```
