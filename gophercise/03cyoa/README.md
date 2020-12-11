# Exercise #3: Choose your own adventure 练习3: 选择你自己的冒险

[![exercise status: released](https://camo.githubusercontent.com/a528b56710632f59fe033b8dc08705252a9d5f8b793ee7107cb3bb2d05962890/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f65786572636973652532307374617475732d72656c65617365642d677265656e2e7376673f7374796c653d666f722d7468652d6261646765)](https://gophercises.com/exercises/cyoa) [![demo: ->](https://camo.githubusercontent.com/aac41289ff69b7a5b4932925354ffd899a7fd8d9e69eb0e1e4eedbcbe3474c4f/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64656d6f2d2545322538362539322d626c75652e7376673f7374796c653d666f722d7468652d6261646765)](https://gophercises.com/demos/cyoa/)

## Exercise details 运动细节

[Choose Your Own Adventure](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure) is (was?) a series of books intended for children where as you read you would occasionally be given options about how you want to proceed. For instance, you might read about a boy walking in a cave when he stumbles across a dark passage or a ladder leading to an upper level and the reader will be presented with two options like:

选择你自己的冒险是(过去?)这是一系列专为儿童设计的书籍，在你阅读的过程中，你会偶尔得到一些关于如何继续下去的选择。例如，你可能会读到一个男孩在山洞里行走，当他绊倒在通往上一层的黑暗通道或梯子时，读者会有两种选择:

- Turn to page 44 to go up the ladder. 翻到第44页去爬梯子
- Turn to page 87 to venture down the dark passage. 翻到第87页，沿着黑暗的通道冒险

The goal of this exercise is to recreate this experience via a web application where each page will be a portion of the story, and at the end of every page the user will be given a series of options to choose from (or be told that they have reached the end of that particular story arc).

这个练习的目的是通过一个 web 应用程序重新创建这种体验，每个页面都是故事的一部分，在每个页面的结尾，用户将有一系列的选项可以选择(或者被告知他们已经到达了故事弧的结尾)。

Stories will be provided via a JSON file with the following format:

故事将通过 JSON 文件提供，格式如下:

```go
{
  // Each story arc will have a unique key that represents
  // the name of that particular arc.
  "story-arc": {
    "title": "A title for that story arc. Think of it like a chapter title.",
    "story": [
      "A series of paragraphs, each represented as a string in a slice.",
      "This is a new paragraph in this particular story arc."
    ],
    // Options will be empty if it is the end of that
    // particular story arc. Otherwise it will have one or
    // more JSON objects that represent an "option" that the
    // reader has at the end of a story arc.
    "options": [
      {
        "text": "the text to render for this option. eg 'venture down the dark passage'",
        "arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
      }
    ]
  },
  ...
}
```

*See [gopher.json](https://github.com/gophercises/cyoa/blob/master/gopher.json) for a real example of a JSON story. I find that seeing the real JSON file really helps answer any confusion or questions about the JSON format.*

有关 JSON 故事的真实示例，请参见 gopher. JSON。我发现，看到真正的 JSON 文件确实有助于解答任何有关 JSON 格式的困惑或问题。

You are welcome to design the code however you want. You can put everything in a single `main` package, or you can break the story into its own package and use that when creating your http handlers.

欢迎您以任何方式设计代码。您可以将所有内容放在一个单独的主包中，或者可以将故事分解成它自己的包，并在创建 http 处理程序时使用它。

The only real requirements are:

唯一真正的要求是:

1. Use the 使用`html/template` package to create your HTML pages. Part of the purpose of this exercise is to get practice using this package. 这个练习的部分目的是练习使用这个包
2. Create an 创建一个`http.Handler` to handle the web requests instead of a handler function. 来处理 web 请求，而不是处理函数
3. Use the 使用`encoding/json` package to decode the JSON file. You are welcome to try out third party packages afterwards, but I recommend starting here. 包来解码 JSON 文件。您可以在之后尝试第三方软件包，但我建议从这里开始

A few things worth noting:

值得注意的是:

- Stories could be cyclical if a user chooses options that keep leading to the same place. This isn't likely to cause issues, but keep it in mind. 如果用户选择的选项一直指向相同的位置，故事可能是循环的。这不太可能引起问题，但要记住
- For simplicity, all stories will have a story arc named "intro" that is where the story starts. That is, every JSON file will have a key with the value 为了简单起见，所有的故事都有一个名为“ intro”的故事弧，这是故事的开始。也就是说，每个 JSON 文件都有一个带有值的键`intro` and this is where your story should start. 你的故事应该从这里开始
- Matt Holt's JSON-to-Go is a really handy tool when working with JSON in Go! Check it out - Matt Holt 的 JSON-to-Go 在 Go 中使用 JSON 时是一个非常方便的工具https://mholt.github.io/json-to-go/

## Bonus 额外奖励

As a bonus exercises you can also:

作为奖励，你还可以:

1. Create a command-line version of our Choose Your Own Adventure application where stories are printed out to the terminal and options are picked via typing in numbers ("Press 1 to venture ..."). 创建一个命令行版本的我们的选择自己的冒险应用程序，其中的故事是打印到终端和选项是通过键入数字(“按1冒险... ”)
2. Consider how you would alter your program in order to support stories starting form a story-defined arc. That is, what if all stories didn't start on an arc named 考虑如何修改程序以支持从故事定义的弧线开始的故事。也就是说，如果所有的故事都不是以一个叫做`intro`? How would you redesign your program or restructure the JSON? This bonus exercises is meant to be as much of a thought exercise as an actual coding one. ？您将如何重新设计您的程序或重新构造 JSON？这个奖金练习意味着作为一个实际的编码一样多的思想练习