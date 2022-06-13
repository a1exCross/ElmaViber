package main

import (
	"log"
	"net/http"
)

var V = &Viber{}

func main() {
	V = Session(SessionParams{
		Token: "4eb3cc1a9f67e49c-e5dc9ffa44b129fd-5e2df787b55ec7a",
		Funcs: FuncList{
			NewMessage:       MessageFromUser,
			SeenMessage:      SeenMessage,
			DeliveredMessage: DeliveredMessage,
			Subscribed:       SubscribedUser,
			Unsibcribed:      UnsibcribedUser,
			Failed:           FailedMessage,
		},
		Sender: Sender{
			Name: "TestViber",
		},
		StartMessage: "Здравствуйте!",
	})

	go func() {
		err := V.SetWebhook(SetWebhookParams{
			Webhook: "https://248e-89-254-227-86.ngrok.io/webhook",
			Events:  []Event{Delivered, Failed, Message, Seen, Subscribed, Unsubscribed, ConversationStarted},
			/* SendName:  true,
			SendPhoto: true, */
		})

		if err != nil {
			log.Println(err)
		}
	}()

	http.HandleFunc("/webhook", V.HandleFunc)
	http.ListenAndServe(":80", nil)
}

func MessageFromUser(m MessageEvent) {
	switch m.Event {
	case Message:
		{
			e := V.GetTypeMessage(m)
			switch e {
			case Text:
				{
					log.Println(Text)
					log.Println("Пользователь", m.Sender.Name, "отправил сообщение", m.Text.Text, "типа", e)
				}

			case Picture:
				{
					log.Println(Picture)
					log.Println(m.Picture.FileName, e, m.Picture.Media, m.Picture.Thumbnail)
				}

			case Video:
				{
					log.Println(Video)
					log.Println(m.Video.FileName, m.Video.Media, m.Video.Thumbnail)
				}
			case File:
				{
					log.Println(File)
					log.Println(m.File.FileName)
				}

			case Location:
				{
					log.Println(Location)
					log.Println(m.Location.Location.Lat, m.Location.Location.Lon)
				}
			case Contact:
				{
					log.Println(Contact)
					log.Println(m.Contact.Contact.Name, m.Contact.Contact.PhoneNumber)
				}
			case Sticker:
				{
					log.Println(Sticker)
					log.Println(m.Sticker.StickerID, m.Sticker.Media)
				}
			case URL:
				{
					log.Println(URL)
					//ХЗ КАК ЭТО РАБОТАЕТ
				}
			}
		}
	}

	keyb := GetKeyboard()

	keyb.DefaultHeight = false
	keyb.BgColor = "#FFFFFF"

	keyb.AddButton(Button{
		Columns:    2,
		Rows:       2,
		BgColor:    "#2db9b9",
		Text:       "click me",
		ActionBody: "TEXTTT",
		Map: &Map{
			Latitude:  "1.123",
			Longitude: "2.432",
		},
	})

	_, err := V.SendMessageText(SendMessageTextParams{
		MessageParams: MessageParams{
			Receiver: m.Sender.ID,
			Keyboard: &keyb,
		},
		Text: "Helllomann",
	})

	if err != nil {
		log.Println(err)
	}

	/* V.SendMessagePicture(SendMessagePictureParam{
		MessageParams: MessageParams{
			Sender:   V.Sender,
			Receiver: m.Sender.ID,
		},
		Text:  "pict",
		Media: "https://static.independent.co.uk/2021/12/07/10/PRI213893584.jpg",
	}) */

	/* 	V.SendMessageVideo(SendMessageVideoParam{
		MessageParams: MessageParams{
			Receiver: m.Sender.ID,
			Sender:   V.Sender,
		},
		Media: "https://dump.video/i/nGZsAH.mp4",
		Size:  1,
	}) */

	/* V.SendMessageFile(SendMessageFileParam{
		MessageParams: MessageParams{
			Receiver: m.Sender.ID,
			Sender:   V.Sender,
		},
		Media:    "http://d.zaix.ru/tuXC.jpg",
		Filename: "_iOe4_DihIE.jpg",
		Size:     1,
	}) */

	/* V.SendMessageContact(SendMessageContactParam{
		MessageParams: MessageParams{
			Receiver: m.Sender.ID,
			Sender:   V.Sender,
		},
		Contacts: Contacts{
			Name:        "Ahmed",
			PhoneNumber: "88005553535",
		},
	}) */

	/* err := V.SendMessageLocation(SendMessageLocationParam{
		MessageParams: MessageParams{
			Receiver: m.Sender.ID,
			Sender:   V.Sender,
		},
		Location: Locate{
			Lat: "56,1",
			Lon: "51,1",
		},
	})

	if err != nil {
		log.Println(err)
	} */

	/* r, _ := V.GetAccountInfo()
	log.Println(r) */
	/* 	V.GetUserDetails(m.Sender.ID) */
	//V.GetOnline([]string{m.Sender.ID})

	/* V.SendMessageSticker(SendMessageStickerParam{
		MessageParams: MessageParams{
			Receiver: m.Sender.ID,
			Sender:   V.Sender,
		},
		StickerID: 419500,
	}) */

	/* V.SendMessageText(SendMessageTextParam{
		MessageParams: MessageParams{
			BroadcastList: []string{"CWpGLLw+q8mk2mpzx0+mTA=="},
		},
		Text: "hallo",
	}) */
}

func SeenMessage(e Events) {
	log.Println("Сообщение", e.MessageToken, "было прочитано пользователем с id", e.UserID)
}

func DeliveredMessage(e Events) {
	log.Println("Сообщение", e.MessageToken, "было доставлено")
}

func SubscribedUser(e Events) {
	log.Println("Пользователь", e.User.ID, "подписался")
}

func UnsibcribedUser(e Events) {
	log.Println("Пользователь", e.UserID, "отписался")
}

func FailedMessage(e Events) {
	log.Println("Не доставлено", e.MessageToken)
}
