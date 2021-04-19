<template>
  <el-row class="header">
    <el-col :span="2">
      <div class="logo">
        <router-link to="/">
          <img src="../../assets/logo.png" alt="logo" />
        </router-link>
      </div>
    </el-col>
    <el-col :offset="12" :span="10">
      <el-menu mode="horizontal">
        <el-menu-item v-for="item in Menu" :key="item.title">
          <router-link :to="item.link"> {{ item.title }} </router-link>
        </el-menu-item>
      </el-menu>
    </el-col>
  </el-row>
</template>

<script>
export default {
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
.header {
    background:#FFFFFF;
  a {
    text-decoration: none;
    font-size: 16px;
    &:hover {
      color: #00a8ff;
    }
  }
  .logo {
    height: 60px;
    img {
      height: 100%;
    }
  }

  .el-menu--horizontal {
    border: none;
  }
}
</style>