# CSS 布局

## Display

`display` 属性的值，[参考MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/display)：

```css
/* precomposed values */
display: block;
display: inline;
display: inline-block;
display: flex;
display: inline-flex;
display: grid;
display: inline-grid;
display: flow-root;

/* box generation */
display: none;
display: contents;

/* two-value syntax */
display: block flow;
display: inline flow;
display: inline flow-root;
display: block flex;
display: inline flex;
display: block grid;
display: inline grid;
display: block flow-root;

/* other values */
display: table;
display: table-row; /* all table elements have an equivalent CSS display value */
display: list-item;

/* Global values */
display: inherit;
display: initial;
display: revert;
display: revert-layer;
display: unset;
```

### `none`

设置为 none 的元素既不会占据空间，也无法显示，相当于该元素不存在。 

### 三剑客比较

| 属性              | `inline`           | `block`                                                      | `inline-block` |
| ----------------- | ------------------ | ------------------------------------------------------------ | -------------- |
| width与height     | 无效               | 有效                                                         | 有效           |
| margin 与 padding | 左右有效，上下无效 | 有效                                                         | 有效           |
| 换行              | 共用一行           | 独占一行，默认会继承父元素的宽度，高度一般以子元素撑开的高度为准 | 共用一行       |

接下来我们重点讲 Flex 布局与 Grid 布局。

## Flex

请阅读：Chris Coyier. [A Complete Guide to Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/). 

### 基础知识

![00-basic-terminology](img/00-basic-terminology.svg)

- **容器与条目**。
- **主轴与交叉轴**。主轴：浮动的方向是主轴方向，交叉轴是与主轴垂直的方向。
- **起点与终点**。
- **盒子的尺寸**。

浮动布局不是单一的属性就能完成的，是一些列属性共同完成的，包括容器与条目。

### 配置容器属性

- 设置浮动显示。

```css
.container {
  display: flex; /* or inline-flex */
}
```

- 设置浮动方向与换行，[例子](https://codepen.io/team/css-tricks/pen/bEajLE/1ea1ef35d942d0041b0467b4d39888d3)。

```css
.container {
  display: flex; /* or inline-flex */
  flex-direction: row | row-reverse | column | column-reverse;
  flex-wrap: nowrap | wrap | wrap-reverse;
}
```

两个属性可以合并成：

```css
.container {
  display: flex;
  flex-flow: column wrap;
}
```

- 条目主轴排列方式

![justify-content](img/justify-content.svg)

```css
.container {
  display: flex;
  flex-flow: row wrap;
  justify-content:flex-start|flex-end|center|space-between|space-around|space-evenly
}
```

- 交叉轴方向排列方式

![align-items](img/align-items.svg)

`align-content`

![align-content](img/align-content.svg)

```css
.container {
  align-items: stretch|flex-start|flex-end|center|baseline;/*单行*/
  align-content:flex-start | flex-end | center | space-between | space-around | space-evenly | stretch;/*只有 wrap 才有效*/
}
```

- 间距

![gap](img/gap-1.svg)

```css
gap: 10px;
gap: 10px 20px; /*行间距，列间距*/
row-gap: 10px;
column-gap: 20px;
```

### 条目优先级

```css
.item {
  order: 5; /* default is 0 */
}
```

`order` 值越小越排在后面。

### 伸展与收缩

使用 `flex-grow` 调整空间占比。设置了`flex-grow` 的元素会获取剩下的空间。

使用 `flex-shrink: 3;` 属性使得该 `item` 相对于其项目的收缩 `1/3`。[示例](https://www.w3school.com.cn/cssref/pr_flex-shrink.asp)

### 调整自己位置

```css
align-self: auto | flex-start | flex-end | center | baseline | stretch;
```

## Grid

请阅读：Chris House. [A Complete Guide to Grid](https://css-tricks.com/snippets/css/complete-guide-grid/).

Grid 布局又称网格布局，他实际上就是把二维平面划分成行与列，形成网格。Gide 布局与 flex 布局一样，需要同时设置容器与项目的属性。

### 术语

![terms-grid-line](img/terms-grid-area.svg)

- **容器与条目**。容器用于存放条目。就好比书架上的书，停车场场的车。
- **栅格线**。即上图中的黑线。
- **栅格**。就是图中的每个格子。
- **轨道(Track)**。两条相邻网格线之间的空间，你可以把它们想象成网格的列或行。
- **区域（Area）**。任意网格区域，例如上图黄色区域。

### 容器属性

首先把 `display` 设置成 `grid`。

```css
.container {
  display: grid | inline-grid;
}
```

接下来用 `grid-template-columns` 与 `grid-template-rows` 划分网格。

```html
<div class="container">
  <div class="red">1</div>
  <div class="green">2</div>
  <div class="blue">3</div>
</div>
```

```css
.container {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr; /*repeat(3, 1fr)*/
}

.red {
  background-color: red;
}

.green {
  background-color: green;
}

.blue {
  background-color: blue;
}
```

上述例子的含义是 `fr` 表示把容器分为三列，每一列占一份。

**定义列的技巧：**

```css
grid-template-columns: repeat(auto-fill, 100px); /*每一列 100px, 列数就是容器宽度/每列宽度*/
grid-template-columns: 1fr 1fr minmax(100px, 1fr);/*列的最大与最小宽度*/
grid-template-columns: 100px auto 100px; /*自动宽度，一般是实际内容宽度*/
grid-template-columns: [c1] 100px [c2] 100px [c3] auto [c4]; /*网格线重命名*/
```

### 行与列的间距
```css
.container {
  row-gap: 20px;
  column-gap: 20px;
  gap: 20px 20px;
}
```

### 命名区域

```css
.container {
  display: grid;
  grid-template-columns: 50px 50px 50px 50px;
  grid-template-rows: auto;
  grid-template-areas: 
    "header header header header"
    "main main . sidebar"
    "footer footer footer footer";
}
```

### 对齐方式

元素内部：

```css
justify-items: start | end | center | stretch;
align-items: start | end | center | stretch;
place-items: center; /*快速居中*/
```

元素之间：

```css
justify-content: start | end | center | stretch | space-around | space-between | space-evenly;

align-content: start | end | center | stretch | space-around | space-between | space-evenly;
```

### 项目占位

通过指定项目左右或上下网格线的形式划分区域。

```css
.item-1 {
  grid-column-start: 2;
  grid-column-end: 4;
}

/* 命名网格线*/
.item-1 {
  grid-column-start: header-start;
  grid-column-end: header-end;
}

/*跨越网格线数，end = start+2*/
.item-1 {
  grid-column-start: span 2;
}

/* grid-column属性是grid-column-start和grid-column-end的合并简写形式
 * grid-row属性是grid-row-start属性和grid-row-end的合并简写形式。
 */
.item {
  grid-column: <start-line> / <end-line>;
  grid-row: <start-line> / <end-line>;
}

/*指定区域名*/
.item-1 {
  grid-area: e;
}
```

## BFC

官方定义：**BFC（Block Formatting Context）块格式化上下文**， 是Web页面的可视CSS渲染的一部分，是块盒子的布局过程发生的区域，也是浮动元素与其他元素交互的区域。

说人话：**BFC就是页面上的一个隔离的独立容器，容器里面的子元素不会影响到外面的元素。** 我们经常使用到BFC，只不过不知道它是BFC而已。

### BFC的性质

1. 内部的块级元素会在垂直方向一个接着一个地放置。
2. Box垂直方向上的距离由 margin 决定。属于同一个 BFC 的两个相邻的Box的margin会发生重叠。
3. 每个盒子的左外边框紧挨着包含块的左边框，即使浮动元素也是如此。
4. BFC的区域不会与float box重叠。
5. BFC就是页面上的一个隔离的独立容器，容器里面的子元素不会影响到外面的元素，反之亦然。
6. 计算BFC的高度时，浮动子元素也参与计算。

### 如何创建一个BFC

我们介绍接种常用的方法，[完整示例链接](https://developer.mozilla.org/en-US/docs/Web/Guide/CSS/Block_formatting_context)。

| 属性       | 值                                                  |
| ---------- | --------------------------------------------------- |
| `float`    | 除了 `none`，可以是：`left, rigth`                  |
| `position` | `absolute, fixed`                                   |
| `overflow` | 除了`visible, clip` ，可以是：`auto, sroll, hidden` |
| `display`  | 很多，例如：`inline-block,table-cell, flow-root`    |
| `contain`  | `layout, content, paint`                            |

 `display:flex;`或` display:inline-flex` 的直接孩子也是BFC，除了他们自己也是 `flex, grid, table`。

 `display:grid;`或` display:inline-grid` 的直接孩子也是BFC，除了他们自己也是 `flex, grid, table`。

# 参考文章

- WebChang. [CSS中的BFC是什么？怎么用？](https://juejin.cn/post/7031065166317879310). 稀土掘金.
- Chris Coyier. [A Complete Guide to Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/). 
- Chris House. [A Complete Guide to Grid](https://css-tricks.com/snippets/css/complete-guide-grid/).
- CUGGZ. [彻底理解CSS Flexbox布局，看这一篇就够了！](https://juejin.cn/post/7004622232378966046). 稀土掘金.
- CUGGZ. [高频前端面试题汇总之CSS篇](https://juejin.cn/post/6905539198107942919). 稀土掘金.
- [Box model in CSS](https://cssreference.io/box-model/)