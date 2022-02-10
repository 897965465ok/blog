
<script setup lang="ts">
import { RouterView } from 'vue-router';
import asideVue from "../aside/aside.vue";
import headerVue from "../header/header.vue"
import * as echarts from "echarts";
</script>
<template>
    <el-container class="h-full">
        <asideVue></asideVue>
        <el-Main>
            <headerVue></headerVue>
            <RouterView></RouterView>
            <el-Footer></el-Footer>
        </el-Main>
    </el-container>
</template>