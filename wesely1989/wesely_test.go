package wesely1989

import (
	"testing"
	"math"
	"fmt"
)

// Results from Wesely (1989) table 3
var (
	SO2 = [][]float64{
		{120, 130, 160, 370, 1000, 90, 1100},
		{1300, 1300, 1300, 1300, 1400, 90, 1200},
		{1100, 1100, 1100, 1100, 1100, 90, 1000},
		{980, 980, 980, 990, 1000, 1000, 1000},
		{260, 280, 330, 620, 1100, 90, 1000}}

	O3 = [][]float64{
		{100, 110, 130, 320, 950, 950, 580},
		{430, 460, 520, 700, 1200, 940, 570},
		{390, 420, 460, 600, 950, 760, 500},
		{550, 610, 700, 1000, 3000, 3000, 3000},
		{180, 190, 220, 430, 940, 810, 520}}

	NO2 = [][]float64{
		{120, 130, 160, 480, 2800, 2700, 2300},
		{1800, 1800, 1800, 1900, 2600, 2400, 2100},
		{1700, 1700, 1700, 1800, 2300, 2200, 1900},
		{3800, 3800, 3900, 4300, 9500, 9500, 9500},
		{270, 290, 350, 850, 2400, 2300, 2000}}

	H2O2 = [][]float64{
		{80, 90, 110, 250, 640, 80, 80},
		{400, 430, 480, 640, 1000, 90, 80},
		{360, 380, 420, 540, 830, 80, 80},
		{390, 420, 460, 610, 980, 980, 980},
		{150, 170, 190, 370, 740, 80, 80}}

	ALD = [][]float64{
		{320, 330, 370, 790, 9999, 9999, 9999},
		{9999, 9999, 9999, 9999, 9999, 9999, 9999},
		{9999, 9999, 9999, 9999, 9999, 9999, 9999},
		{9999, 9999, 9999, 9999, 9999, 9999, 9999},
		{510, 540, 630, 1600, 9999, 9999, 9999}}

	HCHO = [][]float64{
		{100, 110, 130, 450, 6600, 1300, 1300},
		{8100, 8100, 8100, 8100, 8100, 1300, 1300},
		{7800, 7800, 7800, 7800, 7800, 1300, 1300},
		{2800, 2800, 2800, 2800, 2800, 2800, 2800},
		{250, 270, 330, 1000, 7500, 1300, 1300}}

	OP = [][]float64{
		{120, 130, 160, 480, 2800, 2500, 2100},
		{1800, 1800, 1800, 1900, 2600, 2300, 2000},
		{1700, 1700, 1700, 1800, 2300, 2100, 1800},
		{3500, 3600, 3600, 4000, 8000, 8000, 8000},
		{270, 290, 350, 840, 2400, 2100, 1900}}

	PAA = [][]float64{
		{140, 160, 190, 570, 2800, 2300, 2000},
		{1800, 1800, 1800, 1900, 2600, 2100, 1800},
		{1700, 1700, 1700, 1800, 2300, 1900, 1700},
		{3300, 3300, 3400, 3700, 6800, 6800, 6800},
		{320, 350, 420, 950, 2400, 2000, 1800}}

	FORM = [][]float64{
		{30, 30, 30, 40, 40, 0, 0},
		{130, 140, 150, 160, 190, 0, 0},
		{130, 130, 140, 160, 180, 0, 0},
		{270, 300, 330, 450, 660, 660, 660},
		{50, 60, 60, 80, 90, 0, 0}}

	NH3 = [][]float64{
		{70, 80, 100, 310, 2600, 430, 430},
		{3200, 3200, 3200, 3200, 3300, 430, 430},
		{2900, 2900, 2900, 2900, 2900, 430, 430},
		{1500, 1500, 1500, 1500, 1500, 1500, 1500},
		{180, 190, 230, 670, 2800, 430, 430}}
	PAN = [][]float64{
		{190, 200, 250, 700, 2800, 2700, 2300},
		{1800, 1800, 1800, 1900, 2600, 2400, 2100},
		{1700, 1700, 1700, 1800, 2400, 2200, 1900},
		{3800, 3900, 4000, 4400, 9700, 9700, 9700},
		{400, 430, 510, 1000, 2400, 2300, 2000}}
	HNO3 = [][]float64{
		{100, 110, 140, 340, 1000, 90, 90},
		{1300, 1300, 1300, 1300, 1300, 1400, 90, 90},
		{1100, 1100, 1100, 1100, 1100, 90, 90},
		{980, 980, 980, 990, 1000, 1000, 1000},
		{230, 250, 290, 580, 1100, 90, 90}}
)

func TestWesely(t *testing.T) {
	const iLandUse = 3                       // deciduous forest
	Ts := []float64{25, 10, 2, 0, 10}        // Surface Temperature [C]
	Garr := []float64{800, 500, 300, 100, 0} // Solar radiation [W m-2]
	Θ := 0.                                  // Slope [radians]

	polNames := []string{"SO2", "O3", "NO2", "H2O2", "ALD", "HCHO", "OP", "PAA", "FORM", "NH3", "PAN", "HNO3"}
	testData := [][][]float64{SO2, O3, NO2, H2O2, ALD, HCHO, OP, PAA, FORM, NH3, PAN, HNO3}
	gasData := []*GasData{So2Data, O3Data, No2Data, H2o2Data, AldData, HchoData, OpData, PaaData, HchoData, Nh3Data, PanData, Hno3Data}

	rain, dew := false, false
	for iPol, pol := range polNames {
		polData := testData[iPol]
		isSO2, isO3 := false, false
		if pol == "SO2" {
			isSO2 = true
		}
		if pol == "O3" {
			isO3 = true
		}
		for iSeason := 0; iSeason < 5; iSeason++ {
			for ig, G := range Garr {
				r_c := SurfaceResistance(gasData[iPol], G, Ts[iSeason], Θ,
					iSeason, iLandUse, rain, dew, isSO2, isO3)
				if math.Abs(r_c - polData[iSeason][ig]) > 50 {
				fmt.Printf("%v, %v, %v: %.0f, %g\n",pol, iSeason, G,r_c,polData[iSeason][ig])
				t.Fail()
			}
			}
			r_c := SurfaceResistance(gasData[iPol], 0., Ts[iSeason], Θ,
				iSeason, iLandUse, rain, true, isSO2, isO3) // dew
			if math.Abs(r_c - polData[iSeason][5]) > 50 {
				fmt.Printf("%v, %v, %v: %.0f, %g\n",pol,iSeason,"dew", r_c,polData[iSeason][5])
				t.Fail()
			}
			
			r_c = SurfaceResistance(gasData[iPol], 0., Ts[iSeason], Θ,
				iSeason, iLandUse, true, dew, isSO2, isO3) // rain
			if math.Abs(r_c - polData[iSeason][6]) > 50 {
				fmt.Printf("%v, %v, %v: %.0f, %g\n",pol,iSeason,"rain", r_c,polData[iSeason][6])
				t.Fail()
			}
		}
	}
}