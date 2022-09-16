# CSS 文字排版

CSS 与文字排版的样式非常多。

## 颜色

```html
    <p style="color:transparent;">Hello World</p>
    <p style="color:red;">Hello World</p>
    <p style="color:#05ffb0;">Hello World</p>
    <p style="color:rgb(50, 115, 220);">Hello World</p>
    <p style="color:rgba(0, 0, 0, 0.5);">Hello World</p>
    <p style="color:hsl(14, 100%, 53%);">Hello World</p>
    <p style="color:hsla(14, 100%, 53%, 0.6);">Hello World</p>
```

## 字体族

字体族即一类字体， 其值主要有6种：

- `serif`：衬线字体，衬线的笔画有粗有细的变化，在每一笔画上都自有风格，笔画末端会有修饰，强调艺术感，适合用于博客，旅游，文化，艺术类网站
- `sans-serif`: 无衬线字体 字体工整方正，给人正式的感觉，适合政务类，企业类网站使用。
- `monospace`: 等宽字体
- `cursive`: 手写字体
- `fantasy`: 奇幻字体
- `system-ui`: 系统UI字体

`font-family` 属性可以把多个字体名称作为一个“回退”系统来保存。如果浏览器不支持第一个字体，则会尝试下一个。也就是说，font-family 属性的值是用于某个元素的字体族名称或/及类族名称的一个优先表。浏览器会使用它可识别的第一个值。

```html
<h1>CSS 字体族</h1>
<p style="font-family: serif;">The quick brown fox jumps over the lazy dog. 前端靓仔 1234567890.</p>
<p style="font-family: sans-serif;">The quick brown fox jumps over the lazy dog. 前端靓仔 1234567890.</p>
<p style="font-family: monospace;">The quick brown fox jumps over the lazy dog. 前端靓仔 1234567890.</p>
<p style="font-family: cursive;">The quick brown fox jumps over the lazy dog. 前端靓仔 1234567890.</p>
<p style="font-family: fantasy;">The quick brown fox jumps over the lazy dog. 前端靓仔 1234567890.</p>
<p style="font-family: system-ui;">The quick brown fox jumps over the lazy dog. 前端靓仔 1234567890.</p>
```

常见用法：

```css
body, html {
    -webkit-font-smoothing: unset!important;
    font-family: -apple-system,system-ui,Segoe UI,Roboto,Ubuntu,Cantarell,Noto Sans,sans-serif,BlinkMacSystemFont,Helvetica Neue,PingFang SC,Hiragino Sans GB,Microsoft YaHei,Arial!important;
}
```

## 字体大小

 设置字体大小的 css 属性是：`font-size`

```css
font-size: 20px;
font-size: 1.2em;   /* 相对于父元素*/
font-size: 90%;
font-size: [smaller, larger];
font-size: 1.2rem;  /* 相对于 html*/
font-size: [xx-small,x-small,small,medium,large,x-large,x-large,xx-large];
```

## 斜体

```css
font-style: [normal,italic,oblique];
```

## 变形

```css
font-variant: [normal,small-caps];
```

## 字体粗细

```css
font-weight: normal;
font-weight: 600;
```
可以用值或对应的名字：
- **100** Thin
- **200** Extra Light
- **300** Light
- **400** Normal
- **500** Medium
- **600** Semi Bold
- **700** Bold
- **800** Extra Bold
- **900** Ultra Bold

## 字母间距, 单词间距

```css
letter-spacing: [normal, 2px, 0.1em];
word-spacing: 2em;
```

## 行高

```css
line-height: normal; // 取值方式类似
```

## 文字对齐

```css
text-align: [left, right, center, justify];
```

## 文字下划线

```css
text-decoration: [none,underline,line-through,overline];
text-decoration: underline dotted red;/* 设置线条样式与颜色 */
```

[MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/text-decoration)

## 首行缩进

```css
text-indent: [0, 40px, -2em];
```

## 文字阴影

```css
text-shadow: 2px 4px 10px red;
```

分别代表：向右偏移量，向下偏移量，模糊半径，颜色。

## 文字大小写

```css
text-transform: none;
text-transform: capitalize;
text-transform: uppercase;
text-transform: lowercase;
```

## 文字换行

```css
white-space : nowrap;
```

| 值       | 描述                                                         |
| -------- | ------------------------------------------------------------ |
| normal   | 默认。空白字符会被浏览器忽略。                               |
| pre      | 空白字符会被浏览器全部保留。其行为方式类似 HTML 中的 `<pre>` 标签。 |
| nowrap   | 文本不会换行，文本会在在同一行上继续，直到遇到 `<br>` 标签为止。 |
| pre-wrap | CSS 2.1新增保留空白符序列，但是正常地进行换行。              |
| pre-line | CSS 2.1新增合并空白符序列，但是保留换行符。                  |

## 文字溢出

```css
text-overflow: clip;
text-overflow: ellipsis;
overflow: hidden;
```

# 参考文章

- [Typography in CSS](https://cssreference.io/typography)

