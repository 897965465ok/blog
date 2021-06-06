<template>
  <el-container>
    <el-header>
      <Header></Header>
      
    </el-header>
    <el-main>
      <router-view />
    </el-main>
    <el-footer>
      <Footer></Footer>
    </el-footer>
      
  </el-container>
</template>
<script>
import Header from "@/components/Header/Header.vue";
import Footer from "@/components/Footer/Footer.vue";
export default {
  name: "home",
  components: {
    Header,
    Footer,
  },
};
</script>
<style lang="scss" scoped>
.el-container {
  width: 1260px;
  height: 100%;
  .el-main {
    height: 100%;
    width: 100%;
    
  }
  margin: auto;
}
</style>