# flex布局

## flex-direction 水平方向

 ```
 flex-direction 水平方向
 1. row  左边开始排重小到大
 2. row-reverse 右边开始排重大到小
 3. column  上开始排重小到大 
 4. column-reverse 重下开始排重大到小
```

## flex-wrap 排列方式

```
flex-wrap 换行
1. nowrap 默认不换行
2. wrap  换行
3. wrap-reverse 换行不过是有顺序的是后面的先上去 123456  会变成654321
```

## flex-flow  direction 和wrap 间写方式

```
flex-flow   direction 和wrap 简写 
flex-flow: row wrap ; //重左到右  换行
```

## justify-content 主轴对齐方式

```
justify-content 主轴对齐方式
flex-start（默认值）左对齐
flex-end右对齐
center 水平居中
space-between 左右两边贴近盒子中间平均分配空白间隔
space-around 平均空白 前后都有平均空白
```

## align-items  垂直对齐

```
align-items  垂直对齐
1. stretch  在没有给高属性的时候自动拉伸 等于height:100%;
2. flex-start //上对齐
3. flex-end  //下对齐
4. center 垂直居中
5. baseline 根据第一个子元素字体高度来对齐
```

## align-content 多行(按行算),必须有flex-flow:row wrap;属性

```
align-content (多条轴对齐方式) (按行算),
必须flex-direction:row; ,flex-wrap:wrap;
flex-start  重左开始 
flex-end 重右开始
center   上下居中
space-between 左右两边贴近盒子中间平均分配空白间隔
space-around 平均空白 前后都有平均空白
```

## flex子元素属性

```
order 元素的排列顺序,子元素排队属性,值越大排最后面 默认是0;
flex-basis 子元素元素占据主轴的空间  默认是px 覆盖width属性; 
flex-grow :倍数默认是0 ,如果设置可以设置为2意思比其他元素大1倍;
flex-shrink 元素的缩小比例  空间按比例缩小默认是1 //目前不知道;
```

## 其他

```
flex:(flex-row,flex-shrink,flex-basis)的简写。
默认auto (0,1,auto) 如果设为none就是这样(0,0,auto)。 
align-self 设定自身的对齐方式 自己的对齐不算其他子元素
```

那我们就从使用 flex 如何实现三栏布局开始吧。

想要实现三栏等高布局，且两边的侧栏宽度固定而中间一栏占用剩余的空间，如下代码就足够了：

```
<style>
  section {display: flex;}
  .left-side,
  .right-side {width: 200px;}
  .content {flex-grow: 1;}
</style>
<section>
  <div class="left-side"></div>
  <div class="content"></div>
  <div class="right-side"></div>
</section>
```

其中 section 元素的宽度将会像 block 元素一样尽量的宽，对外面的元素来说，它的行为很像一个 block 块。三个元素会从左往右占据父元素的空间（这很显然）。左右侧边栏的宽度都是 200px，中间 .content 元素的宽度将会占据 section 元素的剩余宽度。

另外，section 的高度会自动被最高的一个子元素撑开，同时其它子元素的高度也会被拉到跟 section 元素一样高，而如果给 section 元素设置了高度，而所有子元素的高度设置为 auto ，所有的子元素也都会自动跟父元素一样高，这简直就是在传统布局中做梦都想要的功能！

总之，在高度方面，flex 的表现是相当符合直觉的。

另外，如果不给 flex 子元素设置宽度和 flex-grow，它会尽量的窄。

## flex-grow 的计算方式

上面 demo 中最值得注意的是 .content 元素的 flex-grow 属性，设置为 1 它就可以占满水平剩余空间。这也是本文的重点：讲清 flex-grow 与 flex-shrink 属性的详细计算方式。

flex-grow 属性决定了父元素在空间分配方向上还有剩余空间时，如何分配这些剩余空间。其值为一个权重（也称扩张因子），默认为 0（纯数值，无单位），剩余空间将会按照这个权重来分配。

比如剩余空间为 x，三个元素的 flex-grow 分别为 a，b，c。设 sum 为 a + b + c。那么三个元素将得到剩余空间分别是 x _ a / sum, x _ b / sum, x * c / sum，是为权重也。

举个例子：

父元素宽度 500px，三个子元素的 width 分别为 100px，150px，100px。

于是剩余空间为 150px

三个元素的 flex-grow 分别是 1，2，3，于是 sum 为 6

则三个元素所得到的多余空间分别是：

- 150 * 1 / 6 = 25px
- 150 * 2 / 6 = 50px
- 150 * 3 / 6 = 75px

三个元素最终的宽度分别为 125px，200px，175px。

- 100px + 25px = 125px
- 150px + 50px = 200px
- 100px + 75px = 175px

可以打开这个 [demo](http://link.zhihu.com/?target=https%3A//xieranmaya.github.io/blog/flex-grow-shrink.html)（下文中所有的 demo 都在这个页面） 然后用开发工具查看一下。注意不要用截图工具量，可能量不准，因为高分屏和放大等诸多因素都会影响测量结果。

然而！不止这些，还有一种情况：

当所有元素的 flex-grow 之和小于 1 的时候（注意是 1，也就是说每个元素的 flex-grow 都是一个小数如 0.2 这样的），上面式子中的 sum 将会使用 1 来参与计算，而不论它们的和是多少。也就是说，当所有的元素的 flex-grow 之和小于 1 的时候，剩余空间不会全部分配给各个元素。

实际上用来分配的空间是 sum(flex-grow) / 1 * 剩余空间，这些用来分配的空间依然是按 flex-grow 的比例来分配。

还是上面一个例子，但是三个元素的 flex-grow 分别是 0.1，0.2，0.3，那么计算公式将变成下面这样：

- 150 * 0.1 / 1 = 15px
- 150 * 0.2 / 1 = 30px
- 150 * 0.3 / 1 = 45px

150px - 15px - 30px - 45px = 60px，即还有 60px 没有分配给任何子元素。

三个元素的最终宽度分别为：

- 100px + 15px = 115px
- 150px + 30px = 180px
- 100px + 45px = 145px

如上所述即是 flex-grow 的计算方式。

另外，flex-grow 还会受到 max-width 的影响。如果最终 grow 后的结果大于 max-width 指定的值，max-width 的值将会优先使用。同样会导致父元素有部分剩余空间没有分配。

## flex-shrink 的计算方式

前文已经说到，flex 几乎一次性解决了前端布局的所有问题。

那么既然可以在空间有多余时把多余空间分配给各个子元素，当然也可以在空间不够时让各个子元素收缩以适应有限的空间了。

这就是 flex-shrink 属性的作用。

你可能会觉得 flex-shrink 的计算方式跟 flex-grow 很类似，然而事情并没有这么简单。

flex-shrink 属性定义空间不够时各个元素如何收缩。其值默认为 1。很多文章对此基本是一笔带过：“flex-shrink 属性定义了元素的收缩系数”，根本就不说它具体是怎么计算的。

flex-shrink 定义的仅仅只是元素宽度变小的一个权重分量。

每个元素具体收缩多少，还有另一个重要因素，即它本身的宽度。

举个例子：

父元素 500px。三个子元素分别设置为 150px，200px，300px。

三个子元素的 flex-shrink 的值分别为 1，2，3。

首先，计算子元素溢出多少：150 + 200 + 300 - 500 = -150px。

那这 -150px 将由三个元素的分别收缩一定的量来弥补。

具体的计算方式为：每个元素收缩的权重为其 flex-shrink 乘以其宽度。

所以总权重为 1 _ 150 + 2 _ 200 + 3 * 300 = 1450

三个元素分别收缩：

- 150 _ 1(flex-shrink) _ 150(width) / 1450 = -15.5
- 150 _ 2(flex-shrink) _ 200(width) / 1450 = -41.4
- 150 _ 3(flex-shrink) _ 300(width) / 1450 = -93.1

三个元素的最终宽度分别为：

- 150 - 15.5 = 134.5
- 200 - 41.4 = 158.6
- 300 - 93.1 = 206.9

同样，当所有元素的 flex-shrink 之和小于 1 时，计算方式也会有所不同：

此时，并不会收缩所有的空间，而只会收缩 flex-shrink 之和相对于 1 的比例的空间。

还是上面的例子，但是 flex-shrink 分别改为 0.1，0.2，0.3。

于是总权重为 145（正好缩小 10 倍，略去计算公式）。

三个元素收缩总和并不是 150px，而是只会收缩 150px 的 (0.1 + 0.2 + 0.3) / 1 即 60% 的空间：90px。

每个元素收缩的空间为：

- 90 _ 0.1(flex-shrink) _ 150(width) / 145 = 9.31
- 90 _ 0.2(flex-shrink) _ 200(width) / 145 = 24.83
- 90 _ 0.3(flex-shrink) _ 300(width) / 145 = 55.86

三个元素的最终宽度分别为：

- 150 - 9.31 = 140.69
- 200 - 24.83 = 175.17
- 300 - 55.86 = 244.14

当然，类似 flex-grow，flex-shrink 也会受到 min-width 的影响。
