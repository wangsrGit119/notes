
## 基础文档
https://github.com/hiyouga/LLaMA-Factory/blob/main/README_zh.md


## 选择微调模型

Qwen-1.8B-Chat


## 模型导出为 旧格式 .bin的

## 转换为 guff


python llama.cpp/convert.py qw-1.8-suc --outfile qw-1.8-suc.gguf --vocab-type bpe --pad-vocab
 --outtype q8_0 可选参数
 
https://github.com/ggerganov/llama.cpp/discussions/2948


## ollama 加载

```
FROM ./qwen1_5-0_5b-chat-q5_k_m.gguf

TEMPLATE """{{ if .System }}<|im_start|>system
{{ .System }}<|im_end|>{{ end }}<|im_start|>user
{{ .Prompt }}<|im_end|>
<|im_start|>assistant
"""

PARAMETER stop "<|im_start|>"
PARAMETER stop "<|im_end|>"

```

