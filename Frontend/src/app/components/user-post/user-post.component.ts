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
  @Input() likes!: number;
  @Input() comments!: number;
  @Input() postText!: string;
  @Input() userAvatar!: string;
  @Input() userFirstName!: string;
  @Input() userLastName!: string;
  @Input() hasReacted!: boolean;
  @Input() postId!: number;
  @Input() userId!: number;
  public shouldShowCommentText = false;
  public commentText = '';
  public shouldShowComments = false;
  public replyList: ReplyModel[] = [];

  constructor(
    private userService: UserServiceService,
    private postService: PostService,
    private toasterService: ToastrService,
  ) {
    console.log('temp for linter');
  }

  ngOnInit(): void {
    this.getReplies();
  }

  getReplies(): void {
    this.postService.getReplies(this.postId).subscribe((result) => {
      console.log('the replies: ', result);
      this.replyList = result ?? [];
      this.comments = this.replyList ? this.replyList.length : 0;
    });
  }

  reactToPost(reactionType: string): void {
    this.postService.addReaction(reactionType, this.postId, this.userId).subscribe((response) => {
      console.log(response);
    });
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
        console.log('comment in backend hopefully');
      });
    }

    console.log('sending comment to backend');
  }

  cancel(): void {
    this.shouldShowCommentText = false;
    console.log('other cancel functionality such as clearing the current text');
  }

  closeComments(): void {
    this.shouldShowComments = false;
  }
}
