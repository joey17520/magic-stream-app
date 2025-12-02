# MagicStream - 云原生电影流媒体平台

MagicStream 是一个现代化的云原生电影流媒体平台，采用微服务架构和云原生技术栈，提供高性能、可扩展的电影浏览、播放和智能推荐功能。

## 🎯 项目亮点（云原生工程师视角）

### 🚀 云原生架构优势

- **容器化部署**：完整的 Docker 容器化支持，一键部署
- **微服务架构**：前后端分离，API 驱动设计
- **可观测性**：集成 Prometheus + Grafana 监控体系
- **弹性伸缩**：基于容器编排的自动扩缩容能力
- **配置管理**：环境变量驱动的配置系统

### 📊 技术深度亮点

- **高性能 Go 后端**：Gin 框架 + MongoDB，支持高并发请求
- **现代化 React 前端**：Vite 构建，响应式设计
- **AI 增强推荐**：集成 LangChain + DeepSeek API 的智能推荐系统
- **完整监控体系**：业务指标 + 系统指标全方位监控
- **生产级安全**：JWT 认证 + CORS + 结构化日志

## 🏗️ 技术架构

### 前端技术栈

- **React 18** + **Vite** - 现代化前端开发体验
- **Axios** - 封装 HTTP 客户端，支持自动重试和拦截器
- **Context API** - 轻量级状态管理
- **CSS Modules** - 组件化样式管理
- **ESLint + Prettier** - 代码质量和格式化

### 后端技术栈

- **Golang 1.25** - 高性能并发处理
- **Gin Web Framework** - 轻量级 HTTP 框架
- **MongoDB 6+** - 文档型数据库，灵活的数据模型
- **JWT + Refresh Token** - 双重认证机制
- **Prometheus + Grafana** - 监控和可视化
- **Zap Logger** - 高性能结构化日志
- **LangChain** - AI 集成框架

### 云原生基础设施

- **Docker** - 容器化部署
- **Docker Compose** - 多服务编排
- **Prometheus** - 指标收集和告警
- **Grafana** - 数据可视化和仪表盘
- **Nginx** - 反向代理和负载均衡

## 📁 项目结构

```
MagicStream/
├── client/                          # 前端React应用
│   ├── src/
│   │   ├── components/              # React组件（模块化组织）
│   │   ├── context/                 # React Context状态管理
│   │   ├── hooks/                   # 自定义Hooks
│   │   ├── api/                     # API配置和封装
│   │   └── assets/                  # 静态资源
│   ├── Dockerfile                   # 前端容器化配置
│   ├── nginx.conf                   # Nginx配置
│   └── package.json
├── server/                          # 后端Golang微服务
│   ├── config/                      # 配置管理（环境变量驱动）
│   ├── controllers/                 # 业务控制器
│   ├── models/                      # 数据模型
│   ├── routes/                      # API路由（保护/非保护路由分离）
│   ├── middlewares/                 # 中间件（认证、监控、CORS等）
│   ├── database/                    # 数据库连接和操作
│   ├── utils/                       # 工具函数（日志、令牌等）
│   ├── dashboard/                   # 监控仪表盘配置
│   │   ├── grafana/                 # Grafana配置
│   │   │   ├── dashboards/          # 业务监控仪表盘
│   │   │   └── provisioning/        # 自动配置
│   │   └── prometheus/              # Prometheus配置
│   ├── Dockerfile                   # 后端容器化配置
│   └── go.mod
├── docker-compose.yml               # 多服务编排配置
├── .env.example                     # 环境变量模板
├── CONFIGURATION_README.md          # 配置详细说明
└── README.MD
```

## 🚀 快速开始

### 环境要求

- Docker 20.10+ 和 Docker Compose 2.0+
- 或 Node.js 18+ + Go 1.25+ + MongoDB 6+（本地开发）

### 一键部署（推荐）

```bash
# 克隆项目
git clone https://github.com/joey17520/magic-stream-app.git
cd magic-stream-app

# 复制环境变量配置
cp .env.example .env
# 编辑.env文件配置必要的环境变量

# 启动所有服务
docker-compose up -d
```

### 服务访问地址

- **前端应用**：http://localhost:80
- **后端 API**：http://localhost:8088
- **API 文档**：http://localhost:8088/swagger/index.html
- **Grafana 监控**：http://localhost:3000 (admin/admin)
- **Prometheus**：http://localhost:9090

### 分步部署（开发环境）

#### 1. 后端服务

```bash
cd server
go mod download
go run main.go
```

#### 2. 前端应用

```bash
cd client
npm install
npm run dev
```

## 🔧 核心功能模块

### 🎬 电影管理

- 电影 CRUD 操作（管理员）
- 电影分类和标签
- 电影详情展示
- 电影搜索和筛选

### 📺 播放功能

- YouTube 视频集成
- 播放器控制（播放/暂停/音量）
- 播放进度保存
- 观看历史记录

### ⭐ 评分评论系统

- 用户评分（1-5 星）
- 文字评论发布
- 评论情感分析（AI 驱动）
- 评论回复和互动

### 🎯 智能推荐

- 基于观看历史的协同过滤
- AI 驱动的个性化推荐
- 实时推荐更新
- 多维度推荐策略

### 🔐 用户认证

- JWT + Refresh Token 双重认证
- 密码加密存储（bcrypt）
- 会话管理
- 权限控制（用户/管理员）

## 📊 监控与可观测性

### 业务指标监控

- HTTP 请求速率和延迟
- 数据库操作性能
- 用户活跃度统计
- 电影观看趋势

### 系统指标监控

- 服务健康状态
- 资源使用率（CPU/内存）
- 网络 I/O 统计
- 错误率和异常检测

### Grafana 仪表盘

项目包含预配置的监控仪表盘：

- **MagicStream Overview**：核心业务指标概览
- **API 性能监控**：HTTP 请求详细分析
- **数据库性能**：MongoDB 操作监控
- **用户行为分析**：用户活跃度和转化

## ⚙️ 配置管理

### 环境变量配置

项目采用环境变量驱动的配置管理，支持以下配置项：

```bash
# 服务器配置
PORT=8088
GIN_MODE=release
LOG_LEVEL=info

# 数据库配置
MONGODB_URI=mongodb://mongodb:27017
DATABASE_NAME=magicstream

# 安全配置
SECRET_KEY=your-secret-key
SECRET_REFRESH_KEY=your-refresh-secret-key

# CORS配置
ALLOWED_ORIGINS=http://localhost:80,http://localhost:5173

# AI服务配置
DEEPSEEK_API_KEY=your-deepseek-api-key

# 业务配置
RECOMMENDED_MOVIE_LIMIT=5
```

### 配置验证

- 必需配置项验证
- 配置合理性检查
- 敏感信息保护
- 配置变更日志

## 🐳 容器化部署

### Docker 镜像构建

```bash
# 构建后端镜像
docker build -t magicstream-server ./server

# 构建前端镜像
docker build -t magicstream-client ./client
```

### Docker Compose 编排

```yaml
version: "3.8"
services:
  mongodb:
    image: mongo:6
    volumes:
      - mongodb_data:/data/db

  api:
    build: ./server
    depends_on:
      - mongodb
    environment:
      - MONGODB_URI=mongodb://mongodb:27017

  client:
    build: ./client
    depends_on:
      - api

  prometheus:
    image: prom/prometheus
    volumes:
      - ./server/dashboard/prometheus:/etc/prometheus

  grafana:
    image: grafana/grafana
    depends_on:
      - prometheus
```

## 🔍 API 文档

### RESTful API 设计

- **标准化响应格式**：统一成功/错误响应结构
- **版本控制**：API 版本管理
- **分页和过滤**：标准化的查询参数
- **速率限制**：API 访问频率控制

### 主要 API 端点

```
GET    /api/v1/health          # 健康检查
GET    /api/v1/metrics         # Prometheus指标
POST   /api/v1/register        # 用户注册
POST   /api/v1/login           # 用户登录
GET    /api/v1/movies          # 电影列表
GET    /api/v1/movies/:id      # 电影详情
POST   /api/v1/reviews         # 发布评论
GET    /api/v1/recommended     # 推荐电影
```

## 🧪 测试策略

### 测试金字塔

- **单元测试**：业务逻辑和工具函数
- **集成测试**：API 端点和数据库交互
- **E2E 测试**：完整用户流程测试

### 质量保证

- 代码覆盖率报告
- 静态代码分析
- 安全漏洞扫描
- 性能基准测试

## 📈 性能优化

### 后端优化

- 数据库查询优化（索引、聚合管道）
- 连接池管理
- 缓存策略（Redis 集成准备）
- 异步处理（耗时操作）

### 前端优化

- 代码分割和懒加载
- 图片优化和 CDN
- 服务端渲染（SSR）准备
- PWA 支持

## 🔄 CI/CD 流水线

### GitHub Actions 工作流

- 自动化测试和构建
- Docker 镜像构建和推送
- 安全扫描和代码质量检查
- 自动化部署到云平台

### 部署策略

- 蓝绿部署
- 金丝雀发布
- 滚动更新
- 健康检查和自愈

## 🛡️ 安全特性

### 认证和授权

- JWT 令牌验证
- Refresh Token 轮换
- 角色基于访问控制（RBAC）
- API 密钥管理

### 数据安全

- 密码加密存储（bcrypt）
- 敏感数据脱敏
- SQL/NoSQL 注入防护
- XSS 和 CSRF 防护

### 网络安全

- HTTPS 强制
- CORS 策略配置
- 速率限制
- DDoS 防护

## 🤝 贡献指南

我们欢迎社区贡献！请遵循以下流程：

1. **Fork 项目**并创建功能分支
2. **编写测试**：确保新功能有测试覆盖
3. **遵循代码规范**：使用项目约定的代码风格
4. **提交 PR**：包含清晰的描述和测试结果
5. **代码审查**：通过审查后合并

### 开发规范

- 提交信息遵循 Conventional Commits
- 保持代码简洁和可维护
- 编写清晰的文档和注释
- 及时更新 CHANGELOG

## 📄 许可证

本项目基于 MIT 许可证开源 - 查看[LICENSE](LICENSE)文件了解详情。

## 📞 联系方式

- **项目维护者**：JoeyWong
- **邮箱**：joey17520@163.com
- **GitHub**：[joey17520](https://github.com/joey17520)

## 🙏 致谢

感谢所有为这个项目做出贡献的开发者！特别感谢以下开源项目：

- Gin Web Framework
- React 和 Vite 团队
- MongoDB 社区
- Prometheus 和 Grafana 项目

---

**MagicStream** - 构建下一代云原生流媒体平台！ 🚀

_最后更新：2025 年 12 月_
