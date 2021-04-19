<template>
  <el-row class="main">
    <el-row class="main-content">
      <el-col class="content-left" :span="16">
        <el-row>
          <el-menu mode="horizontal">
            <el-menu-item
              v-for="item in recommen.tags && recommen.tags"
              :key="item.ID"
            >
              <span v-if="item.article_tag" @click="check(item.article_tag)">
                {{ item.article_tag }}
              </span>
            </el-menu-item>
          </el-menu>
        </el-row>
        <el-row class="card-wrapper">
          
          <Articler
            v-for="item in tmp"
            :key="item.uuid"
            :article="item"
          ></Articler>

        </el-row>
        <el-row>
          <el-col class="pagination">
            <el-pagination
              @current-change="change"
              layout="prev, pager, next"
              :current-page="current"
              :page-count="len"
            >
            </el-pagination>
          </el-col>
        </el-row>
      </el-col>
      <el-col class="content-right" :span="8">
        <h3 class="sub-title">#精选文章</h3>
        <ReArticle
          v-for="item in articles.slice(9, 14)"
          :key="item.uuid"
          :article="item"
        ></ReArticle>
        <!-- <el-row>
          <el-image v-if="recommen.imagUrl" :src="recommen.imagUrl"></el-image>
        </el-row> -->
        <!-- <h3 class="sub-title">#热门标签</h3>
        <el-row class="friends-link">
          <waterfall :col="3" :data="favorites" :gutterWidth="10">
            <div
              class="friends-item"
              v-for="(item, index) in favorites"
              :key="index"
            >
              <el-button>
                {{ item.userName }}
              </el-button>
            </div>
          </waterfall>
        </el-row> -->
      </el-col>
    </el-row>
  </el-row>
</template>
<script>
import { mapState, mapActions } from "vuex";
export default {
  name: "note",
  data() {
    return {
      favorites: [
        {
          userName: "动漫之家",
          id: 120,
          link: "链接",
        },
      ],
      tmp: [],
      current: 1,
    };
  },
  async mounted() {

    await this.setSimilar([]);

  },

  computed: {
    ...mapState({
      recommen: (state) => state.recommen,
      articles: (state) => (state.articlers.length ? state.articlers : []),
      similar: (state) => (state.similar.length ? state.similar : []),
    }),
    // 监听
    listen() {
      return {
        articles: this.articles,
        similar: this.similar,
      };
    },
    // 获取数量
    len() {
      return this.similar.length
        ? Math.ceil(this.similar.length / 8)
        : Math.ceil(this.articles.length / 8);
    },
  },
  watch: {
    listen(listen) {
      this.tmp =
        listen.similar.length > 0
          ? listen.similar.slice(0, 7)
          : listen.articles.slice(0, 7);
    },
  },
  methods: {
    ...mapActions(["setSimilar"]),
    change(index) {
      let end  =  index * 6
      let start =   end - 6
      this.tmp =
        this.similar.length > 0
          ? this.similar.slice(start, end)
          : this.articles.slice(start, end);

      
    },
    async check(articleTag) {
      this.current = 0;
      let similar = this.articles.filter((item) => item.tag == articleTag);
      await this.setSimilar(similar);
    },
  },
};
</script>

<style lang="scss" scoped>
.main {
  height: 100%;
  background: #ffffff;
  .content-left {
    padding: 16px;
  }

  .el-menu--horizontal > .el-menu-item {
    height: 35px;
    line-height: 35px;
    a {
      text-decoration: none;
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
  .pagination {
    display: flex;
    justify-content: center;
  }
}
</style>