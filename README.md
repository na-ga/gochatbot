gochatbot
====

Gochatbot is a client library to access [UserLocal Chatbot API](http://ai.userlocal.jp/).

Installation
----

```
go get github.com/na-ga/gochatbot
```

Quick start
----

To create a default client:

```go
import(
  "github.com/na-ga/gochatbot"
)

func main() {
  client, _ := gochatbot.NewClient()
  res, err := client.Chat("Hello")
  ...
}
```

Custom Client
----

To create a configured client:

```go
import(
  "github.com/na-ga/gochatbot"
)

func main() {
  config := ChatbotConfig{APIKey: "sample"}
  client, _ := gochatbot.NewClientWithConfig(config)
  res, err := client.Chat("Hello")
  ...
}
```

Example
----

Chatbot vs Chatbot:

```go
import(
  "github.com/na-ga/gochatbot"
)

func main() {
  client, _ := gochatbot.NewClient()
  message := "Hello"
  fmt.Printf("==> %s\n", message)
  for i := 0; i < 30; i++ {
    res, _ := client.Chat(message)
    message = res.Result
    fmt.Printf("<== %s\n", message)
  }
  
  // ==> Hello
  // <== 何度も何度もありがとう！
  // <== すごーっ
  // <== あれれ？
  // <== なーん
  // <== えーと...
  // <== ふーん？
  // <== え？
  // <== 同じこと何回聞くのー
  // <== えーw
  // <== ほほほー
  // <== 呼んだ?
  // <== 呼んでないー
  // <== なんて呼ぶの？
  // <== 呼んじゃおっかなー笑
  // <== にゃにかね？
  // <== にゃんでも …
  // <== ふふふ
  // <== ふふーっ
  // <== かわいいっ
  // <== でしょでしょー！！
  // <== うん
  // <== ん?
  // <== んー？
  // <== 全然わからないw
  // <== はーい
  // <== (´・ω・｀)
  // <== なでなで
  // <== して？
  // <== なんで？
  // <== なにがー？
}
```
