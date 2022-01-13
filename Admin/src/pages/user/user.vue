<template>
    <el-row>
    <el-table :data="state.articles"  :cell-style="(state.center as any)"  brother class="w-full min-h-full">
        <el-table-column prop="CreatedAt" label="创建时间" width="150"></el-table-column>
        <el-table-column prop="ID" label="文章ID"  ></el-table-column>
        <el-table-column prop="name" label="文章名字"  ></el-table-column>
        <el-table-column prop="like" label="点赞"  ></el-table-column>
        <el-table-column prop="tag" label="分类"  ></el-table-column>
        <el-table-column prop="whatch_number" label="查看次数"  ></el-table-column>
    </el-table>
    <el-row>
    <paginationVue></paginationVue>
    </el-row>
    </el-row>
</template>
<script  setup lang='ts'>
import { reactive, defineProps, onMounted } from 'vue'
import * as api from "../../api/index"
import paginationVue from '../../components/pagination.vue'
const state = reactive({
    articles: [],
    center:{
      "  text-align": "center"
    }
})
onMounted(async () => {
    let { data } = await api.getArticles({
        limit: 10,
        offset: 1
    })
    state.articles = data.result
})
</script>