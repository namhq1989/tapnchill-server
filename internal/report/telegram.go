package report

import (
	"fmt"
	"github.com/namhq1989/go-utilities/appcontext"
	"time"
)

func (r Report) sendTelegramMessage(ctx *appcontext.AppContext, message string) error {
	_, err := r.telegramClient.R().
		SetQueryParams(map[string]string{
			"chat_id":    r.telegramChatID,
			"text":       message,
			"parse_mode": "HTML",
		}).
		SetResult(map[string]interface{}{}).
		Get(fmt.Sprintf("/bot%s/sendMessage", r.telegramBotToken))

	if err != nil {
		ctx.Logger().Error("[report] error when send telegram message", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Info("[report] telegram message sent", appcontext.Fields{"message": message})

	return nil
}

func (r Report) NewUserSignedInWithGoogle(ctx *appcontext.AppContext, userID, email string) error {
	// Prepare formatted message
	message := fmt.Sprintf(
		"🎉 <b>New User Signed In With Google</b>\n"+
			"<b>Time:</b> %s\n"+
			"<b>User ID:</b> <code>%s</code>\n"+
			"<b>Email:</b> %s",
		time.Now().Format(time.RFC1123),
		userID,
		email,
	)

	return r.sendTelegramMessage(ctx, message)
}

func (r Report) NewUserFeedback(ctx *appcontext.AppContext, feedbackID, content string) error {
	displayContent := content
	if len(content) > 500 {
		displayContent = content[:497] + "..."
	}

	// Prepare formatted message
	message := fmt.Sprintf(
		"📝 <b>New User Feedback Received</b>\n"+
			"<b>Time:</b> %s\n"+
			"<b>Feedback ID:</b> <code>%s</code>\n"+
			"<b>Content:</b>\n\n%s",
		time.Now().Format(time.RFC1123),
		feedbackID,
		displayContent,
	)

	return r.sendTelegramMessage(ctx, message)
}
