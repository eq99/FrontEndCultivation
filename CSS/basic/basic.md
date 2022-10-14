# CSS 基础知识

本章主要讲解 CSS 的一些基础知识，理解这些基础知识能让我们写出更好的CSS。

## 样式的引入

CSS 有三种引入方法：分别是行内样式，内部样式表，外部样式表

### 行内样式

行内样式也叫内联式样式，行内样式通过标签的 `style` 属性设置样式，如下所示：

```css
<h1 style="color:red;">style 属性的应用</h1>
<p style="font-size:14px;color:green;">直接在 HTML 标签中设置的样式</p>
<p style="display:none;">影藏的内容</p>
```

行内样式的优先级较高，可以用来设置某些特殊样式或者微调元素的样式。不建议把所有的样式放在行内样式表中，因为过多的样式会让元素变得“臃肿不堪”，使文档的文档结构不明晰，不利于维护。

### 内部样式表

内部样式表是把样式写在一个 `style` 标签内部：

```css
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>内部样式表</title>
  <style>
    h1 {
      color: red;
    }

    p {
      font-size: 14px;
      color: green;
    }
  </style>
</head>

<body>
  <h1>style 属性的应用</h1>
  <p>直接在 HTML 标签中设置的样式</p>
</body>

</html>
```

内部样式表的方式实现了结构与样式的分离，这样做的好处是使得 HTML 专注于搭建网页结构，而CSS更加专注于元素样式的调整。这种方式在学习或演示CSS属性的时候非常有用，或者调试一个小功能的时候也非常方便。目前有一些前端框架就是采用的这种方式，例如 Vue，Svelte。这种写法在应对大型网站的时候不够友好，因为大型网站有很多页面，有许多CSS样式可以共用，如果每个页面都要重写一遍或者手动赋值，那种感觉也是挺痛苦的。因此就有了独立的CSS文件。

### 外部样式表

 外部样式表指的是把样式写在一个单独的样式文件中，然后通过文件路径引用：

```html
<link rel="stylesheet" href="reset.css">
<link rel="stylesheet" href="utils.css">
<link rel="stylesheet" href="main.css">
```

把样式单独的写在样式文件中，可以在需要的地方通过 `link` 标签引入，修改的时候，只需要修改一次，所有依赖该文件的网页都能生效。除此之外，还可以把自己的样式表以库的形式发布到网上，其他有需要的网站也可以引用，真的不要太方便。

## CSS 选择器

把样式写在样式表中（`style` 标签或者单独的样式文件中）需要解决一个问题：这个样式是写给谁的，或者谁需要这个样式，如果把样式必做各式各样的衣服，那么给元素应用样式的过程就像穿衣服的一样。穿衣服的过程就是应用样式的过程，CSS 通过 CSS 选择器给元素应用样式。

### 标签选择器

标签选择器通过元素标签名字选择标签，例如：

```css
h1 {
  color: red;
}
```

会把所有的 h1 标签的字体颜色设置为红色。

### id 选择器

id 选择器会选择具有某个 id 值得标签：

```html
<style>
  #tom {
    color: grey;
  }

  #garfield {
    color: orange;
  }
</style>

<h1 id="tom">我是汤姆猫</h1>
<h1 id="garfield">我是加菲猫</h1>
```

id 选择器的语法格式是 `#id-name`，一般来说 id 在一个网页中是唯一的，如果有两个元素有一样的 id 值会发生什么呢？那就是所有具有该 id 值的元素都会穿上相应样式。

### 类选择器

类选择器通过 `class` 属性应用相应的样式，它的声明方法是：

```css
<style>
  .grey {
    color: grey;
  }

  .orange {
    color: orange;
  }
</style>

<h1 class="grey">我是汤姆猫</h1>
<h1 class="orange">我是加菲猫</h1>
```

类选择通过 `.class-name` 声明类样式，类样式可以组合：

```css
<style>
  .grey {
    color: grey;
  }

  .orange {
    color: orange;
  }

  .underline {
    text-decoration: underline;
  }
  
  .small{
    font-size: 14px;
  }
</style>

<h1 class="grey underline">我是汤姆猫</h1>
<h1 class="orange small">我是加菲猫</h1>
```

每种样式有自己独特的效果，通过组合不同的样式来实现一个复杂的效果，就好像女生穿搭一样：涂个口红显得成熟，配个项链显得富贵，配个蝴蝶结显得可爱，配个丝袜显得性感，想要什么样的效果就搭配什么样式，所以女孩子特别喜欢学前端，对吧。

### 属性选择器

属性选择器可以选择具有某种属性的元素，语法如下所示：

```css
<style>
  [hello] {
    color: green;
  }
</style>

<h1 hello>我是汤姆猫</h1>
<h1>我是加菲猫</h1>
```

当然还可以设置属性值：

```css
<style>
  [hello=world] {
    color: green;
  }
</style>

<h1 hello="world">我是汤姆猫</h1>
<h1 hello>我是加菲猫</h1>
```

### 后代选择器

除了单独的选择某个元素，还可以通过元素的之间的嵌套关系选择，最常用的是后代选择器：

```css
<style>
  .summer .plant {
    color: green;
  }

  .autumn .plant {
    color: gold;
  }
</style>

<div class="summer">
  <div class="plant">玉米</div>
</div>

<div class="autumn">
  <div class="plant">玉米</div>
</div>
```

两个独立的选择器之间加一个空格表示：什么里面的什么，或者说，谁谁谁的后代的后代。反正空格表示后代关系，隔着也是可以的，只要在里面就能生效。上面里的例子表示夏天里的植物和秋天里的植物各是一个样式。

### 直接孩子选择器

直接孩子选择器的语法是 `parent>child`:

```html
<style>
  .summer>.plant {
    color: green;
  }

  .autumn>.plant {
    color: gold;
  }
</style>

<div class="summer">
  <div class="north">
    <div class="plant">玉米</div>
  </div>

</div>

<div class="autumn">
  <div class="plant">玉米</div>
</div>
```

隔代了就不管用了，必须是直接的孩子结点。

### 兄弟关系

相邻兄弟选择器 `one+two+three`：

```html
<style>
  .one+.two {
    color: red;
  }

  .two+.three {
    color: green;
  }

  .one+.three {
    color: blue;
  }
</style>

<div class="box">
  <div class="one">大娃</div>
  <div class="two">二娃</div>
  <div class="three">三娃</div>
</div>
```

特别注意大娃和三娃隔了个二娃，三娃的蓝色就不生效了。

### 小结

以上小节总结了一些常用的选择器，CSS 选择器实在是太多了，更多的选择器可以阅读参考文章[^1]。

## CSS选择器优先级

CSS （Cascading Style Sheets）叠层样式表，某一个元素可能同时设置了很多样式，甚至有些样式是冲突的，到底用哪个样式呢？这时候浏览器就需要根据优先级来设置样式了。打开浏览器调试工具，在 Styles 一栏会显示元素的样式，并且优先级高的样式会显示在上面，优先级较低的样式会显示在下面，并且后面有很多相同的样式样式属性被划掉了，这是因为高优先级的样式覆盖了低优先级样式的缘故。

![image-20221013223318715](image-20221013223318715.png)

样式选择器优先级归纳：

| **选择器**       | **格式**                | **优先级权重** |
| ---------------- | ----------------------- | -------------- |
| id选择器         | `#id`                   | 100            |
| 类选择器         | `.classname`            | 10             |
| 属性选择器       | `div[attr=value]`       | 10             |
| 伪类选择器       | `li:last-child`         | 10             |
| 标签选择器       | `element`               | 1              |
| 伪元素选择器     | `li::after`             | 1              |
| 相邻兄弟选择器   | `h1+p`                  | 0              |
| 相邻兄弟们选择器 | `h1~p`                  | 0              |
| 子选择器         | `ul>li`                 | 0              |
| 后代选择器       | `li a`                  | 0              |
| 通配符选择器     | `*`                     | 0              |
| 自带的样式       | 浏览器自带              | 1              |
| 行内样式         | `style="color: red;"`   | 1000           |
| `!important`     | `color: red;!important` | 1000           |
| 继承             | 继承                    | 0              |

选择器的特殊性由选择器本身的组成确定，特殊性值表述为4个部分，如0,0,0,0
一个选择器的具体特殊性如下确定：
		内联声明的特殊性，加1,0,0,0
		对于选择器中给定的ID属性值，加0,1,0,0
		对于选择器中给定的各个类属性，属性选择，或伪类，加0,0,1,0
		对于选择器中的给定的各个元素和伪元素，加0,0,0,1
		通配符选择器的特殊性为0,0,0,0
		结合符(,)对选择器的特殊性没有一点贡献
		继承没有特殊性
	
	例：div[id="test"] (0,0,1,1) 和 #test(0,1,0,0)
	特殊性不进位，1,0,0,0大于所有以0开头的特殊性(比如0,1111,0,0)
	如果多个规则与同一个元素匹配，而有些声明互相冲突时，特殊性越大的越占优势
-----------------------------------
©著作权归作者所有：来自51CTO博客作者叹之的原创作品，请联系作者获取转载授权，否则将追究法律责任
CSS优先级判定规则
https://blog.51cto.com/u_15295346/3022331





对于选择器的**优先级**：

- 内联样式：1000

**注意事项：**

- `!important` 声明的样式的优先级最高；
- 如果优先级相同，则最后出现的样式生效；
- 继承得到的样式的优先级最低；
- 通用选择器（*）、子选择器（>）和相邻同胞选择器（+）并不在这四个等级中，所以它们的权值都为 0 ；
- 样式表的来源不同时，优先级顺序为：内联样式 > 内部样式 > 外部样式 > 浏览器用户自定义样式 > 浏览器默认样式。

### CSS 样式继承性

**一、无继承性的属性**

1) **display**：规定元素应该生成的框的类型
2) **文本属性**：

- vertical-align：垂直文本对齐
- text-decoration：规定添加到文本的装饰
- text-shadow：文本阴影效果
- white-space：空白符的处理
- unicode-bidi：设置文本的方向

3) **盒子模型的属性**：width、height、margin、border、padding
4) **背景属性**：background、background-color、background-image、background-repeat、background-position、background-attachment
5) **定位属性**：float、clear、position、top、right、bottom、left、min-width、min-height、max-width、max-height、overflow、clip、z-index
6) **生成内容属性**：content、counter-reset、counter-increment
7) **轮廓样式属性**：outline-style、outline-width、outline-color、outline
8) **页面样式属性**：size、page-break-before、page-break-after
9) **声音样式属性**：pause-before、pause-after、pause、cue-before、cue-after、cue、play-during

**二、有继承性的属性**

1. **字体系列属性**

- font-family：字体系列
- font-weight：字体的粗细
- font-size：字体的大小
- font-style：字体的风格

2. **文本系列属性**

- text-indent：文本缩进
- text-align：文本水平对齐
- line-height：行高
- word-spacing：单词之间的间距
- letter-spacing：中文或者字母之间的间距
- text-transform：控制文本大小写（就是uppercase、lowercase、capitalize这三个）
- color：文本颜色

3. **元素可见性**

- visibility：控制元素显示隐藏

4. **列表布局属性**

- list-style：列表风格，包括list-style-type、list-style-image等

5. **光标属性**

- cursor：光标显示为何种形态

###  display 的属性值及其作用

| **属性值**   | **作用**                                                   |
| ------------ | ---------------------------------------------------------- |
| none         | 元素不显示，并且会从文档流中移除。                         |
| block        | 块类型。默认宽度为父元素宽度，可设置宽高，换行显示。       |
| inline       | 行内元素类型。默认宽度为内容宽度，不可设置宽高，同行显示。 |
| inline-block | 默认宽度为内容宽度，可以设置宽高，同行显示。               |
| list-item    | 像块类型元素一样显示，并添加样式列表标记。                 |
| table        | 此元素会作为块级表格来显示。                               |
| inherit      | 规定应该从父元素继承display属性的值。                      |

### display 的 block、inline和inline-block的区别

（1）**block：** 会独占一行，多个元素会另起一行，可以设置width、height、margin和padding属性；

（2）**inline：** 元素不会独占一行，设置width、height属性无效。但可以设置水平方向的margin和padding属性，不能设置垂直方向的padding和margin；

（3）**inline-block：** 将对象设置为inline对象，但对象的内容作为 对象呈现，之后的内联对象会被排列在同一行内。

对于行内元素和块级元素，其特点如下：

**（1）行内元素**

- 设置宽高无效；
- 可以设置水平方向的margin和padding属性，不能设置垂直方向的padding和margin；
- 不会自动换行；

**（2）块级元素**

- 可以设置宽高；
- 设置margin和padding都有效；
- 自动换行；
- 多个块状，默认排列从上到下。

（3）inline-block

- 可以设置宽高；
- 设置margin和padding都有效；
- 不会自动换行，后面可以有其他元素。

[案例](https://www.w3school.com.cn/css/css_inline-block.asp)

### 隐藏元素的方法有哪些

- **display: none**：渲染树不会包含该渲染对象，因此该元素不会在页面中占据位置，也不会响应绑定的监听事件。
- **visibility: hidden**：元素在页面中仍占据空间，但是不会响应绑定的监听事件。
- **opacity: 0**：将元素的透明度设置为 0，以此来实现元素的隐藏。元素在页面中仍然占据空间，并且能够响应元素绑定的监听事件。
- **position: absolute**：通过使用绝对定位将元素移除可视区域内，以此来实现元素的隐藏。
- **z-index: 负值**：来使其他元素遮盖住该元素，以此来实现隐藏。
- **clip/clip-path** ：使用元素裁剪的方法来实现元素的隐藏，这种方法下，元素仍在页面中占据位置，但是不会响应绑定的监听事件。
- **transform: scale(0,0)**：将元素缩放为 0，来实现元素的隐藏。这种方法下，元素仍在页面中占据位置，但是不会响应绑定的监听事件。

### transition和animation的区别

- **transition是过度属性**，强调过度，它的实现需要触发一个事件（比如鼠标移动上去，焦点，点击等）才执行动画。它类似于flash的补间动画，设置一个开始关键帧，一个结束关键帧。
- **animation是动画属性**，它的实现不需要触发事件，设定好时间之后可以自己执行，且可以循环一个动画。它也类似于flash的补间动画，但是它可以设置多个关键帧（用@keyframe定义）完成动画。

### 对盒模型的理解

CSS3中的盒模型有以下两种：标准盒子模型、怪异盒子模型（IE盒子模型）。

 ![盒子模型](img/box-size.jpg)  

盒模型都是由四个部分组成的，分别是 margin、border、padding 和 content。

标准盒模型和怪异盒模型的区别在于设置 width 和 height 时，所对应的范围不同：

- 标准盒模型的 width 和 height 属性的范围只包含了content，
- 怪异盒模型的 width 和 height 属性的范围包含了border、padding和content。

可以通过修改元素的 box-sizing 属性来改变元素的盒模型：

```css
box-sizing: border-box;  /*标准盒子模块*/
box-sizing: content-box; /*怪异*/
```

作者：CUGGZ
链接：https://juejin.cn/post/6905539198107942919
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

## 参考文章

[^1]: MDN. [CSS selectors](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Selectors). 
