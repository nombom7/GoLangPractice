package main

import (
        "fmt"
        "log"
        "net/http"
        "text/template"
)

type Person struct {
        Id string
        Name string
}

// HTML テンプレートを読み込む関数
func readerTemplate(w http.ResponseWriter, r *http.Request) {

        person := Person{Id: "1", Name: "Mohohewo"}

        // docker 上にコピーしたtemplateファイルを指定する(ファイルは、後で作成する)
        parsedTemplate, err := template.ParseFiles("/go/src/github.com/nombom7/GoLangPractice/templates/template.html")
        if err != nil {
                log.Fatalf("Not found file: %v", err)
        }
		// テンプレートに Person の値を渡す
        err = parsedTemplate.Execute(w, person)
        if err != nil {
                log.Printf("Error occurred while executing the template or writing its output: %v", err)
                return
        }
}

func main() {
        fmt.Println("API Server Start ... ")
        // ハンドラー登録
        http.HandleFunc("/", readerTemplate)
        // localhost:8080 で動くようにする
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                log.Fatal("error starting http server: ", err)
                return
        }
}