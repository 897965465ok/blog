<template>
  <div>
    <div class="viewer"></div>
    <InputBox></InputBox>
    <Comment :comments="comments"></Comment>
  </div>
</template>
<script>
import Comment from "./Comment.vue";
import InputBox from "./InputBox.vue";
import {getComments} from "../../api/BlogApi"
export default {
  name: "About",
  props: {
    comments: {
      type: Array,
      default: [],
    },
  },
  components: {
    Comment,
    InputBox,
  },
  async mounted() {
    let articlerPath = this.$route.query.articlerPath.replace(
      "http://localhost:3800",
      ""
    );
    let { data } = await this.$api.get("/v1" + articlerPath);
    const { Editor } = toastui;
    const { codeSyntaxHighlight } = Editor.plugin;
    const viewer = Editor.factory({
      el: document.querySelector(".viewer"),
      viewer: true,
      height: "100%",
      plugins: [[codeSyntaxHighlight, { highlighter: Prism }]],
    });
    viewer.setMarkdown(data);
    // viewer.invoke("setMarkdown", data);
    // 请求评论 // 先不写
    let articleId = this.$route.query.uuid;
    this.comments = await getComments(articleId);
  },
};
</script>
<style lang="scss" scoped>
/deep/ .toastui-editor-contents {
  font-size: 18px;
  background: #ffffff;
  h1,
  h2,
  h3,
  h4,
  h5,
  h6 {
    padding: 10px;
    border-bottom: none;
  }
  pre {
    margin: 0px;
    padding: 12px;
    background-color: transparent;
  }
  @apply bg-gray-50;
}
</style>
