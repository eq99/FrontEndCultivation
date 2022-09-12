# HTML 标签之功能标签

功能标签指的是不提供肉眼可见的内容的标签，这类标签对构建网页起到辅助作用，但作用不可忽略。

## `script` 标签

`script` 标签是最重要的功能性标签之一，它用于引入 `JavaScript` 脚本。

主要有两种用法：

**第一种，通过 `src` 属性引入** ：

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
    <script src="./hello-world.js"></script>
</body>
</html>
```

然后在同一文件夹下建立一个文件 `hello-world.js`

```js
alert("hello");
```

浏览器打开即可看到弹出框。

这种引入方法有两个重要属性：

`defer`：单独加载，不影响HTML解析，在HTML解析完之后运行，会按相对顺序执行。在 DOMContentLoaded

事件之前完成。

`async`：单独加载，单独运行（特立独行），两者关系如下图所示：

![defer-async](img/defer-async.jpg)

![script-load](img/script-load.jpg)

第二种，在标签内部直接写 JS 代码：

```html
    <script>
        let msg = "哎哟，你干嘛！";
        alert(msg);
    </script>
```

## `meta` 标签

`meta` 标签放在 `head` 标签里面，用于设置网页的源信息。这些源信息一方面可以控制网页行为，另一方面可以做SEO，或者分享的网页格式。

```html
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
```

`meta` 标签共有两个重要属性：

**http-equiv 属性**

`meta` 标签的 `http-equiv` 属性语法格式是： `＜meta http-equiv='参数' content='参数变量值'＞`，请看如下案例：


1. `<meta http-equiv='Set-Cookie' content='cookievalue=xxx; expires=Friday,12-Jan-2001 18:18:18 GMT; path=/'>`:如果网页过期，那么存盘的cookie将被删除。必须使用GMT的时间格式。
2. `<meta http-equiv='expires' content='时间' >`：用于设定网页的到期时间。一旦网页过期，必须到服务器上重新传输。
3. `<meta http-equiv='Refresh' content='3; https://www.bilibili.com'>`：告诉浏览器在【数字】秒后跳转到【一个网址】
4. `<meta http-equiv='content-Type' content='text/html; charset=utf-8'>`：设定页面使用的字符集。
    `<meta charset='utf-8'>`：在HTML5中设定字符集的简写写法。
5. `<meta http-equiv='Pragma' content='no-cache'>`：禁止浏览器从本地计算机的缓存中访问页面内容。访问者将无法脱机浏览。
6. `<meta http-equiv='Window-target' content='_top'>`：用来防止别人在iframe(框架)里调用自己的页面，这也算是一个非常实用的属性。
7. `<meta http-equiv='X-UA-Compatible' content='IE=edge,chrome=1'>` :强制浏览器按照特定的版本标准进行渲染。但不支持IE7及以下版本。如果是ie浏览器就用最新的ie渲染，如果是双核浏览器就用chrome内核。

**name属性**

`meta` 标签的 `name` 属性语法格式是： `＜meta name='参数' content='具体的参数值'＞`，常见的案例如下：
1. `<meta name='viewport' content='width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no'>`：在移动设备浏览器上，禁用缩放（zooming）功能，用户只能滚动屏幕。
2. `<meta name='description' content="">`：告诉搜索引擎，当前页面的主要内容是xxx。
3. `<meta name='keywords' content="">`：告诉搜索引擎，当前页面的关键字。
4. `<meta name='author' content="">`：告诉搜索引擎，标注网站作者是谁。
5. `<meta name='copyright' content="">`：标注网站的版权信息。



## `link` 标签

`link` 标签的作用就是定义当前文档与链接的关系，语法格式：`<link rel="link类型" href="链接">`

常见案例：
1. `<link rel="stylesheet" href="//s1.hdslb.com/bfs/static/jinkela/long/laputa-css/map.css">`：引入样式表。
2. `<link rel="dns-prefetch" href="//s1.hdslb.com">`: DNS Prefetch 是一种 DNS 预解析技术。当你浏览网页时，浏览器会在加载网页时对网页中的域名进行解析缓存，这样在你单击当前网页中的连接时就无需进行 DNS 的解析，减少用户等待时间，提高用户体验。
3. `<link rel="shortcut icon" href="https://www.bilibili.com/favicon.ico?v=1">`: 引入图标。
4. `<link rel="apple-touch-icon" href="//static.hdslb.com/mobile/img/512.png">`：现今移动设备越来越多，苹果为iOS设备配备了apple-touch-icon私有属性，添加该属性，在iPhone,iPad,iTouch的safari浏览器上可以使用添加到主屏按钮将网站添加到主屏幕上，方便用户以后访问。
5. `<link rel="canonical" href="https://www.bilibili.com/">`：规范化网址，提高网站搜索引擎排名。