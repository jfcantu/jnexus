import { Injectable } from '@angular/core';
import { DataSet, Node, Edge, Network } from 'vis';
import { NetworkState, Server } from '../routing-service/entities';
import { RoutingService } from '../routing-service/routing.service';

@Injectable({
  providedIn: 'root',
})
export class MapService {
  public networkInstance: Network;
  public nodeData: DataSet<Node> = new DataSet<Node>();
  public activeEdgeData: DataSet<Edge> = new DataSet<Edge>();
  public allEdgeData: DataSet<Edge> = new DataSet<Edge>();

  private networkOptions = {
    nodes: {
      shape: 'box',
      color: 'blue',
      font: {
        color: 'white',
      },
    },
    edges: {
      width: 2,
      smooth: {
        enabled: true,
        type: 'continuous',
        roundness: 0,
      },
    },
    physics: {
      enabled: true,
      barnesHut: {
        avoidOverlap: 1,
      },
    },
  };

  constructor(private routingService: RoutingService) {
    this.routingService.networkState.subscribe((newMap) => {
      if (newMap == null) {
        return;
      }
      this.updateMap(newMap);
    });
  }

  public render = (container: any) => {
    this.routingService.refreshStatus();

    this.networkInstance = new Network(
      container,
      { nodes: this.nodeData, edges: this.activeEdgeData },
      this.networkOptions
    );
  };

  public updateMap = (newState: NetworkState) => {
    let newNodeData = new DataSet<any>();
    let newEdgeData = new DataSet<any>();

    // Iterate through known servers and add them to the new node list
    for (let server of this.routingService.serverData.values()) {
      let newNode = <any>{
        id: server.id,
        label: server.properties.name,
        data: server,
      };

      // Set color based on server type
      if (server.labels.includes('hub')) {
        newNode.color = 'green';
      } else if (server.labels.includes('service')) {
        newNode.color = 'lime';
        newNode.font = { color: 'black' };
      }

      // ... but if it's split (according to the latest state), it will always be red
      if (newState.splitServers.includes(server.id)) {
        newNode.color = 'red';
      }

      this.nodeData.update(newNode);
    }

    // Add all links to map
    newState.links.forEach((link) => {
      let newEdge = <any>{
        id: link.id,
        from: link.start,
        to: link.end,
        data: link,
      };

      if (link.type == 'SECONDARY') {
        newEdge.color = { color: 'yellow' };
      }

      this.allEdgeData.update(newEdge);
      
      if (link.properties.status && link.properties.status == 'ACTIVE') {
        this.activeEdgeData.update(newEdge);
      }
    });
  };

  public getLinkedNodes = (id: number): Server[] =>  {
    let retList: Server[] = [];
    
    for (let edge of <any>this.allEdgeData.get({
      filter: function(item) {
        return ((item.from == id) || (item.to == id));
      }
    })) {
            // get the ID of the other node
            let distantNodeID = (edge.from == id) ? edge.to : edge.from;
            let distantNode = (<any>this.nodeData.get(distantNodeID)).data;

            retList.push(distantNode)
    }

    return retList;
  }
}
