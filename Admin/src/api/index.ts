import { api } from "./require"
import qs from "qs"

interface T {
     result: object[],
     code: number,
     message: string,
     count: number
}

export const getBanner = async (params: paging) => {
     let response = await api.post<T>("admin/getbanner", null, {
          params
     })
     console.log(response)
     return { result: response.data.result, message: response.data.message, count: response.data.count }
}

export const wallhaven = async (params: wallhaven) => {
     let { data } = await api.get("admin/wallhaven", {
          params
     });
     console.log(data)
     return data.result;
};

export const bannerPlush: appendBanner = async (bannerIds) => {
     let form = new FormData()
     form.append("bannerIds", JSON.stringify(bannerIds))
     return await api.post("admin/appendbanner", form, {
          headers: {
               "Content-Type": 'application/x-www-form-urlencoded'
          }
     });
};

export const deleteBanner = async (ID: string) => {
     let { data } = await api.post<
          Promise<{
               code: number,
               message: string
          }>
     >("admin/deletebanner", null, {
          params: {
               id: ID
          }
     });
     return data
}

export const getArticles = async (params: paging) => {
     return await await api.get("/v1/articles", {
          params
     });
}

export const getUser = async (params: paging) => {
     return await await api.post("/admin/queryuser", null,{
          params
     });
}




// export const getOauthInfo = async userInfo => {
//      let url;
//      const { access_token, scope, origin } = userInfo;
//      origin == "gitee" ? (url = "https://gitee.com/api/v5/user") : (url = "");
//      return await api.get(url, {
//           params: {
//                access_token
//           }
//      });
// };
