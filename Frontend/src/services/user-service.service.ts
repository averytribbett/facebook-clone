import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { UserModel } from 'src/models/user-model';

@Injectable({
  providedIn: 'root'
})
export class UserServiceService {

  constructor(private httpClient: HttpClient) { }

  getAllUsers(): Observable<UserModel[]>{
    return this.httpClient.get<UserModel[]>("api/users");
  }

  getUser(userId: number): Observable<UserModel>{
    return this.httpClient.get<UserModel>("api/user/" + userId);
  }

  getUserByUsername(username: String): Observable<UserModel>{
    return this.httpClient.get<UserModel>("api/username/" + username);
  }

  addNewUser(user: UserModel): Observable<boolean>{
    return this.httpClient.put<boolean>("api/user/addNewUser", user);
  }
}
