package restutil

import (
	"encoding/json"
	"fmt"

	"github.com/SevenTV/ServerGo/src/mongo/datastructure"
	"github.com/SevenTV/ServerGo/src/utils"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status  int
	Message string
}

func (e *ErrorResponse) Send(c *fiber.Ctx, placeholders ...string) error {
	if len(placeholders) > 0 {
		e.Message = fmt.Sprintf(e.Message, placeholders)
	}

	b, _ := json.Marshal(e)
	return c.Status(e.Status).Send(b)
}

func createErrorResponse(status int, message string) ErrorResponse {
	return ErrorResponse{
		status, message,
	}
}

var (
	ErrUnknownEmote   = createErrorResponse(404, "Unknown Emote")
	ErrUnknownUser    = createErrorResponse(404, "Unknown User")
	MalformedObjectId = createErrorResponse(400, "Malformed Object ID")
	ErrInternalServer = createErrorResponse(500, "Internal Server Error (%s)")
)

func CreateEmoteResponse(emote *datastructure.Emote, owner *datastructure.User) EmoteResponse {
	// Generate URLs
	urls := make([][]string, 4)
	for i := 1; i <= 4; i++ {
		a := make([]string, 2)
		a[0] = fmt.Sprintf("%d", i)
		a[1] = utils.GetCdnURL(emote.ID.Hex(), int8(i))

		urls[i-1] = a
	}

	// Generate simple visibility value
	simpleVis := []string{}
	for vis, s := range emoteVisibilitySimpleMap {
		if !utils.BitField.HasBits(int64(emote.Visibility), int64(vis)) {
			continue
		}

		simpleVis = append(simpleVis, s)
	}

	// Create the final response
	response := EmoteResponse{
		ID:               emote.ID.Hex(),
		Name:             emote.Name,
		Owner:            nil,
		Visibility:       emote.Visibility,
		VisibilitySimple: &simpleVis,
		Mime:             emote.Mime,
		Status:           emote.Status,
		Tags:             emote.Tags,
		Width:            emote.Width,
		Height:           emote.Height,
		URLs:             urls,
	}
	if owner != nil {
		response.Owner = CreateUserResponse(owner)
	}

	return response
}

var emoteVisibilitySimpleMap = map[int32]string{
	datastructure.EmoteVisibilityPrivate:                  "PRIVATE",
	datastructure.EmoteVisibilityGlobal:                   "GLOBAL",
	datastructure.EmoteVisibilityUnlisted:                 "UNLISTED",
	datastructure.EmoteVisibilityOverrideFFZ:              "OVERRIDE_FFZ",
	datastructure.EmoteVisibilityOverrideBTTV:             "OVERRIDE_BTTV",
	datastructure.EmoteVisibilityOverrideTwitchSubscriber: "OVERRIDE_TWITCH_SUBSCRIBER",
	datastructure.EmoteVisibilityOverrideTwitchGlobal:     "OVERRIDE_TWITCH_GLOBAL",
}

type EmoteResponse struct {
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	Owner            *UserResponse `json:"owner"`
	Visibility       int32         `json:"visibility"`
	VisibilitySimple *[]string     `json:"visibility_simple"`
	Mime             string        `json:"mime"`
	Status           int32         `json:"status"`
	Tags             []string      `json:"tags"`
	Width            [4]int16      `json:"width"`
	Height           [4]int16      `json:"height"`
	URLs             [][]string    `json:"urls"`
}

func CreateUserResponse(user *datastructure.User) *UserResponse {
	response := UserResponse{
		ID:          user.ID.Hex(),
		Login:       user.Login,
		DisplayName: user.DisplayName,
		Role:        datastructure.GetRole(user.RoleID),
	}

	return &response
}

type UserResponse struct {
	ID          string             `json:"id"`
	TwitchID    string             `json:"twitch_id"`
	Login       string             `json:"login"`
	DisplayName string             `json:"display_name"`
	Role        datastructure.Role `json:"role"`
}
