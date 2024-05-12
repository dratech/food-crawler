package main

import (
	"food-crawler/internal/bot"
	"food-crawler/internal/downloader"
	"food-crawler/internal/uploader"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	discordToken := os.Getenv("DISCORD_TOKEN")
	dropboxToken := os.Getenv("DROPBOX_TOKEN")

	tiktok := downloader.NewTikTok()
	dropbox := uploader.NewDropbox(dropboxToken)
	discord := bot.NewDiscord(discordToken, tiktok, dropbox)

	discord.Run()
	defer discord.Stop()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
