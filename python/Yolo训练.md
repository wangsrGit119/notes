

## paddleyolo
https://github.com/PaddlePaddle/PaddleYOLO
https://gitee.com/monkeycc/PaddleYOLO?_from=gitee_search

## 数据集
1.垃圾
https://github.com/Nicola115/garbage_classification
https://github.com/pedropro/TACO

## 版本对比

| yolo版本 | 简要概述 | 官网GitHub链接 | 性能对比 | 场景推荐使用 | 产出公司 |
| --- | --- | --- | --- | --- | --- |
| YOLOv1 | 奠定了整个YOLO系列的基础，使用单个卷积神经网络同时预测多个边界框和类别。 | https://github.com/pjreddie/darknet | 速度快，但精度较低。适用于实时性要求高的场景，如人脸检测。 | 人脸检测、行人检测、车辆检测等。 | Joseph Redmon |
| YOLOv2 | 采用了Batch Normalization和High Resolution Classifier，提高了检测精度。 | https://github.com/pjreddie/darknet | 精度提高，速度略有下降。适用于对精度要求较高的场景，如自动驾驶。 | 自动驾驶、安防监控等。 | Joseph Redmon |
| YOLOv3 | 采用了多尺度预测，使用了残差块和特征金字塔网络，提高了检测精度。 | https://github.com/pjreddie/darknet | 精度进一步提高，速度略有下降。适用于对精度要求更高的场景，如智能家居、智能医疗等。 | 智能家居、智能医疗、智能交通等。 | Joseph Redmon |
| YOLOv4 | 采用了CSPNet和SPP结构，使用了Mish激活函数和DropBlock正则化，提高了检测精度。 | https://github.com/AlexeyAB/darknet | 精度进一步提高，速度略有下降。适用于对精度要求更高的场景，如智能家居、智能医疗等。 | 智能家居、智能医疗、智能交通等。| Alexey Bochkovskiy |
| YOLOv5 | 采用了自适应卷积、Swish激活函数、PANet特征融合等技术，提高了检测速度和精度。 | https://github.com/ultralytics/yolov5 | 精度和速度都有所提升。适用于对速度和精度都有要求的场景，如无人机航拍、工业质检等。 | 无人机航拍、工业质检、安防监控等。| Ultralytics |
| YOLOv6 | 采用了SAM模块、CSP模块、FSAF损失函数等技术，提高了检测速度和精度。 | [https://github.com/WongKinYiu/yolov6 ](https://github.com/meituan/YOLOv6)| 精度和速度都有所提升。适用于对速度和精度都有要求的场景，如无人机航拍、工业质检等。| Wong Kin Yiu |
| YOLOv7| 采用了CSP模块|
| YOLOv8| 采用了CSP模块|https://docs.ultralytics.com/quickstart/#install

## yolov8 教程
https://github.com/bubbliiiing/yolov8-pytorch
https://docs.ultralytics.com/quickstart/#use-with-python
https://blog.csdn.net/u012863603/article/details/128706974
https://aitechtogether.com/article/48256.html 
https://docs.ultralytics.com/usage/python/

## 基本环境 
py 1.11.3

## start
pip install ultralytics

利用conda安装出现下列警告：(将对应的加到path即可)
 WARNING: The script tqdm.exe is installed in 'C:\Users\12156\AppData\Roaming\Python\Python311\Scripts' which is not on PATH.
  Consider adding this directory to PATH or, if you prefer to suppress this warning, use --no-warn-script-location.
  WARNING: The script isympy.exe is installed in 'C:\Users\12156\AppData\Roaming\Python\Python311\Scripts' which is not on PATH.
  Consider adding this directory to PATH or, if you prefer to suppress this warning, use --no-warn-script-location.
  WARNING: The scripts fonttools.exe, pyftmerge.exe, pyftsubset.exe and ttx.exe are installed in 'C:\Users\12156\AppData\Roaming\Python\Python311\Scripts' which is not on PATH.
  Consider adding this directory to PATH or, if you prefer to suppress this warning, use --no-warn-script-location.
  WARNING: The script normalizer.exe is installed in 'C:\Users\12156\AppData\Roaming\Python\Python311\Scripts' which is not on PATH.
  Consider adding this directory to PATH or, if you prefer to suppress this warning, use --no-warn-script-location.
  WARNING: The scripts convert-caffe2-to-onnx.exe, convert-onnx-to-caffe2.exe and torchrun.exe are installed in 'C:\Users\12156\AppData\Roaming\Python\Python311\Scripts' which is not on PATH.
  Consider adding this directory to PATH or, if you prefer to suppress this warning, use --no-warn-script-location.
  WARNING: The scripts ultralytics.exe and yolo.exe are installed in 'C:\Users\12156\AppData\Roaming\Python\Python311\Scripts' which is not on PATH.
  Consider adding this directory to PATH or, if you prefer to suppress this warning, use --no-warn-script-location.
S

## 模型


| Model                                                                                | size<br><sup>(pixels) | mAP<sup>val<br>50-95 | Speed<br><sup>CPU ONNX<br>(ms) | Speed<br><sup>A100 TensorRT<br>(ms) | params<br><sup>(M) | FLOPs<br><sup>(B) |
| ------------------------------------------------------------------------------------ | --------------------- | -------------------- | ------------------------------ | ----------------------------------- | ------------------ | ----------------- |
| [YOLOv8n](https://github.com/ultralytics/assets/releases/download/v0.0.0/yolov8n.pt) | 640                   | 37.3                 | 80.4                           | 0.99                                | 3.2                | 8.7               |
| [YOLOv8s](https://github.com/ultralytics/assets/releases/download/v0.0.0/yolov8s.pt) | 640                   | 44.9                 | 128.4                          | 1.20                                | 11.2               | 28.6              |
| [YOLOv8m](https://github.com/ultralytics/assets/releases/download/v0.0.0/yolov8m.pt) | 640                   | 50.2                 | 234.7                          | 1.83                                | 25.9               | 78.9              |
| [YOLOv8l](https://github.com/ultralytics/assets/releases/download/v0.0.0/yolov8l.pt) | 640                   | 52.9                 | 375.2                          | 2.39                                | 43.7               | 165.2             |
| [YOLOv8x](https://github.com/ultralytics/assets/releases/download/v0.0.0/yolov8x.pt) | 640                   | 53.9                 | 479.1                          | 3.53                                | 68.2               | 257.8             |


### 基础用法

https://docs.ultralytics.com/usage/python/

### 数据标注工具


[pip install labelme](https://github.com/wkentaro/labelme/releases)


### 执行下面代码会自动下载所需的模型

1.预测
> 会自动检测模型文件的存在性，不存在则下载
> 预测示例，运行完成会在当前目录的 runs文件夹输出

```
from ultralytics import YOLO
from PIL import Image
import cv2
model = YOLO('./models/yolov8n.pt')
im1 = Image.open("./images/bus.jpg")
results = model.predict(source=im1, save=True)  # save plotted images

```
2.训练


