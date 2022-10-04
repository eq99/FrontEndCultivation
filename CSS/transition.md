# CSS 过度动画

## 什么是过度动画

过度动画指的是元素样式从一个状态到另一个状态的过度动画。这里涉及到一个关键的概念，就是元素的样式发生改变，样式怎么会发生改变呢？最长见的改变样式的方式就是给相同元素添加伪类样式：

```html
<div class="box"></div>
```

```css
.box {
  width: 50px;
  height: 50px;
  background-color: pink;
  transition-property: all;
  transition-duration: 0.8s;
}
.box:hover {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  background-color: skyblue;
}
```

当鼠标移动到元素上方时，元素就会使用`.box:hover` 里的样式，从普通样式到 `hover` 样式，或者从`hover` 样式回到普通样式，这里都涉及到样式的改变，所谓的过度动画就是就是这两种样式之间转换的过度动画。在不使用 过渡时，两种样式之间的转变是瞬间完成的，看起来显得生硬，如果添加了过渡动画，那么就会显得柔软缓和，有一种Q弹的感觉。有时候还能构造复杂的动画，例如模拟重力，模拟水滴的形态，模拟风吹草低现牛羊的效果。

CSS 通过 `transition`  属性定义过度动画。主要有如下属性：

| 属性                         | 含义                        |
| ---------------------------- | --------------------------- |
| `transition-property`        | 指定使用过渡效果的 CSS 属性 |
| `transition-duration`        | 设置过渡动画持续时间        |
| `transition-timing-function` | 设置动画的时间函数          |
| `transition-delay`           | 设置动画的延迟时间          |

## transition-property

`transition-property `: 用于指定使用过渡效果的 CSS 属性，例如 `width, height`，默认值为 `all`，也就是只要有CSS属性发生变化，就会触发动画。当然也不是所有的CSS样式值都可以过渡，只有具有中间值的属性才具备过渡效果。常见的属性有：

颜色: color background-color opacity gradient
位置: position left right top bottom
长度: width height
盒子: margin padding outline border box-shadow z-index  
字体: font-size line-height text-indent vertical-align letter-spacing word-spacing font-weight text-shadow 
其他: transform visibility zoom clip

## transition-duration

动画持续时间，持续时间越长，动画看起来越慢。

## transition-timing-function

时间函数，主要有如下取值：

| 值                      | 描述                                                         |
| ----------------------- | ------------------------------------------------------------ |
| `linear`                | 动画从头到尾的速度是相同的。                                 |
| `ease`                  | **默认值**：动画以低速开始，然后加快，在结束前变慢。         |
| `ease-in`               | 动画以低速开始。                                             |
| `ease-out`              | 动画以低速结束。                                             |
| `ease-in-out`           | 动画以低速开始和结束。                                       |
| `cubic-bezier(n,n,n,n)` | 贝塞尔曲线（自定义数值），可到借助[可视化工具](https://cubic-bezier.com/#.17,.67,.83,.67)定制。 |

## transition-delay

表示动画延迟多少时间后触发动画，默认值 0，表示不延迟。

以上属性可以有如下简写形式：

```css
transition: width 1s ease ,heitht 1.5s linear 1s, color 2s ease-in-out ;
```

# 参考文章

- Ashun. [CSS3-transition过渡动画详解](https://juejin.cn/post/6970885478967050254). 稀土掘金.