package services

import (
	"testing"

	"git.wndv.co/sharp/app/models"
)

func TestFormatEmail(t *testing.T) {
	testCases := []struct {
		email        models.EmailView
		bodytemplate string
		expected     string
		desc         string
	}{
		{
			email: models.EmailView{
				Email:     "test55@gmail.com",
				Firstname: "pete",
				Lastname:  "pppt2",
			},
			bodytemplate: "Test {{firstname}}",
			expected:     "Test pete",
			desc:         "Case firstname",
		},
		{
			email: models.EmailView{
				Email:     "test55@gmail.com",
				Firstname: "pete",
				Lastname:  "pppt2",
			},
			bodytemplate: "Test {{firstname}} {{lastname}}",
			expected:     "Test pete pppt2",
			desc:         "Case firstname and lastname",
		},
		{
			email: models.EmailView{
				Email:     "test55@gmail.com",
				Firstname: "pete",
				Lastname:  "pppt2",
			},
			bodytemplate: "Test",
			expected:     "Test",
			desc:         "Case NO firstname and lastname",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			campaign := models.CampaignView{
				Subject:      "Test",
				BodyTemplate: tC.bodytemplate,
			}

			output := FormatEmail(tC.email, campaign)
			if output != tC.expected {
				t.Errorf("body is not matched expect=%v, got=%v", tC.expected, output)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	// 11397 ns/op ~ 0.011397 ms
	for i := 0; i < b.N; i++ {
		email := models.EmailView{
			Email:     "test55@gmail.com",
			Firstname: "pete",
			Lastname:  "pppt2",
		}
		campaign := models.CampaignView{
			Subject:      "Test",
			BodyTemplate: "Test {{firstname}} {{lastname}}Curabitur aliquet quam id dui posuere blandit. Curabitur aliquet quam id dui posuere blandit. Proin eget tortor risus. Vestibulum ac diam sit amet quam vehicula elementum sed sit amet dui. Donec rutrum congue leo eget malesuada. Nulla quis lorem ut libero malesuada feugiat. Proin eget tortor risus. Curabitur arcu erat, accumsan id imperdiet et, porttitor at sem. Praesent sapien massa, convallis a pellentesque nec, egestas non nisi. Proin eget tortor risus. Nulla quis lorem ut libero malesuada feugiat. Curabitur aliquet quam id dui posuere blandit. Nulla porttitor accumsan tincidunt. Mauris blandit aliquet elit, eget tincidunt nibh pulvinar a. Nulla porttitor accumsan tincidunt. Vivamus magna justo, lacinia eget consectetur sed, convallis at tellus. Curabitur arcu erat, accumsan id imperdiet et, porttitor at sem. Mauris blandit aliquet elit, eget tincidunt nibh pulvinar a. Proin eget tortor risus. Nulla porttitor accumsan tincidunt. Curabitur arcu erat, accumsan id imperdiet et, porttitor at sem. Vestibulum ac diam sit amet quam vehicula elementum sed sit amet dui. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur aliquet quam id dui posuere blandit. Nulla porttitor accumsan tincidunt. Sed porttitor lectus nibh. Curabitur arcu erat, accumsan id imperdiet et, porttitor at sem. Cras ultricies ligula sed magna dictum porta. Donec rutrum congue leo eget malesuada. Donec sollicitudin molestie malesuada.",
		}
		_ = FormatEmail(email, campaign)
	}
}
