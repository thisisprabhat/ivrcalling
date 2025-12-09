package models

// IVRConfig holds the IVR flow configuration
type IVRConfig struct {
	IntroText  string      `json:"intro_text"`
	Actions    []IVRAction `json:"actions"`
	EndMessage string      `json:"end_message"`
}

// IVRAction represents a single IVR menu action
type IVRAction struct {
	Key         string `json:"key"`
	Message     string `json:"message"`
	Action      string `json:"action"`
	Description string `json:"description,omitempty"`
	ForwardTo   string `json:"forward_to,omitempty"`
}

// GetIVRConfig returns the configured IVR flow for Q&I
func GetIVRConfig() IVRConfig {
	return IVRConfig{
		IntroText: "Welcome to Q&I! We are transforming education with smart digital tools. " +
			"We help your school digitize teaching and measure true student understanding. " +
			"Our AI-powered platform provides topic analysis and targeted practice to boost academic performance. " +
			"With Q&I, teachers get deeper insights, students learn effectively, and your institution achieves measurable growth. " +
			"Ready to see how Q&I can revolutionize your classrooms?",
		Actions: []IVRAction{
			{
				Key:       "1",
				Message:   "To talk to Q&I team, press 1",
				Action:    "forward",
				ForwardTo: "+917905252436",
			},
			{
				Key:     "2",
				Message: "To know more about Q&I, press 2",
				Action:  "inform",
				Description: "Q&I is an AI-powered educational platform that helps schools digitize teaching and measure student understanding. " +
					"It provides topic analysis and targeted practice to improve academic performance, " +
					"giving teachers deeper insights and students more effective learning experiences.",
			},
			{
				Key:     "3",
				Message: "To hear this message again, press 3",
				Action:  "repeat",
			},
		},
		EndMessage: "Thank you for contacting Q&I. We look forward to helping your school achieve success. Goodbye!",
	}
}
