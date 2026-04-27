package main

import (
	"testing"

	"github.com/montanaflynn/stats"
)

func TestGetDataset(t *testing.T) {
	x, y := getDataset(1, anscombe)

	if len(x) != 11 || len(y) != 11 {
		t.Fatalf("expected length 11, got len(x) = %d and len(y) = %d", len(x), len(y))
	}

	if x[0] != 10 || y[0] != 8.04 {
		t.Errorf("Unexpected values. got (%v, %v)", x[0], y[0])
	}
}

func TestMakePoints(t *testing.T) {
	x := []float64{1, 2, 3}
	y := []float64{4, 5, 6}

	points := makePoints(x, y)

	if len(points) != 3 {
		t.Errorf("Expectedd 3 points, got %d", len(points))
	}

	if points[0].X != 1 || points[0].Y != 4 {
		t.Errorf("unexpected first point %+v", points[0])
	}
}

func TestMakeSeries(t *testing.T) {
	x := []float64{1, 2}
	y := []float64{3}

	_, err := makeSeries(x, y)

	if err == nil {
		t.Fatal("expected error for mismatched x and y lengths")
	}

	x2 := []float64{1, 2, 3}
	y2 := []float64{4, 5, 6}

	series, err := makeSeries(x2, y2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(series) != 3 {
		t.Fatalf("expected 3 points, got %d", len(series))
	}
}

func TestGetRegressionStats(t *testing.T) {
	series := stats.Series{
		{X: 1, Y: 2},
		{X: 3, Y: 4},
		{X: 5, Y: 6},
	}

	slope, intercept, err := getRegressionStats(series)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if slope != 1 {
		t.Errorf("expected slope 2, got %v", slope)
	}

	if intercept != 1 {
		t.Errorf("expected intercept 0, got %v", intercept)
	}

	series2 := stats.Series{
		{X: 1, Y: 2},
		{X: 1, Y: 3},
	}

	_, _, err2 := getRegressionStats(series2)

	if err2 == nil {
		t.Fatal("expected division by zero error")
	}
}

func TestRunRegression(t *testing.T) {
	x := []float64{1, 2, 3}
	y := []float64{2, 4, 6}

	slope, intercept, err := runRegression(x, y)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if slope != 2 {
		t.Errorf("expected slope 2, got %v", slope)
	}

	if intercept != 0 {
		t.Errorf("expected intercept 0, got %v", intercept)
	}
}

func TestPlotData(t *testing.T) {
	err := plotData(anscombe, "test_output.png")
	if err != nil {
		t.Fatalf("plotData failed: %v", err)
	}
}
