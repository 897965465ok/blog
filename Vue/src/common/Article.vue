<template>
  <el-row class="card-item" @click.native="readArticler(article.article_path)">
    <el-col class="card-left" :span="19">
      <el-row class="tager">
        <span>{{ article.tag }}</span>
      </el-row>
      <el-row class="title"
        ><h4>{{ article.name.replace(".md", "") }}</h4></el-row
      >
      <el-row class="paragraph">
        {{ article.article_path }}
      </el-row>
      <el-row class="footer">
        <!-- <span>{{ article.CreatedAt }}</span> -->
        <!-- <span> 浏览:{{ article.whatch_number }}</span> -->
      </el-row>
    </el-col>
    <el-col class="card-right" :span="5">
      <el-image :src="url" lazy></el-image>
    </el-col>
  </el-row>
</template>
<script>
import { mapState } from "vuex";
export default {
  data() {
    return {
      url: "",
    };
  },
  mounted() {
    let random = Math.floor(Math.random() * (200 - 1) + 1);
    this.url = this.pictures[random];
  },
  computed: mapState({
    pictures: (state) => (state.pictures.length ? state.pictures : []),
  }),
  watch: {
    pictures: (value) => value,
  },

  props: ["article"],
  name: "Articler",
  methods: {
    readArticler(articlerPath) {
      this.$router.push({ path: "markdown", query: { articlerPath } });
    },
  },
};
</script>

<style lang="scss" scoped>
.card-item {
  cursor: pointer;
  height: 140px;
  margin: 1px 0px;
  padding: 6px 0px;
  box-sizing: border-box;
  background-color: #fff;
  border: none;
  border-bottom: solid 1px #999;
   &:hover {
      background: #F4F5F5;
    }
  .tager {
    font-size: 14px;
    color: #999;
    &:hover {
      color: #52a8ff;
    }
  }
  .card-left {
    height: 140px;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    .title {
      font-size: 16px;
      color: #333;
      font-weight: 700;
      &:hover {
        color: #52a8ff;
      }
    }
    .paragraph,
    .footer {
      color: #999;
    }
  }
  .card-right {
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    .el-image {
      height: 100%;
      width: 100%;
      border-radius: 5px;
    }
    img {
      height: 100%;
      width: 100%;
    }
  }
}
</style>