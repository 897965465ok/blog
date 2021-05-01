<template>
  <el-row class="card-item" @click.native="readArticler(article.article_path)">
    <el-row class="item-wrapper" :span="24">
      <el-col class="card-left" :span="19">
        <el-row class="title">
          <h4>{{ article.name.replace(".md", "") }}</h4>
        </el-row>
        <el-row class="paragraph">
          <!-- {{ article.article_path }} -->
        </el-row>
        <el-row class="footer">
          <div class="footer-left">
            <span class="item-tag">{{ article.tag }}</span>
            <span class="iconfont icon-zan">5</span>
          </div>
          <div class="footer-right">
            <span class="iconfont icon-yuedu">
              {{ article.whatch_number }}</span
            >
            <span class="iconfont icon-shijian"> {{ article.CreatedAt }}</span>
          </div>
        </el-row>
      </el-col>
      <el-col class="card-right" :span="5">
        <el-image :src="url" lazy>
          <div slot="error" class="image-slot">
            <i class="el-icon-picture-outline"></i>
          </div>
        </el-image>
      </el-col>
    </el-row>
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
@import "../assets/icon/iconfont.css";
.card-item {
  cursor: pointer;
  height: 135px;
  box-sizing: border-box;
  background-color: #fff;
  border: none;
  border-bottom: solid 1px #999;
  display: flex;
  align-items: center;
  .item-wrapper {
    height: 115px;
    width: 100%;
    .card-left {
      height: 100%;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      .title {
        font-size: 15px;
        color: #333;
        font-weight: 700;
        &:hover {
          color: #52a8ff;
        }
      }
      .footer {
        width: 100%;
        color: #999;
        display: flex;
        .footer-left,
        .footer-right {
          flex: 1 0 50%;
          span {
            margin: 0px 6px;
          }
          .item-tag {
            padding: 4px;
            display: inline-block;
            line-height: 13px;
            border: solid salmon 1px;
            border-radius: 14px;
            color: salmon;
            font-size: 12px;
          }
        }
        .footer-right {
          display: flex;
          justify-content: flex-end;
        }
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
}
</style>