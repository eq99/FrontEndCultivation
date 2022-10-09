

## `a` 标签详解

`a` 标签的作用是超链接引用，最重要的属性是 `href`:

```html
<a href="https://bilibili.com">网址直接跳转</a> 
<a href="images/1.webp">相对路径跳转</a>
<a href="hello.exe">如果是打不开的文件，则下载之</a>
<a href="">空链接是刷新</a>
<a href="#">#是回到顶部</a>
<a href="javascript:;">禁止链接跳转</a>
```

`a` 标签还有一个属性：`target`:

```html
<a href="https://bilibili.com" tatget="_blank">新标签打开网页</a> 
<a href="https://bilibili.com" tatget="_self">当前标签打开网页</a> 
```

### `table` 标签

表格相关的标签，[示例](https://www.w3schools.com/tags/tag_tbody.asp)：

```html
<table>
  <caption>Monthly savings</caption>
  <thead>
    <tr>
      <th>Month</th>
      <th>Savings</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>January</td>
      <td>$100</td>
    </tr>
    <tr>
      <td>February</td>
      <td>$80</td>
    </tr>
  </tbody>
  <tfoot>
    <tr>
      <td>Sum</td>
      <td>$180</td>
    </tr>
  </tfoot>
</table>
```

### `img` 图片

图片元素，[示例](https://www.w3schools.com/tags/tag_figure.asp)：

```html
<figure>
  <img src="https://imgs.lolimi.cn/?img=https://t2cy.com/d/file/phone/list/pic/2021-07-19/6913fa858c0dc48f9f7160ffb9430af5.jpg" alt="Trulli" style="width:100%" >
  <figcaption>Fig.1 碧蓝航线可畏礼服cos梳妆中的大小姐。</figcaption>
</figure>
```

(图片来源：https://acgdh.org/4409.html)

💯【面试题】img 的srcset 属性的作⽤？

响应式页面中经常用到根据屏幕密度设置不同的图片。这时就用到了 img 标签的srcset属性。srcset属性用于设置不同屏幕密度下，img 会自动加载不同的图片。用法如下：

```html
<img src="image-128.png" srcset="image-256.png 2x" />
```

使用上面的代码，就能实现在屏幕密度为1x的情况下加载image-128.png, 屏幕密度为2x时加载image-256.png。

按照上面的实现，不同的屏幕密度都要设置图片地址，目前的屏幕密度有1x,2x,3x,4x四种，如果每一个图片都设置4张图片，加载就会很慢。所以就有了新的srcset标准。代码如下：

```html
<img src="image-128.png"
     srcset="image-128.png 128w, image-256.png 256w, image-512.png 512w"
     sizes="(max-width: 360px) 340px, 128px" />
```

其中srcset指定图片的地址和对应的图片质量。sizes用来设置图片的尺寸零界点。对于 srcset 中的 w 单位，可以理解成图片质量。如果可视区域小于这个质量的值，就可以使用。浏览器会自动选择一个最小的可用图片。

sizes语法如下：

```html
sizes="[media query] [length], [media query] [length] ... "
```

sizes就是指默认显示128px, 如果视区宽度大于360px, 则显示340px。

【对比】图片与背景图片使用场景？

背景图片写在 css 里面，一般用在静态场景，例如装饰。

如果图片需要经常变动，则使用 `img`。

###  列表标签

列表标签分为有序标签（ol）和无序（ul），内部的条目放在 `li` (list item)标签中，如下图所示：

```html
<ol>
  <li>第一项</li>
  <li>第二项</li>
  <li>第三项</li>
</ol>

<ul>
  <li>第一项</li>
  <li>第二项</li>
  <li>第三项</li>
</ul>
```

他们的区别是：ol 的条目之前有数字，而 ul 的条目之前是小圆点。

除此之外还有自定义列表 `dl`:

```html
<dl>
  <dt>列表头</dt>
  <dd>列表项</dd>
  <dd>列表项</dd>
  <dd>列表项</dd>
  <dd>列表项</dd>
</dl>
```

【属性】列表相关的属性有：

- 使用 `type` 指定序号类型：

```html
<ol type="i">
  <li>Introduction</li>
  <li>List of Grievances</li>
  <li>Conclusion</li>
</ol>

<!-- 
a: 小写字母序
A: 大写字母序
i: 小写罗马字母
I: 大写罗马字母
-->
```



- 可以用 `start` 属性设置开始序号，使用 `reversed` 开启倒排序 ：

```html
<ol start="5" reversed>
  <li>Toast pita, leave to cool, then slice down the edge.</li>
  <li>
    Fry the halloumi in a shallow, non-stick pan, until browned on both sides.
  </li>
  <li>Wash and chop the salad.</li>
  <li>Fill pita with salad, hummus, and fried halloumi.</li>
</ol>
```

- 使用 `value` 个列表项标号：

```html
<ol>
  <li value="2">Toast pita, leave to cool, then slice down the edge.</li>
  <li value="4">
    Fry the halloumi in a shallow, non-stick pan, until browned on both sides.
  </li>
  <li value="6">Wash and chop the salad.</li>
  <li value="8">Fill pita with salad, hummus, and fried halloumi.</li>
</ol>
```



列表相关的样式属性：

```css
ol, ul {
  list-style-type: square|disc|circle|decimal; /*可以用来控制列表项开始样式*/
  list-style-image: url(example.png); /*在列表项开头使用用图片*/
  list-style-position: inside|outside ; /* 里面或外面 */
}

/*简写形式*/
ul {
  list-style: square url(example.png) inside;
}
```



### `details` 标签

详情展示：[示例](https://www.w3schools.com/tags/tag_details.asp)

```html
<details>
  <summary>Epcot Center</summary>
  <p>Epcot is a theme park at Walt Disney World Resort featuring exciting attractions, international pavilions, award-winning fireworks and seasonal special events.</p>
</details>
```

## 带有特定样式的行内元素

```html
<p>My favorite color is <del>blue</del> <ins>red</ins>!<strong>This text is important!</strong></p>

<p>This is normal text - <b>and this is bold text</b>.</p>
<p><small>This is some smaller text.</small></p>

<p>Press <kbd>Ctrl</kbd> + <kbd>C</kbd> to copy text (Windows).</p>

<p><cite>The Scream</cite> by Edward Munch. Painted in 1893.</p>

The <abbr title="World Health Organization">WHO</abbr> was founded in 1948.

<p>你可以用 <code>console.log()</code> 打印消息.</p>

<p>Open from <time>10:00</time> to <time>21:00</time> every weekday.</p>

<p>I have a date on <time datetime="2008-02-14 20:00">Valentines day</time>.</p>

<p><b>Note:</b> The time element does not render as anything special in any of the major browsers.</p>
```

# 语义化布局标签

什么是语义化标签呢？请看如下两张图（来自[MDN](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Document_and_website_structure)）：

网站布局草图：

![site-structure](img/site-structure.png)

HTML布局图：

![sample-website](img/sample-website.png)

网页布局主要涉及到如下一些语义化的标签：

```html
<header></header>  头部
<nav></nav>  导航栏
<section></section>  区块（有语义化的div）
<main></main>  主要区域
<article></article>  主要内容
<aside></aside>  侧边栏
<footer></footer>  底部
```

💯【面试题】请说说对 HTML5 语义化的理解？

**语义化是指根据内容的结构化（内容语义化），选择合适的标签（代码语义化）**。通俗来讲就是用正确的标签做正确的事情。

语义化的优点如下：

- 对机器友好，带有语义的文字表现力丰富，更适合搜索引擎的爬虫爬取有效信息，有利于SEO。除此之外，语义类还支持读屏软件，根据文章可以自动生成目录；
- 对开发者友好，使用语义类标签增强了可读性，结构更加清晰，开发者能清晰的看出网页的结构，便于团队的开发与维护。

💯【面试题】`src` 和 `href` 的区别？

`src` 和 `href` 都是**用来引用外部的资源**，它们的区别如下：

- **src：** 表示对资源的引用，它指向的内容会嵌入到当前标签所在的位置。src会将其指向的资源下载并应⽤到⽂档内，如请求js脚本。当浏览器解析到该元素时，会暂停其他资源的下载和处理，直到将该资源加载、编译、执⾏完毕，所以⼀般js脚本会放在页面底部。
- **href：** 表示超文本引用，它指向一些网络资源，建立和当前元素或本文档的链接关系。当浏览器识别到它他指向的⽂件时，就会并⾏下载资源，不会停⽌对当前⽂档的处理。 常用在 a、link 等标签上。



## video

```
<video width="658" height="444" poster="http://www.youname.com/images/first.png" autoplay="autoplay" preload="none" controls="controls">
  <source src="http://www.youname.com/images/first.ogv" />
  <source src="http://www.youname.com/images/first.ogg" />
</video>
```

这段代码在页面中定义了一个视频，此视频的预览图为poster的属性值，显示浏览器的默认媒体控制栏，预加载视频的元数据，循环播放，宽度为900像素，高度为240像素。

第一选择视频地址为第一个source标签的src属性值，视频类别为Ogg视频，视频编码译码器为Theora，音频编码译码器为Vorbis，播放媒 介为显示器；后面的依次类推。

## audio

```html
<audio controls>
  <source src="/static/i/html/horse.ogg" type="audio/ogg">
  <source src="/static/i/html/horse.mp3" type="audio/mpeg">
  您的浏览器不支持 audio 元素
</audio>
```

## track[^2]

`<track>` 标签为媒体元素（比如 `<audio>` and `<video>`）规定外部文本轨道，也就是字幕，字幕格式采用 WebVTT 格式（.vtt 格式文件）。这个元素用于规定字幕文件或其他包含文本的文件，当媒体播放时，这些文件是可见的。使用示例：

```html
<video controls src="https://interactive-examples.mdn.mozilla.net/media/cc0-videos/friday.mp4">
  <track default kind="captions" srclang="en" src="https://interactive-examples.mdn.mozilla.net/media/examples/friday.vtt">
</video>
```

# 参考文章

1. [「2021」高频前端面试题汇总之HTML篇](https://juejin.cn/post/6905294475539513352)

[^2]:[The Embed Text Track element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track). MDN.

