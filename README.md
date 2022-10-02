# golang CI

一个基于pipeline基本思想的自动化部署工具(利用github进行自动化部署)


## v1.0

可支持pipeline语法

pipeline{
      parameters:{
          repository:""(should be ssh link),
          ssh:"absolutely location",
          branch:default(main)
          dockerfileName:
}



`}