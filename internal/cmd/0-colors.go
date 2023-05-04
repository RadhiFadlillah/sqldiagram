package cmd

var bgColors = []string{
	"#fef2f2", // red 50
	"#fff7ed", // orange 50
	"#fffbeb", // amber 50
	"#f7fee7", // lime 50
	"#f0fdf4", // green 50
	"#f0fdfa", // teal 50
	"#f0f9ff", // sky 50
	"#eff6ff", // blue 50
	"#eef2ff", // indigo 50
	"#f5f3ff", // violet 50
	"#fdf4ff", // fuchsia 50
	"#fdf2f8", // pink 50
}

var fgColors = []string{
	"#b91c1c", // red 700
	"#c2410c", // orange 700
	"#b45309", // amber 700
	"#4d7c0f", // lime 700
	"#15803d", // green 700
	"#0f766e", // teal 700
	"#0369a1", // sky 700
	"#1d4ed8", // blue 700
	"#4338ca", // indigo 700
	"#6d28d9", // violet 700
	"#a21caf", // fuchsia 700
	"#be185d", // pink 700
}

func getBgColor(idx int) string {
	return getColor(idx, bgColors)
}

func getFgColor(idx int) string {
	return getColor(idx, fgColors)
}

func getColor(idx int, colors []string) string {
	nColor := len(colors)

	for idx < 0 {
		idx += nColor
	}

	for idx >= nColor {
		idx -= nColor
	}

	return colors[idx]
}
