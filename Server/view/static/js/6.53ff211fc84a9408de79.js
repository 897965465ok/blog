webpackJsonp([6],{lb4t:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var r=a("lC5x"),n=a.n(r),s=a("J0Oq"),i=a.n(s),c={name:"Gossip",data:function(){return{articles:[]}},created:function(){var t=this;return i()(n.a.mark(function e(){var a,r;return n.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,t.$api.get("/v1/gossip");case 2:a=e.sent,r=a.data,t.articles=r.articles;case 5:case"end":return e.stop()}},e,t)}))()}},o={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-row",{staticClass:"gossip"},[a("el-row",t._l(t.articles,function(e){return a("el-col",{key:e.id,staticClass:"card",attrs:{span:8}},[a("div",{staticClass:"picture"}),t._v(" "),a("h4",[t._v(t._s(e.paragraph))])])}),1),t._v(" "),a("el-col",{staticClass:"pagination"},[a("el-pagination",{attrs:{layout:"prev, pager, next",total:1e3}})],1)],1)},staticRenderFns:[]};var l=a("C7Lr")(c,o,!1,function(t){a("p3oK")},"data-v-680ea6f3",null);e.default=l.exports},p3oK:function(t,e){}});