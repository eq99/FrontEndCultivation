# HTML `a` 标签详解

`a` 标签是浏览器中最常见的标签，它最主要的作用是从一个页面跳转到另一个页面，那么除了超链接这个最主要的功能，你还知道 `a` 标签的哪些功能呢？

## `a` 标签的属性

### href

`a` 标签有一个至关重要的属性：`href`，如果没有设置这个属性，他就会和最普通的行内元素一样 `span` 一样，没有特殊的功能。

```html
<a href="https://bilibili.com">网址直接跳转</a> 
<a href="images/1.webp">相对路径跳转</a>
<a href="hello.exe">如果是打不开的文件，则下载之</a>
<a href="">空链接是刷新</a>
<a href="#">#是回到顶部</a>
<a href="javascript:;">禁止链接跳转</a>
```

`href` 的工作原理[^1]：

1）a 标签的 href 属性值如果是以 http 开头的，那么浏览器会马上去解释该网址：首先会在本地机器去找一个 hosts 文件， 如果在 hosts 文件上该域名没有对应的主机，那么浏览器就去到对应的 dns 服务器去寻找该域名对应的主机号。如果找到了对应的主机，那么该请求就会发给对应的主机，最后打开该页面。
2）如果 a 标签的 href 属性值没有以任何协议开头，那么浏览就会启动 file 协议解释器去解释该资源路径。
3）如果 a 标签的 href 属性值并不是以 http 协议，而且其他的一些协议，那么这时候浏览器就会到我们本地电脑的注册表中去查找是否有处理这种协议的应用程序，如果有，那么马上启动该应用程序处理该协议。常用的协议：

```html
<a href="tel:12345678">拨打电话</a>
<a href="mailto:12345678@qq.com">打开邮箱应用</a>
```

如果点击相关的链接，浏览器会弹出一个对话框，或者直接打开默认的应用程序。这个功能非常有意思，读者可以自己尝试一下。

###  target

`a` 标签最主要的功能是页面跳转，跳转的目标页面可以在当前窗口打开，也可以在新窗口打开，新窗口这个功能可以由 `target` 属性控制：

```html
<a href="https://bilibili.com" tatget="_blank">新窗口打开网页</a> 
<a href="https://bilibili.com" tatget="_self">当前窗口打开网页</a> 
```

【注意】属性值前面有一个下划线。

## `a` 标签的嵌套

`a` 标签是一个行内元素，可以当做内容用于其他块级元素的子元素，也在它的内部也可以嵌套块级元素，这样点击一块内容可以跳转到其他页面（卡片链接），**但是 `a` 标签不能嵌套另一个 `a` 标签**：

```html
<a href="">
  <h3>出Bug了</h3>
  <span class="box">
    <a href="">哈哈哈</a>
  </span>
</a>
```

在浏览器中被解析成：

```html
<a href>
  <h3>出Bug了</h3>
  <span class="box">
  </span>
</a>
<a href="">哈哈哈</a>
```

正常人也不会这么嵌套，但标签嵌套多了难免会犯傻，所以记住这个特例对调试找错有帮助。

## `a`  标签的样式

### 下划线

`a` 标签典型的特点是自带一个下划线。下划线效果是可以通过 css 设置的：

```css
a {
  text-decoration: none; /*去掉下划线*/
}

/*加上下划线*/
a {
  text-decoration: underline;
}

/*还能设置线型样式*/
a {
  text-decoration: underline dotted red;
}
```

更多样式可以参考：MDN [text-decoration](https://developer.mozilla.org/en-US/docs/Web/CSS/text-decoration).

### 字体颜色

`a` 标签的默认字体颜色是蓝色，我们知道，字体样式：例如字体大小，行高，颜色等属性是可以继承父元素的，但如果该元素带有某种浏览器自带的特殊样式（user agent style），那么该属性就不会继承，由于这个原因， `a` 标签的颜色设置需要我们自己初始化成：

```css
a {
  color: inherit;
}
```

这样 `a` 标签也可以继承父元素的颜色了，不需要单独设置。

## 盒子属性

`a` 标签是一个行内元素，它和其他行内元素一样：

- 设置宽高无效
- 设置上下外边距无效
- 设置上下内边距会把背景撑开，但不能撑开内容。

但有时候网页需要一个长的像按钮一样的链接，按钮需要设置高度与宽度、内外边距，可以在 a 标签外面套一层会计标签，也可以把它设置成行内块元素解决：

```css
a {
  display: inline-block;
  
  height: 32px;
  line-height: 32px;
  padding: 4px 16px;
  border-radius: 5px;
  
  /*辅助样式*/
  text-decoration: none;
  color: deeppink;
  background-color: lightpink;
}
```

### 动态效果

`a` 标签可以根据用户的行为设置某些样式，例如当鼠标停在链接上时，鼠标的形状会变成手型的。这个效果其实是：

```css
a:hover {
  cursor: pointer;
}

/*改成其他的就变了*/
a:hover {
  cursor: auto;
}
```

更多鼠标的指针样式可以参考 MDN [cursor](https://developer.mozilla.org/en-US/docs/Web/CSS/cursor)。

值得一提的是 `:hover`，这表示鼠标飘过该元素时的状态，称为伪类选择器。类似的还有：

```css
/* 带有 href 的a标签会生效, 用得比较少了*/
a:link {
  color: brown;
}

/* visited 点过的链接会生效，用的比较少了*/
a:visited {
  color: aquamarine;
}

/* 鼠标滑过时生效，使用非常多 */
a:hover {
  color: blueviolet;
}

/* 鼠标长按元素的生效 */
a:active {
  color: darkcyan;
}
```

一般来说，上述伪类样式只会用到一个，但是如果要同时设置以上四个样式时，需要注意声明的先后顺序。如果一个链接没有被访问过，那么它有可能同时拥有 link , hover 两个属性，不能让 link 写在后面，否则不管 hover 是否，都会显示link的样式；同样的道理，如果一个链接已经被访问过了，那么它有可能同时拥有visited , hover 两个属性，hover 要在 visited 的后面[^2]。

【技巧】记忆方法：条件苛刻的放在后面。

## 参考文章

[^1]:前端路上的小兵. [详解HTML的a标签（超链接标签） ](https://www.cnblogs.com/shcrk/p/9279960.html). 博客园.
[^2]:杜俊成. [a 标签的伪类的正确顺序，以及原因](https://www.cnblogs.com/dujuncheng/p/8a44d12f5dec0def9518321af4e71c22.html). 博客园. 



> ♥ 非常感谢大家的点赞与关注。文章整理自网络，若有疏漏请在评论区指正。
>
> ★ 本文开源（[Github链接](https://github.com/xiayulu/frontend-all-in-one)）欢迎参与贡献，转载需注明出处：小星森. [HTML a 标签详解](https://zhuanlan.zhihu.com/p/573177374). 知乎.
>
> ★ 商业合作请发私信或邮件：zuiaiqiansen@163.com，并注明主题：商业合作。
>
> ★ 宣传做的好，产品差不了，宣传做的棒，产品No.1。我是前端小星森，让用户看到最伟大的产品。本人会一点点前端，如果贵公司有一流的产品或服务需要前端工程师展示，或需要一个前端工程师实现您的创业梦想，请发邮件：zuiaiqiansen@163.com，并注明主题：招聘前端工程师。
