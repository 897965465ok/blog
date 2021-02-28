<template>
  <el-row class="footer">
    <el-col :span="8">
      <el-menu mode="horizontal">
        <el-menu-item v-for="item in Menu" :key="item.title">
          <router-link :to="item.link"> {{ item.title }} </router-link>
        </el-menu-item>
      </el-menu>
    </el-col>
    <el-col :offset="8" :span="8">
      <el-menu mode="horizontal">
        <el-menu-item>
          <div class="top_ver">
            ©The website was set up on 2021/01/3 . version v1.0
          </div>
        </el-menu-item>
      </el-menu>
    </el-col>
    <JumpTop></JumpTop>
  </el-row>
</template>
<script>
import JumpTop from "./Jump/Jump"
export default {
  components:{
   JumpTop
  },
  data: () => ({
    Menu: [
      {
        title: "首页",
        link: "/",
      },
      {
        title: "笔记",
        link: "/note",
      },
      {
        title: "关于我",
        link: "/about",
      },
      {
        title: "登陆",
        link: "/login",
      },
    ],
  }),
};
</script>

<style lang="scss" scoped>
.footer {
  a {
    text-decoration: none;
    font-size: 16px;
    &:hover {
      color: #00a8ff;
    }
  }

  .el-menu--horizontal {
    border: none;
  }
}
</style>