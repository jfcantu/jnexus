package db

const getAllNodesQuery = `
MATCH (n)
RETURN ID(n), LABELS(n), PROPERTIES(n)
`

const getAllLinksQuery = `
MATCH (a)-[r]->(b)
RETURN DISTINCT ID(r),ID(a),ID(b),TYPE(r),PROPERTIES(r)
`

const getActiveLinksQuery = `
MATCH (a)-[r]->(b)
WHERE r.status="ACTIVE"
RETURN DISTINCT ID(r),ID(a),ID(b),TYPE(r),PROPERTIES(r)
`

const getLinkQuery = `
MATCH (a:node)-[r]-(b:node)
WHERE a.name = $server1 AND b.name = $server2
RETURN TYPE(r), r.status
`

const clearRouteStatesQuery = `
MATCH ()-[r]-()
SET r.status = "INACTIVE";
MATCH ()-[s:SECONDARY]-()
DELETE s
`

const updateLinkQuery = `
MATCH (a:node {name: $server1}),(b:node {name: $server2})
MERGE (a)-[r:%s]-(b)
SET r.status = $newStatus
`

const deleteLinkQuery = `
MATCH (a:node)-[r]-(b:node)
WHERE a.name = $server1 AND b.name = $server2
DELETE r
`

const getSplitNodesQuery = `
MATCH (a)-[r]-()
WITH a, collect(r) AS links
WHERE ALL(link IN links WHERE link.status="INACTIVE")
RETURN ID(a), LABELS(a), PROPERTIES(a)
`
