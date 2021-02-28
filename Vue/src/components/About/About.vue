<template>
  <el-row class="about">
    <el-row class="user">
      <h3>#Jiang</h3>
    </el-row>
    <el-row class="pictuer">
      <img class="el-image" src="../../assets/jerry.png" />
    </el-row>
    <el-row class="content">
      <el-col :span="8" class="content-item">
        <p class="iconfont icon-nianling"></p>
        <span>基本信息</span>
        <el-row class="text">
          <p>Jiang|男|28|南宁</p>
          <p>18176283201</p>
          <p>897965465@qq.com</p>
          <p>博客</p>
          <p>Github</p>
        </el-row>
      </el-col>
      <el-col :span="8" class="content-item">
        <p class="iconfont icon-jineng-copy"></p>
        <span>技能清单</span>
        <el-row class="text">
          <p>html</p>
          <p>css</p>
          <p>Vue</p>
          <p>javascript</p>
          <p>electron-Vue</p>
        </el-row>
      </el-col>
      <el-col :span="8" class="content-item">
        <p class="iconfont icon-xiangmu"></p>
        <span>工作经历</span>
        <el-row class="text">
          <p>微致汇网络科技有限公司</p>
          <p>2019年3月|二年工作经验</p>
        </el-row>
      </el-col>
    </el-row>
  </el-row>
</template>
<script>
import Profile from "@/components/Main/Profile/Profile";
export default {
  name: "About",
  components: {
    Profile,
  },
  async created() {},
  methods: {},
};
</script>

<style lang="scss" scoped>
@import "../../assets/icon/iconfont.css";
.about {
  height: 100vh;
  .user {
    height: 35px;
  }
  .pictuer {
    width: 100%;
    text-align: center;
    .el-image {
      height: 150px;
      width: 150px;
      border-radius: 50%;
    }
  }
  .content {
    .content-item {
      text-align: center;
      .iconfont{
        font-size: 60px;
      }
      .text{
        height: 120px;
      }
    }
  }
}
</style>