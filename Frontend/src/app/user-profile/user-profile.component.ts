import { DOCUMENT } from '@angular/common';
import {
  Component,
  EventEmitter,
  Inject,
  Input,
  Output,
  QueryList,
  ViewChild,
  ViewChildren,
} from '@angular/core';
import { FormGroup } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';
import {
  EditOptions,
  profileEditOptions,
  UserModel,
} from 'src/models/user-model';
import { PostModel } from 'src/models/post-model';
import { UserServiceService } from 'src/services/user-service.service';
import { ActivatedRoute, Router } from '@angular/router';
import { FriendServiceService } from 'src/services/friend-service.service';
import { PostService } from 'src/services/posts-service.service';
import { forkJoin } from 'rxjs';
import {
  MatChipEvent,
  MatChipListbox,
  MatChipOption,
  MatChipSelectionChange,
} from '@angular/material/chips';
import { AuthService } from '@auth0/auth0-angular';
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
  @ViewChild('chipList') chipList!: MatChipListbox;
  @ViewChildren(MatChipOption) chipOptions!: QueryList<MatChipOption>;

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
  public isEditing = false;
  public isEditingOptionSelected = false;
  public isSubmitEditButtonDisabled = false;
  public newProfileInfo = '';
  public editOptions = profileEditOptions;
  public currentAttributeBeingEdited = '';
  public file: File | null = null;
  public previewUrl: string | null = null;
  public username: string | null = null;
  public profileImageUrl = 'http://localhost:3000/uploads/default.png';
  public selectedFile: File | null = null;
  public loggedInUserIsAdmin = true;
  public profileBeingViewedIsAdmin = false;

  constructor(
    private userService: UserServiceService,
    private route: ActivatedRoute,
    @Inject(DOCUMENT) public document: Document,
    private friendService: FriendServiceService,
    private toaster: ToastrService,
    private postService: PostService,
    public auth: AuthService,
    private router: Router,
  ) { }

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
        this.loggedInUserIsAdmin = this.userService.getBoolean();
        this.route.params.subscribe((params) => {
          this.profileBeingViewedUsername = params['profileUser'];
          this.isDeveloper = params['isDeveloperMode'];
        });
        if (
          this.profileBeingViewedUsername == this.userService.loggedInUsername
        ) {
          this.shouldHaveWriteAccess = true;
        }
        const loggedInUserIdObservable = this.userService.getUserByUsername(
          this.userService.loggedInUsername,
        );
        const profileBeingViewedObservable = this.userService.getUserByUsername(
          this.profileBeingViewedUsername,
        );

        forkJoin([
          loggedInUserIdObservable,
          profileBeingViewedObservable,
        ]).subscribe((result) => {
          if (result[0]?.id !== undefined) {
            this.loggedInUserId = result[0].id;
          } else {
            this.loggedInUserId = 0;
          }
          this.selectedUser = result[1];
          this.userService.userIsAdmin(this.selectedUser.id as number).subscribe(result => {
            this.profileBeingViewedIsAdmin = result;
          });
          this.profileBeingViewedUsername = this.selectedUser.username;
          this.loadProfilePicture(this.profileBeingViewedUsername);
          this.getUserPosts();
          this.retrieveAndUpdateFriendLists();
        });
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
        this.getFriendLists(result[0], result[1]);
      },
    );
  }

  public getUserPosts(): void {
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
    this.isEditing = true;
  }

  public cancelEditing(): void {
    this.isEditing = false;
    this.clearEditing(false);
  }

  public setProfileEditingParameters(
    attribute: string,
    isSelected: boolean,
  ): void {
    this.isEditingOptionSelected = isSelected;
    this.currentAttributeBeingEdited = attribute ?? '';
    if (!this.isEditingOptionSelected) {
      this.newProfileInfo = '';
    }
  }

  public editProfileAttribute(): void {
    this.isSubmitEditButtonDisabled = true;
    if (this.currentAttributeBeingEdited === EditOptions.Bio) {
      this.userService
        .editProfileBio(this.newProfileInfo, this.profileBeingViewedUsername)
        .subscribe(() => {
          this.selectedUser.bio = this.newProfileInfo;
          this.clearEditing();
        });
    } else if (this.currentAttributeBeingEdited === EditOptions.FirstName) {
      this.userService
        .editProfileFirstName(
          this.newProfileInfo,
          this.userService.loggedInUsername,
        )
        .subscribe(() => {
          this.selectedUser.firstName = this.newProfileInfo;
          this.clearEditing();
        });
    } else if (this.currentAttributeBeingEdited === EditOptions.LastName) {
      this.userService
        .editProfileLastName(
          this.newProfileInfo,
          this.userService.loggedInUsername,
        )
        .subscribe(() => {
          this.selectedUser.lastName = this.newProfileInfo;
          this.clearEditing();
        });
    } else {
      this.toaster.error(
        "I'm not sure how you did it, but whatever option you selected doesn't exist... try again, pal",
      );
    }
  }

  public clearEditing(requiresMessage = true): void {
    if (requiresMessage) {
      this.toaster.show(
        `${this.currentAttributeBeingEdited} updated successfully`,
      );
    }
    this.newProfileInfo = '';
    this.isEditingOptionSelected = false;
    this.isSubmitEditButtonDisabled = false;
    const chipOption = this.chipOptions
      .toArray()
      .find((chip) => chip.value === this.currentAttributeBeingEdited);
    if (chipOption) {
      chipOption.deselect();
    }
    if (this.currentAttributeBeingEdited !== EditOptions.Bio) {
      this.getUserPosts();
    }
    this.currentAttributeBeingEdited = '';
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

  uploaded(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
      this.selectedFile = input.files[0];
      const reader = new FileReader();
      reader.onload = (e: any) => {
        this.profileImageUrl = e.target.result;
      };
      reader.readAsDataURL(this.selectedFile);
    }
  }

  uploadFile() {
    if (this.selectedFile) {
      const formData = new FormData();
      formData.append('file', this.selectedFile);
      formData.append('username', this.userService.loggedInUsername);

      this.userService.uploadProfilePicture(formData).subscribe({
        next: (response) => {
          this.profileImageUrl = response.profileImageUrl;
          this.selectedFile = null;
        },
        error: (error) => {
          console.error('Error uploading profile picture:', error);
        },
      });
    }
    location.reload();
  }

  loadProfilePicture(username: string): void {
    this.userService.getProfilePicture(username).subscribe(
      (response) => {
        this.profileImageUrl = this.userService.getProfilePictureUrl(
          response.imageName,
        );
      },
      (error) => {
        console.error('Error loading profile picture', error);
      },
    );
  }

  triggerFileInput() {
    if (this.shouldHaveWriteAccess) {
      const fileInput = document.querySelector(
        'input[type="file"]',
      ) as HTMLElement;
      fileInput.click();
    }
  }

  handleFileSelection(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
      this.selectedFile = input.files[0];
      const reader = new FileReader();
      reader.onload = (e: any) => {
        this.profileImageUrl = e.target.result;
      };
      reader.readAsDataURL(this.selectedFile);
      this.uploadFile();
    }
  }

  public makeUserAdmin(): void {
    this.userService.makeUserAdmin(this.selectedUser.id as number, this.loggedInUserId).subscribe(result => {
      this.toaster.show('User successfully made an admin!');
      this.profileBeingViewedIsAdmin = true;
    });
  }

  public revokeUserAdmin(): void {
    this.userService.revokeUserAdmin(this.selectedUser.id as number, this.loggedInUserId as number).subscribe(result => {
      this.toaster.show('User admin privileges successfully revoked!');
      this.profileBeingViewedIsAdmin = false;
    });
  }

  public deleteUser(): void {
    this.toaster.show('User deletion in progress, please be patient, this may take several moments')
    this.userService.deleteUser(this.selectedUser.username, this.loggedInUserId as number).subscribe(() => {
      this.toaster.show('User deleted');
      this.router.navigate(['/home']);
    });
  }
}
