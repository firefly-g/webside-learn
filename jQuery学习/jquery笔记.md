1.jQuery入口函数

- jquery的入口函数不必等待所有外部的资源加载完毕，只需DOM结构渲染完毕即可执行内部代码，相当于js中的DOMContentLoaded。与js中的load事件不同，load是等页面文档，js文件、css文件。图片加载完毕才执行内部代码。

- 第一种写法：
  		$(document).ready(function(){
               ...
  		})

- 第二种写法：

  ​       $(function(){

  ​		   ...

  ​      })

2.jQuery对象和DOM对象

- DOM对象：用原生js获取过来的对象      var myDiv=document.querySelector('div');

- jQuery对象：用jQuery获取过来的元素对象       $('div');

- jQuery对象的本质是：利用$对DOM对象包装后产生的对象（伪数组形式存储）

- 区别：DOM对象只能使用原生的js属性和方法，jQuery对象只能使用jQuery方法

- 注意：DOM对象与jQuery对象是可以相互转换的

  > 原生js中含有的一些属性和方法在jQuery中未封装，要想使用时需要把jQuery对象转化为DOM对象。

- 转换方法：

  > jquery转为DOM对象：
  >
  > （1）$('div')[index]           index是索引号
  >
  > （2）$('div').get(index)     index是索引号
  >
  > 
  >
  > DOM对象转jquery对象：
  >
  >    （1） $(DOM对象)



3.jQuery常用API

- 选择器： $('选择器')    

  > $('#id')   		 id选择器
  >
  > $('*')     		 匹配所有元素
  >
  > $('.class')		同一类class元素
  >
  > $('div')			同一类标签的所有元素
  >
  > $('div ,p,li')	 选取多个元素   并集选择器
  >
  > $('li.current')  交集选择器

- 样式：

- 效果：

- 属性操作：

- 文本属性值：

- 元素操作：

- 尺寸、位置操作：