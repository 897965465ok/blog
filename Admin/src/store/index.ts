import { createStore } from "vuex"
import * as api from "../api/index"
import { qsTime } from "../util/index"
import { ElMessage } from 'element-plus'
export const store = createStore({
    state() {
        return {
            Banner: [],
            BannerLenght: 0,
            wallhavenPicture: [],
            wallhavenLength: 0,
            tableLoading: false,
            loading: false,
            toDelete: [],
            toDeleteLength: 0
        }
    },
    mutations: {
        LOAD_BANNER(state: any, data) {
            let { count, result } = data;
            state.BannerLenght = count;
            state.Banner = result.map((item: any) => {
                item.CreatedAt = qsTime(item.CreatedAt)
                return item
            })
        },
        DELETE_BANNER(state: any, ID) {
            state.Banner.forEach((item: any, index: number) => {
                if (item.ID == ID) {
                    state.Banner.splice(index, 1)
                }
            })
            ElMessage.success("成功")
        },
        WALLHAVEN(state: any, { count, imgs }) {
            state.wallhavenPicture = imgs
            state.wallhavenLength = count
        },
        TO_DELETE(state, data) {
            let { count, result } = data;
            state.toDeleteLength = count;
            state.toDelete = result.map((item: any) => {
                item.CreatedAt = qsTime(item.CreatedAt)
                item.uuid = item.id
                return item
            })
        }
    },
    actions: {
        loadBanner(context, { limit, offset }) {
            context.state.tableLoading = !context.state.tableLoading
            api.getBanner({
                limit: limit,
                offset: offset
            }).then(data => {
                context.commit('LOAD_BANNER', data)
                context.state.tableLoading = !context.state.tableLoading
            })
        },
        toDelete(context, { limit, offset }) {
            context.state.tableLoading = !context.state.tableLoading
            api.getBanner({
                limit: limit,
                offset: offset
            }).then(data => {
                context.commit('TO_DELETE', data)
                context.state.tableLoading = !context.state.tableLoading
            })
        },
        deleteBanner(context, ID: string) {
            api.deleteBanner(ID).then(() => {
                context.commit('DELETE_BANNER', ID)
            })
        },
        wallhaven(context, { limit, offset }) {
            context.state.loading = !context.state.loading
            api.wallhaven({
                limit: limit, //12
                offset: offset // 1
            }).then(({ count, imgs }) => {
                context.commit('WALLHAVEN', { count, imgs })
                context.state.loading = !context.state.loading
            })
        }

    },
    getters: {}
})