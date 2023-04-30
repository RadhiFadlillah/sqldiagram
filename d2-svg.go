package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/go-shiori/dom"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

func renderD2Svg(codes string) ([]byte, error) {
	// Prepare text measurer
	ruler, err := textmeasure.NewRuler()
	if err != nil {
		return nil, fmt.Errorf("textmeasure error: %w", err)
	}

	// Prepare ELK layout
	layout := func(ctx context.Context, g *d2graph.Graph) error {
		return d2elklayout.Layout(ctx, g, &d2elklayout.ConfigurableOpts{
			Algorithm:       "layered",
			NodeSpacing:     20.0,
			Padding:         "[top=50,left=50,bottom=50,right=50]",
			EdgeNodeSpacing: 50.0,
			SelfLoopSpacing: 50.0,
		})
	}

	// Generate diagram
	diagram, _, err := d2lib.Compile(context.Background(), codes, &d2lib.CompileOptions{
		Ruler:  ruler,
		Layout: layout})
	if err != nil {
		return nil, fmt.Errorf("d2 compile error: %w", err)
	}

	// Convert diagram to SVG
	svg, err := d2svg.Render(diagram, &d2svg.RenderOpts{
		Pad:     d2svg.DEFAULT_PADDING,
		ThemeID: d2themescatalog.NeutralDefault.ID})
	if err != nil {
		return nil, fmt.Errorf("d2 svg error: %w", err)
	}

	// Stylize SVG for better clarity
	svg, err = stylizeD2Svg(svg)
	if err != nil {
		return nil, fmt.Errorf("d2 styling error: %w", err)
	}

	return svg, nil
}

func stylizeD2Svg(svg []byte) ([]byte, error) {
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
