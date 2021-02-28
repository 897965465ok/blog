# mathjax_tutorial_mdeditor

---


#Cmd Markdown 公式指导手册

标签： Tutorial

---

> 点击跳转至 [Cmd Markdown 简明语法手册](https://www.zybuluo.com/mdeditor?url=https://www.zybuluo.com/static/editor/md-help.markdown) ，立刻开始 Cmd Markdown 编辑阅读器的记录和写作之旅！


---

> 本文为 MathJax 在 Cmd Markdown 环境下的语法指引。


> Cmd Markdown 编辑阅读器支持 ![](https://g.yuque.com/gr/latex?%5CLaTeX#card=math&code=%5CLaTeX) 编辑显示支持，例如：![](https://g.yuque.com/gr/latex?%5Csum_%7Bi%3D1%7D%5En%20a_i%3D0#card=math&code=%5Csum_%7Bi%3D1%7D%5En%20a_i%3D0)，访问 [MathJax](http://meta.math.stackexchange.com/questions/5020/mathjax-basic-tutorial-and-quick-reference) 以参考更多使用方法。


> 右键点击每一个公式，选择 **[Show Math As] → [TeX Commands]** 以查看该公式的命令详情。


[TOC]

---

# 一、公式使用参考

## 1．如何插入公式

![](https://g.yuque.com/gr/latex?%5CLaTeX#card=math&code=%5CLaTeX) 的数学公式有两种：行中公式和独立公式。行中公式放在文中与其它文字混编，独立公式单独成行。

行中公式可以用如下方法表示：

:        $ 数学公式 $

独立公式可以用如下方法表示：

:        $$ 数学公式 $$

自动编号的公式可以用如下方法表示：

:    若需要手动编号，参见 [大括号和行标的使用](#14%E5%A4%A7%E6%8B%AC%E5%8F%B7%E5%92%8C%E8%A1%8C%E6%A0%87%E7%9A%84%E4%BD%BF%E7%94%A8) 。

:       \begin{equation}

数学公式

\label{eq:当前公式名}

\end{equation}

**自动编号后的公式可在全文任意处使用 `\eqref{eq:公式名}` 语句引用。**

- 例子：

```
$ J_\alpha(x) = \sum_{m=0}^\infty \frac{(-1)^m}{m! \Gamma (m + \alpha + 1)} {\left({ \frac{x}{2} }\right)}^{2m + \alpha} \text {，行内公式示例} $
```

- 
显示：$ J_\alpha(x) = \sum_{m=0}^\infty \frac{(-1)^m}{m! \Gamma (m + \alpha + 1)} {\left({ \frac{x}{2} }\right)}^{2m + \alpha} \text {，行内公式示例} $

- 
例子：


```
$$ J_\alpha(x) = \sum_{m=0}^\infty \frac{(-1)^m}{m! \Gamma (m + \alpha + 1)} {\left({ \frac{x}{2} }\right)}^{2m + \alpha} \text {，独立公式示例} $$
```

- 
显示：$$ J_\alpha(x) = \sum_{m=0}^\infty \frac{(-1)^m}{m! \Gamma (m + \alpha + 1)} {\left({ \frac{x}{2} }\right)}^{2m + \alpha} \text {，独立公式示例} $$

- 
例子：


```
在公式 \eqref{eq:sample} 中，我们看到了这个被自动编号的公式。

\begin{equation}
E=mc^2 \text{，自动编号公式示例}
\label{eq:Sample}
\end{equation}
```

- 显示：

![](https://g.yuque.com/gr/latex?%E5%9C%A8%E5%85%AC%E5%BC%8F%20%5Ceqref%7Beq%3Asample%7D%20%E4%B8%AD%EF%BC%8C%E6%88%91%E4%BB%AC%E7%9C%8B%E5%88%B0%E4%BA%86%E8%BF%99%E4%B8%AA%E8%A2%AB%E8%87%AA%E5%8A%A8%E7%BC%96%E5%8F%B7%E7%9A%84%E5%85%AC%E5%BC%8F%E3%80%82%0A#card=math&code=%E5%9C%A8%E5%85%AC%E5%BC%8F%20%5Ceqref%7Beq%3Asample%7D%20%E4%B8%AD%EF%BC%8C%E6%88%91%E4%BB%AC%E7%9C%8B%E5%88%B0%E4%BA%86%E8%BF%99%E4%B8%AA%E8%A2%AB%E8%87%AA%E5%8A%A8%E7%BC%96%E5%8F%B7%E7%9A%84%E5%85%AC%E5%BC%8F%E3%80%82%0A)

\begin{equation}

E=mc^2 \text{，自动编号公式示例}

\label{eq:sample}

\end{equation}

##2．如何输入上下标

`^` 表示上标, `_` 表示下标。如果上下标的内容多于一个字符，需要用 `{}` 将这些内容括成一个整体。上下标可以嵌套，也可以同时使用。

- 例子：

```
$$ x^{y^z}=(1+{\rm e}^x)^{-2xy^w} $$
```

- 显示：$$ xz}=(1+{\rm e}{-2xy^w} $$

另外，如果要在左右两边都有上下标，可以用 `\sideset` 命令。

- 例子：

```
$$ \sideset{^1_2}{^3_4}\bigotimes $$
```

- 显示：$$\sideset{3_4}\bigotimes$$

##3．如何输入括号和分隔符

`()`、`[]` 和 `|` 表示符号本身，使用 `\{\}` 来表示 `{}` 。当要显示大号的括号或分隔符时，要用 `\left` 和 `\right` 命令。

一些特殊的括号：

| 输入 | 显示 | 输入 | 显示 |
| :---: | :---: | :---: | :---: |
| \\langle | ![](https://g.yuque.com/gr/latex?%5Clangle#card=math&code=%5Clangle) | \\rangle | ![](https://g.yuque.com/gr/latex?%5Crangle#card=math&code=%5Crangle) |
| \\lceil | ![](https://g.yuque.com/gr/latex?%5Clceil#card=math&code=%5Clceil) | \\rceil | ![](https://g.yuque.com/gr/latex?%5Crceil#card=math&code=%5Crceil) |
| \\lfloor | ![](https://g.yuque.com/gr/latex?%5Clfloor#card=math&code=%5Clfloor) | \\rfloor | ![](https://g.yuque.com/gr/latex?%5Crfloor#card=math&code=%5Crfloor) |
| \\lbrace | ![](https://g.yuque.com/gr/latex?%5Clbrace#card=math&code=%5Clbrace) | \\rbrace | ![](https://g.yuque.com/gr/latex?%5Crbrace#card=math&code=%5Crbrace) |


- 例子：

```
$$ f(x,y,z) = 3y^2z \left( 3+\frac{7x+5}{1+y^2} \right) $$
```

- 显示：$$ f(x,y,z) = 3y^2z \left( 3+\frac{7x+5}{1+y^2} \right) $$

有时候要用 `\left.` 或 `\right.` 进行匹配而不显示本身。

- 例子：

```
$$ \left. \frac{{\rm d}u}{{\rm d}x} \right| _{x=0} $$
```

- 显示：$$ \left. \frac{{\rm d}u}{{\rm d}x} \right| _{x=0} $$

##4．如何输入分数

通常使用 `\frac {分子} {分母}` 命令产生一个分数，分数可嵌套。

便捷情况可直接输入 `\frac ab` 来快速生成一个 ![](https://g.yuque.com/gr/latex?%5Cfrac%20ab#card=math&code=%5Cfrac%20ab) 。

如果分式很复杂，亦可使用 `分子 \over 分母` 命令，此时分数仅有一层。

- 例子：

```
$$\frac{a-1}{b-1} \quad and \quad {a+1\over b+1}$$
```

- 显示：$$\frac{a-1}{b-1} \quad and \quad {a+1\over b+1}$$

##5．如何输入开方

使用 `\sqrt [根指数，省略时为2] {被开方数}` 命令输入开方。

- 例子：

```
$$\sqrt{2} \quad and \quad \sqrt[n]{3}$$
```

- 显示：$$\sqrt{2} \quad and \quad \sqrt[n]{3}$$

##6．如何输入省略号

数学公式中常见的省略号有两种，`\ldots` 表示与文本底线对齐的省略号，`\cdots` 表示与文本中线对齐的省略号。

- 例子：

```
$$f(x_1,x_2,\underbrace{\ldots}_{\rm ldots} ,x_n) = x_1^2 + x_2^2 + \underbrace{\cdots}_{\rm cdots} + x_n^2$$
```

- 显示：$$f(x_1,x_2,\underbrace{\ldots}_{\rm ldots} ,x_n) = x_1^2 + x_2^2 + \underbrace{\cdots}_{\rm cdots} + x_n^2$$

##7．如何输入矢量

使用 `\vec{矢量}` 来自动产生一个矢量。也可以使用 `\overrightarrow` 等命令自定义字母上方的符号。

- 例子：

```
$$\vec{a} \cdot \vec{b}=0$$
```

- 
显示：$$\vec{a} \cdot \vec{b}=0$$

- 
例子：


```
$$\overleftarrow{xy} \quad and \quad \overleftrightarrow{xy} \quad and \quad \overrightarrow{xy}$$
```

- 显示：$$\overleftarrow{xy} \quad and \quad \overleftrightarrow{xy} \quad and \quad \overrightarrow{xy}$$

##8．如何输入积分

使用 `\int_积分下限^积分上限 {被积表达式}` 来输入一个积分。

例子：

```
$$\int_0^1 {x^2} \,{\rm d}x$$
```

显示：$$\int_0^1 {x^2} ,{\rm d}x$$

本例中 `\,` 和 `{\rm d}` 部分可省略，但建议加入，能使式子更美观。

##9．如何输入极限运算

使用 `\lim_{变量 \to 表达式} 表达式` 来输入一个极限。如有需求，可以更改 `\to` 符号至任意符号。

例子：

```
$$ \lim_{n \to +\infty} \frac{1}{n(n+1)} \quad and \quad \lim_{x\leftarrow{示例}} \frac{1}{n(n+1)} $$
```

显示：$$\lim_{n \to +\infty} \frac{1}{n(n+1)} \quad and \quad \lim_{x\leftarrow{示例}} \frac{1}{n(n+1)}$$

##10．如何输入累加、累乘运算

使用 `\sum_{下标表达式}^{上标表达式} {累加表达式}` 来输入一个累加。

与之类似，使用 `\prod` `\bigcup` `\bigcap` 来分别输入累乘、并集和交集。

此类符号在行内显示时上下标表达式将会移至右上角和右下角。

- 例子：

```
$$\sum_{i=1}^n \frac{1}{i^2} \quad and \quad \prod_{i=1}^n \frac{1}{i^2} \quad and \quad \bigcup_{i=1}^{2} R$$
```

- 显示：$$\sum_{i=1}^n \frac{1}{i^2} \quad and \quad \prod_{i=1}^n \frac{1}{i^2} \quad and \quad \bigcup_{i=1}^{2} R$$

##11．如何输入希腊字母

输入 `\小写希腊字母英文全称` 和 `\首字母大写希腊字母英文全称` 来分别输入小写和大写希腊字母。

**对于大写希腊字母与现有字母相同的，直接输入大写字母即可。**

| 输入 | 显示 | 输入 | 显示 | 输入 | 显示 | 输入 | 显示 |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| \\alpha | ![](https://g.yuque.com/gr/latex?%5Calpha#card=math&code=%5Calpha) | A | ![](https://g.yuque.com/gr/latex?A#card=math&code=A) | \\beta | ![](https://g.yuque.com/gr/latex?%5Cbeta#card=math&code=%5Cbeta) | B | ![](https://g.yuque.com/gr/latex?B#card=math&code=B) |
| \\gamma | ![](https://g.yuque.com/gr/latex?%5Cgamma#card=math&code=%5Cgamma) | \\Gamma | ![](https://g.yuque.com/gr/latex?%5CGamma#card=math&code=%5CGamma) | \\delta | ![](https://g.yuque.com/gr/latex?%5Cdelta#card=math&code=%5Cdelta) | \\Delta | ![](https://g.yuque.com/gr/latex?%5CDelta#card=math&code=%5CDelta) |
| \\epsilon | ![](https://g.yuque.com/gr/latex?%5Cepsilon#card=math&code=%5Cepsilon) | E | ![](https://g.yuque.com/gr/latex?E#card=math&code=E) | \\zeta | ![](https://g.yuque.com/gr/latex?%5Czeta#card=math&code=%5Czeta) | Z | ![](https://g.yuque.com/gr/latex?Z#card=math&code=Z) |
| \\eta | ![](https://g.yuque.com/gr/latex?%5Ceta#card=math&code=%5Ceta) | H | ![](https://g.yuque.com/gr/latex?H#card=math&code=H) | \\theta | ![](https://g.yuque.com/gr/latex?%5Ctheta#card=math&code=%5Ctheta) | \\Theta | ![](https://g.yuque.com/gr/latex?%5CTheta#card=math&code=%5CTheta) |
| \\iota | ![](https://g.yuque.com/gr/latex?%5Ciota#card=math&code=%5Ciota) | I | ![](https://g.yuque.com/gr/latex?I#card=math&code=I) | \\kappa | ![](https://g.yuque.com/gr/latex?%5Ckappa#card=math&code=%5Ckappa) | K | ![](https://g.yuque.com/gr/latex?K#card=math&code=K) |
| \\lambda | ![](https://g.yuque.com/gr/latex?%5Clambda#card=math&code=%5Clambda) | \\Lambda | ![](https://g.yuque.com/gr/latex?%5CLambda#card=math&code=%5CLambda) | \\mu | ![](https://g.yuque.com/gr/latex?%5Cmu#card=math&code=%5Cmu) | M | ![](https://g.yuque.com/gr/latex?M#card=math&code=M) |
| \\nu | ![](https://g.yuque.com/gr/latex?%5Cnu#card=math&code=%5Cnu) | N | ![](https://g.yuque.com/gr/latex?N#card=math&code=N) | \\xi | ![](https://g.yuque.com/gr/latex?%5Cxi#card=math&code=%5Cxi) | \\Xi | ![](https://g.yuque.com/gr/latex?%5CXi#card=math&code=%5CXi) |
| o | ![](https://g.yuque.com/gr/latex?o#card=math&code=o) | O | ![](https://g.yuque.com/gr/latex?O#card=math&code=O) | \\pi | ![](https://g.yuque.com/gr/latex?%5Cpi#card=math&code=%5Cpi) | \\Pi | ![](https://g.yuque.com/gr/latex?%5CPi#card=math&code=%5CPi) |
| \\rho | ![](https://g.yuque.com/gr/latex?%5Crho#card=math&code=%5Crho) | P | ![](https://g.yuque.com/gr/latex?P#card=math&code=P) | \\sigma | ![](https://g.yuque.com/gr/latex?%5Csigma#card=math&code=%5Csigma) | \\Sigma | ![](https://g.yuque.com/gr/latex?%5CSigma#card=math&code=%5CSigma) |
| \\tau | ![](https://g.yuque.com/gr/latex?%5Ctau#card=math&code=%5Ctau) | T | ![](https://g.yuque.com/gr/latex?T#card=math&code=T) | \\upsilon | ![](https://g.yuque.com/gr/latex?%5Cupsilon#card=math&code=%5Cupsilon) | \\Upsilon | ![](https://g.yuque.com/gr/latex?%5CUpsilon#card=math&code=%5CUpsilon) |
| \\phi | ![](https://g.yuque.com/gr/latex?%5Cphi#card=math&code=%5Cphi) | \\Phi | ![](https://g.yuque.com/gr/latex?%5CPhi#card=math&code=%5CPhi) | \\chi | ![](https://g.yuque.com/gr/latex?%5Cchi#card=math&code=%5Cchi) | X | ![](https://g.yuque.com/gr/latex?X#card=math&code=X) |
| \\psi | ![](https://g.yuque.com/gr/latex?%5Cpsi#card=math&code=%5Cpsi) | \\Psi | ![](https://g.yuque.com/gr/latex?%5CPsi#card=math&code=%5CPsi) | \\omega | ![](https://g.yuque.com/gr/latex?%5Comega#card=math&code=%5Comega) | \\Omega | ![](https://g.yuque.com/gr/latex?%5COmega#card=math&code=%5COmega) |


**部分字母有变量专用形式，以 `\var-` 开头。**

| 小写形式 | 大写形式 | 变量形式 | 显示 |
| :---: | :---: | :---: | :---: |
| \\epsilon | E | \\varepsilon | ![](https://g.yuque.com/gr/latex?%5Cepsilon%20%5Cmid%20E%20%5Cmid%20%5Cvarepsilon#card=math&code=%5Cepsilon%20%5Cmid%20E%20%5Cmid%20%5Cvarepsilon) |
| \\theta | \\Theta | \\vartheta | ![](https://g.yuque.com/gr/latex?%5Ctheta%20%5Cmid%20%5CTheta%20%5Cmid%20%5Cvartheta#card=math&code=%5Ctheta%20%5Cmid%20%5CTheta%20%5Cmid%20%5Cvartheta) |
| \\rho | P | \\varrho | ![](https://g.yuque.com/gr/latex?%5Crho%20%5Cmid%20P%20%5Cmid%20%5Cvarrho#card=math&code=%5Crho%20%5Cmid%20P%20%5Cmid%20%5Cvarrho) |
| \\sigma | \\Sigma | \\varsigma | ![](https://g.yuque.com/gr/latex?%5Csigma%20%5Cmid%20%5CSigma%20%5Cmid%20%5Cvarsigma#card=math&code=%5Csigma%20%5Cmid%20%5CSigma%20%5Cmid%20%5Cvarsigma) |
| \\phi | \\Phi | \\varphi | ![](https://g.yuque.com/gr/latex?%5Cphi%20%5Cmid%20%5CPhi%20%5Cmid%20%5Cvarphi#card=math&code=%5Cphi%20%5Cmid%20%5CPhi%20%5Cmid%20%5Cvarphi) |


##12．如何输入其它特殊字符

> **若需要显示更大或更小的字符，在符号前插入 `\large` 或 `\small` 命令。**


> 若找不到需要的符号，使用 [![](https://g.yuque.com/gr/latex?%5Crm%7BDetexify%5E2%7D#card=math&code=%5Crm%7BDetexify%5E2%7D)](http://detexify.kirelabs.org/classify.html) 来画出想要的符号。


###(1)．关系运算符

| 输入 | 显示 | 输入 | 显示 | 输入 | 显示 | 输入 | 显示 |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| \\pm | ![](https://g.yuque.com/gr/latex?%5Cpm#card=math&code=%5Cpm) | \\times | ![](https://g.yuque.com/gr/latex?%5Ctimes#card=math&code=%5Ctimes) | \\div | ![](https://g.yuque.com/gr/latex?%5Cdiv#card=math&code=%5Cdiv) | \\mid | ![](https://g.yuque.com/gr/latex?%5Cmid#card=math&code=%5Cmid) |
| \\nmid | ![](https://g.yuque.com/gr/latex?%5Cnmid#card=math&code=%5Cnmid) | \\cdot | ![](https://g.yuque.com/gr/latex?%5Ccdot#card=math&code=%5Ccdot) | \\circ | ![](https://g.yuque.com/gr/latex?%5Ccirc#card=math&code=%5Ccirc) | \\ast | ![](https://g.yuque.com/gr/latex?%5Cast#card=math&code=%5Cast) |
| \\bigodot | ![](https://g.yuque.com/gr/latex?%5Cbigodot#card=math&code=%5Cbigodot) | \\bigotimes | ![](https://g.yuque.com/gr/latex?%5Cbigotimes#card=math&code=%5Cbigotimes) | \\bigoplus | ![](https://g.yuque.com/gr/latex?%5Cbigoplus#card=math&code=%5Cbigoplus) | \\leq | ![](https://g.yuque.com/gr/latex?%5Cleq#card=math&code=%5Cleq) |
| \\geq | ![](https://g.yuque.com/gr/latex?%5Cgeq#card=math&code=%5Cgeq) | \\neq | ![](https://g.yuque.com/gr/latex?%5Cneq#card=math&code=%5Cneq) | \\approx | ![](https://g.yuque.com/gr/latex?%5Capprox#card=math&code=%5Capprox) | \\equiv | ![](https://g.yuque.com/gr/latex?%5Cequiv#card=math&code=%5Cequiv) |
| \\sum | ![](https://g.yuque.com/gr/latex?%5Csum#card=math&code=%5Csum) | \\prod | ![](https://g.yuque.com/gr/latex?%5Cprod#card=math&code=%5Cprod) | \\coprod | ![](https://g.yuque.com/gr/latex?%5Ccoprod#card=math&code=%5Ccoprod) | \\backslash | ![](https://g.yuque.com/gr/latex?%5Cbackslash#card=math&code=%5Cbackslash) |


###(2)．集合运算符

| 输入 | 显示 | 输入 | 显示 | 输入 | 显示 |
| :---: | :---: | :---: | :---: | :---: | :---: |
| \\emptyset | ![](https://g.yuque.com/gr/latex?%5Cemptyset#card=math&code=%5Cemptyset) | \\in | ![](https://g.yuque.com/gr/latex?%5Cin#card=math&code=%5Cin) | \\notin | ![](https://g.yuque.com/gr/latex?%5Cnotin#card=math&code=%5Cnotin) |
| \\subset | ![](https://g.yuque.com/gr/latex?%5Csubset#card=math&code=%5Csubset) | \\supset | ![](https://g.yuque.com/gr/latex?%5Csupset#card=math&code=%5Csupset) | \\subseteq | ![](https://g.yuque.com/gr/latex?%5Csubseteq#card=math&code=%5Csubseteq) |
| \\supseteq | ![](https://g.yuque.com/gr/latex?%5Csupseteq#card=math&code=%5Csupseteq) | \\bigcap | ![](https://g.yuque.com/gr/latex?%5Cbigcap#card=math&code=%5Cbigcap) | \\bigcup | ![](https://g.yuque.com/gr/latex?%5Cbigcup#card=math&code=%5Cbigcup) |
| \\bigvee | ![](https://g.yuque.com/gr/latex?%5Cbigvee#card=math&code=%5Cbigvee) | \\bigwedge | ![](https://g.yuque.com/gr/latex?%5Cbigwedge#card=math&code=%5Cbigwedge) | \\biguplus | ![](https://g.yuque.com/gr/latex?%5Cbiguplus#card=math&code=%5Cbiguplus) |


###(3)．对数运算符

| 输入 | 显示 | 输入 | 显示 | 输入 | 显示 |
| :---: | :---: | :---: | :---: | :---: | :---: |
| \\log | ![](https://g.yuque.com/gr/latex?%5Clog#card=math&code=%5Clog) | \\lg | ![](https://g.yuque.com/gr/latex?%5Clg#card=math&code=%5Clg) | \\ln | ![](https://g.yuque.com/gr/latex?%5Cln#card=math&code=%5Cln) |


###(4)．三角运算符

| 输入 | 显示 | 输入 | 显示 | 输入 | 显示 |
| :---: | :---: | :---: | :---: | :---: | :---: |
| 30^\\circ | ![](https://g.yuque.com/gr/latex?30%5E%5Ccirc#card=math&code=30%5E%5Ccirc) | \\bot | ![](https://g.yuque.com/gr/latex?%5Cbot#card=math&code=%5Cbot) | \\angle A | ![](https://g.yuque.com/gr/latex?%5Cangle%20A#card=math&code=%5Cangle%20A) |
| \\sin | ![](https://g.yuque.com/gr/latex?%5Csin#card=math&code=%5Csin) | \\cos | ![](https://g.yuque.com/gr/latex?%5Ccos#card=math&code=%5Ccos) | \\tan | ![](https://g.yuque.com/gr/latex?%5Ctan#card=math&code=%5Ctan) |
| \\csc | ![](https://g.yuque.com/gr/latex?%5Ccsc#card=math&code=%5Ccsc) | \\sec | ![](https://g.yuque.com/gr/latex?%5Csec#card=math&code=%5Csec) | \\cot | ![](https://g.yuque.com/gr/latex?%5Ccot#card=math&code=%5Ccot) |


###(5)．微积分运算符

| 输入 | 显示 | 输入 | 显示 | 输入 | 显示 |
| :---: | :---: | :---: | :---: | :---: | :---: |
| \\int | ![](https://g.yuque.com/gr/latex?%5Cint#card=math&code=%5Cint) | \\iint | ![](https://g.yuque.com/gr/latex?%5Ciint#card=math&code=%5Ciint) | \\iiint | ![](https://g.yuque.com/gr/latex?%5Ciiint#card=math&code=%5Ciiint) |
| \\iiiint | ![](https://g.yuque.com/gr/latex?%5Ciiiint#card=math&code=%5Ciiiint) | \\oint | ![](https://g.yuque.com/gr/latex?%5Coint#card=math&code=%5Coint) | \\prime | ![](https://g.yuque.com/gr/latex?%5Cprime#card=math&code=%5Cprime) |
| \\lim | ![](https://g.yuque.com/gr/latex?%5Clim#card=math&code=%5Clim) | \\infty | ![](https://g.yuque.com/gr/latex?%5Cinfty#card=math&code=%5Cinfty) | \\nabla | ![](https://g.yuque.com/gr/latex?%5Cnabla#card=math&code=%5Cnabla) |


###(6)．逻辑运算符

| 输入 | 显示 | 输入 | 显示 | 输入 | 显示 |
| :---: | :---: | :---: | :---: | :---: | :---: |
| \\because | ![](https://g.yuque.com/gr/latex?%5Cbecause#card=math&code=%5Cbecause) | \\therefore | ![](https://g.yuque.com/gr/latex?%5Ctherefore#card=math&code=%5Ctherefore) |  |  |
| \\forall | ![](https://g.yuque.com/gr/latex?%5Cforall#card=math&code=%5Cforall) | \\exists | ![](https://g.yuque.com/gr/latex?%5Cexists#card=math&code=%5Cexists) | \\not\\subset | ![](https://g.yuque.com/gr/latex?%5Cnot%5Csubset#card=math&code=%5Cnot%5Csubset) |
| \\not< | ![](https://g.yuque.com/gr/latex?%5Cnot%3C#card=math&code=%5Cnot%3C) | \\not> | ![](https://g.yuque.com/gr/latex?%5Cnot%3E#card=math&code=%5Cnot%3E) | \\not= | ![](https://g.yuque.com/gr/latex?%5Cnot%3D#card=math&code=%5Cnot%3D) |


###(7)．戴帽符号

| 输入 | 显示 | 输入 | 显示 |
| :---: | :---: | :---: | :---: |
| \\hat{xy} | ![](https://g.yuque.com/gr/latex?%5Chat%7Bxy%7D#card=math&code=%5Chat%7Bxy%7D) | \\widehat{xyz} | ![](https://g.yuque.com/gr/latex?%5Cwidehat%7Bxyz%7D#card=math&code=%5Cwidehat%7Bxyz%7D) |
| \\tilde{xy} | ![](https://g.yuque.com/gr/latex?%5Ctilde%7Bxy%7D#card=math&code=%5Ctilde%7Bxy%7D) | \\widetilde{xyz} | ![](https://g.yuque.com/gr/latex?%5Cwidetilde%7Bxyz%7D#card=math&code=%5Cwidetilde%7Bxyz%7D) |
| \\check{x} | ![](https://g.yuque.com/gr/latex?%5Ccheck%7Bx%7D#card=math&code=%5Ccheck%7Bx%7D) | \\breve{y} | ![](https://g.yuque.com/gr/latex?%5Cbreve%7By%7D#card=math&code=%5Cbreve%7By%7D) |
| \\grave{x} | ![](https://g.yuque.com/gr/latex?%5Cgrave%7Bx%7D#card=math&code=%5Cgrave%7Bx%7D) | \\acute{y} | ![](https://g.yuque.com/gr/latex?%5Cacute%7By%7D#card=math&code=%5Cacute%7By%7D) |


###(8)．连线符号

| 输入 | 显示 |
| :---: | :---: |
| \\fbox{a+b+c+d} | ![](https://g.yuque.com/gr/latex?%5Cfbox%7Ba%2Bb%2Bc%2Bd%7D#card=math&code=%5Cfbox%7Ba%2Bb%2Bc%2Bd%7D) |
| \\overleftarrow{a+b+c+d} | ![](https://g.yuque.com/gr/latex?%5Coverleftarrow%7Ba%2Bb%2Bc%2Bd%7D#card=math&code=%5Coverleftarrow%7Ba%2Bb%2Bc%2Bd%7D) |
| \\overrightarrow{a+b+c+d} | ![](https://g.yuque.com/gr/latex?%5Coverrightarrow%7Ba%2Bb%2Bc%2Bd%7D#card=math&code=%5Coverrightarrow%7Ba%2Bb%2Bc%2Bd%7D) |
| \\overleftrightarrow{a+b+c+d} | ![](https://g.yuque.com/gr/latex?%5Coverleftrightarrow%7Ba%2Bb%2Bc%2Bd%7D#card=math&code=%5Coverleftrightarrow%7Ba%2Bb%2Bc%2Bd%7D) |
| \\underleftarrow{a+b+c+d} | ![](https://g.yuque.com/gr/latex?%5Cunderleftarrow%7Ba%2Bb%2Bc%2Bd%7D#card=math&code=%5Cunderleftarrow%7Ba%2Bb%2Bc%2Bd%7D) |
| \\underrightarrow{a+b+c+d} | ![](https://g.yuque.com/gr/latex?%5Cunderrightarrow%7Ba%2Bb%2Bc%2Bd%7D#card=math&code=%5Cunderrightarrow%7Ba%2Bb%2Bc%2Bd%7D) |
| \\underleftrightarrow{a+b+c+d} | ![](https://g.yuque.com/gr/latex?%5Cunderleftrightarrow%7Ba%2Bb%2Bc%2Bd%7D#card=math&code=%5Cunderleftrightarrow%7Ba%2Bb%2Bc%2Bd%7D) |
| \\overline{a+b+c+d} | ![](https://g.yuque.com/gr/latex?%5Coverline%7Ba%2Bb%2Bc%2Bd%7D#card=math&code=%5Coverline%7Ba%2Bb%2Bc%2Bd%7D) |
| \\underline{a+b+c+d} | ![](https://g.yuque.com/gr/latex?%5Cunderline%7Ba%2Bb%2Bc%2Bd%7D#card=math&code=%5Cunderline%7Ba%2Bb%2Bc%2Bd%7D) |
| \\overbrace{a+b+c+d}^{Sample} | ![](https://g.yuque.com/gr/latex?%5Coverbrace%7Ba%2Bb%2Bc%2Bd%7D%5E%7BSample%7D#card=math&code=%5Coverbrace%7Ba%2Bb%2Bc%2Bd%7D%5E%7BSample%7D) |
| \\underbrace{a+b+c+d}_{Sample} | ![](https://g.yuque.com/gr/latex?%5Cunderbrace%7Ba%2Bb%2Bc%2Bd%7D_%7BSample%7D#card=math&code=%5Cunderbrace%7Ba%2Bb%2Bc%2Bd%7D_%7BSample%7D) |
| \\overbrace{a+\\underbrace{b+c}_{1.0}+d}^{2.0} | ![](https://g.yuque.com/gr/latex?%5Coverbrace%7Ba%2B%5Cunderbrace%7Bb%2Bc%7D_%7B1.0%7D%2Bd%7D%5E%7B2.0%7D#card=math&code=%5Coverbrace%7Ba%2B%5Cunderbrace%7Bb%2Bc%7D_%7B1.0%7D%2Bd%7D%5E%7B2.0%7D) |
| \\underbrace{a\\cdot a\\cdots a}_{b\\text{ times}} | ![](https://g.yuque.com/gr/latex?%5Cunderbrace%7Ba%5Ccdot%20a%5Ccdots%20a%7D_%7Bb%5Ctext%7B%20times%7D%7D#card=math&code=%5Cunderbrace%7Ba%5Ccdot%20a%5Ccdots%20a%7D_%7Bb%5Ctext%7B%20times%7D%7D) |
| \\underrightarrow{1℃/min} | ![](https://g.yuque.com/gr/latex?%5Cunderrightarrow%7B1%E2%84%83%2Fmin%7D#card=math&code=%5Cunderrightarrow%7B1%E2%84%83%2Fmin%7D) |


###(9)．箭头符号

- 
推荐使用符号：

|输入|显示|输入|显示|输入|显示|

|:--:|:--:|:--:|:--:|:--:|:--:|

|\to|![](https://g.yuque.com/gr/latex?%5Cto#card=math&code=%5Cto)|\mapsto|![](https://g.yuque.com/gr/latex?%5Cmapsto#card=math&code=%5Cmapsto)|

|\implies|![](https://g.yuque.com/gr/latex?%5Cimplies#card=math&code=%5Cimplies)|\iff|![](https://g.yuque.com/gr/latex?%5Ciff#card=math&code=%5Ciff)|\impliedby|![](https://g.yuque.com/gr/latex?%5Cimpliedby#card=math&code=%5Cimpliedby)|

- 
其它可用符号：

|输入|显示|输入|显示|

|:--:|:--:|:--:|:--:|

|\uparrow|![](https://g.yuque.com/gr/latex?%5Cuparrow#card=math&code=%5Cuparrow)|\Uparrow|![](https://g.yuque.com/gr/latex?%5CUparrow#card=math&code=%5CUparrow)|

|\downarrow|![](https://g.yuque.com/gr/latex?%5Cdownarrow#card=math&code=%5Cdownarrow)|\Downarrow|![](https://g.yuque.com/gr/latex?%5CDownarrow#card=math&code=%5CDownarrow)|

|\leftarrow|![](https://g.yuque.com/gr/latex?%5Cleftarrow#card=math&code=%5Cleftarrow)|\Leftarrow|![](https://g.yuque.com/gr/latex?%5CLeftarrow#card=math&code=%5CLeftarrow)|

|\rightarrow|![](https://g.yuque.com/gr/latex?%5Crightarrow#card=math&code=%5Crightarrow)|\Rightarrow|![](https://g.yuque.com/gr/latex?%5CRightarrow#card=math&code=%5CRightarrow)|

|\leftrightarrow|![](https://g.yuque.com/gr/latex?%5Cleftrightarrow#card=math&code=%5Cleftrightarrow)|\Leftrightarrow|![](https://g.yuque.com/gr/latex?%5CLeftrightarrow#card=math&code=%5CLeftrightarrow)|

|\longleftarrow|![](https://g.yuque.com/gr/latex?%5Clongleftarrow#card=math&code=%5Clongleftarrow)|\Longleftarrow|![](https://g.yuque.com/gr/latex?%5CLongleftarrow#card=math&code=%5CLongleftarrow)|

|\longrightarrow|![](https://g.yuque.com/gr/latex?%5Clongrightarrow#card=math&code=%5Clongrightarrow)|\Longrightarrow|![](https://g.yuque.com/gr/latex?%5CLongrightarrow#card=math&code=%5CLongrightarrow)|

|\longleftrightarrow|![](https://g.yuque.com/gr/latex?%5Clongleftrightarrow#card=math&code=%5Clongleftrightarrow)|\Longleftrightarrow|![](https://g.yuque.com/gr/latex?%5CLongleftrightarrow#card=math&code=%5CLongleftrightarrow)|


##13．如何进行字体转换

若要对公式的某一部分字符进行字体转换，可以用 `{\字体 {需转换的部分字符}}` 命令，其中 `\字体` 部分可以参照下表选择合适的字体。一般情况下，公式默认为意大利体 ![](https://g.yuque.com/gr/latex?italic#card=math&code=italic) 。

示例中 **全部大写** 的字体仅大写可用。

| 输入 | 说明 | 显示 | 输入 | 说明 | 显示 |
| :---: | :---: | :---: | :---: | :---: | :---: |
| \\rm | 罗马体 | ![](https://g.yuque.com/gr/latex?%5Crm%7BSample%7D#card=math&code=%5Crm%7BSample%7D) | \\cal | 花体 | ![](https://g.yuque.com/gr/latex?%5Ccal%7BSAMPLE%7D#card=math&code=%5Ccal%7BSAMPLE%7D) |
| \\it | 意大利体 | ![](https://g.yuque.com/gr/latex?%5Cit%7BSample%7D#card=math&code=%5Cit%7BSample%7D) | \\Bbb | 黑板粗体 | ![](https://g.yuque.com/gr/latex?%5CBbb%7BSAMPLE%7D#card=math&code=%5CBbb%7BSAMPLE%7D) |
| \\bf | 粗体 | ![](https://g.yuque.com/gr/latex?%5Cbf%7BSample%7D#card=math&code=%5Cbf%7BSample%7D) | \\mit | 数学斜体 | ![](https://g.yuque.com/gr/latex?%5Cmit%7BSAMPLE%7D#card=math&code=%5Cmit%7BSAMPLE%7D) |
| \\sf | 等线体 | ![](https://g.yuque.com/gr/latex?%5Csf%7BSample%7D#card=math&code=%5Csf%7BSample%7D) | \\scr | 手写体 | ![](https://g.yuque.com/gr/latex?%5Cscr%7BSAMPLE%7D#card=math&code=%5Cscr%7BSAMPLE%7D) |
| \\tt | 打字机体 | ![](https://g.yuque.com/gr/latex?%5Ctt%7BSample%7D#card=math&code=%5Ctt%7BSample%7D) |  |  |  |
| \\frak | 旧德式字体 | ![](https://g.yuque.com/gr/latex?%5Cfrak%7BSample%7D#card=math&code=%5Cfrak%7BSample%7D) |  |  |  |


转换字体十分常用，例如在积分中：

- 例子：

```
\begin{array}{cc}
\mathrm{Bad} & \mathrm{Better} \\
\hline \\
\int_0^1 x^2 dx & \int_0^1 x^2 \,{\rm d}x
\end{array}
```

- 显示：

\begin{array}{cc}

\mathrm{Bad} & \mathrm{Better} \

\hline \

\int_0^1 x^2 dx & \int_0^1 x^2 ,{\rm d}x

\end{array}

注意比较两个式子间 ![](https://g.yuque.com/gr/latex?dx#card=math&code=dx) 与 ![](https://g.yuque.com/gr/latex?%7B%5Crm%20d%7D%20x#card=math&code=%7B%5Crm%20d%7D%20x) 的不同。

使用 `\operatorname` 命令也可以达到相同的效果，详见 [定义新的符号 \operatorname](#1%E5%AE%9A%E4%B9%89%E6%96%B0%E7%9A%84%E7%AC%A6%E5%8F%B7-operatorname) 。

##14．大括号和行标的使用

使用 `\left` 和 `\right` 来创建自动匹配高度的 (圆括号)，[方括号] 和 {花括号} 。

在每个公式末尾前使用 `\tag{行标}` 来实现行标。

- 例子：

```
$$
f\left(
   \left[ 
     \frac{
       1+\left\{x,y\right\}
     }{
       \left(
          \frac{x}{y}+\frac{y}{x}
       \right)
       \left(u+1\right)
     }+a
   \right]^{3/2}
\right)
\tag{行标}
$$
```

- 显示：

![](https://g.yuque.com/gr/latex?f%5Cleft(%0A%20%20%20%5Cleft%5B%20%0A%20%20%20%20%20%5Cfrac%7B%0A%20%20%20%20%20%20%201%2B%5Cleft%5C%7Bx%2Cy%5Cright%5C%7D%0A%20%20%20%20%20%7D%7B%0A%20%20%20%20%20%20%20%5Cleft(%0A%20%20%20%20%20%20%20%20%20%20%5Cfrac%7Bx%7D%7By%7D%2B%5Cfrac%7By%7D%7Bx%7D%0A%20%20%20%20%20%20%20%5Cright)%0A%20%20%20%20%20%20%20%5Cleft(u%2B1%5Cright)%0A%20%20%20%20%20%7D%2Ba%0A%20%20%20%5Cright%5D%5E%7B3%2F2%7D%0A%5Cright)%0A%5Ctag%7B%E8%A1%8C%E6%A0%87%7D%0A#card=math&code=f%5Cleft%28%0A%20%20%20%5Cleft%5B%20%0A%20%20%20%20%20%5Cfrac%7B%0A%20%20%20%20%20%20%201%2B%5Cleft%5C%7Bx%2Cy%5Cright%5C%7D%0A%20%20%20%20%20%7D%7B%0A%20%20%20%20%20%20%20%5Cleft%28%0A%20%20%20%20%20%20%20%20%20%20%5Cfrac%7Bx%7D%7By%7D%2B%5Cfrac%7By%7D%7Bx%7D%0A%20%20%20%20%20%20%20%5Cright%29%0A%20%20%20%20%20%20%20%5Cleft%28u%2B1%5Cright%29%0A%20%20%20%20%20%7D%2Ba%0A%20%20%20%5Cright%5D%5E%7B3%2F2%7D%0A%5Cright%29%0A%5Ctag%7B%E8%A1%8C%E6%A0%87%7D%0A)

如果你需要在不同的行显示对应括号，可以在每一行对应处使用 `\left.` 或 `\right.` 来放一个"影子"括号：

- 例子：

```
$$
\begin{aligned}
a=&\left(1+2+3+  \cdots \right. \\
& \cdots+ \left. \infty-2+\infty-1+\infty\right)
\end{aligned}
$$
```

- 显示：

![](https://g.yuque.com/gr/latex?%5Cbegin%7Baligned%7D%0Aa%3D%26%5Cleft(1%2B2%2B3%2B%20%20%5Ccdots%20%5Cright.%20%5C%5C%0A%26%20%5Ccdots%2B%20%5Cleft.%20%5Cinfty-2%2B%5Cinfty-1%2B%5Cinfty%5Cright)%0A%5Cend%7Baligned%7D%0A#card=math&code=%5Cbegin%7Baligned%7D%0Aa%3D%26%5Cleft%281%2B2%2B3%2B%20%20%5Ccdots%20%5Cright.%20%5C%5C%0A%26%20%5Ccdots%2B%20%5Cleft.%20%5Cinfty-2%2B%5Cinfty-1%2B%5Cinfty%5Cright%29%0A%5Cend%7Baligned%7D%0A)

如果你需要将行内显示的分隔符也变大，可以使用 `\middle` 命令：

- 例子：

```
$$
\left\langle  
  q
\middle\|
  \frac{\frac{x}{y}}{\frac{u}{v}}
\middle| 
   p 
\right\rangle
$$
```

- 显示：

![](https://g.yuque.com/gr/latex?%5Cleft%5Clangle%20%20%0A%20%20q%0A%5Cmiddle%5C%7C%0A%20%20%5Cfrac%7B%5Cfrac%7Bx%7D%7By%7D%7D%7B%5Cfrac%7Bu%7D%7Bv%7D%7D%0A%5Cmiddle%7C%20%0A%20%20%20p%20%0A%5Cright%5Crangle%0A#card=math&code=%5Cleft%5Clangle%20%20%0A%20%20q%0A%5Cmiddle%5C%7C%0A%20%20%5Cfrac%7B%5Cfrac%7Bx%7D%7By%7D%7D%7B%5Cfrac%7Bu%7D%7Bv%7D%7D%0A%5Cmiddle%7C%20%0A%20%20%20p%20%0A%5Cright%5Crangle%0A)

##15．其它命令

###(1)．定义新的符号 \operatorname

查询 [关于此命令的定义](http://meta.math.stackexchange.com/questions/5020/mathjax-basic-tutorial-and-quick-reference/15077#15077) 和 [关于此命令的讨论](http://meta.math.stackexchange.com/search?q=operatorname) 来进一步了解此命令。

- 例子：

```
$$ \operatorname{Symbol} A $$
```

- 显示： $$\operatorname{Symbol} A$$

###(2)．添加注释文字 \text

在 `\text {文字}` 中仍可以使用 `$公式$` 插入其它公式。

- 例子：

```
$$ f(n)= \begin{cases} n/2, & \text {if $n$ is even} \\ 3n+1, & \text{if $n$ is odd} \end{cases} $$
```

- 显示：

![](https://g.yuque.com/gr/latex?f(n)%3D%20%5Cbegin%7Bcases%7D%20n%2F2%2C%20%26%20%5Ctext%20%7Bif%20%24n%24%20is%20even%7D%20%5C%5C%203n%2B1%2C%20%26%20%5Ctext%7Bif%20%24n%24%20is%20odd%7D%20%5Cend%7Bcases%7D%20%0A#card=math&code=f%28n%29%3D%20%5Cbegin%7Bcases%7D%20n%2F2%2C%20%26%20%5Ctext%20%7Bif%20%24n%24%20is%20even%7D%20%5C%5C%203n%2B1%2C%20%26%20%5Ctext%7Bif%20%24n%24%20is%20odd%7D%20%5Cend%7Bcases%7D%20%0A)

###(3)．在字符间加入空格

有四种宽度的空格可以使用： `\,`、`\;`、`\quad` 和 `\qquad` 。

- 例子：

```
$$ a \, b \mid a \; b \mid a \quad b \mid a \qquad b $$
```

- 显示：$$ a , b \mid a ; b \mid a \quad b \mid a \qquad b $$

当然，使用 `\text {n个空格}` 也可以达到同样效果。

###(4)．更改文字颜色

使用 `\color{颜色}{文字}` 来更改特定的文字颜色。

更改文字颜色 **需要浏览器支持** ，如果浏览器不知道你所需的颜色，那么文字将被渲染为黑色。

对于较旧的浏览器（HTML4与CSS2），以下颜色是被支持的：

| 输入 | 显示 | 输入 | 显示 |
| :---: | :---: | :---: | :---: |
| black | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bblack%7D%7Btext%7D#card=math&code=%5Ccolor%7Bblack%7D%7Btext%7D) | grey | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bgrey%7D%7Btext%7D#card=math&code=%5Ccolor%7Bgrey%7D%7Btext%7D) |
| silver | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bsilver%7D%7Btext%7D#card=math&code=%5Ccolor%7Bsilver%7D%7Btext%7D) | white | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bwhite%7D%7Btext%7D#card=math&code=%5Ccolor%7Bwhite%7D%7Btext%7D) |
| maroon | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bmaroon%7D%7Btext%7D#card=math&code=%5Ccolor%7Bmaroon%7D%7Btext%7D) | red | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bred%7D%7Btext%7D#card=math&code=%5Ccolor%7Bred%7D%7Btext%7D) |
| yellow | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Byellow%7D%7Btext%7D#card=math&code=%5Ccolor%7Byellow%7D%7Btext%7D) | lime | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Blime%7D%7Btext%7D#card=math&code=%5Ccolor%7Blime%7D%7Btext%7D) |
| olive | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bolive%7D%7Btext%7D#card=math&code=%5Ccolor%7Bolive%7D%7Btext%7D) | green | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bgreen%7D%7Btext%7D#card=math&code=%5Ccolor%7Bgreen%7D%7Btext%7D) |
| teal | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bteal%7D%7Btext%7D#card=math&code=%5Ccolor%7Bteal%7D%7Btext%7D) | auqa | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bauqa%7D%7Btext%7D#card=math&code=%5Ccolor%7Bauqa%7D%7Btext%7D) |
| blue | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bblue%7D%7Btext%7D#card=math&code=%5Ccolor%7Bblue%7D%7Btext%7D) | navy | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bnavy%7D%7Btext%7D#card=math&code=%5Ccolor%7Bnavy%7D%7Btext%7D) |
| purple | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bpurple%7D%7Btext%7D#card=math&code=%5Ccolor%7Bpurple%7D%7Btext%7D) | fuchsia | ![](https://g.yuque.com/gr/latex?%5Ccolor%7Bfuchsia%7D%7Btext%7D#card=math&code=%5Ccolor%7Bfuchsia%7D%7Btext%7D) |


对于较新的浏览器（HTML5与CSS3），额外的124种颜色将被支持：

输入 `\color {#rgb} {text}` 来自定义更多的颜色，其中 `#rgb` 的 `r` `g` `b` 可输入 `0-9` 和 `a-f` 来表示红色、绿色和蓝色的纯度（饱和度）。

- 例子：

```
\begin{array}{|rrrrrrrr|}\hline
\verb+#000+ & \color{#000}{text} & & &
\verb+#00F+ & \color{#00F}{text} & & \\
& & \verb+#0F0+ & \color{#0F0}{text} &
& & \verb+#0FF+ & \color{#0FF}{text}\\
\verb+#F00+ & \color{#F00}{text} & & &
\verb+#F0F+ & \color{#F0F}{text} & & \\
& & \verb+#FF0+ & \color{#FF0}{text} &
& & \verb+#FFF+ & \color{#FFF}{text}\\
\hline
\end{array}
```

- 
显示：

\begin{array}{|rrrrrrrr|}\hline

\verb+#000+ & \color{#000}{text} & & &

\verb+#00F+ & \color{#00F}{text} & & \

& & \verb+#0F0+ & \color{#0F0}{text} &

& & \verb+#0FF+ & \color{#0FF}{text}\

\verb+#F00+ & \color{#F00}{text} & & &

\verb+#F0F+ & \color{#F0F}{text} & & \

& & \verb+#FF0+ & \color{#FF0}{text} &

& & \verb+#FFF+ & \color{#FFF}{text}\

\hline

\end{array}

- 
例子：


```
\begin{array}{|rrrrrrrr|}
\hline
\verb+#000+ & \color{#000}{text} & \verb+#005+ & \color{#005}{text} & \verb+#00A+ & \color{#00A}{text} & \verb+#00F+ & \color{#00F}{text}  \\
\verb+#500+ & \color{#500}{text} & \verb+#505+ & \color{#505}{text} & \verb+#50A+ & \color{#50A}{text} & \verb+#50F+ & \color{#50F}{text}  \\
\verb+#A00+ & \color{#A00}{text} & \verb+#A05+ & \color{#A05}{text} & \verb+#A0A+ & \color{#A0A}{text} & \verb+#A0F+ & \color{#A0F}{text}  \\
\verb+#F00+ & \color{#F00}{text} & \verb+#F05+ & \color{#F05}{text} & \verb+#F0A+ & \color{#F0A}{text} & \verb+#F0F+ & \color{#F0F}{text}  \\
\hline
\verb+#080+ & \color{#080}{text} & \verb+#085+ & \color{#085}{text} & \verb+#08A+ & \color{#08A}{text} & \verb+#08F+ & \color{#08F}{text}  \\
\verb+#580+ & \color{#580}{text} & \verb+#585+ & \color{#585}{text} & \verb+#58A+ & \color{#58A}{text} & \verb+#58F+ & \color{#58F}{text}  \\
\verb+#A80+ & \color{#A80}{text} & \verb+#A85+ & \color{#A85}{text} & \verb+#A8A+ & \color{#A8A}{text} & \verb+#A8F+ & \color{#A8F}{text}  \\
\verb+#F80+ & \color{#F80}{text} & \verb+#F85+ & \color{#F85}{text} & \verb+#F8A+ & \color{#F8A}{text} & \verb+#F8F+ & \color{#F8F}{text}  \\
\hline
\verb+#0F0+ & \color{#0F0}{text} & \verb+#0F5+ & \color{#0F5}{text} & \verb+#0FA+ & \color{#0FA}{text} & \verb+#0FF+ & \color{#0FF}{text}  \\
\verb+#5F0+ & \color{#5F0}{text} & \verb+#5F5+ & \color{#5F5}{text} & \verb+#5FA+ & \color{#5FA}{text} & \verb+#5FF+ & \color{#5FF}{text}  \\
\verb+#AF0+ & \color{#AF0}{text} & \verb+#AF5+ & \color{#AF5}{text} & \verb+#AFA+ & \color{#AFA}{text} & \verb+#AFF+ & \color{#AFF}{text}  \\
\verb+#FF0+ & \color{#FF0}{text} & \verb+#FF5+ & \color{#FF5}{text} & \verb+#FFA+ & \color{#FFA}{text} & \verb+#FFF+ & \color{#FFF}{text}  \\
\hline
\end{array}
```

- 显示：

\begin{array}{|rrrrrrrr|}

\hline

\verb+#000+ & \color{#000}{text} & \verb+#005+ & \color{#005}{text} & \verb+#00A+ & \color{#00A}{text} & \verb+#00F+ & \color{#00F}{text}  \

\verb+#500+ & \color{#500}{text} & \verb+#505+ & \color{#505}{text} & \verb+#50A+ & \color{#50A}{text} & \verb+#50F+ & \color{#50F}{text}  \

\verb+#A00+ & \color{#A00}{text} & \verb+#A05+ & \color{#A05}{text} & \verb+#A0A+ & \color{#A0A}{text} & \verb+#A0F+ & \color{#A0F}{text}  \

\verb+#F00+ & \color{#F00}{text} & \verb+#F05+ & \color{#F05}{text} & \verb+#F0A+ & \color{#F0A}{text} & \verb+#F0F+ & \color{#F0F}{text}  \

\hline

\verb+#080+ & \color{#080}{text} & \verb+#085+ & \color{#085}{text} & \verb+#08A+ & \color{#08A}{text} & \verb+#08F+ & \color{#08F}{text}  \

\verb+#580+ & \color{#580}{text} & \verb+#585+ & \color{#585}{text} & \verb+#58A+ & \color{#58A}{text} & \verb+#58F+ & \color{#58F}{text}  \

\verb+#A80+ & \color{#A80}{text} & \verb+#A85+ & \color{#A85}{text} & \verb+#A8A+ & \color{#A8A}{text} & \verb+#A8F+ & \color{#A8F}{text}  \

\verb+#F80+ & \color{#F80}{text} & \verb+#F85+ & \color{#F85}{text} & \verb+#F8A+ & \color{#F8A}{text} & \verb+#F8F+ & \color{#F8F}{text}  \

\hline

\verb+#0F0+ & \color{#0F0}{text} & \verb+#0F5+ & \color{#0F5}{text} & \verb+#0FA+ & \color{#0FA}{text} & \verb+#0FF+ & \color{#0FF}{text}  \

\verb+#5F0+ & \color{#5F0}{text} & \verb+#5F5+ & \color{#5F5}{text} & \verb+#5FA+ & \color{#5FA}{text} & \verb+#5FF+ & \color{#5FF}{text}  \

\verb+#AF0+ & \color{#AF0}{text} & \verb+#AF5+ & \color{#AF5}{text} & \verb+#AFA+ & \color{#AFA}{text} & \verb+#AFF+ & \color{#AFF}{text}  \

\verb+#FF0+ & \color{#FF0}{text} & \verb+#FF5+ & \color{#FF5}{text} & \verb+#FFA+ & \color{#FFA}{text} & \verb+#FFF+ & \color{#FFF}{text}  \

\hline

\end{array}

###(5)．添加删除线

使用删除线功能必须声明 `$$` 符号。

在公式内使用 `\require{cancel}` 来允许 **片段删除线** 的显示。

声明片段删除线后，使用 `\cancel{字符}`、`\bcancel{字符}`、`\xcancel{字符}` 和 `\cancelto{字符}` 来实现各种片段删除线效果。

- 例子：

```
$$
\require{cancel}\begin{array}{rl}
\verb|y+\cancel{x}| & y+\cancel{x}\\
\verb|\cancel{y+x}| & \cancel{y+x}\\
\verb|y+\bcancel{x}| & y+\bcancel{x}\\
\verb|y+\xcancel{x}| & y+\xcancel{x}\\
\verb|y+\cancelto{0}{x}| & y+\cancelto{0}{x}\\
\verb+\frac{1\cancel9}{\cancel95} = \frac15+& \frac{1\cancel9}{\cancel95} = \frac15 \\
\end{array}
$$
```

- 显示：

![](https://g.yuque.com/gr/latex?%5Crequire%7Bcancel%7D%0A%5Cbegin%7Barray%7D%7Brl%7D%0A%5Cverb%7Cy%2B%5Ccancel%7Bx%7D%7C%20%26%20y%2B%5Ccancel%7Bx%7D%5C%5C%0A%5Cverb%7C%5Ccancel%7By%2Bx%7D%7C%20%26%20%5Ccancel%7By%2Bx%7D%5C%5C%0A%5Cverb%7Cy%2B%5Cbcancel%7Bx%7D%7C%20%26%20y%2B%5Cbcancel%7Bx%7D%5C%5C%0A%5Cverb%7Cy%2B%5Cxcancel%7Bx%7D%7C%20%26%20y%2B%5Cxcancel%7Bx%7D%5C%5C%0A%5Cverb%7Cy%2B%5Ccancelto%7B0%7D%7Bx%7D%7C%20%26%20y%2B%5Ccancelto%7B0%7D%7Bx%7D%5C%5C%0A%5Cverb%2B%5Cfrac%7B1%5Ccancel9%7D%7B%5Ccancel95%7D%20%3D%20%5Cfrac15%2B%26%20%5Cfrac%7B1%5Ccancel9%7D%7B%5Ccancel95%7D%20%3D%20%5Cfrac15%20%5C%5C%0A%5Cend%7Barray%7D%0A#card=math&code=%5Crequire%7Bcancel%7D%0A%5Cbegin%7Barray%7D%7Brl%7D%0A%5Cverb%7Cy%2B%5Ccancel%7Bx%7D%7C%20%26%20y%2B%5Ccancel%7Bx%7D%5C%5C%0A%5Cverb%7C%5Ccancel%7By%2Bx%7D%7C%20%26%20%5Ccancel%7By%2Bx%7D%5C%5C%0A%5Cverb%7Cy%2B%5Cbcancel%7Bx%7D%7C%20%26%20y%2B%5Cbcancel%7Bx%7D%5C%5C%0A%5Cverb%7Cy%2B%5Cxcancel%7Bx%7D%7C%20%26%20y%2B%5Cxcancel%7Bx%7D%5C%5C%0A%5Cverb%7Cy%2B%5Ccancelto%7B0%7D%7Bx%7D%7C%20%26%20y%2B%5Ccancelto%7B0%7D%7Bx%7D%5C%5C%0A%5Cverb%2B%5Cfrac%7B1%5Ccancel9%7D%7B%5Ccancel95%7D%20%3D%20%5Cfrac15%2B%26%20%5Cfrac%7B1%5Ccancel9%7D%7B%5Ccancel95%7D%20%3D%20%5Cfrac15%20%5C%5C%0A%5Cend%7Barray%7D%0A)

使用 `\require{enclose}` 来允许 **整段删除线** 的显示。

声明整段删除线后，使用 `\enclose{删除线效果}{字符}` 来实现各种整段删除线效果。

其中，删除线效果有 `horizontalstrike`、`verticalstrike`、`updiagonalstrike` 和 `downdiagonalstrike`，可叠加使用。

- 例子：

```
$$
\require{enclose}\begin{array}{rl}
\verb|\enclose{horizontalstrike}{x+y}| & \enclose{horizontalstrike}{x+y}\\
\verb|\enclose{verticalstrike}{\frac xy}| & \enclose{verticalstrike}{\frac xy}\\
\verb|\enclose{updiagonalstrike}{x+y}| & \enclose{updiagonalstrike}{x+y}\\
\verb|\enclose{downdiagonalstrike}{x+y}| & \enclose{downdiagonalstrike}{x+y}\\
\verb|\enclose{horizontalstrike,updiagonalstrike}{x+y}| & \enclose{horizontalstrike,updiagonalstrike}{x+y}\\
\end{array}
$$
```

- 显示：

![](https://g.yuque.com/gr/latex?%5Crequire%7Benclose%7D%5Cbegin%7Barray%7D%7Brl%7D%0A%5Cverb%7C%5Cenclose%7Bhorizontalstrike%7D%7Bx%2By%7D%7C%20%26%20%5Cenclose%7Bhorizontalstrike%7D%7Bx%2By%7D%5C%5C%0A%5Cverb%7C%5Cenclose%7Bverticalstrike%7D%7B%5Cfrac%20xy%7D%7C%20%26%20%5Cenclose%7Bverticalstrike%7D%7B%5Cfrac%20xy%7D%5C%5C%0A%5Cverb%7C%5Cenclose%7Bupdiagonalstrike%7D%7Bx%2By%7D%7C%20%26%20%5Cenclose%7Bupdiagonalstrike%7D%7Bx%2By%7D%5C%5C%0A%5Cverb%7C%5Cenclose%7Bdowndiagonalstrike%7D%7Bx%2By%7D%7C%20%26%20%5Cenclose%7Bdowndiagonalstrike%7D%7Bx%2By%7D%5C%5C%0A%5Cverb%7C%5Cenclose%7Bhorizontalstrike%2Cupdiagonalstrike%7D%7Bx%2By%7D%7C%20%26%20%5Cenclose%7Bhorizontalstrike%2Cupdiagonalstrike%7D%7Bx%2By%7D%5C%5C%0A%5Cend%7Barray%7D%0A#card=math&code=%5Crequire%7Benclose%7D%5Cbegin%7Barray%7D%7Brl%7D%0A%5Cverb%7C%5Cenclose%7Bhorizontalstrike%7D%7Bx%2By%7D%7C%20%26%20%5Cenclose%7Bhorizontalstrike%7D%7Bx%2By%7D%5C%5C%0A%5Cverb%7C%5Cenclose%7Bverticalstrike%7D%7B%5Cfrac%20xy%7D%7C%20%26%20%5Cenclose%7Bverticalstrike%7D%7B%5Cfrac%20xy%7D%5C%5C%0A%5Cverb%7C%5Cenclose%7Bupdiagonalstrike%7D%7Bx%2By%7D%7C%20%26%20%5Cenclose%7Bupdiagonalstrike%7D%7Bx%2By%7D%5C%5C%0A%5Cverb%7C%5Cenclose%7Bdowndiagonalstrike%7D%7Bx%2By%7D%7C%20%26%20%5Cenclose%7Bdowndiagonalstrike%7D%7Bx%2By%7D%5C%5C%0A%5Cverb%7C%5Cenclose%7Bhorizontalstrike%2Cupdiagonalstrike%7D%7Bx%2By%7D%7C%20%26%20%5Cenclose%7Bhorizontalstrike%2Cupdiagonalstrike%7D%7Bx%2By%7D%5C%5C%0A%5Cend%7Barray%7D%0A)

此外， `\enclose` 命令还可以产生包围的边框和圆等，参见 [MathML Menclose Documentation](https://developer.mozilla.org/en-US/docs/Web/MathML/Element/menclose) 以查看更多效果。

#二、矩阵使用参考

##1．如何输入无框矩阵

在开头使用 `begin{matrix}`，在结尾使用 `end{matrix}`，在中间插入矩阵元素，每个元素之间插入 `&` ，并在每行结尾处使用 `\\` 。

使用矩阵时必须声明 `$` 或 `$$` 符号。

- 例子：

```
$$
        \begin{matrix}
        1 & x & x^2 \\
        1 & y & y^2 \\
        1 & z & z^2 \\
        \end{matrix}
$$
```

- 显示：

![](https://g.yuque.com/gr/latex?%20%20%20%20%20%20%20%20%5Cbegin%7Bmatrix%7D%0A%20%20%20%20%20%20%20%201%20%26%20x%20%26%20x%5E2%20%5C%5C%0A%20%20%20%20%20%20%20%201%20%26%20y%20%26%20y%5E2%20%5C%5C%0A%20%20%20%20%20%20%20%201%20%26%20z%20%26%20z%5E2%20%5C%5C%0A%20%20%20%20%20%20%20%20%5Cend%7Bmatrix%7D%0A#card=math&code=%20%20%20%20%20%20%20%20%5Cbegin%7Bmatrix%7D%0A%20%20%20%20%20%20%20%201%20%26%20x%20%26%20x%5E2%20%5C%5C%0A%20%20%20%20%20%20%20%201%20%26%20y%20%26%20y%5E2%20%5C%5C%0A%20%20%20%20%20%20%20%201%20%26%20z%20%26%20z%5E2%20%5C%5C%0A%20%20%20%20%20%20%20%20%5Cend%7Bmatrix%7D%0A)

##2．如何输入边框矩阵

在开头将 `matrix` 替换为 `pmatrix` `bmatrix` `Bmatrix` `vmatrix` `Vmatrix` 。

- 例子：

```
$ \begin{matrix} 1 & 2 \\ 3 & 4 \\ \end{matrix} $
$ \begin{pmatrix} 1 & 2 \\ 3 & 4 \\ \end{pmatrix} $
$ \begin{bmatrix} 1 & 2 \\ 3 & 4 \\ \end{bmatrix} $
$ \begin{Bmatrix} 1 & 2 \\ 3 & 4 \\ \end{Bmatrix} $
$ \begin{vmatrix} 1 & 2 \\ 3 & 4 \\ \end{vmatrix} $
$ \begin{Vmatrix} 1 & 2 \\ 3 & 4 \\ \end{Vmatrix} $
```

- 显示：

|matrix|pmatrix|bmatrix|Bmatrix|vmatrix|Vmatrix|

|:--:|:--:|:--:|:--:|:--:|:--:|

|$ \begin{matrix} 1 & 2 \ 3 & 4 \ \end{matrix} ![](https://g.yuque.com/gr/latex?%7C#card=math&code=%7C) \begin{pmatrix} 1 & 2 \ 3 & 4 \ \end{pmatrix} ![](https://g.yuque.com/gr/latex?%7C#card=math&code=%7C) \begin{bmatrix} 1 & 2 \ 3 & 4 \ \end{bmatrix} ![](https://g.yuque.com/gr/latex?%7C#card=math&code=%7C) \begin{Bmatrix} 1 & 2 \ 3 & 4 \ \end{Bmatrix} ![](https://g.yuque.com/gr/latex?%7C#card=math&code=%7C) \begin{vmatrix} 1 & 2 \ 3 & 4 \ \end{vmatrix} ![](https://g.yuque.com/gr/latex?%7C#card=math&code=%7C) \begin{Vmatrix} 1 & 2 \ 3 & 4 \ \end{Vmatrix} $|

　　　　

##3．如何输入带省略符号的矩阵

使用 `\cdots` ![](https://g.yuque.com/gr/latex?%5Ccdots#card=math&code=%5Ccdots) , `\ddots` ![](https://g.yuque.com/gr/latex?%5Cddots#card=math&code=%5Cddots) , `\vdots` ![](https://g.yuque.com/gr/latex?%5Cvdots#card=math&code=%5Cvdots) 来输入省略符号。

- 例子：

```
$$
        \begin{pmatrix}
        1 & a_1 & a_1^2 & \cdots & a_1^n \\
        1 & a_2 & a_2^2 & \cdots & a_2^n \\
        \vdots & \vdots & \vdots & \ddots & \vdots \\
        1 & a_m & a_m^2 & \cdots & a_m^n \\
        \end{pmatrix}
$$
```

- 显示：

![](https://g.yuque.com/gr/latex?%20%20%20%20%20%20%20%20%5Cbegin%7Bpmatrix%7D%0A%20%20%20%20%20%20%20%201%20%26%20a_1%20%26%20a_1%5E2%20%26%20%5Ccdots%20%26%20a_1%5En%20%5C%5C%0A%20%20%20%20%20%20%20%201%20%26%20a_2%20%26%20a_2%5E2%20%26%20%5Ccdots%20%26%20a_2%5En%20%5C%5C%0A%20%20%20%20%20%20%20%20%5Cvdots%20%26%20%5Cvdots%20%26%20%5Cvdots%20%26%20%5Cddots%20%26%20%5Cvdots%20%5C%5C%0A%20%20%20%20%20%20%20%201%20%26%20a_m%20%26%20a_m%5E2%20%26%20%5Ccdots%20%26%20a_m%5En%20%5C%5C%0A%20%20%20%20%20%20%20%20%5Cend%7Bpmatrix%7D%0A#card=math&code=%20%20%20%20%20%20%20%20%5Cbegin%7Bpmatrix%7D%0A%20%20%20%20%20%20%20%201%20%26%20a_1%20%26%20a_1%5E2%20%26%20%5Ccdots%20%26%20a_1%5En%20%5C%5C%0A%20%20%20%20%20%20%20%201%20%26%20a_2%20%26%20a_2%5E2%20%26%20%5Ccdots%20%26%20a_2%5En%20%5C%5C%0A%20%20%20%20%20%20%20%20%5Cvdots%20%26%20%5Cvdots%20%26%20%5Cvdots%20%26%20%5Cddots%20%26%20%5Cvdots%20%5C%5C%0A%20%20%20%20%20%20%20%201%20%26%20a_m%20%26%20a_m%5E2%20%26%20%5Ccdots%20%26%20a_m%5En%20%5C%5C%0A%20%20%20%20%20%20%20%20%5Cend%7Bpmatrix%7D%0A)

##4．如何输入带分割符号的矩阵

详见"[数组使用参考](#%E4%BA%94%E6%95%B0%E7%BB%84%E4%B8%8E%E8%A1%A8%E6%A0%BC%E4%BD%BF%E7%94%A8%E5%8F%82%E8%80%83)"。

- 例子：

```
$$
\left[
    \begin{array}{cc|c}
      1&2&3\\
      4&5&6
    \end{array}
\right]
$$
```

- 显示：

![](https://g.yuque.com/gr/latex?%5Cleft%5B%0A%20%20%20%20%5Cbegin%7Barray%7D%7Bcc%7Cc%7D%0A%20%20%20%20%20%201%262%263%5C%5C%0A%20%20%20%20%20%204%265%266%0A%20%20%20%20%5Cend%7Barray%7D%0A%5Cright%5D%0A#card=math&code=%5Cleft%5B%0A%20%20%20%20%5Cbegin%7Barray%7D%7Bcc%7Cc%7D%0A%20%20%20%20%20%201%262%263%5C%5C%0A%20%20%20%20%20%204%265%266%0A%20%20%20%20%5Cend%7Barray%7D%0A%5Cright%5D%0A)

其中 `cc|c` 代表在一个三列矩阵中的第二和第三列之间插入分割线。

##5．如何输入行中矩阵

若想在一行内显示矩阵，

使用`\bigl(\begin{smallmatrix} ... \end{smallmatrix}\bigr)`。

- 例子：

```
这是一个行中矩阵的示例 $\bigl( \begin{smallmatrix} a & b \\ c & d \end{smallmatrix} \bigr)$ 。
```

- 显示：这是一个行中矩阵的示例 ![](https://g.yuque.com/gr/latex?%5Cbigl(%20%5Cbegin%7Bsmallmatrix%7D%20a%20%26%20b%20%5C%5C%20c%20%26%20d%20%5Cend%7Bsmallmatrix%7D%20%5Cbigr)#card=math&code=%5Cbigl%28%20%5Cbegin%7Bsmallmatrix%7D%20a%20%26%20b%20%5C%5C%20c%20%26%20d%20%5Cend%7Bsmallmatrix%7D%20%5Cbigr%29) 。

#三、方程式序列使用参考

##1．如何输入一个方程式序列

人们经常想要一列整齐且居中的方程式序列。使用 `\begin{align}…\end{align}` 来创造一列方程式，其中在每行结尾处使用 `\\` 。

使用方程式序列无需声明公式符号 `$` 或 `$$` 。

请注意 `{align}` 语句是 **自动编号** 的，使用 `{align*}` 声明停止自动编号(^wuyufei批注)。

- 例子：

```
\begin{align}
\sqrt{37} & = \sqrt{\frac{73^2-1}{12^2}} \\
 & = \sqrt{\frac{73^2}{12^2}\cdot\frac{73^2-1}{73^2}} \\ 
 & = \sqrt{\frac{73^2}{12^2}}\sqrt{\frac{73^2-1}{73^2}} \\
 & = \frac{73}{12}\sqrt{1 - \frac{1}{73^2}} \\ 
 & \approx \frac{73}{12}\left(1 - \frac{1}{2\cdot73^2}\right)
\end{align}
```

- 显示：

\begin{align}

\sqrt{37} & = \sqrt{\frac{732}} \

& = \sqrt{\frac{732}\cdot\frac{732}} \

& = \sqrt{\frac{732}}\sqrt{\frac{732}} \

& = \frac{73}{12}\sqrt{1 - \frac{1}{73^2}} \

& \approx \frac{73}{12}\left(1 - \frac{1}{2\cdot73^2}\right)

\end{align}

本例中每行公式的编号续自 [如何插入公式](#1%E5%A6%82%E4%BD%95%E6%8F%92%E5%85%A5%E5%85%AC%E5%BC%8F) 中的自动编号公式 \eqref{eq:sample} 。

##2．在一个方程式序列的每一行中注明原因

在 `{align}` 中灵活组合 `\text` 和 `\tag` 语句。`\tag` 语句编号优先级高于自动编号。

- 例子：

```
\begin{align}
   v + w & = 0  &\text{Given} \tag 1\\
   -w & = -w + 0 & \text{additive identity} \tag 2\\
   -w + 0 & = -w + (v + w) & \text{equations $(1)$ and $(2)$}
\end{align}
```

- 显示：

\begin{align}

v + w & = 0  &\text{Given} \tag 1\

-w & = -w + 0 & \text{additive identity} \tag 2\

-w + 0 & = -w + (v + w) & \text{equations ![](https://g.yuque.com/gr/latex?(1)#card=math&code=%281%29) and ![](https://g.yuque.com/gr/latex?(2)#card=math&code=%282%29)}

\end{align}

本例中第一、第二行的自动编号被 `\tag` 语句覆盖，第三行的编号为自动编号。

#四、条件表达式使用参考

##1．如何输入一个条件表达式

使用 `begin{cases}` 来创造一组条件表达式，在每一行条件中插入 `&` 来指定需要对齐的内容，并在每一行结尾处使用 `\\`，以 `end{cases}` 结束。

条件表达式无需声明 `$` 或 `$$` 符号。

- 例子：

```
$$
        f(n) =
        \begin{cases}
        n/2,  & \text{if $n$ is even} \\
        3n+1, & \text{if $n$ is odd}
        \end{cases}
$$
```

- 显示：

![](https://g.yuque.com/gr/latex?%20%20%20%20%20%20%20%20f(n)%20%3D%0A%20%20%20%20%20%20%20%20%5Cbegin%7Bcases%7D%0A%20%20%20%20%20%20%20%20n%2F2%2C%20%20%26%20%5Ctext%7Bif%20%24n%24%20is%20even%7D%20%5C%5C%0A%20%20%20%20%20%20%20%203n%2B1%2C%20%26%20%5Ctext%7Bif%20%24n%24%20is%20odd%7D%0A%20%20%20%20%20%20%20%20%5Cend%7Bcases%7D%0A#card=math&code=%20%20%20%20%20%20%20%20f%28n%29%20%3D%0A%20%20%20%20%20%20%20%20%5Cbegin%7Bcases%7D%0A%20%20%20%20%20%20%20%20n%2F2%2C%20%20%26%20%5Ctext%7Bif%20%24n%24%20is%20even%7D%20%5C%5C%0A%20%20%20%20%20%20%20%203n%2B1%2C%20%26%20%5Ctext%7Bif%20%24n%24%20is%20odd%7D%0A%20%20%20%20%20%20%20%20%5Cend%7Bcases%7D%0A)

##2．如何输入一个左侧对齐的条件表达式

若想让文字在 **左侧对齐显示** ，则有如下方式：

- 例子：

```
$$
        \left.
        \begin{array}{l}
        \text{if $n$ is even:}&n/2\\
        \text{if $n$ is odd:}&3n+1
        \end{array}
        \right\}
        =f(n)
$$
```

- 显示：

![](https://g.yuque.com/gr/latex?%20%20%20%20%20%20%20%20%5Cleft.%0A%20%20%20%20%20%20%20%20%5Cbegin%7Barray%7D%7Bl%7D%0A%20%20%20%20%20%20%20%20%5Ctext%7Bif%20%24n%24%20is%20even%3A%7D%26n%2F2%5C%5C%0A%20%20%20%20%20%20%20%20%5Ctext%7Bif%20%24n%24%20is%20odd%3A%7D%263n%2B1%0A%20%20%20%20%20%20%20%20%5Cend%7Barray%7D%0A%20%20%20%20%20%20%20%20%5Cright%5C%7D%0A%20%20%20%20%20%20%20%20%3Df(n)%0A#card=math&code=%20%20%20%20%20%20%20%20%5Cleft.%0A%20%20%20%20%20%20%20%20%5Cbegin%7Barray%7D%7Bl%7D%0A%20%20%20%20%20%20%20%20%5Ctext%7Bif%20%24n%24%20is%20even%3A%7D%26n%2F2%5C%5C%0A%20%20%20%20%20%20%20%20%5Ctext%7Bif%20%24n%24%20is%20odd%3A%7D%263n%2B1%0A%20%20%20%20%20%20%20%20%5Cend%7Barray%7D%0A%20%20%20%20%20%20%20%20%5Cright%5C%7D%0A%20%20%20%20%20%20%20%20%3Df%28n%29%0A)

##3．如何使条件表达式适配行高

在一些情况下，条件表达式中某些行的行高为非标准高度，此时使用 `\\[2ex]` 语句代替该行末尾的 `\\` 来让编辑器适配。

- 例子：

|不适配[2ex]|

|:--:|

|

```
$$
f(n) = 
\begin{cases}
\frac{n}{2},  & \text{if $n$ is even} \\
3n+1, & \text{if $n$ is odd}
\end{cases}
$$
```
| 适配[2ex] |
| :---: |
|  |


```
$$
f(n) = 
\begin{cases}
\frac{n}{2},  & \text{if $n$ is even} \\[2ex]
3n+1, & \text{if $n$ is odd}
\end{cases}
$$
```

- 显示：

|不适配[2ex]|

|:--:|

|$$

f(n) =

\begin{cases}

\frac{n}{2},  & \text{if ![](https://g.yuque.com/gr/latex?n#card=math&code=n) is even} \

3n+1, & \text{if ![](https://g.yuque.com/gr/latex?n#card=math&code=n) is odd}

\end{cases}

![](https://g.yuque.com/gr/latex?%7C%0A%0A%7C%E9%80%82%E9%85%8D%5B2ex%5D%7C%0A%7C%3A--%3A%7C%0A%7C#card=math&code=%7C%0A%0A%7C%E9%80%82%E9%85%8D%5B2ex%5D%7C%0A%7C%3A--%3A%7C%0A%7C)

f(n) =

\begin{cases}

\frac{n}{2},  & \text{if ![](https://g.yuque.com/gr/latex?n#card=math&code=n) is even} \[2ex]

3n+1, & \text{if ![](https://g.yuque.com/gr/latex?n#card=math&code=n) is odd}

\end{cases}

![](https://g.yuque.com/gr/latex?%7C%0A%0A**%E4%B8%80%E4%B8%AA%20%60%5Bex%5D%60%20%E6%8C%87%E4%B8%80%E4%B8%AA%20%22X-Height%22%EF%BC%8C%E5%8D%B3x%E5%AD%97%E6%AF%8D%E9%AB%98%E5%BA%A6%E3%80%82%E5%8F%AF%E4%BB%A5%E6%A0%B9%E6%8D%AE%E6%83%85%E5%86%B5%E6%8C%87%E5%AE%9A%E5%A4%9A%E4%B8%AA%20%60%5Bex%5D%60%EF%BC%8C%E5%A6%82%20%60%5B3ex%5D%60%E3%80%81%60%5B4ex%5D%60%20%E7%AD%89%E3%80%82**%0A%E5%85%B6%E5%AE%9E%E5%8F%AF%E4%BB%A5%E5%9C%A8%E4%BB%BB%E4%BD%95%E5%9C%B0%E6%96%B9%E4%BD%BF%E7%94%A8%20%60%5C%5C%5B2ex%5D%60%20%E8%AF%AD%E5%8F%A5%EF%BC%8C%E5%8F%AA%E8%A6%81%E4%BD%A0%E8%A7%89%E5%BE%97%E5%90%88%E9%80%82%E3%80%82%0A%0A%23%E4%BA%94%E3%80%81%E6%95%B0%E7%BB%84%E4%B8%8E%E8%A1%A8%E6%A0%BC%E4%BD%BF%E7%94%A8%E5%8F%82%E8%80%83%0A%0A%23%231%EF%BC%8E%E5%A6%82%E4%BD%95%E8%BE%93%E5%85%A5%E4%B8%80%E4%B8%AA%E6%95%B0%E7%BB%84%E6%88%96%E8%A1%A8%E6%A0%BC%0A%0A%E9%80%9A%E5%B8%B8%EF%BC%8C%E4%B8%80%E4%B8%AA%E6%A0%BC%E5%BC%8F%E5%8C%96%E5%90%8E%E7%9A%84%E8%A1%A8%E6%A0%BC%E6%AF%94%E5%8D%95%E7%BA%AF%E7%9A%84%E6%96%87%E5%AD%97%E6%88%96%E6%8E%92%E7%89%88%E5%90%8E%E7%9A%84%E6%96%87%E5%AD%97%E6%9B%B4%E5%85%B7%E6%9C%89%E5%8F%AF%E8%AF%BB%E6%80%A7%E3%80%82%E6%95%B0%E7%BB%84%E5%92%8C%E8%A1%A8%E6%A0%BC%E5%9D%87%E4%BB%A5%20%60begin%7Barray%7D%60%20%E5%BC%80%E5%A4%B4%EF%BC%8C%E5%B9%B6%E5%9C%A8%E5%85%B6%E5%90%8E%E5%AE%9A%E4%B9%89%E5%88%97%E6%95%B0%E5%8F%8A%E6%AF%8F%E4%B8%80%E5%88%97%E7%9A%84%E6%96%87%E6%9C%AC%E5%AF%B9%E9%BD%90%E5%B1%9E%E6%80%A7%EF%BC%8C%60c%60%20%60l%60%20%60r%60%20%E5%88%86%E5%88%AB%E4%BB%A3%E8%A1%A8%E5%B1%85%E4%B8%AD%E3%80%81%E5%B7%A6%E5%AF%B9%E9%BD%90%E5%8F%8A%E5%8F%B3%E5%AF%B9%E9%BD%90%E3%80%82%E8%8B%A5%E9%9C%80%E8%A6%81%E6%8F%92%E5%85%A5%E5%9E%82%E7%9B%B4%E5%88%86%E5%89%B2%E7%BA%BF%EF%BC%8C%E5%9C%A8%E5%AE%9A%E4%B9%89%E5%BC%8F%E4%B8%AD%E6%8F%92%E5%85%A5%20%60%7C%60%20%EF%BC%8C%E8%8B%A5%E8%A6%81%E6%8F%92%E5%85%A5%E6%B0%B4%E5%B9%B3%E5%88%86%E5%89%B2%E7%BA%BF%EF%BC%8C%E5%9C%A8%E4%B8%8B%E4%B8%80%E8%A1%8C%E8%BE%93%E5%85%A5%E5%89%8D%E6%8F%92%E5%85%A5%20%60%5Chline%60%20%E3%80%82%E4%B8%8E%E7%9F%A9%E9%98%B5%E7%9B%B8%E4%BC%BC%EF%BC%8C%E6%AF%8F%E8%A1%8C%E5%85%83%E7%B4%A0%E9%97%B4%E5%9D%87%E9%A1%BB%E8%A6%81%E6%8F%92%E5%85%A5%20%60%26%60%20%EF%BC%8C%E6%AF%8F%E8%A1%8C%E5%85%83%E7%B4%A0%E4%BB%A5%20%60%5C%5C%60%20%E7%BB%93%E5%B0%BE%EF%BC%8C%E6%9C%80%E5%90%8E%E4%BB%A5%20%60end%7Barray%7D%60%20%E7%BB%93%E6%9D%9F%E6%95%B0%E7%BB%84%E3%80%82%0A%E4%BD%BF%E7%94%A8%E5%8D%95%E4%B8%AA%E6%95%B0%E7%BB%84%E6%88%96%E8%A1%A8%E6%A0%BC%E6%97%B6%E6%97%A0%E9%9C%80%E5%A3%B0%E6%98%8E%20%60%24%60%20%E6%88%96%20%60%24%24%60%20%E7%AC%A6%E5%8F%B7%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%5Cbegin%7Barray%7D%7Bc%7Clcr%7D%0An%20%26%20%5Ctext%7B%E5%B7%A6%E5%AF%B9%E9%BD%90%7D%20%26%20%5Ctext%7B%E5%B1%85%E4%B8%AD%E5%AF%B9%E9%BD%90%7D%20%26%20%5Ctext%7B%E5%8F%B3%E5%AF%B9%E9%BD%90%7D%20%5C%5C%0A%5Chline%0A1%20%26%200.24%20%26%201%20%26%20125%20%5C%5C%0A2%20%26%20-1%20%26%20189%20%26%20-8%20%5C%5C%0A3%20%26%20-20%20%26%202000%20%26%201%2B10i%0A%5Cend%7Barray%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7Barray%7D%7Bc%7Clcr%7D%0An%20%26%20%5Ctext%7B%E5%B7%A6%E5%AF%B9%E9%BD%90%7D%20%26%20%5Ctext%7B%E5%B1%85%E4%B8%AD%E5%AF%B9%E9%BD%90%7D%20%26%20%5Ctext%7B%E5%8F%B3%E5%AF%B9%E9%BD%90%7D%20%5C%5C%0A%5Chline%0A1%20%26%200.24%20%26%201%20%26%20125%20%5C%5C%0A2%20%26%20-1%20%26%20189%20%26%20-8%20%5C%5C%0A3%20%26%20-20%20%26%202000%20%26%201%2B10i%0A%5Cend%7Barray%7D%0A%0A%23%232%EF%BC%8E%E5%A6%82%E4%BD%95%E8%BE%93%E5%85%A5%E4%B8%80%E4%B8%AA%E5%B5%8C%E5%A5%97%E7%9A%84%E6%95%B0%E7%BB%84%E6%88%96%E8%A1%A8%E6%A0%BC%0A%0A%E5%A4%9A%E4%B8%AA%E6%95%B0%E7%BB%84%2F%E8%A1%A8%E6%A0%BC%E5%8F%AF%20**%E4%BA%92%E7%9B%B8%E5%B5%8C%E5%A5%97**%20%E5%B9%B6%E7%BB%84%E6%88%90%E4%B8%80%E7%BB%84%E6%95%B0%E7%BB%84%2F%E4%B8%80%E7%BB%84%E8%A1%A8%E6%A0%BC%E3%80%82%0A%E4%BD%BF%E7%94%A8%E5%B5%8C%E5%A5%97%E5%89%8D%E5%BF%85%E9%A1%BB%E5%A3%B0%E6%98%8E%20%60%24%24%60%20%E7%AC%A6%E5%8F%B7%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A#card=math&code=%7C%0A%0A%2A%2A%E4%B8%80%E4%B8%AA%20%60%5Bex%5D%60%20%E6%8C%87%E4%B8%80%E4%B8%AA%20%22X-Height%22%EF%BC%8C%E5%8D%B3x%E5%AD%97%E6%AF%8D%E9%AB%98%E5%BA%A6%E3%80%82%E5%8F%AF%E4%BB%A5%E6%A0%B9%E6%8D%AE%E6%83%85%E5%86%B5%E6%8C%87%E5%AE%9A%E5%A4%9A%E4%B8%AA%20%60%5Bex%5D%60%EF%BC%8C%E5%A6%82%20%60%5B3ex%5D%60%E3%80%81%60%5B4ex%5D%60%20%E7%AD%89%E3%80%82%2A%2A%0A%E5%85%B6%E5%AE%9E%E5%8F%AF%E4%BB%A5%E5%9C%A8%E4%BB%BB%E4%BD%95%E5%9C%B0%E6%96%B9%E4%BD%BF%E7%94%A8%20%60%5C%5C%5B2ex%5D%60%20%E8%AF%AD%E5%8F%A5%EF%BC%8C%E5%8F%AA%E8%A6%81%E4%BD%A0%E8%A7%89%E5%BE%97%E5%90%88%E9%80%82%E3%80%82%0A%0A%23%E4%BA%94%E3%80%81%E6%95%B0%E7%BB%84%E4%B8%8E%E8%A1%A8%E6%A0%BC%E4%BD%BF%E7%94%A8%E5%8F%82%E8%80%83%0A%0A%23%231%EF%BC%8E%E5%A6%82%E4%BD%95%E8%BE%93%E5%85%A5%E4%B8%80%E4%B8%AA%E6%95%B0%E7%BB%84%E6%88%96%E8%A1%A8%E6%A0%BC%0A%0A%E9%80%9A%E5%B8%B8%EF%BC%8C%E4%B8%80%E4%B8%AA%E6%A0%BC%E5%BC%8F%E5%8C%96%E5%90%8E%E7%9A%84%E8%A1%A8%E6%A0%BC%E6%AF%94%E5%8D%95%E7%BA%AF%E7%9A%84%E6%96%87%E5%AD%97%E6%88%96%E6%8E%92%E7%89%88%E5%90%8E%E7%9A%84%E6%96%87%E5%AD%97%E6%9B%B4%E5%85%B7%E6%9C%89%E5%8F%AF%E8%AF%BB%E6%80%A7%E3%80%82%E6%95%B0%E7%BB%84%E5%92%8C%E8%A1%A8%E6%A0%BC%E5%9D%87%E4%BB%A5%20%60begin%7Barray%7D%60%20%E5%BC%80%E5%A4%B4%EF%BC%8C%E5%B9%B6%E5%9C%A8%E5%85%B6%E5%90%8E%E5%AE%9A%E4%B9%89%E5%88%97%E6%95%B0%E5%8F%8A%E6%AF%8F%E4%B8%80%E5%88%97%E7%9A%84%E6%96%87%E6%9C%AC%E5%AF%B9%E9%BD%90%E5%B1%9E%E6%80%A7%EF%BC%8C%60c%60%20%60l%60%20%60r%60%20%E5%88%86%E5%88%AB%E4%BB%A3%E8%A1%A8%E5%B1%85%E4%B8%AD%E3%80%81%E5%B7%A6%E5%AF%B9%E9%BD%90%E5%8F%8A%E5%8F%B3%E5%AF%B9%E9%BD%90%E3%80%82%E8%8B%A5%E9%9C%80%E8%A6%81%E6%8F%92%E5%85%A5%E5%9E%82%E7%9B%B4%E5%88%86%E5%89%B2%E7%BA%BF%EF%BC%8C%E5%9C%A8%E5%AE%9A%E4%B9%89%E5%BC%8F%E4%B8%AD%E6%8F%92%E5%85%A5%20%60%7C%60%20%EF%BC%8C%E8%8B%A5%E8%A6%81%E6%8F%92%E5%85%A5%E6%B0%B4%E5%B9%B3%E5%88%86%E5%89%B2%E7%BA%BF%EF%BC%8C%E5%9C%A8%E4%B8%8B%E4%B8%80%E8%A1%8C%E8%BE%93%E5%85%A5%E5%89%8D%E6%8F%92%E5%85%A5%20%60%5Chline%60%20%E3%80%82%E4%B8%8E%E7%9F%A9%E9%98%B5%E7%9B%B8%E4%BC%BC%EF%BC%8C%E6%AF%8F%E8%A1%8C%E5%85%83%E7%B4%A0%E9%97%B4%E5%9D%87%E9%A1%BB%E8%A6%81%E6%8F%92%E5%85%A5%20%60%26%60%20%EF%BC%8C%E6%AF%8F%E8%A1%8C%E5%85%83%E7%B4%A0%E4%BB%A5%20%60%5C%5C%60%20%E7%BB%93%E5%B0%BE%EF%BC%8C%E6%9C%80%E5%90%8E%E4%BB%A5%20%60end%7Barray%7D%60%20%E7%BB%93%E6%9D%9F%E6%95%B0%E7%BB%84%E3%80%82%0A%E4%BD%BF%E7%94%A8%E5%8D%95%E4%B8%AA%E6%95%B0%E7%BB%84%E6%88%96%E8%A1%A8%E6%A0%BC%E6%97%B6%E6%97%A0%E9%9C%80%E5%A3%B0%E6%98%8E%20%60%24%60%20%E6%88%96%20%60%24%24%60%20%E7%AC%A6%E5%8F%B7%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%5Cbegin%7Barray%7D%7Bc%7Clcr%7D%0An%20%26%20%5Ctext%7B%E5%B7%A6%E5%AF%B9%E9%BD%90%7D%20%26%20%5Ctext%7B%E5%B1%85%E4%B8%AD%E5%AF%B9%E9%BD%90%7D%20%26%20%5Ctext%7B%E5%8F%B3%E5%AF%B9%E9%BD%90%7D%20%5C%5C%0A%5Chline%0A1%20%26%200.24%20%26%201%20%26%20125%20%5C%5C%0A2%20%26%20-1%20%26%20189%20%26%20-8%20%5C%5C%0A3%20%26%20-20%20%26%202000%20%26%201%2B10i%0A%5Cend%7Barray%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7Barray%7D%7Bc%7Clcr%7D%0An%20%26%20%5Ctext%7B%E5%B7%A6%E5%AF%B9%E9%BD%90%7D%20%26%20%5Ctext%7B%E5%B1%85%E4%B8%AD%E5%AF%B9%E9%BD%90%7D%20%26%20%5Ctext%7B%E5%8F%B3%E5%AF%B9%E9%BD%90%7D%20%5C%5C%0A%5Chline%0A1%20%26%200.24%20%26%201%20%26%20125%20%5C%5C%0A2%20%26%20-1%20%26%20189%20%26%20-8%20%5C%5C%0A3%20%26%20-20%20%26%202000%20%26%201%2B10i%0A%5Cend%7Barray%7D%0A%0A%23%232%EF%BC%8E%E5%A6%82%E4%BD%95%E8%BE%93%E5%85%A5%E4%B8%80%E4%B8%AA%E5%B5%8C%E5%A5%97%E7%9A%84%E6%95%B0%E7%BB%84%E6%88%96%E8%A1%A8%E6%A0%BC%0A%0A%E5%A4%9A%E4%B8%AA%E6%95%B0%E7%BB%84%2F%E8%A1%A8%E6%A0%BC%E5%8F%AF%20%2A%2A%E4%BA%92%E7%9B%B8%E5%B5%8C%E5%A5%97%2A%2A%20%E5%B9%B6%E7%BB%84%E6%88%90%E4%B8%80%E7%BB%84%E6%95%B0%E7%BB%84%2F%E4%B8%80%E7%BB%84%E8%A1%A8%E6%A0%BC%E3%80%82%0A%E4%BD%BF%E7%94%A8%E5%B5%8C%E5%A5%97%E5%89%8D%E5%BF%85%E9%A1%BB%E5%A3%B0%E6%98%8E%20%60%24%24%60%20%E7%AC%A6%E5%8F%B7%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A)

% outer vertical array of arrays 外层垂直表格

\begin{array}{c}

% inner horizontal array of arrays 内层水平表格

\begin{array}{cc}

% inner array of minimum values 内层"最小值"数组

\begin{array}{c|cccc}

\text{min} & 0 & 1 & 2 & 3\

\hline

0 & 0 & 0 & 0 & 0\

1 & 0 & 1 & 1 & 1\

2 & 0 & 1 & 2 & 2\

3 & 0 & 1 & 2 & 3

\end{array}

&

% inner array of maximum values 内层"最大值"数组

\begin{array}{c|cccc}

\text{max}&0&1&2&3\

\hline

0 & 0 & 1 & 2 & 3\

1 & 1 & 1 & 2 & 3\

2 & 2 & 2 & 2 & 3\

3 & 3 & 3 & 3 & 3

\end{array}

\end{array}

% 内层第一行表格组结束

\

% inner array of delta values 内层第二行Delta值数组

\begin{array}{c|cccc}

\Delta&0&1&2&3\

\hline

0 & 0 & 1 & 2 & 3\

1 & 1 & 0 & 1 & 2\

2 & 2 & 1 & 0 & 1\

3 & 3 & 2 & 1 & 0

\end{array}

% 内层第二行表格组结束

\end{array}

![](https://g.yuque.com/gr/latex?%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A#card=math&code=%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A)

% outer vertical array of arrays 外层垂直表格

\begin{array}{c}

% inner horizontal array of arrays 内层水平表格

\begin{array}{cc}

% inner array of minimum values 内层"最小值"数组

\begin{array}{c|cccc}

\text{min} & 0 & 1 & 2 & 3\

\hline

0 & 0 & 0 & 0 & 0\

1 & 0 & 1 & 1 & 1\

2 & 0 & 1 & 2 & 2\

3 & 0 & 1 & 2 & 3

\end{array}

&

% inner array of maximum values 内层"最大值"数组

\begin{array}{c|cccc}

\text{max}&0&1&2&3\

\hline

0 & 0 & 1 & 2 & 3\

1 & 1 & 1 & 2 & 3\

2 & 2 & 2 & 2 & 3\

3 & 3 & 3 & 3 & 3

\end{array}

\end{array}

% 内层第一行表格组结束

\

% inner array of delta values 内层第二行Delta值数组

\begin{array}{c|cccc}

\Delta&0&1&2&3\

\hline

0 & 0 & 1 & 2 & 3\

1 & 1 & 0 & 1 & 2\

2 & 2 & 1 & 0 & 1\

3 & 3 & 2 & 1 & 0

\end{array}

% 内层第二行表格组结束

\end{array}

![](https://g.yuque.com/gr/latex?%0A%23%233%EF%BC%8E%E5%A6%82%E4%BD%95%E8%BE%93%E5%85%A5%E4%B8%80%E4%B8%AA%E6%96%B9%E7%A8%8B%E7%BB%84%0A%0A%E4%BD%BF%E7%94%A8%20%60%5Cbegin%7Barray%7D%E2%80%A6%5Cend%7Barray%7D%60%20%E5%92%8C%20%60%5Cleft%5C%7B%E2%80%A6%5Cright.%60%20%E6%9D%A5%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E6%96%B9%E7%A8%8B%E7%BB%84%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A#card=math&code=%0A%23%233%EF%BC%8E%E5%A6%82%E4%BD%95%E8%BE%93%E5%85%A5%E4%B8%80%E4%B8%AA%E6%96%B9%E7%A8%8B%E7%BB%84%0A%0A%E4%BD%BF%E7%94%A8%20%60%5Cbegin%7Barray%7D%E2%80%A6%5Cend%7Barray%7D%60%20%E5%92%8C%20%60%5Cleft%5C%7B%E2%80%A6%5Cright.%60%20%E6%9D%A5%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E6%96%B9%E7%A8%8B%E7%BB%84%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A)

\left{

\begin{array}{c}

a_1x+b_1y+c_1z=d_1 \

a_2x+b_2y+c_2z=d_2 \

a_3x+b_3y+c_3z=d_3

\end{array}

\right.

![](https://g.yuque.com/gr/latex?%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A#card=math&code=%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A)

\left{

\begin{array}{c}

a_1x+b_1y+c_1z=d_1 \

a_2x+b_2y+c_2z=d_2 \

a_3x+b_3y+c_3z=d_3

\end{array}

\right.

![](https://g.yuque.com/gr/latex?%0A%E6%88%96%E8%80%85%E4%BD%BF%E7%94%A8%E6%9D%A1%E4%BB%B6%E8%A1%A8%E8%BE%BE%E5%BC%8F%E7%BB%84%20%60%5Cbegin%7Bcases%7D%E2%80%A6%5Cend%7Bcases%7D%60%20%E6%9D%A5%E5%AE%9E%E7%8E%B0%E7%9B%B8%E5%90%8C%E6%95%88%E6%9E%9C%EF%BC%9A%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%5Cbegin%7Bcases%7D%0Aa_1x%2Bb_1y%2Bc_1z%3Dd_1%20%5C%5C%20%0Aa_2x%2Bb_2y%2Bc_2z%3Dd_2%20%5C%5C%20%0Aa_3x%2Bb_3y%2Bc_3z%3Dd_3%0A%5Cend%7Bcases%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7Bcases%7D%0Aa_1x%2Bb_1y%2Bc_1z%3Dd_1%20%5C%5C%20%0Aa_2x%2Bb_2y%2Bc_2z%3Dd_2%20%5C%5C%20%0Aa_3x%2Bb_3y%2Bc_3z%3Dd_3%0A%5Cend%7Bcases%7D%0A%0A%23%E5%85%AD%E3%80%81%E8%BF%9E%E5%88%86%E6%95%B0%E4%BD%BF%E7%94%A8%E5%8F%82%E8%80%83%0A%0A%23%231%EF%BC%8E%E5%A6%82%E4%BD%95%E8%BE%93%E5%85%A5%E4%B8%80%E4%B8%AA%E8%BF%9E%E5%88%86%E5%BC%8F%0A%0A%E5%B0%B1%E5%83%8F%E8%BE%93%E5%85%A5%E5%88%86%E5%BC%8F%E6%97%B6%E4%BD%BF%E7%94%A8%20%60%5Cfrac%60%20%E4%B8%80%E6%A0%B7%EF%BC%8C%E4%BD%BF%E7%94%A8%20%60%5Ccfrac%60%20%E6%9D%A5%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E8%BF%9E%E5%88%86%E6%95%B0%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A#card=math&code=%0A%E6%88%96%E8%80%85%E4%BD%BF%E7%94%A8%E6%9D%A1%E4%BB%B6%E8%A1%A8%E8%BE%BE%E5%BC%8F%E7%BB%84%20%60%5Cbegin%7Bcases%7D%E2%80%A6%5Cend%7Bcases%7D%60%20%E6%9D%A5%E5%AE%9E%E7%8E%B0%E7%9B%B8%E5%90%8C%E6%95%88%E6%9E%9C%EF%BC%9A%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%5Cbegin%7Bcases%7D%0Aa_1x%2Bb_1y%2Bc_1z%3Dd_1%20%5C%5C%20%0Aa_2x%2Bb_2y%2Bc_2z%3Dd_2%20%5C%5C%20%0Aa_3x%2Bb_3y%2Bc_3z%3Dd_3%0A%5Cend%7Bcases%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7Bcases%7D%0Aa_1x%2Bb_1y%2Bc_1z%3Dd_1%20%5C%5C%20%0Aa_2x%2Bb_2y%2Bc_2z%3Dd_2%20%5C%5C%20%0Aa_3x%2Bb_3y%2Bc_3z%3Dd_3%0A%5Cend%7Bcases%7D%0A%0A%23%E5%85%AD%E3%80%81%E8%BF%9E%E5%88%86%E6%95%B0%E4%BD%BF%E7%94%A8%E5%8F%82%E8%80%83%0A%0A%23%231%EF%BC%8E%E5%A6%82%E4%BD%95%E8%BE%93%E5%85%A5%E4%B8%80%E4%B8%AA%E8%BF%9E%E5%88%86%E5%BC%8F%0A%0A%E5%B0%B1%E5%83%8F%E8%BE%93%E5%85%A5%E5%88%86%E5%BC%8F%E6%97%B6%E4%BD%BF%E7%94%A8%20%60%5Cfrac%60%20%E4%B8%80%E6%A0%B7%EF%BC%8C%E4%BD%BF%E7%94%A8%20%60%5Ccfrac%60%20%E6%9D%A5%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E8%BF%9E%E5%88%86%E6%95%B0%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A)

x = a_0 + \cfrac{1^2}{a_1

+ \cfrac{2^2}{a_2

+ \cfrac{3^2}{a_3 + \cfrac{4^4}{a_4 + \cdots}}}}

![](https://g.yuque.com/gr/latex?%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A#card=math&code=%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A)

x = a_0 + \cfrac{1^2}{a_1

+ \cfrac{2^2}{a_2

+ \cfrac{3^2}{a_3 + \cfrac{4^4}{a_4 + \cdots}}}}

![](https://g.yuque.com/gr/latex?%0A%E4%B8%8D%E8%A6%81%E4%BD%BF%E7%94%A8%E6%99%AE%E9%80%9A%E7%9A%84%20%60%5Cfrac%60%20%E6%88%96%20%60%5Cover%60%20%E6%9D%A5%E5%88%9B%E5%BB%BA%EF%BC%8C%E5%90%A6%E5%88%99%E4%BC%9A%E7%9C%8B%E8%B5%B7%E6%9D%A5%20**%E5%BE%88%E6%81%B6%E5%BF%83**%20%E3%80%82%0A%0A-%20%E5%8F%8D%E4%BE%8B%EF%BC%9A%0A%60%60%60%0A#card=math&code=%0A%E4%B8%8D%E8%A6%81%E4%BD%BF%E7%94%A8%E6%99%AE%E9%80%9A%E7%9A%84%20%60%5Cfrac%60%20%E6%88%96%20%60%5Cover%60%20%E6%9D%A5%E5%88%9B%E5%BB%BA%EF%BC%8C%E5%90%A6%E5%88%99%E4%BC%9A%E7%9C%8B%E8%B5%B7%E6%9D%A5%20%2A%2A%E5%BE%88%E6%81%B6%E5%BF%83%2A%2A%20%E3%80%82%0A%0A-%20%E5%8F%8D%E4%BE%8B%EF%BC%9A%0A%60%60%60%0A)

x = a_0 + \frac{1^2}{a_1

+ \frac{2^2}{a_2

+ \frac{3^2}{a_3 + \frac{4^4}{a_4 + \cdots}}}}

![](https://g.yuque.com/gr/latex?%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A#card=math&code=%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A)

x = a_0 + \frac{1^2}{a_1

+ \frac{2^2}{a_2

+ \frac{3^2}{a_3 + \frac{4^4}{a_4 + \cdots}}}}

![](https://g.yuque.com/gr/latex?%0A%E5%BD%93%E7%84%B6%EF%BC%8C%E4%BD%A0%E5%8F%AF%E4%BB%A5%E4%BD%BF%E7%94%A8%20%60%5Cfrac%60%20%E6%9D%A5%E8%A1%A8%E8%BE%BE%E8%BF%9E%E5%88%86%E6%95%B0%E7%9A%84%20**%E7%B4%A7%E7%BC%A9%E8%AE%B0%E6%B3%95**%20%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A#card=math&code=%0A%E5%BD%93%E7%84%B6%EF%BC%8C%E4%BD%A0%E5%8F%AF%E4%BB%A5%E4%BD%BF%E7%94%A8%20%60%5Cfrac%60%20%E6%9D%A5%E8%A1%A8%E8%BE%BE%E8%BF%9E%E5%88%86%E6%95%B0%E7%9A%84%20%2A%2A%E7%B4%A7%E7%BC%A9%E8%AE%B0%E6%B3%95%2A%2A%20%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A)

x = a_0 + \frac{1^2}{a_1+}

\frac{2^2}{a_2+}

\frac{3^2}{a_3 +} \frac{4^4}{a_4 +} \cdots

![](https://g.yuque.com/gr/latex?%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A#card=math&code=%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A)

x = a_0 + \frac{1^2}{a_1+}

\frac{2^2}{a_2+}

\frac{3^2}{a_3 +} \frac{4^4}{a_4 +} \cdots

![](https://g.yuque.com/gr/latex?%0A%E8%BF%9E%E5%88%86%E6%95%B0%E9%80%9A%E5%B8%B8%E9%83%BD%E5%A4%AA%E5%A4%A7%E4%BB%A5%E8%87%B3%E4%BA%8E%E4%B8%8D%E6%98%93%E6%8E%92%E7%89%88%EF%BC%8C%E6%89%80%E4%BB%A5%E5%BB%BA%E8%AE%AE%E5%9C%A8%E8%BF%9E%E5%88%86%E6%95%B0%E5%89%8D%E5%90%8E%E5%A3%B0%E6%98%8E%20%60%24%24%60%20%E7%AC%A6%E5%8F%B7%EF%BC%8C%E6%88%96%E4%BD%BF%E7%94%A8%E5%83%8F%20%60%5Ba0%3Ba1%2Ca2%2Ca3%2C%E2%80%A6%5D%60%20%E4%B8%80%E6%A0%B7%E7%9A%84%E7%B4%A7%E7%BC%A9%E8%AE%B0%E6%B3%95%E3%80%82%0A%0A%23%E4%B8%83%E3%80%81%E4%BA%A4%E6%8D%A2%E5%9B%BE%E8%A1%A8%E4%BD%BF%E7%94%A8%E5%8F%82%E8%80%83%0A%0A%23%231%EF%BC%8E%E5%A6%82%E4%BD%95%E8%BE%93%E5%85%A5%E4%B8%80%E4%B8%AA%E4%BA%A4%E6%8D%A2%E5%9B%BE%E8%A1%A8%0A%0A%E4%BD%BF%E7%94%A8%E4%B8%80%E8%A1%8C%20%60%24%20%5Crequire%7BAMScd%7D%20%24%60%20%E8%AF%AD%E5%8F%A5%E6%9D%A5%E5%85%81%E8%AE%B8%E4%BA%A4%E6%8D%A2%E5%9B%BE%E8%A1%A8%E7%9A%84%E6%98%BE%E7%A4%BA%E3%80%82%0A%E5%A3%B0%E6%98%8E%E4%BA%A4%E6%8D%A2%E5%9B%BE%E8%A1%A8%E5%90%8E%EF%BC%8C%E8%AF%AD%E6%B3%95%E4%B8%8E%E7%9F%A9%E9%98%B5%E7%9B%B8%E4%BC%BC%EF%BC%8C%E5%9C%A8%E5%BC%80%E5%A4%B4%E4%BD%BF%E7%94%A8%20%60begin%7BCD%7D%60%EF%BC%8C%E5%9C%A8%E7%BB%93%E5%B0%BE%E4%BD%BF%E7%94%A8%20%60end%7BCD%7D%60%EF%BC%8C%E5%9C%A8%E4%B8%AD%E9%97%B4%E6%8F%92%E5%85%A5%E5%9B%BE%E8%A1%A8%E5%85%83%E7%B4%A0%EF%BC%8C%E6%AF%8F%E4%B8%AA%E5%85%83%E7%B4%A0%E4%B9%8B%E9%97%B4%E6%8F%92%E5%85%A5%20%60%26%60%20%EF%BC%8C%E5%B9%B6%E5%9C%A8%E6%AF%8F%E8%A1%8C%E7%BB%93%E5%B0%BE%E5%A4%84%E4%BD%BF%E7%94%A8%20%60%5C%5C%60%20%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%24%5Crequire%7BAMScd%7D%24%0A%5Cbegin%7BCD%7D%0A%20%20%20%20A%20%40%3Ea%3E%3E%20B%5C%5C%0A%20%20%20%20%40V%20b%20V%20V%5C%23%20%40VV%20c%20V%5C%5C%0A%20%20%20%20C%20%40%3E%3Ed%3E%20D%0A%5Cend%7BCD%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%24%5Crequire%7BAMScd%7D%24%0A%5Cbegin%7BCD%7D%0A%20%20%20%20A%20%40%3Ea%3E%3E%20B%5C%5C%0A%20%20%20%20%40V%20b%20V%20V%5C%23%20%40VV%20c%20V%5C%5C%0A%20%20%20%20C%20%40%3E%3Ed%3E%20D%0A%5Cend%7BCD%7D%0A%0A%E5%85%B6%E4%B8%AD%EF%BC%8C%60%40%3E%3E%3E%60%20%E4%BB%A3%E8%A1%A8%E5%8F%B3%E7%AE%AD%E5%A4%B4%E3%80%81%60%40%3C%3C%3C%60%20%E4%BB%A3%E8%A1%A8%E5%B7%A6%E7%AE%AD%E5%A4%B4%E3%80%81%60%40VVV%60%20%E4%BB%A3%E8%A1%A8%E4%B8%8B%E7%AE%AD%E5%A4%B4%E3%80%81%60%40AAA%60%20%E4%BB%A3%E8%A1%A8%E4%B8%8A%E7%AE%AD%E5%A4%B4%E3%80%81%60%40%3D%60%20%E4%BB%A3%E8%A1%A8%E6%B0%B4%E5%B9%B3%E5%8F%8C%E5%AE%9E%E7%BA%BF%E3%80%81%60%40%7C%60%20%E4%BB%A3%E8%A1%A8%E7%AB%96%E7%9B%B4%E5%8F%8C%E5%AE%9E%E7%BA%BF%E3%80%81%60%40.%60%E4%BB%A3%E8%A1%A8%E6%B2%A1%E6%9C%89%E7%AE%AD%E5%A4%B4%E3%80%82%0A%E5%9C%A8%20%60%40%3E%3E%3E%60%20%E7%9A%84%20%60%3E%3E%3E%60%20%E4%B9%8B%E9%97%B4%E4%BB%BB%E6%84%8F%E6%8F%92%E5%85%A5%E6%96%87%E5%AD%97%E5%8D%B3%E4%BB%A3%E8%A1%A8%E8%AF%A5%E7%AE%AD%E5%A4%B4%E7%9A%84%E6%B3%A8%E9%87%8A%E6%96%87%E5%AD%97%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%5Cbegin%7BCD%7D%0A%20%20%20%20A%20%40%3E%3E%3E%20B%20%40%3E%7B%5Ctext%7Bvery%20long%20label%7D%7D%3E%3E%20C%20%5C%5C%0A%20%20%20%20%40.%20%40AAA%20%40%7C%20%5C%5C%0A%20%20%20%20D%20%40%3D%20E%20%40%3C%3C%3C%20F%0A%5Cend%7BCD%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7BCD%7D%0A%20%20%20%20A%20%40%3E%3E%3E%20B%20%40%3E%7B%5Ctext%7Bvery%20long%20label%7D%7D%3E%3E%20C%20%5C%5C%0A%20%20%20%20%40.%20%40AAA%20%40%7C%20%5C%5C%0A%20%20%20%20D%20%40%3D%20E%20%40%3C%3C%3C%20F%0A%5Cend%7BCD%7D%0A%0A%E5%9C%A8%E6%9C%AC%E4%BE%8B%E4%B8%AD%EF%BC%8C%20%22very%20long%20label%22%E8%87%AA%E5%8A%A8%E5%BB%B6%E9%95%BF%E4%BA%86%E5%AE%83%E6%89%80%E5%9C%A8%E7%AE%AD%E5%A4%B4%E4%BB%A5%E5%8F%8A%E5%AF%B9%E5%BA%94%E7%AE%AD%E5%A4%B4%E7%9A%84%E9%95%BF%E5%BA%A6%E3%80%82%0A%0A%0A%23%E5%85%AB%E3%80%81%E4%B8%80%E4%BA%9B%E7%89%B9%E6%AE%8A%E7%9A%84%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A1%B9%0A%0A%7C**!!%20%E6%9C%AC%E6%AE%B5%E5%86%85%E5%AE%B9%E4%B8%BA%E4%B8%AA%E4%BA%BA%E7%BF%BB%E8%AF%91%EF%BC%8C%E5%8F%AF%E8%83%BD%E6%9C%89%E4%B8%8D%E5%87%86%E7%A1%AE%E4%B9%8B%E5%A4%84%20!!**%7C%0A%7C%3A--%3A%7C%0A%0AThese%20are%20issues%20that%20won't%20affect%20the%20correctness%20of%20formulas%2C%20but%20might%20make%20them%20look%20significantly%20better%20or%20worse.%20Beginners%20should%20feel%20free%20to%20ignore%20this%20advice%3B%20someone%20else%20will%20correct%20it%20for%20them%2C%20or%20more%20likely%20nobody%20will%20care.%0A%0A%E7%8E%B0%E5%9C%A8%E6%8C%87%E5%87%BA%E7%9A%84%E5%B0%8F%E9%97%AE%E9%A2%98%E5%B9%B6%E4%B8%8D%E4%BC%9A%E5%BD%B1%E5%93%8D%E6%96%B9%E7%A8%8B%E5%BC%8F%E5%8F%8A%E5%85%AC%E5%BC%8F%E7%AD%89%E7%9A%84%E6%AD%A3%E7%A1%AE%E6%98%BE%E7%A4%BA%EF%BC%8C%E4%BD%86%E8%83%BD%E8%AE%A9%E5%AE%83%E4%BB%AC%E7%9C%8B%E8%B5%B7%E6%9D%A5%E6%98%8E%E6%98%BE%E6%9B%B4%E5%A5%BD%E7%9C%8B%E3%80%82%E5%88%9D%E5%AD%A6%E8%80%85%E5%8F%AF%E6%97%A0%E8%A7%86%E8%BF%99%E4%BA%9B%E5%BB%BA%E8%AE%AE%EF%BC%8C%E8%87%AA%E7%84%B6%E4%BC%9A%E6%9C%89%E5%BC%BA%E8%BF%AB%E7%97%87%E6%82%A3%E8%80%85%E6%9B%BF%E4%BD%A0%E4%BB%AC%E6%94%B9%E6%8E%89%E5%AE%83%E7%9A%84%EF%BC%8C%E6%88%96%E8%80%85%E6%9B%B4%E5%8F%AF%E8%83%BD%E5%9C%B0%EF%BC%8C%E6%A0%B9%E6%9C%AC%E6%B2%A1%E4%BA%BA%E5%8F%91%E7%8E%B0%E8%BF%99%E4%BA%9B%E9%97%AE%E9%A2%98%E3%80%82%0A%0ADon't%20use%20%60%5Cfrac%60%20in%20exponents%20or%20limits%20of%20integrals%3B%20it%20looks%20bad%20and%20can%20be%20confusing%2C%20which%20is%20why%20it%20is%20rarely%20done%20in%20professional%20mathematical%20typesetting.%20Write%20the%20fraction%20horizontally%2C%20with%20a%20slash%3A%0A%0A%E5%9C%A8%E4%BB%A5e%E4%B8%BA%E5%BA%95%E7%9A%84%E6%8C%87%E6%95%B0%E5%87%BD%E6%95%B0%E3%80%81%E6%9E%81%E9%99%90%E5%92%8C%E7%A7%AF%E5%88%86%E4%B8%AD%E5%B0%BD%E9%87%8F%E4%B8%8D%E8%A6%81%E4%BD%BF%E7%94%A8%20%60%5Cfrac%60%20%E7%AC%A6%E5%8F%B7%EF%BC%9A%E5%AE%83%E4%BC%9A%E4%BD%BF%E6%95%B4%E6%AE%B5%E5%87%BD%E6%95%B0%E7%9C%8B%E8%B5%B7%E6%9D%A5%E5%BE%88%E6%80%AA%EF%BC%8C%E8%80%8C%E4%B8%94%E5%8F%AF%E8%83%BD%E4%BA%A7%E7%94%9F%E6%AD%A7%E4%B9%89%E3%80%82%E4%B9%9F%E6%AD%A3%E6%98%AF%E5%9B%A0%E6%AD%A4%E5%AE%83%E5%9C%A8%E4%B8%93%E4%B8%9A%E6%95%B0%E5%AD%A6%E6%8E%92%E7%89%88%E4%B8%AD%E5%87%A0%E4%B9%8E%E4%BB%8E%E4%B8%8D%E5%87%BA%E7%8E%B0%E3%80%82%0A%E6%A8%AA%E7%9D%80%E5%86%99%E8%BF%99%E4%BA%9B%E5%88%86%E5%BC%8F%EF%BC%8C%E4%B8%AD%E9%97%B4%E4%BD%BF%E7%94%A8%E6%96%9C%E7%BA%BF%E9%97%B4%E9%9A%94%20%60%2F%60%20%EF%BC%88%E7%94%A8%E6%96%9C%E7%BA%BF%E4%BB%A3%E6%9B%BF%E5%88%86%E6%95%B0%E7%BA%BF%EF%BC%89%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0Ae%5E%7Bi%5Cfrac%7B%5Cpi%7D2%7D%20%5Cquad%20e%5E%7B%5Cfrac%7Bi%5Cpi%7D2%7D%26%20e%5E%7Bi%5Cpi%2F2%7D%20%5C%5C%0A%5Cint_%7B-%5Cfrac%5Cpi2%7D%5E%5Cfrac%5Cpi2%20%5Csin%20x%5C%2Cdx%20%26%20%5Cint_%7B-%5Cpi%2F2%7D%5E%7B%5Cpi%2F2%7D%5Csin%20x%5C%2Cdx%20%5C%5C%0A%5Cend%7Barray%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0Ae%5E%7Bi%5Cfrac%7B%5Cpi%7D2%7D%20%5Cquad%20e%5E%7B%5Cfrac%7Bi%5Cpi%7D2%7D%26%20e%5E%7Bi%5Cpi%2F2%7D%20%5C%5C%0A%5Cint_%7B-%5Cfrac%5Cpi2%7D%5E%5Cfrac%5Cpi2%20%5Csin%20x%5C%2Cdx%20%26%20%5Cint_%7B-%5Cpi%2F2%7D%5E%7B%5Cpi%2F2%7D%5Csin%20x%5C%2Cdx%20%5C%5C%0A%5Cend%7Barray%7D%0A%0AThe%20%60%7C%60%20symbol%20has%20the%20wrong%20spacing%20when%20it%20is%20used%20as%20a%20divider%2C%20for%20example%20in%20set%20comprehensions.%20Use%20%60%5Cmid%60%20instead%3A%0A%0A%60%7C%60%20%E7%AC%A6%E5%8F%B7%E5%9C%A8%E8%A2%AB%E5%BD%93%E4%BD%9C%E5%88%86%E9%9A%94%E7%AC%A6%E6%97%B6%E4%BC%9A%E4%BA%A7%E7%94%9F%E9%94%99%E8%AF%AF%E7%9A%84%E9%97%B4%E9%9A%94%EF%BC%8C%E5%9B%A0%E6%AD%A4%E5%9C%A8%E9%9C%80%E8%A6%81%E5%88%86%E9%9A%94%E6%97%B6%E6%9C%80%E5%A5%BD%E4%BD%BF%E7%94%A8%20%60%5Cmid%60%20%E6%9D%A5%E4%BB%A3%E6%9B%BF%E5%AE%83%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%3A%0A%60%60%60%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0A%5C%7Bx%7Cx%5E2%5Cin%5CBbb%20Z%5C%7D%20%26%20%5C%7Bx%5Cmid%20x%5E2%5Cin%5CBbb%20Z%5C%7D%20%5C%5C%0A%5Cend%7Barray%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0A%5C%7Bx%7Cx%5E2%5Cin%5CBbb%20Z%5C%7D%20%26%20%5C%7Bx%5Cmid%20x%5E2%5Cin%5CBbb%20Z%5C%7D%20%5C%5C%0A%5Cend%7Barray%7D%0A%0AFor%20double%20and%20triple%20integrals%2C%20don't%20use%20%60%5Cint%5Cint%60%20or%20%60%5Cint%5Cint%5Cint%60.%20Instead%20use%20the%20special%20forms%20%60%5Ciint%60%20and%20%60%5Ciiint%60%3A%0A%0A%E4%BD%BF%E7%94%A8%E5%A4%9A%E9%87%8D%E7%A7%AF%E5%88%86%E7%AC%A6%E5%8F%B7%E6%97%B6%EF%BC%8C%E4%B8%8D%E8%A6%81%E5%A4%9A%E6%AC%A1%E4%BD%BF%E7%94%A8%20%60%5Cint%60%20%E6%9D%A5%E5%A3%B0%E6%98%8E%EF%BC%8C%E7%9B%B4%E6%8E%A5%E4%BD%BF%E7%94%A8%20%60%5Ciint%60%20%E6%9D%A5%E8%A1%A8%E7%A4%BA%20**%E4%BA%8C%E9%87%8D%E7%A7%AF%E5%88%86**%20%EF%BC%8C%E4%BD%BF%E7%94%A8%20%60%5Ciiint%60%20%E6%9D%A5%E8%A1%A8%E7%A4%BA%20**%E4%B8%89%E9%87%8D%E7%A7%AF%E5%88%86**%20%E7%AD%89%E3%80%82%E5%AF%B9%E4%BA%8E%E6%97%A0%E9%99%90%E6%AC%A1%E7%A7%AF%E5%88%86%EF%BC%8C%E5%8F%AF%E4%BB%A5%E7%94%A8%20%60%5Cint%20%5Ccdots%20%5Cint%60%20%E8%A1%A8%E7%A4%BA%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0A%5Cint%5Cint_S%20f(x)%5C%2Cdy%5C%2Cdx%20%26%20%5Ciint_S%20f(x)%5C%2Cdy%5C%2Cdx%20%5C%5C%0A%5Cint%5Cint%5Cint_V%20f(x)%5C%2Cdz%5C%2Cdy%5C%2Cdx%20%26%20%5Ciiint_V%20f(x)%5C%2Cdz%5C%2Cdy%5C%2Cdx%0A%5Cend%7Barray%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0A%5Cint%5Cint_S%20f(x)%5C%2Cdy%5C%2Cdx%20%26%20%5Ciint_S%20f(x)%5C%2Cdy%5C%2Cdx%20%5C%5C%0A%5Cint%5Cint%5Cint_V%20f(x)%5C%2Cdz%5C%2Cdy%5C%2Cdx%20%26%20%5Ciiint_V%20f(x)%5C%2Cdz%5C%2Cdy%5C%2Cdx%0A%5Cend%7Barray%7D%0A%0A%24%24%E6%97%A0%E9%99%90%E6%AC%A1%E7%A7%AF%E5%88%86%EF%BC%9A%5Cint%20%5Ccdots%20%5Cint#card=math&code=%0A%E8%BF%9E%E5%88%86%E6%95%B0%E9%80%9A%E5%B8%B8%E9%83%BD%E5%A4%AA%E5%A4%A7%E4%BB%A5%E8%87%B3%E4%BA%8E%E4%B8%8D%E6%98%93%E6%8E%92%E7%89%88%EF%BC%8C%E6%89%80%E4%BB%A5%E5%BB%BA%E8%AE%AE%E5%9C%A8%E8%BF%9E%E5%88%86%E6%95%B0%E5%89%8D%E5%90%8E%E5%A3%B0%E6%98%8E%20%60%24%24%60%20%E7%AC%A6%E5%8F%B7%EF%BC%8C%E6%88%96%E4%BD%BF%E7%94%A8%E5%83%8F%20%60%5Ba0%3Ba1%2Ca2%2Ca3%2C%E2%80%A6%5D%60%20%E4%B8%80%E6%A0%B7%E7%9A%84%E7%B4%A7%E7%BC%A9%E8%AE%B0%E6%B3%95%E3%80%82%0A%0A%23%E4%B8%83%E3%80%81%E4%BA%A4%E6%8D%A2%E5%9B%BE%E8%A1%A8%E4%BD%BF%E7%94%A8%E5%8F%82%E8%80%83%0A%0A%23%231%EF%BC%8E%E5%A6%82%E4%BD%95%E8%BE%93%E5%85%A5%E4%B8%80%E4%B8%AA%E4%BA%A4%E6%8D%A2%E5%9B%BE%E8%A1%A8%0A%0A%E4%BD%BF%E7%94%A8%E4%B8%80%E8%A1%8C%20%60%24%20%5Crequire%7BAMScd%7D%20%24%60%20%E8%AF%AD%E5%8F%A5%E6%9D%A5%E5%85%81%E8%AE%B8%E4%BA%A4%E6%8D%A2%E5%9B%BE%E8%A1%A8%E7%9A%84%E6%98%BE%E7%A4%BA%E3%80%82%0A%E5%A3%B0%E6%98%8E%E4%BA%A4%E6%8D%A2%E5%9B%BE%E8%A1%A8%E5%90%8E%EF%BC%8C%E8%AF%AD%E6%B3%95%E4%B8%8E%E7%9F%A9%E9%98%B5%E7%9B%B8%E4%BC%BC%EF%BC%8C%E5%9C%A8%E5%BC%80%E5%A4%B4%E4%BD%BF%E7%94%A8%20%60begin%7BCD%7D%60%EF%BC%8C%E5%9C%A8%E7%BB%93%E5%B0%BE%E4%BD%BF%E7%94%A8%20%60end%7BCD%7D%60%EF%BC%8C%E5%9C%A8%E4%B8%AD%E9%97%B4%E6%8F%92%E5%85%A5%E5%9B%BE%E8%A1%A8%E5%85%83%E7%B4%A0%EF%BC%8C%E6%AF%8F%E4%B8%AA%E5%85%83%E7%B4%A0%E4%B9%8B%E9%97%B4%E6%8F%92%E5%85%A5%20%60%26%60%20%EF%BC%8C%E5%B9%B6%E5%9C%A8%E6%AF%8F%E8%A1%8C%E7%BB%93%E5%B0%BE%E5%A4%84%E4%BD%BF%E7%94%A8%20%60%5C%5C%60%20%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%24%5Crequire%7BAMScd%7D%24%0A%5Cbegin%7BCD%7D%0A%20%20%20%20A%20%40%3Ea%3E%3E%20B%5C%5C%0A%20%20%20%20%40V%20b%20V%20V%5C%23%20%40VV%20c%20V%5C%5C%0A%20%20%20%20C%20%40%3E%3Ed%3E%20D%0A%5Cend%7BCD%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%24%5Crequire%7BAMScd%7D%24%0A%5Cbegin%7BCD%7D%0A%20%20%20%20A%20%40%3Ea%3E%3E%20B%5C%5C%0A%20%20%20%20%40V%20b%20V%20V%5C%23%20%40VV%20c%20V%5C%5C%0A%20%20%20%20C%20%40%3E%3Ed%3E%20D%0A%5Cend%7BCD%7D%0A%0A%E5%85%B6%E4%B8%AD%EF%BC%8C%60%40%3E%3E%3E%60%20%E4%BB%A3%E8%A1%A8%E5%8F%B3%E7%AE%AD%E5%A4%B4%E3%80%81%60%40%3C%3C%3C%60%20%E4%BB%A3%E8%A1%A8%E5%B7%A6%E7%AE%AD%E5%A4%B4%E3%80%81%60%40VVV%60%20%E4%BB%A3%E8%A1%A8%E4%B8%8B%E7%AE%AD%E5%A4%B4%E3%80%81%60%40AAA%60%20%E4%BB%A3%E8%A1%A8%E4%B8%8A%E7%AE%AD%E5%A4%B4%E3%80%81%60%40%3D%60%20%E4%BB%A3%E8%A1%A8%E6%B0%B4%E5%B9%B3%E5%8F%8C%E5%AE%9E%E7%BA%BF%E3%80%81%60%40%7C%60%20%E4%BB%A3%E8%A1%A8%E7%AB%96%E7%9B%B4%E5%8F%8C%E5%AE%9E%E7%BA%BF%E3%80%81%60%40.%60%E4%BB%A3%E8%A1%A8%E6%B2%A1%E6%9C%89%E7%AE%AD%E5%A4%B4%E3%80%82%0A%E5%9C%A8%20%60%40%3E%3E%3E%60%20%E7%9A%84%20%60%3E%3E%3E%60%20%E4%B9%8B%E9%97%B4%E4%BB%BB%E6%84%8F%E6%8F%92%E5%85%A5%E6%96%87%E5%AD%97%E5%8D%B3%E4%BB%A3%E8%A1%A8%E8%AF%A5%E7%AE%AD%E5%A4%B4%E7%9A%84%E6%B3%A8%E9%87%8A%E6%96%87%E5%AD%97%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%5Cbegin%7BCD%7D%0A%20%20%20%20A%20%40%3E%3E%3E%20B%20%40%3E%7B%5Ctext%7Bvery%20long%20label%7D%7D%3E%3E%20C%20%5C%5C%0A%20%20%20%20%40.%20%40AAA%20%40%7C%20%5C%5C%0A%20%20%20%20D%20%40%3D%20E%20%40%3C%3C%3C%20F%0A%5Cend%7BCD%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7BCD%7D%0A%20%20%20%20A%20%40%3E%3E%3E%20B%20%40%3E%7B%5Ctext%7Bvery%20long%20label%7D%7D%3E%3E%20C%20%5C%5C%0A%20%20%20%20%40.%20%40AAA%20%40%7C%20%5C%5C%0A%20%20%20%20D%20%40%3D%20E%20%40%3C%3C%3C%20F%0A%5Cend%7BCD%7D%0A%0A%E5%9C%A8%E6%9C%AC%E4%BE%8B%E4%B8%AD%EF%BC%8C%20%22very%20long%20label%22%E8%87%AA%E5%8A%A8%E5%BB%B6%E9%95%BF%E4%BA%86%E5%AE%83%E6%89%80%E5%9C%A8%E7%AE%AD%E5%A4%B4%E4%BB%A5%E5%8F%8A%E5%AF%B9%E5%BA%94%E7%AE%AD%E5%A4%B4%E7%9A%84%E9%95%BF%E5%BA%A6%E3%80%82%0A%0A%0A%23%E5%85%AB%E3%80%81%E4%B8%80%E4%BA%9B%E7%89%B9%E6%AE%8A%E7%9A%84%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A1%B9%0A%0A%7C%2A%2A%21%21%20%E6%9C%AC%E6%AE%B5%E5%86%85%E5%AE%B9%E4%B8%BA%E4%B8%AA%E4%BA%BA%E7%BF%BB%E8%AF%91%EF%BC%8C%E5%8F%AF%E8%83%BD%E6%9C%89%E4%B8%8D%E5%87%86%E7%A1%AE%E4%B9%8B%E5%A4%84%20%21%21%2A%2A%7C%0A%7C%3A--%3A%7C%0A%0AThese%20are%20issues%20that%20won%27t%20affect%20the%20correctness%20of%20formulas%2C%20but%20might%20make%20them%20look%20significantly%20better%20or%20worse.%20Beginners%20should%20feel%20free%20to%20ignore%20this%20advice%3B%20someone%20else%20will%20correct%20it%20for%20them%2C%20or%20more%20likely%20nobody%20will%20care.%0A%0A%E7%8E%B0%E5%9C%A8%E6%8C%87%E5%87%BA%E7%9A%84%E5%B0%8F%E9%97%AE%E9%A2%98%E5%B9%B6%E4%B8%8D%E4%BC%9A%E5%BD%B1%E5%93%8D%E6%96%B9%E7%A8%8B%E5%BC%8F%E5%8F%8A%E5%85%AC%E5%BC%8F%E7%AD%89%E7%9A%84%E6%AD%A3%E7%A1%AE%E6%98%BE%E7%A4%BA%EF%BC%8C%E4%BD%86%E8%83%BD%E8%AE%A9%E5%AE%83%E4%BB%AC%E7%9C%8B%E8%B5%B7%E6%9D%A5%E6%98%8E%E6%98%BE%E6%9B%B4%E5%A5%BD%E7%9C%8B%E3%80%82%E5%88%9D%E5%AD%A6%E8%80%85%E5%8F%AF%E6%97%A0%E8%A7%86%E8%BF%99%E4%BA%9B%E5%BB%BA%E8%AE%AE%EF%BC%8C%E8%87%AA%E7%84%B6%E4%BC%9A%E6%9C%89%E5%BC%BA%E8%BF%AB%E7%97%87%E6%82%A3%E8%80%85%E6%9B%BF%E4%BD%A0%E4%BB%AC%E6%94%B9%E6%8E%89%E5%AE%83%E7%9A%84%EF%BC%8C%E6%88%96%E8%80%85%E6%9B%B4%E5%8F%AF%E8%83%BD%E5%9C%B0%EF%BC%8C%E6%A0%B9%E6%9C%AC%E6%B2%A1%E4%BA%BA%E5%8F%91%E7%8E%B0%E8%BF%99%E4%BA%9B%E9%97%AE%E9%A2%98%E3%80%82%0A%0ADon%27t%20use%20%60%5Cfrac%60%20in%20exponents%20or%20limits%20of%20integrals%3B%20it%20looks%20bad%20and%20can%20be%20confusing%2C%20which%20is%20why%20it%20is%20rarely%20done%20in%20professional%20mathematical%20typesetting.%20Write%20the%20fraction%20horizontally%2C%20with%20a%20slash%3A%0A%0A%E5%9C%A8%E4%BB%A5e%E4%B8%BA%E5%BA%95%E7%9A%84%E6%8C%87%E6%95%B0%E5%87%BD%E6%95%B0%E3%80%81%E6%9E%81%E9%99%90%E5%92%8C%E7%A7%AF%E5%88%86%E4%B8%AD%E5%B0%BD%E9%87%8F%E4%B8%8D%E8%A6%81%E4%BD%BF%E7%94%A8%20%60%5Cfrac%60%20%E7%AC%A6%E5%8F%B7%EF%BC%9A%E5%AE%83%E4%BC%9A%E4%BD%BF%E6%95%B4%E6%AE%B5%E5%87%BD%E6%95%B0%E7%9C%8B%E8%B5%B7%E6%9D%A5%E5%BE%88%E6%80%AA%EF%BC%8C%E8%80%8C%E4%B8%94%E5%8F%AF%E8%83%BD%E4%BA%A7%E7%94%9F%E6%AD%A7%E4%B9%89%E3%80%82%E4%B9%9F%E6%AD%A3%E6%98%AF%E5%9B%A0%E6%AD%A4%E5%AE%83%E5%9C%A8%E4%B8%93%E4%B8%9A%E6%95%B0%E5%AD%A6%E6%8E%92%E7%89%88%E4%B8%AD%E5%87%A0%E4%B9%8E%E4%BB%8E%E4%B8%8D%E5%87%BA%E7%8E%B0%E3%80%82%0A%E6%A8%AA%E7%9D%80%E5%86%99%E8%BF%99%E4%BA%9B%E5%88%86%E5%BC%8F%EF%BC%8C%E4%B8%AD%E9%97%B4%E4%BD%BF%E7%94%A8%E6%96%9C%E7%BA%BF%E9%97%B4%E9%9A%94%20%60%2F%60%20%EF%BC%88%E7%94%A8%E6%96%9C%E7%BA%BF%E4%BB%A3%E6%9B%BF%E5%88%86%E6%95%B0%E7%BA%BF%EF%BC%89%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0Ae%5E%7Bi%5Cfrac%7B%5Cpi%7D2%7D%20%5Cquad%20e%5E%7B%5Cfrac%7Bi%5Cpi%7D2%7D%26%20e%5E%7Bi%5Cpi%2F2%7D%20%5C%5C%0A%5Cint_%7B-%5Cfrac%5Cpi2%7D%5E%5Cfrac%5Cpi2%20%5Csin%20x%5C%2Cdx%20%26%20%5Cint_%7B-%5Cpi%2F2%7D%5E%7B%5Cpi%2F2%7D%5Csin%20x%5C%2Cdx%20%5C%5C%0A%5Cend%7Barray%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0Ae%5E%7Bi%5Cfrac%7B%5Cpi%7D2%7D%20%5Cquad%20e%5E%7B%5Cfrac%7Bi%5Cpi%7D2%7D%26%20e%5E%7Bi%5Cpi%2F2%7D%20%5C%5C%0A%5Cint_%7B-%5Cfrac%5Cpi2%7D%5E%5Cfrac%5Cpi2%20%5Csin%20x%5C%2Cdx%20%26%20%5Cint_%7B-%5Cpi%2F2%7D%5E%7B%5Cpi%2F2%7D%5Csin%20x%5C%2Cdx%20%5C%5C%0A%5Cend%7Barray%7D%0A%0AThe%20%60%7C%60%20symbol%20has%20the%20wrong%20spacing%20when%20it%20is%20used%20as%20a%20divider%2C%20for%20example%20in%20set%20comprehensions.%20Use%20%60%5Cmid%60%20instead%3A%0A%0A%60%7C%60%20%E7%AC%A6%E5%8F%B7%E5%9C%A8%E8%A2%AB%E5%BD%93%E4%BD%9C%E5%88%86%E9%9A%94%E7%AC%A6%E6%97%B6%E4%BC%9A%E4%BA%A7%E7%94%9F%E9%94%99%E8%AF%AF%E7%9A%84%E9%97%B4%E9%9A%94%EF%BC%8C%E5%9B%A0%E6%AD%A4%E5%9C%A8%E9%9C%80%E8%A6%81%E5%88%86%E9%9A%94%E6%97%B6%E6%9C%80%E5%A5%BD%E4%BD%BF%E7%94%A8%20%60%5Cmid%60%20%E6%9D%A5%E4%BB%A3%E6%9B%BF%E5%AE%83%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%3A%0A%60%60%60%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0A%5C%7Bx%7Cx%5E2%5Cin%5CBbb%20Z%5C%7D%20%26%20%5C%7Bx%5Cmid%20x%5E2%5Cin%5CBbb%20Z%5C%7D%20%5C%5C%0A%5Cend%7Barray%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0A%5C%7Bx%7Cx%5E2%5Cin%5CBbb%20Z%5C%7D%20%26%20%5C%7Bx%5Cmid%20x%5E2%5Cin%5CBbb%20Z%5C%7D%20%5C%5C%0A%5Cend%7Barray%7D%0A%0AFor%20double%20and%20triple%20integrals%2C%20don%27t%20use%20%60%5Cint%5Cint%60%20or%20%60%5Cint%5Cint%5Cint%60.%20Instead%20use%20the%20special%20forms%20%60%5Ciint%60%20and%20%60%5Ciiint%60%3A%0A%0A%E4%BD%BF%E7%94%A8%E5%A4%9A%E9%87%8D%E7%A7%AF%E5%88%86%E7%AC%A6%E5%8F%B7%E6%97%B6%EF%BC%8C%E4%B8%8D%E8%A6%81%E5%A4%9A%E6%AC%A1%E4%BD%BF%E7%94%A8%20%60%5Cint%60%20%E6%9D%A5%E5%A3%B0%E6%98%8E%EF%BC%8C%E7%9B%B4%E6%8E%A5%E4%BD%BF%E7%94%A8%20%60%5Ciint%60%20%E6%9D%A5%E8%A1%A8%E7%A4%BA%20%2A%2A%E4%BA%8C%E9%87%8D%E7%A7%AF%E5%88%86%2A%2A%20%EF%BC%8C%E4%BD%BF%E7%94%A8%20%60%5Ciiint%60%20%E6%9D%A5%E8%A1%A8%E7%A4%BA%20%2A%2A%E4%B8%89%E9%87%8D%E7%A7%AF%E5%88%86%2A%2A%20%E7%AD%89%E3%80%82%E5%AF%B9%E4%BA%8E%E6%97%A0%E9%99%90%E6%AC%A1%E7%A7%AF%E5%88%86%EF%BC%8C%E5%8F%AF%E4%BB%A5%E7%94%A8%20%60%5Cint%20%5Ccdots%20%5Cint%60%20%E8%A1%A8%E7%A4%BA%E3%80%82%0A%0A-%20%E4%BE%8B%E5%AD%90%EF%BC%9A%0A%60%60%60%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0A%5Cint%5Cint_S%20f%28x%29%5C%2Cdy%5C%2Cdx%20%26%20%5Ciint_S%20f%28x%29%5C%2Cdy%5C%2Cdx%20%5C%5C%0A%5Cint%5Cint%5Cint_V%20f%28x%29%5C%2Cdz%5C%2Cdy%5C%2Cdx%20%26%20%5Ciiint_V%20f%28x%29%5C%2Cdz%5C%2Cdy%5C%2Cdx%0A%5Cend%7Barray%7D%0A%60%60%60%0A%0A-%20%E6%98%BE%E7%A4%BA%EF%BC%9A%0A%5Cbegin%7Barray%7D%7Bcc%7D%0A%5Cmathrm%7BBad%7D%20%26%20%5Cmathrm%7BBetter%7D%20%5C%5C%0A%5Chline%20%5C%5C%0A%5Cint%5Cint_S%20f%28x%29%5C%2Cdy%5C%2Cdx%20%26%20%5Ciint_S%20f%28x%29%5C%2Cdy%5C%2Cdx%20%5C%5C%0A%5Cint%5Cint%5Cint_V%20f%28x%29%5C%2Cdz%5C%2Cdy%5C%2Cdx%20%26%20%5Ciiint_V%20f%28x%29%5C%2Cdz%5C%2Cdy%5C%2Cdx%0A%5Cend%7Barray%7D%0A%0A%24%24%E6%97%A0%E9%99%90%E6%AC%A1%E7%A7%AF%E5%88%86%EF%BC%9A%5Cint%20%5Ccdots%20%5Cint)

Use `\,`, to insert a thin space before differentials; without this ![](https://g.yuque.com/gr/latex?%5CTeX#card=math&code=%5CTeX) will mash them together:

在微分符号前加入 `\,` 来插入一个小的间隔空隙；没有 `\,` 符号的话，![](https://g.yuque.com/gr/latex?%5CTeX#card=math&code=%5CTeX) 将会把不同的微分符号堆在一起。

- 例子：

```
\begin{array}{cc}
\mathrm{Bad} & \mathrm{Better} \\
\hline \\
\iiint_V f(x){\rm d}z {\rm d}y {\rm d}x & \iiint_V f(x)\,{\rm d}z\,{\rm d}y\,{\rm d}x
\end{array}
```

- 显示：

\begin{array}{cc}

\mathrm{Bad} & \mathrm{Better} \

\hline \

\iiint_V f(x){\rm d}z {\rm d}y {\rm d}x & \iiint_V f(x),{\rm d}z,{\rm d}y,{\rm d}x

\end{array}

---

感谢您花费时间阅读这份指导手册，本手册内容可能有疏漏之处，欢迎更改指正。

更多语法请参见：[Cmd Markdown 简明语法手册](https://www.zybuluo.com/mdeditor?url=https://www.zybuluo.com/static/editor/md-help.markdown)，[Cmd Markdown 高阶语法手册](https://www.zybuluo.com/mdeditor?url=https://www.zybuluo.com/static/editor/md-help.markdown#cmd-markdown-%E9%AB%98%E9%98%B6%E8%AF%AD%E6%B3%95%E6%89%8B%E5%86%8C)。

祝您记录、阅读、分享愉快！

Drafted & Translated by [Eric P.](https://ericp.cn/)

2015-10-02
