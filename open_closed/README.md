# Open-Closed（オープン・クローズド）
- ソフトウェア構造は拡張のためにオープンであるが、変更のためにクローズされるべき
- = システムを変更しなくても、システムの動作を拡張できる

# 例
- triangleを追加する場合
- bad
  - sumAreasBadでcase文を追加する必要がある
- good
  - sumAreasGoodを触らずに、追加できる