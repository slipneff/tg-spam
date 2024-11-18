package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/zelenin/go-tdlib/client"
)

func main() {
	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	var (
		apiIdRaw = os.Getenv("API_ID")
		apiHash  = os.Getenv("API_HASH")
	)

	apiId64, err := strconv.ParseInt(apiIdRaw, 10, 32)
	if err != nil {
		log.Printf("strconv.Atoi error: %s", err)
	}

	apiId := int32(apiId64)

	authorizer.TdlibParameters <- &client.SetTdlibParametersRequest{
		UseTestDc:           false,
		DatabaseDirectory:   filepath.Join(".tdlib", "database"),
		FilesDirectory:      filepath.Join(".tdlib", "files"),
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseMessageDatabase:  true,
		UseSecretChats:      false,
		ApiId:               apiId,
		ApiHash:             apiHash,
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
	}

	_, err = client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 1,
	})
	if err != nil {
		log.Printf("SetLogVerbosityLevel error: %s", err)
	}

	tdlibClient, err := client.NewClient(authorizer)
	if err != nil {
		log.Printf("NewClient error: %s", err)
	}

	optionValue, err := client.GetOption(&client.GetOptionRequest{
		Name: "version",
	})
	if err != nil {
		log.Printf("GetOption error: %s", err)
	}

	log.Printf("TDLib version: %s", optionValue.(*client.OptionValueString).Value)

	me, err := tdlibClient.GetMe()
	if err != nil {
		log.Printf("GetMe error: %s", err)
	}

	log.Printf("Me: %s %s [%v]", me.FirstName, me.LastName, me.Usernames)

	updates := client.getupdates

	go func() {
		for update := range updates {
			if update.GetClass() == client.ClassUpdate && update.GetType() == client.TypeUpdateNewMessage {
				message := update.(*client.UpdateNewMessage).Message
				if message.Content.MessageContentType() == client.TypeMessageText {
					log.Printf("New message in chat %d: %s", message.ChatId, message.Content.(*client.MessageText).Text.Text)
				}
			}
		}
	}()

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		tdlibClient.Stop()
		os.Exit(1)
	}()
}
