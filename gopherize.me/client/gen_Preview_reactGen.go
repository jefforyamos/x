// Code generated by myitcv.io/react/cmd/reactGen. DO NOT EDIT.

package main

import "myitcv.io/react"

type PreviewElem struct {
	react.Element
}

func buildPreview(cd react.ComponentDef) react.Component {
	return PreviewDef{ComponentDef: cd}
}

func buildPreviewElem(props PreviewProps, children ...react.Element) *PreviewElem {
	return &PreviewElem{
		Element: react.CreateElement(buildPreview, props, children...),
	}
}

func (p PreviewDef) RendersElement() react.Element {
	return p.Render()
}

// IsProps is an auto-generated definition so that PreviewProps implements
// the myitcv.io/react.Props interface.
func (p PreviewProps) IsProps() {}

// Props is an auto-generated proxy to the current props of Preview
func (p PreviewDef) Props() PreviewProps {
	uprops := p.ComponentDef.Props()
	return uprops.(PreviewProps)
}

func (p PreviewProps) EqualsIntf(val react.Props) bool {
	return p == val.(PreviewProps)
}

var _ react.Props = PreviewProps{}
