<template>
    <el-Aside>
        <el-row class="h-full">
            <el-col class="h-full">
                <el-menu
                    router
                    background-color="#545c64"
                    text-color="#fff"
                    active-text-color="#ffd04b"
                    class="h-full"
                    default-active="1"
                >
                    <el-sub-menu :index="item.text" v-for="(item, index) in MenuList" :key="index">
                        <template v-slot:title>
                            <el-icon>
                                <Icon :icon="item.iconName"></Icon>
                            </el-icon>
                            <span>{{ item.text }}</span>
                        </template>
                        <el-menu-item
                            v-for="child in item.children"
                            :key="child.path"
                            :index="child.path"
                        >
                            <el-icon>
                                <Icon :icon="child.iconName"></Icon>
                            </el-icon>
                            <span>{{ child.text }}</span>
                        </el-menu-item>
                    </el-sub-menu>
                </el-menu>
            </el-col>
        </el-row>
    </el-Aside>
</template>
<script  setup lang="ts">
import { reactive } from 'vue';
import { useRouter, useRoute } from "vue-router"
const router = useRouter()
let { MenuList } = reactive({
    MenuList: [
        {
            iconName: "HomeFilled",
            text: "博客管理",
            children: [
                {
                    iconName: "PictureFilled",
                    text: "轮播图",
                    path: "/banner"
                },
                {
                    iconName: "Setting",
                    text: "首页文章",
                    path: "/article"
                },
                {
                    iconName: "PictureFilled",
                    text: "音乐管理",
                    path: "/music"
                },
                // {
                //     iconName: "Setting",
                //     text: "侧边栏文章",
                //     path: "/article"
                // },
                // {
                //     iconName: "Setting",
                //     text: "喜欢的网站",
                //     path: "/article"
                // }
            ]
        },
        {
            iconName: "Setting",
            text: "后台管理",
            children: [
                {
                    iconName: "PriceTag",
                    text: "用户",
                    path: "/user"
                },
            ]
        },


    ]
})
</script>