import { DOCUMENT } from '@angular/common';
import { Component, EventEmitter, Inject, Input, Output } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';
import { PostModel, UserModel } from 'src/models/user-model';
import { UserServiceService } from 'src/services/user-service.service';
import {MatCardModule} from '@angular/material/card';
import { ActivatedRoute } from '@angular/router';

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
  public fakePosts: PostModel[] = [];
  public loggedInUsername: string = "";
  public profileBeingViewedUsername: string = "";
  public isDeveloper = false;
  public shouldHaveWriteAccess = false;

  constructor (private userService: UserServiceService,
               private route: ActivatedRoute,
               @Inject(DOCUMENT) public document: Document) {}

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
  }

  public showFeedCard(): void {
    this.shouldShowFriendList = false;
    this.shouldShowProfileFeed = true;
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
}
