package resource

import (
	"math/rand"
)

// Resource is a useful resource
type Resource struct {
	Name         string
	Description  string
	MainMaterial string
	Origin       string
	Tags         []string
	Commonality  int
	Value        int
}

// ByTag returns a slice of resources that have the given tag
func ByTag(tag string, from []Resource) []Resource {
	var resources []Resource

	for _, p := range from {
		if p.HasTag(tag) {
			resources = append(resources, p)
		}
	}

	return resources
}

// HasTag returns true if the resource has a given tag
func (resource Resource) HasTag(tag string) bool {
	for _, t := range resource.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// InSlice checks if a given resource is in a slice of resources
func (resource Resource) InSlice(resources []Resource) bool {
	for _, r := range resources {
		if r.Name == resource.Name {
			return true
		}
	}

	return false
}

// Random returns a random resource from a list
func Random(resources []Resource) Resource {
	if len(resources) == 1 {
		return resources[0]
	} else if len(resources) < 1 {
		panic("No resources given")
	}

	resource := resources[rand.Intn(len(resources))]

	return resource
}

// RandomSet returns a slice of random elements of the given resources
func RandomSet(min int, max int, resources []Resource) []Resource {
	var result []Resource
	var resource Resource

	numberOfResources := rand.Intn(max-min) + min
	if numberOfResources > len(resources) {
		numberOfResources = len(resources)
	}

	for i := 0; i < numberOfResources; i++ {
		resource = Random(resources)
		if !resource.InSlice(result) {
			result = append(result, resource)
		}
	}

	return result
}
