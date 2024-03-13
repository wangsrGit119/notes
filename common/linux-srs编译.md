
## armv7

### 构建
```
git clone https://gitee.com/ossrs/srs.git

cd  srs
## 一端启动
docker run --name temp-s --rm --privileged -it registry.cn-hangzhou.aliyuncs.com/ossrs/srs:ubuntu20 bash
## 一端拷贝内容到容器
docker cp trunk/. temp-s:/tmp/srs
## 容器内编译
./configure && make





-------------构建完成后 保存容器为镜像----------------
docker commit 上面容器ID  media-server:1.0
```
### 测试

```
docker run --rm --privileged --network=host  media-server:1.0 ./objs/srs -c conf/srt.conf
```

### 容器使用宿主机摄像头
- docker run --device=/dev/video1 --rm

```
docker run --device=/dev/video0 --network=host  --rm --privileged   media-server:1.0 ./objs/srs -c conf/srt.conf
```

- push.sh脚本
```
#!/bin/bash

if [ $# -ne 1 ]; then
  echo "Usage: $0 <video_device>"
  exit 1
fi

video_device=$1

ffmpeg -i "$video_device" -s 640:480 -c:v libx264 -preset veryfast -maxrate 3000k -bufsize 6000k -pix_fmt yuv420p -c:a aac -f mpegts "srt://127.0.0.1:10080?streamid=#!::h=sss.srt.com.cn,r=live/123456,m=publish"
```
- 运行

```
sh push.sh /dev/video0
```
- 同时启动服务和推流
  ```
  docker run --device=/dev/video0 --network=host --rm --privileged media-server:2.0 sh -c './objs/srs -c conf/realtime.flv.conf & sh push.sh /dev/video0'

  ```
- 播放

```
ffplay  "srt://127.0.0.1:10080?streamid=#!::h=sss.srt.com.cn,r=live/123456,m=request"

```
