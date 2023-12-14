FROM golang:1.21.3

# モジュールを使用して依存関係を管理
ENV GO111MODULE=on

# ワーキングディレクトリの設定
WORKDIR /app

# 依存関係のインストール
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# ソースコードとairの設定ファイルのコピー
COPY / .
COPY .air.toml ./

# .envファイルのコピー
COPY .env ./

# airのインストール
RUN go install github.com/cosmtrek/air@latest

# ポートを開放
EXPOSE 8080

# airでアプリケーションを起動
CMD ["/go/bin/air"]
