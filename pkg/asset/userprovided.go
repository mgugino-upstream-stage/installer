package asset

import (
	"github.com/AlecAivazis/survey"
)

// UserProvided generates an asset that is supplied by a user.
type UserProvided struct {
	AssetName string
	Prompt    survey.Prompt
}

var _ Asset = (*UserProvided)(nil)

// Dependencies returns no dependencies.
func (a *UserProvided) Dependencies() []Asset {
	return []Asset{}
}

// Generate queries for input from the user.
func (a *UserProvided) Generate(map[Asset]*State) (*State, error) {
	var response string
	survey.AskOne(a.Prompt, &response, survey.Required)

	return &State{
		Contents: []Content{{
			Data: []byte(response),
		}},
	}, nil
}

// Name returns the human-friendly name of the asset.
func (a *UserProvided) Name() string {
	return a.AssetName
}
