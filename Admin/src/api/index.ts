import { api } from "./require"
import qs from "qs"

interface T {
     result: object[],
     code: number,
     message: string
}

export const getBanner = async () => {
     let response = await api.post<T>("admin/getbanner")
     console.log(response)
     return {result: response.data.result, message:response.data.message }
}

export const wallhaven = async (params: wallhaven) => {
     let { data } = await api.get("admin/wallhaven", {
          params
     });
     return data.result;
};

export const appendBanner: appendBanner = async (bannerIds) => {
     let form = new FormData()
     form.append("bannerIds", JSON.stringify(bannerIds))
     return await api.post("admin/appendbanner", form, {
          headers: {
               "Content-Type": 'application/x-www-form-urlencoded'
          }
     });
};

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
