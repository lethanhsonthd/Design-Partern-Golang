type Circle struct {
	// DrawingContext for this circle
	DrawingContext Drawer
	// Center of the circle
	Center Point
	// Radius of the circle
	Radius float64
}

// Draw draws a circle
func (circle *Circle) Draw() error {
	rect := Rect{
		Location: Point{
			X: circle.Center.X - circle.Radius,
			Y: circle.Center.Y - circle.Radius,
		},
		Size: Size{
			Width:  2 * circle.Radius,
			Height: 2 * circle.Radius,
		},
	}

	return circle.DrawingContext.DrawEllipseInRect(rect)
}
type Drawer interface {
	// DrawEllipseInRect draws an ellipse in rectanlge
	DrawEllipseInRect(Rect) error
}
// OpenGL drawer
type OpenGL struct{}

// DrawEllipseInRect draws an ellipse in rectangle
func (gl *OpenGL) DrawEllipseInRect(r Rect) error {
	fmt.Printf("OpenGL is drawing ellipse in rect %v", r)
	return nil
}
// Direct2D drawer
type Direct2D struct{}

// DrawEllipseInRect draws an ellipse in rectangle
func (d2d *Direct2D) DrawEllipseInRect(r Rect) error {
	fmt.Printf("Direct2D is drawing ellipse in rect %v", r)
	return nil
}
openGL := &uikit.OpenGL{}
direct2D := &uikit.Direct2D{}

circle := &uikit.Circle{
	Center: uikit.Point{X: 100, Y: 100},
	Radius: 50,
}

circle.DrawingContext = openGL
circle.Draw()

circle.DrawingContext = direct2D
circle.Draw()