package main

import (
	"fmt"
	"os"

	"github.com/wcharczuk/go-chart/v2"

	mva "performance-examples/mutex-vs-atomic"
	"performance-examples/mutex-vs-atomic/app/extra"
)

func main() {
	const minNumGoRoutines = 100
	const maxNumGoRoutines = 10000

	const numGoRoutinesStep = 100
	const numPoints = (maxNumGoRoutines - minNumGoRoutines) / numGoRoutinesStep

	srs := make([]chart.Series, 2)
	for i := 0; i < 2; i++ {
		xv := make([]float64, numPoints)
		yv := make([]float64, numPoints)

		for j, x := 0, minNumGoRoutines; j < numPoints; j, x = j+1, x+numGoRoutinesStep {
			initFunc, getFunc := mva.InitMutexMap, mva.GetDataFromMutexMap
			if i == 1 {
				initFunc, getFunc = mva.InitAtomicMap, mva.GetDataFromAtomicMap
			}

			xv[j] = float64(x)
			yv[j] = float64(extra.MainFunc(j, initFunc, getFunc).Milliseconds())
			fmt.Println(x, xv[j], yv[j])
		}

		name := "mutex"
		if i == 1 {
			name = "atomic"
		}

		srs[i] = chart.ContinuousSeries{
			Name:    name,
			XValues: xv,
			YValues: yv,
		}
	}

	g := chart.Chart{
		Title:  "Mutex VS Atomic",
		Series: srs,
	}
	g.XAxis.Name = "N"
	g.YAxis.Name = "Time (ms)"

	f, err := os.Create(os.Getenv("GRAPH_DIR") + "/mutex-vs-atomic.png")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()
	if err = g.Render(chart.PNG, f); err != nil {
		panic(err)
	}
}
