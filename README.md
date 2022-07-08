# BaseWebProject

![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/navendu-pottekkat/awesome-readme?include_prereleases)
![GitHub issues](https://img.shields.io/github/issues-raw/navendu-pottekkat/awesome-readme)
![GitHub pull requests](https://img.shields.io/github/issues-pr/navendu-pottekkat/awesome-readme)

golang整合Gin Gorm Mysql Redis Jwt的基础web环境

#  项目结构

├─cache
├─config
├─controller
├─env
├─logger
├─middleware
├─model
├─pkg
│  ├─e
│  └─util
└─routes

cache

+ 存放redis相关文件

config

+  读取配置文件config.ini中的内容
+ redis相关配置初始化
+ mysql相关配置初始化

controller

+ 控制层接口

env

+ 全局环境变量

logger

+ 基于logrus的日志封装

middleware

+ cors
+ jwt

model

+ mysql数据库相关配置
+ 实体类

pkg

+ 工具包

routes

+ gin路由

