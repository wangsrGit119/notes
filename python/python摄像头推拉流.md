
## python客户端从摄像头或者第三方流拉取后推流到第三方

```
from pygrabber.dshow_graph import FilterGraph
import cv2
import subprocess

def get_camera_info():
    graph = FilterGraph()
    devices = graph.get_input_devices()

    for index, device in enumerate(devices):
        print(f"Camera {index}: {device}")

get_camera_info()



def init_push_rtmp_handle(fps,width,height):
  rtmp_url = 'srt://192.168.1.172:17042?streamid=#!::h=sss.srt.com.cn,r=live/123456,m=publish'
  # command and params for ffmpeg
  command = ['ffmpeg',
             '-y',
             '-f', 'rawvideo',
             '-vcodec', 'rawvideo',
             '-pix_fmt', 'bgr24',
             '-s', "{}x{}".format(width,height),
             # '-r', str(fps),
             '-i', '-',
             '-c:v', 'libx264',
             '-pix_fmt', 'yuv420p',
             '-preset', 'ultrafast',
             '-f', 'mpegts',
             rtmp_url]

  # using subprocess and pipe to fetch frame data
  p = subprocess.Popen(command, stdin=subprocess.PIPE,shell=False)
  # p.stdin.write(frame.tostring())
  return p

def came_2_rtmp():
    cap = cv2.VideoCapture(1,cv2.CAP_DSHOW)
    fps = int(cap.get(cv2.CAP_PROP_FPS))
    width = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
    height = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))
    pipe = init_push_rtmp_handle(fps,width,height)
    while True:
        ret, frame = cap.read()
        if not ret:
            continue
        pipe.stdin.write(frame.tostring())
        # 按下 'q' 键退出循环
        if cv2.waitKey(1) & 0xFF == ord('q'):
            break



if __name__ == '__main__':

    get_camera_info()
    came_2_rtmp()
    pass
```

## py 使用vlc库拉流

```
pip install python-vlc

```
```

```
