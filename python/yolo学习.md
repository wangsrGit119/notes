

## Get Start

```
git clone https://github.com/mikel-brostrom/yolo_tracking

cd yolo_tracking
```

如果您想运行 YOLOv8、YOLO-NAS 或 YOLOX 示例：
```
git clone https://github.com/mikel-brostrom/yolo_tracking.git
pip install -v -e .
```
但如果您只想导入跟踪模块，您可以简单地：
```
pip install boxmot
```

## 运行追踪示例

yolo8 文档 https://docs.ultralytics.com/modes/track/#tracking

yolo 检测代码以及参数
https://github.com/ultralytics/yolov5/blob/master/detect.py


## 配置参数大全

https://docs.ultralytics.com/usage/cfg/#predict

## 如果单单用yolo则按照下面

1. 图片
```
from ultralytics import YOLO
import cv2
model = YOLO('./models/yolov8n.pt')

results = model('./det-img/test-img.png')  # 对图像进行预测
print(results[0])
# 处理结果列表
for result in results:
    boxes = result.boxes  # 边界框输出的 Boxes 对象
    masks = result.masks  # 分割掩码输出的 Masks 对象
    keypoints = result.keypoints  # 姿态输出的 Keypoints 对象
    probs = result.probs  # 分类输出的 Probs 对象

annotated_frame = results[0].plot()
# 展示带注释的帧
cv2.imshow("YOLOv8 Tracking", annotated_frame)

while True:
	if cv2.waitKey(1) & 0xFF == ord("q"):
		break
	else:
		pass

```
2.视频

```
from ultralytics import YOLO

# 配置追踪参数并运行追踪器
model = YOLO('yolov8n.pt')
def video_process():
	video_path = "./det-videos/person.mp4"
	cap = cv2.VideoCapture(video_path)
	# 存储追踪历史
	track_history = defaultdict(lambda: [])

	# 循环遍历视频帧
	while cap.isOpened():
	    # 从视频读取一帧
	    success, frame = cap.read()

	    if success:
	        # 在帧上运行YOLOv8追踪，持续追踪帧间的物体
	        results = model.track(frame, persist=True)

	        # 获取框和追踪ID
	        boxes = results[0].boxes.xywh.cpu()
	        track_ids = results[0].boxes.id.int().cpu().tolist()

	        # 在帧上展示结果
	        annotated_frame = results[0].plot()

	        # 绘制追踪路径
	        for box, track_id in zip(boxes, track_ids):
	            x, y, w, h = box
	            track = track_history[track_id]
	            track.append((float(x), float(y)))  # x, y中心点
	            if len(track) > 30:  # 在90帧中保留90个追踪点
	                track.pop(0)

	            # 绘制追踪线
	            points = np.hstack(track).astype(np.int32).reshape((-1, 1, 2))
	            cv2.polylines(annotated_frame, [points], isClosed=False, color=(230, 230, 230), thickness=10)

	        # 展示带注释的帧
	        cv2.imshow("YOLOv8 Tracking", annotated_frame)

	        # 如果按下'q'则退出循环
	        if cv2.waitKey(1) & 0xFF == ord("q"):
	            break
	    else:
	        # 如果视频结束则退出循环
	        break

	# 释放视频捕获对象并关闭显示窗口
	cap.release()
	cv2.destroyAllWindows()

```

## 配合二次图像处理包 Supervision

### 官网：https://supervision.roboflow.com/

```shell
pip install supervision[desktop]
pip install tqdm

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

##  使用gpu

> cuda 版本 https://pytorch.org/get-started/locally/
> pip install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cu117

https://github.com/ultralytics/ultralytics/issues/3084

```
import torch

torch.cuda.set_device(0) # Set to your desired GPU number


```

## 测试demo

```python

import supervision as sv
from ultralytics import YOLO
from tqdm import tqdm
import torch
print(torch.cuda.is_available())

# 检测
model = YOLO('yolov8n.pt')
source_video_path = "./det-videos/person.mp4"
target_video_path="./det-videos/aa2.mp4"
confidence_threshold = 0.3
iou_threshold = 0.7

def process_video_tracking():
	tracker = sv.ByteTrack()
	box_annotator = sv.BoxAnnotator()
	frame_generator = sv.get_video_frames_generator(source_path=source_video_path)
	video_info = sv.VideoInfo.from_video_path(video_path=source_video_path)
	print(video_info)
	for frame in tqdm(frame_generator, total=video_info.total_frames):
		results = model(
		    frame, verbose=False, conf=confidence_threshold, iou=iou_threshold
		)[0]
		detections = sv.Detections.from_ultralytics(results)
		detections = tracker.update_with_detections(detections)

		labels = [
		    f"#{tracker_id} {model.model.names[class_id]}"
		    for _, _, _, class_id, tracker_id in detections
		]

		annotated_frame = box_annotator.annotate(
		    scene=frame.copy(), detections=detections, labels=labels
		)
		    


		    


def process_video_camera():
	tracker = sv.ByteTrack()
	results = model.predict(
		source="0",
		verbose=True,## 打印额外信息
		show_labels=True,#在图中显示对象标签
		show_conf=True,#在图中显示对象置信度分数
		# classes=[0,2,3] ,#按照类别过滤	
		stream=True,##输入数据将以流的形式连续地传递给模型进行预测。这通常用于处理视频流，其中模型会在连续的帧上进行实时预测。并返回预测结果。 当此参数为True 则 show参数不生效
		# show=True,## 实时显示检测过程
		conf=confidence_threshold,##这个参数表示置信度阈值（confidence threshold）用于对象检测。它确定被检测到的对象被视为有效的最低置信度分数。置信度得分低于此阈值的对象将被过滤掉
		 iou=iou_threshold,##这个参数代表交并比（Intersection over Union，IoU）阈值。它确定预测的边界框和实际边界框之间所需的最小重叠程度，以将它们视为匹配。交并比得分低于此阈值的边界框将被视为不同的对象。iou_threshold的值被用作此参数的值。
		)
	for result in results:
		result = result.cpu()
		detections = sv.Detections.from_ultralytics(result)
		# detections = detections[detections.class_id == 0] # 过滤人
		# 使用提供的检测更新跟踪器并返回更新的检测结果。更新后的检测结果，现在包含跟踪 ID
		detections = tracker.update_with_detections(detections)
		print("执行解析")
		labels = [
			f"#{tracker_id} {model.model.names[class_id]}"
			for _, _, _, class_id, tracker_id in detections
		]
		
		print(labels)



	




	# with sv.VideoSink(target_path=target_video_path, video_info=video_info) as sink:
	# 	for frame in tqdm(frame_generator, total=video_info.total_frames):
	# 	    results = model(
	# 	        frame, verbose=False, conf=confidence_threshold, iou=iou_threshold
	# 	    )[0]
	# 	    detections = sv.Detections.from_ultralytics(results)
	# 	    detections = tracker.update_with_detections(detections)

	# 	    labels = [
	# 	        f"#{tracker_id} {model.model.names[class_id]}"
	# 	        for _, _, _, class_id, tracker_id in detections
	# 	    ]

	# 	    annotated_frame = box_annotator.annotate(
	# 	        scene=frame.copy(), detections=detections, labels=labels
	# 	    )

	# 	    sink.write_frame(frame=annotated_frame)


	

# 分类
# model = YOLO('yolov8n-cls.pt')
# result = model("./det-img/bus.jpg")[0]
# detections = sv.Detections.from_ultralytics(result)
# classifications = sv.Classifications.from_ultralytics(result)

# print(detections)
# video_info = sv.VideoInfo.from_video_path(video_path='./static/1行人检测测试视频.mp4')
# print(video_info)


if __name__ == '__main__':
	
	process_video_camera()
	pass



```


