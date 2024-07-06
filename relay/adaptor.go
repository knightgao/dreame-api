package relay

import (
	"github.com/knightgao/dreame-api/relay/adaptor"
	"github.com/knightgao/dreame-api/relay/adaptor/aiproxy"
	"github.com/knightgao/dreame-api/relay/adaptor/ali"
	"github.com/knightgao/dreame-api/relay/adaptor/anthropic"
	"github.com/knightgao/dreame-api/relay/adaptor/aws"
	"github.com/knightgao/dreame-api/relay/adaptor/baidu"
	"github.com/knightgao/dreame-api/relay/adaptor/cloudflare"
	"github.com/knightgao/dreame-api/relay/adaptor/cohere"
	"github.com/knightgao/dreame-api/relay/adaptor/coze"
	"github.com/knightgao/dreame-api/relay/adaptor/deepl"
	"github.com/knightgao/dreame-api/relay/adaptor/gemini"
	"github.com/knightgao/dreame-api/relay/adaptor/ollama"
	"github.com/knightgao/dreame-api/relay/adaptor/openai"
	"github.com/knightgao/dreame-api/relay/adaptor/palm"
	"github.com/knightgao/dreame-api/relay/adaptor/tencent"
	"github.com/knightgao/dreame-api/relay/adaptor/xunfei"
	"github.com/knightgao/dreame-api/relay/adaptor/zhipu"
	"github.com/knightgao/dreame-api/relay/apitype"
)

func GetAdaptor(apiType int) adaptor.Adaptor {
	switch apiType {
	case apitype.AIProxyLibrary:
		return &aiproxy.Adaptor{}
	case apitype.Ali:
		return &ali.Adaptor{}
	case apitype.Anthropic:
		return &anthropic.Adaptor{}
	case apitype.AwsClaude:
		return &aws.Adaptor{}
	case apitype.Baidu:
		return &baidu.Adaptor{}
	case apitype.Gemini:
		return &gemini.Adaptor{}
	case apitype.OpenAI:
		return &openai.Adaptor{}
	case apitype.PaLM:
		return &palm.Adaptor{}
	case apitype.Tencent:
		return &tencent.Adaptor{}
	case apitype.Xunfei:
		return &xunfei.Adaptor{}
	case apitype.Zhipu:
		return &zhipu.Adaptor{}
	case apitype.Ollama:
		return &ollama.Adaptor{}
	case apitype.Coze:
		return &coze.Adaptor{}
	case apitype.Cohere:
		return &cohere.Adaptor{}
	case apitype.Cloudflare:
		return &cloudflare.Adaptor{}
	case apitype.DeepL:
		return &deepl.Adaptor{}
	}
	return nil
}
