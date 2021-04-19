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
    <el-col class="card-right" :span="5">
      <el-image :src="url" lazy></el-image>
    </el-col>
  </el-row>
</template>
<script>
import { mapState} from "vuex";
export default {
  data() {
    return {
      url: "",
    };
  },
  mounted(){
    let random =  Math.floor(Math.random()*(200 - 1) + 1)
    this.url = this.pictures[random]
  },
  computed: mapState({
    pictures: (state) => (state.pictures.length ? state.pictures : [])
  }),
  watch: {
    pictures: (value) => value
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
  margin: 6px 0px;
  padding: 0px 0px 0px 10px;
  border-radius: 5px;
  box-sizing: border-box;
  background-color: rgba(255, 255, 255, 0.904);
  border: solid 1px #d0cfd0;
  backdrop-filter: blur(50px);
  // box-shadow: 3px 3px 3px #00A1D6 ;
  transition: all 2s;

  &:hover {
    //  transform: scale(1.1) ;
  }
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