package db

import (
	"fmt"

	"github.com/jfcantu/jnexus/pkg/util"

	"github.com/jfcantu/jnexus/pkg/network"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cast"
)

// GetLinkStatus returns the link status, or nil if the link is not defined.
func (c *Client) GetLinkStatus(server1, server2 string) (*network.Link, error) {
	server1 = util.GetServerName(server1)
	server2 = util.GetServerName(server2)

	result, err := (*c.Session).Run(
		getLinkQuery,
		map[string]interface{}{
			"server1": server1,
			"server2": server2,
		})
	if err != nil {
		return nil, err
	}

	if !result.Next() {
		return nil, nil
	}

	link := new(network.Link)

	link.Type = cast.ToString(result.Record().GetByIndex(0))
	link.Status = cast.ToString(result.Record().GetByIndex(1))

	return link, nil
}

// UpdateLink updates the link status in the DB (creating a SECONDARY link for links not in the route map)
func (c *Client) UpdateLink(server1, server2, newStatus string) error {
	server1 = util.GetServerName(server1)
	server2 = util.GetServerName(server2)

	// LINKS returns the local server as a link - discard that
	if server1 == server2 {
		return nil
	}

	logrus.Infof("Link state change between %v and %v: link is now %v", server1, server2, newStatus)

	currentState, err := c.GetLinkStatus(server1, server2)
	if err != nil {
		return err
	}

	// By default, assume the link is not in the route map
	linkType := "SECONDARY"

	// If the link is in the database - use whatever type it exists as
	if currentState != nil {
		linkType = currentState.Type
	}

	result, err := (*c.Session).Run(
		fmt.Sprintf(updateLinkQuery, linkType),
		map[string]interface{}{
			"server1":   server1,
			"server2":   server2,
			"newStatus": newStatus},
	)
	if err != nil {
		return err
	}

	logrus.Debug(result)
	return nil
}

// DeleteLink deletes a link relationship between two nodes.
func (c *Client) DeleteLink(server1, server2 string) error {
	server1 = util.GetServerName(server1)
	server2 = util.GetServerName(server2)

	result, err := (*c.Session).Run(
		deleteLinkQuery,
		map[string]interface{}{
			"server1": server1,
			"server2": server2,
		},
	)
	if err != nil {
		return err
	}
	logrus.Debug(result)
	return nil
}

// ClearRouteStates clears all route states
func (c *Client) ClearRouteStates() error {
	logrus.Infof("clearing route states")
	result, err := (*c.Session).Run(clearRouteStatesQuery, map[string]interface{}{})
	if err != nil {
		logrus.Warn(err)
		return err
	}
	logrus.Debug(result)
	return nil
}

// GetNetworkStatus compiles the entire graph into a Network struct.
func (c *Client) GetNetworkStatus() (*Network, error) {
	nodes, err := c.GetServers()
	if err != nil {
		return nil, err
	}

	links, err := c.GetActiveLinks()
	if err != nil {
		return nil, err
	}

	return &Network{
		Nodes: nodes,
		Links: links,
	}, nil
}

// GetActiveLinks returns a list of all active links on the network.
func (c *Client) GetActiveLinks() ([]*Link, error) {
	links := []*Link{}

	result, err := (*c.Session).Run(
		getActiveLinksQuery,
		map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	for result.Next() {
		links = append(links, &Link{
			ID:         cast.ToInt(result.Record().GetByIndex(0)),
			Start:      cast.ToInt(result.Record().GetByIndex(1)),
			End:        cast.ToInt(result.Record().GetByIndex(2)),
			Type:       cast.ToString(result.Record().GetByIndex(3)),
			Properties: cast.ToStringMapString(result.Record().GetByIndex(4)),
		})
	}

	return links, nil
}

// GetAllLinks returns a list of all links on the network, active or not.
func (c *Client) GetAllLinks() ([]*Link, error) {
	links := []*Link{}

	result, err := (*c.Session).Run(
		getAllLinksQuery,
		map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	for result.Next() {
		links = append(links, &Link{
			ID:         cast.ToInt(result.Record().GetByIndex(0)),
			Start:      cast.ToInt(result.Record().GetByIndex(1)),
			End:        cast.ToInt(result.Record().GetByIndex(2)),
			Type:       cast.ToString(result.Record().GetByIndex(3)),
			Properties: cast.ToStringMapString(result.Record().GetByIndex(4)),
		})
	}

	return links, nil
}

// GetServers returns a list of all the servers on the network.
func (c *Client) GetServers() ([]*Node, error) {
	nodes := []*Node{}

	result, err := (*c.Session).Run(
		getAllNodesQuery,
		map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	for result.Next() {
		nodes = append(nodes, &Node{
			ID:         cast.ToInt(result.Record().GetByIndex(0)),
			Labels:     cast.ToStringSlice(result.Record().GetByIndex(1)),
			Properties: cast.ToStringMapString(result.Record().GetByIndex(2)),
		})
	}

	return nodes, nil
}

// GetSplitServers returns a list of split servers.
func (c *Client) GetSplitServers() ([]*Node, error) {
	result, err := (*c.Session).Run(
		getSplitNodesQuery,
		map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	nodeList := []*Node{}

	for result.Next() {
		nodeList = append(nodeList, &Node{
			ID:         cast.ToInt(result.Record().GetByIndex(0)),
			Labels:     cast.ToStringSlice(result.Record().GetByIndex(1)),
			Properties: cast.ToStringMapString(result.Record().GetByIndex(2)),
		})
	}

	return nodeList, nil
}
