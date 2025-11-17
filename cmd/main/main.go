package main

import (
	"github.com/qcserestipy/instagram-api-go-client/pkg/utils"
	"github.com/sirupsen/logrus"
)

func main() {

	err := utils.RefreshAccessToken("627978530406248")
	if err != nil {
		logrus.Fatalf("fatal error: %v", utils.ParseAPIError(err, "refresh access token"))
	}
	// reels, err := instagram.GetReels("17841464714098258", nil, nil)
	// if err != nil {
	// 	logrus.Fatalf("fatal error: %v", err)
	// }
	// for _, reel := range reels {
	// 	fmt.Printf("Reel ID: %s, Views: %d, Reach: %d, Likes: %d, Comments: %d, Shares: %d, Saves: %d, Engagement Views: %.2f%%, Caption: %s\n",
	// 		reel.ID, reel.Views, reel.Reach, reel.Likes, reel.Comments, reel.Shares, reel.Saves, reel.EngagementViews, reel.Caption)
	// }

	// stories, err := instagram.GetStories("17841464714098258")
	// if err != nil {
	// 	logrus.Fatalf("fatal error: %v", err)
	// }
	// for _, story := range stories {
	// 	fmt.Printf("Story ID: %s, Views: %d, Reach: %d, Replies: %d, Shares: %d, Navigation: %d, Profile Activity: %d, Profile Visits: %d, Follows: %d, Total Interactions: %d, Engagement Views: %.2f%%\n",
	// 		story.ID, story.Views, story.Reach, story.Replies, story.Shares, story.Navigation, story.ProfileActivity, story.ProfileVisits, story.Follows, story.TotalInteractions, story.EngagementViews)
	// }

	// 	postResp, err := media.CreateCommentOnMedia(mediaItem.ID, "This is a remark from the Instagram API Go Client!")
	// 	if err != nil {
	// 		logrus.Warnf("failed to create comment on media %s: %v", mediaItem.ID, err)
	// 	} else {
	// 		fmt.Printf("Success! Comment created on media %s: %+v\n", mediaItem.ID, postResp)
	// 	}
	// }
}
