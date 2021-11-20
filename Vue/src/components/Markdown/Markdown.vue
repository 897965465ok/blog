<template>
  <div>
    <div class="viewer"></div>
    <Comment></Comment>
  </div>
</template>
<script>

export default {
  name: "About",
  components: {
    Comment:()=>import("./Comment")
  },
  async mounted() {
    let articlerPath = this.$route.query.articlerPath.replace(
      "http://localhost:3800",
      ""
    );
    let { data } = await this.$api.get("/v1" + articlerPath);
    const Viewer = toastui.Editor;
    const viewer = new Viewer({
      el: document.querySelector(".viewer"),
      height: "1200px",
    });
    viewer.setMarkdown(data);
    // viewer.invoke("setMarkdown", data);
    document.querySelectorAll("pre code").forEach((block) => {
      hljs.highlightBlock(block);
    });
  },
};
</script>
<style lang="scss" scoped>
/deep/.viewer {
  .tui-editor-contents {
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
  }
}
</style>