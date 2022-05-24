<script  setup lang="ts">
import { RouterView } from 'vue-router'
import asideVue from "./pages/aside/aside.vue";
import headerVue from "./pages/header/header.vue"
import * as echarts from "echarts";
</script>
<template>
    <el-container class="h-full">
        <asideVue></asideVue>
        <el-Main>
            <headerVue></headerVue>
            <router-view v-slot="{ Component, route }">
                <transition :name="(route.meta.transitionName as string)">
                    <keep-alive v-if="route.meta.keepAlive">
                        <component :is="Component" />
                    </keep-alive>
                    <component :is="Component" v-else />
                </transition>
            </router-view>
            <el-Footer></el-Footer>
        </el-Main>
    </el-container>
</template>

