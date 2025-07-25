### 替换docker默认挂在路径

```
## 先同步数据
sudo rsync -aP /var/lib/docker/ /new/path/docker/
## 修改配置
sudo vi /lib/systemd/system/docker.service
## 找到下面行 增加 --data-root=/new/path/docker  或者直接 在docker配置增加这个配置即可
ExecStart=/usr/bin/dockerd --data-root=/new/path/docker -H fd:// --containerd=/run/containerd/containerd.sock
```

### 镜像加速

https://github.com/cmliu/CF-Workers-docker.io

### pull or build 的时候单独使用代理

有的时候，你使用的 Docker 镜像在 build 和 run 时也需要代理。

参考 官方文档 中的说明，你可以在`~/.docker/config.json` 中设置代理。
```
{
 "proxies": {
   "default": {
     "httpProxy": "http://192.168.x.100:7890",
     "httpsProxy": "http://192.168.x.100:7890",
     "noProxy": "127.0.0.0/8"
   }
 }
}
```
不想使用配置文件则直接设置环境变量

```
docker build --build-arg HTTP_PROXY="http://192.168.x.100:7890" .
docker run --env HTTP_PROXY="http://192.168.50.x:7890" redis
```



### 镜像导出
```
docker save -o [自定义名称].tar [镜像]

##注意名称最好是镜像本身的名称 否则load后名称为None
docker save -o coturn.tar coturn/coturn:4.6

## 如果很大 可以用gzip压缩
docker save ubuntu:latest | gzip > ubuntu_latest.tar.gz

```
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

## docker 创建无网络环境

```

#创建无 internal 网络 
docker network create --driver bridge --internal isolated_net


#使用示例
docker run --network=isolated_net  ...

```
