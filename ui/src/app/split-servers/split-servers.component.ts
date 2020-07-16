import { Component, OnInit } from '@angular/core';
import { NetworkState, Server } from '../routing-service/entities';
import { Observable } from 'rxjs';
import { RoutingService } from '../routing-service/routing.service';
import { MapService } from '../map-service/map.service';

@Component({
  selector: 'app-split-servers',
  templateUrl: './split-servers.component.html',
  styleUrls: ['./split-servers.component.css']
})
export class SplitServersComponent implements OnInit {
  public networkState: Observable<NetworkState>;
  public splitServers: Server[];

  constructor(
    public routingService: RoutingService,
    public mapService: MapService
  ) {
    let newList: Server[] = [];
    this.routingService.networkState.subscribe((newState) => {
      for (let server of newState.splitServers) {
        let mapNode: any = this.mapService.nodeData.get(server);
        newList.push(<Server>mapNode.data);
      }
      
      this.splitServers = newList;
      // testing
      /*
      this.splitServers = [
        (<any>this.mapService.nodeData.get(5)).data,
        (<any>this.mapService.nodeData.get(16)).data
      ]; */
    })
  }

  ngOnInit() {
  }

  public nope() {
    alert("doesn't work yet lol");
  }
}
