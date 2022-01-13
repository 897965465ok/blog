
<script setup lang="ts">
</script>
<template>
    <el-Header>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/' }">
                <template v-slot:default="data">
                    <el-button size="mini" type="default">首页</el-button>
                </template>
            </el-breadcrumb-item>
        </el-breadcrumb>
    </el-Header>
</template>