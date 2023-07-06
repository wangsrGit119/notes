
## 环境
python>=3.8

pip install protobuf==3.20.0
pip install onnx==1.11.0



python -m pip install paddlepaddle==2.4.2 -i https://pypi.tuna.tsinghua.edu.cn/simple
pip install "paddleocr>=2.0.1" # Recommend to use version 2.0.1+

## numpy报错注意
pip uninstall -y numpy
pip install "Numpy==1.23.5"


## demo
```
from paddleocr import PaddleOCR
ocr = PaddleOCR(enable_mkldnn=True,use_angle_cls=True, lang='ch',det_model_dir='./models/ch_PP-OCRv3_det_infer.tar', rec_model_dir='./models/ch_PP-OCRv3_rec_infer.tar', cls_model_dir='./models/ch_ppocr_mobile_v2.0_cls_infer.tar')
ocr.rec_batch_num = 10
ocr.use_mp = Trueimg_path = './1.png'
result = ocr.ocr(img_path, cls=True)
for idx in range(len(result)):
    res = result[idx]
    for line in res:
        print(line[1][0])

```


