
## cuda 和cudnn版本对照

https://docs.nvidia.com/deeplearning/cudnn/reference/support-matrix.html


##  下面是一个Markdown表格，列出了英伟达gpu的具体参数 来源 chatgpt

| GPU型号       | 架构     | CUDA核心数量 | Tensor核心数量 | 内存容量      | 单精度性能 (TFLOPS) | 双精度性能 (TFLOPS) |
| ----------- | ------ | -------- | ---------- | --------- | -------------- | -------------- |
| A100        | Ampere | 6912     | 6912       | 40GB/80GB | 19.5           | 9.7            |
| RTX 4090    | Ampere | 10496    | 328        | 24GB      | 35.6           | 10.6           |
| RTX 3090    | Ampere | 10496    | 328        | 24GB      | 35.6           | 10.6           |
| RTX 3080    | Ampere | 8704     | 272        | 10GB/20GB | 29.8           | 9.4            |
| RTX 3070    | Ampere | 5888     | 184        | 8GB/16GB  | 20.4           | 6.5            |
| V100        | Volta  | 5120     | 640        | 16GB/32GB | 15.7           | 7.8            |
| GTX 1080 Ti | Pascal | 3584     | N/A        | 11GB      | 11.3           | 0.35           |
| T4          | Turing | 2560     | 320        | 16GB      | 8.1            | 0.25           |
| GTX 1060    | Pascal | 1280     | N/A        | 3GB/6GB   | 3.9            | 0.12           |

## 系统GPU使用查看

```shell
nvidia-smi
```

这将显示当前系统中所有 NVIDIA GPU 的详细信息，包括 GPU 的使用率、内存使用情况、进程占用 GPU 的信息等。
```shell
pip install gpustat
```

安装完成后，在终端中运行以下命令来查看 GPU 使用情况：

```shell
gpustat

```
这将显示当前系统中所有 NVIDIA GPU 的简要使用情况。


## cuda版本


https://pytorch.org/get-started/locally/

### 推荐conda 安装

conda install pytorch torchvision torchaudio pytorch-cuda=12.1 -c pytorch -c nvidia
conda install pytorch torchvision torchaudio pytorch-cuda=11.8 -c pytorch -c nvidia
