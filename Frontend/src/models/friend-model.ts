export interface FriendModel {
  userId: string;
  friendId: string;
  friendStatus: string;
}

export const FRIENDS = 'friends';
export const PENDING = 'pending';
export const BLOCKED = 'blocked';
