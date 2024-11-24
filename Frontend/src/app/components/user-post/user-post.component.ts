import { Component, Input } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { ReplyModel } from 'src/models/post-model';
import { PostService } from 'src/services/posts-service.service';
import { UserServiceService } from 'src/services/user-service.service';

@Component({
  selector: 'user-post',
  templateUrl: './user-post.component.html',
  styleUrls: ['./user-post.component.css'],
})
export class UserPostComponent {
  @Input() initialLikes!: number;
  @Input() comments!: number;
  @Input() postText!: string;
  @Input() userAvatar!: string;
  @Input() userFirstName!: string;
  @Input() userLastName!: string;
  @Input() initialHasReacted!: boolean;
  @Input() postId!: number;
  @Input() userId!: number;
  public shouldShowCommentText = false;
  public commentText = '';
  public shouldShowComments = false;
  public replyList: ReplyModel[] = [];
  public hasReacted = false;
  public likes = 0;
  public isLoading = false;
  public showReactionTypes = false;

  constructor(
    private userService: UserServiceService,
    private postService: PostService,
    private toasterService: ToastrService,
  ) {
    console.log('temp for linter');
  }

  ngOnInit(): void {
    this.getReplies();
    this.hasReacted = this.initialHasReacted;
    this.likes = this.initialLikes;
  }

  getReplies(): void {
    this.postService.getReplies(this.postId).subscribe((result) => {
      this.replyList = result ?? [];
      this.comments = this.replyList ? this.replyList.length : 0;
    });
  }

  reactToPost(reactionType: string): void {
    // Early exit if loading (prevents double click)
    if (this.isLoading) return;

    this.isLoading = true;

    // Delete the reaction (toggle effect)
    if (this.hasReacted) {
      this.postService.deleteReaction(this.postId, this.userId).subscribe({
        next: (response) => {
          if (response) {
            this.hasReacted = false;
            this.likes--;
          }
          this.isLoading = false;
        },
        error: () => {
          this.isLoading = false;
        },
      });
    } else {
      this.postService
        .addReaction(reactionType, this.postId, this.userId)
        .subscribe({
          next: (response) => {
            if (response) {
              this.hasReacted = true;
              this.likes++;
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
    // this.postService.getReplies(this.postId).subscribe(result => {
    //   console.log('the replies: ', result);
    //   this.replyList = result;
    //   this.comments = this.replyList.length;
    // });
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
    this.commentText = "";
  }

  closeComments(): void {
    this.shouldShowComments = false;
  }

  setShowReactionTypes(value: boolean): void {
    console.log('showReactionTypes: ', value);
    this.showReactionTypes = value;
  }
}
