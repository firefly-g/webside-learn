/*Tab对象:
1.具有切换功能
2.具有添加功能
3.具有删除功能
4.具有修改功能
*/


/*
Tab栏的切换功能
1. 点击+ 创建新的li和section
2. 把创建的元素追加到最近的父元素中
 */


/*
Tab栏的删除功能
1. 点击X 可以删除当前的li选项卡和当前的section
2. X 没有索引号 但是可以通过其父亲li的索引号
3.核心：点击X号可以删除这个索引号对应的li和section
 */    
 
 
/*
Tab栏的编辑功能
1. 双击选项卡li或者section里面的文字，可以实现修改功能
2.双击事件是：ondblclick
3.（如果双击文字，会默认选定文字）,此时需要双击禁止选中文字
4.解决：window.getSelection?window.getSelection().removeAllRanges();document.selection.empty();
5.核心思路：双击文字的时候，在里面生成一个文本框，当失去焦点的时候或者按下回车，吧文本框的值输给原先元素即可
*/  
 
var that;
class Tab {
	constructor(id){
		that=this;                                                 //获取元素
		this.main=document.querySelector(id);
//		this.lis=this.main.querySelectorAll('li');             //获取大盒子里面的所有小li
//		this.sections=this.main.querySelectorAll('section');    //获取大盒子里面的所有小section	
		this.add=this.main.querySelector('.tabadd');
		this.ul=this.main.querySelector('.firstnav ul:first-child');  //li的父元素
		this.fsection=this.main.querySelector('.tabscon');             //section的父元素
		
//		this.remove=this.main.querySelectorAll('.icon-guanbi');
		this.init();
	}
	
	//获取所有的li和section
	updateNode(){
		this.lis=this.main.querySelectorAll('li');             
		this.sections=this.main.querySelectorAll('section'); 
		this.remove=this.main.querySelectorAll('.icon-guanbi');
		this.spans=this.main.querySelectorAll('.firstnav li span:first-child');
	}
		init(){
		   this.updateNode();													//init 初始化操作让相关的元素绑定事件
		   for(var i=0; i<this.lis.length;i++){
			  this.lis[i].index=i;								//设置li的index属性
			  this.lis[i].onclick=this.toggleTab;               //绑定切换功能
			  this.remove[i].onclick=this.removeTab;
			  this.spans[i].ondblclick=this.editTab;
			  this.sections[i].ondblclick=this.editTab;
		    }
		   this.add.onclick=this.addTab;                        //绑定添加事件
		}
	
	//1.切换
	toggleTab(){
		//console.log(this.index);
        that.clearClass();
		this.className='liactive';                             //点中哪个li，为其添加样式(此时this指向小li)
		that.sections[this.index].className='conactive';       //为section添加响应样式
		
	}
	clearClass(){
		for(var i=0;i<this.lis.length;i++){
			this.lis[i].className='';
			this.sections[i].className='';
			
		}
	}
	//2.添加
	addTab(){
	that.clearClass();
    var random = Math.random();
	var li = '<li class="liactive"><span>新选项卡</span><span class="iconfont icon-guanbi"></span></li>'	;  //(1)创建li元素和section元素
	var section ='<section class="conactive">测试'+random+'</section>';
	that.ul.insertAdjacentHTML('beforeend',li);	                    //(2)追加到对应的父元素里
	that.fsection.insertAdjacentHTML('beforeend',section);
	that.init();
	}
	//3.删除功能
	removeTab(e){
		e.stopPropagation();                          //阻止事件冒泡
		var index=this.parentNode.index;              //获取父元素的索引号
		//console.log(index);
		that.lis[index].remove();                     //根据索引删除对应的li和section
		that.sections[index].remove();
		that.init();                                  //重新初始化
		
		if(document.querySelector('.liactive'))//当我们删除的不是选中状态的li时候，原来的选中状态保持不变
		return;
		else{
		index--;                                      //细节：当删除选定的li标签之后,让其前一个li变为选中状态
		that.lis[index]&&that.lis[index].click();                      //手动调用点击事件，不需要鼠标触发
		}
		
	}
	//4.修改功能
	editTab(e){
		e.stopPropagation();
		var str =this.innerHTML;                        //把原来的html内容获取过来
		                                                //禁止选中文字
    window.getSelection ? window.getSelection().removeAllRanges() : document.selection.empty();
    this.innerHTML='<input type="text"/>';              //文本框是span的第一个孩子
    var input =this.children[0];                        //获取文本框
    input.value=str;									//将原始值赋值给文本框
    input.select();										//将文本框内容设为选中状态
    
    input.onblur=function(){							//当离开文本框时就把文本框里面的只传给span
    this.parentNode.innerHTML = this.value;				//此时的this指向input，将input的值赋给span
    }
    
    //按下回车也可以将文本框内容传给span
      input.onkeyup=function(e){
       if(e.keyCode===13){
    	 this.blur();
         }	                                               //判断回车键 码值为13
      }
	}
}

var tab = new Tab('#tab');





/*
 * 
 在删除操作时出现的问题：
 1.删除标签时 ，借助X图标的父元素li的索引号定位标签。实际操作时x的点击事件会往上冒泡，触发li的点击事件..
 2.当删除最后一位选项卡是，希望页面呈现出前一个选项卡的内容..
 3.但删除非最后一位标签时，希望页面锁定在之前选中状态..
 
 */