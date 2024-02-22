
## 飞桨 官网

https://github.com/PaddlePaddle/PaddleOCR

## 飞桨OCR常见问题

https://github.com/PaddlePaddle/PaddleOCR/blob/release/2.7/doc/doc_ch/FAQ.md

>最常见的比如 ：相机采集的图像为四通道，应该如何处理？
>A: 有两种方式处理：
>如果没有其他需要，可以在解码数据的时候指定模式为三通道，例如如果使用opencv，可以使用cv::imread(img_path, cv::IMREAD_COLOR)。
>如果其他模块需要处理四通道的图像，那也可以在输入PaddleOCR模块之前进行转换，例如使用cvCvtColor(&img,img3chan,CV_RGBA2RGB)。

### 环境
python>=3.8

```
pip install protobuf==3.20.0
pip install onnx==1.11.0

python -m pip install paddlepaddle -i https://pypi.tuna.tsinghua.edu.cn/simple
pip install "paddleocr>=2.0.1" # Recommend to use version 2.0.1+
```
### numpy报错注意

```
pip uninstall -y numpy
pip install "Numpy==1.23.5"
```

### demo

> 不用自己下载 指定文件夹后会自动下载模型
> `enable_mkldnn=True,use_angle_cls=True`参数启用后，暂时在paddle 2.5版本会丢精度 但是 paddlepaddle 2.4正常截至 目前2023 12 4号
```python
from paddleocr import PaddleOCR
ocr = PaddleOCR(enable_mkldnn=True,use_angle_cls=True, lang='ch',det_model_dir='./models/ch_PP-OCRv3_det_infer.tar', rec_model_dir='./models/ch_PP-OCRv3_rec_infer.tar', cls_model_dir='./models/ch_ppocr_mobile_v2.0_cls_infer.tar')
ocr.rec_batch_num = 10
ocr.use_mp = True
img_path = './1.png'
result = ocr.ocr(img_path, cls=True)
for idx in range(len(result)):
    res = result[idx]
    for line in res:
        print(line[1][0])

```

## 推荐 FastDeploy部署使用

https://github.com/PaddlePaddle/FastDeploy/tree/develop/examples/vision/ocr/PP-OCR/cpu-gpu/python

## 训练自己数据集

- 识别训练

https://gitee.com/paddlepaddle/PaddleOCR/blob/release/2.1/doc/doc_ch/recognition.md

https://zhuanlan.zhihu.com/p/435118207


- 环境

   paddlepaddle 2.5
   conda install -c conda-forge libstdcxx-ng
   git clone  paddleOCr
   pip install -r requirement.txt
  







