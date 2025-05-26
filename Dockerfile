# ベースイメージ（Python 3.11）
FROM python:3.11-slim

# 作業ディレクトリ作成
WORKDIR /app

# 依存関係ファイルをコピーしてインストール
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# アプリのコードをコピー
COPY . .

# Cloud Runはポート8080で起動する必要があるっちゃ！
ENV PORT=8080

# アプリを起動
CMD ["python", "main.py"]
