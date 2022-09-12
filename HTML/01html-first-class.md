## HTML 第一课

### HTML的概念

**HTML** 全称为 HyperText Markup Language，译为超文本标记语言。HTML 不是一种通用的编程语言，是一种描述网页结构的**标记语言**。

**理解超文本**

所谓的超文本，有两层含义：

（1）网页除了文字，还能提供图片、音频、视频、动画、多媒体等内容，不止文本。

（2）不仅如此，它还提供超链接，可以从一个文件跳转到另一个文件，这样世界各地主机的文件都可以进行连接。即：超级链接文本。

**理解标记语言**

HTML 不是一种通用的编程语言（例如 C/C++，Java，Python），它是一种描述性的**标记语言**。标记语言是一套标记标签。比如：标签`<a>`表示超链接、标签`<img>`表示图片、标签`<h1>`表示一级标题等等，它们都是属于 HTML 标签。HTML 中的各种标签有各自的功能，属性，它们在网页中各司其职。HTML 标签是直接由浏览器解析执行，然后渲染网页内容。

## HTML的历史

接下来我们来看一下**HTML**从最原始到现在至今整个**HTML语言**的历史发展过程。

- **HTML 1.0**：在1993年6月作为互联网工程工作小组(IETF)工作草案发布,由此超文本标记语言第一版诞生。
- **HTML 2.0**：1995年11月作为 RFC 1866发布，于2000年6月发布之后被宣布已经过时。
- **HTML 3.2**：1997年1月14日，W3C推荐标准。
- **HTML 4.0**：1997年12月18日，W3C推荐标准。
- **HTML 4.01**（微小改进）：1999年12月24日，W3C推荐标准。
- **HTML 5**：HTML5是公认的下一代Web语言，极大地提升了Web在富媒体、富内容和富应用等方面的能力，被喻为终将改变移动互联网的重要推手。 2014年10月28日，W3C推荐标准。

**HTML5**的诞生，标记着互联网时代的深度发展，比如所**HTML5**里面诞生的**音频、视频、图像、动画**等都做了新的标准，它对于浏览器的兼容也是得到了一定的处理，由此可见，**HTML**的整个历史发展目前为止我们所使用的版本主要是99年诞生的**HTML 4.01**以及2014年诞生的 **HTML5**。

## HTML的专有名词

- 网页 ：由各种标记组成的一个页面就叫网页。
- 主页(首页) : 一个网站的起始页面或者导航页面。
- 标记：  比如`<p>`称为开始标记 ，`</p>`称为结束标记，也叫**标签(TAG)**。每个标签都规定好了特殊的含义。
- 元素：完整的标签，比如`<p>内容</p>`，又称为**元素(element)**。
- 属性：给每一个标签所做的辅助信息。
- HTTP：超文本传输协议。用来规定客户端浏览器和服务端交互时数据的一个格式，加密版本：HTTPS。

![grumpy-cat-small](img/grumpy-cat-small.png)

## 书写第一个 HTML 页面

我们打开 VS Code 编辑器，新建一个 HTML 文件 `hello-wrold.html`：

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
	
</body>
</html>
```

上面的内容，就是 html 页面的骨架。我们在此基础之上，新增几个标签，完整代码如下：

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>

<body>
    <h3>我爱前端，前端</h3>
    <a href="https://www.bilibili.com" target="_blank">
        走，跟我去B站搞学习
    </a>
</body>

</html>
```

标签写完之后，我们用 chrome 浏览器打开上面这个`hello-world.html`文件，看看页面效果：

![image](https://user-images.githubusercontent.com/27769596/189525556-c66c2d6e-a011-4b7d-be10-ad75eeb79d2a.png)

虽然页面比较简单，但是我们的第一个网页，是不是很有成就感？接下来我们将学习更多的标签，以此来构建更加复杂的网页。

## 小总结

- HTML 标记语言的含义
- HTML 简短历史
- VSCode写代码，live server

## 常见面试题

HTML5 新特性有哪些？

- 语义化更好的内容标签（header,nav,footer,aside,article,section）
- 表单增强：更多的表单类型，例如：color、email、url、search
- 音频、视频 API (audio,video)
- 画布(Canvas) API
- 地理(Geolocation) API
- 拖拽释放(Drag and drop) API
- 本地离线存储 localStorage 长期存储数据，浏览器关闭后数据不丢失；
- 新的技术 webworker, websocket, Geolocation

# 参考文章

- [初识HTML](https://github.com/qianguyihao/Web/blob/master/01-HTML/03-%E5%88%9D%E8%AF%86HTML.md)
- [MDN 文档](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started)
