webpackJsonp([5],{"Zk/u":function(t,e){},iAFL:function(t,e,s){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var n={name:"Chat",data:function(){return{ws:null,content:"",messageGroup:[]}},methods:{SocketInit:function(){this.ws=function(t){var e=new WebSocket("ws://localhost:3800/ws");function s(t){console.log("webSocket 打开",t),console.log(e.readyState)}function n(t){console.log("webSocket 关闭",t),console.log(e.readyState)}function o(t){console.log("webSocket 错误",t),console.log(e.readyState)}return e.addEventListener("open",s,!1),e.addEventListener("close",n,!1),e.addEventListener("error",o,!1),e.addEventListener("message",t,!1),e}(this.getMessage)},sendMessage:function(t){this.ws.send(t),this.content=""},getMessage:function(t){console.log(t),this.messageGroup.push(t.data)}}},o={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("el-row",{staticClass:"chat"},[n("el-row",{staticClass:"header"},[n("header",{staticClass:"title"},[t._v("\n      聊天室\n    ")])]),t._v(" "),n("el-row",[n("section",{staticClass:"content"},t._l(t.messageGroup,function(e,o){return n("div",{key:o,staticClass:"user-wrapper"},[n("div",{staticClass:"time"},[t._v("\n          "+t._s(new Date)+"\n        ")]),t._v(" "),n("div",{staticClass:"user-info"},[n("div",{staticClass:"user-picture"},[n("img",{staticClass:"el-image",attrs:{src:s("J9T8")}})]),t._v(" "),n("div",{staticClass:"user-nessage"},[t._v(t._s(e))])])])}),0)]),t._v(" "),n("el-row",[n("footer",{staticClass:"footer"},[n("div",{staticClass:"button-tool"},[n("span",{staticClass:"iconfont icon-shijian"}),t._v(" "),n("span",{staticClass:"iconfont icon-shijian"}),t._v(" "),n("span",{staticClass:"iconfont icon-shijian"}),t._v(" "),n("span",{staticClass:"iconfont icon-shijian"})]),t._v(" "),n("textarea",{directives:[{name:"model",rawName:"v-model",value:t.content,expression:"content"}],staticClass:"input-Content",domProps:{value:t.content},on:{input:function(e){e.target.composing||(t.content=e.target.value)}}}),t._v(" "),n("div",{staticClass:"button-wrapper"},[n("button",{on:{click:t.SocketInit}},[t._v("链接")]),t._v(" "),n("button",[t._v("关闭")]),t._v(" "),n("button",{on:{click:function(e){return t.sendMessage(t.content)}}},[t._v("发送")])])])])],1)},staticRenderFns:[]};var a=s("VU/8")(n,o,!1,function(t){s("Zk/u")},"data-v-9a706196",null);e.default=a.exports}});