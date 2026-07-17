## Win
```
choco install mise -y
mise --version
```

## Go
```
mise use -g go@1.24
go version
```


## エラー出た場合
ヘルスチェック
```
mise doctor
```

> 1. mise shims are not on PATH
>    Add this directory to PATH: C:\Users\kakis\AppData\Local\mise\shims


スクリプト
```
$miseShims = "$env:LOCALAPPDATA\mise\shims"
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")

if ($userPath -notlike "*$miseShims*") {
    [Environment]::SetEnvironmentVariable(
        "Path",
        "$userPath;$miseShims",
        "User"
    )
}
```

