import {
  Component,
  ElementRef,
  ViewChild,
  AfterViewInit,
  Output,
  EventEmitter,
} from '@angular/core';
import { DataSet, Node, Edge } from 'vis';
import { MapService } from '../map-service/map.service';

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css'],
})
export class MapComponent implements AfterViewInit {
  @ViewChild('mapContainer') el: ElementRef;
  @Output()
  nodeSelected: EventEmitter<number> = new EventEmitter<number>();

  private nodeData: DataSet<Node> = new DataSet<Node>();
  private edgeData: DataSet<Edge> = new DataSet<Edge>();

  constructor(private mapService: MapService) {}

  ngAfterViewInit() {
    console.log('rendering');
    this.mapService.render(this.el.nativeElement);
    this.mapService.networkInstance.on('selectNode', this.nodeSelectedEvent);
    this.mapService.networkInstance.on('deselectNode', this.nodeDeselectedEvent);
  }

  private nodeSelectedEvent = (node: any) => {
    this.nodeSelected.emit(<number>node.nodes[0]);
  };

  private nodeDeselectedEvent = (node: any) => {
    this.nodeSelected.emit(-1);
  };
}
