<template>
    <el-pagination
        background
        :page-size="pageSize"
        :pager-count="pageSize"
        layout="prev, pager, next,jumper"
        :total="count"
        @current-change="loadTime"
    ></el-pagination>
</template>
<script  setup lang="ts">
import { getCurrentInstance, defineProps } from 'vue'
const { proxy: Instance }: any = getCurrentInstance()
const props = defineProps<{
    callback: string,
    pageSize: number,
    count: number
}>()
const loadTime = (current: number) => {
    Instance.$emit(props.callback, current)
}
</script>