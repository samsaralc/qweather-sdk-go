# 和风天气 Go SDK 测试示例

本目录包含了和风天气Go SDK的各种测试示例，帮助您了解如何使用SDK的各项功能。

## 文件说明

- `test_runner.go` - 测试运行器主程序
- `grid_weather_test.go` - 格点实时天气测试用例
- `weather_test.go` - 实时天气测试用例  
- `auth_test.go` - 身份认证测试用例

## 使用方法

### 1. 运行所有测试

```bash
cd example
go run *.go
```

### 2. 运行特定测试

```bash
# 只测试格点实时天气
go run *.go -test=grid

# 只测试实时天气
go run *.go -test=weather

# 只测试身份认证
go run *.go -test=auth

# 显示帮助信息
go run *.go -help
```

## 准备工作

在运行测试之前，请确保：

1. **获取API凭据**: 在[和风天气开发者平台](https://dev.qweather.com/)注册并获取：
   - JWT Token（推荐）
   - 或 API Key

2. **替换Token**: 将测试代码中的占位符替换为您的实际凭据：
   - `your_jwt_token_here` → 您的JWT Token
   - `your_api_key_here` → 您的API Key

3. **网络连接**: 确保可以访问和风天气API服务器

## 测试内容

### 格点实时天气测试
- 使用经纬度字符串获取天气
- 使用经纬度数值获取天气
- 可选参数测试（语言、单位）
- 错误处理测试
- 不同认证方式测试

### 实时天气测试
- 使用LocationID获取天气
- 使用经纬度坐标获取天气
- 便利方法测试
- 可选参数测试
- 错误处理测试

### 身份认证测试
- JWT认证方式
- API Key认证方式
- 完整配置测试
- 认证方式对比
- 错误处理测试

## 注意事项

1. **认证方式**: 推荐使用JWT认证，安全性更高
2. **错误处理**: 示例包含完整的错误处理演示
3. **数据格式**: 展示了如何解析和使用API返回的数据
4. **最佳实践**: 代码遵循Go语言和SDK的最佳实践

## 常见问题

**Q: 测试时遇到认证失败怎么办？**
A: 检查Token是否正确，是否有访问权限，网络是否正常。

**Q: 为什么推荐使用JWT认证？**
A: JWT认证更安全，支持过期时间，且从2027年开始API Key将受到限制。

**Q: 如何获取LocationID？**
A: 可以通过和风天气的GeoAPI获取，常用城市ID如：北京-101010100，上海-101020100。

## 相关链接

- [SDK文档](../README.md)
- [和风天气开发者平台](https://dev.qweather.com/)
- [身份认证文档](https://dev.qweather.com/docs/resource/auth/)
- [API文档](https://dev.qweather.com/docs/api/) 