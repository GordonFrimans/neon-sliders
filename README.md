# 🎨 Neon Sliders for Fyne

Beautiful neon sliders with smooth animations for Go GUI applications using the Fyne framework.

## ✨ Features

- **6 Color Schemes**: Cyber Green, Electric Blue, Cyber Pink, Orange Fire, Purple Dream, Teal Wave
- **3 Animation Types**: Wave, Pulse, Breathing
- **Configurable Steps**: Precise value control with discrete steps
- **Interaction Modes**: Drag across full area or thumb-only dragging
- **Smooth Transitions**: 60 FPS animation with advanced smoothing functions


## 🚀 Quick Start

### Installation

```bash
go get github.com/GordonFrimans/neon-sliders
```


### Simple Example

```go
package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    neonslider "your-module/neon-sliders"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Neon Slider Demo")
    
    // Create simple green slider
    slider := neonslider.New(0, 100)
    slider.SetValue(50)
    
    content := container.NewVBox(slider)
    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}
```


## 📖 Documentation

### Creating Sliders

```go
// Basic slider
slider := neonslider.New(0, 100)

// With color scheme
slider := neonslider.NewWithColor(0, 100, neonslider.BlueElectric)

// With step values
slider := neonslider.NewWithStep(0, 100, 5) // Step 5

// Full configuration
slider := neonslider.NewWithSettings(0, 100, neonslider.PinkCyber, 
    neonslider.DragFullTrack, neonslider.AnimationWave)
```


### Color Schemes

- `GreenCyber` - Cyber Green
- `BlueElectric` - Electric Blue
- `PinkCyber` - Cyber Pink
- `OrangeFire` - Orange Fire
- `PurpleDream` - Purple Dream
- `TealWave` - Teal Wave


### Animation Types

- `AnimationWave` - Wave animation
- `AnimationPulse` - Pulse effect
- `AnimationBreathing` - Breathing effect


### Methods

```go
// Set value with step consideration
slider.SetValue(75.0)

// Get current value
value := slider.GetValue()

// Configure step
slider.SetStep(2.5)

// Change color scheme
slider.SetColors(neonslider.PurpleDream)

// Set animation type
slider.SetAnimationType(neonslider.AnimationPulse)

// Set drag mode
slider.SetDragMode(neonslider.DragThumbOnly)

// Handle value changes
slider.OnChanged = func(value float64) {
    fmt.Printf("Value changed to: %.2f\n", value)
}
```


## 🎮 Demo Application

Run the full demo to see all capabilities:

```bash
go run example/main.go
```

The demo includes:

- All 6 color schemes with different configurations
- Interactive animation type comparison
- Step functionality demonstration
- Real-time customization controls


## 🔧 Advanced Usage

### Custom Color Schemes

```go
customColors := neonslider.NeonColors{
    PrimaryR: 255, PrimaryG: 50, PrimaryB: 100,
    TrackR: 30, TrackG: 10, TrackB: 20,
    MinIntensity: 0.3, MaxIntensity: 1.0,
    AnimationSpeed: 1.0, GlowRadius: 12,
    WaveAmplitude: 0.8, WaveFrequency: 1.0,
    PulseStrength: 0.6,
    BreathingDepth: 0.5, BreathingSpeed: 0.8,
}

slider := neonslider.NewWithColor(0, 100, customColors)
```


### Step Examples

```go
// No restrictions (any decimal)
slider := neonslider.NewWithStep(0, 100, 0)

// Integer values only
slider := neonslider.NewWithStep(0, 100, 1)

// Multiples of 5
slider := neonslider.NewWithStep(0, 100, 5)

// Decimal steps
slider := neonslider.NewWithStep(0, 10, 0.1)
```


## 📋 Requirements

- Go 1.19+
- Fyne v2.0+



## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Fyne](https://fyne.io/) - Cross-platform GUI toolkit for Go



