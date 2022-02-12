<template>
  <el-col
    :span="8"
    class="card-item"
    @click.native="readArticler(article.article_path, article.uuid)"
  >
    <el-row class="item-wrapper">
      <el-row class="card-content">
        <el-row class="title">
          <h5>{{ article.name.replace(".md", "") }}</h5>
        </el-row>
        <!-- <el-row class="paragraph">
          {{ article.article_path }}
        </el-row>-->
        <el-row class="card-middle">
          <el-image :src="url" lazy>
            <div slot="error" class="image-slot">
              <i class="el-icon-picture-outline"></i>
            </div>
          </el-image>
        </el-row>
        <el-row class="footer">
          <div class="footer-top">
            <span class="item-tag">{{ article.tag }}</span>
            <span
              @click.stop="like(article.uuid, article)"
              class="iconfont icon-zan"
            >{{ article.like }}</span>
          </div>
          <div class="footer-bottom">
            <span class="iconfont icon-yuedu">{{ article.whatch_number }}</span>
            <span class="iconfont icon-shijian">{{ $qsTime(article.CreatedAt) }}</span>
          </div>
        </el-row>
      </el-row>
    </el-row>
  </el-col>
</template>
<script>
import { mapState } from "vuex";
import * as api from "../api/BlogApi"
export default {
  data() {
    return {
      url: "",
    };
  },
  mounted() {
    let random = this.$random(0, 216)
    if (this.pictures[random]) {
      this.url = this.pictures[random].small;
    }
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
    async readArticler(articlerPath, uuid) {
      await api.readArticler(uuid)
      this.$router.push({ path: "markdown", query: { articlerPath, uuid } });
    },
    async like(uuid, article) {
      await api.like(uuid)
      article.like += 1;
    },
  },
};
</script>

<style lang="scss" scoped>
@import "../assets/icon/iconfont.css";
.card-item {
  cursor: pointer;
  box-sizing: border-box;
  background-color: #fff;
  border: none;
  display: flex;
  align-items: center;
  margin-bottom: 12px;
  .item-wrapper {
    width: 100%;

    .card-content {
      height: 100%;
      display: flex;
      flex-direction: column;
      justify-content: space-around;
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
        flex-direction: column;
        .footer-top {
          display: flex;
          justify-content: space-between;
          margin: 6px 0px;
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
        .footer-bottom {
          display: flex;
          justify-content: space-between;
        }
      }
    }
    .card-middle {
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
