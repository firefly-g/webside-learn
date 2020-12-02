#### vue试错之  深浅拷贝

1.vue elementUi在表单中修改数据，表格的数据也跟着修改的问题

+ 点击table中的数据，弹出修改对话框

+  showEditDialog(index,row){

    console.log(row)

    this.editDialogVisible=true

    this.editForm = row        //将浅拷贝，当输入修改数据的同时,table中的数据已经更新

    }

+ 因为row是Object对象类型（引用类型），如果直接赋值的话，就变成了浅拷贝，复制的是地址，导致在表单中改变值的时候table中的数据也跟着改变，所以要进行深拷贝，利用json就可以了，改成下面就行了

+ this.editForm = JSON.parse(JSON.stringify(row)) 