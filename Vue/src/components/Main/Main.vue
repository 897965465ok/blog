<template>
  <el-row class="main">
    <el-row class="main-content">
      <el-col class="content-left" :span="16">
        <el-row class="card-wrapper">

          <Articler
            v-for="item in articles.slice(0, 4)"
            :key="item.uuid"
            :article="item"
          ></Articler>
         
          <Carousel></Carousel>
          <Articler
            v-for="item in articles.slice(4, 9)"
            :key="item.uuid"
            :article="item"
          ></Articler>
        </el-row>
      </el-col>
      <el-col class="content-right" :span="8">
        <el-row class="main-header">
          <el-col class="header-right">
            <Profile></Profile>
          </el-col>
        </el-row>
      
        <ReArticle
          v-for="item in articles.slice(9, 16)"
          :key="item.uuid"
          :article="item"
        ></ReArticle>
   
        <el-row>
          <el-image
            src="https://tva2.sinaimg.cn/large/87c01ec7gy1frmmz605z4j21kw0w0qvh.jpg"
            alt="二维码图片"
            lazy
          ></el-image>
        </el-row>
    
        <el-row class="favorites-link">
          <waterfall :col="3" :data="favorites" :gutterWidth="10">
            <div
              class="favorites-item"
              v-for="(item, index) in favorites"
              :key="index"
            >
              <el-button>
                {{ item.userName }}
              </el-button>
            </div>
          </waterfall>
        </el-row>
      </el-col>
    </el-row>
  </el-row>
</template>
<script>
import Carousel from "./Carousel/Carousel";
import Profile from "./Profile/Profile";
import { mapState, mapActions } from "vuex";
import Vue from "vue";
let loading;
(() => {
  let once = true;
  return () => {
    if (once) {
      loading = Vue.prototype.$loading();
      once = false;
    }
  };
})()();

export default {
  data() {
    return {
      loading: null,
      favorites: [
        {
          userName: "动漫之家",
          id: 120,
          link: "http://www.imomoe.ai",
        },
         {
          userName: "bilibili",
          id: 120,
          link: "https://www.bilibili.com/",
        },
      ],
    };
  },

 async  created() {
      await this.getPictures(await this.$GetUrl());
      await this.setRecommen({
        tags: (await this.$api.get("v1/tags")).data.result
      });
      await this.changeAll((await this.$api.get("v1/articles")).data.result);
    loading.close();
  },
  computed: mapState({
    articles: (state) => (state.articlers.length ? state.articlers : []),
  }),

  components: {
    Carousel,
    Profile,
  },
  methods: {
    ...mapActions(["changeAll", "getPictures", "setRecommen"])
  },
};
</script>

<style lang="scss" scoped>

.main {
  height: 100%;
  background: #ffffff;
  .main-header {
    height: 300px;
    box-sizing: border-box;
    .header-right {
      padding-left: 5px;
    }
  }
  .content-left {
    min-height: 1650px;
    .card-wrapper {
      padding: 12px;
    }
  }
  .revert-wrapper {
    border-bottom: 1px solid #999;
    .message-wrapper {
      margin: 5px 0px;
      height: 60px;
      padding-left: 5px;
      border-left: 5px solid #999;
    }
  }
  .button-wrapper {
    margin: 6px 0px;
    .el-button {
      width: 100%;
    }
  }
  .favorites-link {
    width: 100%;
    .favorites-item {
      padding: 6px;
      margin: 6px;
      border-radius: 6px;
      text-align: center;
      .el-button {
        width: 100%;
      }
    }
  }
}
</style>