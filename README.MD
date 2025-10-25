# MagicStream - 电影流媒体平台

MagicStream 是一个现代化的电影流媒体平台，采用前后端分离架构，提供用户友好的电影浏览、播放和评论功能。

## 🎯 项目概述

MagicStream 是一个全栈电影流媒体应用，结合了现代化的前端技术和强大的后端服务，为用户提供流畅的电影观看体验。

### 主要特性

- 🎬 **电影浏览** - 浏览丰富的电影库
- 📺 **在线播放** - 集成 YouTube 视频播放
- ⭐ **评分系统** - 用户可以对电影进行评分
- 💬 **评论功能** - 分享观影感受
- 🔐 **用户认证** - 安全的用户注册和登录系统
- 🎯 **推荐系统** - 个性化电影推荐
- 📱 **响应式设计** - 适配各种设备屏幕

## 🏗️ 技术架构

### 前端技术栈

- **React 18** - 现代化前端框架
- **Vite** - 快速构建工具
- **Axios** - HTTP 客户端
- **CSS3** - 样式设计
- **React Context** - 状态管理

### 后端技术栈

- **Golang** - 高性能后端语言
- **Gin** - Web 框架
- **MongoDB** - NoSQL 数据库
- **JWT** - 身份验证
- **CORS** - 跨域资源共享

## 📁 项目结构

```
MagicStream/
├── client/                 # 前端React应用
│   ├── src/
│   │   ├── components/     # React组件
│   │   ├── context/        # 状态管理
│   │   ├── hooks/          # 自定义Hooks
│   │   └── api/            # API配置
│   └── package.json
├── server/                 # 后端Golang服务
│   ├── controllers/        # 控制器层
│   ├── models/             # 数据模型
│   ├── routes/             # 路由配置
│   ├── middlewares/        # 中间件
│   ├── database/           # 数据库连接
│   └── utils/              # 工具函数
└── README.MD
```

## 🚀 快速开始

### 环境要求

- Node.js 16+
- Go 1.19+
- MongoDB 5+

### 安装与运行

#### 后端服务

1. 进入 server 目录：

```bash
cd server
```

2. 安装依赖：

```bash
go mod tidy
```

3. 配置环境变量：

```bash
cp .env.example .env
# 编辑.env文件配置数据库连接等
```

4. 启动服务：

```bash
go run main.go
```

#### 前端应用

1. 进入 client 目录：

```bash
cd client
```

2. 安装依赖：

```bash
npm install
```

3. 启动开发服务器：

```bash
npm run dev
```

### 访问应用

- 前端应用：http://localhost:5173
- 后端 API：http://localhost:8088

## 🔧 功能模块

### 用户管理

- 用户注册和登录
- JWT 身份验证
- 用户会话管理

### 电影管理

- 电影信息展示
- 电影搜索和筛选
- 电影详情查看

### 播放功能

- YouTube 视频集成
- 播放器控制
- 播放历史记录

### 评论系统

- 电影评论发布
- 评论查看和回复
- 评分功能

## 📊 API 文档

### 公开路由

- `GET /movies` - 获取电影列表
- `GET /movies/:id` - 获取电影详情
- `POST /register` - 用户注册
- `POST /login` - 用户登录

### 受保护路由

- `POST /movies` - 添加电影（管理员）
- `PUT /movies/:id` - 更新电影信息
- `POST /reviews` - 发布评论
- `GET /recommended` - 获取推荐电影

## 🛠️ 开发指南

### 代码规范

- 前端使用 ESLint 进行代码检查
- 后端遵循 Go 语言最佳实践
- 统一的代码格式化

### 数据库设计

使用 MongoDB 存储以下主要集合：

- `users` - 用户信息
- `movies` - 电影数据
- `reviews` - 评论数据
- `genres` - 电影分类

### 部署说明

项目支持多种部署方式：

- Docker 容器化部署
- 云平台部署（AWS、GCP、Azure）
- 传统服务器部署

## 🤝 贡献指南

我们欢迎社区贡献！请遵循以下步骤：

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

本项目基于 MIT 许可证开源 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 📞 联系方式

- 项目维护者：JoeyWong
- 邮箱：joey17520@163.com

## 🙏 致谢

感谢所有为这个项目做出贡献的开发者！

---

**MagicStream** - 让电影观看体验更美好！ 🎉
