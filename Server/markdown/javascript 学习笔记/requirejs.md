# requirejs

requirejs.config({

paths:{

jquery:"./jquery-1.7.2.min"

}

});

requirejs(["jquery"],function(![](https://g.yuque.com/gr/latex?)%7B%0A#card=math&code=%29%7B%0A)("body").css("background","red");

})

// 这个是其他包里面用到的

define(["jquery"],function ($) {

return{

ooxx:function{};

}

})
