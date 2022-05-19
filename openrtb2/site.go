package openrtb2

import (
	"encoding/json"

	"github.com/mxmCherry/openrtb/v16/adcom1"
)

// 3.2.13 Object: Site
//
// This object should be included if the ad supported content is a website as opposed to a non-browser application.
// A bid request must not contain both a Site and an App object.
// At a minimum, it is useful to provide a site ID or page URL, but this is not strictly required.
type Site struct {

	// Attribute:
	//   id
	// Type:
	//   string; recommended
	// Description:
	//   Exchange-specific site ID.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Site name (may be aliased at the publisher’s request).
	Name string `json:"name,omitempty"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   Domain of the site (e.g., “mysite.foo.com”).
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   cattax
	// Type:
	//   integer
	// Description:
	//   The taxonomy in use. Refer to the AdCOM list List: Category
	//   Taxonomies for values. If no cattax field is supplied IAB Content
	//   Category Taxonomy 1.0 is assumed.
	CatTax adcom1.CategoryTaxonomy `json:"cattax,omitempty"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IABTL content categories of the site.
	//   The taxonomy to be used is defined by the cattax field.
	Cat []string `json:"cat,omitempty"`

	// Attribute:
	//   sectioncat
	// Type:
	//   string array
	// Description:
	//   Array of IABTL content categories that describe the current
	//   section of the site. The taxonomy to be used is defined by
	//   the cattax field.
	SectionCat []string `json:"sectioncat,omitempty"`

	// Attribute:
	//   pagecat
	// Type:
	//   string array
	// Description:
	//   Array of IABTL content categories that describe the current
	//   page or view of the site.
	//   The taxonomy to be used is defined by the cattax field.
	PageCat []string `json:"pagecat,omitempty"`

	// Attribute:
	//   page
	// Type:
	//   string
	// Description:
	//   URL of the page where the impression will be shown.
	Page string `json:"page,omitempty"`

	// Attribute:
	//   ref
	// Type:
	//   string
	// Description:
	//   Referrer URL that caused navigation to the current page
	Ref string `json:"ref,omitempty"`

	// Attribute:
	//   search
	// Type:
	//   string
	// Description:
	//   Search string that caused navigation to the current page.
	Search string `json:"search,omitempty"`

	// Attribute:
	//   mobile
	// Type:
	//   integer
	// Description:
	//   Indicates if the site has been programmed to optimize layout
	//   when viewed on mobile devices, where 0 = no, 1 = yes.
	Mobile int8 `json:"mobile,omitempty"`

	// Attribute:
	//   privacypolicy
	// Type:
	//   integer
	// Description:
	//   Indicates if the site has a privacy policy, where 0 = no, 1 = yes.
	PrivacyPolicy int8 `json:"privacypolicy,omitempty"`

	// Attribute:
	//   publisher
	// Type:
	//   object
	// Description:
	//   Details about the Publisher (Section 3.2.15) of the site.
	Publisher *Publisher `json:"publisher,omitempty"`

	// Attribute:
	//   content
	// Type:
	//   object
	// Description:
	//   Details about the Content (Section 3.2.16) within the site.
	Content *Content `json:"content,omitempty"`

	// Attribute:
	//   keywords
	// Type:
	//   string
	// Description:
	//   Comma separated list of keywords about the site. Only one of
	//  ‘keywords’ or ‘kwarray’ may be present.
	Keywords string `json:"keywords,omitempty"`

	// Attribute:
	//   kwarray
	// Type:
	//   string
	// Description:
	//   Array of keywords about the site. Only one of ‘keywords’ or
	//   ‘kwarray’ may be present.
	KwArray []string `json:"kwarray,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
