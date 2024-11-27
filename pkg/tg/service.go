package tg

import (
	"context"
	"os"

	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/slipneff/gogger/log"
)

func (c *Client) getSessionsString() (string, error) {
	if c.cfg.SessionsPath == "" {
		log.Error(ErrSessionsPathNotSet, "empty session path in cfg")
		return "", ErrSessionsPathNotSet
	}
	if _, err := os.Stat(c.cfg.SessionsPath); err != nil {
		return "", err
	}
	
	file, err := os.Open(c.cfg.SessionsPath + c.Name)
	if err != nil {
		log.Error(err, "open sessions path error")
		return "", err
	}
	defer file.Close()

	session, err := os.ReadFile(file.Name())
	if err != nil {
		log.Error(err, "read file error")
		return "", err
	}
	return string(session), nil
}

func (c *Client) telethonSession() (*session.Data, error) {
	sessions, err := c.getSessionsString()
	if err != nil {
		return nil, err
	}

	data, err := session.TelethonSession(sessions)
	if err != nil {
		log.Error(err, "get telethon session error")
		return nil, err
	}

	return data, nil
}

func (c *Client) getStorage(ctx context.Context) (*session.StorageMemory, error) {
	telethonSession, err := c.telethonSession()
	if err != nil {
		return nil, err
	}

	storage := &session.StorageMemory{}
	loader := session.Loader{Storage: storage}
	if err = loader.Save(ctx, telethonSession); err != nil {
		log.Error(err, "save session error")
		return nil, err
	}

	return storage, nil
}

func (c *Client) GetTelegramClient(ctx context.Context) *telegram.Client {
	st, err := c.getStorage(ctx)
	if err != nil {
		log.Error(err, "get storage error")
		return nil
	}
	return telegram.NewClient(c.cfg.Telegram.AppID, c.cfg.Telegram.AppHash, telegram.Options{
		SessionStorage: st,
	})
}


