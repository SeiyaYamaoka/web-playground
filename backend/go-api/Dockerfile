# ----------------------------------------------------
# 1. ビルダー (build) ステージ
# ----------------------------------------------------
# goの公式イメージを使用してソースコードをビルド
FROM golang:1.21-alpine AS builder

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールファイルをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# プロジェクトの全ファイルをコピー
COPY . .

# バイナリをビルド
# -o: 出力ファイル名 (server)
# -ldflags -s -w: バイナリからデバッグ情報を除去し、軽量化
# ./cmd/server/main.go: エントリーポイント
RUN go build -o server -ldflags="-s -w" ./cmd/server/main.go

# ----------------------------------------------------
# 2. 実行 (run) ステージ
# ----------------------------------------------------
# 軽量なalpine Linuxイメージを使用
FROM alpine:latest

# タイムゾーンデータをインストール
RUN apk --no-cache add ca-certificates tzdata

# ビルドステージで作成したバイナリをコピー
# --from=builder: ビルダーコンテナから取得
COPY --from=builder /app/server /usr/local/bin/server

# コンテナがリッスンするポートを公開
EXPOSE 8080

# コンテナ起動時の実行コマンド
CMD ["/usr/local/bin/server"]