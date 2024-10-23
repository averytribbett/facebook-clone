import { DOCUMENT } from '@angular/common';
import { Component, EventEmitter, Inject, Input, Output } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';
import { PostModel, UserModel } from 'src/models/user-model';
import { UserServiceService } from 'src/services/user-service.service';
import { ActivatedRoute } from '@angular/router';
import { FriendServiceService } from 'src/services/friend-service.service';
import { FriendModel, FRIENDS, BLOCKED, PENDING } from 'src/models/friend-model';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent {

  @Output() backToHomeEmitter$ = new EventEmitter<boolean>();
  @Output() logoutEmitter$ = new EventEmitter<boolean>();
  @Input() selectedUser = {
    id: 0,
    firstName: "Please choose a name...",
    lastName: "",
    username: ""
  } as UserModel;

  public loginForm: FormGroup = new FormGroup({});
  public availableUsers: UserModel[] = [];
  public currentSearchUser: number = 0;
  public shouldShowProfileFeed: boolean = true;
  public shouldShowFriendList: boolean = false;
  public shouldShowFriendRequests: boolean = false;
  public fakePosts: PostModel[] = [];
  public loggedInUsername: string = "";
  public profileBeingViewedUsername: string = "";
  public isDeveloper = false;
  public shouldHaveWriteAccess = false;
  public userFriends: FriendModel[] = [];
  public shouldShowAddFriendButton: boolean = false;
  public shouldShowDeleteFriendButton: boolean = false;
  public userFriendRequests: FriendModel[] = [];

  constructor (private userService: UserServiceService,
               private route: ActivatedRoute,
               @Inject(DOCUMENT) public document: Document,
              private friendService: FriendServiceService,
              private toaster: ToastrService) {}

  ngOnInit(): void {
    this.userService.userAuth.isAuthenticated$.subscribe(result => {
      if (!result){
        this.userService.userAuth.loginWithRedirect({
          appState: { 
            target: "/home"
          },
          authorizationParams: {
            // Default selected login / signup
            screen_hint: 'signin'
          }    
        });
      } else {
        this.route.params.subscribe(params => {
          this.profileBeingViewedUsername = params['profileUser'];
          this.isDeveloper = params['isDeveloperMode'];
        });
        if (this.profileBeingViewedUsername == this.userService.loggedInUsername) {
          this.shouldHaveWriteAccess = true;
        }
        this.userService.getAllUsers().subscribe(result => {
          this.availableUsers = result;
        });
        this.userService.getUserByUsername(this.profileBeingViewedUsername).subscribe(result => {
          this.selectedUser = result;
          this.fakePosts = [
            {
              likes: 3,
              comments: 5,
              postText: "post one",
              userAvatar: "https://i.redd.it/i-got-bored-so-i-decided-to-draw-a-random-image-on-the-v0-4ig97vv85vjb1.png?width=1280&format=png&auto=webp&s=7177756d1f393b6e093596d06e1ba539f723264b",
              userFirstName: this.selectedUser.firstName,
              userLastName: this.selectedUser.lastName
            },
            {
              likes: 10,
              comments: 25,
              postText: "post two",
              userAvatar: "https://i.redd.it/i-got-bored-so-i-decided-to-draw-a-random-image-on-the-v0-4ig97vv85vjb1.png?width=1280&format=png&auto=webp&s=7177756d1f393b6e093596d06e1ba539f723264b",
              userFirstName: this.selectedUser.firstName,
              userLastName: this.selectedUser.lastName
            },
            {
              likes: 30,
              comments: 75,
              postText: " post three",
              userAvatar: "https://i.redd.it/i-got-bored-so-i-decided-to-draw-a-random-image-on-the-v0-4ig97vv85vjb1.png?width=1280&format=png&auto=webp&s=7177756d1f393b6e093596d06e1ba539f723264b",
              userFirstName: this.selectedUser.firstName,
              userLastName: this.selectedUser.lastName
            },
            {
              likes: 3,
              comments: 5,
              postText: "I am really looking forward to seeing Paddington in Peru in January of 2025!",
              userAvatar: "https://i.redd.it/i-got-bored-so-i-decided-to-draw-a-random-image-on-the-v0-4ig97vv85vjb1.png?width=1280&format=png&auto=webp&s=7177756d1f393b6e093596d06e1ba539f723264b",
              userFirstName: this.selectedUser.firstName,
              userLastName: this.selectedUser.lastName
            }
          ];
        });
        this.getFriendLists();
      }
    });
  }

  public getFriendLists(): void {
    this.friendService.getFriendList(this.profileBeingViewedUsername).subscribe(result => {
      this.userFriends = result ? result.filter(x => x.friendStatus === FRIENDS && x.friendId !== this.profileBeingViewedUsername) : [];
      this.userFriendRequests = result ? result.filter(x => x.friendId === this.profileBeingViewedUsername && x.friendStatus === PENDING) : [];
      if (!this.shouldHaveWriteAccess && result) {
        const loggedInUserIsActiveFriend = this.userFriends.find(x => x.userId === this.profileBeingViewedUsername && x.friendId === this.userService.loggedInUsername) !== undefined;
        const loggedInUserHasRequested = this.userFriendRequests.find(x => x.userId === this.userService.loggedInUsername && x.friendStatus === PENDING); 
        this.shouldShowAddFriendButton = !loggedInUserHasRequested && !loggedInUserIsActiveFriend;
        this.shouldShowDeleteFriendButton = this.userFriends.find(x => x.userId === this.profileBeingViewedUsername && x.friendId === this.userService.loggedInUsername && x.friendStatus === FRIENDS) !== undefined;
      }
      if (!result && !this.shouldHaveWriteAccess) {
        this.shouldShowAddFriendButton = true;
      }
    });
  }

  public changeSelectedUser(user: any): void {
    const newUser = this.availableUsers.find(x => x.username === user.target.value) as UserModel;
    this.selectedUser = newUser;
  }

  // this is searching by ID, not by e-mail
  public searchForUser(): void {
    this.userService.getUser(this.currentSearchUser).subscribe(result => {
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
        returnTo: this.document.location.origin
      }
    });
  }

  public editProfile(): void {
    console.log('LET ME IN!!! LET ME EDIT!!!!!!!!!!!!!!!!!!!!!!');
  }

  public addFriend(): void {
    this.friendService.addFriendRequest(this.userService.loggedInUsername, this.profileBeingViewedUsername).subscribe(result => {
      this.shouldShowAddFriendButton = false;
      this.toaster.show("Friend request sent successfully");
    });
  }

  public acceptFriend(friendToAccept: string): void {
    this.friendService.acceptFriendRequest(friendToAccept, this.userService.loggedInUsername).subscribe(result => {
      this.getFriendLists();
    });
  }

  public deleteFriendRequest(friendToDelete: string): void {
    this.friendService.deleteFriendRequest(friendToDelete, this.userService.loggedInUsername).subscribe(result => {
      this.getFriendLists();
    });
  }

  public deleteFriend(friendToDelete: string): void {
    this.friendService.deleteFriend(friendToDelete, this.userService.loggedInUsername).subscribe(result => {
      this.getFriendLists();
    })
  }
}
