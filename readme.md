# 一种实现在机器上多个版本Go SDK兼容的方式 

   这是一种实现在一台机器上多个Go 版本 SDK 的兼容，并且不需要对本机安装的Go版本进行切换，只需传一个参数即可；
   
   ### 操作步骤
   **step1** 
   
   先在机器上安装两个版本的go 版本，在文件夹/usr/local/下面安装两个版本的go sdk，如下，go1.14 和 go1.18
```bash
   [root@localhost local]# ls -al
   drwxr-xr-x.  2 root root   6 Aug  3 11:10 go1.14
   drwxr-xr-x.  2 root root   6 Aug  3 11:10 go1.18
```  
 
   **step2**:
     
    把这个名为 go 的shell 放入 /usr/local/bin 目录下面，并 chmod 777 ，并将/usr/local/bin 加入 系统 PATH
   
```bash
     [root@localhost bin]# pwd
       /usr/local/bin
     [root@localhost bin]# chmod 777 go
     [root@localhost bin]# ls -al
      total 0
      drwxr-xr-x.  2 root root  16 Aug  3 11:18 .
      drwxr-xr-x. 14 root root 159 Aug  3 11:10 ..
      -rwxrwxrwx.  1 root root   0 Aug  3 11:18 go
```  
   
   
   
   **使用**
   
   如果go shell文件里面 Default_Version=1.18 默认就是1.18，如果需要使用 1.14 则需要加 -v 14参数
   ```bash
      [root@localhost bin]# go version
      go version go1.18 linux/amd64
      [root@localhost bin]# go -v 14 version
      go version go1.14 linux/amd64
      
   ```
   
