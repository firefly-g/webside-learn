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

  > 1.普通选择器
  >
  > > $('#id')   		 id选择器
  > >
  > > $('*')     		 匹配所有元素
  > >
  > > $('.class')		同一类class元素
  > >
  > > $('div')			同一类标签的所有元素
  > >
  > > $('div ,p,li')	 选取多个元素   并集选择器
  > >
  > > $('li.current')  交集选择器
  >
  > 
  >
  > 2.层级选择器：
  >
  > > 子代选择器   $('ul>li')     只获取儿子层级的元素
  > >
  > > 后代选择器   $('ul li')	  获取ul下的所有li元素，包括孙子等
  >
  > 
  >
  > * 隐式迭代(重要)：遍历内部DOM元素（伪数组形式存储）的过程就叫做隐式迭代。
  > * 即循环遍历匹配成功的元素，执行相应的方法，简化操作。
  >
  > 2.筛选选择器：
  >
  > > 语法                                         用法                                描述
  > >
  > > :first							  $("li:first")							获取第一个li
  > >
  > > :last							   $("li:last")						    获取最后一个li
  > >
  > > :eq(index)					$("li :eq(2)")			 获取到的li元素中，选择索引号为2的元素（index从0开始）
  > >
  > > :odd							  $('li :odd')				获取到的li元素中，选择索引号为奇数的元素
  > >
  > > :even						 $('li:even')				   获取到的li元素中，选择索引号为偶数的元素
  > >
  > > > 2.1筛选方法：
  > > >
  > > > parent() 				$('li').parent()			查找父级元素           ******
  > > >
  > > > children(' ')	       $('ul').children('li')	相当于$('ul>li') 子代选择器    ******
  > > >
  > > > find(' ')				   $("ul").find('li')		  相当于$('ul  li')  后代选择器  ******
  > > >
  > > > siblings(' ')			 $('.first').siblings("li")   查找兄弟节点，不包括自己   ******
  > > >
  > > > nextAll([expr])			  $('.first').nextAll('li')	 查找当前元素之后所有的同辈元素
  > > >
  > > > prevAll([expr])			 $('.first').prevAll('li')      查找当前元素之前所有的同辈元素
  > > >
  > > > hasClass(class)		$('div').hasClass('protected')  检查当前的元素是否含有某个特定的类
  > > >
  > > > eq(index)				$('li:eq(2)')									相当于$('li:eq(2)')       ******                                               
  > >
  > > 

- 样式：利用css属性操作元素样式，也可以操作类 修改多个样式

  > 1.参数只写属性名时，返回的是属性值
  >
  > $(this).css('color');
  >
  > 2.css参数格式：
  >
  > $(this).css('color','red');
  >
  > 3.参数可以是对象形式，方便设置多组样式。属性名和属性值用逗号隔开，属性可以不用加引号
  >
  > $('div').css({
  >
  > width:400,
  >
  > height:400，
  >
  > backgroundColor:'red'       //如果是复合属性则必须采取驼峰命名，如果值不是数字，则需要加引号
  >
  > })

  注意：原生js中的className会覆盖元素原先里面的类名，jquery里面类操作只对指定的类操作，不影响原先的类名

  $(function(){
  		$('.current').click(function(){
  			//增加类
  			 $(this).addClass('after');
  			//删除类
  			 $(this).removeClass('after')
  			//切换类
  			$(this).toggleClass('after');
  		})
  		})

​     

- 效果：（动画效果）

  > * 显示与隐藏   
  >
  > show([speed,[easing],[fn]])     
  >
  >  hide()  
  >
  > toggle()
  >
  > > 1.参数都可以省略
  > >
  > > 2.speed：’normal' ,'slow','fast' ,毫秒数
  > >
  > > 3.easing:    默认swing，可用linear
  > >
  > > 4.fn:      回调函数，在动画完成时执行的函数，每个元素执行一次
  > >
  > > 
  >
  > * 滑动   slideDown()  slideUp()  slideToggle()
  >
  > * 淡入淡出   fadeIn()  fadeOut()  fadeToggle()  fadeTo()
  >
  > * 自定义动画  animate(params,[speed],[easing],[fn])     ******
  >
  >   > params ： 想要更改的样式属性（对象形式） 属性名可以不用引号，复合属性驼峰式，其余参数可省略
  >   >
  >   > 
  >   >
  >   > 
  >
  > * 事件切换  hover([over,] out)  
  >
  > > over:鼠标移到元素上要触发的函数（相当于mouseenter)
  > >
  > > out:鼠标移出元素要触发的函数（相当于 mouseleave)
  > >
  > > 如果只写了其中一个函数 ，则鼠标经过和离开都将出发这一个函数
  >
  > * 动画队列 及其停止排队的方法：
  >
  > > 描述：动画效果一旦触发就会执行，如果多次触发，就造成多个动画效果排队执行！
  > >
  > > 解决：stop（）方法停止排队  --stop()写到动画或效果的前面 相当于停止结束上一次的动画

- 属性操作：

  > 1.获取元素的固有属性值  prop()      即元素自带的属性，如<a>元素里的href属性
  >
  > ​    语法：prop('属性')
  >
  >    设置属性语法
  >
  > ​     语法：prop('属性'，‘属性值’)
  >
  > 2.获取/设置元素自定义的属性值  attr()
  >
  >    获取：attract('属性值')       设置： attr(’属性‘，’属性值‘)
  >
  > 3.数据缓存 data()    
  >
  >   在指定的元素上存取数据，并不会修改DOM元素 刷新后之前存放的元素都将被移除
  >
  > ​	
  >
  > ​	

- 文本属性值：

- 元素操作：

- 尺寸、位置操作：



~~ jquery的排他思想

* 想要多选一的效果。 所有元素清除样式，当前元素设置样式
* 