<template>
    <el-row v-loading="state.tableLoading">
        <el-row class="input-container">
            <p class="input-text">
                <span>搜索名字</span>
            </p>
            <el-input class="input" placeholder="文章名字" />
            <p class="input-text">
                <span>文章分类</span>
            </p>
            <el-select class="input" v-model="value" placeholder="Select" size="large">
                <el-option
                    v-for="item in options"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                ></el-option>
            </el-select>
            <el-button size="small" type="primary">搜索</el-button>
            <el-button size="small" type="default">重置</el-button>
        </el-row>
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
import { reactive, onMounted,ref } from 'vue'
import paginationVue from '../../components/pagination.vue'
import * as api from "../../api/index"
import { qsTime } from "../../util"
// QueryAllTag
const state = reactive({
    articles: [],
    center: {
        "text-align": "center"
    },
    pageSize: 10,
    count: 0,
    tableLoading: false,

})
const value = ref('')

const options = [
  {
    value: 'Option1',
    label: 'Option1',
  },
  {
    value: 'Option2',
    label: 'Option2',
  },
  {
    value: 'Option3',
    label: 'Option3',
  },
  {
    value: 'Option4',
    label: 'Option4',
  },
  {
    value: 'Option5',
    label: 'Option5',
  },
]



const getArticles = async (current: number = 1) => {
    state.tableLoading = !state.tableLoading
    let { data }: any = await api.getArticles({
        limit: state.pageSize,
        offset: current
    })
    state.count = data.count
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

<style lang="css" scoped>
.input-container {
    @apply flex items-center justify-center;
}
.input {
    @apply w-56 mx-6;
}

.input-text {
    @apply flex items-center  text-sm;
 
}

</style>