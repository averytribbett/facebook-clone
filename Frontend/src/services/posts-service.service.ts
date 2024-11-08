import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { PostModel } from 'src/models/post-model';

@Injectable({
  providedIn: 'root',
})
export class PostService {
  constructor(private httpClient: HttpClient) {}

  getUserPosts(userID: number): Observable<PostModel[]> {
    return this.httpClient.get<PostModel[]>('api/posts/user/' + String(userID));
  }

  getInitialFeedByTime(numOfPosts: number): Observable<PostModel[]> {
    return this.httpClient.get<PostModel[]>(
      'api/posts/initial/' + String(numOfPosts),
    );
  }

  getFeedByTime(numOfPosts: number): Observable<PostModel[]> {
    return this.httpClient.get<PostModel[]>('api/posts/' + String(numOfPosts));
  }

  addPost(userID: number, postText: string): Observable<boolean> {
    return this.httpClient.post<boolean>(
      `api/posts/${userID}/${postText}`,
      null,
    );
  }

  addReaction(emoji: string, post_id: number, user_id: number): Observable<boolean> {
    return this.httpClient.post<boolean>(`api/reactions/addReaction/${emoji}/${post_id}/${user_id}`, null);
  }
}
