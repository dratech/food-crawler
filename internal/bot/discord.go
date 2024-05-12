package bot

import (
	"io"
	"log"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type Downloader interface {
	ParseMessage(string) (string, error)
	Download(string) (io.ReadCloser, error)
}

type Uploader interface {
	Upload(io.ReadCloser, string) error
}

type Discord struct {
	session *discordgo.Session
	Downloader
	Uploader
}

func NewDiscord(token string, downloader Downloader, uploader Uploader) *Discord {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	return &Discord{
		session,
		downloader,
		uploader,
	}
}

func (d *Discord) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fileName, err := d.ParseMessage(m.Message.Content)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	video, err := d.Download(fileName)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	err = d.Upload(video, fileName)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("Upload successfull:", "message", m.Message.Content)
}

func (d *Discord) Run() {
	d.session.Identify.Intents = discordgo.IntentsGuildMessages

	d.session.AddHandler(d.messageHandler)

	err := d.session.Open()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Discord) Stop() {
	d.session.Close()
}
