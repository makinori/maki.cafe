package config

import (
	"log/slog"
	"runtime/debug"
	"time"
)

const (
	Domain = "maki.cafe"

	Description = "site of maki"
	siteImage   = "https://" + Domain + "/images/maki.jpg"

	Email = "maki@hotmilk.space"
	XMPP  = "maki@hotmilk.space"

	Tox    = "F5A8FBAB3464147B2733AA2781B8CE56C24EF93C196E65BA2E142682428C322646DE03EA9648"
	ToxURI = "tox:" + Tox

	GitHubUsername = "makinori"
	GitHubURL      = "https://github.com/" + GitHubUsername

	ForgejoDomain = "git.hotmilk.space"
	ForgejoURL    = "https://" + ForgejoDomain + "/explore/repos"

	AniListUsername = "makinori"
	AniListURL      = "https://anilist.co/user/" + AniListUsername

	MastodonUsername = "@maki@hotmilk.space"
	MastodonURL      = "https://mastodon.hotmilk.space/@maki"

	MatrixUsername = "@maki:hotmilk.space"
	MatrixURL      = "https://matrix.to/#/" + MatrixUsername

	SquirrelsUsername = "@squirrels@hotmilk.space"
	SquirrelsURL      = "https://mastodon.hotmilk.space/@squirrels"

	SecondLifeName = "norimaki.resident"
	SecondLifeUUID = "b7c5f366-7a39-4289-8157-d3a8ae6d57f4"
	SecondLifeURL  = "secondlife:///app/agent/" + SecondLifeUUID + "/about"

	// BackloggdUsername = "maki_nori"
	// BackloggdURL      = "https://backloggd.com/u/" + BackloggdUsername
)

func getGitCommitAndBuildDate() (gitCommit string, gitTime time.Time) {
	info, _ := debug.ReadBuildInfo()
	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			gitCommit = setting.Value[:min(8, len(setting.Value))]
		case "vcs.time":
			var err error
			gitTime, err = time.Parse(time.RFC3339, setting.Value)
			if err != nil {
				panic("failed to parse vcs.time: " + err.Error())
			}
		}
	}
	slog.Info("git", "commit", gitCommit, "time", gitTime)
	return
}

var (
	GitCommit, GitTime = getGitCommitAndBuildDate()
	SiteImage          = siteImage + "?" + GitCommit
)
