# 多阶段构建Dockerfile for React前端应用
# 第一阶段：构建阶段
FROM node:22-alpine AS builder

# 设置Alpine国内镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 设置npm国内镜像源
RUN npm config set registry https://registry.npmmirror.com

# 设置工作目录
WORKDIR /app

# 复制package.json和package-lock.json
COPY package*.json ./

# 安装依赖
RUN npm ci --only=production

# 复制源代码
COPY . .

# 构建应用
RUN npm run build

# 第二阶段：运行阶段
FROM nginx:1.28-alpine

# 设置工作目录
WORKDIR /usr/share/nginx/html

# 从构建阶段复制构建产物
COPY --from=builder /app/dist .

# 复制nginx配置文件
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 80

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:80 || exit 1

# 运行nginx
CMD ["nginx", "-g", "daemon off;"]
