
<script setup lang="ts">
import { RouterView } from 'vue-router';
import asideVue from "../aside/aside.vue";
</script>
<template>
    <el-container class="h-full">
        <asideVue></asideVue>
        <el-Main>
            <el-Header></el-Header>
            <RouterView></RouterView>
            <el-Footer></el-Footer>
        </el-Main>
    </el-container>
</template>