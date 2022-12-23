# ZeroBot-Plugin (DIY)

---

> 自用插件收集。

## ☆简介

本仓库为[ZeroBot-Plugin](https://github.com/FloatTech/ZeroBot-Plugin)的粉丝（也就是我）创建的自制ZeroBot-Plugin合集🤤🤤

学习是一个漫长的过程。本仓库用于记录我在学习过程中写出来的一些姑且可以润起来的小东西。

## ☆使用

`git clone`本仓库，选择你想要的插件文件夹复制到你自己的`ZeroBot-Plugin`里的`plugin`文件夹。在`main.go`中的`import`列表加上：
```
_ "github.com/FloatTech/ZeroBot-Plugin/plugin/插件文件夹名"
```
然后`go build`即可。

## ☆列表

<details>
 <summary>example</summary>

字面意思，供自己复习用的插件模板。用来测试姬气人是否掉线/纪念自己的第一个插件。

</details>
<details>
 <summary>jrrp</summary>

应同学要求写的一个人品插件。特点是一天只有一个结果，发送时附带图片。

时间限制思路提供：[不愿意透露姓名的btNFC](https://github.com/myrnfc)

</details>
<details>
 <summary>gpbt200</summary>

网上冲浪时发现的api，便决定应用于自己的插件之中。~这也是我第一次使用GoLang调用json，也是我第一次使用api制作插件，对于我来说是个很有学习和纪念意义的插件。~

生成的文章质量参差不齐~，要怪就怪api吧（×~

</details>
<details>
 <summary>dice</summary>

一个简单的掷骰小插件。想复刻一下速度回复小姐写的[dice!](https://v2docs.kokona.tech/zh/latest/index.html)的内容，故诞生此插件。（[该bt仓库大概也是这样想的吧](https://github.com/myrnfc/ZeroBot-Plugin-Dice/tree/61ed586fb870d34b07f260c53b3f70d985634d07) [该bt仓库的最新版](https://github.com/myrnfc/ZeroBot-Plugin-Dice)）但又想了想——本来zbp也不是骰娘项目呀——所以放弃全部内容的复刻，仅保留掷骰的基本功能。~手法十分稚嫩还请过路dalao手下留情~

可能接下来还会有关于该插件的更新？
  - [x] 。ra（一个附带成功率的骰子）
  - [ ] 。reply（教学系统（可能会分到其他插件的制作规划里））

</details>
<details>
 <summary>card21</summary>


进入新学校，开始怀念初中时和同学们在课间玩的民间规则21点了，于是复刻了一个。加上debug，整个过程用时一天不到一点。可能在学校里py写久了对写go也有一定帮助？毕竟以往要是有想法复刻某个游戏的话，可没办法用这个速度写出来。总之能进步总是好事！

民间规则21点，顾名思义，规则与在赌场里玩的那些21点不一样。并且目前的版本我也没有引入已有的货币系统。 ~~所以这是一个清纯的21点。~~ 普通用户玩起来可能会不太习惯，不过总体上还是一样的。但是我还是来写一下大致规则好了，防止想要玩这个魔改版游戏的用户在网上找不到规则。

加入游戏后，用户按顺序抽牌，抽完牌后直到本轮游戏结束再也无法抽牌，并且由姬气人当场结算该玩家的分数。所有玩家完成自己的回合后，姬气人将汇总所有玩家的分数，并报出冠军。


**注意：由于程序设计限制，一轮游戏所有玩家总共能抽52张牌，所以参与人数不宜过过过过多。**

**注意：由于程序设计限制，同分者记玩家列表中靠前者为冠军。**

还是希望大家能玩得开心就好。

感谢zbp群内大佬们的细节指导。名字不一一放出，~~有兴趣可以来群玩。~~

</details>

## ☆TODO

其实还有很多想要写的插件，奈何实力不允许，只能慢慢摸索学习了....

话说一直用句号是不是太过严肃了？那么以颜文字结尾吧！

\*\~\\(ΦωΦ)/\~\*\*\~\\(ΦωΦ)/\~\*\*\~\\(ΦωΦ)/\~\*
