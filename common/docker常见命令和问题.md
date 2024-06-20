### 替换docker默认挂在路径


### 镜像加速

https://github.com/cmliu/CF-Workers-docker.io

### pull or build 的时候单独使用代理

有的时候，你使用的 Docker 镜像在 build 和 run 时也需要代理。

参考 官方文档 中的说明，你可以在`~/.docker/config.json` 中设置代理。
```
{
 "proxies": {
   "default": {
     "httpProxy": "http://192.168.50.100:7893",
     "httpsProxy": "http://192.168.50.100:7893",
     "noProxy": "127.0.0.0/8"
   }
 }
}
```
不想使用配置文件则直接设置环境变量

```
docker build --build-arg HTTP_PROXY="http://192.168.50.100:7893" .
docker run --env HTTP_PROXY="http://192.168.50.100:7893" redis
```



### 镜像导出
docker save -o [自定义名称].tar [镜像]

> 注意名称最好是镜像本身的名称 否则load后名称为None
eg: docker save -o coturn.tar coturn/coturn:4.6

### 镜像导入 
```
docker load < xxx.tar
docker tag [镜像ID] [镜像名称]:[标签或版本]
eg:
docker load < coturn.tar
docker tag [前面导入后生成的镜像ID] coturn/coturn:4.6
```
### docker hub操作
> https://hub.docker.com/
> sucwangsr/dockerhubsuke
> push之前创建同名镜像仓库
> 然后修改本地镜像tag为hub中创建push指定的
```s
docker login -u username -p pwd
docker push image:version
docker logout
```

### 清理镜像vol
- 删除所有无用的volume
docker volume rm $(docker volume ls -q)


### 清理日志

```
find /var/lib/docker/containers/ -name *-json.log |xargs du -sh
truncate -s 0 /var/lib/docker/containers/3c1452f817fad2296d1c105112faed89d01feaa4ee094e8622c959e072012f7a/3c1452f817fad2296d1c105112faed89d01feaa4ee094e8622c959e072012f7a-json.log
```



###  镜像临时运行并进入内部
> docker run --rm -it xx:1.0 bash


### 查看容器IP

```
docker inspect --format='{{.Name}} - {{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -aq)
```


###  docker 从容器拷贝文件


### 所有未使用镜像 网络清理

```
docker system prune
```

### ubuntu 容器 执行 更新源GPG ERROR ....报错
> 
```
docker run -it --privileged imageId
```


