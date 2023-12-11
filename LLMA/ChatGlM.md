

## chatglm 官方仓库
https://github.com/THUDM/ChatGLM3(推荐使用)


https://github.com/THUDM/ChatGLM-6B

## 基础环境

git、git lfs、curl

### git lfs linux安装

wget https://github.com/git-lfs/git-lfs/releases/download/v2.13.3/git-lfs-linux-amd64-v2.13.3.tar.gz
tar xvfz git-lfs-linux-amd64-v2.13.3.tar.gz
cd git-lfs-2.13.3
sudo ./install.sh
git lfs version


### 模型下载
GIT_LFS_SKIP_SMUDGE=1
git clone https://huggingface.co/THUDM/chatglm-6b

https://blog.csdn.net/Hello_World1023/article/details/130356044

### 其他依赖

基础依赖：
pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple/
pip install transformers==4.27.1 -i https://pypi.tuna.tsinghua.edu.cn/simple/
pip install gradio -i https://pypi.tuna.tsinghua.edu.cn/simple/



### 结合 langchain


python3.8

https://github.com/thomas-yanxin/LangChain-ChatGLM-Webui
https://openi.pcl.ac.cn/Learning-Develop-Union/LangChain-ChatGLM-Webui.git

安装教程：https://github.com/thomas-yanxin/LangChain-ChatGLM-Webui/blob/master/docs/deploy.md

pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple/
