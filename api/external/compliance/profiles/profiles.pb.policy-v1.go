// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: api/external/compliance/profiles/profiles.proto

package profiles

import policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.compliance.profiles.v1.ProfilesService/Read", "compliance:profiles:storage:{owner}", "read", "GET", "/compliance/profiles/read/{owner}/{name}/version/{version}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*ProfileDetails); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "owner":
					return m.Owner
				case "name":
					return m.Name
				case "version":
					return m.Version
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.profiles.v1.ProfilesService/ReadFromMarket", "compliance:profiles:market", "read", "GET", "/compliance/market/read/{name}/version/{version}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*ProfileDetails); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "owner":
					return m.Owner
				case "name":
					return m.Name
				case "version":
					return m.Version
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.profiles.v1.ProfilesService/Delete", "compliance:profiles:{owner}", "delete", "DELETE", "/compliance/profiles/{owner}/{name}/version/{version}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*ProfileDetails); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "owner":
					return m.Owner
				case "name":
					return m.Name
				case "version":
					return m.Version
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.profiles.v1.ProfilesService/List", "compliance:profiles", "search", "POST", "/compliance/profiles/search", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "sort":
					return m.Sort
				case "owner":
					return m.Owner
				case "name":
					return m.Name
				case "version":
					return m.Version
				default:
					return ""
				}
			})
		}
		return ""
	})
}
