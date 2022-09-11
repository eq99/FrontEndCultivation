## HTML 第一课

### HTML的概念

**HTML** 全称为 HyperText Markup Language，译为超文本标记语言。HTML 不是一种通用的编程语言，是一种描述性的**标记语言**。

**理解超文本**

所谓的超文本，有两层含义：

（1）图片、音频、视频、动画、多媒体等内容，成为超文本，因为它们超出了文本的限制。

（2）不仅如此，它还可以从一个文件跳转到另一个文件，与世界各地主机的文件进行连接。即：超级链接文本。

**理解标记语言**

HTML 不是一种通用的编程语言（例如 C/C++，Java，Python），它是一种描述性的**标记语言**。主要有两层含义：

（1）**标记语言是一套标记标签**。比如：标签`<a>`表示超链接、标签`<img>`表示图片、标签`<h1>`表示一级标题等等，它们都是属于 HTML 标签。

说的通俗一点就是：网页是由网页元素组成的，这些元素是由 HTML 标签描述出来，然后通过浏览器解析，就可以显示丰富的网站内容了。

（2）HTML 标签是直接由浏览器解析执行。

### HTML是负责描述文档结构的语言

HTML 格式的文件是一个纯本文文件（就是用txt文件改名而成），用一些标签来描述语义，这些标签在浏览器页面上是无法直观看到的，所以称之为“超文本标记语言”。

接下来，我们需要学习 HTML 中的各种标签，他们有各自的功能，属性，它们在网页中各司其职。

## HTML的历史

![html中标签发展趋势](http://img.smyhvae.com/20151001_1001.png)

其中，我们专门来对XHTML做一个介绍。

**XHTML介绍：**
XHTML：Extensible Hypertext Markup Language，可扩展超文本标注语言。
XHTML的主要目的是为了取代HTML，也可以理解为HTML的升级版。
HTML的老版本标记书写没有统一的规范，可移植性差，同一套代码可能在某些平台上 ipad、手机、电视无法正常显示。
XHTML与HTML4.0的标记基本上一样，但 XHTML 是更严格的、纯净的 HTML。

## HTML的专有名词

- 网页 ：由各种标记组成的一个页面就叫网页。
- 主页(首页) : 一个网站的起始页面或者导航页面。
- 标记：  比如`<p>`称为开始标记 ，`</p>`称为结束标记，也叫**标签(TAG)**。每个标签都规定好了特殊的含义。
- 元素：完整的标签，比如`<p>内容</p>`，又称为**元素(element)**。
- 属性：给每一个标签所做的辅助信息。
- HTTP：超文本传输协议。用来规定客户端浏览器和服务端交互时数据的一个格式，加密版本：HTTPS。

## 书写第一个 HTML 页面

我们打开 VS Code 编辑器，新建一个 HTML 文件 `hello-wrold.html`：

```html
<!DOCTYPE html>
<html lang="en">
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
    <a href="https://www.bilibili.com" target="__blank">
        走，跟我去B站搞学习
    </a>
</body>

</html>
```

标签写完之后，我们用 chrome 浏览器打开上面这个`hello-world.html`文件，看看页面效果：

![image](https://user-images.githubusercontent.com/27769596/189525556-c66c2d6e-a011-4b7d-be10-ad75eeb79d2a.png)


虽然页面比较简单，但是我们的第一个网页，是不是很有成就感？接下来我们将学习更多的标签，以此来构建更加复杂的网页。

# 参考文章

- [初识HTML](https://github.com/qianguyihao/Web/blob/master/01-HTML/03-%E5%88%9D%E8%AF%86HTML.md)
