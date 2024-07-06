package openai

import (
	"github.com/knightgao/dreame-api/relay/adaptor/ai360"
	"github.com/knightgao/dreame-api/relay/adaptor/baichuan"
	"github.com/knightgao/dreame-api/relay/adaptor/deepseek"
	"github.com/knightgao/dreame-api/relay/adaptor/doubao"
	"github.com/knightgao/dreame-api/relay/adaptor/groq"
	"github.com/knightgao/dreame-api/relay/adaptor/lingyiwanwu"
	"github.com/knightgao/dreame-api/relay/adaptor/minimax"
	"github.com/knightgao/dreame-api/relay/adaptor/mistral"
	"github.com/knightgao/dreame-api/relay/adaptor/moonshot"
	"github.com/knightgao/dreame-api/relay/adaptor/stepfun"
	"github.com/knightgao/dreame-api/relay/adaptor/togetherai"
	"github.com/knightgao/dreame-api/relay/channeltype"
)

var CompatibleChannels = []int{
	channeltype.Azure,
	channeltype.AI360,
	channeltype.Moonshot,
	channeltype.Baichuan,
	channeltype.Minimax,
	channeltype.Doubao,
	channeltype.Mistral,
	channeltype.Groq,
	channeltype.LingYiWanWu,
	channeltype.StepFun,
	channeltype.DeepSeek,
	channeltype.TogetherAI,
}

func GetCompatibleChannelMeta(channelType int) (string, []string) {
	switch channelType {
	case channeltype.Azure:
		return "azure", ModelList
	case channeltype.AI360:
		return "360", ai360.ModelList
	case channeltype.Moonshot:
		return "moonshot", moonshot.ModelList
	case channeltype.Baichuan:
		return "baichuan", baichuan.ModelList
	case channeltype.Minimax:
		return "minimax", minimax.ModelList
	case channeltype.Mistral:
		return "mistralai", mistral.ModelList
	case channeltype.Groq:
		return "groq", groq.ModelList
	case channeltype.LingYiWanWu:
		return "lingyiwanwu", lingyiwanwu.ModelList
	case channeltype.StepFun:
		return "stepfun", stepfun.ModelList
	case channeltype.DeepSeek:
		return "deepseek", deepseek.ModelList
	case channeltype.TogetherAI:
		return "together.ai", togetherai.ModelList
	case channeltype.Doubao:
		return "doubao", doubao.ModelList
	default:
		return "openai", ModelList
	}
}
