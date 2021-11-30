<template>
  <el-row class="profile">
    <el-row class="user">
      <h3>#Jiang</h3>
    </el-row>
    <el-row class="pictuer">
      <img class="el-image" src="../../../assets/jerry.png" />
    </el-row>
    <el-row class="tag">
      <ul style="padding:0px;">
        <li v-for="(item, index) in icon" :key="index">
          <a :href="item.link">
            <i class="iconfont" :class="item.icon"></i>
            <p >{{ item.title }}</p>
          </a>
        </li>
      </ul>
    </el-row>
  </el-row>
</template>
<script>
import { mapState } from "vuex";
export default {
  data() {
    return {
      icon: [
        {
          title: "Github",
          icon: "icon-huaban88",
          link: "https://github.com/897965465ok",
        },
        {
          title: "QQ",
          icon: "icon-qq",
          link: "http://wpa.qq.com/msgrd?v=3&uin=897965465&site=qq&menu=yes",
        },
        // {
        //   title: "Email",
        //   icon: "icon-youxiang",
        //   link:"Mailto:897965465@qq.com",
        // },
        // {
        //   title: "微信",
        //   icon: "icon-weixin",
        //   link: "",
        // },
      ],
    };
  },
  computed: mapState({
    recommen(state) {
      if (state.recommen) {
        return state.recommen;
      }
    },
  }),
};
</script>
<style lang="scss" scoped>
@import "../../../assets/icon/iconfont.css";
.profile {
  height: 300px;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
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
  .tag {
    // height: 70px;
    ul {
      display: flex;
      li {
        flex: 1;
        display: flex;
        flex-direction: column;
        text-align: center;
        list-style: none;
        cursor: pointer;
        margin: 0px 6px;
        a {
          color: #322931;
          font-size: 16px;
          text-decoration: none;
        }
        i {
          font-size: 40px;
        }
      }
    }
  }
}
</style>