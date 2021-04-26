<template>
  <el-row class="card-item"  @click.native="readArticler(article.article_path)">
    <el-col class="card-left">
      <el-row class="title">
        <h4>{{ article.name.replace('.md','') }}</h4>
      </el-row>
      <el-row class="footer">
        <!-- <span>{{ article.CreatedAt }}</span> -->
        <!-- <span> 浏览: {{ article.whatch_number }}</span> -->
      </el-row>
    </el-col>
  </el-row>
</template>
<script>
export default {
  props: ["article"],
  name: "ReArticle",
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
  height: 40px;
  box-sizing: border-box;
  background-color: #fff;
  border: none;
  border-bottom: solid 1px #999;
   &:hover {
      color: #52a8ff;
          background: #F4F5F5;
    }
  .card-left {
    height: 40px;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
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
    img {
      height: 90%;
      width: 90%;
    }
  }
}
</style>