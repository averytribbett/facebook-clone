import { Component, Input } from "@angular/core";

@Component({
  selector: 'user-post',
  templateUrl: './user-post.component.html',
  styleUrls: ['./user-post.component.css'],
})
export class UserPostComponent {
  @Input() likes!: number
  @Input() comments!: number
  @Input() postText!: string
  @Input() userAvatar!: string
  @Input() userFirstName!: string
  @Input() userLastName!: string

  constructor() {}

  reactToPost(reactionType: string): void {
    /** @TODO create post reaction here */
    console.log(reactionType)
  }

  openComments(): void {
    /** @TODO fetch comments from the post here */
    console.log("Toggle open the comments")
  }

  openCommentInput(): void {
    /** @TODO Should open input to leave comment on post */
    console.log("Write a comment here")
  }
}