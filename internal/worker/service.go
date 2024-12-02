package worker

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/gotd/td/tg"
	"github.com/slipneff/gogger/log"
	myTg "github.com/slipneff/tg-spam/pkg/tg"
)

func (s *Worker) CatchLastPost(ctx context.Context, channelName string) {
	api := s.Client.API()
	channel, err := api.ContactsResolveUsername(ctx, channelName)
	if err != nil {
		log.Error(err, "get channel error", channelName)
	}

	if len(channel.Chats) == 0 {
		log.Error(err, fmt.Sprintf("no channel found with username %s", channelName))
	}

	var peer tg.InputPeerClass
	switch c := channel.Chats[0].(type) {
	case *tg.Channel:
		peer = &tg.InputPeerChannel{
			ChannelID:  c.ID,
			AccessHash: c.AccessHash,
		}

	default:
		log.Error(errors.New("unexpected chat type"), "unexpected chat type", channelName)
	}
	mesId, err := s.storage.GetLastMessageID(ctx, channelName)
	if err != nil {
		log.Error(err, "get last message error", channelName)
	}
	var lastPost *tg.Message
	for lastPost.ID != mesId {
		posts, err := api.MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{
			Peer:  peer,
			Limit: 1,
		})
		if err != nil {
			log.Error(err, "get messages history error", channelName)
		}
		switch posts := posts.(type) {
		case *tg.MessagesMessages:
			if len(posts.Messages) == 0 {
				log.Error(err, fmt.Sprintf("no posts found in channel %s", channelName), channelName)
			}
			lastPost = posts.Messages[0].(*tg.Message)
		case *tg.MessagesChannelMessages:
			if len(posts.Messages) == 0 {
				log.Error(err, fmt.Sprintf("no posts found in channel %s", channelName), channelName)
			}
			lastPost = posts.Messages[0].(*tg.Message)
		default:
			log.Error(err, fmt.Sprintf("unexpected messages type: %T", posts), channelName)
		}
	}

	err = s.prepareGoroutines(ctx, 1, peer, lastPost.ID, lastPost.Message)
	if err != nil {
		log.Error(err, "prepare goroutines error", channelName)
	}

	if err := s.storage.SetLastMessageID(ctx, channelName, lastPost.ID); err != nil {
		log.Error(err, "set last message id error", channelName)
	}
}

func (s *Worker) prepareGoroutines(ctx context.Context, n int, peer tg.InputPeerClass, lastPost int, msg string) error {
	sessions, err := s.storage.GetSessions(ctx, n)
	if err != nil {
		return err
	}
	for _, session := range sessions {
		go func(session string) {
			api := myTg.NewClient(session, s.cfg).GetTelegramClient(ctx).API()
			var randomID int64
			if err := binary.Read(rand.Reader, binary.BigEndian, &randomID); err != nil {
				return
			}
			// _, err = api.ChannelsJoinChannel(ctx, peer.(tg.InputChannelClass))
			// if err != nil {
			// 	log.Error(err, "join channel error")
			// 	return
			// }
			com, err := s.gptClient.NewMessage(ctx, msg)
			if err != nil {
				fmt.Println(err)
			}

			_, err = api.MessagesSendMessage(ctx, &tg.MessagesSendMessageRequest{
				Peer:     peer,
				Message:  com,
				RandomID: randomID,
				ReplyTo: &tg.InputReplyToMessage{
					ReplyToMsgID: lastPost,
				},
			})
			if err != nil {
				return
			}

		}(session)
	}
	return nil
}
