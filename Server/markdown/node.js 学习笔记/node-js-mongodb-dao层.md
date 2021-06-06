# node-js-mongodb-dao 层

好久没写东西了,最近一直在忙，这一年有太多不如意，改学的没到，目标太遥远了，努力吧~~，不扯了，进入正题。

```js
const MongoClient = require('mongodb').MongoClient;
function _connerctDB(callback){
	const url = 'mongodb://localhost:27017';
 MongoClient.connect(url,(err,data)=>{//连接数据库
 	if (!err){console.log('连接数据库成功')};
 	callback(data)
 })
}//增加
exports.insertOne=function(json,collectionName,callback){//那个集合，存入的对象
	_connerctDB((data)=>{
 	const db = data.db('user');//获取数据库
 	const collection = db.collection(collectionName);//获取集合
 	collection.insertOne(json,(err, result)=>{
    callback(err,result);//回调
    data.close();
});
 })
}//查找
exports.findAllducment=function(json,collectionName,callback){
	_connerctDB((data)=>{
 	const db = data.db('user');//获取数据库
 	const collection = db.collection(collectionName);//获取集合
 	collection.find(json).skip(20).limit(20).toArray((err, docs)=>{//查找数据库以数组的形式返回,skip(20)和limit(20)做分页的时候用到。。
 		callback(err,docs);
 		data.close();
 	})
 })
} //删除
exports.removeDocument=function(json,collectionName,callback){
	_connerctDB((data)=>{
 	const db = data.db('user');//获取数据库
 	const collection = db.collection(collectionName);//获取集合
 	collection.deleteMany(json,(err, result)=>{
 		callback(err,result);
 		data.close();
 	})
 })
}//改
exports.updateData = function(json, jsonpudate,collectionName,callback){
    //连接到表
    _connerctDB((data)=>{
    const db = data.db('user');//获取数据库
 	const collection = db.collection(collectionName);//获取集合
    //更新数据
    let whereStr = json;//原本的参数
    let updateStr = {$set:jsonpudate};//需要修改的参数
    collection.update(whereStr,updateStr,(err, result)=>{
    	if(err)
    	{
    		console.log('Error:'+ err);
    		return;
    	}
    	callback(err,result);
    	data.close();
    });
})
}
```
