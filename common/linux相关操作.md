
## 磁盘挂载 小于等于2T 用 fdisk

https://www.cnblogs.com/aqicheng/p/14035090.html

## 磁盘挂载 大于2T 用 parted

https://www.cnblogs.com/kevingrace/p/7612741.html

## 开机自启挂载

### 查看分区UUID

```
sudo blkid
```

### 编辑配置文件

```
vi /etc/fstab

-----------------------------追加下面内容---/home/xxx为挂载分区的主机目录
UUID=<your_partition_uuid> /home/xxx ext4 defaults 0 0


```

### 检查挂载配置

```
sudo mount -a
```
