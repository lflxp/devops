# boltDb
boltDb test

https://github.com/coreos/bbolt/
https://npf.io/2014/07/intro-to-boltdb-painless-performant-persistence/
https://www.progville.com/go/bolt-embedded-db-golang/

# devops
运维自动化

# 开发计划

## 第一期

1. 实现功能为主 架构为辅
2. cmdb

* 2.1 主机管理
* 2.2 机房管理
* 2.3 机柜管理
* 2.4 属组管理
* 2.5 Ansible管理
* 2.6 Shell管理 

3. bolt底层存储

* 3.1 数据库所有表查询
* 3.2 数据库范围查询
* 3.3 数据库前缀查询
* 3.4 数据库新增
* 3.5 数据库修改
* 3.6 数据库删除

4. webterminal

* 4.1 webTerminal
* https://github.com/gorilla/websocket/tree/master/examples/echo
* https://github.com/urfave/cli
* https://github.com/elazarl/go-bindata-assetfs
* https://github.com/kr/pty
* https://github.com/yudai/gotty
* https://github.com/joewalnes/websocketd

5. tree

6.用户管理

* 6.1 用户管理
* 6.2 角色管理
* 6.3 权限管理

## 第二期

1. 服务化和架构优化
2. docker化
3. 一键部署
4. 微服务化以及k8s运维

# INSTALL

* git clone http://github.com/lflxp/devops.git

* cd devops/ui

* npm install

* npm run dev

* npm run build

* cd ..

* go build -o devops main.go

* chmod +x devops && ./devops
