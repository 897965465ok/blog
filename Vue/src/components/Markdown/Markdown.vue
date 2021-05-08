<template>
  <div id="viewer"></div>
</template>

<script>
// import hljs from "highlight.js";
// import "highlight.js/styles/tomorrow.css";

export default {
  name: "About",
  async mounted() {
    let articlerPath = this.$route.query.articlerPath.replace(
      "http://localhost:3800",
      ""
    );
    let { data } = await this.$api.get("/v1" + articlerPath);
    const Viewer = toastui.Editor;
    const viewer = new Viewer({
      el: document.querySelector("#viewer"),
      height: "100%",
      initialValue: data,
    });
    viewer.setMarkdown(data)
    // viewer.invoke("setMarkdown", data);
    document.querySelectorAll("pre code").forEach((block) => {
      hljs.highlightBlock(block);
    });
  },
};
</script>
<style lang="scss" scoped>
/deep/#viewer {
  .tui-editor-contents {
    font-size: 14px;
    background: #ffffff;
    h1,
    h2,
    h3,
    h4,
    h5,
    h6 {
      //  background: #F7F4F4;
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