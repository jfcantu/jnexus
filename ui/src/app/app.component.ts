import {
  Component,
  ViewChild,
  AfterViewInit,
  ElementRef,
  ComponentFactoryResolver,
  ApplicationRef,
  Injector,
  ComponentRef,
} from '@angular/core';
import { MatSidenav } from '@angular/material/sidenav';
import { ComponentPortal, Portal, DomPortalOutlet } from '@angular/cdk/portal';
import { ServerInfoComponent } from './server-info/server-info.component';
import { RoutingService } from './routing-service/routing.service';
import { SplitServersComponent } from './split-servers/split-servers.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements AfterViewInit {
  @ViewChild(MatSidenav) sidenavBar: MatSidenav;
  @ViewChild('portalElement') portalElement: ElementRef;
  private portalOutlet: DomPortalOutlet;
  private portalComponent: ComponentRef<any>;

  title = 'routing-console';

  public selectedPortal: Portal<any>;

  serverInfoPortal: ComponentPortal<ServerInfoComponent>;
  splitServerPortal: ComponentPortal<SplitServersComponent>;

  constructor(
    public routingService: RoutingService,
    private componentFactoryResolver: ComponentFactoryResolver,
    private applicationRef: ApplicationRef,
    private injector: Injector
  ) {}

  ngAfterViewInit() {
    this.serverInfoPortal = new ComponentPortal(ServerInfoComponent);
    this.splitServerPortal = new ComponentPortal(SplitServersComponent);
    this.portalOutlet = new DomPortalOutlet(
      this.portalElement.nativeElement,
      this.componentFactoryResolver,
      this.applicationRef,
      this.injector
    );
  }

  public showNodeInfo = (selectedNode: any) => {
    if (selectedNode == -1) {
      this.hideSidebar();
      return;
    }

    this.sidenavBar.open();
    if (this.portalOutlet.hasAttached() == false) {
      this.portalComponent = this.serverInfoPortal.attach(this.portalOutlet);
    }
    this.portalComponent.instance.setServer(selectedNode);
  };

  public showSplitServers = () => {
    // Check if the sidebar is already open
    if (this.sidenavBar.opened) {
      // If it's open to the split server display already - close the sidebar
      if (this.splitServerPortal.isAttached) {
        this.hideSidebar();
        return;
      }

      // If it's open to something else, detach that and attach the split server component
      this.portalOutlet.detach();
      this.portalComponent = this.splitServerPortal.attach(this.portalOutlet);
    }
    // if it's not open - open it to the split server component
    else {
      this.sidenavBar.open();
      this.portalOutlet.detach();
      this.portalComponent = this.splitServerPortal.attach(this.portalOutlet);
    }
  };

  public hideSidebar = () => {
    this.sidenavBar.close();
    this.portalOutlet.detach();
  };
}
