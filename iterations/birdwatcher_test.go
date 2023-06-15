package birdwatcher

import (
    "reflect"
    "testing"
)

func TestBirdCountSum(t *testing.T){
    cases := []struct {
        birdCount []int
        expected int
    }{
        {
            birdCount: []int{1, 1, 2, 3, 5, 0, 2},
            expected: 14,
        },
        {
            birdCount: []int{0, 1, 3, 0, 2, 0, 1},
            expected: 7,
        },
    }
    for _, tc := range cases {
        results := TotalBirdCount(tc.birdCount)
        if !reflect.DeepEqual(results, tc.expected) {
            t.Errorf("got %+v, expected %+v", results, tc.expected)
        }
    }
}

func TestBirdsPerWeek(t *testing.T){
    cases := []struct {
        birdCount []int
        week int
        expected int
    }{
        {
            birdCount: []int{1, 1, 2, 3, 5, 0, 2, 0, 1, 1, 2, 1, 3, 1},
            week: 2,
            expected: 9,
        },
        {
            birdCount: []int{1, 1, 2, 3, 5, 0, 2, 0, 1, 1, 2, 1, 3, 1},
            week: 1,
            expected: 14,
        },
    }
    for _, tc := range cases {
        results := BirdsPerWeek(tc.birdCount, tc.week)
        if !reflect.DeepEqual(results, tc.expected) {
            t.Errorf("got %+v, expected %+v", results, tc.expected)
        }
    }
}

func TestFixBirdCounting(t *testing.T){
    cases := []struct {
        birdCount []int
        expected []int
    }{
        {
        birdCount: []int{0, 1, 2, 0, 1, 3, 2},
        expected: []int{1, 1, 3, 1, 1, 4, 2},
        },
    }
    for _, tc := range cases {
        results := FixBirdCounting(tc.birdCount)
        if !reflect.DeepEqual(results, tc.expected) {
            t.Errorf("got %+v, expected %+v", results, tc.expected)
        }
    }
}