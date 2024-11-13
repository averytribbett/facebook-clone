import { DOCUMENT } from '@angular/common';
import { Component, EventEmitter, Inject, Input, Output } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';
import { UserModel } from 'src/models/user-model';
import { PostModel } from 'src/models/post-model';
import { UserServiceService } from 'src/services/user-service.service';
import { ActivatedRoute } from '@angular/router';
import { FriendServiceService } from 'src/services/friend-service.service';
import { PostService } from 'src/services/posts-service.service';
import { forkJoin } from 'rxjs';
@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css'],
})
export class UserProfileComponent {
  @Output() backToHomeEmitter$ = new EventEmitter<boolean>();
  @Output() logoutEmitter$ = new EventEmitter<boolean>();
  @Input() selectedUser = {
    id: 0,
    firstName: 'Please choose a name...',
    lastName: '',
    username: '',
  } as UserModel;

  public loginForm: FormGroup = new FormGroup({});
  public availableUsers: UserModel[] = [];
  public currentSearchUser = 0;
  public shouldShowProfileFeed = true;
  public shouldShowFriendList = false;
  public shouldShowFriendRequests = false;
  public userPosts: PostModel[] = [];
  public loggedInUsername = '';
  public loggedInUserId = 0;
  public profileBeingViewedUsername = '';
  public isDeveloper = false;
  public shouldHaveWriteAccess = false;
  public userFriends: UserModel[] = [];
  public shouldShowAddFriendButton = false;
  public shouldShowDeleteFriendButton = false;
  public userFriendRequests: UserModel[] = [];

  constructor(
    private userService: UserServiceService,
    private route: ActivatedRoute,
    @Inject(DOCUMENT) public document: Document,
    private friendService: FriendServiceService,
    private toaster: ToastrService,
    private postService: PostService,
  ) {}

  ngOnInit(): void {
    this.userService.userAuth.isAuthenticated$.subscribe((result) => {
      if (!result) {
        this.userService.userAuth.loginWithRedirect({
          appState: {
            target: '/home',
          },
          authorizationParams: {
            // Default selected login / signup
            screen_hint: 'signin',
          },
        });
      } else {
        this.route.params.subscribe((params) => {
          this.profileBeingViewedUsername = params['profileUser'];
          this.isDeveloper = params['isDeveloperMode'];
        });
        if (
          this.profileBeingViewedUsername == this.userService.loggedInUsername
        ) {
          this.shouldHaveWriteAccess = true;
        }
        this.userService.getAllUsers().subscribe((result) => {
          this.availableUsers = result;
        });
        this.userService
          .getUserByUsername(this.userService.loggedInUsername)
          .subscribe((result) => {
            if (result?.id !== undefined) {
              this.loggedInUserId = result.id;
            } else {
              this.loggedInUserId = 0;
            }
          });
        this.userService
          .getUserByUsername(this.profileBeingViewedUsername)
          .subscribe((result) => {
            this.selectedUser = result;
            if (this.selectedUser.id) {
              this.postService
                .getUserPosts(this.selectedUser.id, this.loggedInUserId)
                .subscribe((result) => {
                  if (result && result.length) {
                    this.userPosts = result;
                  } else {
                    this.userPosts = [];
                  }
                });
            }
          });
        this.retrieveAndUpdateFriendLists();
      }
    });
  }

  public retrieveAndUpdateFriendLists(): void {
    const friendListObservable = this.friendService.getFriendList(
      this.profileBeingViewedUsername,
    );
    const friendRequestListObservable = this.friendService.getFriendRequestList(
      this.profileBeingViewedUsername,
    );

    forkJoin([friendListObservable, friendRequestListObservable]).subscribe(
      (result) => {
        console.log('first result: ', result[0]);
        console.log('second result: ', result[1]);
        this.getFriendLists(result[0], result[1]);
      },
    );
  }
  public getFriendLists(
    friendList: UserModel[],
    friendRequestList: UserModel[],
  ) {
    this.userFriends = friendList
      ? friendList.filter((x) => x.username !== this.profileBeingViewedUsername)
      : [];
    this.userFriendRequests = friendRequestList ?? [];
    if (!this.shouldHaveWriteAccess) {
      const loggedInUserIsActiveFriend =
        this.userFriends.find(
          (x) => x.username === this.userService.loggedInUsername,
        ) !== undefined;
      const loggedInUserHasRequested =
        this.userFriendRequests.find(
          (x) => x.username === this.userService.loggedInUsername,
        ) !== undefined;
      this.shouldShowAddFriendButton =
        !loggedInUserHasRequested && !loggedInUserIsActiveFriend;
      this.shouldShowDeleteFriendButton = loggedInUserIsActiveFriend;
    }
    if (!this.userFriendRequests && !this.shouldHaveWriteAccess) {
      this.shouldShowAddFriendButton = true;
    }
  }

  public changeSelectedUser(user: any): void {
    const newUser = this.availableUsers.find(
      (x) => x.username === user.target.value,
    ) as UserModel;
    this.selectedUser = newUser;
  }

  // this is searching by ID, not by e-mail
  public searchForUser(): void {
    this.userService.getUser(this.currentSearchUser).subscribe((result) => {
      this.selectedUser = result;
    });
  }

  public showFriendCard(): void {
    this.shouldShowFriendList = true;
    this.shouldShowProfileFeed = false;
    this.shouldShowFriendRequests = false;
  }

  public showFeedCard(): void {
    this.shouldShowFriendList = false;
    this.shouldShowFriendRequests = false;
    this.shouldShowProfileFeed = true;
  }

  public showFriendRequestCard(): void {
    this.shouldShowFriendList = false;
    this.shouldShowFriendRequests = true;
    this.shouldShowProfileFeed = false;
  }

  public logout(): void {
    this.userService.userAuth.logout({
      logoutParams: {
        returnTo: this.document.location.origin,
      },
    });
  }

  public editProfile(): void {
    console.log('LET ME IN!!! LET ME EDIT!!!!!!!!!!!!!!!!!!!!!!');
  }

  public addFriend(): void {
    this.friendService
      .addFriendRequest(
        this.userService.loggedInUsername,
        this.profileBeingViewedUsername,
      )
      .subscribe(() => {
        this.shouldShowAddFriendButton = false;
        this.toaster.show('Friend request sent successfully');
      });
  }

  public acceptFriend(friendToAccept: string): void {
    this.friendService
      .acceptFriendRequest(friendToAccept, this.userService.loggedInUsername)
      .subscribe(() => {
        this.retrieveAndUpdateFriendLists();
      });
  }

  public deleteFriendRequest(friendToDelete: string): void {
    this.friendService
      .deleteFriendRequest(friendToDelete, this.userService.loggedInUsername)
      .subscribe(() => {
        this.retrieveAndUpdateFriendLists();
      });
  }

  public deleteFriend(friendToDelete: string): void {
    this.friendService
      .deleteFriend(friendToDelete, this.userService.loggedInUsername)
      .subscribe(() => {
        this.retrieveAndUpdateFriendLists();
      });
  }
}
