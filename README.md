# mywaf-admin
[mywaf](https://github.com/medasz/mywaf)
的web管理界面
# 安装
```shell script
1. 安装mariadb数据库
yum install mariadb-server mariadb -y
2. 创建数据库mywaf
create database mywaf
3. 创建用户并赋权给指定数据库
grant all on mywaf.* to 'admin'@'localhost' identified by 'password';
4. 下载最新的编译版本,第一个版本web服务只适用于linux系统
5. 配置app.ini文件
6. 运行
7. 默认用户名/密码:admin/password
```

# 配置
```shell script
根据实际情况修改app.ini中的配置信息
```
# changed
1. 添加web界面配置
2. 提供配置接口
3. 提供规则接口