# BFC 块级格式化

# BFC 块级格式化


### BFC触发方式


1. 根元素，即HTML标签
1. float的值不是none。
1. position的值不是static或者relative。
1. display的值是inline-block、table-cell、flex、table-caption或者inline-flex
1. overflow的值不是visible



### 效果


1. 内部的Box会在垂直方向，一个接一个地放置。
1. Box垂直方向的距离由margin决定。属于同一个BFC的两个相邻Box的margin会发生重叠
1. 每个元素的margin box的左边， 与包含块border box的左边相接触(对于从左往右的格式化，否则相反)。即使存在浮动也是如此。
1. BFC的区域不会与float box重叠。
1. BFC就是页面上的一个隔离的独立容器，容器里面的子元素不会影响到外面的元素。反之也如此。
1. 计算BFC的高度时，浮动元素也参与计算
