# JetKVM 发布指南

## 🚀 自动发布流程

JetKVM 现在支持自动构建和发布到 GitHub Releases！

### 📦 发布内容

每个发布版本包含：
- **jetkvm-arm-linux**: ARM 设备二进制文件（硬件 KVM）
- **jetkvm-x86_64-linux**: X86_64 设备二进制文件（软件模拟）
- **SHA256 校验文件**: 用于验证下载完整性

### 🔧 创建发布

#### 方法 1: 使用发布脚本（推荐）

```bash
# 创建新版本发布
./scripts/create_release.sh v1.0.0

# 脚本会自动：
# 1. 更新版本号
# 2. 创建 Git 标签
# 3. 推送到 GitHub
# 4. 触发自动构建和发布
```

#### 方法 2: 手动创建标签

```bash
# 创建并推送标签
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0

# GitHub Actions 会自动构建和发布
```

#### 方法 3: 手动触发工作流

1. 访问 GitHub Actions 页面
2. 选择 "Release" 工作流
3. 点击 "Run workflow"
4. 输入版本号（如 v1.0.0）
5. 点击运行

### 📋 发布流程

1. **前端构建**: 构建 React UI
2. **多架构编译**: 
   - ARM Linux 二进制文件
   - X86_64 Linux 二进制文件
3. **生成校验和**: SHA256 文件
4. **创建发布**: 自动上传到 GitHub Releases
5. **发布说明**: 自动生成详细的发布说明

### 🔍 监控构建

- **GitHub Actions**: https://github.com/你的用户名/kvm/actions
- **发布页面**: https://github.com/你的用户名/kvm/releases

### 📥 下载和使用

用户可以从 Releases 页面下载：

```bash
# 下载 X86_64 版本
wget https://github.com/你的用户名/kvm/releases/download/v1.0.0/jetkvm-x86_64-linux
chmod +x jetkvm-x86_64-linux

# 验证校验和
wget https://github.com/你的用户名/kvm/releases/download/v1.0.0/jetkvm-x86_64-linux.sha256
sha256sum -c jetkvm-x86_64-linux.sha256

# 运行
./jetkvm-x86_64-linux
```

### 🏗️ 架构支持

| 架构 | 平台 | 功能 | 适用场景 |
|------|------|------|----------|
| ARM | Linux | 完整硬件 KVM | 生产环境、嵌入式设备 |
| X86_64 | Linux | 软件模拟 KVM | 开发、测试、演示 |

### 🐛 故障排除

如果构建失败：

1. 检查 GitHub Actions 日志
2. 确认所有依赖已正确配置
3. 验证构建约束和架构特定代码
4. 检查 Makefile 和构建脚本

### 📚 相关文档

- [X86_64 使用指南](README_X86.md)
- [构建问题解决方案](BUILD_FIXES_X86.md)
- [实现变更日志](CHANGELOG_X86.md)

## 🎯 版本规范

使用语义化版本控制：
- `v1.0.0`: 主要版本（重大变更）
- `v1.1.0`: 次要版本（新功能）
- `v1.0.1`: 补丁版本（错误修复）

## 🔐 安全注意事项

- 所有发布的二进制文件都包含 SHA256 校验和
- 建议用户验证下载文件的完整性
- 发布流程通过 GitHub Actions 自动化，确保构建环境的一致性