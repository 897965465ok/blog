<template>
    <div
      class="markdwon"
      v-html="compileMarkDown(markdown)"
      v-highlight
    >
    </div>
</template>

<script>

import showdown from "showdown";
import hljs from "highlight.js"
import 'highlight.js/styles/atom-one-dark-reasonable.css';
var converter = new showdown.Converter();
//表格显示
converter.setOption("tables", true);
export default {
   // 定义自定义指令 v-highlight 代码高亮
  directives: {
    highlight: {
      update(el) {
        let blocks = el.querySelectorAll("pre code");
        blocks.forEach((block) => {
          hljs.highlightBlock(block);
        });
      },
    },
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
  methods: {
     // 转换markdown语法为html语法
    compileMarkDown(val) {
      return converter.makeHtml(val);
    },
  },
};
</script>
<style lang="scss" scoped>
.markdwon{
  background: #282C34;
  // font-size: 18px;
}
</style>





