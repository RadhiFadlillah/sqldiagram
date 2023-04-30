package main

import (
	"bytes"
	"fmt"

	"github.com/go-shiori/dom"
	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2target"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
)

func renderScript(graph *d2graph.Graph) []byte {
	formatted := d2format.Format(graph.AST)
	return []byte(formatted)
}

func renderSvg(diagram *d2target.Diagram) ([]byte, error) {
	// Convert diagram to SVG
	svg, err := d2svg.Render(diagram, &d2svg.RenderOpts{
		Pad:     d2svg.DEFAULT_PADDING,
		ThemeID: d2themescatalog.NeutralDefault.ID})
	if err != nil {
		return nil, fmt.Errorf("d2 svg error: %w", err)
	}

	// Stylize SVG for better clarity
	svg, err = stylizeSvg(svg)
	if err != nil {
		return nil, fmt.Errorf("d2 styling error: %w", err)
	}

	return svg, nil
}

func stylizeSvg(svg []byte) ([]byte, error) {
	// Parse SVG as HTML document
	doc, err := dom.FastParse(bytes.NewReader(svg))
	if err != nil {
		return nil, err
	}

	// Make all SQL table header bold
	headers := dom.QuerySelectorAll(doc, ".class_header+text.text")
	for _, header := range headers {
		dom.SetAttribute(header, "class", "text-bold")
	}

	// Remove color from column names
	selector := ".class_header+text+text, line+text"
	columnNames := dom.QuerySelectorAll(doc, selector)
	for _, txtColumn := range columnNames {
		dom.SetAttribute(txtColumn, "class", "text")
	}

	// Adjust "important" columns
	selector = ".class_header ~ text + text.fill-N2 + text.fill-AA2:not(:empty)"
	importantTexts := dom.QuerySelectorAll(doc, selector)
	for _, txtKey := range importantTexts {
		// Get nodes for important texts
		txtType := dom.PreviousElementSibling(txtKey)
		txtColumn := dom.PreviousElementSibling(txtType)

		// Make all important texts to bold
		keyValue := dom.TextContent(txtKey)
		if keyValue == "PK" || keyValue == "FK" {
			dom.SetAttribute(txtColumn, "class", "text-bold")
			dom.SetAttribute(txtType, "class", "text-bold fill-N2")
			dom.SetAttribute(txtKey, "class", "text-bold")
		} else {
			dom.SetAttribute(txtKey, "class", "text")
		}

		// Make key color same as header
		header := dom.QuerySelector(txtKey.Parent, ".class_header")
		headerColor := dom.GetAttribute(header, "fill")
		dom.SetAttribute(txtColumn, "fill", headerColor)
		dom.SetAttribute(txtKey, "fill", headerColor)
	}

	// Return the SVG data
	svgNode := dom.QuerySelector(doc, "svg")
	svgRaw := dom.OuterHTML(svgNode)
	return []byte(svgRaw), nil
}
