package servey

type CreateServeyRequest struct {
	Title           string     `json:"title" binding:"required,title"`
	SubTitle        string     `json:"subTitle" binding:"required,subtitle"`
	MetaTitle       string     `json:"metaTitle,omitempty"`
	Description     string     `json:"description,omitempty"`
	Category        string     `json:"category,omitempty"`
	MetaDescription string     `json:"metaDescription,omitempty"` //`json:"metaDescription" json:"metaDescription"`
	IsOptional      bool       `json:"isOptional" binding:"required,isOptional"`
	Options         []Option   `json:"options,omitempty"`
	OptionsLbs      []Option   `json:"optionslbs,omitempty"`
	IsMultiSelect   bool       `json:"isMultiSelect,omitempty"`
	SeqNo           int        `json:"seqNo,omitempty"`
	Images          []string   `json:"images,omitempty"`
	Validation      Validation `json:"validation,omitempty"`
	ProfileKey      string     `json:"profileKey,omitempty"`
	Exclusion       bool       `json:"exclusion,omitempty"`
	Deleted         bool       `json:"deleted,omitempty"`
}

// {
//   "title": "Health & Wellness Survey",
//   "subTitle": "Your daily habits",
//   "metaTitle": "Health Survey 2025",
//   "description": "A survey to understand your daily health habits.",
//   "category": "Health",
//   "metaDescription": "Survey about health and wellness habits.",
//   "isOptional": false,
//   "options": [
//     {
//       "name": "Option 1",
//       "value": "opt1",
//       "icon": "🍎",
//       "description": "Eat fruits daily"
//     },
//     {
//       "name": "Option 2",
//       "value": "opt2",
//       "icon": "🏃",
//       "description": "Exercise regularly"
//     }
//   ],
//   "optionslbs": [
//     {
//       "name": "Option LBS 1",
//       "value": "lbs1",
//       "icon": "💧",
//       "description": "Drink water"
//     }
//   ],
//   "isMultiSelect": true,
//   "seqNo": 1,
//   "images": [
//     "https://example.com/image1.jpg",
//     "https://example.com/image2.jpg"
//   ],
//   "validation": {
//     "min": 1,
//     "max": 3
//   },
//   "profileKey": "health_profile",
//   "exclusion": false,
//   "createdBy": "665f1b2e8e4b1c2a3d4e5f6a",  // Example ObjectID as string
//   "deleted": false,
//   "createdAt": "2025-06-10T12:00:00Z",
//   "updatedAt": "2025-06-10T12:00:00Z"
// }
