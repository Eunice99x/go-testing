package structsmethodsinterfaces

import (
	"testing"
)


func HelperFunc(t testing.TB, got float64,want float64){
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}


func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0,10.0}

	got := Perimeter(rectangle)
	want := 40.0

	HelperFunc(t, got,want)
}

func TestArea(t *testing.T){

	areaTests := []struct{
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12,Height: 6}, want:  72.0},
		{name: "Circle", shape: Circle{Radius: 2.5}, want: 19.634954084936208},
		{name: "Triangle", shape: Triangle{Height: 3, Base: 6.2}, want: 9.3},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
	// checkArea := func(t testing.TB, shape Shape, want float64){
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("rectangle area", func(t *testing.T) {
	// 	rectangle := Rectangle{12,6}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("circle area", func(t *testing.T){
	// 	circle := Circle{2.5}
	// 	checkArea(t, circle, 19.634954084936208)
	// })

}