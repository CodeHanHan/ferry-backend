# ferry-backend

## 1. 目录
```
.
├── apis
│   ├── dept
│   ├── ping
│   ├── post
│   ├── role
│   └── user
├── db
│   ├── db.go
│   ├── ddl
│   ├── Dockerfile
│   ├── errors.go
│   ├── logger
│   ├── Makefile
│   ├── query
│   └── schema
├── debug-main
├── deploy
│   ├── config
│   └── docker-compose
├── Dockerfile
├── docs
│   ├── dev
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── Makefile
├── middleware
│   ├── auth.go
│   ├── check_role.go
│   ├── header.go
│   ├── init.go
│   ├── log.go
│   ├── requestId.go
│   └── sender.go
├── models
│   ├── dept
│   ├── ping
│   ├── post
│   ├── role
│   └── user
├── pkg
│   ├── app
│   ├── captcha
│   ├── config
│   ├── constants
│   ├── form
│   ├── logger
│   ├── mycasbin
│   ├── pi
│   ├── sender
│   ├── token
│   ├── validator
│   └── xmysql
├── README.md
├── routers
│   ├── dept
│   ├── init.go
│   ├── ping
│   ├── post
│   ├── role
│   └── user
├── static
│   ├── Dockerfile
│   ├── Makefile
│   └── nginx-conf
├── utils
│   ├── fileutil
│   ├── idutil
│   ├── imageutil
│   ├── password
│   ├── rand
│   └── stringutil
└── wercker.yml

```

## 2. 启动流程（本地启动）
1. 拉取最新代码
    ```bash
    git clone git@github.com:CodeHanHan/ferry-backend.git
    ```

2. 配置数据库
   ```sh
   make build-db # 创建数据库image

   make db-up # 启动数据库container

   make migrate-up # 数据库迁移
   ```

3. 启动项目
    ```sh
    make dev-up # 启动项目
    ```

## 3. 启动流程（docker-compose）
1. 拉取最新代码
    ```bash
    git clone git@github.com:CodeHanHan/ferry-backend.git
    ```

2. build镜像
    ```sh
    make build-images
    ```

3. 启动
    ```sh
    make compose-up
    ```

## 4. Contributors
<a href="https://github.com/CodeHanHan/ferry-backend/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=CodeHanHan/ferry-backend" />
</a>
