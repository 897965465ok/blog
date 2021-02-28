# javascript5种继承方式

```
//继承方式一
function myobject(){

}
myobject.prototype={
	constructor:object,
	hasOwnProperty:function(){

	}
}
function myEventTarget(){

}
myEventTarget.prototype= new myobject();
myEventTarget.prototype.addEventListener=function(){

}
function mynode(){
}
mynode.prototype= new myEventTarget();
mynode.prototype.creadteElement=function(){

}
var nodnode = new mynode();
```

---


---

var n = new B;

```
<hr>
```

继承模式4

function A(){

this.x=100;

}

A.prototype.getX=function(){

console.log(this.x);

}

function B(){

A.call(this)

}

B.prototype = new A;

B.prototype.constructor = B;

var n = new B;

n.getX();

```
<hr>
```

继承模式5

function A(){

this.x=100;

}

A.prototype.getX=function(){

console.log(this.x);

}

function B(){

A.call(this)

}

B.prototype = Object.create(A.prototype);

B.prototype.constructor = B;

var n = new B;

n.getX();

