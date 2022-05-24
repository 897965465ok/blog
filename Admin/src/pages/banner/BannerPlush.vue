<template>
    <el-dialog v-model="openPlush" class="flex flex-col" top="1vh" width="95vw" center>
        <el-row style="min-height:81vh;" v-loading="loading" :gutter="5" @click="changeItme">
            <el-col class="mask" :span="6" v-for="item in banner" :key="item.uuid">
                <el-image class="cursor-pointer" :uuid="item.uuid" :src="item.large" lazy></el-image>
                <el-col class="mask-item rounded-xl">
                    <Icon v-if="item.check" class="Check" icon="Check"></Icon>
                </el-col>
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
                <el-button size="small" type="success" @click="bannerPlush">点击添加</el-button>
            </el-col>
        </el-row>
    </el-dialog>
</template>
<script  setup lang='ts'>
import {
    reactive,
    toRefs,
    computed,
    toRef,
    toRaw,
    onMounted,
    defineProps,
    defineEmits,
    onActivated,
onDeactivated
} from 'vue'
import { ElMessage } from 'element-plus'
import * as api from "../../api"
import { useStore } from 'vuex';
const store = useStore();
const state = reactive<any>({
    banner: computed(() => store.state.wallhavenPicture),
    pictures: computed(() => store.state.wallhavenLength),
    loading: computed(() => store.state.loading),
    saveList: [],
    count: 0,
})

defineProps({
    openPlush: Boolean
})

onActivated(()=>{
    console.log("进入触发")
})
onDeactivated(()=>{
    console.log("离开触发")
})
const loadImg = (current: number = 1) => {
    store.dispatch('wallhaven', { limit: 12, offset: current })

}
function changeItme(event: MouseEvent) {
    if (event.target != null && (event.target as HTMLImageElement).nodeName == 'IMG') {
        let uuid = (event.target as HTMLImageElement).getAttribute("uuid")
        let node = state.banner.find((item: { uuid: string | null; }) => item.uuid == uuid)
        if (node != undefined && !node.check && state.saveList.findIndex((item: { uuid: string | null; }) => item.uuid == uuid) == -1) {
            node.check = true
            state.saveList.push(node)
            ElMessage.success("成功添加到列表")
        } else {
            ElMessage.warning("已经拥有该图片")
        }
    }
}
async function bannerPlush() {
    let bannerIds = state.saveList.map((item: { uuid: any; }) => item.uuid)
    let result = await api.bannerPlush(bannerIds)
    ElMessage.success(result.data.message)
}
let { banner, pictures, loading } = toRefs(state)
</script>

<style lang="css" scoped>
.mask {
    display: flex;
    height: 100%;
    width: 100%;
    position: relative;
    top: 0;
    left: 0;
    max-width: none;
}
.mask-item {
    position: absolute;
    top: 0;
    left: 0;
    opacity: 0.5;
    background-color: black;
    max-width: none;
}

.Check {
    @apply text-rose-500 w-16 h-16;
}
</style>