<template>
  <el-row>
    <el-row class="comment-wrapper">
      <el-col class="textarea-container">
        <el-input
          placeholder="发一条友善的评论"
          type="textarea"
          resize="none"
          v-model="textarea"
          ref="textarea"

        >
        </el-input>
        <el-popover placement="top-start" trigger="click">
          <div
            @click="checkout($event)"
            :style="enojiStyle"
             v-html="emojiChars"
          ></div>
          <el-button type="primary" style="font-size: 30px" slot="reference"
            >😁</el-button
          >
        </el-popover>
        <el-button type="primary" @click="change" class="comment-submit">发表评论</el-button>
      </el-col>
    </el-row>
    <Replay></Replay>
  </el-row>
</template>
<script>
const EmojiConvertor = require("emoji-js");
const emoji = new EmojiConvertor();
emoji.replace_mode = "unified";
import emojiArr from "./array";
import Replay from "./Reply";
export default {
  components: {
    Replay,
  },
  data() {
    return {
      textarea: "",
      Showemoji: false,
      emojiChars: "",
      enojiStyle: {
        fontSize: "20px",
        width: "280px",
        cursor: "pointer",
      },
    };
  },
  mounted() {
   emojiArr.forEach((item, index) => {
      this.emojiChars += `<span data-idnex=${index} >${emoji.replace_colons(
        item.char
      )}</span>`;
    });
  },
  methods: {
    checkout(event) {
      let index = event.target.getAttribute("data-idnex");
      if (!index) return;
      this.___interText(
        document.querySelector("textarea"),
        emojiArr[index].char
      );
    },
    ___interText(textarea, str) {
      let tclen = textarea.value.length;
      textarea.focus();
      if (typeof document.selection != "undefined") {
        document.selection.createRange().text = str;
         
      } else {
        textarea.value =
          textarea.value.substr(0, textarea.selectionStart) +
          str +
          textarea.value.substring(textarea.selectionStart, tclen);
         
      }
  
    },
    async change() {
      let articleId = this.$route.query.uuid
      let content = document.querySelector("textarea").value
      let result = await this.$comment(articleId,content);
      console.log(result)
    },
  },
};
</script>

<style lang="scss" scoped>
.textarea-container {
  width: 100%;
  height: 60px;
  margin: 10px 0px;
  display: flex;
  .el-button {
    height: 100%;
  }
  .el-textarea {
    width: 90%;
    margin-right: 6px;
    /deep/ .el-textarea__inner {
      height: 100%;
      font-size: 18px;
      color: black;
    }
  }
  .comment-submit {
    margin-left: 6px;
    height: 100%;
    flex: 0 1 120px;
  }
}
</style>