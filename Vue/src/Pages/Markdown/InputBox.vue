<template>
  <el-row>
    <el-row class="comment-wrapper   h-16  ">
      <el-col class="textarea-container  h-full ">
        <el-input
          :placeholder="userName"
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
        <el-button type="primary" @click="change" class="comment-submit"
          >发表评论</el-button
        >
      </el-col>
    </el-row>
  </el-row>
</template>

<script>
const emoji = new EmojiConvertor();
emoji.replace_mode = "unified";
import emojiArr from "@/Pages/Markdown/array";
export default {
  props: {
    reply: {
      type: Object,
      default: null
    },
    userName: {
      type: String,
      default: "发一条友善的评论"
    },
    ReplyUser: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      textarea: "",
      Showemoji: false,
      emojiChars: "",
      enojiStyle: {
        fontSize: "20px",
        width: "280px",
        cursor: "pointer"
      },
      uuid: ""
    };
  },
  mounted() {
    emojiArr.forEach((item, index) => {
      this.emojiChars += `<span data-idnex=${index} >${emoji.replace_colons(
        item.char
      )}</span>`;
    });
    if (this.reply) {
      // this.userName = this.reply.User.name;
      this.uuid = this.reply.uuid;
      this.ReplyUser = true;
    } else {
      this.uuid = this.$route.query.uuid;
    }
  },
  methods: {
    checkout(event) {
      let index = event.target.getAttribute("data-idnex");
      if (!index) return;
      this.selectEmoji(
        this.$refs.textarea.$el.querySelector("textarea"),
        emojiArr[index].char
      );
    },
    selectEmoji(elInput, emoji) {
      let startPos = elInput.selectionStart; // input 第0个字符到选中的字符
      let endPos = elInput.selectionEnd; // 选中的字符到最后的字符
      if (startPos === undefined || endPos === undefined) return;
      let txt = elInput.value;
      // 将表情添加到选中的光标位置
      let result = txt.substring(0, startPos) + emoji + txt.substring(endPos);
      elInput.value = result; // 赋值给input的value
      // 重新定义光标位置
      elInput.focus();
    },
    async change() {
      let content = this.$refs.textarea.$el.querySelector("textarea").value;
      let result = await this.$comment(this.uuid, content, this.ReplyUser);
      console.log(result);
      // 测试方法
      // await this.$shouldJson(this.uuid, content);
    }
  }
};
</script>

<style lang="scss" scoped>
.textarea-container {
  width: 100%;
  height: 100%;
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
