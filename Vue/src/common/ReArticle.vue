<template>
  <el-row class="card-item"  @click.native="readArticler(article.article_path)">
    <el-col class="card-left">
      <el-row class="title">
        <h4>{{ article.name }}</h4>
      </el-row>
      <el-row class="footer">
        <span>{{ article.CreatedAt }}</span>
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
  height: 80px;
  margin: 6px 0px;
  padding: 0px  10px;
  border-radius: 5px;
  box-sizing: border-box;
  border-bottom: 1px solid rgb(238, 238, 238);
  border: solid 1px #999;
  transition: all 1s;
   &:hover{
     transform: scale(1.1) ;
  }
  .card-left {
    height: 80px;
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