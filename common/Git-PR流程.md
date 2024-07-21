
## 如果你从同一项目存储库中的分支创建了拉取请求，那么：

- 确保你已经检查了你的分支（来自 PR 的分支）：`git checkout your-branch`
- 你的分支签出后，你应该做一个`git pull origin main`
- 然后 ` git push origin your-branch `更新 PR。

## 如果您分叉了一个 repo、创建了一个分支并提交了 PR，请按照以下步骤操作：

- 使用原始项目仓库创建一个远程仓库：`git remote add upstream 'url.git.here'`
- 确保你已经检查过你的分支：`git checkout your-branch`
- 从上游获取最新的更改到你的分支：`git pull upstream main`
- 此后，推动从上游获得的更改：`git push origin your-branch`
- 最后，您可以前往 GitHub 页面，确保没有其他内容out-of-date阻止您的 PR。
- 此后，您应该会看到您的 PR 已全部完成，可以进行合并（设置评论之后）。
