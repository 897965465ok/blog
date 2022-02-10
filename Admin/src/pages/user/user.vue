<template>
    <el-row>
        <el-table
            :data="state.articles"
            :cell-style="(state.center as any)"
            brother
            class="w-full min-h-full"
        >
            <el-table-column prop="CreatedAt" label="时间"></el-table-column>
            <el-table-column prop="ID" label="用户ID"></el-table-column>
            <el-table-column prop="name" label="用户名"></el-table-column>
        </el-table>
        <el-row>
            <!-- <paginationVue></paginationVue> -->
        </el-row>
    </el-row>
</template>
<script  setup lang='ts'>
import { reactive, defineProps, onMounted } from 'vue'
import * as api from "../../api/index"
import paginationVue from '../../components/pagination.vue'
const state = reactive({
    articles: [],
    center: {
        "  text-align": "center"
    }
})
onMounted(async () => {
    let { data } = await api.getUser({
        limit: 1,
        offset: 0
    })
    console.log(data)
    state.articles = data.result
})
</script>