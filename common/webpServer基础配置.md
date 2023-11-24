
## 官网

https://docs.webp.sh/usage/usage/

## 安装

```
docker run --privileged=true -d -p 3333:3333 -v /home/webp-server/pics:/opt/pics --name webp-server webpsh/webp-server-go
```

## 注意事项

/home/webp-server/pics 为本地原始文件夹目录

## nginx 增加配置

> 一般拦截是拦截指定的路径下的图片 否则全站的地址图片都会被代理
```
location ~ .*\.(jpg|jpeg|png)$ {
        proxy_pass http://127.0.0.1:3333;
    }
```

docker run --privileged=true -d -p 3333:3333 --restart=always  -v /home/cebon-website/files:/opt/pics --name webp-server webpsh/webp-server-go