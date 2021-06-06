// 使用 Mock
import Mock, { Random } from 'mockjs'


// 返回token 
Mock.mock('api/home', 'get',
    {
        "pictures|5": [
            Random.image('815x300', '#4A7BF7', 'Brea Bennett')
        ],
        "recommend": {
            "pictures": Random.image('400x230', '#F5E4D0', 'BiBi Jones'),
            "tag": ["面试", "小程序", "瞎聊", "热门推荐"],
        },
        "articles|30": [{
            "id|1-1000": 500,
            "picture": Random.image('169x140', '#4A7BF7', 'Gabriella Fox'),
            "title": Random.cword(15),
            "paragraph": Random.cword(40),
            "content": Random.cword(200),
            "time": Random.now('yyyy-MM-dd HH:mm'),
            "count|1-100": 100,
            "tager": "Vue"
        }],
        "applet": Random.image('400x230', '#F5E4D0', 'ashlynn brooke'),
        "revert|5": [{
            "userName": Random.cword(3, 7),
            "content": Random.cword(20),
            "time": Random.now('yyyy-MM-dd HH:mm'),
            "id|1-1000": 120,
        }],
        "friendsLink|21": [{
            "userName": Random.cword(3, 7),
            "id|1-1000": 120,
            "link": ""
        }]
    });
Mock.mock('api/gossip','get',{
    "articles|9": [{
        "id|1-1000": 500,
        "picture": Random.image('400x300', '#4A7BF7', 'Gabriella Fox'),
        "title": Random.cword(15),
        "paragraph": Random.cword(40),
        "content": Random.cword(200),
        "time": Random.now('yyyy-MM-dd HH:mm'),
        "count|1-100": 100,
        "tager": "Vue"
    }],
})
export { Mock }

