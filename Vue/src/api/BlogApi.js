import { api } from "./index";
import axios, { qs } from "axios";
export const getComments = async articleId => {
    let { data: comment } = await api.get("v1/comment", {
        params: { articleId }
    });
    // 如果有评论
    if (Array.isArray(comment.result)) {
        return comment.result.map(father => {
            if (Array.isArray(father.Replys)) {
                father.Replys.map(item => {
                    item.leave = father.UUID;
                    return item;
                });
            }
            return father;
        });
    } else {
        return []
    }
};
export const comment = async (articleId, content, replyArticle, userName) => {
    console.log(articleId, content, replyArticle, userName);
    return await api.post(
        "v1/comment",
        qs.stringify({
            articleId,
            content,
            replyArticle,
            userName
        })
    );
};

export const shouldJson = async (articleId, content) => {
    return await api.post(
        "v1/should",
        qs.stringify({
            articleId,
            content
        })
    );
};

export const generateJSON = async () => {
    try {
        let data = (await api.get("v1/returnjson")).data;
        return JSON.parse(data.json);
    } catch (e) {
        return false;
    }
};

export const wallhaven = async params => {
    try {
        let { data } = await api.get("v1/wallhaven", {
            params: {
                limit: 216,
                offset: 1
            }
        });
        return data.result;
    } catch (e) {
        return false;
    }
};
export const GetUrl = async () => {
    try {
        return await axios.get("https://pixabay.com/api", {
            params: {
                key: "21226858-57a14a3bedc89005c85e668cc",
                per_page: 200,
                //    q:"",
                category: "nature",

                safesearch: true
            }
        });
    } catch (e) {
        return false;
    }
};
export const getOauthInfo = async userInfo => {
    let url;
    const { access_token, scope, origin } = userInfo;
    origin == "gitee" ? (url = "https://gitee.com/api/v5/user") : (url = "");
    return await api.get(url, {
        params: {
            access_token
        }
    });
};
