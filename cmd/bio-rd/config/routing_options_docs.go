// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
// DO NOT EDIT: this file is automatically generated by docgen
package config

import (
	"github.com/projectdiscovery/yamldoc-go/encoder"
)

var RoutingOptionsDoc encoder.Doc

func init() {
	RoutingOptionsDoc.Type = "RoutingOptions"
	RoutingOptionsDoc.Comments[encoder.LineComment] = ""
	RoutingOptionsDoc.Description = ""
	RoutingOptionsDoc.Fields = make([]encoder.Doc, 3)
	RoutingOptionsDoc.Fields[0].Name = "static_routes"
	RoutingOptionsDoc.Fields[0].Type = "[]StaticRoute"
	RoutingOptionsDoc.Fields[0].Note = ""
	RoutingOptionsDoc.Fields[0].Description = "List of static routes to install in the RIB\nStill not implemented"
	RoutingOptionsDoc.Fields[0].Comments[encoder.LineComment] = "List of static routes to install in the RIB"
	RoutingOptionsDoc.Fields[1].Name = "router_id"
	RoutingOptionsDoc.Fields[1].Type = "string"
	RoutingOptionsDoc.Fields[1].Note = ""
	RoutingOptionsDoc.Fields[1].Description = "32-bit number to serve as router id. Must have the format x.x.x.x"
	RoutingOptionsDoc.Fields[1].Comments[encoder.LineComment] = "32-bit number to serve as router id. Must have the format x.x.x.x"
	RoutingOptionsDoc.Fields[2].Name = "autonomous_system"
	RoutingOptionsDoc.Fields[2].Type = "uint32"
	RoutingOptionsDoc.Fields[2].Note = ""
	RoutingOptionsDoc.Fields[2].Description = "32-bit autonomous system number"
	RoutingOptionsDoc.Fields[2].Comments[encoder.LineComment] = "32-bit autonomous system number"
}

func (_ RoutingOptions) Doc() *encoder.Doc {
	return &RoutingOptionsDoc
}

// Getrouting_optionsDoc returns documentation for the file cmd/bio-rd/config/routing_options_docs.go.
func Getrouting_optionsDoc() *encoder.FileDoc {
	return &encoder.FileDoc{
		Name:        "routing_options",
		Description: "",
		Structs: []*encoder.Doc{
			&RoutingOptionsDoc,
		},
	}
}
