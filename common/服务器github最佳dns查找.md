
## 安装nslookup基础命令


```
apt install dnsutils # ubuntu
nslookup github.global.ssl.fastly.net
nslookup github.com
```
主要查看`Non-authoritative answer`下的内容


## 变更host文件

vi /etc/hosts

```
108.160.163.102 github.global.ssl.fastly.net
20.205.243.166 github.com

```

## 重启网络

sudo /etc/init.d/networking restart  # centos

/etc/init.d/nscd restart # ubuntu