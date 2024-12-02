/* eslint-disable prettier/prettier */
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { UserModel } from 'src/models/user-model';

@Injectable({
  providedIn: 'root',
})
export class FriendServiceService {
  constructor(private httpClient: HttpClient) {}

  public getFriendList(username: string): Observable<UserModel[]> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.get<UserModel[]>(
      `${baseUrl}/api/friends/findFriendList/` + username,
    );
  }

  public getFriendRequestList(username: string): Observable<UserModel[]> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.get<UserModel[]>(
      `${baseUrl}/api/friends/findFriendRequestList/` + username,
    );
  }

  public addFriendRequest(
    requestor: string,
    requestee: string,
  ): Observable<boolean> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.put<boolean>(
      `${baseUrl}/api/friends/addPendingFriendship/` + requestor + '/' + requestee,
      null,
    );
  }

  public acceptFriendRequest(
    originalRequestor: string,
    friendWhoAccepted: string,
  ): Observable<boolean> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.get<boolean>(
      `${baseUrl}/api/friends/acceptFriendship/` +
        originalRequestor +
        '/' +
        friendWhoAccepted,
    );
  }

  public deleteFriendRequest(
    originalRequestor: string,
    friendWhoDeleted: string,
  ): Observable<boolean> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.delete<boolean>(
      `${baseUrl}/api/friends/deleteFriendshipRequest/` +
        originalRequestor +
        '/' +
        friendWhoDeleted,
    );
  }

  public deleteFriend(
    friendToDelete: string,
    deleter: string,
  ): Observable<boolean> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.delete<boolean>(
      `${baseUrl}/api/friends/deleteFriendship/` + friendToDelete + '/' + deleter,
    );
  }
}
