# 练习1：测验游戏



## 细节

本练习分为两部分，以帮助简化说明过程并使其更易于解决。第二部分比第一部分困难，因此，如果您陷入困境，可以继续进行另一个问题，然后稍后再回到第二部分。

*注意：我没有像做某些练习那样将其分为多个练习，因为这两个练习合起来只需要约30m即可被截屏。*

### 第1部分

创建一个程序，该程序将读取通过CSV文件提供的测验（以下更多详细信息），然后将测验提供给用户，以跟踪他们正确回答了多少个问题以及错误回答了多少个问题。无论答案是正确还是错误，都应在之后立即询问下一个问题。

CSV文件应默认为`problems.csv`（如下所示的示例），但用户应能够通过标记来自定义文件名。

CSV文件将采用以下格式，其中第一列是问题，而同一行中的第二列是该问题的答案。

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

您可以假设测验会相对较短（少于100个问题），并且只有一个单词/数字答案。

在测验结束时，程序应输出正确的问题总数以及总共有多少个问题。给出无效答案的问题被视为不正确。

**注意：** *CSV文件中可能包含逗号问题。例如：`"what 2+2, sir?",4`是CSV中的有效行。我建议您在Go中查看CSV包，不要尝试编写自己的CSV解析器。*

### 第2部分

修改第1部分中的程序以添加计时器。默认时间限制应为30秒，但也可以通过标记进行自定义。

超过时间限制后，测验应立即停止。也就是说，您不应该等待用户回答一个最后的问题，而是最好完全停止测验，即使您当前正在等待最终用户的回答。

应该要求用户在计时器启动之前按Enter键（或其他某个键），然后将问题一次打印到屏幕上，直到用户提供答案为止。不论答案是正确还是错误，都应询问下一个问题。

在测验结束时，程序仍应输出正确的问题总数以及总共有多少个问题。给出无效答案或未回答的问题被视为不正确。