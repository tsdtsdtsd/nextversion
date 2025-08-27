package nextversion_test

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"github.com/tsdtsdtsd/nextversion/pkg/nextversion"
)

type lastTagTestcase struct {
	RepoPath        string
	ExpectedTagName string
	ExpectedError   error
	ExpectedExists  bool
}

func TestLastTag(t *testing.T) {

	testCases := []lastTagTestcase{
		{
			RepoPath:        "../../fixtures/valid-tag",
			ExpectedTagName: "v0.1.0",
			ExpectedError:   nil,
			ExpectedExists:  true,
		},
		{
			RepoPath:        "../../fixtures/no-valid-tags",
			ExpectedTagName: "",
			ExpectedError:   nil,
			ExpectedExists:  false,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.RepoPath, func(t *testing.T) {

			repo, err := git.PlainOpen(testCase.RepoPath)
			assert.NoError(t, err)

			lastTag, err := nextversion.LastTag(repo)
			assert.ErrorIs(t, err, testCase.ExpectedError)

			assert.Equal(t, testCase.ExpectedExists, lastTag.Exists())
			assert.Equal(t, testCase.ExpectedTagName, lastTag.Name)
		})
	}
}

func TestLastTag_RepoFailures(t *testing.T) {
	_, err := nextversion.LastTag(nil)
	assert.Error(t, err)

	mockRepo := nextversion.NewMockGitRepository(t)
	mockRepo.On("Tags").Return(nil, assert.AnError).Once()
	mockRepo.On("TagObject", mock.Anything).Return(nil, assert.AnError).Once()

	_, err = nextversion.LastTag(mockRepo)
	assert.Error(t, err)

	_, err = mockRepo.TagObject(plumbing.NewHash("test"))
	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}
