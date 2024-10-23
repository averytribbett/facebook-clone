import { DOCUMENT } from '@angular/common';
import { Component, Inject } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { environment } from '../../environment';
import { UserServiceService } from 'src/services/user-service.service';
import {MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import { map, Observable, of, startWith } from 'rxjs';
import { FormControl } from '@angular/forms';
import { DisplayNameUserModel, UserModel } from 'src/models/user-model';
import { ActivatedRoute } from '@angular/router';
import { Router } from '@angular/router';
import { PostService } from 'src/services/posts-service.service';
import { PostModel } from 'src/models/post-model';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css','../../styles.css'],
})
export class HomeComponent {
  public currentLoggedInUsername: string = "";
  public currentSearchOptions: DisplayNameUserModel[] = [];
  public currentUser = {
    id: 0,
    firstName: "Please choose a name...",
    lastName: "",
    username: ""
  } as UserModel;
  public isDevelopmentEnvironment = !environment.production;
  public searchTerm: string = "";
  public isLoggedIn: boolean = false;
  public friendSearchList: DisplayNameUserModel[] = [];
  public friendList: DisplayNameUserModel[] = [];
  filteredOptions: Observable<DisplayNameUserModel[]> = new Observable<DisplayNameUserModel[]>();
  myControl = new FormControl('');
  public shouldShowProfilePage = false;
  public feed: PostModel[] = []

  constructor (public auth: AuthService,
              @Inject(DOCUMENT) public document: Document,
              public userService: UserServiceService,
              public router: Router,
              public postService: PostService) {}

  // a double subscribe like this is probably not best practice, but  for now it works
  ngOnInit(): void {
    this.userService.userAuth.user$.subscribe(result => {
      // check if someone is signing up / logging in via auth0
      if (result) {
        this.currentLoggedInUsername = result.name as string;
        this.userService.setValue(result.name as string);
        this.isLoggedIn = true;

        // check for user in database
        this.userService.getUserByUsername(this.currentLoggedInUsername).subscribe(result => {

          // if user does not exist in DB redirect to profile creation page else show home page
          if (!result || !result.username) {
            this.router.navigate(['/create-profile', result.username, false]);
          }
          else {
            this.router.navigate(['/home']);
            this.currentUser = result;
          }
        });
        this.postService.getInitialFeedByTime(5).subscribe(result => {
          if (result.length) {
            this.feed = result;
          }
        })
      }
    });

    this.myControl.valueChanges.subscribe(searchValue => {
      var newVal = searchValue as string;
      if (newVal){
        this.userService.searchUser(newVal).subscribe(result => {
          let userList: DisplayNameUserModel[] = [];
          if(result){
            for (let user of result) {
              userList.push({
                firstName: user.firstName,
                lastName: user.lastName
              });
            }
            this.friendList = userList;
          }
        });
      }
    });

  }

  login(isSignUp: boolean) {
    this.userService.userAuth.loginWithRedirect({
      appState: { 
        target: this.document.location.pathname
      },
      authorizationParams: {
        // Default selected login / signup
        screen_hint: isSignUp ? 'signup' : 'signin'
      }    
    });
  }

  logout() {
    this.userService.userAuth.logout({ 
      logoutParams: {
        returnTo: this.document.location.origin 
      }
    });
  }

  findFriend(searchFirstName: string, searchLastName: string): void {
    this.userService.searchUserByFirstAndLastName(searchFirstName, searchLastName).subscribe(result => {
      this.router.navigate(['/profile', result.username, false]);
    });
  }

}
