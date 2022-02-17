<template>
    <!-- 轮播图 -->
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
    <BannerPlush
     v-model="openPlush"
    :openPlush="openPlush"
  ></BannerPlush>
    <!-- <BannerRemove v-model="openRemove" :isShow="openRemove"></BannerRemove> -->
    <el-row>
        <el-col>
            <el-button size="small" @click="bannerPlush" :icon="Delete" type="primary">批量增加</el-button>
            <!-- <el-button size="small" :icon="Edit" type="info">修改</el-button> -->
            <el-button size="small" @click="bannerRemove" :icon="Delete" type="danger">批量删除</el-button>
        </el-col>
    </el-row>
    <el-row class="min-h-full" v-loading="tableLoading">
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
        <paginationVue @loadTime="loadBanner" :count="count" :pageSize="5" callback="loadTime"></paginationVue>
    </el-row>
</template> 

<script  setup lang="ts" >
import { reactive, ref, toRefs, toRaw, onMounted, computed } from 'vue'
import { Edit, Share, Delete } from '@element-plus/icons-vue'
import paginationVue from "../../components/pagination.vue"
import BannerPlush from "./BannerPlush.vue"
import BannerRemove from "./BannerRemove.vue"
import { useStore } from 'vuex'
// 表头水平居中：给 el-table 标签绑定 header-cell-style 属性
// <el-table :header-cell-style="{'text-align':'center'}"></el-table>
// 表格内容水平居中：给 el-table 标签绑定 cell-style 属性
// <el-table :cell-style="{textAlign:'center'}"></el-table>
const store = useStore()
const state = reactive({
    openRemove: false,
    openPlush: false
})
const center = { 'text-align': 'center' }
onMounted(() => loadBanner())
let tableData = ref(computed(() => store.state.Banner))
let count = ref(computed(() => store.state.BannerLenght))
let tableLoading = ref(computed(() => store.state.tableLoading))

const loadBanner = (current: number = 1) => {
    store.dispatch('loadBanner', { limit: 5, offset: current })
}
const bannerPlush = () => {
    store.dispatch('wallhaven', { limit: 12, offset: 1 })
    state.openPlush = !state.openPlush
}

const bannerRemove = () => state.openRemove = !state.openRemove

const deleteBanner = (ID) => store.dispatch('deleteBanner', ID)

const { openRemove,openPlush } = toRefs(state)
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

