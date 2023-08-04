
## 按照不同的环境拉镜像

docker pull mpromonet/webrtc-streamer:v0.8.1-amd64


## github

https://github.com/mpromonet/webrtc-streamer

## 启动 容器

- 默认启动
docker rm -f webrtc-streamer && docker run -d --name webrtc-streamer -p 18001:8000  mpromonet/webrtc-streamer:v0.8.1-amd64

- 挂在外部 config.json 参数

docker rm -f webrtc-streamer && docker run -d --name webrtc-streamer -p 18001:8000 -v /home/webrtc-stream/config.json:/app/config.json  mpromonet/webrtc-streamer:v0.8.1-amd64

docker rm -f webrtc-streamer 
docker run --rm --name webrtc-streamer -p 18001:8000 -v /home/webrtc-stream/config.json:/app/config.json  mpromonet/webrtc-streamer:v0.8.1-amd64

## windows/linux 成品文件启动

https://github.com/mpromonet/webrtc-streamer/releases/tag/v0.8.1

webrtc-streamer -S0.0.0.0:3478 -s$(curl -s ifconfig.me):3478 -T0.0.0.0:3479 -tturn:turn@$(curl -s ifconfig.me):3479 -C config.json


- 启动stun和turn实例并使用

> IP为服务器IP，如果公网服务器则为公网IP 

```
# 内网建议这个
docker run --rm -p 3478:3478/udp -p 3478:3478/tcp -p 3479:3479/udp -p 3479:3479/tcp -p 18001:8000 --name webrtc-streamer mpromonet/webrtc-streamer:v0.8.1-amd64 -S0.0.0.0:3478 -s192.168.1.172:3478 -T0.0.0.0:3479 -tturn:turn@192.168.1.172:3479 -C config.json

# 公网建议这个
docker run --rm -p 3478:3478/udp -p 3478:3478/tcp -p 3479:3479/udp -p 3479:3479/tcp -p 18001:8000 --name webrtc-streamer mpromonet/webrtc-streamer:v0.8.1-amd64 -S0.0.0.0:3478 -s$(curl -s ifconfig.me):3478 -T0.0.0.0:3479 -tturn:turn@$(curl -s ifconfig.me):3479 -C config.json
```

## 设置日志级别

level默认为4，可以从0开始
/api/log?level=0



##  前端播放(http即可播放)

```
<html>
<head>
<script src="dist/adapter.min.js" ></script>
<script src="dist/webrtcstreamer.js" ></script>
<style type="text/css">
	#video {
		width: 600px;
		height: 350px;
		
	}
</style>
<script>        
	var webRtcServer      = null;
	window.onload         = function() { 
		webRtcServer      = new WebRtcStreamer("video","http://192.168.1.xx:17003/");
		webRtcServer.connect("rtsp://admin:admin@192.168.1.23");
	}
	window.onbeforeunload = function() { webRtcServer.disconnect(); }
</script>


</head>
<body> 
	<video id="video" autoplay muted/>
</body>
</html>
```
