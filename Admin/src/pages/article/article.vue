<template>
    <el-row  v-loading="state.tableLoading">
        <el-table
           
            style="min-height:80%;"
            :data="state.articles"
            :cell-style="(state.center as any)"
            brother
        >
            <el-table-column prop="CreatedAt" label="创建时间" width="150"></el-table-column>
            <el-table-column prop="ID" label="文章ID"></el-table-column>
            <el-table-column prop="name" label="文章名字"></el-table-column>
            <el-table-column prop="like" label="点赞"></el-table-column>
            <el-table-column prop="tag" label="分类"></el-table-column>
            <el-table-column prop="whatch_number" label="查看次数"></el-table-column>
        </el-table>
        <paginationVue
            @loadTime="getArticles"
            :count="state.count"
            :pageSize="state.pageSize"
            callback="loadTime"
        ></paginationVue>
    </el-row>
</template>
<script  setup lang='ts'>
import { reactive, onMounted } from 'vue'
import paginationVue from '../../components/pagination.vue'
import * as api from "../../api/index"
import { qsTime } from "../../util"
const state = reactive({
    articles: [],
    center: {
        "text-align": "center"
    },
    pageSize: 10,
    count: 0,
    tableLoading: false,

})
const getArticles = async (current: number = 1) => {
    state.tableLoading = !state.tableLoading
    let { data }: any = await api.getArticles({
        limit: state.pageSize,
        offset: current
    })
    state.count = data.count
    console.log(data.count)
    state.articles = data.result.map((item: any) => {
        item.CreatedAt = qsTime(item.CreatedAt)
        return item
    })
    state.tableLoading = !state.tableLoading
}
onMounted(() => {
    getArticles()
})
</script>