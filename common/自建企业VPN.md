
## 自建内网vpn工具 pritunl
> https://zhuanlan.zhihu.com/p/614987076
> 
> https://hub.docker.com/r/jippi/pritunl
> 
> https://www.cnblogs.com/panwenbin-logs/p/17684987.html  教程
```
mkdir -p /home/pritunl/
cd /home/pritunl/
touch pritunl.conf
默认web端口 在 443  https
```

```
docker run \
    -d --name pritunl \
    --privileged \
    --dns 127.0.0.1 \
    --network=host \
    --detach \
    -v /home/pritunl/pritunl.conf:/etc/pritunl.conf \
    -v /home/pritunl/mongodb/:/var/lib/mongodb \
    -v /home/pritunl/pritunl/:/var/lib/pritunl \
    ghcr.io/jippi/docker-pritunl

--------------查看密码---------------
docker exec -it pritunl /bin/bash
pritunl default-password
```

> 客户端下载和配置 在用户列表第一个链接图标点开可以看到
> 或者使用 openvpn https://openvpn.net/community-downloads/
>
###  注意事项
 - 内网部署上述服务
 - 右上角setting公网IP填写外部公网IP `1.2.3.4`
 - pritunl 启动的server的端口一定要和公网映射出去的端口一致 （ .ovpn文件可以看到具体连接端口）否则连接不上 假如你的内网 server端口9777，而pritunl配置公网ip为`1.2.3.4`，那么到时候远程配置中生成的连接地址就是
`remote 1.2.3.4 9777 tcp-client`
 - `Server Settings` 中注意配置是TCP还是UDP，如果内网部署了上述server，通过frp的TCP映射出去，则需要确定是TCP还是UDP
