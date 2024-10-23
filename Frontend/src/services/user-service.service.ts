import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { Observable } from 'rxjs';
import { UserModel } from 'src/models/user-model';

@Injectable({
  providedIn: 'root'
})
export class UserServiceService {
  public loggedInUsername: string = "";

  constructor(private httpClient: HttpClient,
              public userAuth: AuthService) {
                  this.loggedInUsername = localStorage.getItem('myValue') || 'default value';
               }

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

  searchUser(name: string): Observable<UserModel[]> {
    return this.httpClient.get<UserModel[]>("api/user/findUserByName/" + name);
  }

  searchUserByFirstAndLastName(firstName: string, lastName: string): Observable<UserModel> {
    return this.httpClient.get<UserModel>("api/user/findUserByFirstAndLastName/" + firstName + "/" + lastName)
  }

  setValue(value: string) {
    this.loggedInUsername = value;
    localStorage.setItem('myValue', value);
  }

  getValue() {
    return this.loggedInUsername;
  }
}
