
## 启动APP后同时 执行脚本

```
const path = require('path');
const fs = require('fs-extra');
let distfiles = path.resolve(__dirname, '../src/assets/bat/');

const { spawn,exec } = require('child_process');

function executeScriptOnAppStart(targetFolderPath) {
  const scriptPath = path.join(targetFolderPath, 'server.js');
  const scriptProcess = spawn('node', [scriptPath]);
  scriptProcess.stdout.on('data', (data) => {
    console.log(`Script脚本输出：${data}`);
  });

  scriptProcess.stderr.on('data', (data) => {
    console.error(`Script脚本错误：${data}`);
  });

  scriptProcess.on('exit', (code) => {
    console.log(`Script脚本退出，退出码：${code}`);
  });

  app.on('before-quit', () => {
    scriptProcess.kill();
  });
}

function executeBatScriptOnAppStart(targetFolderPath) {
  const scriptPath = path.join(targetFolderPath, 'test.bat');
  const batProcess = exec(scriptPath);

  batProcess.stdout.on('data', (data) => {
    console.log(`Bat脚本输出：${data}`);
  });

  batProcess.stderr.on('data', (data) => {
    console.error(`Bat脚本错误：${data}`);
  });

  batProcess.on('exit', (code) => {
    console.log(`Bat脚本退出，退出码：${code}`);
  });

  app.on('before-quit', () => {
    batProcess.kill();
  });
}

function createTargetDir(){
	const folderName = 'testdir';
	const installationPath = path.dirname(app.getPath('exe'));
	const folderPath = path.join(installationPath, folderName);
	// 检查文件夹是否已存在，如果不存在则创建
	if (!fs.existsSync(folderPath)) {
	fs.mkdirSync(folderPath);
	}
	
	const targetFolderPath = path.join(installationPath, folderName); // 目标文件夹路径
	fs.copySync(distfiles, targetFolderPath);
	executeScriptOnAppStart(targetFolderPath)
	executeBatScriptOnAppStart(targetFolderPath)

}
```
