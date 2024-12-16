export interface UserModel {
  id?: number;
  firstName: string;
  lastName: string;
  bio: string;
  username: string;
}

export interface DisplayNameUserModel {
  firstName: string;
  lastName: string;
  username: string;
}

/** @TODO remove this and replace with new PostModel in ./models/post-model.ts */
export interface PostModel {
  likes: number;
  comments: number;
  postText: string;
  userAvatar: string;
  userFirstName: string;
  userLastName: string;
}

export const profileEditOptions = ['First Name', 'Last Name', 'Bio'];
export const isAdmin = 'isAdmin';

export enum EditOptions {
  FirstName = 'First Name',
  LastName = 'Last Name',
  Bio = 'Bio',
}
