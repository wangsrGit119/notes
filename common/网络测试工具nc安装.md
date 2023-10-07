

### centos rpm 离线包

64位
> http://vault.centos.org/6.6/os/x86_64/Packages/nc-1.84-22.el6.x86_64.rpm
32位
> http://vault.centos.org/6.6/os/i386/Packages/nc-1.84-22.el6.i686.rpm

安装
```
rpm -iUv nc-1.84-22.el6.x86_64.rpm
```

### uos 离线包


### 测试

服务端 
```
nc -ul 1080
```

测试段
```
nc -u 1080
```

测试段回车发送 服务端即可接受信息



