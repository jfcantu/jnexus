package db

import "github.com/sirupsen/logrus"

// HandleNetsplit handles netsplits
func (c *Client) HandleNetsplit(server1, server2 string) error {
	logrus.Infof("NETSPLIT between %v and %v", server1, server2)
	linkData, err := c.GetLinkStatus(server1, server2)
	if err != nil {
		return err
	}

	// If this is a link that's not in the route map - delete it rather than retain it
	if linkData.Type == "SECONDARY" {
		if err := c.DeleteLink(server1, server2); err != nil {
			return err
		}
		return nil
	}

	// Otherwise, mark it as INACTIVE
	c.UpdateLink(server1, server2, "INACTIVE")
	return nil
}
