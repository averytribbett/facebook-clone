import { Component, EventEmitter, Output } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';
import { UserModel } from 'src/models/user-model';
import { UserServiceService } from 'src/services/user-service.service';
import {MatCardModule} from '@angular/material/card';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent {

  @Output() backToHomeEmitter$ = new EventEmitter<boolean>();
  @Output() logoutEmitter$ = new EventEmitter<boolean>();

  public loginForm: FormGroup = new FormGroup({});
  public selectedUser = {
    id: 0,
    firstName: "Please choose a name...",
    lastName: "",
    username: ""
  } as UserModel;
  public availableUsers: UserModel[] = [];
  public currentSearchUser: number = 0;
  public shouldShowProfileFeed: boolean = true;
  public shouldShowFriendList: boolean = false;
  public fakePosts: string[] = [
    "post one",
    "post two",
    "post three",
    "I am really looking forward to seeing Paddington in Peru in January of 2025!"
  ];

  constructor (private toaster: ToastrService, private userService: UserServiceService) {}

  ngOnInit(): void {
    this.userService.getAllUsers().subscribe(result => {
      this.availableUsers = result;
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

  public backToHome(): void {
    this.backToHomeEmitter$.emit(true);
  }

  public logout(): void {
    this.logoutEmitter$.emit(true);
  }
}
