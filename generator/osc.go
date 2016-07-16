package generator

import (
	"fmt"
	"math"
)

// Osc is an oscillator
type Osc struct {
	Shape     WaveType
	Amplitude float32
	DcOffset  float32
	Freq      float32
	// SampleRate
	Fs                int
	PhaseOffset       float32
	CurrentPhaseAngle float32
	phaseAngleIncr    float32
}

// NewOsc returns a new oscillator, note that if you change the phase offset of the returned osc,
// you also need to set the CurrentPhaseAngle
func NewOsc(shape WaveType, hz float32, fs int) *Osc {
	return &Osc{Shape: shape, Amplitude: 1, Freq: hz, Fs: fs, phaseAngleIncr: ((hz * TwoPi) / float32(fs))}
}

// Signal uses the osc to generate a discreet signal
func (o *Osc) Signal(length int) []float32 {
	output := make([]float32, length)
	for i := 0; i < length; i++ {
		output[i] = o.Sample()
	}
	return output
}

// Sample returns the next sample generated by the oscillator
func (o *Osc) Sample() (output float32) {
	if o == nil {
		return
	}

	if o.CurrentPhaseAngle < -math.Pi {
		o.CurrentPhaseAngle += TwoPi
	} else if o.CurrentPhaseAngle > math.Pi {
		o.CurrentPhaseAngle -= TwoPi
	}

	switch o.Shape {
	case WaveSine:
		output = o.Amplitude*Sine(o.CurrentPhaseAngle) + o.DcOffset
	case WaveTriangle:
		output = o.Amplitude*Triangle(o.CurrentPhaseAngle) + o.DcOffset
	case WaveSaw:
		output = o.Amplitude*Sawtooth(o.CurrentPhaseAngle) + o.DcOffset
	case WaveSqr:
		fmt.Println(o.CurrentPhaseAngle)
		output = o.Amplitude*Square(o.CurrentPhaseAngle) + o.DcOffset
	}

	o.CurrentPhaseAngle += o.phaseAngleIncr
	return output
}
