import { Component, OnInit, InjectionToken, Input } from '@angular/core';
import { Server } from '../routing-service/entities';
import { MapService } from '../map-service/map.service';
import { RoutingService } from '../routing-service/routing.service';

@Component({
  selector: 'app-server-info',
  templateUrl: './server-info.component.html',
  styleUrls: ['./server-info.component.css'],
})
export class ServerInfoComponent implements OnInit {
  public server: Server;
  public mapLinks: Server[] = [];
  public activeLinks: Server[] = [];

  constructor(private mapService: MapService, private routingService: RoutingService) {}

  ngOnInit() {}

  public setServer(id: number) {
    this.server = this.routingService.serverData.get(id);
    this.mapLinks = [];
    this.activeLinks = [];

    // Iterate through edges that have this node as either a source or destination
    for (let edge of <any>this.mapService.allEdgeData.get({
      filter: function(item) {
        return ((item.from == id) || (item.to == id));
      }
    })) {
      // get the ID of the other node
      let distantNodeID = (edge.from == id) ? edge.to : edge.from;
      let distantNode = (<any>this.mapService.nodeData.get(distantNodeID)).data;
      
      console.log(distantNode);

      if (edge.data.properties.status == "ACTIVE") {
        this.activeLinks.push(distantNode);
      }

      if (edge.data.type == "PRIMARY" && distantNode.labels.includes("hub")) {
        this.mapLinks.push(distantNode);
      }
    }
  }
}
