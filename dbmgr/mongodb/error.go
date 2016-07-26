package mongodb

import "errors"

var (
	//ErrorNotMatchInValue 期待されている型とは違った際に返されるエラーです
	ErrorNotMatchInValue = errors.New("入力値が期待されている型ではありません")
)
