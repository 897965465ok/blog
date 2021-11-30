<template>
  <el-row class="gossip">
    <el-row>
      <el-col :span="8" class="card" v-for="item in articles" :key="item.id">
        <div class="picture">
          <!-- <img :src="item.picture" /> -->
        </div>
        <h4>{{ item.paragraph }}</h4>
      </el-col>
    </el-row>
    <el-col class="pagination">
      <el-pagination layout="prev, pager, next" :total="1000"> </el-pagination>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "Gossip",
  data() {
    return {
      articles: [],
    };
  },
  async created() {
    let { data } = await this.$api.get("/v1/gossip");
    this.articles = data["articles"];
  },
};
</script>

<style lang="scss" scoped>
.gossip {
  height: 100vh;
  text-align: center;
}
</style>