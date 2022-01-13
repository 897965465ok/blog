<template>
    <!-- <el-row :gutter="5">
        <el-col :offset="6" :span="12">
            <el-carousel>
                <el-carousel-item
                    class="w-full h-full"
                    v-for="(item,index) in tableData "
                    :key="index"
                >
                    <el-image :fit="'fill'" class="w-full h-full" :src="item.large" ></el-image>
                </el-carousel-item>
            </el-carousel>
        </el-col>
    </el-row>-->
    <el-dialog v-model="centerDialogVisible" class="flex flex-col" top="1vh" width="95vw" center>
        <el-row style="min-height:81vh;" v-loading="loading" :gutter="5" @click="changeItme">
            <el-col :span="6" v-for="item in banner" :key="item.uuid">
                <el-image class="w-full cursor-pointer" :uuid="item.uuid" :src="item.large" lazy></el-image>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="24" class="w-full flex justify-center items-center">
                <el-pagination
                    background
                    :page-size="12"
                    :pager-count="12"
                    layout="prev, pager, next"
                    :total="pictures"
                    @current-change="loadImg"
                ></el-pagination>
                <el-button size="small" type="success" @click="appendBanner">点击添加</el-button>
            </el-col>
        </el-row>
    </el-dialog>
    <el-row>
        <el-col>
            <el-button size="small" @click="appendPitrue" type="primary">批量增加</el-button>
            <el-button size="small" :icon="Edit" type="info">修改</el-button>
            <el-button size="small" :icon="Delete" type="danger">删除</el-button>
        </el-col>
    </el-row>
    <el-row class="min-h-full" v-loading="state.tableLoading">
        <el-table :cell-style="(center as any)" :data="tableData">
            <el-table-column prop="ID" label="图片ID" width="80"></el-table-column>
            <el-table-column prop="CreatedAt" label="创建时间" width="150">
                <template v-slot:default="scope">{{ scope.row.CreatedAt }}</template>
            </el-table-column>
            <el-table-column prop="large" label="图片路径" width="180" minHeight="60">
                <template v-slot:default="scope">
                    <el-image :fit="'fill'" class="w-full h-full" :src="scope.row.large"></el-image>
                </template>
            </el-table-column>
            <el-table-column>
                <template v-slot:default="scope">
                    <el-button size="small" :icon="Edit" type="info" circle></el-button>
                    <el-button
                        @click="deleteBanner(scope.row.ID)"
                        type="danger"
                        size="small"
                        :icon="Delete"
                        circle
                    ></el-button>
                </template>
            </el-table-column>
        </el-table>
        <paginationVue
            @loadTime="loadBanner"
            :count="state.count"
            :pageSize="5"
            callback="loadTime"
        ></paginationVue>
    </el-row>
</template> 
<script  setup lang="ts" >
import * as api from "../../api"
import { reactive, toRefs, toRaw, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { qsTime } from "../../util"
import { Edit, Share, Delete, Search, Upload } from '@element-plus/icons-vue'
import paginationVue from "../../components/pagination.vue"

// 表头水平居中：给 el-table 标签绑定 header-cell-style 属性
// <el-table :header-cell-style="{'text-align':'center'}"></el-table>
// 表格内容水平居中：给 el-table 标签绑定 cell-style 属性
// <el-table :cell-style="{textAlign:'center'}"></el-table>

const center = { 'text-align': 'center' }
const state = reactive<dataType>({
    tableData: [],
    centerDialogVisible: false,
    banner: [],
    saveList: [],
    pictures: 0,
    loading: false,
    count: 0,
    tableLoading: false,

})
const loadBanner = async (current: number = 1) => {
    state.tableLoading = !state.tableLoading
    let { result, count } = await api.getBanner({
        limit: 5,
        offset: current
    })
    state.count = count
    state.tableLoading = !state.tableLoading
    state.tableData = result.map((item: any) => {
        item.CreatedAt = qsTime(item.CreatedAt)
        return item
    })

}

onMounted(() => {
    loadBanner()
})

const loadImg = (current: number) => {
    state.loading = !state.loading
    api.wallhaven({
        limit: 12,
        offset: current
    }).then(({ count, imgs }) => {
        state.pictures = count
        state.banner = imgs
        state.loading = !state.loading
    })
}
const appendPitrue = () => {
    state.centerDialogVisible = !state.centerDialogVisible
    state.loading = !state.loading
    api.wallhaven({
        limit: 12,
        offset: 1
    }).then(({ count, imgs }) => {
        state.pictures = count
        state.banner = imgs
        state.loading = !state.loading
    })
}
function changeItme(event: MouseEvent) {
    if (event.target != null && (event.target as HTMLImageElement).nodeName == 'IMG') {
        let uuid = (event.target as HTMLImageElement).getAttribute("uuid")
        let node = state.banner.find((item) => item.uuid == uuid)
        if (node != undefined && state.saveList.findIndex((item) => item.uuid == uuid) == -1) {
            state.saveList.push(node)
            ElMessage.success("成功添加到列表")
        } else {
            ElMessage.warning("重复提交")
        }

    }
}

async function deleteBanner(ID: string) {
    let { message } = await api.deleteBanner(ID)
    state.tableData.forEach((item: any, index: number) => {
        if (item.ID == ID) {
            state.tableData.splice(index, 1)
        }
    })
    ElMessage.success(message)
}

async function appendBanner() {
    let bannerIds = state.saveList.map(item => item.uuid)
    let result = await api.appendBanner(bannerIds)
    ElMessage.success(result.data.message)

}

let { tableData, centerDialogVisible, banner, pictures, loading } = toRefs(state)
</script>
<style scoped>
.el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 150px;
    margin: 0;
    text-align: center;
}

.el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
    background-color: #d3dce6;
}
</style>

