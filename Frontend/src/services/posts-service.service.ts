import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { PostModel } from 'src/models/post-model';

@Injectable({
  providedIn: 'root'
})
export class PostService {

  constructor(private httpClient: HttpClient) { }

  getInitialFeedByTime(numOfPosts: number): Observable<PostModel[]>{
    return this.httpClient.get<PostModel[]>("api/posts/initial/" + String(numOfPosts));
  }
}
