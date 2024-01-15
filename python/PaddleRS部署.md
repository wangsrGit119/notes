



##  注意事项

paddleRS 可本地下载后切换到对应分支然后上传的服务器再进行安装依赖 同理 paddleslim也是

## cuda 12版本


## 快速安装CUDA12.1或者 11.x版本

```
conda create -n pdrs-py39 python=3.9
conda activate pdrs-py39

git clone https://github.com/PaddlePaddle/PaddleRS
cd PaddleRS
git checkout develop

pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple/
python -m pip install paddlepaddle-gpu==2.6.0 -i https://pypi.tuna.tsinghua.edu.cn/simple

pip install -r requirements.txt
pip install .

conda install gdal



```

## paddleslim 拉取不下来

```
git clone --filter=blob:none --quiet https://github.com/PaddlePaddle/PaddleSlim.git
cd PaddleSlim
git checkout -b release/2.5 --track origin/release/2.5

```
然后上传到服务器 根目录执行

```
pip install -r requirements.txt
pip setup.py install
```

## 错误


 - 缺失 libSM.so.6

 https://stackoverflow.com/questions/47113029/importerror-libsm-so-6-cannot-open-shared-object-file-no-such-file-or-directo
 ```
 apt install libsm6 libxext6 libxrender-dev
 
 ```

 - 缺失cv2属性

 > partially initialized module 'cv2' has no attribute '_registerMatType' (most likely due to a circular import)
 > ttributeError: module 'cv2' has no attribute 'INTER_NEAREST'     

 ```
 opencv-python 4.0.9.80
 opencv-contrib-python 4.2.0.32
 ```
 - 提示cuda api版本不对 
```
export LD_LIBRARY_PATH=/usr/local/cuda-11.7/lib64:$LD_LIBRARY_PATH
```

## 运行测试用例

```python
python tutorials/train/image_restoration/lesrcnn.py
```
