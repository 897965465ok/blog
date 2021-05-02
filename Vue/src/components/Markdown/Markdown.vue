<template>
  <Markdown
    class="markdown"
    v-model="markdown"
    isPreview
    :height="1200"
  >
  </Markdown>
</template>
<script>
import Markdown from "vue-meditor";
export default {
  name: "About",
  components: {
    Markdown,
  },
  data() {
    return {
      markdown: "",
    };
  },
  async created() {
    let articlerPath = this.$route.query.articlerPath.replace(
      "http://localhost:3800",
      ""
    );
    let { data } = await this.$api.get("/v1" + articlerPath);
    this.markdown = data;
  },
};
</script>

<style lang="scss" scoped>
.markdown {
  width: 100%;
}
</style>