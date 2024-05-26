# mongodb

## 单例

### 准备docker-compose.yaml

```yaml
version: '3.8'

services:
  mongo6:
    image: mongodb/mongodb-community-server:6.0.15-ubi8
    container_name: mongo6
    environment:
      - TZ=Asia/Shanghai
      - MONGODB_INITDB_ROOT_USERNAME=zeng
      - MONGODB_INITDB_ROOT_PASSWORD=zeng
    ports:
      - "27017:27017"
    volumes:
      - ./data/db:/data/db
    networks:
      - mongo_network

networks:
  mongo_network:
    driver: bridge
```

### 启动容器

```bash
docker-compose up -d
```

## 数据类型

1. **数值型：**
    * Double：双精度浮点数。
    * Int：32位有符号整数。
    * Long：64位有符号整数。
    * Decimal：用于存储需要高精度计算的数值，如货币数据。
2. **日期型：**
    * Date：日期和时间。
    * Timestamp：时间戳，用于MongoDB内部，也可以用于记录文档的修改或创建时间。
3. **字符串型：**
    * String：用于存储文本数据。
    * ObjectId：一个特殊的字符串，作为文档的唯一标识符。
4. **布尔型：**
    * Boolean：用于存储布尔值（true或false）。
5. **数组型：**
    * Array：用于存储多个值，可以是任意类型的数组。
6. **对象型：**
    * Object（或称为内嵌文档）：用于存储嵌套的对象或文档。
7. **二进制数据型：**
    * Binary Data：用于存储二进制数据。
8. **Null类型：**
    * Null：用于表示空值或不存在的字段。
9. **其他类型：**
    * Regular Expression：用于存储正则表达式，以便在文本搜索时使用。
    * Code：用于存储JavaScript代码片段，通常与$where查询操作符一起使用。
    * Symbol：已废弃，在MongoDB 3.0之后的版本中不再支持。
    * MinKey 和 MaxKey：这两个特殊的键分别用于表示BSON类型的最小值和最大值，主要用于创建索引时的边界值。

## 索引类型

1. 单键索引（Single Field Index）：
    * 定义：只对一个字段进行索引。
    * 特点：可以加速针对该字段的查询，是MongoDB中最简单的索引类型。
    * 示例：db.collection.createIndex({name:1}) 将会为 name 字段建立单键索引。
2. 复合索引（Compound Index）：
    * 定义：对多个字段进行索引。
    * 特点：可以加速针对这些字段的查询，且复合索引的字段顺序和索引方向对性能有重要影响。
    * 示例：db.collection.createIndex({name:1, age:-1}) 将会为 name 和 age 两个字段建立复合索引。
3. 多键索引（Multikey Index）：
    * 定义：对数组或嵌套文档的字段进行索引，即会为数组中的每个元素创建索引项。
    * 特点：适用于属性包含数组数据的情况。
    * 支持的数据类型：strings, numbers和nested documents。
4. 地理空间索引（Geospatial Index）：
    * 定义：对包含地理坐标的字段进行索引。
    * 类型：
        - 2dsphere索引：用于存储和查找球面上的点。
        - 2d索引：用于存储和查找平面上的点。
    * 示例：db.collection.createIndex({location:"2dsphere"}) 将会为 location 字段建立地理空间索引。
5. 全文索引（Text Index）：
    * 定义：针对文本字段建立索引，支持文本搜索。
    * 特点：使用全文搜索技术，支持模糊查询、分词搜索等功能。
    * 示例：db.collection.createIndex({content:"text"}) 将会为 content 字段建立文本索引。
6. 哈希索引（Hashed Index）：
    * 定义：将字段值进行哈希处理后进行索引。
    * 特点：在某些特定场景下，如分布式系统中，哈希索引可以提高查询效率。
7. 联合唯一索引（Unique Index）：
    * 定义：对多个字段进行索引，并保证组合值的唯一性。
    * 特点：确保索引字段的组合值在集合中是唯一的。
8. 稀疏索引（Sparse Index）：
    * 定义：只对存在索引字段的文档进行索引，对不存在索引字段的文档不进行索引。
    * 特点：可以节省空间并提高某些查询的性能。
9. TTL索引（TTL Indexes）：
    * 定义：一种特殊的单键索引，支持文档在一定时间之后自动过期删除。
    * 特点：字段类型必须是日期类型，且只能在单字段上建立。
10. 部分索引（Partial Indexes）：
    * 定义：只针对满足特定条件的文档进行索引。
    * 特点：可以提高索引效率，并减少不必要的索引空间。

## 操作用法

### 创建数据库

```js
use
grade
```

### 创建集合

```js
db.createCollection("class")
```

### 插入数据

```js
db.class.insert({name: "xiaozhang", age: 6, sex: "m", hobby: ["draw", "pingpong", "dance"]});
db.class.insert({name: 'xiaoxiang', age: 9, sex: 'm', hobby: ['draw', 'basketball']});
db.class.insert({name: 'xiaowu', age: 10, sex: 'm', hobby: ['sing', 'basketball']});
db.class.insert({name: 'xiaoliu', age: 11, sex: 'm', hobby: ['footboll', 'dance']});
db.class.insert({name: 'xiaozeng', age: 4, sex: 'f', hobby: ['footboll', 'dance', "sing"]});
db.class.insert({name: 'xiaohe', age: 8, sex: 'f', hobby: ['sing', 'dance']});
```

### 查询

### 常用操作符

#### 比较操作符

1. **$eq:** 等于
2. **$ne:** 不等于
3. **$gt:** 大于
4. **$gte:** 大于等于
5. **$lt:** 小于
6. **$lte:** 小于等于

#### 逻辑操作符

* **$and:** 且，多个条件同时满足
* **$or:** 或者，多个条件中满足一种一个即可
* **$not:** 非，匹配不满足的条件
* **$nor:** 非，匹配所有条件都不满足的条件

#### 元素操作符

* **$exists:** 匹配存在某个字段的文档
* **$type:** 匹配字段类型符合指定类型的文档

#### 数组操作符

* **$in:** 指定字段中包含数组的中至少一个值
* **$nin:** 指定字段中不包含数组的中任何一个值
* **$all:** 指定字段中包含数组的中所有的值
* **$size:** 指定字段的数组长度等于指定的值

#### 正则表达操作符

* **$regex:** 指定字段符合正则表达式

#### 文本搜索操作符

* **$text:** 执行全文本搜索
* **$meta:** 获取文本搜索得分

## 例子

### 查询

```js
// 获取class的所有数据
db.class.find();
// 获取class的所有数据，并只显示name，age，sex字段
db.class.find({}, {name: 1, age: 1, sex: 1, _id: 0});
// 获取age为8的数据
db.class.find({age: 8});
// 获取年龄大于10的学生名字,不需要id时需要显示配置
db.class.find({age: {"$gt": 10}}, {name: 1, _id: 0});
// 获取4到8岁学生信息
db.class.find({age: {"$gte": 4, "$lte": 8}});
// 查询6岁男生
db.class.find({age: 6, sex: "m"});
// 查询年龄小于7岁或者大于10岁的学生
db.class.find({"$or": [{age: {"$lt": 7}}, {age: {"$lt": 10}}]});
// 查询8岁或者11岁的学生
db.class.find({"$or": [{age: 8}, {age: 11}]});
// 查询有3项爱好的学生
db.class.find({hobby: {"$size": 3}});
// 查询爱好中有sing或者dance的学生
db.class.find({hobby: {"$in": ["sing", "dance"]}});
// 查询爱好中同时有sing与dance的学生
db.class.find({hobby: {"$all": ["sing", "dance"]}});
// 根据性别与年龄排序
db.class.find().sort({sex: -1, age: 1});
// 获取年龄最大的两个男生
db.class.find({sex: "m"}).sort({age: -1}).limit(2);
// 获取年龄第二小的女生
db.class.find({sex: "f"}).sort({age: 1}).skip(1).limit(1);
// 获取第一个女生
db.class.findOne({sex: "f"});
// 获取学生的所有兴趣爱好
db.class.distinct("hobby");
// 获取女学生的所有兴趣爱好
db.class.distinct("hobby", {sex: "f"});
```

### 数量获取

```js
// 获取男生数量
db.class.count({sex: "m"});
```

### 删除

```js
// 批量删除
db.class.deleteMany({age: {"$gt": 14}});
// 删除一条数据
db.class.deleteOne({sex: "f"});
```

### 修改

```js
// 修改xiaoxiang的爱好
db.class.update({name: "xiaoxiang"}, {"$set": {age: 8, hobby: ["draw", "dance"]}});
// 增加爱好,push应该是单个元素，如果使用数组的话会把数组作为一个元素加入到原先的数组中
db.class.update({name: "xiaoxiang"}, {"$push": {"hobby": "sing"}});
db.class.update({name: "xiaoxiang"}, {"$push": {"hobby": ["citywalk", "ride bike"]}});
// 新增一条属性email
db.class.update({name: "xiaoxiang"}, {"$set": {email: "xiao@163.com"}});
// 删除一条属性email
db.class.update({name: "xiaoxiang"}, {"$unset": {email: null}});
// 删除数组中的第一个元素
db.class.update({name: "xiaoxiang"}, {"$pop": {"hobby": -1}});
// 删除数组中的指定元素
db.class.update({name: "xiaoxiang"}, {"$pull": {hobby: "sing"}});
// 删除数组中的指定的多个元素
db.class.update({name: "xiaoxiang"}, {"$pullAll": {hobby: ["sing", "ride bike"]}});
```

## 新增字段

```js
db.class.update({name: "xiaozhang"}, {"$set": {score: {math: 75, english: 84, chinese: 92}}});
db.class.update({name: "xiaoxiang"}, {"$set": {score: {math: 88, english: 79, chinese: 83}}});
db.class.update({name: "xiaowu"}, {"$set": {score: {math: 59, english: 47, chinese: 63}}});
db.class.update({name: "xiaoliu"}, {"$set": {score: {math: 79, english: 73, chinese: 69}}});
db.class.update({name: "xiaozeng"}, {"$set": {score: {math: 96, english: 89, chinese: 94}}});
db.class.update({name: "xiaohe"}, {"$set": {score: {math: 87, english: 81, chinese: 82}}});
```

## 聚合函数

```js
// 按照性别统计人数
db.class.aggregate([{"$group": {_id: "$sex", total: {"$sum": 1}}}]);
// 按照年龄进行分组，并筛选出年龄相同的人
db.class.aggregate([{"$group": {_id: "$age", total: {"$sum": 1}}}, {"$match": {"total": {"$gt": 1}}}]);
// 统计每名男生的语文成绩
db.class.aggregate([{"$match": {sex: "m"}}, {"$project": {_id: 0, name: 1, "score.chinese": 1}}]);
// 统计女生的英语成绩，并按照倒序排列
db.class.aggregate([{"$match": {sex: "f"}}, {
    "$project": {
        _id: 0,
        name: 1,
        english_score: "$score.english"
    }
}, {"$sort": {english_score: -1}}]);
```