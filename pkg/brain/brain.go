package brain

import (
	"strconv"
	"strings"

	"github.com/avvero/the_gamers_guild_bot/internal/utils"
)

type Brain struct {
	memory       *Memory
	randomFactor bool
}

func NewBrain(memory *Memory, randomFactor bool) *Brain {
	return &Brain{memory: memory, randomFactor: randomFactor}
}

func (brain *Brain) Decision(chatId int64, text string) (respond bool, response string) {
	if !utils.Contains([]string{"0", "-1001733786877", "245851441", "-578279468"}, strconv.FormatInt(chatId, 10)) {
		return true, "Mr Moony presents his compliments to Professor Snape, and begs him to keep his abnormally large nose out of other people’s business."
	}
	if brain.randomFactor {
		if utils.RandomUpTo(100) == 0 {
			return new(SenselessPhrasesIntention).Express(text)
		}
		if len(text) > 5 && !strings.Contains(text, " ") && utils.RandomUpTo(10) == 0 {
			return new(HuefyLastWordIntention).Express(text)
		} else if len(text) > 5 && utils.RandomUpTo(50) == 0 {
			return new(HuefyIntention).Express(text)
		}
		if len(text) > 14 && utils.RandomUpTo(100) == 0 {
			return NewKhaleesifyIntention().Express(text)
		}
	}
	text = strings.ToLower(text)
	if text == "gg" {
		return true, "gg"
	}
	if brain.normalizeRu(text) == "нет" {
		return true, "пидора ответ"
	}
	if strings.Contains(text, "morrowind") ||
		strings.Contains(text, "морровинд") ||
		strings.Contains(text, "моровинд") {
		return true, "Morrowind - одна из лучших игр эва"
	}
	if text == "er" ||
		text == "ер" ||
		text == "эр" ||
		strings.Contains(text, "elden ring") ||
		strings.Contains(text, " er ") ||
		strings.Contains(text, " ер ") ||
		strings.Contains(text, " эр ") {
		return true, "Elden Ring - это величие"
	}
	if strings.Contains(text, "купил") {
		return true, "А не пиздишь? Аренда это не покупка"
	}
	if strings.Contains(text, "spotify") || strings.Contains(text, "спотифай") {
		return true, "Эти пидоры Антону косарик должны за подписку"
	}
	if strings.Contains(brain.normalizeEn(text), "devops") ||
		strings.Contains(brain.normalizeRu(text), "девопс") {
		return true, "Девопсы не нужны"
	}
	if strings.Contains(brain.normalizeRu(text), "трансформация") ||
		strings.Contains(brain.normalizeRu(text), "трансформацию") ||
		strings.Contains(brain.normalizeRu(text), "трансформации") {
		tokens := map[string]string{
			"трансформация": "оргия гомогеев",
			"трансформацию": "оргию гомогеев",
			"трансформации": "оргии гомогеев",
		}
		result := text
		for k, v := range tokens {
			result = strings.Replace(result, k, v, -1)
		}
		return true, result
	}
	if strings.Contains(text, "java") ||
		strings.Contains(text, "джаба") ||
		strings.Contains(text, "джава") {
		return true, "джава-хуява, а я работаю на го"
	}
	if strings.Contains(text, "блокир") {
		return true, "пусть себе анус заблокируют"
	}
	if strings.Contains(text, "опять") {
		return true, "не опять, а снова"
	}
	return false, ""
}

func (brain *Brain) normalizeRu(text string) string {
	result := text
	for k, v := range brain.GetNormalizationMap() {
		result = strings.Replace(result, k, v, -1)
	}
	return result
}

func (brain *Brain) normalizeEn(text string) string {
	result := text
	for k, v := range brain.GetNormalizationMap() {
		result = strings.Replace(result, v, k, -1)
	}
	return result
}

func (brain *Brain) GetNormalizationMap() map[string]string {
	return brain.memory.normalisationMap
}
