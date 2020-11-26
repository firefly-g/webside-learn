1.构造函数和原型

> * 构造函数：初始化对象，总与new一起使用。我们可以把对象中的有些公共的属性和方法抽出来，封装到这个函数里面。
>
>   new 在执行时会做四件事情：
>
>   1. 在内存中创建一个新的空对象
>   2. 让this指向这个空对象
>   3. 指向构造函数代码，给新对象添加属性和方法
>   4. 返回这个新对象（所以构造函数里不需要return）
>
>   构造函数缺点：有浪费内存的问题。
>
> * 构造函数原型·：prototype
>
>   构造函数通过原型分配的函数，是所有对象所共享的。
>
>   1. JavaScript规定，每一个构造函数都有一个prototype属性，这个prototype就是一个对象，这个对象所有的属性和方法都会被构造函数调用。
>   2.  我们可以把那些不变的方法。直接定义在prototype对象上，这样所有的对象的实例就可以共享这些方法。
>
> 
>



2.原型链

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

    可以通过原型对象，对原来的内置对象进行扩展自定义的方法。比如给数组增加自定义求偶数和的功能

 注意： 数组和字符串内置对象不能给原型对象覆盖操作 Array.prototype={} 

             ,只能是Array.prototype.xxx=function() { } 的形式



#### 继承

+ 在es6之前 没有extends继承 ，我们可以通过 构造函数 + 原型对象 模拟实现继承，被称为 组合继承
+ es6之后通过类 实现面向对象编程

#### 类的本质

1. class的本质还是function

2. 类的所有方法都定义在类的prototype属性上
3. 类创建的实例，里面也有__proto __ ,指向类的prototype原型对象
4. 新的class写法只是让原型对象的写法更加清晰，更像面向对象的编程写法
5. 所以es6的类其实就是语法糖
6. 语法糖：就是一种便捷写法，如有两种方法可以实现同样的功能，但是一种写法更简单，方便，那么这个方法就是语法糖。



#### ES5新增的方法

* 数组方法

  - 迭代（遍历）方法：forEach()  map()   filter()  some()  every()

  > array.forEach(function( currentValue , index , arr ))
  >
  > currentValue :数组当前的值
  >
  > index ：数组当前的索引
  >
  > arr：数组对象本身

  - filter方法创建一个新数组，新数组中的元素是通过检查指定数组中符合条件的所有元素，主要用于筛选数组

  > array.filter (function( currentValue , index , arr ))
  >
  > 注意：直接返回一个新数组

  - some()方法 用于检查数组中的元素是否满足指定条件，（查找数组中是否有满足条件的元素）

  > array.some(function( currentValue , index , arr ))
  >
  > 如果找到第一个满足条件的元素，则终止循环返回 true,找不到就返回 false

* 字符串方法

* 

* 对象方法

