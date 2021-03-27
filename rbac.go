package runscope

import "fmt"

type Role struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	UUID        string   `json:"uuid"`
}

// ListRoles list all the roles. See https://api.blazemeter.com/api-monitoring/#list-roles
func (client *Client) ListRoles(teamID string) ([]*Role, error) {
	resource, error := client.readResource("roles", teamID,
		fmt.Sprintf("/teams/%s/roles", teamID))
	if error != nil {
		return nil, error
	}

	roles, error := getRolesFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return roles, nil
}

// CreateRole creates a new role with specified permissions. See https://api.blazemeter.com/api-monitoring/#create-role
func (client *Client) CreateRole(role *Role, teamID string) (*Role, error) {
	newRole, err := client.createResource(role, "role", role.Name,
		fmt.Sprintf("/teams/%s/roles", teamID))
	if err != nil {
		return nil, err
	}

	roleInfo, err := getRoleFromResponse(newRole.Data)
	if err != nil {
		return nil, err
	}

	return roleInfo, nil
}

// DeleteRole delete runscope role. See https://api.blazemeter.com/api-monitoring/#delete-a-role
func (client *Client) DeleteRole(role *Role, teamID string) error {
	return client.deleteResource("role", role.Name, fmt.Sprintf("/teams/%s/roles/%s", teamID, role.UUID))
}

// ReadRole creates a new role with specified permissions. See https://api.blazemeter.com/api-monitoring/#role-details
func (client *Client) ReadRole(role *Role, teamID string) (*Role, error) {
	resource, err := client.readResource("role", role.Name,
		fmt.Sprintf("/teams/%s/roles/%s", teamID, role.UUID))
	if err != nil {
		return nil, err
	}

	readRole, err := getRoleFromResponse(resource.Data)
	if err != nil {
		return nil, err
	}

	return readRole, nil
}

// UpdateRole update a runscope role. See https://api.blazemeter.com/api-monitoring/#modify-role
func (client *Client) UpdateRole(role *Role, teamID string) (*Role, error) {
	resource, error := client.updateResource(role, "role", role.UUID,
		fmt.Sprintf("/teams/%s/roles/%s", teamID, role.UUID))
	if error != nil {
		return nil, error
	}

	readRole, error := getRoleFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return readRole, nil
}

func getRolesFromResponse(response interface{}) ([]*Role, error) {
	var roles []*Role
	err := decode(&roles, response)
	return roles, err
}

func getRoleFromResponse(response interface{}) (*Role, error) {
	role := new(Role)
	err := decode(role, response)
	return role, err
}
