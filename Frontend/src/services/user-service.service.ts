/* eslint-disable prettier/prettier */
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { Observable } from 'rxjs';
import { UserModel } from 'src/models/user-model';

@Injectable({
  providedIn: 'root',
})
export class UserServiceService {
  public loggedInUsername = '';

  constructor(
    private httpClient: HttpClient,
    public userAuth: AuthService,
  ) {
    this.loggedInUsername = localStorage.getItem('myValue') || 'default value';
  }

  getAllUsers(): Observable<UserModel[]> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.get<UserModel[]>(`${baseUrl}/api/users`);
  }

  getUser(userId: number): Observable<UserModel> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.get<UserModel>(`${baseUrl}/api/user/` + userId);
  }

  getUserByUsername(username: string): Observable<UserModel> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.get<UserModel>(`${baseUrl}/api/username/` + username);
  }

  addNewUser(user: UserModel): Observable<boolean> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.put<boolean>(`${baseUrl}/api/user/addNewUser`, user);
  }

  searchUser(name: string): Observable<UserModel[]> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.get<UserModel[]>(`${baseUrl}/api/user/findUserByName/` + name);
  }

  searchUserByFirstAndLastName(
    firstName: string,
    lastName: string,
  ): Observable<UserModel> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.get<UserModel>(
      `${baseUrl}/api/user/findUserByFirstAndLastName/` + firstName + '/' + lastName,
    );
  }

  editProfileBio(newBio: string, username: string): Observable<boolean> {
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.patch<boolean>(
      `${baseUrl}/api/user/editBio/${newBio}/${username}`,
      null,
    );
  }

  editProfileFirstName(
    newFirstName: string,
    username: string,
  ): Observable<boolean> {
    console.log('we are editing the first name');
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.patch<boolean>(
      `${baseUrl}/api/user/editFirstName/${newFirstName}/${username}`,
      null,
    );
  }

  editProfileLastName(
    newLastName: string,
    username: string,
  ): Observable<boolean> {
    console.log('we are editing the last name');
    const baseUrl = 'http://71.89.73.28:3000';
    return this.httpClient.patch<boolean>(
      `${baseUrl}/api/user/editLastName/${newLastName}/${username}`,
      null,
    );
  }

  setValue(value: string) {
    this.loggedInUsername = value;
    localStorage.setItem('myValue', value);
  }

  getValue() {
    return this.loggedInUsername;
  }
}
