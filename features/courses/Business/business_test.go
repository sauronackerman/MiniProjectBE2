package Business

import (
	"RestfulAPIElearningVideo/features/courses"
	"RestfulAPIElearningVideo/features/courses/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	crsData mocks.Data
	crsBusiness courses.Business
	crsValue courses.CourseCore
	cvid2 []courses.VideoCore
)

func TestMain(m *testing.M) {
	crsBusiness = NewCourseBusiness(&crsData)
	crsValue = courses.CourseCore{
		ID:          0,
		Title:       "",
		Description: "",
		PlaylistID:  "",
		Videos:      nil,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	cvid2 = []courses.VideoCore{
		{
			ID:        0,
			Title:     "",
			CourseID:  "",
			VideoID:   "",
			Duration:  "",
			Task:      courses.Task{},
			UpdatedAt: time.Time{},
			CreatedAt: time.Time{},
		},
	}
	os.Exit(m.Run())

}

func TestCreateCourse(t *testing.T)  {
	t.Run("create course sukses", func(t *testing.T) {
		crsData.On("InsertCourse", mock.Anything).Return(crsValue, nil) .Once()
		_, _, status := crsBusiness.CreateCourse(crsValue)
		assert.Equal(t, status, http.StatusOK)
	})
	t.Run("create course gagal", func(t *testing.T) {
		crsData.On("InsertCourse", mock.Anything).Return(crsValue, errors.New("err")) .Once()
		_, err, status := crsBusiness.CreateCourse(crsValue)
		assert.Equal(t, status, http.StatusInternalServerError)
		assert.Error(t, err)
	})
}

func TestAddVideoToCourse(t *testing.T)  {
	t.Run("create course sukses", func(t *testing.T) {
		crsData.On("GetPlaylistIdforVideo", mock.Anything, mock.AnythingOfType("string")).Return(cvid2, nil) .Once()
		data, _,_ := crsBusiness.AddVideoToCourse(context.Background(), "1")
		assert.Equal(t, data, cvid2)
	})
	t.Run("create course gagal", func(t *testing.T) {
		crsData.On("GetPlaylistIdforVideo", mock.Anything, mock.AnythingOfType("string")).Return(cvid2, errors.New("err")) .Once()
		_, err, status := crsBusiness.AddVideoToCourse(context.Background(), "2")
		assert.Equal(t, status, http.StatusInternalServerError)
		assert.Error(t, err)
	})
}

func TestDeleteCourseById(t *testing.T)  {
	t.Run("delete course sukses", func(t *testing.T) {
		crsData.On("DeleteCourseDataById", mock.AnythingOfType("string")).Return(crsValue, nil) .Once()
		data, _ := crsBusiness.DeleteCourseById(crsValue.PlaylistID)
		assert.Equal(t, data, crsValue)
	})
}


