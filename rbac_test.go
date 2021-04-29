package runscope

import (
	"testing"
)

func TestListRoles(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()

	listRoles, err := client.ListRoles(teamID)
	if err != nil {
		t.Error(err)
	}

	if listRoles[0].Name != "Administrators" {
		t.Errorf("Expected to have role %s got %s", "Administrators", listRoles[0].Name)
	}
}

func TestCreateReadRole(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	newRole := &Role{Name: "test_role", Permissions: []string{"team:people:view"}}

	newRole, err := client.CreateRole(newRole, teamID)
	if err != nil {
		t.Error(err)
	}

	defer client.DeleteRole(newRole, teamID) // nolint: errcheck

	readRole, err := client.ReadRole(newRole, teamID)
	if err != nil {
		t.Error(err)
	}

	if readRole.Name != "test_role" {
		t.Errorf("Expected role name %s, actual %s", "test_role", readRole.Name)
	}

	if readRole.Permissions[0] != "team:people:view" {
		t.Errorf("Expected role permission %s, actual %s", "team:people:view", readRole.Permissions[0])
	}
}

func TestUpdateRole(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	newRole := &Role{Name: "test_role", Permissions: []string{"team:people:view"}}

	newRole, err := client.CreateRole(newRole, teamID)
	if err != nil {
		t.Error(err)
	}

	defer client.DeleteRole(newRole, teamID) // nolint: errcheck

	newRole.Permissions = []string{"team:billing:view"}
	updatedRole, err := client.UpdateRole(newRole, teamID)
	if err != nil {
		t.Error(err)
	}

	readRole, err := client.ReadRole(updatedRole, teamID)

	if err != nil {
		t.Error(err)
	}

	if readRole.Permissions[0] != "team:billing:view" {
		t.Errorf("Expected role permission %s, actual %s", "team:people:view", readRole.Permissions[0])
	}
}
