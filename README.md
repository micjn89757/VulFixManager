# VulFixMiner


## 后端
go-zero


## 前端
vue3


## 提交规范
>  type(scope): subject
1. type（必须） : commit 的类别，只允许使用下面几个标识：
- feat : 新功能
- fix : 修复bug
- docs : 文档改变
- style : 代码格式改变
- refactor : 某个已有功能重构
- perf : 性能优化
- test : 增加测试
- build : 改变了build工具 如 grunt换成了 npm
- revert : 撤销上一次的 commit
- chore : 构建过程或辅助工具的变动
- docs : 文档修改
- merge : 代码合并
- sync : 同步主线或分支的Bug

2. scope: 用于说明 commit 影响的范围，比如数据层、控制层、视图层等等，视项目不同而不同。

3. subject: commit 的简短描述，不超过50个字符。

例如:
> feat(context): add 'comments' api
> fix(controller): fix bug in 'comments' api
> docs(readme): update readme.md
> style(router): format router.js