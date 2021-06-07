##### 基于Gin + Vue + ANTDV的前后端分离权限管理系统 
[Gin文档](https://gin-gonic.com/zh-cn/docs/)

[ANTDV 官方文档](https://pro.antdv.com/)

###### ✨ 简介
- 基础的多角色权限系统，颗粒度细到每一个操作类型
- RESTful API
- 基于golang的GIN框架
- JWT认证及Casbin的访问控制模型
- 基于gorm的mysql数据库存储

###### 预览
<p align="center">
  <img  src="https://github.com/MarshallTeach/basic-antd/blob/main/assets/images/screen.png">
</p>

###### 🎬 在线体验
> test  /  123456
>
>演示地址：[预览](http://115.29.226.247:8011)
>
>前端代码：[传送门](https://github.com/MarshallTeach/basic-antd-front)

###### 项目启动
- 新建"basic_antd"数据库，将config的sql文件导入数据库，修改config中的数据库用户名跟密码
- go get
- go run main.go start

###### 编译
- windows: go build
- linux: set CGO_ENABLED=0
- set GOOS=linux
- go build

###### 部署
- windows: 双击exe文件运行
- linux: nohup ./basic_antd start -c configs/setting.yml &

###### 想说的话
- 如果有问题或者BUG或者某个函数有更好的算法或者写法，欢迎提出issue||pr，共同进步
- 如果这个项目有帮助到你的地方，欢迎star~

<p align="center">
  <img  src="https://github.com/MarshallTeach/basic-antd/blob/main/assets/images/xmn.gif" width="200px" height="200px">
</p>

###### ToDo
- 使用redis缓存优化代码
- 加入im服务

###### 更新日志
- 20210607-优化项目结构