

## 官方模型

https://github.com/openai/whisper




## 实时转写

https://github.com/davabase/whisper_real_time

## 基础教程

```python
import whisper
# 指定模型下载位置
download_path = "./models"
model = whisper.load_model("small",download_root=download_path)

# 前置提示词 可以纠正中文简体繁体

prompt='以下是普通话的句子。'
# 输出语音的文本
result = model.transcribe("1682489036846_transMedia.mp3",language="zh",verbose=True,initial_prompt=prompt)
print(result["text"])

```

## 参数

