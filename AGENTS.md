# AGENTS.md - CFData-WEB 开发指南

## Go 工具链
- 路径: `/opt/homebrew/bin/go`（本地开发）
- 版本: Go 1.26.3

## 编译测试版
每次修改完成后，主动编译测试版到 `release_assets/cfdata-test`，覆盖原文件：
```bash
mkdir -p /root/project/CFData-WEB/release_assets && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 /opt/homebrew/bin/go build -trimpath -ldflags="-s -w" -o /tmp/cfdata-test-tmp ./cmd/cfdata/ && mv /tmp/cfdata-test-tmp /root/project/CFData-WEB/release_assets/cfdata-test
```

## 本地开发编译
```bash
go build -trimpath -ldflags="-s -w" -o cfdata ./cmd/cfdata/
```

## 运行
```bash
go run ./cmd/cfdata/
```

## Git 操作规则
- **禁止**在未获得明确许可前执行 `git commit` 或 `git push`
- **禁止**在任何情况下操作 `main` 主分支，主分支仅限用户本人操作
- 所有修改均在 `beta` 等非主分支上完成

## 禁止修改的文件
- `.github/` 目录下的任何文件（GitHub Actions 配置）
- 顶层目录中注释包含"原版源码"的文件
