webpackJsonp([5],{T3Ax:function(e,r){},mJTh:function(e,r,t){"use strict";Object.defineProperty(r,"__esModule",{value:!0});var a=t("Xxa5"),n=t.n(a),s=t("exGp"),o=t.n(s),l=t("mw3O"),i=t.n(l),u={name:"Login",data:function(){return{form:{username:"",password:"",email:""},rules:{username:[{required:!0,message:"账号不可为空",trigger:"blur"}],password:[{required:!0,message:"密码不可为空",trigger:"blur"}]}}},methods:{onSubmit:function(e){var r,t=this;this.$refs[e].validate((r=o()(n.a.mark(function e(r){var a,s,o,l,u,m;return n.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:if(!r){e.next=9;break}return a=t.form,s=a.username,o=a.password,l=a.email,e.next=4,t.$api.post("/v1/login",i.a.stringify({userName:s,password:o,email:l}),{headers:{"Content-Type":"application/x-www-form-urlencoded"}});case 4:u=e.sent,200==(m=u.data).code&&(localStorage.setItem("token","Bearer "+m.data.token),t.$router.push("/")),e.next=10;break;case 9:return e.abrupt("return",!1);case 10:case"end":return e.stop()}},e,t)})),function(e){return r.apply(this,arguments)}))}}},m={render:function(){var e=this,r=e.$createElement,t=e._self._c||r;return t("el-row",[t("el-form",{ref:"loginForm",staticClass:"login-box",attrs:{model:e.form,rules:e.rules,"label-width":"80px"}},[t("h3",{staticClass:"login-title"},[e._v("欢迎登录")]),e._v(" "),t("el-form-item",{attrs:{label:"账号",prop:"username"}},[t("el-input",{attrs:{type:"text",placeholder:"请输入账号"},model:{value:e.form.username,callback:function(r){e.$set(e.form,"username",r)},expression:"form.username"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"密码",prop:"password"}},[t("el-input",{attrs:{type:"password",placeholder:"请输入密码"},model:{value:e.form.password,callback:function(r){e.$set(e.form,"password",r)},expression:"form.password"}})],1),e._v(" "),t("el-form-item",[t("el-button",{attrs:{type:"primary"},on:{click:function(r){return e.onSubmit("loginForm")}}},[e._v("登录")])],1)],1)],1)},staticRenderFns:[]};var p=t("VU/8")(u,m,!1,function(e){t("T3Ax")},"data-v-ddb6febe",null);r.default=p.exports}});