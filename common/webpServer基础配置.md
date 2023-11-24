
## 官网

https://docs.webp.sh/usage/usage/

## 安装

```
docker run --privileged=true -d -p 3333:3333 -v /home/webp-server/pics:/opt/pics --name webp-server webpsh/webp-server-go
```

## 自定义配置
> 注意路径内部的 `IMG_MAP`,表示多路径映射  映射不同的文件夹，按照此配置，则原路径的图片也可以在多个不同物理储存中或者第三方url中

```json
{
  "HOST": "0.0.0.0",
  "PORT": "3333",
  "IMG_PATH": "./pics",
  "EXHAUST_PATH": "./exhaust",
  "IMG_MAP": {
    "/2": "./pics2",
    "/3": "./pics3"
  },
  "ALLOWED_TYPES": ["jpg","png","jpeg","bmp","gif","svg","heic","nef"],
  "ENABLE_AVIF": false,
  "ENABLE_EXTRA_PARAMS": false
}

```

```shell
docker run --privileged=true -d -p 3333:3333 -v /home/webp-server/file2:/opt/pics2 -v /home/webp-server/files3:/opt/pics3 --name webp-server webpsh/webp-server-go
```

## 注意事项

/home/webp-server/pics 为本地原始文件夹目录

## nginx 增加配置

> 一般拦截是拦截指定的路径下的图片 否则全站的地址图片都会被代理
```
location ~ ^/_nuxt/img/.*\.(jpg|jpeg|png)$ {
        proxy_set_header Host $http_host; 
        rewrite ^/.*?([^/]+)$ /$1 break;
        proxy_pass http://127.0.0.1:3333; 
    }
```

docker run --privileged=true -d -p 3333:3333 --restart=always  -v /home/cebon-website/files:/opt/pics --name webp-server webpsh/webp-server-go
