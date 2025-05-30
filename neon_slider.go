// Package neonslider предоставляет красивый неоновый слайдер с плавной анимацией для Fyne приложений
package neonslider

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// SliderDragMode определяет режим взаимодействия со слайдером
type SliderDragMode int

const (
	// DragThumbOnly - перетаскивание только за ползунок
	DragThumbOnly SliderDragMode = iota
	// DragFullTrack - перетаскивание по всей области слайдера
	DragFullTrack
)

// String возвращает строковое представление режима перетаскивания
func (mode SliderDragMode) String() string {
	switch mode {
	case DragThumbOnly:
		return "Только ползунок"
	case DragFullTrack:
		return "Вся область"
	default:
		return "Неизвестный режим"
	}
}

// AnimationType определяет тип анимации слайдера
type AnimationType int

const (
	AnimationWave AnimationType = iota
	AnimationPulse
	AnimationBreathing
)

// String возвращает строковое представление типа анимации
func (anim AnimationType) String() string {
	switch anim {
	case AnimationWave:
		return "Волновая"
	case AnimationPulse:
		return "Пульсация"
	case AnimationBreathing:
		return "Дыхание"
	default:
		return "Неизвестная"
	}
}

// NeonColors определяет цветовую схему неонового слайдера
type NeonColors struct {
	// Основные RGB компоненты неонового цвета (0-255)
	PrimaryR, PrimaryG, PrimaryB uint8

	// RGB компоненты цвета трека (0-255)
	TrackR, TrackG, TrackB uint8

	// Параметры анимации - УВЕЛИЧЕНЫ для большей заметности
	MinIntensity   float64 // Минимальная яркость свечения (0.0-1.0)
	MaxIntensity   float64 // Максимальная яркость свечения (0.0-1.0)
	AnimationSpeed float64 // Скорость анимации (1.0 = нормальная)

	// Визуальные эффекты
	GlowRadius float32 // Радиус свечения в пикселях

	// Параметры для разных анимаций - УСИЛЕНЫ
	WaveAmplitude  float64 // Амплитуда волны (0.0-1.0)
	WaveFrequency  float64 // Частота волны
	PulseStrength  float64 // Сила пульсации (0.0-1.0)
	BreathingDepth float64 // Глубина "дыхания"
	BreathingSpeed float64 // Скорость дыхания
}

// ИСПРАВЛЕННЫЕ предустановленные цветовые схемы с усиленными анимациями
var (
	// GreenCyber - киберпанк зеленый с заметной анимацией
	GreenCyber = NeonColors{
		PrimaryR: 0, PrimaryG: 255, PrimaryB: 150,
		TrackR: 15, TrackG: 25, TrackB: 30,
		MinIntensity: 0.5, MaxIntensity: 1.0, // УВЕЛИЧЕН диапазон
		AnimationSpeed: 0.8, GlowRadius: 10,
		WaveAmplitude: 0.6, WaveFrequency: 1.0, // УСИЛЕНО
		PulseStrength:  0.5,                      // УСИЛЕНО
		BreathingDepth: 0.4, BreathingSpeed: 0.8, // УСИЛЕНО
	}

	// BlueElectric - электрический синий с яркой анимацией
	BlueElectric = NeonColors{
		PrimaryR: 0, PrimaryG: 150, PrimaryB: 255,
		TrackR: 10, TrackG: 15, TrackB: 35,
		MinIntensity: 0.4, MaxIntensity: 1.0, // УВЕЛИЧЕН диапазон
		AnimationSpeed: 0.9, GlowRadius: 12,
		WaveAmplitude: 0.7, WaveFrequency: 0.9, // УСИЛЕНО
		PulseStrength:  0.6,                      // УСИЛЕНО
		BreathingDepth: 0.5, BreathingSpeed: 0.7, // УСИЛЕНО
	}

	// PinkCyber - розовый киберпанк с энергичной анимацией
	PinkCyber = NeonColors{
		PrimaryR: 255, PrimaryG: 0, PrimaryB: 150,
		TrackR: 35, TrackG: 10, TrackB: 25,
		MinIntensity: 0.3, MaxIntensity: 1.0, // МАКСИМАЛЬНЫЙ диапазон
		AnimationSpeed: 1.2, GlowRadius: 14,
		WaveAmplitude: 0.8, WaveFrequency: 1.1, // МАКСИМУМ
		PulseStrength:  0.7,                      // МАКСИМУМ
		BreathingDepth: 0.6, BreathingSpeed: 0.9, // МАКСИМУМ
	}

	// OrangeFire - оранжевый огонь с мощной анимацией
	OrangeFire = NeonColors{
		PrimaryR: 255, PrimaryG: 100, PrimaryB: 0,
		TrackR: 35, TrackG: 20, TrackB: 5,
		MinIntensity: 0.4, MaxIntensity: 1.0, // УВЕЛИЧЕН диапазон
		AnimationSpeed: 1.0, GlowRadius: 12,
		WaveAmplitude: 0.6, WaveFrequency: 1.3, // УСИЛЕНО
		PulseStrength:  0.5,                      // УСИЛЕНО
		BreathingDepth: 0.4, BreathingSpeed: 1.0, // УСИЛЕНО
	}

	// PurpleDream - фиолетовая мечта с мистической анимацией
	PurpleDream = NeonColors{
		PrimaryR: 200, PrimaryG: 0, PrimaryB: 255,
		TrackR: 30, TrackG: 10, TrackB: 40,
		MinIntensity: 0.3, MaxIntensity: 1.0, // МАКСИМАЛЬНЫЙ диапазон
		AnimationSpeed: 0.8, GlowRadius: 15,
		WaveAmplitude: 0.8, WaveFrequency: 1.2, // УСИЛЕНО
		PulseStrength:  0.6,                      // УСИЛЕНО
		BreathingDepth: 0.7, BreathingSpeed: 0.5, // УСИЛЕНО
	}

	// TealWave - бирюзовая волна с плавной мощной анимацией
	TealWave = NeonColors{
		PrimaryR: 0, PrimaryG: 255, PrimaryB: 200,
		TrackR: 10, TrackG: 30, TrackB: 25,
		MinIntensity: 0.3, MaxIntensity: 1.0, // МАКСИМАЛЬНЫЙ диапазон
		AnimationSpeed: 1.1, GlowRadius: 16,
		WaveAmplitude: 0.9, WaveFrequency: 0.8, // МАКСИМУМ
		PulseStrength:  0.6,                      // УСИЛЕНО
		BreathingDepth: 0.8, BreathingSpeed: 0.6, // МАКСИМУМ
	}
)

// Easing функции для плавной анимации
func smoothstep(t float64) float64 {
	return t * t * (3.0 - 2.0*t)
}

func smootherstep(t float64) float64 {
	return t * t * t * (t*(t*6.0-15.0) + 10.0)
}

func easeInOutCubic(t float64) float64 {
	if t < 0.5 {
		return 4 * t * t * t
	}
	return 1 - math.Pow(-2*t+2, 3)/2
}

// ВОССТАНОВЛЕНА: roundToStep округляет значение к ближайшему шагу
func roundToStep(value, min, max, step float64) float64 {
	if step <= 0 {
		return value
	}

	// Ограничиваем значение в пределах min и max
	if value < min {
		value = min
	}
	if value > max {
		value = max
	}

	// Вычисляем количество шагов от min
	steps := math.Round((value - min) / step)

	// Возвращаем значение, округленное к ближайшему шагу
	result := min + steps*step

	// Дополнительная проверка границ после округления
	if result < min {
		result = min
	}
	if result > max {
		result = max
	}

	return result
}

// NeonSlider представляет неоновый слайдер с анимацией
type NeonSlider struct {
	widget.BaseWidget

	// Основные параметры слайдера
	Min, Max, Value float64       // Минимальное, максимальное и текущее значения
	Step            float64       // ВОССТАНОВЛЕНО: Шаг изменения значения (0 = без ограничений)
	OnChanged       func(float64) // Callback при изменении значения

	// Параметры анимации (внутренние)
	glowIntensity  float64   // Текущая интенсивность свечения
	pulsePhase     float64   // Фаза пульсации
	shimmerPhase   float64   // Фаза мерцания
	lastUpdateTime time.Time // Время последнего обновления

	// Состояние взаимодействия
	isDragging    bool           // Флаг перетаскивания
	DragMode      SliderDragMode // Режим перетаскивания
	AnimationType AnimationType  // Тип анимации

	// Визуальные настройки
	Colors NeonColors // Цветовая схема

	// Геометрия (внутренние параметры)
	thumbCenter fyne.Position       // Центр ползунка
	thumbSize   float32             // Размер ползунка
	renderer    *neonSliderRenderer // Рендерер
}

// New создает новый неоновый слайдер с базовыми настройками
func New(min, max float64) *NeonSlider {
	return NewWithColorAndMode(min, max, GreenCyber, DragFullTrack)
}

// ВОССТАНОВЛЕНО: NewWithStep создает слайдер с указанным шагом
func NewWithStep(min, max, step float64) *NeonSlider {
	return NewWithColorAndModeAndStep(min, max, step, GreenCyber, DragFullTrack, AnimationWave)
}

// NewWithColor создает слайдер с указанной цветовой схемой
func NewWithColor(min, max float64, colors NeonColors) *NeonSlider {
	return NewWithColorAndMode(min, max, colors, DragFullTrack)
}

// NewWithSettings создает слайдер с настройками анимации
func NewWithSettings(min, max float64, colors NeonColors,
	dragMode SliderDragMode, animType AnimationType) *NeonSlider {
	slider := NewWithColorAndMode(min, max, colors, dragMode)
	slider.AnimationType = animType
	return slider
}

// NewWithColorAndMode создает слайдер с полной настройкой (без шага)
func NewWithColorAndMode(min, max float64, colors NeonColors, dragMode SliderDragMode) *NeonSlider {
	return NewWithColorAndModeAndStep(min, max, 0, colors, dragMode, AnimationWave)
}

// ВОССТАНОВЛЕНО: NewWithColorAndModeAndStep создает слайдер с полной настройкой включая шаг
func NewWithColorAndModeAndStep(min, max, step float64, colors NeonColors,
	dragMode SliderDragMode, animType AnimationType) *NeonSlider {
	slider := &NeonSlider{
		Min:            min,
		Max:            max,
		Value:          min,
		Step:           step, // ВОССТАНОВЛЕНО: Устанавливаем шаг
		DragMode:       dragMode,
		Colors:         colors,
		AnimationType:  animType,
		thumbSize:      32,
		glowIntensity:  colors.MinIntensity,
		lastUpdateTime: time.Now(),
	}

	slider.ExtendBaseWidget(slider)
	return slider
}

// SetValue устанавливает значение слайдера с учетом шага
func (n *NeonSlider) SetValue(value float64) {
	// ВОССТАНОВЛЕНО: Применяем шаг при установке значения
	if n.Step > 0 {
		value = roundToStep(value, n.Min, n.Max, n.Step)
	} else {
		// Ограничиваем значение диапазоном без шага
		if value < n.Min {
			value = n.Min
		}
		if value > n.Max {
			value = n.Max
		}
	}

	// Обновляем значение и вызываем callback
	if n.Value != value {
		n.Value = value
		if n.OnChanged != nil {
			n.OnChanged(value)
		}
	}

	n.Refresh()
}

// GetValue возвращает текущее значение слайдера
func (n *NeonSlider) GetValue() float64 {
	return n.Value
}

// ВОССТАНОВЛЕНО: SetStep устанавливает шаг изменения значения
func (n *NeonSlider) SetStep(step float64) {
	if step < 0 {
		step = 0
	}
	n.Step = step

	// Перепроверяем текущее значение с новым шагом
	if step > 0 {
		n.SetValue(n.Value) // Это пересчитает значение с учетом нового шага
	}
}

// ВОССТАНОВЛЕНО: GetStep возвращает текущий шаг
func (n *NeonSlider) GetStep() float64 {
	return n.Step
}

// SetColors изменяет цветовую схему слайдера
func (n *NeonSlider) SetColors(colors NeonColors) {
	n.Colors = colors
	n.glowIntensity = colors.MinIntensity
	n.Refresh()
}

// SetDragMode изменяет режим перетаскивания
func (n *NeonSlider) SetDragMode(mode SliderDragMode) {
	n.DragMode = mode
}

// SetAnimationType изменяет тип анимации
func (n *NeonSlider) SetAnimationType(animType AnimationType) {
	n.AnimationType = animType
}

// StartAnimation запускает анимацию слайдера
func (n *NeonSlider) StartAnimation() {
	go func() {
		ticker := time.NewTicker(16 * time.Millisecond) // 60 FPS
		defer ticker.Stop()
		startTime := time.Now()

		for range ticker.C {
			elapsed := time.Since(startTime).Seconds()

			if elapsed > 86400 { // Сброс каждые 24 часа
				startTime = time.Now()
				elapsed = 0
			}

			fyne.Do(func() {
				n.updateSmoothGlow(elapsed)
				n.Refresh()
			})
		}
	}()
}

// КАРДИНАЛЬНО УЛУЧШЕННЫЕ методы анимации для большей заметности

// updateWaveAnimation - волновая анимация с многослойными эффектами
func (n *NeonSlider) updateWaveAnimation(elapsed float64) {
	baseTime := elapsed * n.Colors.AnimationSpeed * n.Colors.WaveFrequency

	// Основная волна
	mainWave := math.Sin(baseTime * 1.5)
	// Вторичная волна для сложности
	secondWave := math.Sin(baseTime*2.3) * 0.4
	// Быстрое мерцание для живости
	flicker := math.Sin(baseTime*8.0) * 0.15

	// Комбинируем все волны
	combinedWave := (mainWave + secondWave + flicker) * n.Colors.WaveAmplitude

	// Нормализуем и применяем easing
	normalizedWave := (combinedWave + 1.0) / 2.0
	smoothWave := smootherstep(normalizedWave)

	pulseRange := n.Colors.MaxIntensity - n.Colors.MinIntensity
	n.glowIntensity = n.Colors.MinIntensity + pulseRange*smoothWave

	// Обновляем дополнительные фазы для рендера
	n.pulsePhase = smoothWave * 0.4                               // Увеличено
	n.shimmerPhase = ((math.Sin(baseTime*5.5) + 1.0) / 2.0) * 0.3 // Увеличено
}

// updatePulseAnimation - пульсирующая анимация с резкими всплесками
func (n *NeonSlider) updatePulseAnimation(elapsed float64) {
	pulseTime := elapsed * n.Colors.AnimationSpeed * 2.5 // Ускорено

	// Основная пульсация
	rawPulse := math.Sin(pulseTime * math.Pi)
	normalizedPulse := (rawPulse + 1.0) / 2.0
	smoothPulse := easeInOutCubic(normalizedPulse)

	// Добавляем быстрые всплески
	burstTime := elapsed * n.Colors.AnimationSpeed * 7.0
	burst := math.Max(0, math.Sin(burstTime)) * 0.3

	pulseRange := n.Colors.MaxIntensity - n.Colors.MinIntensity
	basePulse := smoothPulse * n.Colors.PulseStrength

	n.glowIntensity = n.Colors.MinIntensity + pulseRange*(basePulse+burst)

	// Сильные эффекты для пульсации
	n.pulsePhase = (smoothPulse + burst) * 0.5 // Увеличено
	n.shimmerPhase = burst * 0.4               // Новый эффект
}

// updateBreathingAnimation - "дыхание" с медленными мощными переходами
func (n *NeonSlider) updateBreathingAnimation(elapsed float64) {
	breathTime := elapsed * n.Colors.BreathingSpeed

	// Медленное основное дыхание
	rawBreath := math.Sin(breathTime * math.Pi * 0.4) // Медленнее
	normalizedBreath := (rawBreath + 1.0) / 2.0
	// Тройное сглаживание для супер-плавности
	smoothBreath := smootherstep(smootherstep(smootherstep(normalizedBreath)))

	// Добавляем тонкое мерцание на пиках
	peakFlicker := 0.0
	if smoothBreath > 0.8 {
		flickerTime := elapsed * n.Colors.AnimationSpeed * 12.0
		peakFlicker = math.Sin(flickerTime) * 0.1 * (smoothBreath - 0.8) * 5.0
	}

	pulseRange := n.Colors.MaxIntensity - n.Colors.MinIntensity
	breathEffect := smoothBreath * n.Colors.BreathingDepth

	n.glowIntensity = n.Colors.MinIntensity + pulseRange*(breathEffect+peakFlicker)

	// Мягкие дополнительные эффекты
	n.pulsePhase = smoothBreath * 0.3  // Мягко
	n.shimmerPhase = peakFlicker * 0.5 // Только на пиках
}

// updateSmoothGlow - главная функция обновления анимации
func (n *NeonSlider) updateSmoothGlow(elapsed float64) {
	if elapsed < 0 || math.IsNaN(elapsed) || math.IsInf(elapsed, 0) {
		return
	}

	// Выбираем анимацию
	switch n.AnimationType {
	case AnimationWave:
		n.updateWaveAnimation(elapsed)
	case AnimationPulse:
		n.updatePulseAnimation(elapsed)
	case AnimationBreathing:
		n.updateBreathingAnimation(elapsed)
	default:
		n.updateWaveAnimation(elapsed) // По умолчанию волновая
	}

	// Гарантируем границы
	n.glowIntensity = math.Max(n.Colors.MinIntensity,
		math.Min(n.glowIntensity, n.Colors.MaxIntensity))

	// УСИЛЕННЫЙ эффект при перетаскивании
	if n.isDragging {
		dragBoost := (n.Colors.MaxIntensity - n.Colors.MinIntensity) * 0.3 // Увеличено
		maxPossible := n.Colors.MaxIntensity - n.glowIntensity
		if dragBoost > maxPossible {
			dragBoost = maxPossible
		}
		n.glowIntensity += dragBoost
		n.pulsePhase *= 1.5 // Усиливаем все эффекты
		n.shimmerPhase *= 1.5
	}
}

// isPointInThumb проверяет, находится ли точка внутри ползунка
func (n *NeonSlider) isPointInThumb(pos fyne.Position) bool {
	hitRadius := n.thumbSize/2 + 10
	dx := pos.X - n.thumbCenter.X
	dy := pos.Y - n.thumbCenter.Y
	distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))
	return distance <= hitRadius
}

// updateValueFromPosition обновляет значение слайдера на основе позиции мыши
func (n *NeonSlider) updateValueFromPosition(x float32) {
	size := n.Size()
	if size.Width == 0 {
		return
	}

	padding := n.thumbSize / 2
	usableWidth := size.Width - padding*2
	adjustedX := x - padding

	if adjustedX < 0 {
		adjustedX = 0
	}
	if adjustedX > usableWidth {
		adjustedX = usableWidth
	}

	if usableWidth == 0 {
		return
	}

	ratio := float64(adjustedX / usableWidth)
	newValue := n.Min + ratio*(n.Max-n.Min)

	// ВОССТАНОВЛЕНО: Применяем шаг при перетаскивании
	n.SetValue(newValue) // SetValue уже учитывает шаг
}

// Реализация интерфейсов взаимодействия
func (n *NeonSlider) Tapped(e *fyne.PointEvent) {
	switch n.DragMode {
	case DragThumbOnly:
		if n.isPointInThumb(e.Position) {
			n.updateValueFromPosition(e.Position.X)
		}
	case DragFullTrack:
		n.updateValueFromPosition(e.Position.X)
	}
}

func (n *NeonSlider) Dragged(e *fyne.DragEvent) {
	if !n.isDragging {
		startPos := fyne.NewPos(e.Position.X-e.Dragged.DX, e.Position.Y-e.Dragged.DY)
		switch n.DragMode {
		case DragThumbOnly:
			if n.isPointInThumb(startPos) {
				n.isDragging = true
			}
		case DragFullTrack:
			n.isDragging = true
		}
	}

	if n.isDragging {
		n.updateValueFromPosition(e.Position.X)
	}
}

func (n *NeonSlider) DragEnd() {
	n.isDragging = false
}

func (n *NeonSlider) MouseIn(*desktop.MouseEvent)      {}
func (n *NeonSlider) MouseOut()                        {}
func (n *NeonSlider) MouseMoved(e *desktop.MouseEvent) {}

// CreateRenderer создает рендерер для слайдера
func (n *NeonSlider) CreateRenderer() fyne.WidgetRenderer {
	track := canvas.NewRectangle(color.RGBA{R: n.Colors.TrackR, G: n.Colors.TrackG, B: n.Colors.TrackB, A: 255})
	fill := canvas.NewRectangle(color.RGBA{R: n.Colors.PrimaryR, G: n.Colors.PrimaryG, B: n.Colors.PrimaryB, A: 200})
	thumb := canvas.NewCircle(color.RGBA{R: n.Colors.PrimaryR, G: n.Colors.PrimaryG, B: n.Colors.PrimaryB, A: 255})

	track.CornerRadius = 10
	fill.CornerRadius = 10

	renderer := &neonSliderRenderer{
		slider: n,
		track:  track,
		fill:   fill,
		thumb:  thumb,
	}

	n.renderer = renderer
	n.StartAnimation()

	return renderer
}

// neonSliderRenderer отвечает за отрисовку слайдера
type neonSliderRenderer struct {
	slider *NeonSlider
	track  *canvas.Rectangle
	fill   *canvas.Rectangle
	thumb  *canvas.Circle
}

func (r *neonSliderRenderer) Layout(size fyne.Size) {
	if r.track == nil || r.fill == nil || r.thumb == nil {
		return
	}

	trackHeight := float32(20)
	thumbSize := r.slider.thumbSize
	padding := thumbSize / 2
	trackY := (size.Height - trackHeight) / 2

	r.track.Resize(fyne.NewSize(size.Width-padding*2, trackHeight))
	r.track.Move(fyne.NewPos(padding, trackY))

	fillRatio := (r.slider.Value - r.slider.Min) / (r.slider.Max - r.slider.Min)
	if math.IsNaN(fillRatio) || math.IsInf(fillRatio, 0) {
		fillRatio = 0
	}
	fillWidth := float32(fillRatio) * (size.Width - padding*2)

	r.fill.Resize(fyne.NewSize(fillWidth, trackHeight))
	r.fill.Move(fyne.NewPos(padding, trackY))

	thumbX := padding + fillWidth
	thumbY := size.Height / 2
	r.slider.thumbCenter = fyne.NewPos(thumbX, thumbY)

	r.thumb.Resize(fyne.NewSize(thumbSize, thumbSize))
	r.thumb.Move(fyne.NewPos(thumbX-thumbSize/2, thumbY-thumbSize/2))
}

func (r *neonSliderRenderer) MinSize() fyne.Size {
	return fyne.NewSize(250, 80)
}

// УЛУЧШЕННЫЙ Refresh с более яркими и заметными эффектами
func (r *neonSliderRenderer) Refresh() {
	if r.track == nil || r.fill == nil || r.thumb == nil {
		return
	}

	colors := &r.slider.Colors
	intensity := r.slider.glowIntensity
	pulse := r.slider.pulsePhase
	shimmer := r.slider.shimmerPhase

	if intensity < colors.MinIntensity {
		intensity = colors.MinIntensity
	}

	// УСИЛЕННОЕ свечение дорожки
	trackGlow := colors.MinIntensity*0.8 + intensity*0.2 // Больше базового свечения

	r.track.FillColor = color.RGBA{
		R: colors.TrackR,
		G: colors.TrackG,
		B: colors.TrackB,
		A: 255,
	}

	r.track.StrokeColor = color.RGBA{
		R: uint8(float64(colors.PrimaryR) * trackGlow),
		G: uint8(float64(colors.PrimaryG) * trackGlow),
		B: uint8(float64(colors.PrimaryB) * trackGlow),
		A: uint8(100 + trackGlow*155), // Увеличена базовая прозрачность
	}
	r.track.StrokeWidth = float32(2.0 + trackGlow*2.0) // Увеличена толщина

	// МАКСИМАЛЬНО заметная заливка
	fillBrightness := intensity + pulse*0.3 + shimmer*0.2 // УВЕЛИЧЕНЫ коэффициенты
	if fillBrightness < colors.MinIntensity {
		fillBrightness = colors.MinIntensity
	}
	if fillBrightness > colors.MaxIntensity {
		fillBrightness = colors.MaxIntensity
	}

	fillAlpha := uint8(200 + fillBrightness*55) // Увеличена базовая непрозрачность

	r.fill.FillColor = color.RGBA{
		R: uint8(math.Min(255, float64(colors.PrimaryR)*(0.7+fillBrightness*0.3))), // Увеличен диапазон
		G: uint8(math.Min(255, float64(colors.PrimaryG)*(0.7+fillBrightness*0.3))),
		B: uint8(math.Min(255, float64(colors.PrimaryB)*(0.7+fillBrightness*0.3))),
		A: fillAlpha,
	}

	// МОЩНОЕ свечение заливки
	glowIntensity := fillBrightness + pulse*0.4 + shimmer*0.3 // МАКСИМАЛЬНЫЕ эффекты
	if glowIntensity > colors.MaxIntensity {
		glowIntensity = colors.MaxIntensity
	}

	r.fill.StrokeColor = color.RGBA{
		R: uint8(math.Min(255, float64(colors.PrimaryR)*(0.8+glowIntensity*0.5))), // Ярче
		G: uint8(math.Min(255, float64(colors.PrimaryG)*(0.8+glowIntensity*0.5))),
		B: uint8(math.Min(255, float64(colors.PrimaryB)*(0.8+glowIntensity*0.5))),
		A: uint8(150 + glowIntensity*105), // Максимальная видимость
	}
	r.fill.StrokeWidth = colors.GlowRadius * float32(1.0+fillBrightness*0.8) // Увеличен радиус

	// СУПЕР-ЯРКИЙ ползунок
	thumbBrightness := fillBrightness + pulse*0.5 + shimmer*0.4 // МАКСИМУМ
	if thumbBrightness < colors.MinIntensity {
		thumbBrightness = colors.MinIntensity
	}
	if thumbBrightness > colors.MaxIntensity {
		thumbBrightness = colors.MaxIntensity
	}

	// Ядро ползунка - максимально яркое
	thumbCore := color.RGBA{
		R: uint8(math.Min(255, float64(colors.PrimaryR)*(0.8+thumbBrightness*0.2))),
		G: uint8(math.Min(255, float64(colors.PrimaryG)*(0.8+thumbBrightness*0.2))),
		B: uint8(math.Min(255, float64(colors.PrimaryB)*(0.8+thumbBrightness*0.2))),
		A: 255,
	}

	// Свечение ползунка - максимальная видимость
	thumbGlow := color.RGBA{
		R: uint8(math.Min(255, float64(colors.PrimaryR)*(1.0+thumbBrightness*0.5))),
		G: uint8(math.Min(255, float64(colors.PrimaryG)*(1.0+thumbBrightness*0.5))),
		B: uint8(math.Min(255, float64(colors.PrimaryB)*(1.0+thumbBrightness*0.5))),
		A: uint8(180 + thumbBrightness*75), // Очень высокая видимость
	}

	r.thumb.FillColor = thumbCore
	r.thumb.StrokeColor = thumbGlow
	r.thumb.StrokeWidth = colors.GlowRadius * float32(1.2+thumbBrightness*0.8) // Максимальное свечение

	// Принудительное обновление
	canvas.Refresh(r.track)
	canvas.Refresh(r.fill)
	canvas.Refresh(r.thumb)
	r.Layout(r.slider.Size())
}

func (r *neonSliderRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.track, r.fill, r.thumb}
}

func (r *neonSliderRenderer) Destroy() {}
