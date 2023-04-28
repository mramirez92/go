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
		t.Run("car", func(t *testing.T) {
			result := NewCar(tc.speed, tc.batteryDrain)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestTrack(t *testing.T) {
	got := NewTrack(5)
	expected := Track{distance: 5}
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}

}

func TestDrive(t *testing.T) {
	cases := []struct {
		car      Car
		expected Car
	}{
		{
			car: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      100,
				distance:     0,
			},
			expected: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      98,
				distance:     5,
			},
		},
		{
			car: Car{
				speed:        10,
				batteryDrain: 5,
				battery:      0,
				distance:     20,
			},
			expected: Car{
				speed:        10,
				batteryDrain: 5,
				battery:      0,
				distance:     20,
			},
		},
	}
	for _, tc := range cases {
		results := Drive(tc.car)
		if results != tc.expected {
			t.Errorf("got %+v, expected %+v", results, tc.expected)
		}
	}
}

func TestCanFinish(t *testing.T) {
	cases := []struct {
		car      Car
		track    Track
		expected bool
	}{
		{
			car: Car{
				speed:        10,
				batteryDrain: 5,
				battery:      0,
				distance:     20,
			},
			track: Track{
				distance: 10,
			},
			expected: false,
		},
		{
			car: Car{
				speed:        7,
				batteryDrain: 5,
				battery:      80,
				distance:     20,
			},
			track: Track{
				distance: 10,
			},
			expected: true,
		},
	}
	for _, tc := range cases {
		results := CanFinish(tc.car, tc.track)
		if results != tc.expected {
			t.Errorf("got %+v, expected %+v", results, tc.expected)
		}
	}

}
