# Single Responsibility（単一責任の原則）
- 1つのクラスは1つの責務を持つべきであるという原則
- 各機能を独立させ、変更が必要な場合も他の部分に影響を与えにくくする

# 例
- メールを送って、DBに保存するロジック
- bad
  - EmailService structはdbとsmtp両方の責任がある
  - Sendはdbとsmtpのデータ送信責任がある

- good
  - dbとsmtpの責務が別々になった
  - さらにEmailRepositoryインタフェースを使うことで、上位レイヤー（ここだとdomain（ビジネスロジック） -> infrastructure層（DB書込）のdomain）はリポジトリのインタフェースのみ参照するためリポジトリ内部でどんなDBを使用しているのかとか、どんな SQL や ORM を使用しているのかなどを知らなくて済む
