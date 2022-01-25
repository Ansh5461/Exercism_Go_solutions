package elon
import "fmt"
// TODO: define the 'Drive()' method
func(c *Car) Drive() {
    if c.battery < c.batteryDrain {
        return
    }
    c.battery = c.battery - c.batteryDrain
    c.distance = c.distance + c.speed
}
// TODO: define the 'DisplayDistance() string' method
func(c *Car) DisplayDistance() string {
    var s string
    s = fmt.Sprintf("Driven %d meters", c.distance)
    return s
}
// TODO: define the 'DisplayBattery() string' method
func(c *Car) DisplayBattery() string{
    s := fmt.Sprintf("Battery at %d%%",c.battery)
    return s
}
// TODO: define the 'CanFinish(trackDistance int) bool' method
func(c *Car) CanFinish(trackDistance int) bool {
    var dc int
    if c.battery < c.batteryDrain {
        return false
    }
    dc = c.battery/c.batteryDrain
    s := dc* c.speed
    if s >= trackDistance {
        return true
    } else {
    return false
    }
}
