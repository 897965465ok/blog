<template>
  <el-row class="main">
    <el-row class="main-content">
      <el-col class="content-left" :span="16">
        <el-row class="card-wrapper">
          <h3 class="sub-title">#最新文章</h3>
          <Articler
            v-for="item in articles.slice(0, 4)"
            :key="item.uuid"
            :article="item"
          ></Articler>
          <h3 class="sub-title">#博主推荐</h3>
          <Carousel ></Carousel>
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
        <h3 class="sub-title">#特别推荐</h3>
        <ReArticle
          v-for="item in articles.slice(9, 16)"
          :key="item.uuid"
          :article="item"
        ></ReArticle>
        <h3 class="sub-title">#小程序</h3>
        <el-row>
          <el-image
            src="https://tva2.sinaimg.cn/large/87c01ec7gy1frmmz605z4j21kw0w0qvh.jpg"
            alt="二维码图片"
            lazy
          ></el-image>
        </el-row>
        <h3 class="sub-title">#个人收藏网站</h3>
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
      ],
    };
  },

  async created() {
    await this.getjson();
    await this.getTages();
    await this.changeAll((await this.$api.get("v1/articles")).data.result);
    loading.close();
  },
  computed: mapState({
    pictures: (state) => (state.pictures.length ? state.pictures : []),
    articles: (state) => (state.articlers.length ? state.articlers : []),
  }),
  watch: {
    pictures: (value) => value,
  },
  components: {
    Carousel,
    Profile,
  },
  methods: {
    ...mapActions(["changeAll", "getPictures", "setRecommen"]),
    async getjson() {
      for (let index = 0; index < 5; index++) {
        let item = await this.$generateJSON();
        await this.getPictures(item);
      }
    },
    async getTages() {
      this.setRecommen({
        tags: (await this.$api.get("v1/tags")).data.result,
        imagUrl: (await this.$generateJSON()).imgurl,
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.main {
  height: 100%;
  .main-header {
    height: 300px;
    box-sizing: border-box;
    .header-right {
      padding-left: 5px;
    }
  }
  .content-left {
    min-height: 1650px;
  }
  .revert-wrapper {
    // padding: 6px;
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