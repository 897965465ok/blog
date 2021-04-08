<template>
  <el-row class="card-item" @click.native="readArticler(article.article_path)">
    <el-col class="card-left" :span="19">
      <el-row class="tager">
        <span>#</span>
        <span>{{ article.tag }}</span>
      </el-row>
      <el-row class="title"
        ><h4>{{ article.name }}</h4></el-row
      >
      <el-row class="paragraph">
        <!-- {{ article.article_path }} -->
      </el-row>
      <el-row class="footer">
        <span>{{ article.CreatedAt }}</span>
        <!-- <span> 浏览:{{ article.whatch_number }}</span> -->
      </el-row>
    </el-col>
    <!-- <el-col class="card-right" :span="5">
      <el-image :src="pic" lazy></el-image>
    </el-col> -->
  </el-row>
</template>
<script>
export default {
  data() {
    return {
      pic: "",
    };
  },
  async created() {
    // 去除 \ufeff
    // let data = await this.$generateJSON();
    // this.pic = data.imgurl;
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
  margin-bottom: 6px;
  border-radius: 5px;
  box-sizing: border-box;
  border-bottom: 1px solid rgb(238, 238, 238);
  .card-left {
    height: 140px;
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