const initState = {
    HotList: [],
    favoriteList: [],
    token: '',
    userInfo: {},
    bannerUrl: [],
    ALLarticle: [],
    tagerList: []
}

export default (state = initState, action) => {
    switch (action.type) {
        // 设置 文件列表
        case 'SET-ARTICLES': {
            return Object.assign({}, state, action)
        }
        default: {
            return state
        }
    }
}