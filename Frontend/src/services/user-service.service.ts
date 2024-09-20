import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { UserModel } from 'src/models/user-model';

@Injectable({
  providedIn: 'root'
})
export class UserServiceService {
  httpClient: any;

  constructor() { }

  getAllUsers(): Observable<UserModel[]>{
    return this.httpClient.sendUserProfile("/users");
  }
}
