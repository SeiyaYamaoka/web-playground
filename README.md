# web-playground

## ディレクトリ構成

```
web-playground/
├─ README.md
├─ docker-compose.yml
├─ .env.example
│
├─ backend/
│  └─ go-api/
│     ├─ cmd/
│     │  └─ server/
│     │     └─ main.go
│     ├─ internal/
│     │  ├─ http/        # ハンドラ・ルーティング
│     │  ├─ service/     # ビジネスロジック
│     │  ├─ repository/  # DBアクセス（使うなら）
│     │  ├─ domain/
│     │  ├─ model/
│     │  └─ config/
│     ├─ pkg/            # 共通ユーティリティ（ログ、レスポンス共通処理など）
│     ├─ Dockerfile
│     └─ go.mod / go.sum
│
├─ frontend/
│  ├─ react-app/
│  │  ├─ src/
│  │  │  ├─ api/         # APIクライアント（axios/fetch ラッパ）
│  │  │  ├─ components/
│  │  │  ├─ pages/
│  │  │  ├─ hooks/
│  │  │  └─ types/
│  │  ├─ public/
│  │  ├─ Dockerfile
│  │  ├─ package.json
│  │  └─ vite.config.ts  # or webpack 等
│  │
│  └─ vue-app/
│     ├─ src/
│     │  ├─ api/         # Reactと似せておくと見返しやすい
│     │  ├─ components/
│     │  ├─ pages/
│     │  ├─ composables/ # Vueのhooks的なやつ
│     │  └─ types/
│     ├─ public/
│     ├─ Dockerfile
│     ├─ package.json
│     └─ vite.config.ts
│
├─ shared/
│  ├─ api-spec/
│  │  └─ openapi.yaml    # そのうちOpenAPI書きたくなったとき用
│  └─ schemas/
│     └─ sample.json
│
└─ docs/
   ├─ architecture.md     # 全体構成メモ（ポート、パス、コンテナ名など）
   ├─ backend-notes.md    # Go側で試したことメモ
   ├─ react-notes.md      # Reactで試したことメモ
   └─ vue-notes.md        # Vueで試したことメモ
```
