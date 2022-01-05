<template>
    <el-row :gutter="5"  >
        <el-col :offset="2" :span="20" >
            <el-carousel height="320px">
                <el-carousel-item
                    class="w-full h-full"
                    v-for="(item,index) in tableData "
                    :key="index"
                >
                    <el-image :fit="'fill'" class="w-full h-full" :src="item.large"></el-image>
                </el-carousel-item>
            </el-carousel>
        </el-col>
    </el-row>
    <el-table :data="tableData" style="width:100%">
        <el-table-column>
            <template v-slot:header>
                <el-button @click="appendPitrue" size="small">添加图片</el-button>
            </template>
        </el-table-column>
    
        <el-table-column>
            <template v-slot:default="scope">
                <el-button size="small">修改图片</el-button>
                <el-button size="small">删除图片</el-button>
            </template>
        </el-table-column>
    </el-table>
    <el-dialog v-model="centerDialogVisible" class="flex flex-col" top="1vh" width="95vw" center>
        <el-row :gutter="5" @click="changeItme">
            <el-col :span="6" v-for="item in banner" :key="item.uuid">
                <el-image class="w-full cursor-pointer" :uuid="item.uuid" :src="item.large"></el-image>
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
</template>
<script  setup lang="ts" >
import * as api from "../../api"
import { reactive, toRefs, toRaw, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
const state = reactive<dataType>({
    tableData: [],
    centerDialogVisible: false,
    banner: [],
    saveList: [],
    pictures: 0,
})
const loadImg = (current: number) => {
    api.wallhaven({
        limit: 12,
        offset: current
    }).then(({ count, imgs }) => {
        state.pictures = count
        state.banner = imgs
    })
}
const appendPitrue = () => {
    state.centerDialogVisible = !state.centerDialogVisible
    api.wallhaven({
        limit: 12,
        offset: 1
    }).then(({ count, imgs }) => {
        state.pictures = count
        state.banner = imgs
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

async function appendBanner() {
    let bannerIds = state.saveList.map(item => item.uuid)
    let result = await api.appendBanner(bannerIds)
    if (result.code == 200) {
        ElMessage.success(result.message)
    } else {
        ElMessage.warning(result.message)
    }

}
onMounted(async () => {
    let { result } = await api.getBanner()
    state.tableData = result
})

let { tableData, centerDialogVisible, banner, pictures } = toRefs(state)
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

