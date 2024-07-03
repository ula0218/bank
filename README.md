# Bank

## Skills
* Backend: GoLang, Gin
* Frontend: React
* Deployment: EC2
* Version Control: Git
* Containerization: Docker
  
## Host（已停止EC2 instance )
```
http://52.194.190.91/
```
## Setup Instructions

### csv檔案
政府資料公開平台下載最新『金融機構基本資料查詢』
```
https://data.gov.tw/dataset/6041
```
## Deploy
### 1. 反向代理設定
* 有需要再設定
* 在 instance 的終端執行：
```
apt update
apt install nginx
systemctl start nginx
systemctl enable nginx
systemctl status nginx

vim /etc/nginx/nginx.conf
systemctl restart nginx  #設定好後重啟
```

* `nginx.conf`:
```
server {
    listen 80;
    server_name 52.194.190.91;  #自行更換(EC2)

    location / {
        proxy_pass http://localhost:3456;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /api/ {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```
### 2.vim編輯frontend/src/App.js
* 在 instance 的終端執行：

修改axios.get(`http://localhost:8080/`)
```
axios.get(`http://52.194.190.91:8080/`)
```
修改axios.get(`http://localhost:8080/${bankCode}/branches`)
```
axios.get(`http://52.194.190.91:8080/${bankCode}/branches`)
```

## Docker
```
docker-compose -f docker-compose.yml down # 停止當前運行容器
docker-compose -f docker-compose.yml build
docker-compose -f docker-compose.yml up -d
```

