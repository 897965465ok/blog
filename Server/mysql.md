# mysql

```
mysql -uroot -p    //连接Mysql -uroot(用户) -p(密码)
exit //退出
show databases; 显示所有数控库注意后面那个单引号 
show databases like '匹配模式'
_:匹配当前位置单个字符
%:匹配指定位置多个字符

获取以my 开头的全部数据库：'my%'
获取M开头，后面第一个字母不确定 最后为database 的数据库; 'm database';
获取以database结尾的数据库'%databse'

use (数据库名) //进入数据库
alter database (数据库名字) charset (字符集) //修改字符集

show tables; //查看当前数据库的所有表

create database (数据库名) charset utf8 字符集 //创建数据库
库选项 :数据库的相关属性
字符集 : charset 字符集,代表着当前数据库指定的字符集
校对集: collate 校对集
列:create database test2 charset utf8;

drop database  (数据库名;) //删除数据库
列:drop database test2;

Myseql 表/列可以改名,databases不能改名

如果忘记数据放那个文件怎么办可以到根目录找my.ini 找到datadir 后面就是路径了
create table stu(                       //创建表语句 
    snum int,
    sname varchar(10)
)engine myisam charset utf8;                  //engine 是指引擎

drop table stu; //删除表stu

rename table stu to newstu; 修改表名字

清空表数据:truncate(表名字);
truncate 和delete 的区别;
truncate 相当与删表后再重建一张同样解构的表,操作后再得到一张全新的表,
delete 是从删除所有的层面来操作的

truncate 相当于把旧的学籍表扔了重画一张 
delete 相当于用橡皮把学籍表的数据库擦掉

delete from class where id=2; //删除id等于二的 如果没有条件将删除整张表

select * from newstu; 打印出这张表

set names gbk;设置编码

 
create table (表名) (字段名 字段类型 [字段属性])[表选项] //创建表
表选项 
1. engine:存储引擎
2. charset:字符集 只对自己表有效(级别比数据库高级)
3. Collate ：校对集

//创建表
create table (新表) like 表名://只要使用数据库.表名就可以在任何数据库下访问其他数据库的表名了

列：
create table newclass like test.class //复制表结构不复制数据

列：
create table class( //创建表   如果没有进入数据库可以这样写 demo.class
    id int primary key auto_increment,
    sname varchar(10) not null default'',
    gender char(1) not null default '',
    company varchar(20) not null default'',
    salay decimal(6,2) not null default 0.00,
    fanbu smallint not null default 0
)engine myisam charset utf8; //表选项


insert into (表名字) values  插入表数据
(1,'zhangsan'), //ID  值
(2,'lisi'),
(3,'wangwu);

也可以按属性名插入 insert into (表名字) values
insert into class
    -> (sname,gender) //按照属性名
    -> values 
    -> ("小明","男");

 insert into my values('小红'),('小黄'); //插入多行

update class  set  (属性名)=20000 //修改数据
where    (表达式)       id=6 and name="小明"; 可以别的条件 name 或age都行不一定是id
列: update class set fanbu = 300  where id=9;

select (属性名)可以多个 from (表名) where (条件) //查询语句

列:select sname,company,salay from class where id=6; // 查询的属性名 查询的表达式  


建表的过程就是定义字段的过程 就键表头的过程
mysql 三大列类型
数值型:
   整形
   Tinyint/smallint/mediumint/int/bigint(M)
   小浮点型
    Float(M,D)/decimal(M,D)
   M叫精度 ->代表总位数,而D是标度代表小数位
    FLOAT[(M[, D])]数据类型是单精度浮点数， 默认大小为24位数字，精度大约7位数字（经测试为6位），当设置M<=24等于4字节，否则自动转换为DOUBLE类型;同时设置M和D时不进行自动转换。
   字符类型  所以一般账号存钱的都是用分来存
    Char(M)/Varchar(M)/Text文本类型/blob

    Char(6)固定长度如果不够按空格补齐取出的时候在把右侧空格删掉 //(1)是字节

    Varchar(100)不是固定长度 但是在数据库中有个隐藏的1-2字节的标识位 标识字节的长度 缺点也是这个标识位他会占用内存  每次查找都是先找到标识位 如果不够按空格补齐取出的时候在把最后空格删掉 //(100)是字符数
    text :不必给默认值 

    Blob :是二进制类型,用来存储图像等信息 意义:2进制,0-255 都有可能 blob 在于防止因为字符集问题,导致信息丢失;
    

   时间日期类型
   Date 日期/ time 时间/ datetiem 时间时间类型/ year年类型

   Date 存放 年-月-日 范围1000-01-01到9999-12-31

   tiem '00:00:00' ;

   datetiem  1000-01-01 00:00:00
   
   timestamp  CURRENT_TIMESTAMP 使用这个默认值时会取出当前服务器时间

tee D:/3-26.sql  //记录今天命令行 到硬盘;

desc ooxx; 查看ooxx表

往行里面增加一个属性名
alter table (表名)ooxx add score tinyint(5)类型后面的5代表宽度  unsigned(修饰这个代表无符号) (是否可以为空)not null (默认值)default 0;

alter table ooxx add gender char(1) not null default '' after name; 如果想加在name后面这样写 

如果想新建一个列且在表的最前面，用first

删除列：

alter table ooxx drop name; 删除name列;

修改列类型：

alter table (表名) modify (列名) (新类型) not null default '';

alter table ooxx modify name char(4) not null default '';

修改列名及列类型:
alter table (表名) change (旧列名) (新列名) (新类型) (新参数)

alter table test3 change name name2 char(10) not null default '';

(M) unsigned zerofill
M代表补0的宽度 和zerofill 配合使用
+-------+-----+-------+---------+
| sname | age | score | message |
+-------+-----+-------+---------+
| 小     |  18 |    28 |   00005 | <-5代表5个宽度
+-------+-----+-------+---------+

zero代表0 fill代表填充

查看某个表的建表语句
show create table (表名);
show create table gooods;

把ecshop 中的商品表数据批量导入测试goods表
insert into demo.goods
select 
goods_id,cat_id,goods_sn,goods_name,click_count,goods_number,market_price,shop_price,add_time,is_best,is_new,is_hot from test.ecs_goods;

create table goods( 
    goods_id mediumint(8) unsigned not null auto_increment,
    cat_id smallint(5) unsigned not null default '0',
    goods_sn varchar(60) not null default '0',
    goods_name varchar(120) not null default '',
    click_count int(10) unsigned not null default '0',
    goods_number smallint(5) unsigned not null default '0',
    market_price decimal(10,2) unsigned not null default '0.00',
    shop_price decimal(10,2) unsigned not null default '0.00',
    add_time int(10) unsigned not null default '0',
    is_best tinyint(1) unsigned not null default '0',
    is_new tinyint(1) unsigned not null default '0',
    is_hot tinyint(1) unsigned not null default '0',
   PRIMARY KEY(goods_id) //这个东西是组件
)engine=myisam default charset=utf8;

show tables; //查看所有表
show tables like '匹配模式' //和数据库匹配一样
desc class;// 显示表结构

alter table (表名) charset  utf8;//修改字符集
rename table (旧表名) to (新表名) //修改表名字
alter table (表名) add (新字段) (列类型) 比如 int(10) not null default 0; //增加列 
alter table (表名) change (旧列名) (新列名) int(10) not null default 0; //修改列名字
alter table (表名) drop name ;//删除 字段名
drop table (表名,表名); //删除表  批量用,逗号隔开

insert into (表名) (字段,字段) values(值,值) //插入数据
insert into values() //这样插入必须一一对应

select (字段) from(表名) //搜索所有字段
select name,age, from class //指定字段
select * from where  shop_price>3000 where 条件 //搜索符合条件的字段
delete from (表名) where 没有条件表示全部删除 
update (表名) set (字段) where (条件)

show variables like 'character_set_%' //查看字符集
client //客户端字符集
connection//中间层字符集
database //数据层字符集
results//返回数据结果字符集
如果使用 set names utf8 下列三种都会变成 gbk
| character_set_results    | gbk  
| character_set_connection | gbk    
| character_set_client     | gbk         
如果只想修改一种怎么办 set character_set_client=utf8;

mysql +-/* 碰到字符串自动转成0

在mysql 中有项规定 mysql 记录的字节长度不能超过65535个字节

列属性

有称为字段属性, 在mysql 中一共有6个属性: null 默认值，列描述，主键，唯一键和自动增长
not null  表示字段不可为空
default 表示默认值 如果你想触发默认值values('jojo',default) 这样写
comment 专门给开发人员进行维护的一个注释说明 comment '字段描述'
主键:一个表的唯一的ID只能有一个 语法 primary key 
方案1： 直接在需要当做主键的字段之后,增加primary key 属性来确定主键
列: create table my_pri1(
 username varchar(10) primary key
)charset utf8;

方案2: 在所有的字段之后增加primary key 选项：primary key(字段信息)

列: create table my_pri2(
 username varchar(10),
  primary key(username)  
)charset utf8;

建表后增加
alter table (表名) add primar key(字段);

删除主键
alter table my_pri3 drop primary key;


+----------+-------------+------+-----+---------+-------+
| Field    | Type        | Null | Key | Default | Extra |
+----------+-------------+------+-----+---------+-------+
| username | varchar(10) | NO   | PRI |         |       |
+----------+-------------+------+-----+---------+-------+
PRI 主键描述

复合主键:

案列:有一张学生选修课表: 一个学生可以有多个选修课,一个选修课也可以由多个学生来选：但是一个学生在一个选修课中只有一个成绩.

中间表的时候用到
create table my_score(
student_no char(10),
course_no char(10),
score tinyint not null,
primary key (student_no,course_no) //复合主键  两个主键合起来不能重复
) charset utf8;

主键约束
主键一旦增加,那么对对应的字段有数据要求
1. 当前字段对应的数据不能为空
2. 当前字段对应的数据不能重复

insert into my_score values('001','数学',100);
insert into my_score values('002','语文',90);
insert into my_score values('003','英语',100);

主键分类

业务主键: 主键所在的字段,具有的业务意义(学生id)

逻辑组件: 自然增长的整型(应用广泛)

自动增长
自动增长：auto increment, 当给定某个字段该属性之后,
该列的数据在没有提供正确的数据的时候 系统会根据之前已经存在的数据进行自动增加后,
填充数据.
通常自动增长用于逻辑主键.

原理

自动增长的原理:
1. 在系统中有维护一组数据,用来报错当前舒勇了自动增长属性的字段,
记住当前对应的数据值,在给定一个指定的步长
2. 当用户进行数据的插入的时候,如果没有给定值,系统在原始值上在加步长变成新的数据
3. 自动增长的触发: 给定属性字段没有提供值
4. 自动增长只适用于数值

使用自动增长
基本语法 :在字段之后增加一个属性 auto increment

列:
create table auto(
    id int  primary key auto_increment, //自动增长
    name varchar(10) not null comment '用户名',
    pass varchar(50) not null comment '密码'
) charset utf8;

1. 一张表只能拥有一个自增长  

查看自动增长 show table (表名)

-- 修改自增长 
alter table auto auto_increment =10;

-- 删除自增长

 alter table auto modify id int;


-- 显示 自增长的初始值
show variables like 'auto increment%';

-- 增加自增长 
alter table (表名) modify id int auto_increment;


唯一键 
 unique key 用来保证对应的字段中的数据唯一的.

1. 唯一键在一张表中可以有多个
2. 唯一键允许字段数据为null null可以为多个(null不参与比较)

 创建唯一键
1. 直接在表字段之后增加unique[key]
2. 唯一键可以为多个也可以为NULL

 create table unique1(
  id int primary key auto_increment,
  username varchar(10) unique  //增加唯一键
 )charset utf8;

也可以这样:

create table unique2(
id int primary key auto_increment,
  username varchar(10),
  unique key(username) //增加唯一键
)

create table unique3( 
id int primary key auto_increment,
  username varchar(10)
)
alter table unique3  add  unique key(username); //建表后增加

.查看唯一键
desc  unique1;

+----------+-------------+------+-----+---------+----------------+
| Field    | Type        | Null | Key | Default | Extra          |
+----------+-------------+------+-----+---------+----------------+
| id       | int(11)     | NO   | PRI | NULL    | auto_increment |
| username | varchar(10) | YES  | UNI | NULL    |                |
+----------+-------------+------+-----+---------+----------------+

insert into unique1 values(null,'小明');

删除唯一键 

先找出键名字
show create table unique1;

| unique1 | CREATE TABLE `unique1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)  //第一个就是唯一键名字
) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 |

alter table 表名 drop index (唯一键名字); //删除唯一键

# 修改

先删除在增加(无法修改)

复合唯一键
唯一键与主键一样可以使用多个字段来共同保证唯一性
一般主键都是单一字段(逻辑主键),而其他需要唯一性的内容都是由唯一键来处理;


表关系

表关系: 表与表之间有什么样的关系, 每种关系应该如何设计表结构

一对一

一对一：一张表中的一条记录一与另外一张表中最多有一条明确的关系:通常,此设计方案保证在两张表中使用同样的主键即可

 学生表
 +----------+-------------+------+-----+---------+----------------+
| 学生id  (PRI)| 姓名   | 年龄   | 性别 | 婚否    | 住址 |
| username | varchar(10) | YES  | UNI | NULL    |                | 
+----------+-------------+------+-----+---------+----------------+
表的使用过程中：常用的信息经常去查询,而不常用的信息偶尔才会用到.

解决方案:将两张表拆分,常见的一张表 ,不常见的一张表

常用表：
+----------+-------------+------+-----+---------+----------------+
| 学生id  (PRI)| 姓名   | 年龄   | 性别 | 
+----------+-------------+------+-----+---------+----------------+
不常用表:
+----------+-------------+------+-----+---------+----------------+
| 学生id(PRI)| 籍贯   | 婚否   | 住址 |
+----------+-------------+------+-----+---------+----------------+

#一对多

一对多,通常也叫作多对一的关系。通常一对多的关系设计的方案,在‘多’关系的表中去维护一个字段,这个字段是‘一’关系的主键。

母亲表
+----------+-------------+------+-----+---------+----------------+
| 母亲id  (PRI)| 姓名   | 年龄   | 性别 | 
+----------+-------------+------+-----+---------+----------------+

孩子表
+----------+-------------+------+-----+---------+----------------+
| 孩子id  (PRI)| 姓名   | 年龄   | 性别 | 母亲ID  //子表记录母亲ID
+----------+-------------+------+-----+---------+----------------+

多对多

多对多: 一张表中的一条记录在另外一张表中科院匹配多条记录,反过来也一样,
多对多的关系如果按照多对一的关系维护:就会出现一个字段中有多个其他表的主键,
在访问的时候就会带来不便。
既然通过两张表自己增加字段解决不了问题，那么久通过第三张表了解决.

母亲表
+----------+-------------+------+-----+---------+----------------+
| 母亲id  (PRI)| 姓名   | 年龄   | 性别 | 
| T1| 张三 | 30  | 男 | 
| T2| 韩梅梅 | 20  | 女 | 
  T3| 杰伦 | 50  | 男 | 
+----------+-------------+------+-----+---------+----------------+
中间表
+----------+-------------+------+-----+---------+----------------+
| ID  | 母亲id | 儿子id   
| 1   |  T1| S1 | //通过中间表就可以找到所有子id 也可以通过 子ID找到父信息
| 2   |  T2| S1 |
| 3   |  T1| S1 |
| 4   |  T3| S1 |
+----------+-------------+------+-----+---------+----------------+

孩子表
+----------+-------------+------+-----+---------+----------------+
| 孩子id  (PRI)| 姓名   | 年龄   | 性别
| S1| 小明 | 30  | 男 | 
| S2| 小红 | 20  | 女 | 
| S3| 武藤兰 | 20  | 女 |
| S4| 东尼大木 | 90  | 男 |
+----------+-------------+------+-----+---------+----------------+

主键冲突 

create table student(
  id varchar(10) primary key,
  name varchar(10) not null
)charset utf8;
insert into student values
('1','小红'),
('2','小黑'),
('3','小白'),
('4','小黄');

冲突解决方案 
1. 
 insert into (表名) values('值','值') on duplicate  key update (字段) = (新值)
 insert into student values('1','小红') on duplicate key update name ='小明'; //如果主键有冲突 的时候把name 改成小明

 2. 
 rplace into (表名) values(值,值); //第二种方法
 replace  into student values('1','小红');



#  复制数据

insert into [表名字(字段)] select (字段) from (表名字)

列:
insert into simple(name) select name from simple;

修改列数据

update (表名) set (字段名) = (新值)  where (条件) limit (数量)

update my set username = '小白' where username = '小黄' limit 1;

删除数据

1、删除数据的时候尽量不要全部删除 、应该使用where进行判定

1、删除数据的时候可以使用limit来限制要删除的具体数量

delete  from (表名) where(条件)  limit(数量);

delete from my where id > 1 limit 1;

删除后 auto increment 没有重置为重一开始 怎么办 可以使用 truncate (表名)解决


查询数据

完整查询指令

select  (字段列表) from (表名) where(条件) group by (分组) having (条件) order by (排序) limit (限制)

select * from my;

select distinct * from goods;  //distinct去重

字段列表;有的时候需要从多张表中获取数据，在获取数据的时候，可能在不同表中有同名的字段，需要将同名的字段命名成不同名的：别名 alias

 (字段名) as (别名)

 select distinct name as name1, name as name2 from my;


 from  数据源 只有是一个符号二维表结构的数据即可
 1、单标数据

 from （表名）

 2、  多表获取数据 
  from (表1,表2)  多表

  select * from my, goods;
  结果: 两张表的记录数相乘，字段数拼接
  本质：从第一表取出一条记录，去拼凑第二张表的所有记录，保留所有结果
  好像没什么用

  动态数据

  from 后面跟的数据不是一个实体表，而是一个从表中查询出来得到的二维结果表(子查询) 
  基本语法: from (select (字段列表) from (表)) as (别名)

  select * from (select name from my ) as ooxx;

  where 子句
  where 子句：用来重表获取数据后根据运算条件筛选 
  
  group by  表示分组：根据指定的字段,将数据进行分组:分组的目标是为了统计

  基本语法：
  group by 字段名;
  group by goodsid //根据字段名分组 但是只会取出第一条

  select  * from student group by classId;
  
  利用一些统计函数：(聚合函数)
  count()// 如果值为*求绝对行数,如果为值不求null
  avg(); 求平均值
  sum(); 求和
  max() 求最大值
  min():求最小值

select goods_name,cat_id,sum(goods_number) from goods group by cat_id;

group_concat(字段)//合并字段



多分组
将数据按照某个字段进行分组之后，对已经分组的数据进行再次分组


基本语法: group by 字段1，字段2 //先按照字段1进行排序，之后在按照字段2进行排序

分组排序 mysql 中分组默认有排序功能:按照分组字段进行排序，默认是升序
基本语法 group 字段 [asc|desc] , 字段[asc|desc]

回溯统计
当分组进行多分组之后,往上统计的过程中，需要进行层层上报，将这种层层上报统计的过程称之为回溯统计,每一次分组向上统计的过程都会产生一次新的统计数据而且当前数据对应的分组字段为NULL
基本语法: group 字段{asc|desc} with rollup;

select goods_id ,count(*) from goods group by goods_id with rollup;

Having 子句
Having 的本质和where一样, 是用来进行数据条件筛选。

where 不支持聚合函数 所以使用这个
select goods_id,count(*) from goods group by cat_id having count( *)>=4;

order by 子句
order by 排序：根据校对规则对数据进行排序
基本语法： order by (字段)[asc|desc] //升序，默认的

-- 商城按照商品价格排序
select * from  goods order by shop_price desc;

order by 也可以像group by 一样进行多字段排序先按照第一个字段排序在按照第二个
select * from goods order by shop_price desc,cat_id desc limit 5;

limit //限制数量

基本语法: 
limit 
offset //偏移量 重那开始  
length //多少条

select * from goods order by shop_price desc,cat_id desc limit 0,5;

in 代表多个选择条件 比如 我要查找cat_id类 cat_id in('6','4') //找出6和4类

列：select goods_name,cat_id from goods where cat_id in('4','6') 

ls 运算符

ls 是专门用来判断是否为Null 
语法:  is null | is not null 

列: 
select goods_name,cat_id from goods where cat_id is null  

like 运算符: 是用模糊匹配

_: 匹配对应的单个字符
%：匹配多个字符
select * from student where name like '小%'

. 联合查询

. 基本概念

 联合查询是可合并多个相似的选择查询的结果集。等同于将一个表追加到另一个表,从而实现将两个表的查询结果到一起，使用谓词为union 或unionall。

 联合查询: 将多个查询的结果合并到一起(纵向合并):字段数不变，多个查询的记录数合并。


.应用场景

1 、 将同一张表中的不同结果(需要对应多条查询语句来实现)、 合并到一起展示数据 
男生身高升序排序、女人身高降序排序
2 、 最常见：在数据量大的情况下、会对表进行分表操作、需要对每张表进行部数据统计、
使用联合查询来将数据存放到一起显示
qq1 表获取在线数据  
qq2 表获取在线数据
 ——》 将所有在线数据显示出来

 . 基本语法 //作用是将两个查询语句查询后合成
 select 语句
 union[union 选项]
 select 语句;
 union 选项:
 distinct 去重 去掉完全重复的数据(默认的)
 all 保存所有的结果

--- 获取商品价格
 
列：//
select * from goods  
union all  //默认为(distinct)  
select * from goods;
注意细节：union理论上只有保证字段数一样,不需要每次拿到的数据对应的字段类型一致.
永远只保留第一个select语句对应的字段名字。

order by 的使用

1、在联合查询中,如果要使用order by , 那么对应的select语句必须使用括号扩起来
2、order by 在联合查询中使用必须要配合 limit (数量)
列: 
(select * from goods where  cat_id  = 16 order by shop_price asc limit 10)
union all
(select * from goods where  cat_id  = 4 order by  shop_price  desc limit 10);


.连接查询
连接查询：将多张表连到一起进行查询(会导致记录数行和字段数列发生改变)

. 连接查询的意义
在关系型数据库设计过程中，实体（表）与实体之间是存在很多联系的。在关系型数据库表的设计过程中，
遵循着关系来设计：一对一， 一对多，和多对多，通常在实际操作的过程中需要利用这层关系来保证数据的完整性。

.连接查询分类
 
1. 交叉连接
 交叉连接： 将两张表的数据与另外一张表彼此交叉
 1 、从第一张表依次取出每一条记录
 2 、取出每一条记录之后，与另外一张表的全部记录挨个匹配
 3 、没有任何匹配条件，所有的结果都会进行保留
 4 、记录数 = 第一张表记录数*第二张表记录数; 字段数 = 第一张表字段数 + 第二张表字段数（笛卡尔积）
-- 交叉连接
select * from student cross join class;
公式   总记录 =  (表一每一条数据 * 表二总行数)+  (表二每一条数据 * 表一总行数)
应用： 没有实际应用

#内连接
内连接：inner join 从一张表中取出所有的记录去另外一张表中匹配：利用匹配条件进行匹配，
成功了则保留，失败了放弃。

。原理
1、从第一张表中取出一条记录，然后去另外一张表中进行匹配
2、利用匹配条件进行匹配：
2.1、匹配到：保留，继续向下匹配
2.2、匹配失败：向下继续，如果全表匹配失败，结束
3. 因为表的设计通常容易产生同名字段,尤其是ID所以为了避免重名的出现错误,通常使用(表名.字段名)，
来确保唯一性。
4. 通常，如果条件中使用到对应的表名，而表名通常比较长，所以可以通过表别名来简化
5. 内连接匹配的时候，必须保证匹配到才会有数据
6. 内连接因为不强制必须使用条件(on)因此可以在数据匹配完成之后,使用where条件来限制，效果与on一样(建议使用on)


. 语法
基本语法：表1[inner] join 表2 on 匹配条件;

-- select * from (表一) inner join (表二) on (条件) //基本语法
-- select * from student inner join class; //可以，没有条件 （默认为笛卡尔积);
-- select * from student inner join class on class.class_id = student.student_id; // 利用表一的数据依次匹配表二的数据
-- select * from student as a inner join class b on b.class_id = a.student_id;  // 利用别名查询
-- select * from student as a inner join class b where b.class_id = a.student_id;  // 也可以使用where 但是不建议用

.应用
 内连接通常是在对数据有精确要求的地方使用:必须保证两种表中都能进行数据匹配.

 外连接：

 外连接： outer join, 按照某一张表作为主表(表中所有记录最后都会保留),根据条件去连接另外一张表，
 从而得到目标数据。

 外连接分为两种:  左外连接(left join)和右外连接(right join)
 左连接：左表是主表
 右连接：右表是主表
 
 .原理
 1、确定连接主表：左连接就是left join 左边的表为主表：right join 就是右边为主表
 2、拿主表的每一条记录，去匹配另外一张表(从表)的每一条记录
 3、如果满足匹配条件：保留，不满足即不保留
 4、如果主表记录在从表中一条都没有匹配成功，那么也要保留该记录：从表对应的字段值都为null

 .语法
  基本语法：
  左连接：主表left join 从表 on 连接条件
  右连接：从表 right join 主表 on 连接条件
  左连接对应的主表数据在左边；右连接对应的主表数据在右边

 select * from student as a left  join class b on b.class_id = a.student_id; 
 select * from student as a right  join class b on b.class_id = a.student_id; 
  .应用 
  非常常用的一种获取方式：作为数据获取对应主表以及其他数据(关联)
  
 自然连接

.Using 关键字是在连接查询中用来代替对应的on关键字的，进行条件匹配

. 原理
1、在连接查询时，使用on的地方用using代替
2、使用using的前提是对应的两张表连接的字段是同名(类似自然连接自动匹配)
3、如果使用using 关键字，那么对应的同名字段，最终结果中只会保留一个。

.语法
基本语法: 表1(inner,left,right) join 表2 using(同名字段);

select * from student as a left  join class b using(id);

. 子查询 
. 什么是子查询
. 子查询概念
 子查询： shu query 

子查询是一种常用计算机语言 select-sql 语言中嵌套查询下层的程序模块 。 当一个查询是另一个查询的条件时，称之为子查询。
子查询：指在一条select语句中，嵌入了另外一条select 语句, 那么被嵌入的 select 语句称之为子查询语句。

。主查询概念

主查询：主要的查询对象，第一条select,确定的用户所有获取的数据目标(数据源),
已经要具体得到的字段信息

。 子查询和主查询的关系
1、子查询是嵌入到主查询中的，
2、子查询是辅助主查询的：要么作为条件，要么作为数据源
3、子查询其实可以独立存在：是一条完整的select语句

. 子查询分类

. 按功能划分
1、标量子查询：子查询返回的结果是一个数据(一行一列)
   .概念
   .标量子查询：子查询得到结果是一个数据(一行一列)
   .语法
    基本语法:select *from 数据源 where 条件判断 =/<>(select 字段名from 数据源 where 条件判断)// 子查询得到的结果只有一个值
--知道一个商品的名字,想知道他属于是哪个经销商的
-- 通过商品表获取他的id 
-- 通过id 获取经销商
-- 标量子查询实现
-- 列:
select * from goods where goods_id = (select goods_id from goods where goods_name = '诺基亚N85');
//需求决定主查询，条件决定子查询

2、列子查询：返回的结果是一列(一列多行)
. 概念

 列子查询： 子查询得到的结果是一列数据(一列多行)

 基本语法：

 主查询 where 条件 in(列子查询);

 -- 想获取已经是火热商品的所有商品ID
 -- 1. 找出商品中所有是hot 的商品
 -- 2. 通过商品找出ID

select * from goods where goods_sn  in(select goods_sn from goods where (is_hot = 1)); //查询语句有问题明天问

3、行子查询：返回的结果是一行(一行多列)
  .概念
  行子查询：子查询返回的结果是一行多列

  .行元素
  行元素：字段元素是指一个字段对应的值，行元素对应的就是多个字段：多个字段合起来作为
  为一个元素参与运算，把这种情况称之为行元素。

 .语法
 基本语法：
 主查询 where 条件 (构造一个行元素) = (行子查询);

-- 需求：获取商品价格最大，且是新的商品
-- 1、 求出商品最大的值;
-- 2、 求出商品是最新的值;
-- 3、 求出对应的商品;
列：
select * from goods where (shop_price,is_new) = (select max(shop_price),(is_new = 1) from goods);

错误示范：

select * from goods having shop_price = max(shop_price) and (is_new = 1);

-- 1、 having 是在group 之后的：所以使用having 就默认会有group by 语句
-- 2、 group by 一旦执行只会返回一行

. 总结 
已经学过三个字查询: 常见的三个子查询
标量子查询。列子查询和行子查询：都属于where 子查询

4、表子查询: 返回的结果是多行多列(多行多列)
 
 .概念
 表子查询: 子查询返回的结果是多行多列,表子查询与行子查询非常相似,只是行子查询需要产生行元素，而表子查询没有。
 
 行子查询用于where 条件判断:where 子查询
 表子查询用于from 数据源:from 子查询


 . 语法

  基本语法:

  select 字段表 from (表子查询)[where][group by][having][order by][limit]

-- 获取每个班上最高身高的学生
-- 1、将每个班最高的学生排在前边 ：order by 
-- 2、在针对结果进行group by ：保留每组第一个
select * from (select * from student order by height desc) as temp group by class_id;

5、exists子查询: 返回的结果是1或0(类似布尔操作)
 .概念
  exists 子查询：查询返回的结果只有0或者1 1 代表成立 ， 0 代表不成立

  . 语法
  
   基本语法 ： where exists(查询语句); //exists 就是根据查询得到的结果进行判断：如果结果存在返回1否则返回0
  
  where 1 :永远为真

 列: select *from class as c where exists(select stu_id from student as s where s.calss_id = c.class_id)

. 按位置分
where子查询: 子查询出现的位置在where条件中
from 子查询: 在查询出现的位置在from 数据源中(做数据源)

子查询中特定关键字的使用

. in 
主查询 where 条件in(列子查询)

.any  .some :任意一个 
 any(任意一个)


.all 全部的 
 = all(列子查询): 等于里面所有


整库数据备份与还原
整库数据备份也叫SQL数据备份：备份的结果都是SQL指令
在MYSQL 中提供一了一个专门用于备份sql的客户端：mysqldump.exe

.应用场景

sql备份是一种mysql非常常见的备份与还原方式,sql备份不只是备份数据，还备份对应的sql指令（表结构）：即便是数据库遭到毁灭性破坏(数据库被删),那么利用sql备份依然可以实现数据还原。

SQL备份因为需要备份结构，因此产生的备份文件特别大，因此不适合大型数据备份,也不适合数据变换频繁型数据库备份.

.应用方案

.sql备份

SQL备份用到的是专门的备份客户端，因此还没有与数据库服务器进行连接

基本语法：mysqldump.exe - hPup 数据库名字(表1表2)>备份文件地址

备份可以有三种形式：

1、整库备份 (只需要提供数据库名字)

-- 整库备份  mysqldump -hlocalhost -p3306 -uroot -p897965465 test > G:\test.sql

-- 整库还原 mysql -uroot -p897965465 demo < G:\test.sql
-- 整库还原 source g:/test.sql; 这个是要进入数据库才能操作

2、单表备份 ：数据库后面跟一张表

3、多表备份 ：数据库后面跟多张表

-- 多表备份 mysqldump -uroot -p897965465 (数据库) 表名 表名 > G:\demo2.sql

用户权限管理

用户权限管理：在不同的项目中给不同开发者不同的操作权限，为了保证数据库数据的安全。
通常，一个用户的密码不会长期不变所以需要经常变更数据库用户密码来确保用户本身安全(mysql 客户端用户)

.用户管理

mysql 需要客户端进行连接认证才能进行服务器操作：需要用户信息。mysql中的所有的用户信息都是保存在mysql数据库下的user表中。

-- select * from mysql.user

在mysql 中,对用户管理中，是由对应的Host和User共同组成主键来区分用胡
User: 代表用户的用户名
Host :代表允许访问的客户端(IP或者主机地址).如果host 使用*代表所有用户可以访问

. 创建用户
 
 有两种方法创建用户：
 1、直接使用root 用户创建mysql.user表中插入记录(不推荐)
 2、专门创建用户的sql指令
 基本语法：create user 用户名 identified by '明文密码';
 create user '(root)'@'(ip地址 %代表所有id)' identified by '(用户密码)'
 -- create user 'root1'@'%' identified by '123456'; // 创建用户
 -- select * from mysql.user   //查看用户

. 删除用户  
 注意: mysql 中user 是带着host本身的(具有唯一性)

 基本语法:drop user 用户名@host;
 drop user root @ '%'
 drop user user1@'%';

. 修改用户密码
 Mysql 中提供了多种修改的方式:
 1、 使用专门的修改密码的指令
 基本语法：set password for 用户 =password('新密码')
 -- set password for 'user1'@'%' = password('654321')
 2、使用更新语句update来修改表
 基本语法： update mysql.user set password = password('新密码') where user = " and  host =";

 
. 权限管理

在mysql 中将权限管理分为三类
1、 数据权限: 增删改查(select|update|delete|insert)
2、 结构权限：结构操作(create|drop)
3、 管理权限：(create user| grant| revoke):通常只给管理员

. 授予权限：grant

 将权限分配给指定用户

 基本语法:grant [那种权限,drop] on (数据库).(表名) to (那个用户) //all privileges 代表选择全部,如果想选择全部数据库可以用*号代替。

 -- grant all privileges  on *.*  to 'demo'@'%';

. 取消权限：revoke
基本语法: revoke [权限列表] on (数据库).(表名) from (用户)。
revoke select on *.* from 'demo'@'%';

. 刷新权限：flush
flush 刷新 ,将权限写入表里面去
基本语法: flush privileges;

.密码丢失解决方案

如果忘记root 用户密码，就需要去找回或者重置root 用户密码

1、停止服务
net stop mysql
2、重新启动服务:mysql.exe --skip-grant-tables//启动服务器但跳过权限

3、当前启动的服务器没有权限概念:非常危险任何客户端不需要任何信息都可以登录，而且是root权限：新开cmd使用mysql.exe登录即可

4、修改root用户的密码: 指定 用户名@host
update mysql.user set password = password('root') where user='root' and
gist = 'localhost';

5、 关闭服务 net stop mysql 重启服务  net start mysql


.外键

.外键概念

如果公共关键字在一个关系中是主关键字，那么这个公共关键字被称为另一个关系的外键.由此可见，外键表示两个关系之间的相关联系。以另一个关系的外键作主关键字的表被称为主表，具有此外键的表被称为主表的从表。外键又称作外关键字。

外键: foreig key
一张表(A)中有一个字段,保存的值是指向另外一张表(B)的主键
B:主表
A:从表

.外键的操作

.增加外键

mysql中提供了两种方式增加外键

方案1:在创建的时候增加外键(类似主键)

// 基本语法: 在字段之后增加一条语句 
create table student(
 id int primary key auto_increment,
 name varchar(20),
 teacher_id int,
 constraint student_tea foreign key(teacher_id) references teacher(id) 
)ENGINE=INNODB DEFAULT CHARSET=utf8;


方案2:建表后增加外键

alter table 从表 add [constraint `外键名`] foreign key(外键字段) references主表(主键);

列： alter table child add foreign key(father_id) references father(id);


删除外键: alter table (从表) drop foreign key (外键名)    // 敲入命令行show create table (从表就可以看到CONSTRAINT 后面就是外键名了); 
列： alter table student drop foreign key student_tea;

mysql> desc my_foreign;
+----------+-------------+------+-----+---------+----------------+
| Field    | Type        | Null | Key | Default | Extra          |
+----------+-------------+------+-----+---------+----------------+
| id       | int(11)     | NO   | PRI | NULL    | auto_increment |
| name     | varchar(10) | NO   |     | NULL    |                |
| class_id | int(11)     | YES  | MUL | NULL    |                |
+----------+-------------+------+-----+---------+----------------+

MUL: 代表多索引,外键本身是一个索引,外键要求外键字段本身也是一种普通索引

create table  child (  //第一种
child_id int(10) not null references father(child_id) //增加外键 这种可以不用主表主键
)ENGINE=INNODB DEFAULT CHARSET=utf8;

create table  child ( //第二种
foreign key(child_id) references father(id) //这种必须要主表主键不然会报错
)ENGINE=INNODB DEFAULT CHARSET=utf8; 


 // 敲入命令行show create table child; 

child | CREATE TABLE `child` (  
  
  KEY `child_id` (`child_id`),   // 创建外键时自动增加的普通索引

 //这张表自动创建  CONSTRAINT `child_ibfk_1` 这个东西叫外键索引

  CONSTRAINT `child_ibfk_1` FOREIGN KEY (`child_id`) REFERENCES `father` (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8 |


--修改外键
alter table my_student add constraint `student_classibkf_1 foreign key`

-- 添加外键的时候，设置级联更新
on update cascade
alter table student add foreign key(teacher_id) references teacher(id) on update cascade;

-- 添加外键的时候,设置级联删除
on delete cascade
alter table student add foreign key(teacher_id) references teacher(id) on update cascade on delete cascade;


视图基本操作

.创建视图
视图的本质是sql指令(select语句)
基本语法:create view 视图名字 as select 指令//可以是单表数据，也可以是连接查询,联合查询或者子查询

.使用视图
视图是一张虚拟表；可以把视图当做表操作但是视图本身没有数据。是临时执行select语句得到的对应结果。视图主要用户查询操作结果

基本语法：select 字段列表from 视频名字[子句]
select * from student_view

.修改视图
 修改视图：本质是修改视图对应的查询语句
 基本语法:alter view 视图名字 as 新 select 指令:

 -- alter view student as select * from studnet  using(class_id)

 .删除视图
 基本语法:dop view (视图名字)

 dop view student_view 

 ## 事务
1. 事务的基本介绍
   1.概念:
   * 如果一个包含多个步骤的业务操作,被事务管理，那么这些操作要么同时成功，要么同时失败
2. 操作:
  1. 开启事务:start transaction;
  2. 回滚:rollback;
  3. 提交:commit; 手动提交

      列子: 张三给李四转账500块
      1. 查询张三账户余额是否大于500
      2. 张三账户 金额 -500
      3. 李四账户 金额 +500 
 这个一整套流程 如果中间有一步出错就回滚 如果没错就提交
 3. mysql 数据库中事务默认自动提交
   事务提交的两种方式：
   *自动提交
    *mysql 就是自动提交的 ,一条mysql语句会自动提交一次事务
    *手动提交
     * 需要先开启事务，在提交
   *修改事务的默认提交方式:
    *查看事务的默认提交方式:select@@autocommit; --1代表自动提交 0 代表手动提交
    *修改默认提交方式：set@@autocommit = 0;

2. 事务的四大特征:
 1.原子性:是不可分割的小操作单位,要么同时成功,要么同时失败.
 2.持久性:当事务提交或回滚后,数据库会持久化保存数据.
 3.隔离性:多个事务之间.相互独立
 4.一致性:事务操作前后,数据总量不变
 3. 事务的隔离级别
  * 概念:多个事务之间隔离的，相互独立的。但是如果多个事务操作同一批数据，则会引发一些问题,设置不同的隔离级别就可以解决这些问题。
  *存在问题:
  1. 脏读:一个事务,读取到另一个事务中没有提交的数据
  2. 不可重复读(虚读):在同一个事务中,两次读取到的数据不一样.
  3. 幻读:一个事务操作(DML)数据表中所有记录,另一个事务添加了一条数据，则第一个事务查询不到自己的修改
  *隔离级别 
   1. read uncommitted :读未提交的
   * 产生问题:脏读、不可重复读、幻读
   2. read committed:读已提交
   *产生问题:不可重复读、幻读
   3. repeatable read:可重复读(mysql默认)
   *产生的问题:幻读
   4. serializable:串行化
   * 可以解决所有的问题

   * 注意隔离级别从小到大安全性越来越高,但是效率越来越低
    * select @@tx_isolation;
  *数据库设置隔离级别:
   * set global transaction isolation level 级别字符串;
* 演示:
  set global transaction isolation level read uncommitted;
  start transaction;
  -- 转账操作
  update account set qian = qian -500 where id = 1;
  update account set qian = qian +500 where id = 2;

  .变量

  .系统变量
  mysql 本质是一种编程语言,需要很多变量来保存数据。mysql中很多的属性控制都是通过mysql中固有的变量来实现的。

  系统内部定义的变量,系统变量针对所有用户(mysql 客户端) 有效.
  
  查看系统所有变量 show variables;
  mysql允许用户使用select查询变量的数据值
  基本语法: select @@(变量名);
  
  修改系统变量:分为两种修改模式

  1、局部修改(会话级别):只针对当前自己客户端当次连接有效
  基本语法:set 变量名 = 新值
  2、全局修改:针对所有客户端进行修改
  基本语法:set blobal 变量名 = 值 或者set @@global.变量名 = 值
  全局修改后需要重启客户端才能生效

  .会话变量
   会话变量也称为之用户变量, 会话变量跟mysql客户端是绑定的，设置的变量，只对当前用户使用的客户端生效.
   
   定义用户变量: set @变量名=值; 或者 set @变量名 :=值 

  在mysql 中因为没有比较符号==, 所有用=代替比较符号,有时候在赋值的时候,会报错,
  mysql为了避免系统分不清楚是赋值还是比较特定增加一个变量的赋值符号 :=

  mysql 是专门存储数据的：允许将数据从表中取出存储在变量中:查询得到的数据必须只能是一行数据
  (一个变量对应一个字段的值):mysql没有数组
  
  语法
  select @name:= stu_name , @age :=stu_age from my_student limit 1;

  select stu_name,stu_age from my_student order by stu_height dese limit 1 into @name @age

  查看变量：select @name,@age;

  .局部变量
  作用范围在begin到end语句块之间。 在改语句块里设置的变量，declare语句专门用于局部变量。
  1、局部变量是使用declare关键字声明
  2、局部变量declare语句出现的位置一定是begin 和end之间(begin end是在大型语句中使用的:函数/存储过程/触发器)
  3、声明语法:declare变量名数据类型 [属性];

  .流程结构
  流程结构:代码的执行顺序
  
  .if分支

  .基本语法
  if在mysql中有两种基本用法
  1、 用在select查询中，当做一种条件进行判断
  基本语法:if(条件,为真结果,为假结果)
  select * ,if(stu_age)>20,'符合','不符合'> as judge from my_student where stu_age >20
  2、用在复杂的语句块中(函数/存储过程/触发器)

  if 条件表达式 then 
     满足条件要执行的语句;
     End if
  .复合语法
 符合语法:代码的判断存在两面性，两面都有对应的代码执行.
 基本语法:
 if条件表达式 then
  满足条件要执行的语句
  else
  不满足条件要执行的语句
  End if;
  whlie循环
  .基本语法
    while 条件 do
    要执行的代码
    End while
   基本语法
   结构标识符
   .函数
   在mysql中。函数分为两类系统函数(内置函数)和自定义函数
   不管内置函数还是自定义函数都是使用select 函数名(参数列表);
   .内置函数
   .字符串函数
   char length():判断字符串的字符数
   length():判断字符串的字节数(与字符集)
   Concat():连接字符串
   instr():判断字符在目标字符串中是否存在存在返回位置不存在返回0
   lcase():全部小写
   left()：从左侧指定位置开始截取字符串
   ltrim():消除左边对应的空格
   mid()从中间指定位置开始截取,如果不指定截取长度直接到最后
   .时间函数
   Now():返回当前时间,日期时间
   Curdate():返回当前日期
   Curtime():返回当前时间
   Datediff():判断两个日期之间的天数差距,阐述日期必须使用字符串格式(用引号)
   Date_add(日期,interval 时间数字 type):进行时间的增加
   .数学函数
   abs():取绝对值
   ceiling():向上取整
   fllor()：向下取整
   pow():求次方
   rand()：随机数0-1之间
   round()：四舍五入
   . 其他函数
   Md5():对数据进行MD5加密
   version():获取版本号
   database():显示当前所在数据库
   uuId():生成一个唯一标识符(自增长):自增长是单表唯一，UUID是整库(数据唯一同时空间唯一)
.自定义函数
自定义函数:用户自己定义的函数
函数：实现某种功能的语句块(由多条语句组成)
1、函数内部的每条指令都是一个独立的个体:需要符合语句定义规范：需要语句结束符分号;
2、函数是一个整体，并且函数是在调用的时候才会被执行，那么当设计函数的时候，意味着整体不能中断:
3、Mysql 一旦见到语句结束符分号，就会自动开始执行
解决方案:在定义函数前，尝试修改临时的语句结束符
基本语法:delimiter
修改临时语句结束符：delimiter新符号[可以使用系统非内置即可$$]
中间为正常sql指令:使用分号结束(系统不会执行:不认识的分号)
使用新符号结束
修改会语句结束符:delimiter;

. 创建函数
自定义函数包含几个要素:function关键字,函数名,参数,确认函数返回值类型,函数体,返回值

函数定义基本语法：
修改语句结束符
Create function 函数名(形参) return 返回值类型
Begin 
    //函数体
    Return 返回值数据;//数据必须与结构中定义的返回值类型一致
End
修改结束符
修改语句结束符(改回来)

--- 创建自定义函数
--- 修改语句结束符
delimiter $$
create function my_func() returns int
begin 
return 10;
end
--结束
$$
--修改语句结束(改回来)
delimiter ;

并不是所有的函数都需要begin和end：如果函数只有一条指令(return),那么可以省略begin和end
---最简函数
create function my_func1() returns int 
return 10;

形参: 在mysql中需要为函数的形参指定数据类型（基本形参可以有多个）
基本语法：变量名 字段类型

create function my_func3(int_1 int) returns int 
return int_1;

.查看函数
1、可以通过查看function状态，查看所有的函数

--查看所有函数 show function status;

2、查看函数的创建语句：show create function (函数名字)

.调用函数

自定义函数的调用和与内置函数的调用是一样的:select 函数名(参数);
select  my_function();


.删除函数

-- drop function my_func1; 删除函数

.注意事项
1、 自定义函数是属于用户级别的：只有当前可短对应的数据库中科院使用
2、 可以在不同的数据库下看到对应的函数，但是不可以调用
3、 自动函数:通常是为了将多行代码集合到一起解决一个重复性的问题
4、 函数因为必须规范返回值：那么在函数内部不能使用select指令:因为一旦执行就会得到一个结果(result set) :select 字段 into @变量(唯一可用)

.函数的列子
-- 需求：从1开始，直到用户传入的对应的值为止，自动求和：凡是5的倍数都不要。
设计：
1、创建函数
2、需要一个形参:确定要累加
2、需要定义一个变量来保存对应的结果
4、内部需要一个循环实现迭代累加
5、循环内部需要进行条件判断控制: 5的倍数

-- 创建一个自动求和的函数
-- 修改语句结束符
delimitert $$
-- 创建函数
create function my_sum(end_value int)  return int
begin 
    -- 声明变量(局部变量):如果使用declare声明变量:必须在函数体其他语句之前
    declare res int default 0;
    declare i int default 1;
   -- 循环处理
   mywhile:while i <= end_value do
         -- 判断当前数据是否合理
         if i % 5 = then
                  -- 5的倍数不要
                  set i = i+1;
                  iterate mywhile;
         end if;
         -- 修改变量:累加加过
         set res = res +1;
         set i = i +1;
       end while mywhile; 
       -- 返回值
       return res;
end
-- 结束
$$
-- 修改语句结束符
delimitert ; 

定义函数完成后;

调用函数: select my_sum(100),my_sum(-100);


.变量作用域
 
 变量作用域:变量能够使用的区域范围

 . 局部作用域

 使用declare 关键字声明(在结构体内：函数/存储过程/触发器),并且只能在结构体内部使用

 1、 declare 关键字声明的变量没有任何符号修饰,就是普通字符串,如果在外部访问该变量,
 系统会自动认为是字段

 .会话作用域
  用户定义的,使用@符号定义的变量,使用set关键字  

  会话作用域: 在当前用户当次连接有效,只要在本连接中，
  任何地方都可以使用(可以在结构内部，也可以跨库)

  .全局作用域

  所有的客户端所有的连接有效：需要使用全局符号来定义
  set global 变量名 = 值;
  set @@global/变量名 = 值;

  通常,在sql编程的时候，不会使用自定义变量来控制全局，
  一般都是自定义会话变量或者在结构中使用局部变量来解决问题。


  .存储过程
  
  .存储过程概念

   存储过程(stored procedure) 是在大型数据库系统中，一组为了完成特定功能的sql语句集,
   存储在数据库中,经过第一次编译后再次调用不需要再次编译(效率比较高)，用户通过指定存储过程的名字并给出参数(如果该存储过程带有参数)来执行它。存储过程就是数据库中的一个重要对象(针对sql编程而言)。

   存储过程：简称过程

   .与函数的区别

   .相同点

    1、存储过程和函数目的都是为了可重复执行操作数据的SQL语句的集合
    2、存储过程函数都是一次编译，后续执行

   .不同点
   1、标识符不同。 函数的标识符为：function，过程为：procedure。
   2、函数中有返回值，且必须返回，而过程中没有返回值
   3、 过程无返回值类型，不能将结果直接赋值给变量;函数有返回值类型，调用时，除了在select，中，必须将返回值赋值给变量。
   4、函数可以通过select语句中使用,而过程不能:函数是使用select调用,过程不是




   .存储过程操作
    
   .创建过程

   基本语法
   Create procedure 过程名字([参数列表])
   Begin
     过程体
     End
   结束符

  如果过程体中只有一条指令，那么可以省略begin 和end

  -- 创建过程

  create procedure my_pro()
  select * from child; //创建过程完成
-- 修改语句结束符
delimiter $$ 
create procedure my_pro2()
begin 
--求1-100和
 declare i int default 1;
 declare sum int default 0; -- 局部变量
 set @sum = 0 --会话变量
 -- 开始循环获取结果
 while i <  = 100 do 
           -- 求和
           set @sum = @sum +i
           set i = i+1;
end while;
-- 显示结果
select @sum;
end 
$$
delimiter ;

.查看过程
查看过程与查看函数完全一样:除了关键字

查看全部存储过程:show procedures status[like \`pattern\`];

查看过程创建语句: show create procedure 过程名字;

.调用过程

过程：没有返回值，select 不可能调用

调用过程有专门的语法: call 过程名(参数列表)

. 删除过程
 基本语法： drop procedure 过程名字;


 .存储过程的形参类型
  存储过程也允许提供参数(形参和实参):存储的参数也和函数一样,需要指定其类型。 但是存储过程对参数还有额外的要求;自己的参数分类

  .In 
   表示参数从外部传入到里面使用(过程内部使用):可以是直接数据也可以是保存数据的变量

   .Out

   表示参数是从过程里面把数据保存到变量中,然后交给外部使用：传入的必须是变量
   如果说传入的out变量本身在外部有数据,那么进入过程之后，第一件事就是被清空,设为NULL

   .Inout
   数据可以从外部传入到过程内部使用，同时内部操作之后，又会将数据返还给外部.


   参数使用级别语法(形参)
   过程类型 变量名 数据类型;//In Int_1 Int

   实例

   -- 创建三个外部变量
   set@n1 = 1;
   set@n2 = 2;
   set@n3 = 3;

   -- 创建过程

   -- 修改语句结束符
   delimiter $$
   create procedure my_pro3(in int_1 int , out int_2 int , inout int_3 int)
   begin
    --查看三个传入进来的数据的值
    select int_1,int_2,int_3;
    -- 修改三个变量的值
    set int_1 = 10;
    set int_2 = 100;
    set int_3 = 1000;

     select int_1,int_2,int_3;

     -- 查看会话变量

     select @n1,@n2,@3;

     -- 修改会话变量
     set @n1 = 'a';
     set @n2 = 'b';
     set @n3 = 'c';

    select @n1,@n2,@3;

   end

   $$

   delimiter ;

   --调用过程 

   call my_pro3(@n1,@n2,@n3)


   分析结果：out类型的数据会被清空,其他正常


   在执行过程之后，再次查看会话变量(外部)

   .触发器 

   .触发器概念
   触发器是一种特殊类型的存储过程，它不同于我们前面介绍过的存储过程。
   触发器主要是通过事件进行触发而被执行的，而存储过程可以直接调用
   触发器：trigger, 是一种非常接近于js中的事件的知识。提前给某张表的所有记录(行)绑定
   一段代码，如果该行的操作满足条件(触发),这段提前准备好的代码会自动执行。


   。 作用

   1、可以在写入数据表前，强制校验或转换数据 (保证数据安全)
   2、触发器发生错误时，异动的结果会被撤销(如果触发器执行错误,那么前面用户已经执行成功的操作，就会被撤销：事物安全)
   3、部分数据管理系统可以针对数据定义语言（DDL）使用触发器，称为DDL触发器
   4、可依照特定的情况,替换异动的指令(instead of)。(mysql 不支持)

   。 触发器优缺点
   
   .优点

   1、触发器可通过数据库中的相关表实现联级更改。(如果某张表的数据改变，可以利用触发器来实现其他表的无痕操作【用户不知道】)

   .缺点
   1、对触发器过分的依赖，势必影响数据库的结构，同时增加了维护的复杂度。
   2、造成数据在程序层面不可控 (PHP层)



   .触发器基本语法

    .创建触发器
     . 基本语法
    
Create trigger 触发器名字 触发时机 触发事件 on 表 for each row
Begin 

Emd

触发对象: on 表 for each row , 触发器绑定实质是表中的所有行，因此每一行发生指定的改变的时候，就会触发触发器

.触发时机

触发时机：每张表中对应的行都会有不同的状态,当sql指令发生的时候,都会令行中数据发生改变,每一行总会有相同两种状态：数据操作前和操作后

Before： 在表中数据发生改变前的状态
After: 在表中数据已经发生改变后的状态

.触发事件
 
 触发事件:mysql中触发器针对的目标是数据发生改变，对应的操作只有写作操作(增删改)

 insert:插入操作

 update:更新操作

 Delete: 删除操作

 .注意事项
  一张表中，每一个触发时机绑定的触发事件对应的触发器类型只能有一个, 
  一张表中只能有一个对应after insert 创触发器,
  因此,一张表中最多的触发器只能有6个:before insert, before update, before delete ,after insert , ager update ,after delete


  .查看触发器

  需求:有两张表,一张是商品表,一张是订单表(保留商品ID),每次订单生成,商品表中对应的库存就应该发生变化.

  1、 创建两张表:商品表和订单表
  2、 创建触发器:如果订单表发生数据插入，那么对应的商品就应该减少库存
  Create 名字 after insert on my orders for each row


 --创建触发器
 delimiter $$
 create trigger after_insert_order_t after insert on my_orders for each row
 begin 
 --更新商品库存
 update my_goods set inv = inv - new.goods_num  where id = new.goods_id;
 end 
 $$
 delimiter ;

 --删除触发器
  drop trigger after_inser_order_t;

-- 自动扣除商品库存的触发器
delimiter $$
create table a_i_o_t insert on my_order for each row 
begin
        -- 更新商品库存:new 代表着新增的订单
        update my_goods set inv = inv - new.goods_num where id = new.goods_id;
end 
$$
delimiter ;
.查看触发器

 1、 查看全部触发器
 Show triggers;
 2、查看触发器的创建语句
 Show create trigger 触发器名字;

 3、 触发触发器
  想办法让触发器执行:让触发器指定的表中,对应的时机发生对应的操作即可。
  1、表为my_orders
  2、在插入之后
  3、插入操作
  列：insert into my_orders value(null,1,1);

4、删除触发器
基本语法: drop trigger after_inser_order_t;
```

create table my_goods(

id int primary key   auto_increment,

name varchar(20) not null,

inv int

) charset utf8;

create table my_orders(

id int primary key auto_increment,

goods_id int not null,

goods_num int not null

)charset utf8;

insert into my_goods values(null,'手机',1000),(null,'电脑',5000),(null,'游戏机',100);

```

.触发器应用

. 记录关键字: new 、old
触发器针对的是数据表中的每条记录(每行),每行在数据操作前后都有一个对应的状态
触发器在执行前就讲对应的状态获取到了,将没有操作之前的状态(数据)都保存到old关键字中,
而操作后的状态都放在new中。

在触发器中，可以通过old和new来获取绑定表中对应的记录数据。
基本语法：关键字。字段名
old和new 并不是所有触发器都有:
insert: 插入前全为空,没有old
Delete:清空数据,没有new

. 商品自动扣除库存
delimiter $$
create table a_i_o_t insert on my_order for each row 
begin
        -- 更新商品库存:new 代表着新增的订单
      update my_goods set inv = inv - new.goods_num where id = new.goods_id;
end 
$$
delimiter ;

-- 判读库存
delimiter $$
create trigger b_i_o_t before insert on my_orders for each row
begin
-- 取出存储数据进行判断
select inv from my_goods where id = new.goods_id into @inv;
-- 判断  
if @inv < new.goods_num then 
          -- 中段操作 :暴力解决，主动出错
          insert into xxx value ('xxx');
end if ;
end 
$$
delimiter ;
```
