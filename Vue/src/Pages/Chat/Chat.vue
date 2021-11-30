<template>
  <el-row class="chat">
    <el-row class="header">
      <header class="title">
        聊天室
      </header>
    </el-row>
    <el-row>
      <section class="content">
        <div
          class="user-wrapper"
          v-for="(item, index) in messageGroup"
          :key="index"
        >
          <div class="time">
            {{ new Date() }}
          </div>
          <div class="user-info">
            <div class="user-picture">
              <img class="el-image" src="../../assets/jerry.png" />
            </div>
            <div class="user-nessage">{{ item }}</div>
          </div>
        </div>
      </section>
    </el-row>
    <el-row>
      <footer class="footer">
        <div class="button-tool">
          <span class="iconfont icon-shijian"></span>
          <span class="iconfont icon-shijian"></span>
          <span class="iconfont icon-shijian"></span>
          <span class="iconfont icon-shijian"></span>
        </div>
        <textarea v-model="content" class="input-Content"> </textarea>
        <div class="button-wrapper">
          <button @click="SocketInit">链接</button>
          <button>关闭</button>
          <button @click="sendMessage(content)">发送</button>
        </div>
      </footer>
    </el-row>
  </el-row>
</template>
<script>
import { useWebSocket } from "./ws.js";
export default {
  name: "Chat",
  data() {
    return {
      ws: null,
      content: "",
      messageGroup: []
    };
  },

  methods: {
    SocketInit() {
      this.ws = useWebSocket(this.getMessage);
    },
    sendMessage(content) {
      this.ws.send(content);
      this.content = "";
    },
    getMessage(event) {
        console.log(event)
        this.messageGroup.push(event.data);
      
    }
  }
};
</script>

<style lang="scss" scoped>
@import "../../assets/icon/iconfont.css";
.chat {
  padding: 6px;
  border-radius: 0.5rem;
  border: #898989 solid 1px;
  background-color: #ffffff;
}
.header {
  height: 2rem;
  background-color: #22a4ff;
  text-align: center;
  border-radius: 0.5rem;
  .title {
    line-height: 2rem;
    color: #ffff;
  }
}
.content {
  margin: 6px 0px;
  height: 30rem;
  border-radius: 0.5rem;
  border: #898989 solid 1px;
  overflow-y: scroll;
  .user-wrapper {
    display: flex;
    flex-direction: column;
    margin: 6px;
    .time {
      text-align: left;
      margin: 2px;
    }
    .user-info {
      display: flex;
      .user-picture {
        .el-image {
          border-radius: 50%;
          height: 5rem;
        }
        margin: 0px 12px;
      }
      .user-nessage {
        height: 5rem;
        line-height: 5rem;
      }
    }
  }
}
.footer {
  height: 13rem;
  border-radius: 0.5rem;
  border: #898989 solid 1px;
  background-color: #ffffff;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  .input-Content {
    outline-style: none;
    resize: none;
    flex-grow: 2;
    border: none;
    background-color: #ffffff;
  }
  .button-wrapper {
    display: flex;
    justify-content: flex-end;
    button {
      margin: 6px;
      width: 6rem;
      height: 1.5rem;
      border-radius: 0.5rem;
      border: #898989 solid 1px;
      color: #ffffff;
      background-color: #22a4ff;
      outline: none;
      cursor: pointer;
    }
  }
  .button-tool {
    height: 1.5rem;
    color: #ffffff;
    background-color: #22a4ff;
  }
}
</style>
