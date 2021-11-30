<template>
  <el-row>
    <el-form
      ref="loginForm"
      :model="form"
      :rules="rules"
      label-width="80px"
      class="login-box"
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
      <!-- <el-form-item label="电子邮箱" prop="email">
        <el-input type="txt" placeholder="请输入邮箱" v-model="form.email" />
      </el-form-item> -->
      <el-form-item>
        <el-button type="primary" @click="onSubmit('loginForm')"
          >登录</el-button
        >
      </el-form-item>
    </el-form>
  </el-row>
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
        email: ""
      },

      // 表单验证，需要在 el-form-item 元素中增加 prop 属性
      rules: {
        username: [
          { required: true, message: "账号不可为空", trigger: "blur" }
        ],
        password: [{ required: true, message: "密码不可为空", trigger: "blur" }]
        // email: [
        //   { required: true, message: "邮箱不能为空", trigger: "blur" },
        //   {
        //     pattern: /\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*/,
        //     message: "请正确输入邮箱",
        //   },
        // ],
      }
    };
  },
  methods: {
    onSubmit(formName) {
      // 为表单绑定验证功能
      this.$refs[formName].validate(async valid => {
        if (valid) {
          let { username: userName, password, email } = this.form;
          let { data: result } = await this.$api.post(
            "/v1/login",
            qs.stringify({ userName, password, email }),
            { headers: { "Content-Type": "application/x-www-form-urlencoded" } }
          );
          console.log(result)
          if (result.code == 200) {
            localStorage.setItem("token", "Bearer " + result.data.token);
            this.$router.push("/");
          }
        } else {
          return false;
        }
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.login-box {
  border: 1px solid #dcdfe6;
  width: 350px;
  margin: 180px auto;
  padding: 35px 35px 15px 35px;
  border-radius: 5px;
  -webkit-border-radius: 5px;
  -moz-border-radius: 5px;
  box-shadow: 0 0 25px #909399;
}

.login-title {
  text-align: center;
  margin: 0 auto 40px auto;
  color: #303133;
}
</style>
