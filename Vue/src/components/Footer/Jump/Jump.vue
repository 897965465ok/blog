<template>
  <el-row class="jump-top">
    <p class="el-icon-caret-top" @click="fly"></p>
  </el-row>
</template>
<script>
let scrollToBottom = {
  getScrollTop: function () {
    let scrollTop = 0,
      bodyScrollTop = 0,
      documentScrollTop = 0;
    if (document.body) {
      bodyScrollTop = document.body.scrollTop;
    }
    if (document.documentElement) {
      documentScrollTop = document.documentElement.scrollTop;
    }
    scrollTop =
      bodyScrollTop - documentScrollTop > 0 ? bodyScrollTop : documentScrollTop;
    return scrollTop;
  },
  getScrollHeight: function () {
    let scrollHeight = 0,
      bodyScrollHeight = 0,
      documentScrollHeight = 0;
    if (document.body) {
      bodyScrollHeight = document.body.scrollHeight;
    }
    if (document.documentElement) {
      documentScrollHeight = document.documentElement.scrollHeight;
    }
    scrollHeight =
      bodyScrollHeight - documentScrollHeight > 0
        ? bodyScrollHeight
        : documentScrollHeight;
    return scrollHeight;
  },
  getClientHeight: function () {
    let windowHeight = 0;
    if (document.compatMode == "CSS1Compat") {
      windowHeight = document.documentElement.clientHeight;
    } else {
      windowHeight = document.body.clientHeight;
    }
    return windowHeight;
  },
  onScrollEvent: function (succeed, fail) {
    let This = this;
    window.addEventListener("scroll", () => {
      if (
        This.getScrollTop() + This.getClientHeight() >=
        This.getScrollHeight()
      ) {
        typeof succeed == "function" && succeed.call(this);
      } else {
        typeof fail == "function" && fail.call(this);
      }
    });
  },
};

let succeed = () =>
  (document.querySelector(".jump-top").style.visibility = "visible");
let fail = () =>
  (document.querySelector(".jump-top").style.visibility = "hidden");
succeed = Vue.prototype.$debounce(succeed, 500, true);
fail = Vue.prototype.$debounce(fail, 500, true);
scrollToBottom.onScrollEvent(succeed, fail);
export default {
  methods: {
    fly() {
      let interval = setInterval(() => {
        //   document.documentElement.scrollTop  获取滚动条
        if (document.documentElement.scrollTop == 0) {
          clearInterval(interval);
        }
        window.scrollTo(0, document.documentElement.scrollTop - 5);
      });
    },
  },
};
</script>
<style lang="scss" scoped>
.jump-top {
  visibility: hidden;
  position: fixed;
  bottom: 10%;
  right: 5%;
  display: flex;
  justify-content: center;
  align-items: center;
  p {
    cursor: pointer;
    border: #999 solid 1px;
    border-radius: 50%;
    font-size: 40px;
    height: 40px;
    width: 40px;
    color: cadetblue;
  }
}
</style>