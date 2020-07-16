export class Server {
  public id: number;
  public labels: string[];
  public properties: any;
}

export class Link {
  public id: number;
  public start: number;
  public end: number;
  public type: string;
  public properties: any;
}

export class NetworkState {
  public links: Link[];
  public splitServers: number[] = [];
}
