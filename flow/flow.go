package flow

import (
	"encoding/json"
	"errors"
)

type Node struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Initialized bool   `json:"initialized"`
	Position    struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"position"`
	Data struct {
		Label           string `json:"label"`
		ToolbarVisible  bool   `json:"toolbarVisible"`
		ToolbarPosition string `json:"toolbarPosition"`
		Icon            string `json:"icon"`
		Color           string `json:"color"`
		Description     string `json:"description"`
		Subject         struct {
			ObjectType          string  `json:"object_type"`
			SubjectType         string  `json:"subject_type"`
			RoleID              *int    `json:"role_id"`
			UserID              *string `json:"user_id"`
			IsReadOnly          bool    `json:"is_read_only"`
			IsSignatureNeeded   bool    `json:"is_signature_needed"`
			IsSubjectChangeable bool    `json:"is_subject_changeable"`
			OrgRoleID           *int    `json:"org_role_id"`
			OrgID               *string `json:"org_id"`
		} `json:"subject"`
		Triggers []string `json:"triggers,omitempty"`
	} `json:"data"`
	Style struct {
		Width  string `json:"width,omitempty"`
		Height string `json:"height,omitempty"`
	} `json:"style,omitempty"`
	ParentNode   string `json:"parentNode,omitempty"`
	Extent       string `json:"extent,omitempty"`
	ExpandParent bool   `json:"expandParent,omitempty"`
	Draggable    bool   `json:"draggable,omitempty"`
}

type Edge struct {
	ID           string  `json:"id"`
	Type         string  `json:"type"`
	Source       string  `json:"source"`
	Target       string  `json:"target"`
	SourceHandle string  `json:"sourceHandle"`
	TargetHandle *string `json:"targetHandle"`
	Data         struct {
	} `json:"data"`
	Label     string `json:"label"`
	Animated  bool   `json:"animated"`
	MarkerEnd struct {
		Type string `json:"type"`
	} `json:"markerEnd"`
	SourceX float64 `json:"sourceX"`
	SourceY float64 `json:"sourceY"`
	TargetX float64 `json:"targetX"`
	TargetY float64 `json:"targetY"`
}

type Data struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Flow struct {
	Data        Data
	Starts      []Node
	Ends        []Node
	Actions     []Node
	Triggers    []Node
	Edges       []Edge
	WithVote    bool
	WithEndVote bool
	VoteIndex   int
}

func NewFlow(rawData []byte) (*Flow, error) {
	var data Data
	if err := json.Unmarshal(rawData, &data); err != nil {
		return nil, err
	}

	f := &Flow{
		Data:        data,
		Starts:      []Node{},
		Ends:        []Node{},
		Actions:     []Node{},
		Triggers:    []Node{},
		Edges:       data.Edges,
		WithVote:    false,
		WithEndVote: false,
		VoteIndex:   -1,
	}

	if err := f.initialize(); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *Flow) initialize() error {
	if len(f.Data.Nodes) == 0 || len(f.Data.Edges) == 0 {
		return errors.New("invalid data")
	}

	for i, node := range f.Data.Nodes {
		switch node.Type {
		case "start":
			f.Starts = append(f.Starts, node)
		case "end":
			f.Ends = append(f.Ends, node)
		case "action":
			f.Actions = append(f.Actions, node)
			if node.Data.Subject.ObjectType == "VOTE" {
				f.WithVote = true
				f.VoteIndex = i
			}
			if node.Data.Subject.ObjectType == "END_VOTE" {
				f.WithEndVote = true
			}
		case "trigger":
			f.Triggers = append(f.Triggers, node)
		default:
			// Optionally log warning
		}
	}

	return nil
}

func (f *Flow) GetNodeByID(id string) *Node {
	for _, node := range f.Data.Nodes {
		if node.ID == id {
			return &node
		}
	}
	return nil
}

func (f *Flow) GetNodeIndex(id string) int {
	for i, node := range f.Data.Nodes {
		if node.ID == id {
			return i
		}
	}
	return -1
}

type NextStep struct {
	Node
	SourcePortLabel *string `json:"sourcePortLabel"`
	Open            bool    `json:"open"`
}

func (f *Flow) GetNextSteps(node Node) []NextStep {
	var nextSteps []NextStep

	if node.Type == "action" {
		var triggerNodes []Node
		for _, t := range f.Triggers {
			if t.ParentNode == node.ID {
				triggerNodes = append(triggerNodes, t)
			}
		}

		for _, trigger := range triggerNodes {
			var outgoingEdge *Edge
			for _, edge := range f.Edges {
				if edge.Source == trigger.ID {
					outgoingEdge = &edge
					break
				}
			}

			if outgoingEdge != nil {
				targetNode := f.GetNodeByID(outgoingEdge.Target)
				if targetNode != nil {
					sourcePortLabel := trigger.Data.Label
					nextSteps = append(nextSteps, NextStep{
						Node:            *targetNode,
						SourcePortLabel: &sourcePortLabel,
						Open:            false,
					})
				}
			}
		}
	} else {
		var outgoingEdges []Edge
		for _, edge := range f.Edges {
			if edge.Source == node.ID {
				outgoingEdges = append(outgoingEdges, edge)
			}
		}

		for _, edge := range outgoingEdges {
			targetNode := f.GetNodeByID(edge.Target)
			if targetNode != nil {
				nextSteps = append(nextSteps, NextStep{
					Node:            *targetNode,
					SourcePortLabel: nil,
					Open:            false,
				})
			}
		}
	}

	return nextSteps
}
