/*
 * rach
 *
 * This is an an organization app. it helps keeps books and other items in location.
 *
 * API version: 1.0.0
 * Contact: aravind@agastya-library.in
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Object struct {

	Id int64 `json:"id,omitempty"`

	// category of object
	Category string `json:"category,omitempty"`
}

// AssertObjectRequired checks if the required fields are not zero-ed
func AssertObjectRequired(obj Object) error {
	return nil
}

// AssertRecurseObjectRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Object (e.g. [][]Object), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseObjectRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aObject, ok := obj.(Object)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertObjectRequired(aObject)
	})
}
