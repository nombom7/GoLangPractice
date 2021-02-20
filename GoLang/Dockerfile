# golang のイメージを取得
FROM golang:1.9.2

# ソースのコピー先
ENV SRC_DIR=/go/src/github.com/nombom7/GoLangPractice

# コンパイルしたファイルの格納場所
ENV GOBIN=/go/bin

# RUN, や ENTRYPOINT などを実行するワークディレクトリを指定
WORKDIR $GOBIN

# ソースをコピーする
ADD . $SRC_DIR

# コンパイル
RUN cd /go/src;
RUN go install github.com/nombom7/GoLangPractice;

# プログラム実行
ENTRYPOINT ["./nombom7/GoLangPractice"]

# 8080 ポートを解放
EXPOSE 8080

