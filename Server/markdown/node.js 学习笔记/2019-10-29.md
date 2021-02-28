# 2019-10-29

2018-08-08 00:32:56

MongoDB 简介

MongDB 是一款跨平台，面向文档的数据库，可以实现高性能，高可用性，并且能够轻松扩展，是一个基于分布式文件存储的开源数据库系统。在高负载的情况下，添加更多的节点，可以保证服务器性能。

MongoDB 也是一个介于关系数据库和非关系数据库之间的产品，是非关系数据库当中功能最丰富，最像关系数据库的。不采用关系模型主要是为了获得更好的拓展性，MongoDB 不再有“行”(row)的概念，其运行方式主要基于两个概念：集合(collection)与文档(document)。

MongoDB 的特点

MongoDB 的特点包括面向集合存储、模式自由、丰富的查询语句和多级索引、复制集机制、易于水平拓展、可插入存储引擎、跨平台多语言支持等。

MongoDB 安装简单，提供了面向文档存储功能，操作起来比较容易上手。

MongoDB 提供了复制、高可用性和自动分片功能。如果负载增加（需要更多的存储空间和更强的处理能力），它可以分布在计算机网络中的其他节点上，这就是所谓的分片。

Mongo 支持丰富的查询表达式。查询指令使用 JSON 形式的标记，可轻易查询文档中内嵌的对象及数组。

MongoDB 支持各种编程语言：Ruby、Python、Java、C++、PHP、C# 等各种语言。

MongoDB 适用领域

MongoDB 可以为 Web 应用提供可拓展的高性能数据存储解决方案。MongoDB 主要适用领域有网站数据、分布式场景、数据缓存和 JSON 文档格式存储。适合大数据量、高并发、弱事务的互联网应用，其内置的水平拓展机制提供了从几百万到十亿级别的数据处理能力，可以很好地满足 Web2.0 和移动互联网应用数据存储的要求。

MongoDB 4.0 的安装

MongoDB 提供了 Linux 平台上的安装包，可以从官方网站 [http://www.mongodb.org/downloads](http://www.mongodb.org/downloads) 下载。本次我们选择使用最新版 MongoDB4.0 来安装并实验。

下载 MongoDB4.0 软件包

wget [https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-4.0.0.tgz](https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-4.0.0.tgz)

tar zxvf mongodb-linux-x86_64-4.0.0.tgz -C /opt

mv /opt/mongodb-linux-x86_64-4.0.0/ /usr/local/mongodb

创建MongoDB 的数据存储目录，日志存储目录，配置文件

# mkdir /data/mongodb1

# mkdir /data/logs/mongodb

# touch /data/logs/mongodb/mongodb1.log

# chmod -R 777 /data/logs/mongodb/mongodb1.log

# vim /usr/local/mongodb/mongodb1.conf  //添加下面的行

dbpath=/data/mongodb1    #数据存储目录

logpath=/data/logs/mongodb/mongodb1.log    #日志文件

port=27017    #默认服务器端口

logappend=true    #使用追加方式写日志

fork=true    #后台运行

maxConns=5000  #最大同时连接数，默认2000

storageEngine=mmapv1    #指定存储引擎为内存映射文件

设置内核参数

echo 0 > /proc/sys/vm/zone_reclaim_mode

sysctl -w vm.zone_reclaim_mode=0    #永久设置

echo never > /sys/kernel/mm/transparent_hugepage/enabled

echo never > /sys/kernel/mm/transparent_hugepage/defrag

设置系统环境变量，方便使用

echo 'export MONGODB_HOME=/usr/local/mongodb' >> /etc/profile

echo 'export PATH=![](https://g.yuque.com/gr/latex?PATH%3A#card=math&code=PATH%3A)MONGODB_HOME/bin' >> /etc/profile

source /etc/profile

启动 mongodb 服务进程，查看端口（默认27017）

mongod --config /usr/local/mongodb/mongodb1.conf  #开启MongoDB

mongod --config /usrlocal/mongodb/mongodb1.conf --shutdown  #停止MongoDB

netstat -ntap | grep mongod

mongo --port 27017

# 进入mongo数据库，若是不指定端口默认进27017端口

创建多实例

在单台服务器资源充分的情况下，可以使用多实例，以便充分使用服务器资源（只需要修改数据存储目录、日志文件及端口号并创建相应目录即可）

cd /usr/local/mongodb/

cp mongodb1.conf mongodb2.conf

vim mongodb2.conf   #修改如下

# dbpath=/data/mongodb2    #数据存储目录

# logpath=/data/logs/mongodb/mongodb2.log    #日志文件

# port=27018    #默认服务器端口

# logappend=true    #使用追加方式写日志

# fork=true    #后台运行

# maxConns=5000  #最大同时连接数，默认2000

# storageEngine=mmapv1    #指定存储引擎为内存映射文件

mkdir /data/mongodb2

touch /data/logs/mongodb/mongodb2.log

chmod -R 777 /data/logs/mongodb/mongodb2.log

mongod --config /usr/local/mongodb/mongodb2.conf  #开启多实例

MongoDB 逻辑存储结构

MongoDB 的逻辑结构主要由文档（document）、集合（collection）和数据库（database）三部分组成。其中文档是 MongoDB 的核心概念，它是 MongoDB 逻辑存储的最小单元，相当于关系型数据库中的一行记录，多个文档组成集合，集合相当于关系型数据库中的表的概念，多个几个组成数据库。

SQL 术语/概念	MongoDB 术语/概念	解释/说明

database	database	数据库

table	collection	数据库表/集合

row	document	数据记录行/文档

column	field	数据字段/域

index	index	索引

table joins		表连接，MongoDB不支持

primary key	primary key	主键，MongoDB自动将 _id字段设置为主键

MongoDB 基本操作

登录、退出

#本地登录（默认实例端口号：--port=27017，可以不写）

> mongo

#登录远程主机的实例

mongo --host 192.168.1.2 --port =27017

#退出MongoDB

exit

数据库操作


#创建数据库（如果数据库不存在，则创建数据库，否则切换到指定数据库）

> use school

#查看所有数据库

show dbs

#删除school数据库

use school

db.dropDatabase()

#显示数据库操作命令

db.help()

集合

MongoDB 的数据保存在集合中，所有存储在集合中的数据都是 Binary JSON 格式，简称 BSON。BSON 是一种类似于 JSON 的二进制形式的存储格式。


#创建info集合

> db.createcollection('info')

#查看集合

方法一：

show tabels

方法二：

show colletctions

#显示info集合操作命令

db.info.help()

文档（增、删、改、查）

插入

#插入一条记录

db.info.insert({"id":1,"name":"jack","hobby":["game","talk","sport"]})

#向指定集合中插入一条文档数据

db.collection.insertOne()

#向指定集合中插入多条文档数据

db.collection.insertMany()

#通过循环批量插入数据

for(var i=1;i<100;i++)db.info.insert({"id":i,"name":"jack"+i})

删除


#删除info集合中id=1的文档

> db.info.remove({"id":"1"})

修改


#修改info集合id=1的name值为"zhangsan"文档

db.info.update({"id":"1"},{$set:{"name":"wzn"}})

查询

#查询info集合所有文档

> db.info.find()

#查询info集合id为1的文档

db.info.findOne({id:1})

#统计记录数

db.info.count()

备份、恢复数据库

导入导出

导出：mongoexport

导入：mongoimport

选项： -d 指定数据库名称；-c 指定集合名称；-f 指定导出哪些列；-o 指定要导出的文件名；-q：指定导出数据的过滤条件。具体命令通过 --help 查看。

备份与恢复

备份：mongodump

恢复：mongorestore

选项：

1 ： -h 指定Mongodb所在服务器的地址也可以指定端口。（例：-h 127.0.0.1:27017）

2 ： -d ：需要备份的数据库实例。

3 ： -o ：备份数据存放的目录 该目录需提前创建。


复制数据库

> db.copyDatabase("db1","db2")     //复制数据库db1 到db2 中

克隆集合

把数据库 db1 的 info 集合克隆到实例端口：27018


# mongo --port 27018

# db.runCommand({"cloneCollection":"db1.info","from":"localhost:27017"})

MongoDB 安全管理

MongoDB 安全管理主要包括 MongoDB 的安全访问控制以及用户权限分配。

限定监听特定 IP 和端口

# vim /usr/local/mongodb/mongodb1.conf

# 只绑定内网卡地址

bind_ip=localhost(ip)

# 只监听指定的端口

port=27017

授权启动

内置数据库用户：read、readWrite

数据库管理角色：dbAdmin、dbOwner、userAdmin

超级用户角色：root

#在db1数据库创建超级用户root,密码：123123

> use db1

db.createUser({"user":"root","pwd":"123123","roles":["root"]})

exit


# 

#关闭 mongodb 服务

mongod -f /usr/local/mongodb/mongodb1.conf --shutdown

# 

# 使用带认证参数的方式启动 mongodb 服务

mongod -f /usr/local/mongodb/mongodb1.conf --auth

# 

# 此时查询数据不显示内容，需要进行授权认证

> use db1

db.auth("root":"123123")


# 

# 当然，实际情况中，我们可以修改配置文件，

#这样别人在访问我们的 MongoDB 时，我们可以指定一个有相应权限的用户给他进行登录并操作

vim /usr/local/mongodb/mongodb1.conf

auth=true    //添加

进程管理

查看当前正在运行的进程(获取高消耗资源的进程opid）

> db.currentOp()

终止正在运行的高消耗资源的进程（括号内跟上面获取到的opid值）

db.killOp(opid)

MongoDB 监控

查看数据库实例的状态信息

db.serverStatus()

查看当前数据库的统计信息

db.status()

查看集合统计信息

db.users.status()

查看集合大小

db.users.dataSize()

此外，我们可以通过 Web 界面查看系统监控信息，只需要修改配置文件


# vim /usr/local/mongodb/mongodb1.conf

httpinterface=true    //添加

通过 Web 页面可以看到：

1)当前 MongoDB 的所有连接

2)各个数据库和集合的访问统计，包括：Reads、Writes、Queries、GetMores、Inserts、Updates、Removes

3)写锁的状态

4)日志文件的最后几百行

5)所有的 MongoDB 命令

第三方监控工具

我们可以通过 Nagios 中配置使用 MongoDB 插件来监控 MongoDB 数据库。
