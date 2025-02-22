package rules

import (
	"github.com/bluesky-social/indigo/automod"
)

func DefaultRules() automod.RuleSet {
	rules := automod.RuleSet{
		PostRules: []automod.PostRuleFunc{
			//MisleadingURLPostRule,
			//MisleadingMentionPostRule,
			ReplyCountPostRule,
			BadHashtagsPostRule,
			//TooManyHashtagsPostRule,
			//AccountDemoPostRule,
			AccountPrivateDemoPostRule,
			GtubePostRule,
			KeywordPostRule,
			ReplySingleKeywordPostRule,
			AggressivePromotionRule,
			IdenticalReplyPostRule,
			DistinctMentionsRule,
		},
		ProfileRules: []automod.ProfileRuleFunc{
			GtubeProfileRule,
			KeywordProfileRule,
		},
		RecordRules: []automod.RecordRuleFunc{
			InteractionChurnRule,
		},
		RecordDeleteRules: []automod.RecordRuleFunc{
			DeleteInteractionRule,
		},
		IdentityRules: []automod.IdentityRuleFunc{
			NewAccountRule,
		},
	}
	return rules
}
