
## 付费产品 buildjet

【官网】https://app.buildjet.com/
【使用】 Github授权登录

## 用法

1.在`GitHub`的`Action yaml`文件替换`runs-on`环境即可

2.可替换环境配置：https://buildjet.com/for-github-actions/docs/runners/hardware ，下面是截至 20230925的价格表

### AMD

| vCPU   | RAM   | YAML runner tag                 | Cost         |
|--------|-------|---------------------------------|--------------|
| 2 vCPU | 8 GB  | buildjet-2vcpu-ubuntu-2204      | $0.004 / min |
| 4 vCPU | 16 GB | buildjet-4vcpu-ubuntu-2204      | $0.008 / min |
| 8 vCPU | 32 GB | buildjet-8vcpu-ubuntu-2204      | $0.016 / min |
| 16 vCPU| 64 GB | buildjet-16vcpu-ubuntu-2204     | $0.032 / min |
| 32 vCPU| 64 GB | buildjet-32vcpu-ubuntu-2204     | $0.048 / min |

如果您想使用Ubuntu 20.04，请在runner tag中将2204替换为2004。

### ARM

| vCPU   | RAM   | YAML runner tag                   | Cost         |
|--------|-------|-----------------------------------|--------------|
| 2 vCPU | 3 GB  | buildjet-2vcpu-ubuntu-2204-arm    | $0.004 / min |
| 4 vCPU | 6 GB  | buildjet-4vcpu-ubuntu-2204-arm    | $0.008 / min |
| 8 vCPU | 12 GB | buildjet-8vcpu-ubuntu-2204-arm    | $0.016 / min |
| 16 vCPU| 24 GB | buildjet-16vcpu-ubuntu-2204-arm   | $0.032 / min |
| 32 vCPU| 48 GB | buildjet-32vcpu-ubuntu-2204-arm   | $0.064 / min |


