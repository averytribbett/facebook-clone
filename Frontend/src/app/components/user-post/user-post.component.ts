import { Component, Input } from '@angular/core';
import { PostService } from 'src/services/posts-service.service';

@Component({
  selector: 'user-post',
  templateUrl: './user-post.component.html',
  styleUrls: ['./user-post.component.css'],
})
export class UserPostComponent {
  @Input() likes!: number;
  @Input() comments!: number;
  @Input() postText!: string;
  @Input() userAvatar!: string;
  @Input() userFirstName!: string;
  @Input() userLastName!: string;
  @Input() postId!: number;
  @Input() userId!: number;

  constructor(private postService: PostService) {
    console.log('Temporary for linter');
  }

  reactToPost(reactionType: string): void {
    this.postService.addReaction(reactionType, this.postId, this.userId).subscribe((response) => {
      console.log(response);
    });
  }

  openComments(): void {
    /** @TODO fetch comments from the post here */
    console.log('Toggle open the comments');
  }

  openCommentInput(): void {
    /** @TODO Should open input to leave comment on post */
    console.log('Write a comment here');
  }
}
