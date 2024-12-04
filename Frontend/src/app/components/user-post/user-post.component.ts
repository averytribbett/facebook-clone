import { Component, Input } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { ReplyModel } from 'src/models/post-model';
import { PostService } from 'src/services/posts-service.service';
import { UserServiceService } from 'src/services/user-service.service';

export type ReactionType = 'thumbs_up' | 'thumbs_down' | 'heart';

@Component({
  selector: 'user-post',
  templateUrl: './user-post.component.html',
  styleUrls: ['./user-post.component.css'],
})
export class UserPostComponent {
  @Input() initialReactions!: number;
  @Input() comments!: number;
  @Input() postText!: string;
  @Input() userFirstName!: string;
  @Input() userLastName!: string;
  @Input() authorUsername!: string;
  @Input() initialReactionByUser!: string;
  @Input() postId!: number;
  @Input() userId!: number;
  @Input() loggedInUserId!: number;
  public shouldShowCommentText = false;
  public commentText = '';
  public shouldShowComments = false;
  public replyList: ReplyModel[] = [];
  public reactionByUser = '';
  public reactions = 0;
  public isLoading = false;
  public showReactionTypes = false;
  public profileImageUrl = 'http://localhost:3000/uploads/default.png';

  constructor(
    private userService: UserServiceService,
    private postService: PostService,
    private toasterService: ToastrService,
  ) {
    console.log('temp for linter');
  }

  ngOnInit(): void {
    this.getReplies();
    this.reactionByUser = this.initialReactionByUser;
    this.reactions = this.initialReactions;
    this.userService
      .getProfilePicture(this.authorUsername)
      .subscribe((result) => {
        this.profileImageUrl = this.userService.getProfilePictureUrl(
          result.imageName,
        );
      });
  }

  getReplies(): void {
    this.postService.getReplies(this.postId).subscribe((result) => {
      this.replyList = result ?? [];
      this.comments = this.replyList ? this.replyList.length : 0;
    });
  }

  reactToPost(reactionType: ReactionType, isMainButton = false): void {
    // Early exit if loading (prevents double click)
    if (this.isLoading) return;

    this.isLoading = true;

    // Delete the reaction (auto toggle when clicking the main button and the reaction already exists)
    if (
      this.reactionByUser === reactionType ||
      (isMainButton && this.reactionByUser)
    ) {
      this.postService
        .deleteReaction(this.postId, this.loggedInUserId)
        .subscribe({
          next: (response) => {
            if (response) {
              this.reactionByUser = '';
              this.reactions--;
            }
            this.isLoading = false;
          },
          error: () => {
            this.isLoading = false;
          },
        });
    } else if (this.reactionByUser) {
      // Update the reaction when reaction already exists
      this.postService
        .updateReaction(reactionType, this.postId, this.loggedInUserId)
        .subscribe({
          next: (response) => {
            if (response) {
              this.reactionByUser = reactionType;
            }
            this.isLoading = false;
          },
          error: () => {
            this.isLoading = false;
          },
        });
    } else {
      this.postService
        .addReaction(reactionType, this.postId, this.loggedInUserId)
        .subscribe({
          next: (response) => {
            if (response) {
              this.reactionByUser = reactionType;
              this.reactions++;
            }
            this.isLoading = false;
          },
          error: () => {
            this.isLoading = false;
          },
        });
    }
  }

  openComments(): void {
    this.shouldShowComments = true;
    /** @TODO fetch comments from the post here */
    console.log('Toggle open the comments');
    this.postService.getReplies(this.postId).subscribe((result) => {
      console.log('the replies: ', result);
      this.replyList = result;
      this.comments = this.replyList.length;
    });
  }

  openCommentInput(): void {
    this.shouldShowCommentText = true;
    /** @TODO Should open input to leave comment on post */
    console.log('Write a comment here');
  }

  comment(): void {
    const newComment = {
      postId: this.postId,
      userId: this.userService.loggedInUsername,
      replyText: this.commentText,
    } as ReplyModel;

    if (!newComment.replyText) {
      this.toasterService
        .error(`Do you think we are going to waste precious resources on an empty comment?
                                 Go back to school you moron.`);
    } else {
      this.postService.addComment(newComment).subscribe(() => {
        this.shouldShowCommentText = false;
        this.commentText = '';
        this.getReplies();
      });
    }
  }

  cancel(): void {
    this.shouldShowCommentText = false;
    this.commentText = '';
  }

  closeComments(): void {
    this.shouldShowComments = false;
  }

  setShowReactionTypes(value: boolean): void {
    setTimeout(() => {
      this.showReactionTypes = value;
    }, 100);
  }
}
