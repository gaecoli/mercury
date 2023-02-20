

<center font-size='15px'><font size='50'>Mercury</font></center>

1. 基本介绍

   Mercury 是一个基于 React 和 gin 开发的集成工具平台，集成 jwt 鉴权；Todo-Lists；记账；上传；文章记录以及照片墙等功能。

2. 项目功能点

```
- login
	- jwt
- todo-list
- monopoly (记账)
- upload
	- files
	- photo
- articles
	- markdown
	- rich text-editor
- ideas
- error-page
- global
```

3. 技术选型
   - 前端：基于 React 的 antd-design构建页面；
   - 后端：基于 Gin 搭建 restful API；
   - 数据库：MySQL (gorm 实现对数据库的增删改查)
   - 缓存：Redis (记录用户 jwt token，实现单点登录；记录用户访问过的页面信息)
4. 主要功能
   - 用户登录
   - todo-list
   - 记账
   - 引入图表汇总（ECharts）
   - 上传文件（断点续传）和图片
   - 记录文章
   - 灵光乍现的点子 (ideas)
   - 错误页面
   - 国际化

