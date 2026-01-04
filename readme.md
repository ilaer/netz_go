# Netz

Netz是一个用Go语言开发的局域网服务发现工具，用于自动发现和识别局域网内的主机、数据库等服务。

## 功能特点

- 自动发现局域网内的主机
- 扫描开放端口
- 识别服务类型和版本
- 检测数据库服务
- 生成扫描报告

## 系统要求

- Go 1.21或更高版本
- 支持的操作系统：Linux, macOS, Windows
- 需要管理员/root权限（用于网络扫描）

## 快速开始

### 安装

```bash
# 克隆仓库
git clone https://github.com/yourusername/netz.git
cd netz

# 安装依赖
go mod download

# 编译
go build -o netz cmd/netz/main.go
```

### 使用方法

1. 基本扫描：
```bash
./netz scan --network 192.168.1.0/24
```

2. 快速扫描：
```bash
./netz scan --quick
```

3. 详细扫描：
```bash
./netz scan --network 192.168.1.0/24 --ports 1-1000 --detect-db
```

4. 导出报告：
```bash
./netz scan --network 192.168.1.0/24 --output report.json
```

## 配置说明

配置文件位于 `config/config.yaml`，主要配置项包括：

```yaml
scanner:
  timeout: 5s
  threads: 10
  ports: [21,22,23,80,443,3306,5432,6379,27017]

database:
  mysql:
    enabled: true
    timeout: 3s
  postgres:
    enabled: true
    timeout: 3s
  redis:
    enabled: true
    timeout: 2s
  mongodb:
    enabled: true
    timeout: 3s

output:
  format: json
  path: ./reports
```

## 开发文档

详细的开发文档请参考 [docs/development.md](docs/development.md)

## 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 免责声明

本工具仅用于合法的网络管理和安全测试目的。使用本工具进行未经授权的网络扫描可能违反相关法律法规。使用者需自行承担使用本工具的风险和责任。 