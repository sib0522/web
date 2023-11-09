# Game Server and Admin Page

### Environment
**・language**：Go<br>
**・framework**：Go Echo<br>
**・database**：MySQL<br>
**・infrastructure**：AWS<br>

### Directories
**・app**：ゲームサーバーロジック<br>
　**・api**：クライアントとサーバー間通信で使用するAPIを定義、ルーター設定<br>
　**・application**：domainを取得して各APIに対するサービスを定義<br>
　**・constants**：applicationで使用するconstantを定義<br>
　**・domain**：dbからデータを取得、更新などを行う<br>
　**・util**：ロジック実装のためのutil<br>

**・database**：データベースに接続、クエリー処理を行う<br>

**・router**：ルーター設定<br>

**・forms**：HTMLレンダリング<br>

**・web**：管理画面<br>
　**・models**：ビジネスロジック<br>
　**・views**：controllerからデータなどを受け取りブラウザに表示する<br>
　**・controller**：modelからデータを取得したりviewをRenderする<br>
