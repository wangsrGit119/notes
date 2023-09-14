

## get start

git clone https://github.com/mikel-brostrom/yolo_tracking

cd yolo_tracking

如果您想运行 YOLOv8、YOLO-NAS 或 YOLOX 示例：

git clone https://github.com/mikel-brostrom/yolo_tracking.git
pip install -v -e .
但如果您只想导入跟踪模块，您可以简单地：

pip install boxmot

## 运行追踪示例

yolo8 文档 https://docs.ultralytics.com/modes/track/#tracking

yolo 检测代码以及参数
https://github.com/ultralytics/yolov5/blob/master/detect.py



## supervision


```shell
pip install supervision[desktop]

pip install ultralytics
```

```python

import supervision as sv
from ultralytics import YOLO

# model = YOLO('yolov8n.pt')
model = YOLO('yolov8n-cls.pt')
result = model("./det-img/bus.jpg")[0]
# detections = sv.Detections.from_ultralytics(result)
classifications = sv.Classifications.from_ultralytics(result)

print(classifications)
# video_info = sv.VideoInfo.from_video_path(video_path='./static/1行人检测测试视频.mp4')
# print(video_info)

```