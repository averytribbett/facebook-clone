import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { Observable } from 'rxjs';
import { isAdmin, UserModel } from 'src/models/user-model';

@Injectable({
  providedIn: 'root',
})
export class UserServiceService {
  public loggedInUsername = '';
  public loggedInUserIsAdmin = false;

  constructor(
    private httpClient: HttpClient,
    public userAuth: AuthService,
  ) {
    this.loggedInUsername = localStorage.getItem('myValue') || 'default value';
  }

  getAllUsers(): Observable<UserModel[]> {
    return this.httpClient.get<UserModel[]>('api/users');
  }

  getUser(userId: number): Observable<UserModel> {
    return this.httpClient.get<UserModel>('api/user/' + userId);
  }

  getUserByUsername(username: string): Observable<UserModel> {
    return this.httpClient.get<UserModel>('api/username/' + username);
  }

  addNewUser(user: UserModel): Observable<boolean> {
    return this.httpClient.put<boolean>('api/user/addNewUser', user);
  }

  searchUser(name: string): Observable<UserModel[]> {
    return this.httpClient.get<UserModel[]>('api/user/findUserByName/' + name);
  }

  searchUserByFirstAndLastName(
    firstName: string,
    lastName: string,
  ): Observable<UserModel> {
    return this.httpClient.get<UserModel>(
      'api/user/findUserByFirstAndLastName/' + firstName + '/' + lastName,
    );
  }

  editProfileBio(newBio: string, username: string): Observable<boolean> {
    return this.httpClient.patch<boolean>(
      `/api/user/editBio/${newBio}/${username}`,
      null,
    );
  }

  editProfileFirstName(
    newFirstName: string,
    username: string,
  ): Observable<boolean> {
    return this.httpClient.patch<boolean>(
      `/api/user/editFirstName/${newFirstName}/${username}`,
      null,
    );
  }

  editProfileLastName(
    newLastName: string,
    username: string,
  ): Observable<boolean> {
    return this.httpClient.patch<boolean>(
      `/api/user/editLastName/${newLastName}/${username}`,
      null,
    );
  }

  userIsAdmin(userId: number): Observable<boolean> {
    return this.httpClient.get<boolean>(`/api/checkAdmin/${userId}`);
  }

  setValue(value: string) {
    this.loggedInUsername = value;
    localStorage.setItem('myValue', value);
  }

  setBoolean(value: boolean): void {
    localStorage.setItem(isAdmin, value.toString());
  }

  getBoolean(): boolean {
    const storedValue = localStorage.getItem(isAdmin);
    return storedValue === 'true';
  }

  getValue() {
    return this.loggedInUsername;
  }

  uploadProfilePicture(
    formData: FormData,
  ): Observable<{ profileImageUrl: string }> {
    return this.httpClient.post<{ profileImageUrl: string }>(
      '/upload',
      formData,
    );
  }

  getProfilePicture(username: string): Observable<{ imageName: string }> {
    return this.httpClient.get<{ imageName: string }>(
      `/getProfilePicture?username=${username}`,
    );
  }

  getProfilePictureUrl(imageName: string): string {
    return `/uploads/${imageName}`;
  }

  makeUserAdmin(pendingAdminUserId: number, adminUserId: number): Observable<boolean> {
    return this.httpClient.put<boolean>(`/api/makeAdmin/${pendingAdminUserId}/${adminUserId}`, null);
  }

  revokeUserAdmin(pendingRevokeAdminUserId: number, adminUserId: number): Observable<boolean> {
    return this.httpClient.delete<boolean>(`/api/unmakeAdmin/${pendingRevokeAdminUserId}/${adminUserId}`);
  }

  deleteUser(usernameToDelete: string, adminUserId: number): Observable<boolean> {
    return this.httpClient.delete<boolean>(`/api/deleteUserAdmin/${usernameToDelete}/${adminUserId}`);
  }
}
