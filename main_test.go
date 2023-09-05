package main

import "testing"

func TestIsAgeRestricted(t *testing.T) {

	testCases := []struct {
		name           string
		videoId        string
		expectedResult bool
	}{
		{
			name:           "IsRestricted_1",
			videoId:        "sxF9PGRiabw",
			expectedResult: true,
		},
		{
			name:           "IsRestricted_2",
			videoId:        "ieumolVsiWI",
			expectedResult: true,
		},
		{
			name:           "IsRestricted_3",
			videoId:        "tSig_iJmJT0",
			expectedResult: true,
		},
		{
			name:           "IsRestricted_4",
			videoId:        "cdDxXQcyvZM",
			expectedResult: true,
		},
		{
			name:           "IsRestricted_5",
			videoId:        "tYrONr5OH7Y",
			expectedResult: true,
		},
		{
			name:           "NotRestricted_1",
			videoId:        "ksrH_DK_jt8",
			expectedResult: false,
		},
		{
			name:           "NotRestricted_2",
			videoId:        "JpowNxr1n8k",
			expectedResult: false,
		},
		{
			name:           "NotRestricted_3",
			videoId:        "GLWBM8kvNvQ",
			expectedResult: false,
		},
		{
			name:           "NotRestricted_4",
			videoId:        "MTnGKnTIbys",
			expectedResult: false,
		},
		{
			name:           "NotRestricted_5",
			videoId:        "DdVUB-70mjs",
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			result := IsAgeRestricted(testCase.videoId)

			if result != testCase.expectedResult {

				t.Errorf("Test failed, Want=%t, got=%t  ", testCase.expectedResult, result)

			}

			t.Logf("%s test is passed", testCase.name)

		})

	}

}
