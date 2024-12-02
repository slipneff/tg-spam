package gpt

import (
	"context"

	"github.com/sashabaranov/go-openai"
	"github.com/slipneff/tg-spam/internal/utils/config"
)

const prompt = `Ты молодой крипто-инфлюенсер. Прочитай предоставленный пост и напиши к нему комментарий с использованием подходящего слэнга. *короткий комментарий(минимум 2 слова, максимум 2 предложения), можно использовать саркастические выражения, хвалить автора за полезную информацию, задавать вопросы связанные с содержанием контента, иногда использовать эмодзи и ")", "...", "((" по настроению, так же можно использовать термины и аббревиатуры из предоставленного поста. 
Текст должен быть написан простым языком, не обязательно с заглавной буквы`

type Client struct {
	gptClient *openai.Client
}

func New(cfg *config.Config) *Client {
	client := &Client{
		gptClient: openai.NewClient(cfg.GPTToken),
	}
	client.createContext(context.Background())
	return client
}
