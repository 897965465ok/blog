<template>
    <div class="w-full h-full">
        <div class="w-full h-1/2"></div>
        <div class="w-full h-1/2"></div>
    </div>
</template>
<script  setup lang='ts'>
import { reactive, defineProps, onMounted } from 'vue'
import * as echarts from 'echarts';
onMounted(() => {
    // 基于准备好的dom，初始化echarts实例
    const Chart = echarts.init((document.querySelector("#count") as any));
    // 绘制图表
    Chart.setOption(
        {
            xAxis: {
                type: 'category',
                data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
            },
            yAxis: {
                type: 'value'
            },
            series: [
                {
                    data: [150, 230, 224, 218, 135, 147, 260],
                    type: 'line'
                }
            ]
        }
    );
})
</script>