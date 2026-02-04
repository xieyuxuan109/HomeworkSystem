# 红岩网校作业管理系统（Redrock Assessment System）

## 项目简介

本项目是一个面向红岩网校内部使用的作业管理系统，支持「老登」（管理员/讲师）发布与批改作业、「小登」（学员）提交作业并查看评语的完整流程。系统按部门划分权限，涵盖后端、前端、SRE、产品、视觉设计、Android、iOS七个方向，并集成了 AI 初评功能以辅助教学评估。

---

## 已实现功能清单

### 基础功能

- **用户模块**
  - 用户注册（含用户名、密码、昵称、部门）
  - 用户登录（返回 Access Token + Refresh Token）
  - 获取当前用户信息（含 `department` 与 `department_label`）
  - 注销账号（软删除）
  - 密码使用 bcrypt 加盐哈希存储
  - JWT 双 Token 认证机制（含刷新逻辑）

- **作业模块**
  - 老登可发布作业（设置标题、描述、部门、截止时间、是否允许补交）
  - 按部门筛选 + 分页查询作业列表
  - 查看作业详情（含发布者信息、提交人数、当前用户提交状态）
  - 同部门老登可修改或删除作业（已处理并发控制）

- **提交模块**
  - 小登提交作业（自动判断是否迟交）
  - 小登查看自己的所有提交及评语
  - 老登查看本部门所有学员的提交
  - 老登批改作业（填写分数、评语、标记是否优秀）
  - 所有用户可查看优秀作业列表（支持按部门筛选）

- **其他要求**
  - 统一响应格式：`{ code, message, data }`
  - 部门枚举值与中文标签同时返回（如 `"backend"` → `"后端"`）
  - 规范的 Git 提交记录（Conventional Commits 风格）
  - 使用 Gin 框架开发
  - 完整 API 文档（见链接）

### 进阶功能

- **AI 作业初评（AI Review）**
  - 老登可通过 `/submission/:id/aiReview` 接口触发 AI 对作业内容的自动分析
  - 调用大模型 API对代码质量、结构、规范性进行初步评价
  - 返回 AI 生成的评语建议与推荐分数，供老登参考
  - 支持文本/链接形式的作业内容分析
  - 此外，还在本地通过docker ollama部署了本地ai，也可以进行ai评价
---

## 技术栈说明

- **后端语言**：Go 1.25+
- **Web 框架**：Gin
- **数据库**：MySQL 8.0
- **ORM**：GORM
- **认证机制**：JWT（Access Token + Refresh Token）
- **密码安全**：bcrypt 加盐哈希
- **AI 集成**：通过 HTTP 调用第三方大模型 API（如 DeepSeek、Qwen等）
- **部署方式**：Docker 容器化（含 `docker-compose.yaml`）或本地部署
---

## 项目结构说明
```
homework-system/
├── api/ # API 文档与接口定义
├── cmd/
│ └── main.go # 程序入口
├── configs/ # 配置文件（数据库、JWT、AI等）
├── dao/ # 数据访问层（User, Homework, Submission）
├── docker/ #docker部署相关文件
├── handler/ # HTTP 请求处理器
├── middleware/ # 中间件（JWT 认证、权限校验）
├── models/ # GORM 模型定义
├── pkg/
│ ├── jwt/ # JWT 工具
│ ├── response/ # 统一响应封装
│ └── password.go/#密码加密相关操作
├── router/ # 路由注册
├── service/ # 业务逻辑层
├── go.mod
├── go.sum
└── README.md
```
---
## 本地运行指南

1. **下载**
  ```
   git clone https://github.com/xieyuxuan109/HomeworkSystem.git
  ```
2. **运行**
  ```
  cd HomeworkSystem/docker
  docker load -i homeworksystem.tar//自己制作的镜像,但由于镜像过大，需要的时候可以私我要镜像，也可以根据dockerfile自行构建
  cp .env.example .env
  docker-compose up -d
  docker-compose ps
  ```
  ```
  运行成功图片(刚开始需要等mysql容器初始化完成第二个容器才会启动)
  ```
  ![运行成功图片](/img/05.png)
3.**测试**
  ```
  导入postman测试脚本和环境变量
  测试ai功能时候需修改.env文件里面的DASHSCOPE_API_KEY，仅支持千问
  ```
  ```
  调试成功图片
  ```
  ![调试成功图片](/img/01.png)
  ```
  调试后docker logs日志
  ```
  ![调试后docker logs日志](/img/02.png)
  


## API文档，由postman直接导出，并且由自动化测试代码，所有由两个文件，一个是测试API，一个是环境变量
["Postman API文件链接"](https://github.com/xieyuxuan109/HomeworkSystem/tree/main/api "Postman API文件")
  ```
  导入两个文档后效果
  ```
  ![导入两个文档后效果](/img/03.png)
  ![导入两个文档后效果](/img/04.png)
