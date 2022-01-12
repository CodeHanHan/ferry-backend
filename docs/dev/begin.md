# ferry-backend

## 1. 目录
```
.
├── apis  // api接口
│   └── ping  // ping组接口
├── db  // 数据库相关
│   ├── ddl  // 创建数据库
│   ├── logger  // gorm Interface
│   ├── query  // 数据库查询
│   └── schema  // 数据库表定义
├── deploy  // 部署、配置相关
│   ├── config  // 项目配置
│   └── docker-compose  // docker-compose部署
├── docs
│   └── dev // 开发文档
├── middleware  // 中间件
├── models  // orm Model
│   └── ping
│   └── ping
├── pkg  // 相关包 
│   ├── app  // 请求回答相关
│   ├── config  // 加载配置
│   ├── captcha // 生成验证码
│   ├── err  
│   ├── form  // 请求表单
│   ├── logger  // 日志
│   ├── mycasbin // 权限认证
│   ├── pi  // 全局
│   ├── token // jwt token身份验证
│   ├── validator // 表单字段验证
│   └── xmysql // 数据库连接
├── routers  // 路由
│   ├── ping  // ping路由组
│   └── user  // user路由组
└── utils  // 工具包
    ├── idutil // uuid
    ├── password // 密码加密&验证
    └── rand // 随机
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

## 4. 访问接口
本地启动的项目将监听`10000`端口，docker-compose启动的项目将监听`10001`端口。

这里以`/ping/create`路由为例简单介绍接口的访问过程。
![20211122144749](https://raw.githubusercontent.com/lich-Img/blogImg/master/img/20211122144749.png)

我的开发环境为WSL2（Ubuntu18.04）,在这里我将使用`wsl`这个dns访问wsl主机。
如果我要访问`/ping/create`这个接口，那么我访问的地址是：
```
POST http://wsl:10001/ping/create
```
访问这个接口时，表单应包含`message`，否则无法验证通过。一切顺利的话，程序将返回你发送的message值加上`, too`后缀。如`hello` -> `hello, too`。

## 5. 调试（vscode）
打开vscode，点击调试选项卡    
![20211122145329](https://raw.githubusercontent.com/lich-Img/blogImg/master/img/20211122145329.png)

点击`create a launch.json file`创建`launch.json`文件。    
![20211122145727](https://raw.githubusercontent.com/lich-Img/blogImg/master/img/20211122145727.png)

创建一个`attach to local process`配置  
![20211122145827](https://raw.githubusercontent.com/lich-Img/blogImg/master/img/20211122145827.png)

创建完成后配置如下：
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Attach to Process",
            "type": "go",
            "request": "attach",
            "mode": "local",
            "processId": 0
        }
    ]
}
```

确保项目此时已经**本地启动**

设置断点：  
![20211122150044](https://raw.githubusercontent.com/lich-Img/blogImg/master/img/20211122150044.png)

开始调试  
![20211122150008](https://raw.githubusercontent.com/lich-Img/blogImg/master/img/20211122150008.png)

附加进程  
![20211122150138](https://raw.githubusercontent.com/lich-Img/blogImg/master/img/20211122150138.png)

访问`http://wsl:10000/ping/create`，命中断点  
![20211122150325](https://raw.githubusercontent.com/lich-Img/blogImg/master/img/20211122150325.png)
