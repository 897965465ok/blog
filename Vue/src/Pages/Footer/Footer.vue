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
            <a   href='https://beian.miit.gov.cn' target="_blank">备案号:  桂ICP备2021004482号-1</a>
            ©The website was set up on 2021/01/3 . version v2.0
          </div>
        </el-menu-item>
      </el-menu>
    </el-col>
    <JumpTop></JumpTop>
  </el-row>
</template>
<script>
import JumpTop from "./Jump/Jump";
export default {
  components: {
    JumpTop,
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
        title: "登陆",
        link: "/login",
      },
    ],
  }),
};
</script>

<style lang="scss" scoped>
.footer {
  background: #ffffff;

  a {
    text-decoration: none;
    font-size: 16px;

    &:hover {
      color: #00a8ff;
    }
  }

  .top_ver {
    display: flex;

    a {
      margin-right: 30px;
    }
  }

  .el-menu--horizontal {
    border: none;
  }
}
</style>