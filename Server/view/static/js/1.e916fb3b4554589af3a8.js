webpackJsonp([1],{"8np1":function(t,e){},J9T8:function(t,e,r){t.exports=r.p+"static/img/jerry.7a6f41e.png"},Jz7M:function(t,e){},"j3r+":function(t,e,r){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var a=r("4YfN"),s=r.n(a),n=r("lC5x"),i=r.n(n),c=r("J0Oq"),l=r.n(c),o=r("SJI6"),u={name:"Carousel",data:function(){return{carouselUrl:[]}},computed:Object(o.mapState)({pictures:function(t){return t.pictures.length?t.pictures:[]}}),watch:{pictures:function(t){var e=Math.floor(194*Math.random()+1);return this.carouselUrl=t.slice(e,e+5),t}}},f={render:function(){var t=this.$createElement,e=this._self._c||t;return e("el-carousel",{staticClass:"carousel",attrs:{"indicator-position":"none",interval:3e3,height:"300px",type:this.type}},this._l(this.carouselUrl,function(t,r){return e("el-carousel-item",{key:r},[e("el-image",{attrs:{src:t,height:"300px"}})],1)}),1)},staticRenderFns:[]};var v,p=r("C7Lr")(u,f,!1,function(t){r("xHCS")},null,null).exports,m=r("xxrf"),h=r("lRwf"),d=r.n(h),_=void 0;(v=!0,function(){v&&(_=d.a.prototype.$loading(),v=!1)})();var g={data:function(){return{loading:null,favorites:[{userName:"动漫之家",id:120,link:"http://www.imomoe.ai"}]}},created:function(){var t=this;return l()(i.a.mark(function e(){return i.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.t0=t,e.next=3,t.$GetUrl();case 3:return e.t1=e.sent,e.next=6,e.t0.getPictures.call(e.t0,e.t1);case 6:return e.t2=t,e.next=9,t.$api.get("v1/tags");case 9:return e.t3=e.sent.data.result,e.t4={tags:e.t3},e.next=13,e.t2.setRecommen.call(e.t2,e.t4);case 13:return e.t5=t,e.next=16,t.$api.get("v1/articles");case 16:return e.t6=e.sent.data.result,e.next=19,e.t5.changeAll.call(e.t5,e.t6);case 19:_.close();case 20:case"end":return e.stop()}},e,t)}))()},computed:Object(o.mapState)({articles:function(t){return t.articlers.length?t.articlers:[]}}),components:{Carousel:p,Profile:m.a},methods:s()({},Object(o.mapActions)(["changeAll","getPictures","setRecommen"]))},C={render:function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("el-row",{staticClass:"main"},[r("el-row",{staticClass:"main-content"},[r("el-col",{staticClass:"content-left",attrs:{span:16}},[r("el-row",{staticClass:"card-wrapper"},[r("h3",{staticClass:"sub-title"},[t._v("#最新文章")]),t._v(" "),t._l(t.articles.slice(0,4),function(t){return r("Articler",{key:t.uuid,attrs:{article:t}})}),t._v(" "),r("h3",{staticClass:"sub-title"},[t._v("#博主推荐")]),t._v(" "),r("Carousel"),t._v(" "),t._l(t.articles.slice(4,9),function(t){return r("Articler",{key:t.uuid,attrs:{article:t}})})],2)],1),t._v(" "),r("el-col",{staticClass:"content-right",attrs:{span:8}},[r("el-row",{staticClass:"main-header"},[r("el-col",{staticClass:"header-right"},[r("Profile")],1)],1),t._v(" "),r("h3",{staticClass:"sub-title"},[t._v("#特别推荐")]),t._v(" "),t._l(t.articles.slice(9,16),function(t){return r("ReArticle",{key:t.uuid,attrs:{article:t}})}),t._v(" "),r("h3",{staticClass:"sub-title"},[t._v("#小程序")]),t._v(" "),r("el-row",[r("el-image",{attrs:{src:"https://tva2.sinaimg.cn/large/87c01ec7gy1frmmz605z4j21kw0w0qvh.jpg",alt:"二维码图片",lazy:""}})],1),t._v(" "),r("h3",{staticClass:"sub-title"},[t._v("#个人收藏网站")]),t._v(" "),r("el-row",{staticClass:"favorites-link"},[r("waterfall",{attrs:{col:3,data:t.favorites,gutterWidth:10}},t._l(t.favorites,function(e,a){return r("div",{key:a,staticClass:"favorites-item"},[r("el-button",[t._v("\n              "+t._s(e.userName)+"\n            ")])],1)}),0)],1)],2)],1)],1)},staticRenderFns:[]};var w=r("C7Lr")(g,C,!1,function(t){r("Jz7M")},"data-v-21ef752b",null);e.default=w.exports},xHCS:function(t,e){},xxrf:function(t,e,r){"use strict";var a=r("SJI6"),s={data:function(){return{icon:[{title:"Github",icon:"icon-huaban88",link:"https://github.com/897965465ok"},{title:"QQ",icon:"icon-qq",link:"http://wpa.qq.com/msgrd?v=3&uin=897965465&site=qq&menu=yes"}]}},computed:Object(a.mapState)({recommen:function(t){if(t.recommen)return t.recommen}})},n={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-row",{staticClass:"profile"},[a("el-row",{staticClass:"user"},[a("h3",[t._v("#Jiang")])]),t._v(" "),a("el-row",{staticClass:"pictuer"},[a("img",{staticClass:"el-image",attrs:{src:r("J9T8")}})]),t._v(" "),a("el-row",{staticClass:"tag"},[a("ul",t._l(t.icon,function(e,r){return a("li",{key:r},[a("a",{attrs:{href:e.link}},[a("i",{staticClass:"iconfont",class:e.icon}),t._v(" "),a("p",[t._v(t._s(e.title))])])])}),0)])],1)},staticRenderFns:[]};var i=r("C7Lr")(s,n,!1,function(t){r("8np1")},"data-v-79a037c7",null);e.a=i.exports}});