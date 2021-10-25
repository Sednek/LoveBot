package bot

import (
	service "LoveBot/LoveBot/pkg/service"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
	"net"
)

type BotRepo interface {
	BotListen()error
	LoveBot()error
}
type botRepo struct {
	bot     *tgbotapi.BotAPI
	s 		service.Service
	updates tgbotapi.UpdatesChannel
}

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Хочу короткий комплимент!!!"),
		tgbotapi.NewKeyboardButton("Хочу длинный комплимент"),
	))

var SimpleCompliments = []string {
		"красивая",
		"умная",
		"заботливая",
		"привлекательная",
		"сексуальная",
		"добрая",
		"нежная",
		"милая",
		"очаровательная",
		"обворожительная",
		"неповторимая",
		"неописуемая",
		"незабываемая",
		"неотразимая",
		"шикарная",
		"ослепительная",
		"страстная",
		"недоступная",
		"божественная",
		"завораживающая",
		"ангельская",
		"лучезарная",
		"сексапильная",
		"яркая",
		"пушистая",
		"обалденная",
		"сногсшибательная",
		"стройная",
		"обольстительная",
		"кокетливая",
		"утончённая",
		"грациозная",
		"весёлая",
		"энергичная",
		"креативная",
		"стильная",
		"коммуникабельная",
		"тактичная",
		"любвиобильная",
		"романтичная",
		"разносторонняя",
		"сказочная",
		"симпатичная",
		"пылкая",
		"единственная",
		"ласковая",
		"сладенькая",
		"умопомрачительная",
		"желанная",
		"непредсказуемая",
		"загадочная",
		"цветущая",
		"безупречная",
		"гармоничная",
		"отзывчивая",
		"совершенная",
		"лучшая",
		"скромная",
		"изысканная",
		"шаловливая",
		"отпадная",
		"искренная",
		"дружелюбная",
		"понимающая",
		"экстравагантная",
		"мечтательная",
		"ароматная",
		"искромётная",
		"чистолюбивая",
		"манящая",
		"восторженная",
		"бескорыстная",
		"непосредственная",
		"соблазнительная",
		"одурманивающая",
		"жизнерадостнаяя",
		"прелестная",
		"улыбчивая",
		"застенчивая",
		"зажигательная",
		"честная",
		"возбуждающая",
		"чистосердечная",
		"игривая",
		"обаятельная",
		"офигительная",
		"целеустремлённая",
		"дивная",
		"женственная",
		"блаженная",
		"бесподобная",
		"лучезарная",
		"ненаглядная",
		"необходимая",
		"изумительная",
		"сказочная",
		"трогательная",
		"миниатюрная",
		"любимая",
		"самая-самая",
	}

func NewBotRepo(bot *tgbotapi.BotAPI, s service.Service, updates tgbotapi.UpdatesChannel) (BotRepo, error) {
	return &botRepo{
		bot: bot,
		s: s,
		updates: updates,
	}, nil
}

func Connect(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println(err)
		return bot, err;
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot, err
}

func (br *botRepo) BotListen()error{
	for update := range br.updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		userIP := net.ParseIP(update.Message.Text)
		if update.Message.Text == "/start"{
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я бот, который проверяет информацию по заданному тобой айпи адресу.\nНапиши мне айпи, которое ты хочешь проверить.")
			_,err := br.bot.Send(msg)
			if err != nil{
				log.Println(err)
			}
		} else {
			if userIP == nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не верно введен айпи адрес")
				_,err := br.bot.Send(msg)
				if err != nil{
					log.Println(err)
				}
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, " Идет получение данных по заданному айпи адресу")
				_,err := br.bot.Send(msg)
				if err != nil{
					log.Println(err)
				}
				infoIP, err := br.s.GetIpInfo(update.Message.Text)
				if err != nil {
					log.Println(err)
				}
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(
					"Проверяемый айпи: %s\n"+
						"Тип айпи: %s\n"+
						"Код континента: %s\n"+
						"Континент: %s\n"+
						"Код страны: %s\n"+
						"Страна: %s\n"+
						"код региона: %s\n"+
						"Регион: %s\n"+
						"Город: %s\n"+
						"Почтовый индекс: %s\n"+
						"Широта: %f\n"+
						"Долгота: %f",
					infoIP.Ip,
					infoIP.Ip_type,
					infoIP.Continent_code,
					infoIP.Continent_name,
					infoIP.Country_code,
					infoIP.Country_name,
					infoIP.Region_code,
					infoIP.Region_name,
					infoIP.City,
					infoIP.Zip,
					infoIP.Latitude,
					infoIP.Longitude))
				_,err = br.bot.Send(msg)
				if err != nil{
					log.Println(err)
				}
			}
		}

	}
	return nil
}
func(br *botRepo) LoveBot()error{
	for update := range br.updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text ==  "Привет"{
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет пупс ❤")
			msg.ReplyMarkup = numericKeyboard
			br.bot.Send(msg)
			if update.Message.Text == "Хочу короткий комплимент!!!"{
				complimentMsg := tgbotapi.NewMessage(update.Message.Chat.ID, SimpleCompliments[rand.Intn(len(SimpleCompliments))])
				br.bot.Send(complimentMsg)
			}
		}
	}
	return nil
}