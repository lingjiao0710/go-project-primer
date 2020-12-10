# Exercise #2: URL Shortener 练习2: 网址缩短

[![exercise status: released](https://camo.githubusercontent.com/a528b56710632f59fe033b8dc08705252a9d5f8b793ee7107cb3bb2d05962890/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f65786572636973652532307374617475732d72656c65617365642d677265656e2e7376673f7374796c653d666f722d7468652d6261646765)](https://gophercises.com/exercises/urlshort)

## Exercise details 运动细节

The goal of this exercise is to create an [http.Handler](https://golang.org/pkg/net/http/#Handler) that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

这个练习的目的是创建一个 http。它会查看任何传入 web 请求的路径，并决定是否应该将用户重定向到一个新页面，就像 URL 压缩服务一样。

For instance, if we have a redirect setup for `/dogs` to `https://www.somesite.com/a-story-about-dogs` we would look for any incoming web requests with the path `/dogs` and redirect them.

例如，如果我们有一个重定向设置`/dogs` 到 `https://www.somesite.com/a-story-about-dogs` ，我们将寻找任何传入的带有 path`/dogs` 的 web 请求并重定向它们。

To complete this exercises you will need to implement the stubbed out methods in [handler.go](https://github.com/gophercises/urlshort/blob/master/handler.go). There are a good bit of comments explaining what each method should do, and there is also a [main/main.go](https://github.com/gophercises/urlshort/blob/master/main/main.go) source file that uses the package to help you test your code and get an idea of what your program should be doing.

为了完成这个练习，您需要在 handler.go 中实现这些已经存根的方法。有很多注释解释了每个方法应该做什么，还有一个 `main/main.go`源文件，使用包来帮助您测试您的代码，并得到您的程序应该做什么的想法。

I suggest first commenting out all of the code in main.go related to the `YAMLHandler` function and focusing on implementing the `MapHandler` function first.

我建议首先注释掉 main.go 中与 `YAMLHandler` 函数相关的所有代码，然后重点实现 `MapHandler` 函数。

Once you have that working, focus on parsing the YAML using the [gopkg.in/yaml.v2](https://godoc.org/gopkg.in/yaml.v2) package. *Note: You will need to `go get` this package if you don't have it already.*

一旦你做到了这一点，集中精力使用 gopkg.in/YAML.v2语言包来解析 YAML。注意: 如果你还没有这个包裹，你需要`go get`。

After you get the YAML parsing down, try to convert the data into a map and then use the MapHandler to finish the YAMLHandler implementation. Eg you might end up with some code like this:

完成 YAML 解析之后，尝试将数据转换为映射，然后使用 MapHandler 完成 YAMLHandler 实现。你可能会得到这样的代码:

```go
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
  parsedYaml, err := parseYAML(yaml)
  if err != nil {
    return nil, err
  }
  pathMap := buildMap(parsedYaml)
  return MapHandler(pathMap, fallback), nil
}
```

But in order for this to work you will need to create functions like `parseYAML` and `buildMap` on your own. This should give you ample experience working with YAML data.

但是为了实现这个功能，你需要自己创建像 `parseYAML` 和 `buildMap` 这样的函数。这将为您提供处理 YAML 数据的丰富经验。

## Bonus 额外奖励

As a bonus exercises you can also...

作为奖励，你也可以..。

1. Update the 更新[main/main.go](https://github.com/gophercises/urlshort/blob/master/main/main.go) source file to accept a YAML file as a flag and then load the YAML from a file rather than from a string. 源文件接受一个 YAML 文件作为标志，然后从一个文件而不是从一个字符串加载 YAML
2. Build a JSONHandler that serves the same purpose, but reads from JSON data. 构建一个用于相同目的，但是从 JSON 数据读取的 JSONHandler
3. Build a Handler that doesn't read from a map but instead reads from a database. Whether you use BoltDB, SQL, or something else is entirely up to you. 构建一个不从地图读取，而是从数据库读取的 Handler。是否使用 BoltDB、 SQL 或其他语言完全取决于您自己。