package stephandler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

type newBody struct {
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Content     string       `json:"content,omitempty"`
	Category    data.StepCat `json:"category,omitempty"`
	XPAward     int          `json:"xp_award,omitempty"`
	CoinsAward  int          `json:"coins_award,omitempty"`
	EnergyCost  int          `json:"energy_cost,omitempty"`

	Parent      string `json:"parent,omitempty"`
	BranchIndex int    `json:"branch_index,omitempty"`
}

func New(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var body newBody
	json.NewDecoder(r.Body).Decode(&body)

	if (body == newBody{}) {
		return handler.ErrEmptyBody
	}

	step := data.Step{
		Name:        body.Name,
		Tag:         data.Tagify(body.Name, true),
		Description: body.Description,
		Content:     body.Content,
		Category:    body.Category,
		XPAward:     body.XPAward,
		CoinsAward:  body.CoinsAward,
		EnergyCost:  body.EnergyCost,
		Children:    [][]string{},
	}
	err := stores.Steps.Put(step)
	if err != nil {
		return err
	}

	// parenting
	if body.Parent == "" { // return early if parent is empty
		handler.WriteJSON(w, 200, step)
		return nil
	}

	if strings.HasPrefix(body.Parent, "adv:") {
		ptag := strings.Replace(body.Parent, "adv:", "", 1)
		parent, err := stores.Adventures.One(db.Query{"tag": ptag})
		if err != nil {
			return err
		}

		parent.Steps = append(parent.Steps, step.Tag)
		err = stores.Adventures.Update(db.Query{"tag": parent.Tag}, *parent)
		if err != nil {
			return err
		}
	} else if strings.HasPrefix(body.Parent, "step:") {
		ptag := strings.Replace(body.Parent, "step:", "", 1)
		parent, err := stores.Steps.One(db.Query{"tag": ptag})
		if err != nil {
			return err
		}

		if body.BranchIndex < 0 {
			return handler.ErrOutOfRange
		} else if body.BranchIndex > len(parent.Children) {
			return handler.ErrOutOfRange
		}

		if body.BranchIndex == len(parent.Children) { // in case we are making a new branch
			parent.Children = append(parent.Children, []string{})
		}

		parent.Children[body.BranchIndex] = append(parent.Children[body.BranchIndex], step.Tag)

		err = stores.Steps.Update(db.Query{"tag": parent.Tag}, *parent)
		if err != nil {
			return err
		}
	} else {
		return handler.ErrInvalidParentPrefix
	}

	handler.WriteJSON(w, 200, step)
	return nil
}
