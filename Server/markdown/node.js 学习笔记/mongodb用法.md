# mongodb用法

```
mongo  //进入mongodb 控制台
help  //帮助
exit // 退出
show dbs //显示所有数据库
use test //进入数据库
db.createCollection('test') //创建集合
show collections //显示所有集合
show dbs 
db.stats() //显示数据库属性
db.dropDatabase() //删除数据库 不写名字表示删除当前数据库
db.goods.renameCollection("good") //修改集合名字
db.good.drop()//删除这个集合
db.good.insert({'name':1})//插入集合
db.good.insertOne({"a": 3})//插入一条
db.good.insertMany([{"b": 3}, {'c': 4}])//插入多条数据
db.good.find()//查找文档 
(>) 大于 - $gt
(<) 小于 - $lt
(>=) 大于等于 - $gte
(<= ) 小于等于 - $lte
$eq (==)  $ne (!=)
db.good.distinct("type") //应该属于查找分类的 查找type 在集合里面有多少种值 
db.good.find({age:{$gt:4}}) 
db.goods.find({ORI_PRICE:{$gt:50},GOOD_TYPE:1})//and用法
db.goods.find({$or:[{name:'小明'},{age:{$gt:5}}]}) //or
db.goods.find({name:'小明',$or[{name:'小红'},{age:{$gt:50}}]}) 混合使用
db.goods.find({name:'小明'},{age:0}) 排除模式查询
db.good.count()//计算出集合中的数量
db.good.remove()//删除集合中的
db.good.deleteMany({ name : "A" })//删除全部 条件相等的文档
db.good.deleteOne({})//删除一条
skip(), limilt(), sort()三个放在一起执行的时候，执行的顺序是先 sort(), 然后是 skip()，最后是显示的 limit()。 limit 显示数量 skip跳过几条数量 sort({age:-1})排序
db.good.updateOne({name:'小明'},{$set:{age:66}) //查找条件name等于小明修改age这个属性
db.good.updateMany({name:'小明'},{$set:{age:66}) //查找条件name等于小明修改age这个属性
db.good.remove({'name':'李四'},{justOne:true})
db.good.deleteOne()
db.good.deleteMany()
db.stus.aggregate([{$group:{_id:'$GOOD_TYPE',age:{$sum:1}}}])//$sum代表数量的统计{ "_id" : 0, "name" : 325 } { "_id" : 1, "name" : 309 } 表示GOOD_TYPE有两种类型0和1 后面是数量表示0的商品有325
db,stus.createrIndex()
$inc:增加
db.goods.pudate({name:'小明'},{$inc:{rank:1}}
$mul:相乘
db.goods.pudate({name:'小明'},{$mul:{rank:1}}
$rename:改名
db.goods.pudate({name:'小明'},{$rename:{"name":"name2"}}
$set:新增或修改
db.goods.pudate({name:'小明'},{$set:{"papa":"爱你哦"}}
$unset:字段删除
db.goods.pudate({name:'小明'},{$unset:{"papa":"爱你哦"}}
mongodump -h 127.0.0.1(主机) -d test(数据库名) -o G:\mongodbData(存放的地方) //导出数据
mongorestore -h 127.0.0.1 -d dbname --dir G:\mongodbData\test //导入后面要加上以前导出来的名字

mongoimport --db 数据库名字 --collection 集合名 --file 文件路径

order.aggregate([
    {
        $match: { name: '母亲' } //搜索条件
    },
    {
        $lookup: {
            from: 'item', //需要查找的子collestions
            localField: 'parent_id', //父id
            foreignField: 'child_id',//子id
            as: 'items'//返回的列表名称
        }
    }, { $project: { items: 1 } }], function (err, data) {
        console.log(err)
        console.log(data[0].items) //在这里可以打印出
    })
populate//查找关联
```
