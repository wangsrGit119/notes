

## 简易安装 

https://space.bilibili.com/12566101/

##  使用 google colab部署
https://github.com/camenduru/stable-diffusion-webui-colab

## 在线提示词
https://www.thomas.io/stable-diffusion-prompt-generator

## 模型概要
Stable diffusion v1.4和v1.5：可以作为通用模型使用，用于生成真实的图像。

F222：适用于肖像，有生成裸体的倾向，可以在提示中包含“裙子”和“牛仔裤”等衣服术语。

Anything V3：适用于生成高质量的动漫风格图像，可以在文本提示中使用danbooru标签。

Open Journey：适用于一般用途，可以使用mdjrny-v4样式的触发关键字。

Dreamshaper：适用于生成介于逼真和计算机图形之间的肖像插图风格。

ChilloutMix：适用于生成高质量的亚洲女性照片，可以与韩国嵌入ulzzang-6500-v1一起使用，生成像韩国流行音乐女孩一样的女孩。

Waifu-diffusion：适用于生成日本动漫风格。

Robo Diffusion：适用于将每个主题转换为机器人。

Mo-di-diffusion：适用于生成类似于皮克斯风格的图像

Momoko 模型是一种人工智能模型，它是根据日本画家MOMOKO的风格训练的


## 非docker 自己安装


1.conda环境 -》python 3.10.6
conda create -p D:\ProgramData\miniconda3\envs\sd_3106 python=3.10.6
1. git clone https://github.com/AUTOMATIC1111/stable-diffusion-webui
2. 设置pip镜像源 pip config set global.index-url https://mirrors.aliyun.com/pypi/simple/
3. 先升级对应环境的pip 比如：D:\python-workspace\stable-diffusion-webui\venv\Scripts\python.exe -m pip install --upgrade pip
4. 直接执行：webui-user.bat （执行过程中可以挂代理，否则报错按照下面报错指南链接更改加速地址）
6. pip install xformer 
7. 查看显卡信息 cmd输入：wmic path win32_VideoController get name 或输入 dxdiag 回车
8. 如何是NVIDIA 则需下载 cuda；如果是：AMD Radeon Graphics （NViD的GPU架构和cuda对应 可用网上查看对应关系）
9.  cuda 下载位置 https://developer.nvidia.com/cuda-downloads

## 报错指南
上面第四报错：
```
 error: subprocess-exited-with-error
  × python setup.py egg_info did not run successfully.
```
可能是权限问题，管理员权限打开 conda 同时打开代理 全局模式或者https://savokiss.com/tech/sdw-install-error.html


## docker
https://github.com/AbdBarho/stable-diffusion-webui-docker

> git clone https://github.com/AbdBarho/stable-diffusion-webui-docker
> cd stable-diffusion-webui-docker
> docker-compose --profile download up --build
> docker-compose --profile auto-cpu up --build

## AMD GPU运行

https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Install-and-Run-on-AMD-GPUs


## prompt
https://mpost.io/best-100-stable-diffusion-prompts-the-most-beautiful-ai-text-to-image-prompts/


## 相关教程
[controlnet教程](https://www.uisdc.com/controlnet)

[controlnet教程2](https://www.uisdc.com/controlnet-2)

https://mp.weixin.qq.com/s?__biz=MzI2NTQ0MjY5Nw==&mid=2247484878&idx=1&sn=8d60fb4f6acc075c9ff77cb77a0cc0c0&chksm=ea9c0128ddeb883e13d7c7d52ef2b180323479679a0008f1ed3b8099ddee993da23f4f55ccc6&scene=21#wechat_redirect

[cn模型下载](https://huggingface.co/lllyasviel/ControlNet/tree/main/models)

## B端风格模型
https://civitai.com/models/38558/ddiconlora

https://civitai.com/models/100056/bdicon

