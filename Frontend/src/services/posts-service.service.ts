import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { PostModel, ReplyModel } from 'src/models/post-model';
import { ReactionType } from 'src/app/components/user-post/user-post.component';

@Injectable({
  providedIn: 'root',
})
export class PostService {
  constructor(private httpClient: HttpClient) {}

  getUserPosts(
    userID: number,
    loggedInUserId: number,
  ): Observable<PostModel[]> {
    return this.httpClient.get<PostModel[]>(
      'api/posts/user/' + String(userID) + '/' + String(loggedInUserId),
    );
  }

  getInitialFeedByTime(
    numOfPosts: number,
    loggedInUserId: number,
  ): Observable<PostModel[]> {
    return this.httpClient.get<PostModel[]>(
      'api/posts/initial/' + String(numOfPosts) + '/' + String(loggedInUserId),
    );
  }

  getFeedByTime(
    numOfPosts: number,
    loggedInUserId: number,
  ): Observable<PostModel[]> {
    return this.httpClient.get<PostModel[]>(
      'api/posts/' + String(numOfPosts) + '/' + String(loggedInUserId),
    );
  }

  addPost(userID: number, postText: string): Observable<boolean> {
    return this.httpClient.post<boolean>(
      `api/posts/${userID}/${postText}`,
      null,
    );
  }

  addComment(reply: ReplyModel): Observable<boolean> {
    return this.httpClient.post<boolean>(`api/posts/reply`, reply);
  }

  getReplies(postId: number): Observable<ReplyModel[]> {
    return this.httpClient.get<ReplyModel[]>(
      'api/posts/getAllReplies/' + postId,
    );
  }

  addReaction(
    emoji: ReactionType,
    post_id: number,
    user_id: number,
  ): Observable<boolean> {
    return this.httpClient.post<boolean>(
      `api/reactions/addReaction/${emoji}/${post_id}/${user_id}`,
      null,
    );
  }

  updateReaction(
    emoji: ReactionType,
    post_id: number,
    user_id: number,
  ): Observable<boolean> {
    return this.httpClient.put<boolean>(
      `api/reactions/updateReaction/${emoji}/${post_id}/${user_id}`,
      null,
    );
  }

  deleteReaction(post_id: number, user_id: number): Observable<boolean> {
    return this.httpClient.delete<boolean>(
      `api/reactions/deleteReaction/${post_id}/${user_id}`,
    );
  }

  deletePost(postId: number, adminId: number): Observable<boolean> {
    return this.httpClient.delete<boolean>(`/api/deletePostAdmin/${postId}/${adminId}`);
  }
}
