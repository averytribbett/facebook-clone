export interface PostModel {
  id: number,
  text: string,
  authorId: number,
  authorFirstName: string,
  authorLastName: string,
  reactCount: number,
  replyCount: number,
  // author profile picture
}