# 🎯 JetKVM X86_64 HDMI 输出功能指南

## 🚀 **功能概述**

JetKVM X86_64 版本现在支持完整的 HDMI 输出功能！你可以将 JetKVM 的 Web 界面直接显示在连接的 HDMI 显示器上，实现真正的本地 KVM 体验。

## ✨ **主要特性**

### 🖥️ **HDMI 显示输出**
- ✅ **全屏显示** - JetKVM Web 界面在 HDMI 显示器上全屏显示
- ✅ **自动配置** - 自动检测和配置 HDMI 输出设备
- ✅ **高分辨率** - 支持 1920x1080 分辨率
- ✅ **实时更新** - 界面实时同步，支持远程和本地同时操作

### 🔧 **技术实现**
- **X Server** - 自动启动和配置 X 服务器
- **Chromium 浏览器** - Kiosk 模式全屏显示
- **DRM/Framebuffer** - 直接硬件显示输出
- **自动安装** - 自动安装所需软件包

### 🌐 **Web API 控制**
- `GET /hdmi/status` - 获取 HDMI 输出状态
- `POST /hdmi/enable` - 启用 HDMI 输出
- `POST /hdmi/disable` - 禁用 HDMI 输出
- `POST /hdmi/toggle` - 切换 HDMI 输出状态
- `GET/POST /hdmi/config` - HDMI 配置管理

## 📋 **系统要求**

### **支持平台**
- ✅ **Linux X86_64** - 完全支持
- ❌ **其他平台** - 暂不支持

### **硬件要求**
- 🖥️ **HDMI 输出端口** - 系统必须有 HDMI 输出
- 💾 **显示驱动** - 支持 DRM 的显卡驱动
- 🔌 **HDMI 显示器** - 连接的 HDMI 显示器

### **软件依赖**
程序会自动安装以下软件包：
- `xorg` - X Window 系统
- `openbox` - 轻量级窗口管理器
- `chromium-browser` - 浏览器
- `unclutter` - 隐藏鼠标光标
- `xdotool` - X Window 工具

## 🚀 **使用方法**

### **方法1：通过 Web API 启用**

```bash
# 启用 HDMI 输出
curl -X POST http://10.10.10.12:8080/hdmi/enable

# 检查状态
curl http://10.10.10.12:8080/hdmi/status

# 禁用 HDMI 输出
curl -X POST http://10.10.10.12:8080/hdmi/disable
```

### **方法2：通过配置文件**

编辑 `/userdata/kvm_config.json`：

```json
{
  "hdmi_output_enabled": true,
  "hdmi_output_auto_start": true
}
```

### **方法3：通过 Web 界面**

1. 访问 JetKVM Web 界面
2. 进入设置页面
3. 找到 "HDMI 输出" 选项
4. 启用 HDMI 输出功能

## 🔧 **配置选项**

### **配置参数**

| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `hdmi_output_enabled` | boolean | `false` | 是否启用 HDMI 输出 |
| `hdmi_output_auto_start` | boolean | `false` | 是否开机自动启动 |

### **API 响应示例**

```json
{
  "enabled": true,
  "available": true,
  "display_device": "card0-Virtual-1",
  "resolution": "1920x1080",
  "web_url": "http://localhost:8080"
}
```

## 🛠️ **故障排除**

### **常见问题**

#### **1. HDMI 输出不工作**
```bash
# 检查显示设备
ls /sys/class/drm/

# 检查 X 服务器日志
tail -f /tmp/jetkvm-xorg.log

# 检查进程状态
ps aux | grep -E "(Xorg|chromium)"
```

#### **2. 显示器无信号**
- 确保 HDMI 线缆连接正确
- 检查显示器输入源设置
- 尝试重启 HDMI 输出：
```bash
curl -X POST http://10.10.10.12:8080/hdmi/disable
curl -X POST http://10.10.10.12:8080/hdmi/enable
```

#### **3. 浏览器启动失败**
```bash
# 手动启动浏览器测试
DISPLAY=:1 chromium-browser --kiosk http://localhost:8080
```

### **日志查看**

```bash
# JetKVM 主日志
journalctl -f -u jetkvm

# X 服务器日志
tail -f /tmp/jetkvm-xorg.log

# 系统显示相关日志
dmesg | grep -i drm
```

## 🎯 **使用场景**

### **1. 本地 KVM 控制台**
- 将 JetKVM 连接到显示器
- 直接在显示器上操作 KVM 功能
- 无需额外的电脑或网络连接

### **2. 双模式操作**
- HDMI 本地显示 + 网络远程访问
- 本地和远程可同时操作
- 灵活的使用方式

### **3. 演示和培训**
- 大屏幕展示 KVM 操作
- 培训和演示用途
- 会议室部署

## 📊 **性能特点**

- **启动时间**: 约 5-10 秒
- **分辨率**: 1920x1080@30fps
- **延迟**: 极低延迟本地显示
- **资源占用**: 轻量级实现

## 🔄 **自动化集成**

### **开机自启动**
```json
{
  "hdmi_output_enabled": true,
  "hdmi_output_auto_start": true
}
```

### **脚本控制**
```bash
#!/bin/bash
# 自动启用 HDMI 输出脚本

# 等待系统就绪
sleep 10

# 启用 HDMI 输出
curl -X POST http://localhost:8080/hdmi/enable

echo "HDMI 输出已启用"
```

## 🎉 **总结**

JetKVM X86_64 的 HDMI 输出功能为你提供了：

- 🖥️ **本地显示** - 直接在 HDMI 显示器上使用
- 🌐 **远程访问** - 网络访问功能保持不变
- 🔧 **灵活配置** - 多种启用和配置方式
- 🚀 **自动化** - 开机自启动和 API 控制

现在你的 JetKVM 不仅是一个网络 KVM 解决方案，更是一个完整的本地+远程双模式 KVM 系统！

---

**🎯 享受你的新 HDMI 输出功能吧！** 🚀