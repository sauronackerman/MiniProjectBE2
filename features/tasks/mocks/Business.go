package mocks

import (
	"RestfulAPIElearningVideo/features/tasks"
	"github.com/stretchr/testify/mock"
)

type Business struct {
	mock.Mock
}

func (_m *Business) CreateTaskByVideoId(data tasks.TaskCore) (tasks.TaskCore, error, int) {
	ret := _m.Called(data)

	var m1 tasks.TaskCore
	if rf, ok := ret.Get(0).(func( tasks.TaskCore) tasks.TaskCore); ok {
		m1 = rf(data)
	} else {
		m1 = ret.Get(0).(tasks.TaskCore)
	}

	var m2 error
	if rf, ok := ret.Get(1).(func( tasks.TaskCore) error); ok {
		m2 = rf(data)
	} else {
		m2 = ret.Error(1)
	}

	var m3 int
	if rf, ok := ret.Get(2).(func( tasks.TaskCore) int); ok {
		m3 = rf(data)
	} else {
		m3 = ret.Get(2).(int)
	}

	return m1, m2, m3
}

func (_m *Business) FindTaskByVideoId(videoId string) (tasks.TaskCore, error, int) {
	ret := _m.Called(videoId)

	var m1 tasks.TaskCore
	if rf, ok := ret.Get(0).(func( string) tasks.TaskCore); ok {
		m1 = rf(videoId)
	} else {
		m1 = ret.Get(0).(tasks.TaskCore)
	}

	var m2 error
	if rf, ok := ret.Get(1).(func( string) error); ok {
		m2 = rf(videoId)
	} else {
		m2 = ret.Error(1)
	}

	var m3 int
	if rf, ok := ret.Get(2).(func( string) int); ok {
		m3 = rf(videoId)
	} else {
		m3 = ret.Get(2).(int)
	}

	return m1, m2, m3
}

func (_m *Business) DeleteTask(videoId string) (tasks.TaskCore, error) {
	ret := _m.Called(videoId)
	var r0 tasks.TaskCore
	if rf, ok := ret.Get(0).(func(string) tasks.TaskCore); ok {
		r0 = rf(videoId)
	} else {
		r0 = ret.Get(0).(tasks.TaskCore)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(videoId)
	} else {
		r1 = ret.Error(1)
	}
	return r0,r1
}

