import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { FriendModel } from 'src/models/friend-model';
import { UserModel } from 'src/models/user-model';

@Injectable({
  providedIn: 'root'
})
export class FriendServiceService {

  constructor(private httpClient: HttpClient) { }

  public getFriendList(username: string): Observable<UserModel[]> {
    return this.httpClient.get<UserModel[]>("api/friends/findFriendList/" + username);
  }

  public getFriendRequestList(username: string): Observable<UserModel[]> {
    return this.httpClient.get<UserModel[]>("api/friends/findFriendRequestList/" + username);
  }

  public addFriendRequest(requestor: string, requestee: string): Observable<boolean> {
    return this.httpClient.put<boolean>("/api/friends/addPendingFriendship/" + requestor + "/" + requestee, null);
  }

  public acceptFriendRequest(originalRequestor: string, friendWhoAccepted: string): Observable<boolean> {
    return this.httpClient.get<boolean>("/api/friends/acceptFriendship/" + originalRequestor + "/" + friendWhoAccepted);
  }

  public deleteFriendRequest(originalRequestor: string, friendWhoDeleted: string): Observable<boolean> {
    return this.httpClient.delete<boolean>("/api/friends/deleteFriendshipRequest/" + originalRequestor + "/" + friendWhoDeleted)
  }

  public deleteFriend(friendToDelete: string, deleter: string): Observable<boolean> {
    return this.httpClient.delete<boolean>("/api/friends/deleteFriendship/" + friendToDelete + "/" + deleter)
  }
}
