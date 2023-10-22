package ApiUserLogin

type Request struct {
	Uuid string
}

type Response struct {
	Uuid  string
	Level uint32
	Exp   uint64
}

// Requestから送られてきたuuidが空だったら、新しいユーザーデータを作成し、それをResponseとして送る
// Requestから値が入っているuuidが送られてきたら、サーバー側でuuidのデータを取得し、それをResponseとして送る
