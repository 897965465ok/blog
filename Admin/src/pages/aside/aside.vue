<template>
    <el-Aside>
        <el-row class="h-full">
            <el-col class="h-full">
                <el-menu class="h-full" default-active="1">
                    <el-menu-item
                        :index="index"
                        v-for="(item,index) in MenuList"
                        :key="item.path"
                        @click="jump(item.path)"
                    >
                        <el-icon>
                            <Icon :icon="item.iconName"></Icon>
                        </el-icon>
                        <span>{{ item.text }}</span>
                    </el-menu-item>
                </el-menu>
            </el-col>
        </el-row>
    </el-Aside>
</template>
<script  setup lang="ts">
import { reactive } from 'vue';
import { useRouter, useRoute } from "vue-router"
const router = useRouter()
const route = useRoute()
const jump = (path: string) => {
    router.push(path)
}
let { MenuList } = reactive({
    MenuList: [
        {
            iconName: "HomeFilled",
            text: "首页",
            path: "/"
        },
        {
            iconName: "PictureFilled",
            text: "轮播图管理",
            path: "/banner"
        },
        {
            iconName: "PriceTag",
            text: "标签管理",
            path: "/tags"
        },
        {
            iconName: "Setting",
            text: "文章管理",
            path: "/article"
        }
    ]
})
</script>