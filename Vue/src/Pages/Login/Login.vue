<template>
  <div
    class="
      container
      flex
      h-screen
      w-full
      bg-white
      justify-center
      content-center
      shadow-sm
    "
  >
    <el-form
      ref="loginForm"
      :model="form"
      :rules="rules"
      class="bg-white w-4/12 h-4/12"
    >
      <h3 class="login-title">欢迎登录</h3>
      <el-form-item label="账号" prop="username">
        <el-input
          type="text"
          placeholder="请输入账号"
          v-model="form.username"
        />
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input
          type="password"
          placeholder="请输入密码"
          v-model="form.password"
        />
      </el-form-item>
      <el-form-item class="flex justify-between">
        <el-button @click="onSubmit('loginForm')">登录</el-button>
        <el-button class="iconfont icon-huaban88" @click="oauth()"></el-button>
        <el-button class="iconfont icon-mayun" @click="oauth()"></el-button>
        <el-button class="iconfont icon-qq" @click="oauth()"></el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import qs from "qs";
export default {
  name: "Login",
  data() {
    return {
      form: {
        username: "",
        password: "",
        email: "",
        url: "",
      },

      // 表单验证，需要在 el-form-item 元素中增加 prop 属性
      rules: {
        username: [
          { required: true, message: "账号不可为空", trigger: "blur" },
        ],
        password: [
          { required: true, message: "密码不可为空", trigger: "blur" },
        ],
        // email: [
        //   { required: true, message: "邮箱不能为空", trigger: "blur" },
        //   {
        //     pattern: /\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*/,
        //     message: "请正确输入邮箱",
        //   },
        // ],
      },
    };
  },

  methods: {
    onSubmit(formName) {
      // 为表单绑定验证功能
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          let { username: userName, password, email } = this.form;
          let { data: result } = await this.$api.post(
            "/v1/login",
            qs.stringify({ userName, password, email }),
            { headers: { "Content-Type": "application/x-www-form-urlencoded" } }
          );
          console.log(result);
          if (result.code == 200) {
            localStorage.setItem("token", "Bearer " + result.data.token);
            this.$router.push("/");
          }
        } else {
          return false;
        }
      });
    },
    openwindow(url, name, iWidth, iHeight) {
      var url; //转向网页的地址;
      var name; //网页名称，可为空;
      var iWidth; //弹出窗口的宽度;
      var iHeight; //弹出窗口的高度;
      var iTop = (window.screen.availHeight - 30 - iHeight) / 2; //获得窗口的垂直位置;
      var iLeft = (window.screen.availWidth - 10 - iWidth) / 2; //获得窗口的水平位置;
      return window.open(
        url,
        name,
        "height=" +
          iHeight +
          ",,innerHeight=" +
          iHeight +
          ",width=" +
          iWidth +
          ",innerWidth=" +
          iWidth +
          ",top=" +
          iTop +
          ",left=" +
          iLeft +
          ",toolbar=no,menubar=no,scrollbars=auto,resizeable=no,location=no,status=no"
      );
    },
    oauth() {
      this.$nextTick(() => {
        let client_id =
          "a480ace6ca2d676d30f1c932936da15c256db9692834a5107ca910586470ccc1";
        let redirect_uri = "http://www.mrjiang.work/v1/oauth";
        let url = `https://gitee.com/oauth/authorize?client_id=${client_id}&redirect_uri=${redirect_uri}&response_type=code`;
        let newWindows = this.openwindow(url, "", "500", "500");
        window.addEventListener("message",event=> {
          if (event.origin === "http://www.mrjiang.work") {
   
            console.log(event.data)
                     newWindows.close();
          }
        });
      });
    },
  },
};
</script>
<style lang="scss" scoped>
@import "../../assets/icon/iconfont.css";
.login-title {
  text-align: center;
  margin: 0 auto 40px auto;
  color: #303133;
}
</style>
