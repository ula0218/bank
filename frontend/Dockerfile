# 使用官方 Node.js 映像
FROM node:14-alpine

# 設置工作目錄
WORKDIR /app

# 複製 package.json 與 package-lock.json 並安裝依賴
COPY package*.json ./
RUN npm install --silent

# 複製所有源代碼到容器中
COPY . .

# 建立生產環境構建
RUN npm run build

# 暴露端口，用於本地測試
EXPOSE 3001

# 定義啟動命令
CMD ["npm", "start"]
