
## mediamtx
https://github.com/bluenviron/mediamtx

## rtsp推流

```
ffmpeg -re -stream_loop -1 -i in1.mp4 -c:v copy -c:a aac -f rtsp rtsp://localhost:8554/10001

```
