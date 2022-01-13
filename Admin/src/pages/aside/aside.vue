<template>
    <el-Aside>
        <el-row class="h-full">
            <el-col class="h-full">
                <el-menu class="h-full" default-active="1">
                    <el-menu-item
                        :index="item.path"
                        v-for="item in MenuList"
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
            text: "轮播图",
            path: "/banner"
        },
        {
            iconName: "PriceTag",
            text: "用户",
            path: "/user"
        },
        {
            iconName: "Setting",
            text: "文章",
            path: "/article"
        }
    ]
})
</script>