package aiproxy

import "github.com/knightgao/dreame-api/relay/adaptor/openai"

var ModelList = []string{""}

func init() {
	ModelList = openai.ModelList
}
