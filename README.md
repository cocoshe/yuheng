# yuheng

yuheng ☁️ is a convenient and friendly backend service to manage the users' and companies' information.


## Directory Structure

```shell
.
├── Dockerfile                  # Dockerfile打包docker镜像
├── README.md                   # 说明文档
├── accus_img                   # 图片
├── appeal_img                  # 图片
├── application.example.yaml    # 配置文件
├── application.yaml            # 配置文件
├── config                      # 导入配置
│   └── config.go
├── dao                         # 持久化层
│   ├── AccusationDao.go
│   └── userDao.go
├── docs                        # api文档(swagger启动)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── drivers                     # 驱动层
│   └── mysqldb.go
├── go.mod                      # 包管理
├── go.sum
├── handlers                    # 业务层 (接口实现)
│   ├── addAtten.               # 添加关注
│   ├── attenInfo.go            # 查看关注信息
│   ├── company                 # 公司路由组
│   │   ├── appeal.go           # 发起申诉
│   │   └── history.go          # 历史申诉记录
│   ├── delAtten.go             # 取消关注
│   ├── display.go              # 查看用户信息
│   ├── displayRank.go          # 查看排名
│   ├── displaybad.go           # 查看异常公司
│   ├── getFeatures.go          # 获取特征类名
│   ├── gvmt                    # 公司路由组
│   │   ├── changeAccusStatus.go # 修改投诉状态
│   │   ├── changeAplStatus.go   # 修改申诉状态
│   │   ├── changeStatus.go      # 修改公司状态
│   │   ├── changeThreshold.go   # 修改阈值
│   │   ├── getAccusationList.go # 获取投诉列表
│   │   └── getAppealList.go     # 获取申诉列表
│   ├── info.go                  # 查看公司相关排放量等信息
│   ├── login.go                 # 登录
│   ├── pong.go                  # 测试(debug)
│   ├── runModel.go              # 重定向给flask应用接口
│   ├── selfCheck.go             # 重定向给flask应用，进行自检
│   └── visitor                  # 访客路由组
│       ├── history.go           # 查看历史投诉信息
│       └── upload.go            # 上传投诉信息
├── main.go                      # 主程序入口
├── middleware                   # 中间件
│   ├── CompanyMiddleware.go     # 公司鉴权控制
│   ├── GvmtMiddleware.go        # 管理人员鉴权控制
│   ├── VisitorMiddleware.go     # 访客鉴权控制
│   └── const.go                 # 等级/状态常量
├── models                       # 模型层
│   ├── accusation.go            # 投诉信息
│   ├── allData.go               # 数据库模型
│   ├── company.go               # 公司信息
│   ├── templates.go             # 消息绑定模板
│   └── user.go                  # 用户信息
├── routers                      # 路由
│   └── router.go                # 路由创建
├── server                       # 服务
│   └── server.go                # 服务启动
├── templates                    # 模板测试
│   └── index.html               # 模板测试
├── test.xlsx                    # 自检测试数据
└── utils                        # 工具
    └── jwt.go                   # jwt鉴权

```


## Quick Start
1. Adjust the `application.yaml` file according to your needs (see application.yaml.example for more details).

2. Install the dependencies:
```shell
go mod tidy
```

3. Run the application:
```shell
go run main.go
```

## Docker

```dockerfile
docker pull cocoshe/yuheng:api
```

if you want to make your own docker image:
```dockerfile
docker build -t cocoshe/yuheng:tagname .
docker push cocoshe/yuheng:tagname
```

## API doc
http://localhost:port/swagger/index.html
http://psyqlk.space/api/v1/swagger/index.html