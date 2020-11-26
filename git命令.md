git命令之-----分支

- git branch -v                      查看分支

- git branch  [分支名]          创建一个新分支

- git checkout [分支名]        切换分支

- git branch -a   查看远程分支

- git checkout -b XXX     新建一个分支，并切换到那个分支上

- git checkout XXX (主干)    git merge XXX (分支名)     切换至主干，在合并分支

- git branch -d xxx (分支名)   删除文件

- git status 查看冲突文件

  > 从分支A获取另一个分支B的部分代码
  >
  > 1. 切换至保留分支A
  >
  > 2. git checkout B xxx.file
  > 3. commit and push

git命令之-----合并分支

　　第一步：切换到接受修改的分支（被合并，增加新内容）上

　　　git  checkout　 [分支名（master）] 

　　第二步：执行merge命令

　　　git  merge 　 [分支名（hot_fix）] 



git命令之-----解决冲突

　　 第一步：编辑文件  删除特殊符号 

​         第二步：把冲突解决，保存退出

   	 第三步：git add [文件名] 告知冲突已解决

​		第四步：git commit  -m  '日志信息 '

​		（此时commit一定不能带具体文件名） 



git命令之-----克隆  git clone [文件名]

​		第一步：完整的把远程库下载到本地

​		第二步：创建origin远程地址别名

​		第三步：初始化本地库



   git命令之-----拉取

​       pull = fetch + merge

1. ​    git  fetch [远程库地址别名] [远程分支名]
2. ​    git merge [远程库地址别名/远程分支名]
3. ​    git pull     [远程库地址别名] [远程分支名]



git命令之-----查看提交历史

1. **git log** - 查看历史提交记录。

   **git blame <file>** - 以列表形式查看指定文件的历史修改记录。

2. ```
    git log --oneline  查看历史记录的简介版本
   ```







> git add [文件名]                    添加到暂存区
>
> git commit  -m  ' '                 提交到本地库
>
> git commit -amend			 当提交时发现漏掉了文件，可以重新提交



git基本原理：哈希算法（加密算法），git底层使用SHA-1算法。

明文-----------（哈希）------------->密文    ：哈希算法不可逆

git文件管理机制：Git把数据看作是小型文件系统的一组快照，每次提交更新时Git都会对当前的全部文件制作一个快照并保存这个快照的索引。文件无修改时，不再存储该文件，而是保留一个链接指向之前存储的文件。其工作方式可以称为快照流。





git remote add origin https://github/firefly-g/huashan.git       [将远程仓库起个origin别名代替]

git push [origin 别名]  [master 分支名]       [推送文件到远程库，并指定分支]



跨团队协作

1. fork

2. clone到本地，进行修改，然后推送到远程
3. Pull request
4. New pull request
5. Create pull request

