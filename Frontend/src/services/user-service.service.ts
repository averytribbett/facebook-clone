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
}
