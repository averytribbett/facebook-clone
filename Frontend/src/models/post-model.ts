export interface PostModel {
  id: number;
  text: string;
  authorId: number;
  authorFirstName: string;
  authorLastName: string;
  reactionCount: number;
  replyCount: number;
  reactionByUser: string;
  // author profile picture
}

export interface ReplyModel {
  postId: number;
  userId: string;
  replyText: string;
  replierFirstName?: string;
  replierLastName?: string;
  replierUsername: string;
}
