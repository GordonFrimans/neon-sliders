// Full-featured test application for neon sliders
package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	neonslider "neonslider" // Replace with your module
)

func main() {
	myApp := app.NewWithID("com.neonslider.demo")
	myApp.Settings().SetTheme(&darkTheme{})

	myWindow := myApp.NewWindow("🎨 Neon Sliders - Full Demo")
	myWindow.Resize(fyne.NewSize(1300, 900))
	myWindow.CenterOnScreen()

	content := createFullDemo()
	myWindow.SetContent(container.NewScroll(content))
	myWindow.ShowAndRun()
}

func createFullDemo() *fyne.Container {
	// === SECTION 1: Basic color schemes ===
	basicSection := createBasicColorsSection()

	// === SECTION 2: Animation demo ===
	animationSection := createAnimationDemoSection()

	// === SECTION 3: Step sliders ===
	stepSection := createStepDemoSection()

	// === SECTION 4: Interactive customization ===
	customSection := createCustomizationSection()

	return container.NewVBox(
		widget.NewLabelWithStyle("🎨 Neon Sliders - Full Demo",
			fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabel("All features: colors, animations, Step, interactive settings"),
		widget.NewSeparator(),

		basicSection,
		widget.NewSeparator(),
		animationSection,
		widget.NewSeparator(),
		stepSection,
		widget.NewSeparator(),
		customSection,
	)
}

// Demo of all color schemes
func createBasicColorsSection() *widget.Card {
	// Create sliders with preset colors
	greenSlider := neonslider.NewWithColor(0, 100, neonslider.GreenCyber)
	greenSlider.SetValue(25)

	blueSlider := neonslider.NewWithColorAndMode(0, 100, neonslider.BlueElectric, neonslider.DragThumbOnly)
	blueSlider.SetValue(45)

	pinkSlider := neonslider.NewWithColor(0, 100, neonslider.PinkCyber)
	pinkSlider.SetValue(65)

	orangeSlider := neonslider.NewWithColor(0, 100, neonslider.OrangeFire)
	orangeSlider.SetValue(85)

	purpleSlider := neonslider.NewWithColor(0, 100, neonslider.PurpleDream)
	purpleSlider.SetValue(35)

	tealSlider := neonslider.NewWithColor(0, 100, neonslider.TealWave)
	tealSlider.SetValue(75)

	// Value labels
	greenLabel := widget.NewLabelWithStyle("25%", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	blueLabel := widget.NewLabelWithStyle("45%", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	pinkLabel := widget.NewLabelWithStyle("65%", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	orangeLabel := widget.NewLabelWithStyle("85%", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	purpleLabel := widget.NewLabelWithStyle("35%", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	tealLabel := widget.NewLabelWithStyle("75%", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Bind handlers
	greenSlider.OnChanged = func(v float64) { greenLabel.SetText(fmt.Sprintf("%.1f%%", v)) }
	blueSlider.OnChanged = func(v float64) { blueLabel.SetText(fmt.Sprintf("%.1f%%", v)) }
	pinkSlider.OnChanged = func(v float64) { pinkLabel.SetText(fmt.Sprintf("%.1f%%", v)) }
	orangeSlider.OnChanged = func(v float64) { orangeLabel.SetText(fmt.Sprintf("%.1f%%", v)) }
	purpleSlider.OnChanged = func(v float64) { purpleLabel.SetText(fmt.Sprintf("%.1f%%", v)) }
	tealSlider.OnChanged = func(v float64) { tealLabel.SetText(fmt.Sprintf("%.1f%%", v)) }

	content := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 🎨 All Color Schemes"),

		container.NewGridWithColumns(2,
			// Left column
			container.NewVBox(
				widget.NewLabelWithStyle("🟢 Cyber Green", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel("Stable bright glow | Full area"),
				greenLabel, greenSlider,

				widget.NewSeparator(),

				widget.NewLabelWithStyle("🟣 Pink Cyber", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel("Energetic animation | Full area"),
				pinkLabel, pinkSlider,

				widget.NewSeparator(),

				widget.NewLabelWithStyle("💜 Purple Dream", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel("Mystical glow | Full area"),
				purpleLabel, purpleSlider,
			),

			// Right column
			container.NewVBox(
				widget.NewLabelWithStyle("🔵 Electric Blue", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel("Smooth animation | Thumb only"),
				blueLabel, blueSlider,

				widget.NewSeparator(),

				widget.NewLabelWithStyle("🟠 Orange Fire", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel("Powerful glow | Full area"),
				orangeLabel, orangeSlider,

				widget.NewSeparator(),

				widget.NewLabelWithStyle("💎 Teal Wave", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel("Smooth waves | Full area"),
				tealLabel, tealSlider,
			),
		),
	)

	return widget.NewCard("🎨 Color Schemes", "All preset colors with enhanced animation", content)
}

// Demo of animation types
func createAnimationDemoSection() *widget.Card {
	// Create sliders with different animations
	waveSlider := neonslider.NewWithSettings(0, 100, neonslider.PurpleDream,
		neonslider.DragFullTrack, neonslider.AnimationWave)
	waveSlider.SetValue(40)

	pulseSlider := neonslider.NewWithSettings(0, 100, neonslider.TealWave,
		neonslider.DragFullTrack, neonslider.AnimationPulse)
	pulseSlider.SetValue(60)

	breathSlider := neonslider.NewWithSettings(0, 100, neonslider.BlueElectric,
		neonslider.DragFullTrack, neonslider.AnimationBreathing)
	breathSlider.SetValue(80)

	// Labels
	waveLabel := widget.NewLabelWithStyle("40%", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	pulseLabel := widget.NewLabelWithStyle("60%", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	breathLabel := widget.NewLabelWithStyle("80%", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Handlers
	waveSlider.OnChanged = func(v float64) { waveLabel.SetText(fmt.Sprintf("%.1f%%", v)) }
	pulseSlider.OnChanged = func(v float64) { pulseLabel.SetText(fmt.Sprintf("%.1f%%", v)) }
	breathSlider.OnChanged = func(v float64) { breathLabel.SetText(fmt.Sprintf("%.1f%%", v)) }

	content := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 🎭 Animation Types"),
		widget.NewLabel("Each type creates a unique visual effect"),

		container.NewVBox(
			widget.NewLabelWithStyle("🌊 Wave Animation", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			widget.NewLabel("Multi-layer waves with fast shimmer | Purple Dream"),
			waveLabel, waveSlider,
		),

		widget.NewSeparator(),

		container.NewVBox(
			widget.NewLabelWithStyle("💓 Pulse Animation", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			widget.NewLabel("Rhythmic oscillations with sharp bursts | Teal Wave"),
			pulseLabel, pulseSlider,
		),

		widget.NewSeparator(),

		container.NewVBox(
			widget.NewLabelWithStyle("🫁 Breathing", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			widget.NewLabel("Slow smooth transitions + peak shimmer | Electric Blue"),
			breathLabel, breathSlider,
		),
	)

	return widget.NewCard("🎬 Animations", "Demo of all animation effect types", content)
}

// Demo of Step functionality
func createStepDemoSection() *widget.Card {
	// Sliders with different steps
	noStepSlider := neonslider.NewWithStep(0, 100, 0) // No limits
	noStepSlider.SetColors(neonslider.GreenCyber)
	noStepSlider.SetValue(33.7)

	step1Slider := neonslider.NewWithStep(0, 100, 1) // Integers
	step1Slider.SetColors(neonslider.BlueElectric)
	step1Slider.SetAnimationType(neonslider.AnimationPulse)
	step1Slider.SetValue(47.3) // Will be rounded to 47

	step5Slider := neonslider.NewWithStep(0, 100, 5) // Step 5
	step5Slider.SetColors(neonslider.PinkCyber)
	step5Slider.SetAnimationType(neonslider.AnimationWave)
	step5Slider.SetValue(23) // Will be rounded to 25

	step10Slider := neonslider.NewWithStep(0, 100, 10) // Step 10
	step10Slider.SetColors(neonslider.OrangeFire)
	step10Slider.SetAnimationType(neonslider.AnimationBreathing)
	step10Slider.SetValue(37) // Will be rounded to 40

	// Labels
	noStepLabel := widget.NewLabelWithStyle("33.7", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	step1Label := widget.NewLabelWithStyle("47", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	step5Label := widget.NewLabelWithStyle("25", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	step10Label := widget.NewLabelWithStyle("40", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Handlers
	noStepSlider.OnChanged = func(v float64) { noStepLabel.SetText(fmt.Sprintf("%.2f", v)) }
	step1Slider.OnChanged = func(v float64) { step1Label.SetText(fmt.Sprintf("%.0f", v)) }
	step5Slider.OnChanged = func(v float64) { step5Label.SetText(fmt.Sprintf("%.0f", v)) }
	step10Slider.OnChanged = func(v float64) { step10Label.SetText(fmt.Sprintf("%.0f", v)) }

	content := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 📏 Step Demo"),
		widget.NewLabel("Try dragging sliders and see value rounding"),

		container.NewGridWithColumns(2,
			// Left column
			container.NewVBox(
				widget.NewLabelWithStyle("🎯 No limits (Step: 0)", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel("Any decimal values | Green"),
				noStepLabel, noStepSlider,

				widget.NewSeparator(),

				widget.NewLabelWithStyle("📊 Step 5", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel("0, 5, 10, 15, 20... | Pink + Waves"),
				step5Label, step5Slider,
			),

			// Right column
			container.NewVBox(
				widget.NewLabelWithStyle("🔢 Integers (Step: 1)", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel("1, 2, 3, 4, 5... | Blue + Pulse"),
				step1Label, step1Slider,

				widget.NewSeparator(),

				widget.NewLabelWithStyle("📈 Step 10", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel("0, 10, 20, 30... | Orange + Breathing"),
				step10Label, step10Slider,
			),
		),
	)

	return widget.NewCard("🎯 Step Control", "Discrete steps for precise value control", content)
}

// Interactive customization section
func createCustomizationSection() *widget.Card {
	// Create customizable slider
	customSlider := neonslider.NewWithStep(0, 100, 0)
	customSlider.SetValue(50)

	valueLabel := widget.NewLabelWithStyle("50.0", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	infoLabel := widget.NewLabelWithStyle("Step: 0 | Color: Green | Animation: Wave",
		fyne.TextAlignCenter, fyne.TextStyle{})

	// Declare selector variables beforehand
	var colorSelect *widget.Select
	var animSelect *widget.Select

	// Function to update info label
	updateInfoLabel := func() {
		stepText := "0"
		if customSlider.GetStep() > 0 {
			stepText = fmt.Sprintf("%.1f", customSlider.GetStep())
		}

		colorText := "Green"
		if colorSelect != nil && colorSelect.Selected != "" {
			switch colorSelect.Selected {
			case "🟢 Cyber Green":
				colorText = "Green"
			case "🔵 Electric Blue":
				colorText = "Blue"
			case "🟣 Pink Cyber":
				colorText = "Pink"
			case "🟠 Orange Fire":
				colorText = "Orange"
			case "💜 Purple Dream":
				colorText = "Purple"
			case "💎 Teal Wave":
				colorText = "Teal"
			}
		}

		animText := "Wave"
		if animSelect != nil && animSelect.Selected != "" {
			switch animSelect.Selected {
			case "🌊 Wave":
				animText = "Wave"
			case "💓 Pulse":
				animText = "Pulse"
			case "🫁 Breathing":
				animText = "Breathing"
			}
		}

		infoLabel.SetText(fmt.Sprintf("Step: %s | Color: %s | Animation: %s",
			stepText, colorText, animText))
	}

	customSlider.OnChanged = func(v float64) {
		if customSlider.GetStep() > 0 {
			valueLabel.SetText(fmt.Sprintf("%.0f", v))
		} else {
			valueLabel.SetText(fmt.Sprintf("%.2f", v))
		}
	}

	// Color scheme selector
	colorSelect = widget.NewSelect([]string{
		"🟢 Cyber Green",
		"🔵 Electric Blue",
		"🟣 Pink Cyber",
		"🟠 Orange Fire",
		"💜 Purple Dream",
		"💎 Teal Wave",
	}, func(selected string) {
		var colors neonslider.NeonColors
		switch selected {
		case "🟢 Cyber Green":
			colors = neonslider.GreenCyber
		case "🔵 Electric Blue":
			colors = neonslider.BlueElectric
		case "🟣 Pink Cyber":
			colors = neonslider.PinkCyber
		case "🟠 Orange Fire":
			colors = neonslider.OrangeFire
		case "💜 Purple Dream":
			colors = neonslider.PurpleDream
		case "💎 Teal Wave":
			colors = neonslider.TealWave
		default:
			colors = neonslider.GreenCyber
		}
		customSlider.SetColors(colors)
		updateInfoLabel()
	})
	colorSelect.SetSelected("🟢 Cyber Green")

	// Animation selector
	animSelect = widget.NewSelect([]string{
		"🌊 Wave",
		"💓 Pulse",
		"🫁 Breathing",
	}, func(selected string) {
		switch selected {
		case "🌊 Wave":
			customSlider.SetAnimationType(neonslider.AnimationWave)
		case "💓 Pulse":
			customSlider.SetAnimationType(neonslider.AnimationPulse)
		case "🫁 Breathing":
			customSlider.SetAnimationType(neonslider.AnimationBreathing)
		}
		updateInfoLabel()
	})
	animSelect.SetSelected("🌊 Wave")

	// Step selector
	stepSelect := widget.NewSelect([]string{
		"0 (no limits)",
		"0.1 (tenths)",
		"0.5 (halves)",
		"1 (integers)",
		"2 (even)",
		"5 (fives)",
		"10 (tens)",
		"25 (quarters)",
	}, func(selected string) {
		var step float64
		switch selected {
		case "0 (no limits)":
			step = 0
		case "0.1 (tenths)":
			step = 0.1
		case "0.5 (halves)":
			step = 0.5
		case "1 (integers)":
			step = 1
		case "2 (even)":
			step = 2
		case "5 (fives)":
			step = 5
		case "10 (tens)":
			step = 10
		case "25 (quarters)":
			step = 25
		default:
			step = 0
		}
		customSlider.SetStep(step)
		updateInfoLabel()
	})
	stepSelect.SetSelected("0 (no limits)")

	// Interaction mode selector
	dragSelect := widget.NewSelect([]string{
		"📏 Full area",
		"🎯 Thumb only",
	}, func(selected string) {
		switch selected {
		case "📏 Full area":
			customSlider.SetDragMode(neonslider.DragFullTrack)
		case "🎯 Thumb only":
			customSlider.SetDragMode(neonslider.DragThumbOnly)
		}
		updateInfoLabel()
	})
	dragSelect.SetSelected("📏 Full area")

	// Test buttons
	testButtons := container.NewGridWithColumns(4,
		widget.NewButton("13.7", func() { customSlider.SetValue(13.7) }),
		widget.NewButton("33.3", func() { customSlider.SetValue(33.3) }),
		widget.NewButton("66.6", func() { customSlider.SetValue(66.6) }),
		widget.NewButton("87.2", func() { customSlider.SetValue(87.2) }),
	)

	content := container.NewVBox(
		widget.NewRichTextFromMarkdown("### ⚙️ Interactive Settings"),
		widget.NewLabel("Configure all slider parameters in real time"),

		container.NewGridWithColumns(2,
			// Left column
			container.NewVBox(
				widget.NewLabel("Color scheme:"),
				colorSelect,
				widget.NewLabel("Animation type:"),
				animSelect,
			),

			// Right column
			container.NewVBox(
				widget.NewLabel("Step (increment):"),
				stepSelect,
				widget.NewLabel("Interaction mode:"),
				dragSelect,
			),
		),

		widget.NewSeparator(),

		widget.NewLabel("Test values:"),
		testButtons,

		widget.NewSeparator(),

		infoLabel,
		valueLabel,
		customSlider,

		widget.NewLabel("💡 Change settings above and see how the slider changes"),
	)

	return widget.NewCard("🔧 Settings", "Full control over all parameters", content)
}

// Enhanced dark theme for neon effects
type darkTheme struct{}

func (d *darkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 5, G: 8, B: 12, A: 255} // Very dark background
	case theme.ColorNameButton:
		return color.RGBA{R: 20, G: 30, B: 40, A: 255}
	case theme.ColorNamePrimary:
		return color.RGBA{R: 0, G: 255, B: 150, A: 255}
	case theme.ColorNameFocus:
		return color.RGBA{R: 0, G: 255, B: 200, A: 120}
	case theme.ColorNameHover:
		return color.RGBA{R: 30, G: 40, B: 50, A: 255}
	case theme.ColorNameInputBackground:
		return color.RGBA{R: 15, G: 20, B: 25, A: 255}
	case theme.ColorNamePlaceHolder:
		return color.RGBA{R: 100, G: 120, B: 140, A: 255}
	case theme.ColorNamePressed:
		return color.RGBA{R: 0, G: 200, B: 120, A: 255}
	case theme.ColorNameSelection:
		return color.RGBA{R: 0, G: 150, B: 255, A: 80}
	case theme.ColorNameSeparator:
		return color.RGBA{R: 40, G: 50, B: 60, A: 255}
	case theme.ColorNameShadow:
		return color.RGBA{R: 0, G: 0, B: 0, A: 120}
	case theme.ColorNameMenuBackground:
		return color.RGBA{R: 10, G: 15, B: 20, A: 255}
	case theme.ColorNameOverlayBackground:
		return color.RGBA{R: 0, G: 0, B: 0, A: 150}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (d *darkTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (d *darkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (d *darkTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 13
	case theme.SizeNameCaptionText:
		return 11
	case theme.SizeNameHeadingText:
		return 22
	case theme.SizeNameSubHeadingText:
		return 17
	default:
		return theme.DefaultTheme().Size(name)
	}
}
