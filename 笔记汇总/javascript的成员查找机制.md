#### JavaScript的成员查找机制

+ 当访问一个对象的属性（或方法）时，首先查找这个对象自身有没有这个属性      [ ldh.sex='男' ]
+ 没有 就查找他的原型 （即 __ proto __指向的prototype对象）                  [ Star.prototype.sex='男']
+ 还没有 查找原型对象的原型  （Object 的原型对象 ）						[ Object.prototype.sex='男 ]
+ 以此类推 直到找到Object为止 （null )
+ __ proto __对象原型的意义在于为对象成员查找机制提供一个方向 或者说一条线



#### 原型对象中的this指向问题

+ call( ) 调用此函数，修改函数运行时的this指向问题

  function.call ( thisArg , arg1 , arg2 , ...)

  thisArg :当前调用函数this的指向对象

  agr1,agr2:传递的其他参数

+ 







  扩展内置对象：

​    可以通过原型对象，对原来的内置对象进行扩展自定义的方法。比如给数组增加自定义求偶数和的功能

 注意： 数组和字符串内置对象不能给原型对象覆盖操作 Array.prototype={} 

​             ,只能是Array.prototype.xxx=function() { } 的形式



#### 继承

+ 在es6之前 没有extends继承 ，我们可以通过 构造函数 + 原型对象 模拟实现继承，被称为 组合继承

