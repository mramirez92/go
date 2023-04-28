package speed

import (
    "reflect"
    "testing"
)

func TestNewCar(t *testing.T) {
    cases := []struct {
        speed        int
        batteryDrain int
        expected     Car
    }{
        {
            speed:        5,
            batteryDrain: 5,
            expected: Car{
                speed:        5,
                batteryDrain: 5,
                battery:      100,
                distance:     0,
            },
        },
    }
    for _, tc := range cases {
        t.Run("car", func(t *testing.T){
            result := NewCar(tc.speed, tc.batteryDrain)
            if !reflect.DeepEqual(result, tc.expected) {
                t.Errorf("expected %v, but got %v", tc.expected, result)
            }
        })
    }
}

