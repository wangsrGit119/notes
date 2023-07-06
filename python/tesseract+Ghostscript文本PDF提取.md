## 基础步骤


### Tesseract
https://tesseract-ocr.github.io/tessdoc/
下载并安装Tesseract OCR库，windows可以从以下网址获取安装程序：https://github.com/UB-Mannheim/tesseract/wiki
linux 版本最新版本需要更新源 ： sources.list.d
安装完成后，将Tesseract的安装路径添加到系统环境变量中。
如果ubuntu安装5版本：


### Ghostscript
下载安装即可
https://ghostscript.com/releases/gsdnld.html

### 代码中调用
pdf分割-》ocr-》文本提取

### dockerfile

```
FROM python:3.9-slim-buster

RUN cp -a /etc/apt/sources.list /etc/apt/sources.list.bak
RUN sed -i "s@http://.*deb.debian.org@https://repo.huaweicloud.com@g" /etc/apt/sources.list 
RUN sed -i "s@http://.*archive.debian.org@http://repo.huaweicloud.com@g" /etc/apt/sources.list
RUN sed -i "s@http://.*security.debian.org@http://repo.huaweicloud.com@g" /etc/apt/sources.list


RUN apt-get -y update &&  apt-get install -y apt-transport-https ghostscript lsb-release wget gnupg gnupg2 gnupg1 mesa-utils

RUN echo "deb https://notesalexp.org/tesseract-ocr5/$(lsb_release -cs)/ $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/notesalexp.list

RUN wget -O - https://notesalexp.org/debian/alexp_key.asc --no-check-certificate | apt-key add -

RUN apt-get -y update && apt-get install -y tesseract-ocr tesseract-ocr-eng tesseract-ocr-chi-sim

WORKDIR /app
COPY requirements.txt /app/
RUN /usr/local/bin/python -m pip install --upgrade pip -i https://pypi.douban.com/simple

RUN pip install --no-cache-dir --upgrade pip -r requirements.txt -i https://pypi.douban.com/simple
RUN pip install python-multipart -i https://pypi.douban.com/simple
COPY . /app

CMD ["python", "tessertact-ocr.py"]

EXPOSE 18080

```

## 伪代码

```python
import os,sys,io
import tempfile
import cv2

import pytesseract
from datetime import datetime
from fastapi import FastAPI, Body, UploadFile
from fastapi.responses import JSONResponse, StreamingResponse
from fastapi.staticfiles import StaticFiles
from fastapi.middleware.cors import CORSMiddleware
import uvicorn
from concurrent.futures import ThreadPoolExecutor, as_completed
import multiprocessing
import uuid


app = FastAPI(docs_url="/documentation", redoc_url=None)
#------------------全局异常处理----------------------------
@app.exception_handler(Exception)
async def validation_exception_handler(request, err):
    # print(request)
    base_error_message = f"Failed to execute: {request.method}: {request.url}"
    return JSONResponse(status_code=400, content=getResErrData(f'{err}'))

#-----------------------------中间件--------------------------------------
app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_credentials=True, #如果为True则allow_origins不能为*
    allow_methods=['*'],
    allow_headers=['*'],
    expose_headers=['*']
)

def getResData(code,data,msg):
    return {'code':code,'data':data,'msg':msg}
def getResErrData(msg):
    return {'code':400,'data':None,'msg':msg}
def getResSuccessData(data):
    return {'code':200,'data':data,'msg':'success'}



cpu_count = multiprocessing.cpu_count()
print("系统CPU=========>",cpu_count)

if sys.platform.startswith('linux'):
    print('当前操作系统是Linux')
    gs_path = 'gs'
elif sys.platform == 'darwin':
    print('当前操作系统是macOS')
    gs_path = 'gs'
elif sys.platform == 'win32':
    print('当前操作系统是Windows')
    # 设置Ghostscript路径
    gs_path = 'D:/gs/gs10.00.0/bin/gswin64c'
    # 设置GTesseract路径
    pytesseract.pytesseract.tesseract_cmd = 'D:\\Program Files\\Tesseract-OCR\\tesseract'
else:
    print('无法判断当前操作系统')
    gs_path = 'gs'


langs = pytesseract.get_languages()
print("tesseract识别语言=========>", langs)
version = pytesseract.get_tesseract_version()
print("tesseract版本=========>", version)

# 配置Tesseract识别参数
custom_config = r'--oem 3 --psm 6'


async def pdf_ocr_extract(pdf_path):
    # 使用Ghostscript将PDF文件转换为图像
    with tempfile.TemporaryDirectory(dir='./temp/') as temp_dir:
        cmd = f'{gs_path} -dNOPAUSE -dBATCH -sDEVICE=png16m -r300 -sOutputFile={temp_dir}/%d.png {pdf_path}'
        os.system(cmd)
        # 使用Tesseract识别图像中的文本
        text = ''
        start_time = datetime.now()
        with ThreadPoolExecutor(max_workers=4) as executor:
            futures = []
            for i, filename in enumerate(sorted(os.listdir(temp_dir))):
                if filename.endswith('.png'):
                    img_path = os.path.join(temp_dir, filename)
                    future = executor.submit(ocr_image, img_path)
                    futures.append(future)
            for future in as_completed(futures):
                text += future.result()
        end_time = datetime.now()
        elapsed_time = end_time - start_time
        print(f"总共用时{elapsed_time.total_seconds()}秒")
    return text



def ocr_image(img_path):
    print(f"开始处理文件：{img_path}")
    start_time = datetime.now()
    img = cv2.imread(img_path)
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
    text = pytesseract.image_to_string(gray, lang='chi_sim', config=custom_config)
    end_time = datetime.now()
    elapsed_time = end_time - start_time
    print(f"文件{img_path}处理完成，用时{elapsed_time.total_seconds()}秒")
    return text


@app.post('/ocr/pdf',tags=["ocr"],summary="pdf文本提取")
async def importPdfData(file: UploadFile):
    text = ''
    with tempfile.TemporaryDirectory(dir='./temp/') as temp_dir:
        pdf_path = os.path.join(temp_dir, str(uuid.uuid1())+file.filename)
        with open(pdf_path, 'wb') as pdf:
            pdf.write(await file.read())
        text = await pdf_ocr_extract(pdf_path)
        return getResSuccessData(text)

@app.post('/ocr/image',tags=["ocr"],summary="图片文本提取")
async def importImageData(file: UploadFile):
    contents = await file.read()
    nparr = np.frombuffer(contents, np.uint8)
    img = cv2.imdecode(nparr, cv2.IMREAD_COLOR)
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
    text = pytesseract.image_to_string(gray, lang='chi_sim', config=custom_config)
    return getResSuccessData(text)


if __name__ == '__main__':
    uvicorn.run(app,host='0.0.0.0',port=18080)
    # 测试代码
    # pdf_path = 'bc.pdf'
    # text = pdf_ocr_extract(pdf_path)
    # print(text)


```

