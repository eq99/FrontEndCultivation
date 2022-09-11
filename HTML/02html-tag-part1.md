# HTML 排版标签

排版标签指的是划分文章结构的一类标签。

## 标题

用于标注**标题**的标签有6个，[示例](https://www.w3schools.com/tags/tag_hn.asp)：

```html
<h1>This is heading 1</h1>
<h2>This is heading 2</h2>
<h3>This is heading 3</h3>
<h4>This is heading 4</h4>
<h5>This is heading 5</h5>
<h6>This is heading 6</h6>
```

## 内容

用于标记**段落**的标签，[示例](https://www.w3schools.com/tags/tag_p.asp)：

```html
<p>This is a paragraph.</p>
```

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

列表，分为两类，有序列表和无序列表：

先看有序列表，[示例](https://www.w3schools.com/tags/tag_ol.asp)：

```html
<ol>
  <li>Coffee</li>
  <li>Tea</li>
  <li>Milk</li>
</ol>
```

再看无序列表，[示例](https://www.w3schools.com/tags/tag_ul.asp)：

```html
<ul>
  <li>Coffee</li>
  <li>Tea</li>
  <li>Milk</li>
</ul>
```

图片元素，[示例](https://www.w3schools.com/tags/tag_figure.asp)：

```html
<figure>
  <img src="https://imgs.lolimi.cn/?img=https://t2cy.com/d/file/phone/list/pic/2021-07-19/6913fa858c0dc48f9f7160ffb9430af5.jpg" alt="Trulli" style="width:100%">
  <figcaption>Fig.1 碧蓝航线可畏礼服cos梳妆中的大小姐。</figcaption>
</figure>
```

图片来源：https://acgdh.org/4409.html

详情展示：[示例](https://www.w3schools.com/tags/tag_details.asp)：

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

