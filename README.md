## sshctl

### 依赖

---
linux  go环境

### 简要功能说明（详情看帮助）

---
- 支持隐藏密码输入执行远程命令
- 支持批量主机执行多条命令等
- 支持主机地址和命令文件配置执行
- 支持远程复制文件内容到本地
- 支持远程主机ping功能


### 实现需求

---
- 此工具用go实现要跨平台开箱即用，木有依赖
- 简单无脑，无需学习，直接堆命令行即可
- 支持并发

### 使用说明

---

```bash
git clone https://github.com/LiuXiangBiao/sshctl.git
```
```bash
go mod tidy
```
```bash
go build
```
```bash
mv sshctl /gopath/bin/
```

---

#### 详情 sshctl -h 获取帮助信息
