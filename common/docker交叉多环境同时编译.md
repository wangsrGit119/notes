
## docker buildx 安装

```

docker buildx install

```

## 开始构建
> --push 表示推送到登录的hub仓库
```
docker buildx build --platform linux/arm64,linux/amd64  -t sucwangsr/dyn-contain:1.0 . --push
```
