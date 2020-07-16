import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Server, Link, NetworkState } from './entities';
import { BehaviorSubject, Observable } from 'rxjs';
import { Network } from 'vis';

@Injectable({
  providedIn: 'root',
})
export class RoutingService {
  private _networkState: BehaviorSubject<NetworkState> = new BehaviorSubject<
    NetworkState
  >(null);
  private _splitServers: BehaviorSubject<number[]> = new BehaviorSubject<
    number[]
  >([]);
  private _serverData: Map<number, Server> = new Map<number, Server>();

  constructor(private http: HttpClient) {
    this.getServerData();
  }

  public get networkState(): Observable<NetworkState> {
    return this._networkState.asObservable();
  }

  public get splitServers(): Observable<number[]> {
    return this._splitServers.asObservable();
  }

  public get serverData(): Map<number, Server> {
    return this._serverData;
  }

  public refreshStatus = async () => {
    let newState = new NetworkState();

    // Retrieve all links
    const linkRequest = this.http
      .get<any>(`http://pyhakon.net:7443/links/all`, {
        responseType: 'json',
      })
      .toPromise();

    try {
      await linkRequest;
    } catch (e) {
      return linkRequest;
    }

    linkRequest.then((data) => {
      newState.links = data;
    });

    const splitRequest = this.http
      .get<any>(`http://pyhakon.net:7443/servers/split`, {
        responseType: 'json',
      })
      .toPromise();

    try {
      await splitRequest;
    } catch (e) {
      return splitRequest;
    }

    splitRequest.then((data) => {
      newState.splitServers = <number[]>data;
    });

    // Update observables
    this._networkState.next(newState);
    this._splitServers.next(newState.splitServers);
  };

  private getServerData = async () => {
    // Retrieve server info
    const dataRequest = this.http
      .get<any>(`http://pyhakon.net:7443/servers`, {
        responseType: 'json',
      })
      .toPromise();

    try {
      await dataRequest;
    } catch (e) {
      return dataRequest;
    }

    dataRequest.then((data) => {
      for (let v of <Server[]>data) {
        this._serverData.set(v.id, v);
      }
    });
  }
}
