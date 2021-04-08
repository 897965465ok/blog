<template>
  <div class="markdown">
    <mavon-editor
      v-model="content"
      :subfield="false"
      :editable="false"
      :toolbarsFlag="false"
      :boxShadow="false"
      defaultOpen="preview"
      codeStyle="vs2015"
    />
  </div>
</template>
<script>
export default {
  name: "About",
  data() {
    return {
      content: "",
    };
  },
  async created() {
    let articlerPath = this.$route.query.articlerPath.replace(
      "http://localhost:3800",
      ""
    );
    let { data } = await this.$api.get("/v1" + articlerPath);
    this.content = data;
  },
};
</script>

<style lang="scss" scoped>
.markdown {
  width: 100%;
  min-height: 80vh;
  .mavonEditor {
    width: 100%;
    height: 100%;
    background: #ffffff;
  }
}
/deep/.markdown-body pre {
  font-size: 17px;
}
</style>